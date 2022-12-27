package list

import (
	"context"
	"fmt"
	"log"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/input-stream/cli/pkg/config"
	"github.com/input-stream/cli/stream/input/v1beta1"
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
			# Uploads a file to 'redteam' channel of 'messaging' channel type
			$ stream-cli chat upload-file --channel-type messaging --channel-id redteam --user-id "user-1" --file "./snippet.txt"
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := config.GetConfig(cmd).GetInputsClient(cmd)
			if err != nil {
				return err
			}

			// inputType, _ := cmd.Flags().GetString("type")

			resp, err := client.ListInputs(context.Background(), &v1beta1.ListInputsRequest{})
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
