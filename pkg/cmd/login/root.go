package login

import (
	"log"

	"github.com/input-stream/cli/pkg/config"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with server",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := config.GetConfig(cmd)
			log.Println("login:", c.APIKey)
			return nil
		},
	}
	return cmd
}
