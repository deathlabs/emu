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
	controlsControlAcronyms []string
)

var (
	getControlsCmd = &cobra.Command{
		Use:   "controls",
		Short: "Get data about controls",
		RunE:  getControls,
	}
)

func getControls(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		response *http.Response
		params   url.Values
		system   models.System
		systems  []models.System
	)

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {

		endpoint = fmt.Sprintf("%s/api/systems/%d/controls", config.Data.URL, system.ID)

		params = url.Values{}

		params.Set("acronyms", strings.Join(controlsControlAcronyms, ","))
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
	getControlsCmd.PersistentFlags().StringSliceVarP(&controlsControlAcronyms, "control-acronyms", "", []string{}, "Control acronyms")
}
