/*
Copyright © 2026 Vic Fernandez III <@cyberphor>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
	ppsPageIndexDefault = 0
	ppsPageSizeDefault  = 20000
)

var (
	ppsPageIndex int
	ppsPageSize  int
)

var (
	getPpsCmd = &cobra.Command{
		Use:   "pps",
		Short: "Get data about ports, protocols, and services",
		RunE:  getPps,
	}
)

func getPps(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if ppsPageIndex != ppsPageIndexDefault {
		params.Set("pageIndex", strconv.Itoa(ppsPageIndex))
	}

	if ppsPageSize != ppsPageSizeDefault {
		if ppsPageSize > 20000 {
			return fmt.Errorf("the Page Size cannot exceed %d", ppsPageIndexDefault)
		}
		params.Set("pageSize", strconv.Itoa(ppsPageSize))
	}

	// Filter systems based on system IDs provided via the root-level --system-ids flag.
	// If no system IDs are provided, this will return all systems for the active profile.
	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	// Loop through the filtered systems and get PPS data for each one.
	for _, system = range systems {
		// Define the endpoint for getting PPS data for the current system.
		endpoint = fmt.Sprintf("%s/api/systems/%d/ports-protocols", config.Data.URL, system.ID)

		if len(params) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
		}

		response, err = emass.Get(system.ConfigProfile, endpoint)
		if err != nil {
			return err
		}

		// Print the response in the specified output format.
		err = output.Response(response, config.OutputFormat)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	// Define flags for the "emu get pps" subcommand.
	getPpsCmd.PersistentFlags().IntVarP(&ppsPageIndex, "page-index", "", ppsPageIndexDefault, "Page index")
	getPpsCmd.PersistentFlags().IntVarP(&ppsPageIndex, "page-size", "", ppsPageSizeDefault, "Page size")
}
