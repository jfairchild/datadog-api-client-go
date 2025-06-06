// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeWidgetDefinitionType Type of the change widget.
type ChangeWidgetDefinitionType string

// List of ChangeWidgetDefinitionType.
const (
	CHANGEWIDGETDEFINITIONTYPE_CHANGE ChangeWidgetDefinitionType = "change"
)

var allowedChangeWidgetDefinitionTypeEnumValues = []ChangeWidgetDefinitionType{
	CHANGEWIDGETDEFINITIONTYPE_CHANGE,
}

// GetAllowedValues returns the list of possible values.
func (v *ChangeWidgetDefinitionType) GetAllowedValues() []ChangeWidgetDefinitionType {
	return allowedChangeWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeWidgetDefinitionType(value)
	return nil
}

// NewChangeWidgetDefinitionTypeFromValue returns a pointer to a valid ChangeWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeWidgetDefinitionTypeFromValue(v string) (*ChangeWidgetDefinitionType, error) {
	ev := ChangeWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeWidgetDefinitionType: valid values are %v", v, allowedChangeWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedChangeWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeWidgetDefinitionType value.
func (v ChangeWidgetDefinitionType) Ptr() *ChangeWidgetDefinitionType {
	return &v
}
