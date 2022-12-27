package list

import (
	"fmt"
	"log"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/input-stream/cli/build/stack/inputstream/v1beta1"
	"github.com/input-stream/cli/pkg/config"
)

func NewCmds() []*cobra.Command {
	return []*cobra.Command{
		listInputsCmd(),
	}
}

func listInputsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list --type [type]",
		Short: "List inputs",
		Long: heredoc.Doc(`
			Stream will not block any file types from uploading, however, different
			clients may handle different types differently or not at all.
			You can set a more restrictive list for your application if needed.
			The maximum file size is 100MB.
			Stream will allow any file extension. If you want to be more restrictive
			for an application, this is can be set via API or by logging into your dashboard.
		`),
		Example: heredoc.Doc(`
			# List all
			$ istream input list
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.GetConfig(cmd)
			client, err := cfg.GetInputsClient(cmd)
			if err != nil {
				return err
			}

			// inputType, _ := cmd.Flags().GetString("type")
			ctx, cancel := cfg.GetClientCallContext(time.Second * 5)
			defer cancel()

			resp, err := client.ListInputs(ctx, &v1beta1.ListInputsRequest{})
			if err != nil {
				return fmt.Errorf("listing inputs: %w", err)
			}
			for _, input := range resp.Input {
				log.Println(input.Id, input.Title, input.Status)
			}
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringP("type", "t", "", "[option] Input type filter")

	return cmd
}
