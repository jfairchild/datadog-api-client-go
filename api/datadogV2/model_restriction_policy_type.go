// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionPolicyType Restriction policy type.
type RestrictionPolicyType string

// List of RestrictionPolicyType.
const (
	RESTRICTIONPOLICYTYPE_RESTRICTION_POLICY RestrictionPolicyType = "restriction_policy"
)

var allowedRestrictionPolicyTypeEnumValues = []RestrictionPolicyType{
	RESTRICTIONPOLICYTYPE_RESTRICTION_POLICY,
}

// GetAllowedValues returns the list of possible values.
func (v *RestrictionPolicyType) GetAllowedValues() []RestrictionPolicyType {
	return allowedRestrictionPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RestrictionPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RestrictionPolicyType(value)
	return nil
}

// NewRestrictionPolicyTypeFromValue returns a pointer to a valid RestrictionPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRestrictionPolicyTypeFromValue(v string) (*RestrictionPolicyType, error) {
	ev := RestrictionPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RestrictionPolicyType: valid values are %v", v, allowedRestrictionPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RestrictionPolicyType) IsValid() bool {
	for _, existing := range allowedRestrictionPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RestrictionPolicyType value.
func (v RestrictionPolicyType) Ptr() *RestrictionPolicyType {
	return &v
}
