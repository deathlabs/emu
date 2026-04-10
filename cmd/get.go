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
	"net/url"
	"strings"

	"github.com/deathlabs/emu/emass"
	"github.com/deathlabs/emu/models"
	"github.com/spf13/cobra"
)

var (
	assessmentProcedures []string
	controlIDs           []string
	getCmd               = &cobra.Command{
		Use:   "get",
		Short: "Get data about systems, controls, approvals, artifacts, roles, and workflows",
	}
	getControlsCmd = &cobra.Command{
		Use:   "controls",
		Short: "Get data about controls",
		RunE:  getControls,
	}
	getApprovalsCmd = &cobra.Command{
		Use:   "approvals",
		Short: "Get data about approvals",
		RunE:  getApprovals,
	}
	getArtifactsCmd = &cobra.Command{
		Use:   "artifacts",
		Short: "Get data about artifacts",
		RunE:  getArtifacts,
	}
	getResultsCmd = &cobra.Command{
		Use:   "results",
		Short: "Get data about test results",
		RunE:  getResults,
	}
	getRolesCmd = &cobra.Command{
		Use:   "roles",
		Short: "Get data about roles",
		RunE:  getRoles,
	}
	getSystemsCmd = &cobra.Command{
		Use:   "systems",
		Short: "Get data about systems",
		RunE:  getSystems,
	}
	getWorkflowsCmd = &cobra.Command{
		Use:   "workflows",
		Short: "Get data about workflows",
		RunE:  getWorkflows,
	}
	role         string
	roleCategory string
	policy       string
)

func getArtifacts(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		label    string
		system   models.System
		systems  []models.System
	)

	systems, err = filterSystems(config, activeProfileName, systemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/artifacts", config.URL, system.ID)
		label = fmt.Sprintf("system %d", system.ID)
		err = emass.FetchAndPrint(system.ConfigProfile, endpoint, label, outputFormat)
		if err != nil {
			return err
		}
	}

	return nil
}

func getControls(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		label    string
		params   url.Values
		system   models.System
		systems  []models.System
	)

	systems, err = filterSystems(config, activeProfileName, systemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/controls", config.URL, system.ID)
		params = url.Values{}
		params.Set("acronyms", strings.Join(controlIDs, ","))
		if len(params) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
		}

		label = fmt.Sprintf("system %d", system.ID)
		err = emass.FetchAndPrint(system.ConfigProfile, endpoint, label, outputFormat)
		if err != nil {
			return err
		}
	}

	return nil
}

func getApprovals(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		label    string
		system   models.System
		systems  []models.System
	)

	systems, err = filterSystems(config, activeProfileName, systemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/approval/cac", config.URL, system.ID)
		label = fmt.Sprintf("system %d", system.ID)
		err = emass.FetchAndPrint(system.ConfigProfile, endpoint, label, outputFormat)
		if err != nil {
			return err
		}
	}

	return nil
}

func getResults(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		label    string
		params   url.Values
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
	)

	profiles, err = filterProfiles(config, activeProfileName)
	if err != nil {
		return err
	}

	for _, profile = range profiles {
		endpoint = fmt.Sprintf("%s/api/test-results", config.URL)
		params = url.Values{}
		if len(controlIDs) != 0 {
			params.Set("controlAcronyms", strings.Join(controlIDs, ","))
		}
		if len(assessmentProcedures) != 0 {
			params.Set("assessmentProcedures", strings.Join(assessmentProcedures, ","))
		}
		if len(params) > 0 {
			endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
		}

		label = fmt.Sprintf("profile %s", profile.Name)
		err = emass.FetchAndPrint(profile, endpoint, label, outputFormat)
		if err != nil {
			return err
		}
	}

	return nil
}

func getRoles(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		label    string
		params   url.Values
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
	)

	profiles, err = filterProfiles(config, activeProfileName)
	if err != nil {
		return err
	}

	for _, profile = range profiles {
		endpoint = fmt.Sprintf("%s/api/system-roles", config.URL)
		label = fmt.Sprintf("profile %s", profile.Name)

		if roleCategory != "" {
			if role == "" {
				return fmt.Errorf("profile %s: a category and role are required", profile.Name)
			}
			endpoint = fmt.Sprintf("%s/%s", endpoint, roleCategory)
			params = url.Values{}
			params.Set("role", role)
			if policy != "" {
				params.Set("policy", policy)
			}
			if len(params) > 0 {
				endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
			}
		}

		label = fmt.Sprintf("profile %s", profile.Name)
		err = emass.FetchAndPrint(profile, endpoint, label, outputFormat)
		if err != nil {
			return err
		}
	}

	return nil
}

func getSystems(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		label    string
		profile  models.ConfigProfile
		profiles []models.ConfigProfile
		system   models.System
		systems  []models.System
	)

	if len(systemIDs) != 0 {
		systems, err = filterSystems(config, activeProfileName, systemIDs)
		if err != nil {
			return err
		}

		for _, system = range systems {
			endpoint = fmt.Sprintf("%s/api/systems/%d", config.URL, system.ID)
			label = fmt.Sprintf("system %d", system.ID)
			err = emass.FetchAndPrint(system.ConfigProfile, endpoint, label, outputFormat)
			if err != nil {
				return err
			}
		}
	} else {
		profiles, err = filterProfiles(config, activeProfileName)
		if err != nil {
			return err
		}

		for _, profile = range profiles {
			endpoint = fmt.Sprintf("%s/api/systems", config.URL)
			label = fmt.Sprintf("profile %s", profile.Name)
			err = emass.FetchAndPrint(profile, endpoint, label, outputFormat)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getWorkflows(cmd *cobra.Command, args []string) error {
	var (
		endpoint string
		err      error
		label    string
		system   models.System
		systems  []models.System
	)

	systems, err = filterSystems(config, activeProfileName, systemIDs)
	if err != nil {
		return err
	}

	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/approval/pac", config.URL, system.ID)
		label = fmt.Sprintf("system %d", system.ID)
		err = emass.FetchAndPrint(system.ConfigProfile, endpoint, label, outputFormat)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	getControlsCmd.PersistentFlags().StringSliceVarP(&controlIDs, "control-ids", "", []string{}, "Control IDs")

	getResultsCmd.PersistentFlags().StringSliceVarP(&controlIDs, "control-ids", "", []string{}, "Control IDs")
	getResultsCmd.PersistentFlags().StringSliceVarP(&assessmentProcedures, "assessment-procedures", "", []string{}, "Assessment procedures")

	getRolesCmd.Flags().StringVarP(&roleCategory, "category", "", "", "PAC, CAC, or Other")
	getRolesCmd.Flags().StringVarP(&role, "role", "", "", "ISO, ISSM, SCA, Auditor, AO, etc. (required if --category is used)")
	getRolesCmd.Flags().StringVarP(&policy, "policy", "", "", "RMF, DIACAP, or Reporting")

	getCmd.AddCommand(getApprovalsCmd)
	getCmd.AddCommand(getArtifactsCmd)
	getCmd.AddCommand(getControlsCmd)
	getCmd.AddCommand(getResultsCmd)
	getCmd.AddCommand(getRolesCmd)
	getCmd.AddCommand(getSystemsCmd)
	getCmd.AddCommand(getWorkflowsCmd)

	rootCmd.AddCommand(getCmd)
}
