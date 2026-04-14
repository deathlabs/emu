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
package emass

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/deathlabs/emu/models"
)

func getHTTPStatusCodeDescription(response *http.Response) string {
	switch response.StatusCode {
	case 200:
		return "request succeeded (applicable to partial successes and the response body depends on the request method)"
	case 201:
		return "request was fulfilled and resulted in one or more new resources being successfully created on the server."
	case 400:
		return "request could not be understood by the server due to incorrect syntax or an unexpected format."
	case 401:
		return "request failed to provide suitable authentication"
	case 403:
		return "request was blocked due to a lack of client permissions for the API or to a specific API endpoint"
	case 404:
		return "request failed because the URL provided did not match any available API endpoints"
	case 405:
		return "request was made with a verb (GET, POST, etc.) that is not permitted for the API endpoint specified"
	case 411:
		return "request was of type POST and failed to provide the server information about the data/content length being submitted"
	case 409:
		return "request failed because too much data was requested in a single batch (this error is specific to eMASS)"
	case 500:
		return "server encountered an unexpected condition that prevented it from fulfilling the request"
	default:
		return fmt.Sprintf("Unknown HTTP status code: %d", response.StatusCode)
	}
}

// Get sends an HTTP GET request using the endpoint and profile specified and returns the HTTP response.
func Get(profile models.ConfigProfile, endpoint string) (*http.Response, error) {
	var (
		client   *http.Client
		err      error
		request  *http.Request
		response *http.Response
	)

	// Get an HTTPS client for the specified profile.
	client, err = getHTTPClient(profile)
	if err != nil {
		return nil, err
	}

	// Create an HTTP request for the specified endpoint.
	request, err = http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Set the headers required for the HTTP request.
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("api-key", profile.APIKey)
	request.Header.Set("user-uid", profile.UserUID)

	// Send the HTTP request.
	response, err = client.Do(request)
	if err != nil {
		return nil, err
	}

	// Check the HTTP response status code.
	if response.StatusCode != 200 && response.StatusCode != 201 {
		return nil, fmt.Errorf("HTTP status code %d: %s", response.StatusCode, getHTTPStatusCodeDescription(response))
	}

	// Return the HTTP response.
	return response, nil
}

// Post sends an HTTP POST request using the endpoint, profile, and body specified and returns the HTTP response.
func Post(profile models.ConfigProfile, endpoint string, body *bytes.Buffer, contentType string) (*http.Response, error) {
	var (
		client   *http.Client
		err      error
		request  *http.Request
		response *http.Response
	)

	// Get an HTTPS client for the specified profile.
	client, err = getHTTPClient(profile)
	if err != nil {
		return nil, err
	}

	// Create an HTTP request for the specified endpoint.
	request, err = http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		return nil, err
	}

	// Set the headers required for the HTTP request.
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("api-key", profile.APIKey)
	request.Header.Set("user-uid", profile.UserUID)

	// Send the HTTP request.
	response, err = client.Do(request)
	if err != nil {
		return nil, err
	}

	// Check the HTTP response status code.
	if response.StatusCode != 200 && response.StatusCode != 201 {
		return nil, fmt.Errorf("HTTP status code %d: %s", response.StatusCode, getHTTPStatusCodeDescription(response))
	}
	// Return the HTTP response.
	return response, nil
}
