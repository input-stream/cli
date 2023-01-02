package get

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
		getInputCmd(),
	}
}

func getInputCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get --id [id] --title [title] --login [login] --owner [owner]",
		Short: "Get a single input by ID",
		Long: heredoc.Doc(`
		`),
		Example: heredoc.Doc(`
			# Get by ID
			$ istream input get --id 50286c35-5c15-4d11-8989-da82237453fb

			# Get by Title
			$ istream input get --title todo
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.GetConfig(cmd)
			client, err := cfg.GetInputsClient(cmd)
			if err != nil {
				return err
			}

			ctx, cancel := cfg.GetClientCallContext(time.Second * 5)
			defer cancel()

			return runGet(ctx, client, cmd)
		},
	}

	flags := cmd.Flags()
	flags.String("id", "", "[optional] ID")
	flags.String("title", "", "[optional] Title")
	flags.String("login", "", "[optional] Login Name")
	flags.String("owner", "", "[optional] Owner Name")

	return cmd
}

func runGet(ctx context.Context, client ispb.InputsClient, cmd *cobra.Command) error {
	id, _ := cmd.Flags().GetString("id")
	title, _ := cmd.Flags().GetString("title")
	owner, _ := cmd.Flags().GetString("owner")
	login, _ := cmd.Flags().GetString("login")
	input, err := client.GetInput(ctx, &ispb.GetInputRequest{
		Filter: &ispb.InputFilterOptions{
			Id:        id,
			TitleSlug: title,
			Owner:     owner,
			Login:     login,
		},
	})
	if err != nil {
		return fmt.Errorf("listing inputs: %w", err)
	}
	return protobuf.WritePrettyJSON(cmd.OutOrStdout(), input)
}
