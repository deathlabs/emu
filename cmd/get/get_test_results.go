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
