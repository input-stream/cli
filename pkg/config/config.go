package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configDir = "istream"
)

type Config struct {
	// Viper uses `yaml` for serializing the object into a file.
	// And then uses `mapstructure` to deserialize into an actual Config object.

	// APIKey is sent with each request
	APIKey string `yaml:"api-key" mapstructure:"INPUTSTREAM_API_KEY"`
}

func GetConfig(cmd *cobra.Command) *Config {
	cfg := &Config{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		cmd.PrintErr("malformed configuration: " + err.Error())
		os.Exit(1)
	}
	return cfg
}

func GetInitConfig(cmd *cobra.Command, cfgPath *string) func() {
	return func() {
		var configPath string

		if *cfgPath != "" {
			// Use config file from the flag.
			configPath = *cfgPath
		} else {
			// Otherwise use UserConfigDir
			dir, err := os.UserConfigDir()
			if err != nil {
				cmd.PrintErr(err)
				os.Exit(1)
			}

			configPath = filepath.Join(dir, configDir, "config.yaml")
		}

		log.Println("config file path:", configPath)
		viper.SetConfigFile(configPath)

		err := viper.ReadInConfig()
		if err != nil && os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Dir(configPath), 0o755)
			if err != nil {
				cmd.PrintErr(err)
				os.Exit(1)
			}

			f, err := os.Create(configPath)
			if err != nil {
				cmd.PrintErr(err)
				os.Exit(1)
			}

			f.Close()
		}
		if err != nil {
			cmd.PrintErr(err)
			os.Exit(1)
		}
	}
}
