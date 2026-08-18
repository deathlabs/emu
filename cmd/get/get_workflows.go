package get

import (
	"fmt"
	"net/http"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

var (
	getWorkflowsCmd = &cobra.Command{
		Use:   "workflows",
		Short: "Get data about workflows in the Package Approval Chain (PAC)",
		RunE:  getWorkflows,
	}
)

func getWorkflows(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		response *http.Response
		system   models.System
		systems  []models.System
	)

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {

		endpoint = fmt.Sprintf("%s/api/systems/%d/approval/pac", config.Data.URL, system.ID)

		response, err = emass.Get(system.ConfigProfile, endpoint)
		if err != nil {
			return err
		}

		err = output.Response(response, config.OutputFormat)
		if err != nil {
			return err
		}
	}

	return nil
}
