package types

type RequestConfig struct {
	Method       string            `json:"method"`
	URL          string            `json:"url"`
	Headers      map[string]string `json:"headers"`
	Body         string            `json:"body,omitempty"`
	Timeout      int               `json:"timeout"`
	OutputFormat string            `json:"output_format"`
}
