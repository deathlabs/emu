package get

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

const (
	cloudResourceResultsPageIndexDefault = 0
	cloudResourceResultsPageSizeDefault  = 20000
)

var (
	cloudResourceResultsPageIndex int
	cloudResourceResultsPageSize  int
)

var (
	getCloudResourceResultsCmd = &cobra.Command{
		Use:   "cloud-resource-results",
		Short: "Get data about cloud resource results",
		RunE:  getCloudResourceResults,
	}
)

func getCloudResourceResults(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if cloudResourceResultsPageIndex != cloudResourceResultsPageIndexDefault {
		params.Set("pageIndex", strconv.Itoa(cloudResourceResultsPageIndex))
	}

	if cloudResourceResultsPageSize != cloudResourceResultsPageSizeDefault {
		if cloudResourceResultsPageSize > 20000 {
			return fmt.Errorf("the Page Size cannot exceed %d", cloudResourceResultsPageIndexDefault)
		}
		params.Set("pageSize", strconv.Itoa(cloudResourceResultsPageSize))
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/cloud-resources", config.Data.URL, system.ID)

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
	getCloudResourceResultsCmd.PersistentFlags().IntVarP(&cloudResourceResultsPageIndex, "page-index", "", cloudResourceResultsPageIndexDefault, "Page index")
	getCloudResourceResultsCmd.PersistentFlags().IntVarP(&cloudResourceResultsPageSize, "page-size", "", cloudResourceResultsPageSizeDefault, "Page size")
}
