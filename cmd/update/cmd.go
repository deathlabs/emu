package update

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "update",
		Short: "Update data",
	}
)

func init() {
	Cmd.AddCommand(updateArtifactCmd)
}
