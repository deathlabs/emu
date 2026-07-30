package models

type CloudResourceResult struct {
	Provider          string             `json:"provider" yaml:"provider"`
	ResourceId        string             `json:"resourceId" yaml:"resourceId"`
	ResourceName      string             `json:"resourceName" yaml:"resourceName"`
	ResourceType      string             `json:"resourceType" yaml:"resourceType"`
	InitiatedBy       string             `json:"initiatedBy" yaml:"initiatedBy"`
	CspAccountId      string             `json:"cspAccountId" yaml:"cspAccountId"`
	CspRegion         string             `json:"cspRegion" yaml:"cspRegion"`
	IsBaseline        bool               `json:"isBaseline" yaml:"isBaseline"`
	Tags              map[string]string  `json:"tags" yaml:"tags"`
	ComplianceResults []ComplianceResult `json:"complianceResults" yaml:"complianceResults"`
}
