// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetDefinitionType Type of the host map widget.
type HostMapWidgetDefinitionType string

// List of HostMapWidgetDefinitionType.
const (
	HOSTMAPWIDGETDEFINITIONTYPE_HOSTMAP HostMapWidgetDefinitionType = "hostmap"
)

var allowedHostMapWidgetDefinitionTypeEnumValues = []HostMapWidgetDefinitionType{
	HOSTMAPWIDGETDEFINITIONTYPE_HOSTMAP,
}

// GetAllowedValues returns the list of possible values.
func (v *HostMapWidgetDefinitionType) GetAllowedValues() []HostMapWidgetDefinitionType {
	return allowedHostMapWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostMapWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostMapWidgetDefinitionType(value)
	return nil
}

// NewHostMapWidgetDefinitionTypeFromValue returns a pointer to a valid HostMapWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostMapWidgetDefinitionTypeFromValue(v string) (*HostMapWidgetDefinitionType, error) {
	ev := HostMapWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostMapWidgetDefinitionType: valid values are %v", v, allowedHostMapWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostMapWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedHostMapWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostMapWidgetDefinitionType value.
func (v HostMapWidgetDefinitionType) Ptr() *HostMapWidgetDefinitionType {
	return &v
}
