package config

import (
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/input-stream/cli/stream/input/v1beta1"
)

const (
	configDir      = "istream"
	defaultApiHost = "api.input.stream"
	defaultApiPort = "443"
)

type Config struct {
	// conn is the client connection, if it has been initialized.
	conn *grpc.ClientConn

	// Viper uses `yaml` for serializing the object into a file.
	// And then uses `mapstructure` to deserialize into an actual Config object.

	// APIKey is sent with each request
	APIKey string `yaml:"api-key" mapstructure:"INPUTSTREAM_API_KEY"`
	// APIHost is the hostname for the API Server
	APIHost string `yaml:"api-host" mapstructure:"INPUTSTREAM_API_HOST"`
	// APIPort is the hostname for the API Server
	APIPort string `yaml:"api-port" mapstructure:"INPUTSTREAM_API_PORT"`
}

func (c *Config) createGrpcConnection() (*grpc.ClientConn, error) {
	pool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("initializing system x509 cert pool: %w", err)
	}
	creds := credentials.NewClientTLSFromCert(pool, "")

	target := fmt.Sprintf("%s:%s", c.APIHost, c.APIPort)
	conn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, fmt.Errorf("dialing api server: %w", err)
	}
	c.conn = conn
	return conn, nil
}

func (c *Config) GetUsersClient(cmd *cobra.Command) (v1beta1.UsersClient, error) {
	conn, err := c.createGrpcConnection()
	if err != nil {
		return nil, err
	}
	return v1beta1.NewUsersClient(conn), nil
}

func (c *Config) GetInputsClient(cmd *cobra.Command) (v1beta1.InputsClient, error) {
	conn, err := c.createGrpcConnection()
	if err != nil {
		return nil, err
	}
	return v1beta1.NewInputsClient(conn), nil
}

func GetConfig(cmd *cobra.Command) *Config {
	cfg := &Config{
		APIHost: defaultApiHost,
		APIPort: defaultApiPort,
	}
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
