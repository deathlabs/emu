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
package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/deathlabs/emu/emass"
	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
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

	// Filter systems based on system IDs provided via the root-level --system-ids flag.
	// If no system IDs are provided, this will return all systems for the active profile.
	systems, err = filterSystems(config, activeProfileName, systemIDs)
	if err != nil {
		return err
	}

	// Loop through the filtered systems and get controls data for each one.
	for _, system = range systems {
		// Define the endpoint for getting controls data for the current system.
		endpoint = fmt.Sprintf("%s/api/systems/%d/controls", config.URL, system.ID)

		// If control IDs are specified via the --control-ids flag, add them as a query parameter.
		params = url.Values{}

		// If control IDs are specified via the --control-ids flag, add them as a query parameter.
		params.Set("acronyms", strings.Join(controlIDs, ","))
		if len(params) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
		}

		// Make the request for controls data.
		response, err = emass.Get(system.ConfigProfile, endpoint)
		if err != nil {
			return err
		}

		// Print the response in the specified output format.
		err = output.Response(response, outputFormat)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	// Define flags for the "emu get controls" subcommand.
	getControlsCmd.PersistentFlags().StringSliceVarP(&controlIDs, "control-ids", "", []string{}, "Control IDs")

	// Add the "emu get controls" subcommand to the "emu get" command.
	getCmd.AddCommand(getControlsCmd)
}
