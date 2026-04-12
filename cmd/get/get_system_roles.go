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

	"github.com/deathlabs/emu/config"
	"github.com/deathlabs/emu/emass"
	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
)

var (
	rolesCategory string
	rolesRole     string
	rolesPolicy   string
)

var (
	getSystemRolesCmd = &cobra.Command{
		Use:   "system-roles",
		Short: "Get data about system roles",
		RunE:  getSystemRoles,
	}
)

func getSystemRoles(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		response *http.Response
		params   url.Values
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
	)

	// Filter profiles based on the profile name provided via the root-level --profile flag.
	profiles, err = config.FilterProfiles(config.Data, config.ActiveProfileName)
	if err != nil {
		return err
	}

	// Loop through the filtered profiles and get system role data for each one.
	for _, profile = range profiles {
		// Define the endpoint for getting system roles data for the current profile.
		endpoint = fmt.Sprintf("%s/api/system-roles", config.Data.URL)

		// If a role category is specified via the --category flag,
		// add it to the endpoint and set the required parameter.
		if rolesCategory != "" {
			// Return an error if a role argument is not provided.
			if rolesRole == "" {
				return fmt.Errorf("profile %s: a category and role are required", profile.Name)
			}

			// Add the role category to the endpoint.
			endpoint = fmt.Sprintf("%s/%s", endpoint, rolesCategory)

			// Set the required role parameter and optional policy parameter.
			params = url.Values{}
			params.Set("role", rolesRole)
			if rolesPolicy != "" {
				params.Set("policy", rolesPolicy)
			}

			// If query parameters are set, add them to the endpoint.
			if len(params) > 0 {
				endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
			}
		}

		// Make the request for system roles data.
		response, err = emass.Get(profile, endpoint)
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
	// Define flags for the "emu get system-roles" subcommand.
	getSystemRolesCmd.Flags().StringVarP(&rolesCategory, "category", "", "", "PAC, CAC, or Other")
	getSystemRolesCmd.Flags().StringVarP(&rolesRole, "role", "", "", "ISO, ISSM, SCA, Auditor, AO, etc. (required if --category is used)")
	getSystemRolesCmd.Flags().StringVarP(&rolesPolicy, "policy", "", "", "RMF, DIACAP, or Reporting")
}
