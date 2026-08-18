package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	updateArtifactCmd = &cobra.Command{
		Use:   "artifact",
		Short: "Update artifact data",
		Run:   updateArtifact,
	}
)

func updateArtifact(cmd *cobra.Command, args []string) {
	fmt.Println("emu update artifact")
}
