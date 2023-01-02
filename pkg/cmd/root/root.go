package root

import (
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/input-stream/cli/pkg/cmd/input"
	"github.com/input-stream/cli/pkg/cmd/login"
	"github.com/input-stream/cli/pkg/config"
	"github.com/input-stream/cli/pkg/version"
)

var cfgPath = new(string)

func NewCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "istream <command> <subcommand> [flags]",
		Short: "Stream CLI",
		Long:  "Interact with your Stream applications easily",
		Example: heredoc.Doc(`
		`),
		Version: version.FmtVersion(),
	}

	fl := root.PersistentFlags()
	fl.StringVar(cfgPath, "config", "", "[optional] Explicit config file path")
	fl.String("base-dir", "", "[optional] Base directory file path")

	root.AddCommand(
		login.NewRootCmd(),
		input.NewRootCmd(),
	)

	cobra.OnInitialize(config.GetInitConfig(root, cfgPath))

	root.SetOut(os.Stdout)

	return root
}
