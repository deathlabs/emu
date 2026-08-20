package get

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "get",
		Short: "Get data",
	}
)

func init() {
	Cmd.AddCommand(getApplicationsCmd)
	Cmd.AddCommand(getArtifactsCmd)
	Cmd.AddCommand(getCloudResourcesCmd)
	Cmd.AddCommand(getCloudResourceResultsCmd)
	Cmd.AddCommand(getConfigCmd)
	Cmd.AddCommand(getContainersCmd)
	Cmd.AddCommand(getContainerScanResultsCmd)
	Cmd.AddCommand(getControlsCmd)
	Cmd.AddCommand(getControlApprovalsCmd)
	Cmd.AddCommand(getDevicesCmd)
	Cmd.AddCommand(getDeviceScanResultsCmd)
	Cmd.AddCommand(getPoamsCmd)
	Cmd.AddCommand(getPpsCmd)
	Cmd.AddCommand(getStaticCodeScansCmd)
	Cmd.AddCommand(getSystemsCmd)
	Cmd.AddCommand(getSystemRolesCmd)
	Cmd.AddCommand(getTestResultsCmd)
	Cmd.AddCommand(getWorkflowsCmd)
}
