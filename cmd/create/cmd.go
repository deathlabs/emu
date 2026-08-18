package create

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "create",
		Short: "Create data",
	}
)

func init() {
	Cmd.AddCommand(createPoamCmd)
}
