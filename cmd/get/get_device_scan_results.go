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
	deviceScanResultsPageIndexDefault = 0
	deviceScanResultsPageSizeDefault  = 20000
)

var (
	deviceScanResultsPageIndex int
	deviceScanResultsPageSize  int
)

var (
	getDeviceScanResultsCmd = &cobra.Command{
		Use:   "device-scan-results",
		Short: "Get data about device scan results",
		RunE:  getDeviceScanResults,
	}
)

func getDeviceScanResults(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if deviceScanResultsPageIndex != deviceScanResultsPageIndexDefault {
		params.Set("pageIndex", strconv.Itoa(deviceScanResultsPageIndex))
	}

	if deviceScanResultsPageSize != deviceScanResultsPageSizeDefault {
		if deviceScanResultsPageSize > 20000 {
			return fmt.Errorf("the Page Size cannot exceed %d", deviceScanResultsPageIndexDefault)
		}
		params.Set("pageSize", strconv.Itoa(deviceScanResultsPageSize))
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/device-scan-results", config.Data.URL, system.ID)

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
	getDeviceScanResultsCmd.PersistentFlags().IntVarP(&deviceScanResultsPageIndex, "page-index", "", deviceScanResultsPageIndexDefault, "Page index")
	getDeviceScanResultsCmd.PersistentFlags().IntVarP(&deviceScanResultsPageSize, "page-size", "", deviceScanResultsPageSizeDefault, "Page size")
}
