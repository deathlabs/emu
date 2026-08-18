package models

type CloudResourceResult struct {
	Provider          string            `json:"provider" yaml:"provider"`
	ResourceId        string            `json:"resourceId" yaml:"resourceId"`
	ResourceName      string            `json:"resourceName" yaml:"resourceName"`
	ResourceType      string            `json:"resourceType" yaml:"resourceType"`
	InitiatedBy       string            `json:"initiatedBy" yaml:"initiatedBy"`
	CspAccountId      string            `json:"cspAccountId" yaml:"cspAccountId"`
	CspRegion         string            `json:"cspRegion" yaml:"cspRegion"`
	IsBaseline        bool              `json:"isBaseline" yaml:"isBaseline"`
	Tags              map[string]string `json:"tags" yaml:"tags"`
	ComplianceResults []struct {
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
	} `json:"complianceResults" yaml:"complianceResults"`
}
