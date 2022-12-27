package input

import (
	"github.com/spf13/cobra"

	"github.com/input-stream/cli/pkg/cmd/input/get"
	"github.com/input-stream/cli/pkg/cmd/input/list"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "input",
		Short: "create, read, update, delete operations on inputs",
	}

	cmd.AddCommand(list.NewCmds()...)
	cmd.AddCommand(get.NewCmds()...)

	return cmd
}
