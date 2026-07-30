package get

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

var (
	controlApprovalsControlAcronyms []string
)

var (
	getControlApprovalsCmd = &cobra.Command{
		Use:   "control-approvals",
		Short: "Get data about control approvals in the Control Approval Chain (CAC)",
		RunE:  getControlApprovals,
	}
)

func getControlApprovals(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if len(controlApprovalsControlAcronyms) > 0 {
		params.Set("controlAcronyms", strings.Join(controlApprovalsControlAcronyms, ","))
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {

		endpoint = fmt.Sprintf("%s/api/systems/%d/approval/cac", config.Data.URL, system.ID)

		if len(params) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
		}

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

func init() {
	getControlApprovalsCmd.PersistentFlags().StringSliceVarP(&controlApprovalsControlAcronyms, "control-acronyms", "", []string{}, "Control acronyms")
}
