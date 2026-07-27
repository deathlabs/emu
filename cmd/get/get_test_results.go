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
	"strings"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
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
		system   models.System
		systems  []models.System
		response *http.Response
	)

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {

		endpoint = fmt.Sprintf("%s/api/systems/%d/test-results", config.Data.URL, system.ID)

		params = url.Values{}

		if len(resultsAssessmentProcedures) != 0 {
			params.Set("assessmentProcedures", strings.Join(resultsAssessmentProcedures, ","))
		}

		if len(resultsCcis) != 0 {
			params.Set("ccis", strings.Join(resultsCcis, ","))
		}

		if len(resultsControlAcronyms) != 0 {
			params.Set("acronyms", strings.Join(resultsControlAcronyms, ","))
		}

		if resultsLatestOnly {
			params.Set("latestOnly", "true")
		}

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
	getTestResultsCmd.PersistentFlags().StringSliceVarP(&resultsControlAcronyms, "control-acronyms", "", []string{}, "Control acronyms")
	getTestResultsCmd.PersistentFlags().StringSliceVarP(&resultsAssessmentProcedures, "assessment-procedures", "", []string{}, "Assessment procedures")
	getTestResultsCmd.PersistentFlags().StringSliceVarP(&resultsCcis, "ccis", "", []string{}, "CCIs")
	getTestResultsCmd.PersistentFlags().BoolVarP(&resultsLatestOnly, "latest-only", "", false, "Return only the latest test result for each control")
}
