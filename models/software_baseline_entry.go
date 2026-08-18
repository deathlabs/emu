package models

// SoftwareBaselineEntry represents a single row in a software baseline CSV file.
type SoftwareBaselineEntry struct {
	ApprovalDate                 *int64  `csv:"approvalDate"                    json:"approvalDate,omitempty"         yaml:"approvalDate,omitempty"`
	ApprovalStatus               string  `csv:"approvalStatus"                  json:"approvalStatus"                 yaml:"approvalStatus"`
	CostPerLicense               float64 `csv:"costPerLicense"                  json:"costPerLicense"                 yaml:"costPerLicense"`
	CriticalAsset                bool    `csv:"criticalAsset"                   json:"criticalAsset"                  yaml:"criticalAsset"`
	CryptographicHash            string  `csv:"cryptographicHash"               json:"cryptographicHash"              yaml:"cryptographicHash"`
	EndOfLifeSupportDate         int64   `csv:"endOfLifeSupportDate"            json:"endOfLifeSupportDate"           yaml:"endOfLifeSupportDate"`
	ExtendedEndOfLifeSupportDate int64   `csv:"extendedEndOfLifeSupportDate"    json:"extendedEndOfLifeSupportDate"   yaml:"extendedEndOfLifeSupportDate"`
	FiscalYear                   string  `csv:"fiscalYear"                      json:"fiscalYear"                     yaml:"fiscalYear"`
	HostingEnvironment           string  `csv:"hostingEnvironment"              json:"hostingEnvironment"             yaml:"hostingEnvironment"`
	InServiceDate                string  `csv:"inServiceDate"                   json:"inServiceDate"                  yaml:"inServiceDate"`
	ItBudgetUii                  string  `csv:"itBudgetUii"                     json:"itBudgetUii"                    yaml:"itBudgetUii"`
	LicenseExpirationDate        int64   `csv:"licenseExpirationDate"           json:"licenseExpirationDate"          yaml:"licenseExpirationDate"`
	LicenseOrContract            string  `csv:"licenseOrContract"               json:"licenseOrContract"              yaml:"licenseOrContract"`
	LicensePoc                   string  `csv:"licensePoc"                      json:"licensePoc"                     yaml:"licensePoc"`
	LicenseRenewalDate           int64   `csv:"licenseRenewalDate"              json:"licenseRenewalDate"             yaml:"licenseRenewalDate"`
	LicenseTerm                  string  `csv:"licenseTerm"                     json:"licenseTerm"                    yaml:"licenseTerm"`
	LicensesUsed                 int     `csv:"licensesUsed"                    json:"licensesUsed"                   yaml:"licensesUsed"`
	Location                     string  `csv:"location"                        json:"location"                       yaml:"location"`
	MaintenanceDate              int64   `csv:"maintenanceDate"                 json:"maintenanceDate"                yaml:"maintenanceDate"`
	Network                      string  `csv:"network"                         json:"network"                        yaml:"network"`
	ParentSystem                 string  `csv:"parentSystem"                    json:"parentSystem"                   yaml:"parentSystem"`
	PopEndDate                   int64   `csv:"popEndDate"                      json:"popEndDate"                     yaml:"popEndDate"`
	Purpose                      string  `csv:"purpose"                         json:"purpose"                        yaml:"purpose"`
	ReleaseDate                  int64   `csv:"releaseDate"                     json:"releaseDate"                    yaml:"releaseDate"`
	RetirementDate               int64   `csv:"retirementDate"                  json:"retirementDate"                 yaml:"retirementDate"`
	SoftwareDependencies         string  `csv:"softwareDependencies"            json:"softwareDependencies"           yaml:"softwareDependencies"`
	SoftwareName                 string  `csv:"softwareName"                    json:"softwareName"                   yaml:"softwareName"`
	SoftwareType                 string  `csv:"softwareType"                    json:"softwareType"                   yaml:"softwareType"`
	SoftwareVendor               string  `csv:"softwareVendor"                  json:"softwareVendor"                 yaml:"softwareVendor"`
	Subsystem                    string  `csv:"subsystem"                       json:"subsystem"                      yaml:"subsystem"`
	TotalLicenseCost             float64 `csv:"totalLicenseCost"                json:"totalLicenseCost"               yaml:"totalLicenseCost"`
	TotalLicenses                int     `csv:"totalLicenses"                   json:"totalLicenses"                  yaml:"totalLicenses"`
	Version                      string  `csv:"version"                         json:"version"                        yaml:"version"`
}
