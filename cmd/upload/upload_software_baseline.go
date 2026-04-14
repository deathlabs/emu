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
package upload

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/deathlabs/emu/config"
	"github.com/deathlabs/emu/emass"
	"github.com/deathlabs/emu/models"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

var (
	softwareBaselineFileType string
	softwareBaselinePath     string
)

var (
	uploadSoftwareBaselineCmd = &cobra.Command{
		Use:   "software-baseline",
		Short: "Upload a software baseline to eMASS",
		RunE:  uploadSoftwareBaseline,
	}
)

func closeExcelFile(file *excelize.File) error {
	var err error

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func uploadSoftwareBaseline(cmd *cobra.Command, args []string) error {
	var (
		body                  *bytes.Buffer
		endpoint              string
		err                   error
		file                  *excelize.File
		jsonData              []byte
		response              *http.Response
		row                   []string
		rows                  [][]string
		softwareBaseline      []models.SoftwareBaselineEntry
		softwareBaselineEntry models.SoftwareBaselineEntry
		system                models.System
		systems               []models.System
	)

	// Check the argument provided for the "softwareBaselinePath" parameter.
	if softwareBaselinePath == "" {
		return errors.New("a file path for the software baseline was not provided")
	}

	// Open the XLSM file specified.
	file, err = excelize.OpenFile(softwareBaselinePath)
	if err != nil {
		return err
	}
	defer closeExcelFile(file)

	// Get all the rows from the "Software" tab in the XLSM file.
	rows, err = file.GetRows("Software")
	if err != nil {
		return err
	}

	// Build a software baseline based on all the rows collected.
	for _, row = range rows[7:] {
		softwareBaselineEntry = models.SoftwareBaselineEntry{
			SoftwareType:   row[1],
			SoftwareVendor: row[2],
			SoftwareName:   row[3],
			Version:        row[4],
			//ParentSystem:                 row[5],
			//Subsystem:                    row[6],
			//Network:                      row[7],
			//HostingEnvironment:           row[8],
			//SoftwareDependencies:         row[9],
			//CryptographicHash:            row[10],
			//InServiceDate:                row[11],
			//ItBudgetUii:                  row[12],
			//FiscalYear:                   row[13],
			//PopEndDate:                   row[14],
			//LicenseOrContract:            row[15],
			//LicenseTerm:                  row[16],
			//CostPerLicense:               row[17],
			//TotalLicenses:                row[18],
			//LicensesUsed:                 row[19],
			//LicensePoc:                   row[20],
			//LicenseRenewalDate:           row[21],
			//ApprovalStatus:               row[22],
			//ApprovalDate:                 row[23],
			//ReleaseDate:                  row[24],
			//MaintenanceDate:              row[25],
			//RetirementDate:               row[26],
			//EndOfLifeSupportDate:         row[27],
			//ExtendedEndOfLifeSupportDate: row[28],
			//CriticalAsset:                row[29],
			//Location:                     row[30],
			//Purpose:                      row[31],
		}
		softwareBaseline = append(softwareBaseline, softwareBaselineEntry)
	}

	// Convert the software baseline to JSON and then a sequence of bytes.
	jsonData, err = json.Marshal(softwareBaseline)
	if err != nil {
		return err
	}
	body = bytes.NewBuffer(jsonData)

	// Filter systems based on system IDs provided via the root-level --system-ids flag.
	// If no system IDs are provided, this will return all systems for the active profile.
	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	// Loop through the filtered systems and upload a software baseline to each one.
	for _, system = range systems {
		endpoint = fmt.Sprintf("%s/api/systems/%d/sw-baseline", config.Data.URL, system.ID)
		response, err = emass.Post(system.ConfigProfile, endpoint, body, "application/json")
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
	// Define flags for the "emu upload software-baseline" subcommand.
	uploadSoftwareBaselineCmd.PersistentFlags().StringVarP(&softwareBaselinePath, "file", "f", "", "File path to the software baseline")
}
