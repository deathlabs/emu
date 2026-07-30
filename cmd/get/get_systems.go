package get

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
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
		params   url.Values
		response *http.Response
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if systemsCoamsID != "" {
		params.Set("coamsId", systemsCoamsID)
	}

	if systemsDitprID != "" {
		params.Set("ditprId", systemsDitprID)
	}

	if systemsIncludeDecommissioned {
		params.Set("includeDecommissioned", "true")
	}

	if systemsIncludeDitprMetrics {
		params.Set("includeDitprMetrics", "true")
	}

	if systemsPolicy != "" {
		params.Set("policy", systemsPolicy)
	}

	if systemsRegistrationType != "" {
		params.Set("registrationType", systemsRegistrationType)
	}

	if systemsReportsForScorecard {
		params.Set("reportsForScorecard", "true")
	}

	if len(config.SystemIDs) != 0 {

		systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
		if err != nil {
			return err
		}

		for _, system = range systems {

			endpoint = fmt.Sprintf("%s/api/systems/%d", config.Data.URL, system.ID)

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
	} else {

		profiles, err = config.FilterProfiles(config.Data, config.ActiveProfileName)
		if err != nil {
			return err
		}

		for _, profile = range profiles {

			endpoint = fmt.Sprintf("%s/api/systems", config.Data.URL)

			if len(params) > 0 {
				endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
			}

			response, err = emass.Get(profile, endpoint)
			if err != nil {
				return err
			}

			err = output.Response(response, config.OutputFormat)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func init() {
	getSystemsCmd.PersistentFlags().StringVarP(&systemsCoamsID, "coams-id", "", "", "COAMS ID")
	getSystemsCmd.PersistentFlags().StringVarP(&systemsDitprID, "ditpr-id", "", "", "DITPR ID")
	getSystemsCmd.PersistentFlags().BoolVarP(&systemsIncludeDecommissioned, "include-decommissioned", "", false, "Include decommissioned systems")
	getSystemsCmd.PersistentFlags().StringVarP(&systemsPolicy, "policy", "", "", "Policy (DIACAP, RMF, or Reporting)")
	getSystemsCmd.PersistentFlags().BoolVarP(&systemsIncludeDitprMetrics, "include-ditpr-metrics", "", false, "Include DITPR metrics (cannot be used in conjunction with --coams-id or --ditpr-id)")
	getSystemsCmd.PersistentFlags().StringVarP(&systemsRegistrationType, "registration-type", "", "", "Registration type (assessAndAuthorize, assessOnly, guest, regular, functional, cloudServiceProvider, commonControlProvider, authorizationToUse, reciprocityAcceptance)")
	getSystemsCmd.PersistentFlags().BoolVarP(&systemsReportsForScorecard, "reports-for-scorecard", "", false, "Return only systems that report to the DOD Cyber Hygiene Scorecard")
}
