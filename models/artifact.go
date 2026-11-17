package models

type ArtifactUpdate struct {
	AssessmentProcedures *string  `mapstructure:"assessmentProcedures" json:"assessmentProcedures,omitempty" yaml:"assessmentProcedures,omitempty"`
	Category             string   `mapstructure:"category"             json:"category"                      yaml:"category"`
	Controls             *string  `mapstructure:"controls"             json:"controls,omitempty"            yaml:"controls,omitempty"`
	Description          *string  `mapstructure:"description"          json:"description,omitempty"         yaml:"description,omitempty"`
	ExpirationDate       *float64 `mapstructure:"expirationDate"       json:"expirationDate,omitempty"      yaml:"expirationDate,omitempty"`
	Filename             string   `mapstructure:"filename"             json:"filename"                      yaml:"filename"`
	IsTemplate           bool     `mapstructure:"isTemplate"           json:"isTemplate"                    yaml:"isTemplate"`
	LastReviewedDate     *float64 `mapstructure:"lastReviewedDate"     json:"lastReviewedDate,omitempty"    yaml:"lastReviewedDate,omitempty"`
	Name                 *string  `mapstructure:"name"                 json:"name,omitempty"                yaml:"name,omitempty"`
	ReferencePageNumber  *string  `mapstructure:"referencePageNumber"  json:"referencePageNumber,omitempty" yaml:"referencePageNumber,omitempty"`
	SignedDate           *float64 `mapstructure:"signedDate"           json:"signedDate,omitempty"          yaml:"signedDate,omitempty"`
	Type                 string   `mapstructure:"type"                 json:"type"                          yaml:"type"`
}

type ArtifactDelete struct {
	Filename string `mapstructure:"filename" json:"filename" yaml:"filename"`
}
