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
	resultsAssessmentProcedures []string
	resultsControlAcronyms      []string
	resultsCcis                 []string
	resultsLatestOnly           bool
)

var (
	getTestResultsCmd = &cobra.Command{
		Use:   "test-results",
		Short: "Get data about test results",
		RunE:  getTestResults,
	}
)

func getTestResults(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
		response *http.Response
	)

	// Filter profiles based on the profile name provided via the root-level --profile flag.
	profiles, err = filterProfiles(config, activeProfileName)
	if err != nil {
		return err
	}

	// Loop through the filtered profiles and get test results data for each one.
	for _, profile = range profiles {
		// Define the endpoint for getting test results data for the current profile.
		endpoint = fmt.Sprintf("%s/api/test-results", config.URL)

		// If control IDs or assessment procedures are specified via the --control-ids
		// and --assessment-procedures flags, add them as query parameters.
		params = url.Values{}

		// If assessment procedures are specified via the --assessment-procedures flag,
		// add them as a query parameter.
		if len(resultsAssessmentProcedures) != 0 {
			params.Set("assessmentProcedures", strings.Join(resultsAssessmentProcedures, ","))
		}

		if len(resultsCcis) != 0 {
			params.Set("ccis", strings.Join(resultsCcis, ","))
		}

		// If control IDs are specified via the --control-ids flag,
		// add them as a query parameter.
		if len(resultsControlAcronyms) != 0 {
			params.Set("acronyms", strings.Join(resultsControlAcronyms, ","))
		}

		if resultsLatestOnly {
			params.Set("latestOnly", "true")
		}

		// If query parameters are set, add them to the endpoint.
		if len(params) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
		}

		// Make the request for test results data.
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

	return nil
}

func init() {
	// Define flags for the "emu get results" subcommand.
	getTestResultsCmd.PersistentFlags().StringSliceVarP(&resultsControlAcronyms, "control-acronyms", "", []string{}, "Control acronyms")
	getTestResultsCmd.PersistentFlags().StringSliceVarP(&resultsAssessmentProcedures, "assessment-procedures", "", []string{}, "Assessment procedures")
	getTestResultsCmd.PersistentFlags().StringSliceVarP(&resultsCcis, "ccis", "", []string{}, "CCIs")
	getTestResultsCmd.PersistentFlags().BoolVarP(&resultsLatestOnly, "latest-only", "", false, "Return only the latest test result for each control")

	// Add the "emu get results" subcommand to the "emu get" command.
	getCmd.AddCommand(getTestResultsCmd)
}
