// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetComparator Comparator to apply.
type WidgetComparator string

// List of WidgetComparator.
const (
	WIDGETCOMPARATOR_EQUAL_TO                 WidgetComparator = "="
	WIDGETCOMPARATOR_GREATER_THAN             WidgetComparator = ">"
	WIDGETCOMPARATOR_GREATER_THAN_OR_EQUAL_TO WidgetComparator = ">="
	WIDGETCOMPARATOR_LESS_THAN                WidgetComparator = "<"
	WIDGETCOMPARATOR_LESS_THAN_OR_EQUAL_TO    WidgetComparator = "<="
)

var allowedWidgetComparatorEnumValues = []WidgetComparator{
	WIDGETCOMPARATOR_EQUAL_TO,
	WIDGETCOMPARATOR_GREATER_THAN,
	WIDGETCOMPARATOR_GREATER_THAN_OR_EQUAL_TO,
	WIDGETCOMPARATOR_LESS_THAN,
	WIDGETCOMPARATOR_LESS_THAN_OR_EQUAL_TO,
}

// GetAllowedValues returns the list of possible values.
func (v *WidgetComparator) GetAllowedValues() []WidgetComparator {
	return allowedWidgetComparatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetComparator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetComparator(value)
	return nil
}

// NewWidgetComparatorFromValue returns a pointer to a valid WidgetComparator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetComparatorFromValue(v string) (*WidgetComparator, error) {
	ev := WidgetComparator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetComparator: valid values are %v", v, allowedWidgetComparatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetComparator) IsValid() bool {
	for _, existing := range allowedWidgetComparatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetComparator value.
func (v WidgetComparator) Ptr() *WidgetComparator {
	return &v
}
