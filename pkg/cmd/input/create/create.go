package create

import (
	"context"
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	ispb "github.com/input-stream/cli/build/stack/inputstream/v1beta1"
	"github.com/input-stream/cli/pkg/config"
	"github.com/input-stream/cli/pkg/protobuf"
)

func NewCmds() []*cobra.Command {
	return []*cobra.Command{
		createInputCmd(),
	}
}

func createInputCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create --title [title] --login [login] --owner [owner]",
		Short: "Create a single input",
		Long: heredoc.Doc(`
		`),
		Example: heredoc.Doc(`
			# Create (all defaults)
			$ istream input create

			# Create (with options)
			$ istream input create --title 'My TODO List'
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.GetConfig(cmd)
			client, err := cfg.GetInputsClient(cmd)
			if err != nil {
				return err
			}

			ctx, cancel := cfg.GetClientCallContext(time.Second * 5)
			defer cancel()

			return runCreate(ctx, client, cmd)
		},
	}

	flags := cmd.Flags()
	flags.String("title", "", "[optional] Title")

	return cmd
}

func runCreate(ctx context.Context, client ispb.InputsClient, cmd *cobra.Command) error {
	title, _ := cmd.Flags().GetString("title")
	input, err := client.CreateInput(ctx, &ispb.CreateInputRequest{
		Input: &ispb.Input{
			Title: title,
		},
	})
	if err != nil {
		return fmt.Errorf("listing inputs: %w", err)
	}
	return protobuf.WritePrettyJSON(cmd.OutOrStdout(), input)
}
