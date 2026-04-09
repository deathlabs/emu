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
package output

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Response(response *http.Response, format string) error {
	var (
		body     []byte
		jsonBody interface{}
		data     interface{}
		err      error
	)

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return err
	}

	// If there is no data, do not print anything.
	data = jsonBody.(map[string]interface{})["data"]
	if data == nil {
		return nil
	}

	switch strings.ToLower(format) {
	case "json":
		ToJSON(data)
		return nil
	case "yaml":
		ToYAML(data)
		return nil
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
}
