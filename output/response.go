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
		body       []byte
		err        error
		parsedBody interface{}
	)

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		return err
	}

	switch strings.ToLower(format) {
	case "json":
		ToJSON(parsedBody.(map[string]interface{})["data"])
		return nil
	case "yaml":
		ToYAML(parsedBody.(map[string]interface{})["data"])
		return nil
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
}
