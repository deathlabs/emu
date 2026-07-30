package upload

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

const (
	cloudResourceResultsProviderDefault   = ""
	cloudResourceResultsResourceIdDefault = ""
)

var (
	cloudResourceResultsProvider   string
	cloudResourceResultsResourceId string
)

var (
	uploadCloudResourceResultsCmd = &cobra.Command{
		Use:   "cloud-resource-results",
		Short: "Upload cloud resource results",
		RunE:  updateCloudResourceResults,
	}
)

func updateCloudResourceResults(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if cloudResourceResultsProvider != cloudResourceResultsProviderDefault {
		if len(cloudResourceResultsProvider) > 100 {
			return fmt.Errorf("the Provider length cannot exceed %d characters", 100)
		}
		params.Set("provider", cloudResourceResultsProvider)
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/cloud-resource-results", config.Data.URL, system.ID)

		if len(params) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
		}

		response, err = emass.Post(system.ConfigProfile, endpoint, nil, nil)
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
