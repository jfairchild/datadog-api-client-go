// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumMetricEventType The type of RUM events to filter on.
type RumMetricEventType string

// List of RumMetricEventType.
const (
	RUMMETRICEVENTTYPE_SESSION   RumMetricEventType = "session"
	RUMMETRICEVENTTYPE_VIEW      RumMetricEventType = "view"
	RUMMETRICEVENTTYPE_ACTION    RumMetricEventType = "action"
	RUMMETRICEVENTTYPE_ERROR     RumMetricEventType = "error"
	RUMMETRICEVENTTYPE_RESOURCE  RumMetricEventType = "resource"
	RUMMETRICEVENTTYPE_LONG_TASK RumMetricEventType = "long_task"
	RUMMETRICEVENTTYPE_VITAL     RumMetricEventType = "vital"
)

var allowedRumMetricEventTypeEnumValues = []RumMetricEventType{
	RUMMETRICEVENTTYPE_SESSION,
	RUMMETRICEVENTTYPE_VIEW,
	RUMMETRICEVENTTYPE_ACTION,
	RUMMETRICEVENTTYPE_ERROR,
	RUMMETRICEVENTTYPE_RESOURCE,
	RUMMETRICEVENTTYPE_LONG_TASK,
	RUMMETRICEVENTTYPE_VITAL,
}

// GetAllowedValues returns the list of possible values.
func (v *RumMetricEventType) GetAllowedValues() []RumMetricEventType {
	return allowedRumMetricEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumMetricEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumMetricEventType(value)
	return nil
}

// NewRumMetricEventTypeFromValue returns a pointer to a valid RumMetricEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumMetricEventTypeFromValue(v string) (*RumMetricEventType, error) {
	ev := RumMetricEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumMetricEventType: valid values are %v", v, allowedRumMetricEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumMetricEventType) IsValid() bool {
	for _, existing := range allowedRumMetricEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumMetricEventType value.
func (v RumMetricEventType) Ptr() *RumMetricEventType {
	return &v
}
