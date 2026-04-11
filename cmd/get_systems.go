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

	"github.com/deathlabs/emu/emass"
	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
)

var (
	systemsCoamsID               string
	systemsDitprID               string
	systemsIncludeDecommissioned bool
	systemsIncludeDitprMetrics   bool
	systemsPolicy                string
	systemsRegistrationType      string
	systemsReportsForScorecard   bool
)

var (
	getSystemsCmd = &cobra.Command{
		Use:   "systems",
		Short: "Get data about systems",
		RunE:  getSystems,
	}
)

func getSystems(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		response *http.Response
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
		system   models.System
		systems  []models.System
	)
	// TODO: add params

	// If system IDs are provided via the root-level --system-ids flag, use them to filter systems.
	// Otherwise, filter profiles based on the active profile name and get all systems for those profiles.
	if len(systemIDs) != 0 {
		// Filter systems based on system IDs provided via the root-level --system-ids flag.
		// If no system IDs are provided, this will return all systems for the active profile.
		systems, err = filterSystems(config, activeProfileName, systemIDs)
		if err != nil {
			return err
		}

		for _, system = range systems {
			// Define the endpoint for getting systems data for the current system.
			endpoint = fmt.Sprintf("%s/api/systems/%d", config.URL, system.ID)

			// Make the request for systems data.
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
	} else {
		// Filter profiles based on the profile name provided via the root-level --profile flag.
		profiles, err = filterProfiles(config, activeProfileName)
		if err != nil {
			return err
		}

		// Loop through the filtered profiles and get systems data for each one.
		for _, profile = range profiles {
			// Define the endpoint for getting systems data for the current profile.
			endpoint = fmt.Sprintf("%s/api/systems", config.URL)

			// Make the request for systems data.
			response, err = emass.Get(profile, endpoint)
			if err != nil {
				return err
			}

			// Print the response in the specified output format.
			err = output.Response(response, outputFormat)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func init() {
	// Define flags for the "emu get systems" subcommand.
	getSystemsCmd.PersistentFlags().StringVarP(&systemsCoamsID, "coams-id", "", "", "COAMS ID")
	getSystemsCmd.PersistentFlags().StringVarP(&systemsDitprID, "ditpr-id", "", "", "DITPR ID")
	getSystemsCmd.PersistentFlags().BoolVarP(&systemsIncludeDecommissioned, "include-decommissioned", "", false, "Include decommissioned systems")
	getSystemsCmd.PersistentFlags().StringVarP(&systemsPolicy, "policy", "", "", "Policy (DIACAP, RMF, or Reporting)")
	getSystemsCmd.PersistentFlags().BoolVarP(&systemsIncludeDitprMetrics, "include-ditpr-metrics", "", false, "Include DITPR metrics (cannot be used in conjunction with --coams-id or --ditpr-id)")
	getSystemsCmd.PersistentFlags().StringVarP(&systemsRegistrationType, "registration-type", "", "", "Registration type (assessAndAuthorize, assessOnly, guest, regular, functional, cloudServiceProvider, commonControlProvider, authorizationToUse, reciprocityAcceptance)")
	getSystemsCmd.PersistentFlags().BoolVarP(&systemsReportsForScorecard, "reports-for-scorecard", "", false, "Return only systems that report to the DOD Cyber Hygiene Scorecard")

	// Add the "emu get systems" subcommand to the "emu get" command.
	getCmd.AddCommand(getSystemsCmd)
}
