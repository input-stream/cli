package login

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"

	"github.com/input-stream/cli/build/stack/auth/v1beta1"
	"github.com/input-stream/cli/pkg/config"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with server",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.GetConfig(cmd)
			client, err := cfg.GetAuthClient(cmd)
			if err != nil {
				return err
			}
			stream, err := client.DeviceLogin(context.Background(), &v1beta1.DeviceLoginRequest{
				DeviceName: "github.com-inputstream-cli",
				// ApiToken:   cfg.APIKey,
			})
			if err != nil {
				return err
			}

			done := make(chan error)
			go func() {
				didOpenOauthURL := false
				for {
					resp, err := stream.Recv()
					if err == io.EOF {
						done <- nil
						return
					}
					if err != nil {
						log.Fatalf("can not receive %v", err)
					}
					if !resp.Completed && resp.OauthUrl != "" && !didOpenOauthURL {
						log.Printf("open: %s", resp.OauthUrl)
						didOpenOauthURL = true
						if err := open.Start(resp.OauthUrl); err != nil {
							done <- err
							return
						}
					} else if resp.Completed && resp.AccessToken != "" {
						cfg.AccessToken = resp.AccessToken
						log.Printf("access token: %s", resp.AccessToken)
						if err := cfg.Write(); err != nil {
							done <- fmt.Errorf("failed to write access token in config: %w", err)
						}
						done <- nil
					}
				}
			}()

			select {
			case err = <-done:
				if err != nil {
					return err
				}
				log.Println("Done.")
			case <-time.After(30 * time.Second):
				return fmt.Errorf("timeout")
			}
			return nil
		},
	}
	return cmd
}
