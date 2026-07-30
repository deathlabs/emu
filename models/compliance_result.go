package models

type ComplianceResult struct {
	CspPolicyDefinitionId    string `json:"cspPolicyDefinitionId" yaml:"cspPolicyDefinitionId"`
	PolicyDefinitionTitle    string `json:"policyDefinitionTitle" yaml:"policyDefinitionTitle"`
	ComplianceCheckTimestamp int64  `json:"complianceCheckTimestamp" yaml:"complianceCheckTimestamp"`
	IsCompliant              bool   `json:"isCompliant" yaml:"isCompliant"`
	Control                  string `json:"control" yaml:"control"`
	AssessmentProcedure      string `json:"assessmentProcedure" yaml:"assessmentProcedure"`
	ComplianceReason         string `json:"complianceReason" yaml:"complianceReason"`
	PolicyDeploymentName     string `json:"policyDeploymentName" yaml:"policyDeploymentName"`
	PolicyDeploymentVersion  string `json:"policyDeploymentVersion" yaml:"policyDeploymentVersion"`
	Severity                 string `json:"severity" yaml:"severity"`
}
