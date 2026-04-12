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

	"github.com/deathlabs/emu/config"
	"github.com/deathlabs/emu/emass"
	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
)

var (
	getWorkflowsCmd = &cobra.Command{
		Use:   "workflows",
		Short: "Get data about workflows in the Package Approval Chain (PAC)",
		RunE:  getWorkflows,
	}
)

func getWorkflows(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		response *http.Response
		system   models.System
		systems  []models.System
	)

	// Filter systems based on system IDs provided via the root-level --system-ids flag.
	// If no system IDs are provided, this will return all systems for the active profile.
	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	// Loop through the filtered systems and get workflow data for each one.
	for _, system = range systems {
		// Define the endpoint for getting workflow data for the current system.
		endpoint = fmt.Sprintf("%s/api/systems/%d/approval/pac", config.Data.URL, system.ID)

		// Make the request for workflow data.
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
