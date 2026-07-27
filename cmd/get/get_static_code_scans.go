/*
Copyright © 2026 Victor Fernandez III <@cyberphor>

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
	staticCodeScansPageIndexDefault = 0
	staticCodeScansPageSizeDefault  = 20000
)

var (
	staticCodeScansPageIndex int
	staticCodeScansPageSize  int
)

var (
	getStaticCodeScansCmd = &cobra.Command{
		Use:   "static-code-scans",
		Short: "Get data about static code scans",
		RunE:  getStaticCodeScans,
	}
)

func getStaticCodeScans(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if staticCodeScansPageIndex != staticCodeScansPageIndexDefault {
		params.Set("pageIndex", strconv.Itoa(staticCodeScansPageIndex))
	}

	if staticCodeScansPageSize != staticCodeScansPageSizeDefault {
		if staticCodeScansPageSize > 20000 {
			return fmt.Errorf("the Page Size cannot exceed %d", staticCodeScansPageIndexDefault)
		}
		params.Set("pageSize", strconv.Itoa(staticCodeScansPageSize))
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/static-code-scans", config.Data.URL, system.ID)

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
	getStaticCodeScansCmd.PersistentFlags().IntVarP(&staticCodeScansPageIndex, "page-index", "", staticCodeScansPageIndexDefault, "Page index")
	getStaticCodeScansCmd.PersistentFlags().IntVarP(&staticCodeScansPageSize, "page-size", "", staticCodeScansPageSizeDefault, "Page size")
}
