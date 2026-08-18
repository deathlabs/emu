package upload

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "upload",
		Short: "Upload data",
	}
)

func init() {
	Cmd.AddCommand(uploadArtifactCmd)
	Cmd.AddCommand(uploadCloudResourceResultsCmd)
	Cmd.AddCommand(uploadContainerSBOMCmd)
	Cmd.AddCommand(uploadDeviceScanCmd)
	Cmd.AddCommand(uploadSoftwareBaselineCmd)
}
