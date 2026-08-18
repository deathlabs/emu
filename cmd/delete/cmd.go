package delete

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete data",
	}
)

func init() {
	Cmd.AddCommand(deleteArtifactCmd)
}
