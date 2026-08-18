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
