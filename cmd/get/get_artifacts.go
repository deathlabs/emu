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
	artifactsFilename             string
	artifactsControlAcronyms      []string
	artifactsAssessmentProcedures []string
	artifactsCcis                 []string
	artifactsSystemOnly           bool
)

var (
	getArtifactsCmd = &cobra.Command{
		Use:   "artifacts",
		Short: "Get data about artifacts",
		RunE:  getArtifacts,
	}
)

func getArtifacts(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		params   url.Values
		response *http.Response
		system   models.System
		systems  []models.System
	)

	params = url.Values{}

	if len(artifactsAssessmentProcedures) > 0 {
		params.Set("assessmentProcedures", strings.Join(artifactsAssessmentProcedures, ","))
	}

	if len(artifactsCcis) > 0 {
		params.Set("ccis", strings.Join(artifactsCcis, ","))
	}

	if len(artifactsControlAcronyms) > 0 {
		params.Set("controlAcronyms", strings.Join(artifactsControlAcronyms, ","))
	}

	if artifactsFilename != "" {
		params.Set("filename", artifactsFilename)
	}

	if artifactsSystemOnly {
		params.Set("systemOnly", "true")
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {

		endpoint = fmt.Sprintf("%s/api/systems/%d/artifacts", config.Data.URL, system.ID)

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
	getArtifactsCmd.PersistentFlags().StringVarP(&artifactsFilename, "filename", "f", "", "Filename")
	getArtifactsCmd.PersistentFlags().StringSliceVarP(&artifactsControlAcronyms, "control-acronyms", "", []string{}, "Control acronyms")
	getArtifactsCmd.PersistentFlags().StringSliceVarP(&artifactsAssessmentProcedures, "assessment-procedures", "", []string{}, "Assessment procedures")
	getArtifactsCmd.PersistentFlags().StringSliceVarP(&artifactsCcis, "ccis", "", []string{}, "CCIs")
	getArtifactsCmd.PersistentFlags().BoolVarP(&artifactsSystemOnly, "system-only", "", false, "Exclude control and AP-level artifacts only")
}
