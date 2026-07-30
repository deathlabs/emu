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
	devicesPageIndexDefault = 0
	devicesPageSizeDefault  = 20000
)

var (
	devicesPageIndex int
	devicesPageSize  int
)

var (
	getDevicesCmd = &cobra.Command{
		Use:   "devices",
		Short: "Get data about devices",
		RunE:  getDevices,
	}
)

func getDevices(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if devicesPageIndex != devicesPageIndexDefault {
		params.Set("pageIndex", strconv.Itoa(devicesPageIndex))
	}

	if devicesPageSize != devicesPageSizeDefault {
		if devicesPageSize > 20000 {
			return fmt.Errorf("the Page Size cannot exceed %d", devicesPageIndexDefault)
		}
		params.Set("pageSize", strconv.Itoa(devicesPageSize))
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/devices", config.Data.URL, system.ID)

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
	getDevicesCmd.PersistentFlags().IntVarP(&devicesPageIndex, "page-index", "", devicesPageIndexDefault, "Page index")
	getDevicesCmd.PersistentFlags().IntVarP(&devicesPageSize, "page-size", "", devicesPageSizeDefault, "Page size")
}
