package test

import (
	"bytes"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func prepareViperConfig() {
	viper.Set("default", "default_app")
}

func GetRootCmdWithSubCommands(c ...*cobra.Command) *cobra.Command {
	prepareViperConfig()

	rootCmd := &cobra.Command{}
	rootCmd.PersistentFlags().String("app", "", "app name")
	rootCmd.AddCommand(c...)
	rootCmd.SetIn(&bytes.Buffer{})
	rootCmd.SetOut(&bytes.Buffer{})
	rootCmd.SetErr(&bytes.Buffer{})

	return rootCmd
}
