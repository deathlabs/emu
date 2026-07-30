package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

const (
	cloudResourceResultsProviderDefault   = ""
	cloudResourceResultsResourceIdDefault = ""
)

var (
	cloudResourceResultsProvider   string
	cloudResourceResultsResourceId string
)

var (
	uploadCloudResourceResultsCmd = &cobra.Command{
		Use:   "cloud-resource-results",
		Short: "Upload cloud resource results",
		RunE:  updateCloudResourceResults,
	}
)

func updateCloudResourceResults(cmd *cobra.Command, args []string) error {
	var (
		endpoint        string
		err             error
		headers         map[string]string
		request         *bytes.Buffer
		requestBody     []byte
		requestBodyData models.CloudResourceResult
		response        *http.Response
		system          models.System
		systems         []models.System
	)

	headers = map[string]string{
		"Content-Type": "application/json",
	}

	requestBodyData = models.CloudResourceResult{
		Provider:     "azure",
		ResourceId:   "/subscriptions/123456789/sample/resource/namespace/default",
		ResourceName: "Storage Resource",
		ResourceType: "Microsoft.storage.table",
		InitiatedBy:  "john.doe.ctr@mail.mil",
		CspAccountId: "123456789",
		CspRegion:    "useast2",
		IsBaseline:   true,
		Tags: map[string]string{
			"test": "testtag",
		},
		ComplianceResults: []models.ComplianceResult{
			{
				CspPolicyDefinitionId:    "/providers/sample/policy/namespace/au11_policy",
				PolicyDefinitionTitle:    "AU-11 - Audit Record Retention",
				ComplianceCheckTimestamp: 1644003780,
				IsCompliant:              false,
				Control:                  "AU-11",
				AssessmentProcedure:      "000167,000168",
				ComplianceReason:         "retention period not configured",
				PolicyDeploymentName:     "testDeployment",
				PolicyDeploymentVersion:  "1.0.0",
				Severity:                 "High",
			},
		},
	}

	requestBody, err = json.Marshal([]models.CloudResourceResult{requestBodyData})
	if err != nil {
		return err
	}

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/cloud-resource-results", config.Data.URL, system.ID)

		request = bytes.NewBuffer(requestBody)
		response, err = emass.Post(system.ConfigProfile, endpoint, headers, request)
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
