package models

import (
	"fmt"
	"strings"
)

type DeviceScanType string

var deviceScanTypes = []string{
	"acasAsrArf",
	"acasConsolidatedArf",
	"acasNessus",
	"disaStigViewerCklCklb",
	"disaStigViewerCmrs",
	"policyAuditor",
	"scapComplianceChecker",
}

func (deviceScanType *DeviceScanType) String() string {
	return string(*deviceScanType)
}

func (deviceScanType *DeviceScanType) Set(value string) error {
	var valid string

	for _, valid = range deviceScanTypes {
		if value == valid {
			*deviceScanType = DeviceScanType(value)
			return nil
		}
	}

	return fmt.Errorf(
		"must be one of: %s",
		strings.Join(deviceScanTypes, ", "),
	)
}

func (deviceScanType *DeviceScanType) Type() string {
	return "deviceScanType"
}
