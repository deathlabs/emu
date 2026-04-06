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
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deathlabs/emu/client"
	"github.com/deathlabs/emu/output"
	"github.com/spf13/cobra"
)

var (
	systemID int
	getCmd   = &cobra.Command{
		Use:   "get",
		Short: "Get evidence from eMASS",
	}
)

func getArtifacts(cmd *cobra.Command, args []string) {
	var (
		body       []byte
		err        error
		httpClient *http.Client
		parsedBody interface{}
		response   *http.Response
		url        string
	)

	httpClient, err = client.New(config.Systems[0].ConfigProfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	url = fmt.Sprintf("%s/api/systems/%d/artifacts", config.URL, systemID)

	response, err = client.Get(httpClient, config.Systems[0].ConfigProfile, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = output.Response(parsedBody.(map[string]interface{})["data"], outputFormat)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func getSystems(cmd *cobra.Command, args []string) {
	var (
		body       []byte
		err        error
		httpClient *http.Client
		parsedBody interface{}
		response   *http.Response
		url        string
	)

	httpClient, err = client.New(config.Systems[0].ConfigProfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	url = fmt.Sprintf("%s/api/systems", config.URL)

	response, err = client.Get(httpClient, config.Systems[0].ConfigProfile, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = output.Response(parsedBody.(map[string]interface{})["data"], outputFormat)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	getCmd.PersistentFlags().IntVarP(&systemID, "system-id", "s", 0, "System ID")
	getCmd.AddCommand(&cobra.Command{
		Use:   "artifacts",
		Short: "Get data about one or more artifacts",
		Run:   getArtifacts,
	})
	getCmd.AddCommand(&cobra.Command{
		Use:   "systems",
		Short: "Get data about one or more systems",
		Run:   getSystems,
	})
	rootCmd.AddCommand(getCmd)
}
