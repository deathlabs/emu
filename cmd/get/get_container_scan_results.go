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
	containerScanResultsPageIndexDefault = 0
	containerScanResultsPageSizeDefault  = 20000
)

var (
	containerScanResultsPageIndex int
	containerScanResultsPageSize  int
)

var (
	getContainerScanResultsCmd = &cobra.Command{
		Use:   "container-scan-results",
		Short: "Get data about container scan results",
		RunE:  getContainerScanResults,
	}
)

func getContainerScanResults(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if containerScanResultsPageIndex != containerScanResultsPageIndexDefault {
		params.Set("pageIndex", strconv.Itoa(containerScanResultsPageIndex))
	}

	if containerScanResultsPageSize != containerScanResultsPageSizeDefault {
		if containerScanResultsPageSize > 20000 {
			return fmt.Errorf("the Page Size cannot exceed %d", containerScanResultsPageIndexDefault)
		}
		params.Set("pageSize", strconv.Itoa(containerScanResultsPageSize))
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/container-scan-results", config.Data.URL, system.ID)

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
	getContainerScanResultsCmd.PersistentFlags().IntVarP(&containerScanResultsPageIndex, "page-index", "", containerScanResultsPageIndexDefault, "Page index")
	getContainerScanResultsCmd.PersistentFlags().IntVarP(&containerScanResultsPageSize, "page-size", "", containerScanResultsPageSizeDefault, "Page size")
}
