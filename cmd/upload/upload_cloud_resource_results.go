package upload

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	uploadCloudResourceResultsCmd = &cobra.Command{
		Use:   "cloud-resource-results",
		Short: "Upload cloud resource results",
		Run:   updateCloudResourceResults,
	}
)

func updateCloudResourceResults(cmd *cobra.Command, args []string) {
	fmt.Println("emu update artifact")
}
