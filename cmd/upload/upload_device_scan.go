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
package upload

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/deathlabs/emu/v4/config"
	"github.com/deathlabs/emu/v4/emass"
	"github.com/deathlabs/emu/v4/models"
	"github.com/deathlabs/emu/v4/output"
	"github.com/spf13/cobra"
)

var (
	deviceScanPath       string
	deviceScanType       models.DeviceScanType
	deviceScanIsBaseline bool
)

var (
	uploadDeviceScanCmd = &cobra.Command{
		Use:   "device-scan",
		Short: "Upload a device scan to eMASS",
		RunE:  uploadDeviceScan,
	}
)

func uploadDeviceScan(cmd *cobra.Command, args []string) error {
	var (
		body            bytes.Buffer
		endpoint        string
		err             error
		file            *os.File
		fileWriter      io.Writer
		headers         map[string]string
		queryParameters url.Values
		response        *http.Response
		system          models.System
		systems         []models.System
		writer          *multipart.Writer
	)

	systems, err = config.FilterSystems(config.Data, config.ActiveProfileName, config.SystemIDs)
	if err != nil {
		return err
	}

	if deviceScanPath == "" {
		return errors.New("an argument for the 'file' parameter is required")
	}

	if deviceScanType == "" {
		return errors.New("an argument for the 'device-scan-type' parameter is required")
	}

	for _, system = range systems {
		writer = multipart.NewWriter(&body)

		fileWriter, err = writer.CreateFormFile("file", filepath.Base(deviceScanPath))
		if err != nil {
			return err
		}

		file, err = os.Open(deviceScanPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return err
		}

		err = file.Close()
		if err != nil {
			return err
		}

		err = writer.Close()
		if err != nil {
			return err
		}

		queryParameters = url.Values{}
		queryParameters.Set("scanType", string(deviceScanType))
		queryParameters.Set("isBaseline", strconv.FormatBool(deviceScanIsBaseline))
		endpoint = fmt.Sprintf("%s/api/systems/%d/device-scan-results?%s", config.Data.URL, system.ID, queryParameters.Encode())

		headers = map[string]string{
			"Content-Type": writer.FormDataContentType(),
		}

		response, err = emass.Post(system.ConfigProfile, endpoint, headers, &body)
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
	uploadDeviceScanCmd.PersistentFlags().StringVarP(&deviceScanPath, "file", "f", "", "File path to the device scan")
	uploadDeviceScanCmd.PersistentFlags().VarP(&deviceScanType, "device-scan-type", "t", "Device scan type (acasAsrArf | acasConsolidatedArf | acasNessus | disaStigViewerCklCklb | disaStigViewerCmrs | policyAuditor | scapComplianceChecker)")
	uploadDeviceScanCmd.PersistentFlags().BoolVarP(&deviceScanIsBaseline, "is-baseline", "", false, "Utilize this parameter if the imported file represents a baseline scan that includes all findings and results. Importing as a baseline scan, which assumes a common set of scan policies are used when conducting a scan, will replace a device's findings for a specific benchmark.")
}
