// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelRequestType Widget request type.
type FunnelRequestType string

// List of FunnelRequestType.
const (
	FUNNELREQUESTTYPE_FUNNEL FunnelRequestType = "funnel"
)

var allowedFunnelRequestTypeEnumValues = []FunnelRequestType{
	FUNNELREQUESTTYPE_FUNNEL,
}

// GetAllowedValues returns the list of possible values.
func (v *FunnelRequestType) GetAllowedValues() []FunnelRequestType {
	return allowedFunnelRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FunnelRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelRequestType(value)
	return nil
}

// NewFunnelRequestTypeFromValue returns a pointer to a valid FunnelRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFunnelRequestTypeFromValue(v string) (*FunnelRequestType, error) {
	ev := FunnelRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FunnelRequestType: valid values are %v", v, allowedFunnelRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FunnelRequestType) IsValid() bool {
	for _, existing := range allowedFunnelRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelRequestType value.
func (v FunnelRequestType) Ptr() *FunnelRequestType {
	return &v
}
