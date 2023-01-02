package input

import (
	"github.com/spf13/cobra"

	"github.com/input-stream/cli/pkg/cmd/input/create"
	"github.com/input-stream/cli/pkg/cmd/input/delete"
	"github.com/input-stream/cli/pkg/cmd/input/edit"
	"github.com/input-stream/cli/pkg/cmd/input/get"
	"github.com/input-stream/cli/pkg/cmd/input/list"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "input",
		Short: "input create, read, update, delete operations",
	}

	cmd.AddCommand(create.NewCmds()...)
	cmd.AddCommand(delete.NewCmds()...)
	cmd.AddCommand(get.NewCmds()...)
	cmd.AddCommand(list.NewCmds()...)
	cmd.AddCommand(edit.NewCmds()...)

	return cmd
}
