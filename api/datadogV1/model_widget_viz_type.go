// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetVizType Whether to display the Alert Graph as a timeseries or a top list.
type WidgetVizType string

// List of WidgetVizType.
const (
	WIDGETVIZTYPE_TIMESERIES WidgetVizType = "timeseries"
	WIDGETVIZTYPE_TOPLIST    WidgetVizType = "toplist"
)

var allowedWidgetVizTypeEnumValues = []WidgetVizType{
	WIDGETVIZTYPE_TIMESERIES,
	WIDGETVIZTYPE_TOPLIST,
}

// GetAllowedValues returns the list of possible values.
func (v *WidgetVizType) GetAllowedValues() []WidgetVizType {
	return allowedWidgetVizTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetVizType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetVizType(value)
	return nil
}

// NewWidgetVizTypeFromValue returns a pointer to a valid WidgetVizType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetVizTypeFromValue(v string) (*WidgetVizType, error) {
	ev := WidgetVizType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetVizType: valid values are %v", v, allowedWidgetVizTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetVizType) IsValid() bool {
	for _, existing := range allowedWidgetVizTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetVizType value.
func (v WidgetVizType) Ptr() *WidgetVizType {
	return &v
}
