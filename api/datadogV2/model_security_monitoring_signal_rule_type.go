// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalRuleType The rule type.
type SecurityMonitoringSignalRuleType string

// List of SecurityMonitoringSignalRuleType.
const (
	SECURITYMONITORINGSIGNALRULETYPE_SIGNAL_CORRELATION SecurityMonitoringSignalRuleType = "signal_correlation"
)

var allowedSecurityMonitoringSignalRuleTypeEnumValues = []SecurityMonitoringSignalRuleType{
	SECURITYMONITORINGSIGNALRULETYPE_SIGNAL_CORRELATION,
}

// GetAllowedValues returns the list of possible values.
func (v *SecurityMonitoringSignalRuleType) GetAllowedValues() []SecurityMonitoringSignalRuleType {
	return allowedSecurityMonitoringSignalRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalRuleType(value)
	return nil
}

// NewSecurityMonitoringSignalRuleTypeFromValue returns a pointer to a valid SecurityMonitoringSignalRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalRuleTypeFromValue(v string) (*SecurityMonitoringSignalRuleType, error) {
	ev := SecurityMonitoringSignalRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalRuleType: valid values are %v", v, allowedSecurityMonitoringSignalRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalRuleType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalRuleType value.
func (v SecurityMonitoringSignalRuleType) Ptr() *SecurityMonitoringSignalRuleType {
	return &v
}
