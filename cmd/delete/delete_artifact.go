package delete

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	deleteArtifactCmd = &cobra.Command{
		Use:   "artifact",
		Short: "Delete an artifact in eMASS",
		Run:   deleteArtifact,
	}
)

func deleteArtifact(cmd *cobra.Command, args []string) {
	fmt.Println("emu delete artifact")
}
