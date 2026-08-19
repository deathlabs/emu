package get

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

var (
	poamsScheduledCompletionDateStart int
	poamsScheduledCompletionDateEnd   string
	poamsControlAcronyms              []string
	poamsAssessmentProcedures         []string
	poamsCcis                         []string
	poamsSystemOnly                   bool
)

var (
	getPoamsCmd = &cobra.Command{
		Use:   "poams",
		Short: "Get data about Plan of Action and Milestones (POA&M) items",
		RunE:  getPoams,
	}
)

func getPoams(cmd *cobra.Command, args []string) error {
	var (
		endpoint        string
		err             error
		params          url.Values
		queryParameters url.Values
		response        *http.Response
		system          models.System
		systems         []models.System
	)

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	queryParameters = url.Values{}

	if poamsScheduledCompletionDateStart != 0 {
		queryParameters.Set("scheduledCompletionDateStart", strconv.Itoa(poamsScheduledCompletionDateStart))
	}

	if poamsScheduledCompletionDateEnd != "" {
		queryParameters.Set("scheduledCompletionDateEnd", poamsScheduledCompletionDateEnd)
	}

	if len(poamsControlAcronyms) > 0 {
		queryParameters.Set("controlAcronyms", strings.Join(poamsControlAcronyms, ","))
	}

	if len(poamsAssessmentProcedures) > 0 {
		queryParameters.Set("assessmentProcedures", strings.Join(poamsAssessmentProcedures, ","))
	}

	if len(poamsCcis) > 0 {
		queryParameters.Set("ccis", strings.Join(poamsCcis, ","))
	}

	queryParameters.Set("systemOnly", strconv.FormatBool(poamsSystemOnly))

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/poams?%s", config.Data.URL, system.ID, queryParameters.Encode())

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
	getPoamsCmd.PersistentFlags().IntVarP(&poamsScheduledCompletionDateStart, "scheduled-completion-date-start", "", 0, "Scheduled completion date start")
	getPoamsCmd.PersistentFlags().StringVarP(&poamsScheduledCompletionDateEnd, "scheduled-completion-date-end", "", "", "Scheduled completion date end")
	getPoamsCmd.PersistentFlags().StringArrayVarP(&poamsControlAcronyms, "control-acronyms", "", []string{}, "Control acronyms")
	getPoamsCmd.PersistentFlags().StringArrayVarP(&poamsControlAcronyms, "assessment-procedures", "", []string{}, "Assessment procedures")
	getPoamsCmd.PersistentFlags().StringArrayVarP(&poamsCcis, "ccis", "", []string{}, "CCIs")
	getPoamsCmd.PersistentFlags().BoolVarP(&poamsSystemOnly, "system-only", "", false, "System only")
}
