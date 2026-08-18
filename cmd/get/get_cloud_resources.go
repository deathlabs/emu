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
	cloudResourcesPageIndexDefault = 0
	cloudResourcesPageSizeDefault  = 20000
)

var (
	cloudResourcesPageIndex int
	cloudResourcesPageSize  int
)

var (
	getCloudResourcesCmd = &cobra.Command{
		Use:   "cloud-resources",
		Short: "Get data about cloud resources",
		RunE:  getCloudResources,
	}
)

func getCloudResources(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if cloudResourcesPageIndex != cloudResourcesPageIndexDefault {
		params.Set("pageIndex", strconv.Itoa(cloudResourcesPageIndex))
	}

	if cloudResourcesPageSize != cloudResourcesPageSizeDefault {
		if cloudResourcesPageSize > 20000 {
			return fmt.Errorf("the Page Size cannot exceed %d", cloudResourcesPageIndexDefault)
		}
		params.Set("pageSize", strconv.Itoa(cloudResourcesPageSize))
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
	getCloudResourcesCmd.PersistentFlags().IntVarP(&cloudResourcesPageIndex, "page-index", "", cloudResourcesPageIndexDefault, "Page index")
	getCloudResourcesCmd.PersistentFlags().IntVarP(&cloudResourcesPageSize, "page-size", "", cloudResourcesPageSizeDefault, "Page size")
}
