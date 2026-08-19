package models

type POAM struct {
	PoamID                      int      `json:"poamId" yaml:"poamId"`
	ControlAcronym              string   `json:"controlAcronym" yaml:"controlAcronym"`
	AssessmentProcedure         string   `json:"assessmentProcedure" yaml:"assessmentProcedure"`
	Severity                    string   `json:"severity" yaml:"severity"`
	RawSeverity                 string   `json:"rawSeverity" yaml:"rawSeverity"`
	Status                      string   `json:"status" yaml:"status"`
	ScheduledCompletionDate     int64    `json:"scheduledCompletionDate" yaml:"scheduledCompletionDate"`
	CompletionDate              *int64   `json:"completionDate" yaml:"completionDate"`
	PocOrganization             string   `json:"pocOrganization" yaml:"pocOrganization"`
	PocLastName                 string   `json:"pocLastName" yaml:"pocLastName"`
	PocFirstName                string   `json:"pocFirstName" yaml:"pocFirstName"`
	PocEmail                    string   `json:"pocEmail" yaml:"pocEmail"`
	PocPhoneNumber              string   `json:"pocPhoneNumber" yaml:"pocPhoneNumber"`
	VulnerabilityDescription    string   `json:"vulnerabilityDescription" yaml:"vulnerabilityDescription"`
	Mitigations                 string   `json:"mitigations" yaml:"mitigations"`
	Comments                    string   `json:"comments" yaml:"comments"`
	Resources                   string   `json:"resources" yaml:"resources"`
	IdentificationSource        []string `json:"identificationSource" yaml:"identificationSource"`
	IdentificationSourceDetails string   `json:"identificationSourceDetails" yaml:"identificationSourceDetails"`
	SecurityChecks              string   `json:"securityChecks" yaml:"securityChecks"`
	Recommendations             string   `json:"recommendations" yaml:"recommendations"`
	RelevanceOfThreat           string   `json:"relevanceOfThreat" yaml:"relevanceOfThreat"`
	Likelihood                  string   `json:"likelihood" yaml:"likelihood"`
	Impact                      string   `json:"impact" yaml:"impact"`
	ImpactDescription           string   `json:"impactDescription" yaml:"impactDescription"`
	ResidualRiskLevel           string   `json:"residualRiskLevel" yaml:"residualRiskLevel"`
	Milestones                  []struct {
		Description             string `json:"description" yaml:"description"`
		ScheduledCompletionDate int64  `json:"scheduledCompletionDate" yaml:"scheduledCompletionDate"`
		Status                  string `json:"status" yaml:"status"`
		StatusComments          string `json:"statusComments" yaml:"statusComments"`
		CompletionDate          *int64 `json:"completionDate" yaml:"completionDate"`
		IsActive                bool   `json:"isActive" yaml:"isActive"`
	} `json:"milestones" yaml:"milestones"`
}
