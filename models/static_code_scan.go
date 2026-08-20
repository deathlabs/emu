package models

type StaticCodeScan struct {
	Application struct {
		ApplicationName string `json:"applicationName" yaml:"applicationName"`
		Version         string `json:"version" yaml:"version"`
	} `json:"application" yaml:"application"`
	ApplicationFindings []struct {
		RawSeverity   string `json:"rawSeverity" yaml:"rawSeverity"`
		CweId         string `json:"cweId" yaml:"cweId"`
		ScanDate      int64  `json:"scanDate" yaml:"scanDate"`
		CodeCheckName string `json:"codeCheckName" yaml:"codeCheckName"`
		Count         int    `json:"count" yaml:"count"`
		ScanType      string `json:"scanType" yaml:"scanType"`
	} `json:"applicationFindings" yaml:"applicationFindings"`
}
