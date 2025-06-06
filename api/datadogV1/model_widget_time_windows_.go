// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetTimeWindows Define a time window.
type WidgetTimeWindows string

// List of WidgetTimeWindows.
const (
	WIDGETTIMEWINDOWS_SEVEN_DAYS     WidgetTimeWindows = "7d"
	WIDGETTIMEWINDOWS_THIRTY_DAYS    WidgetTimeWindows = "30d"
	WIDGETTIMEWINDOWS_NINETY_DAYS    WidgetTimeWindows = "90d"
	WIDGETTIMEWINDOWS_WEEK_TO_DATE   WidgetTimeWindows = "week_to_date"
	WIDGETTIMEWINDOWS_PREVIOUS_WEEK  WidgetTimeWindows = "previous_week"
	WIDGETTIMEWINDOWS_MONTH_TO_DATE  WidgetTimeWindows = "month_to_date"
	WIDGETTIMEWINDOWS_PREVIOUS_MONTH WidgetTimeWindows = "previous_month"
	WIDGETTIMEWINDOWS_GLOBAL_TIME    WidgetTimeWindows = "global_time"
)

var allowedWidgetTimeWindowsEnumValues = []WidgetTimeWindows{
	WIDGETTIMEWINDOWS_SEVEN_DAYS,
	WIDGETTIMEWINDOWS_THIRTY_DAYS,
	WIDGETTIMEWINDOWS_NINETY_DAYS,
	WIDGETTIMEWINDOWS_WEEK_TO_DATE,
	WIDGETTIMEWINDOWS_PREVIOUS_WEEK,
	WIDGETTIMEWINDOWS_MONTH_TO_DATE,
	WIDGETTIMEWINDOWS_PREVIOUS_MONTH,
	WIDGETTIMEWINDOWS_GLOBAL_TIME,
}

// GetAllowedValues returns the list of possible values.
func (v *WidgetTimeWindows) GetAllowedValues() []WidgetTimeWindows {
	return allowedWidgetTimeWindowsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetTimeWindows) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetTimeWindows(value)
	return nil
}

// NewWidgetTimeWindowsFromValue returns a pointer to a valid WidgetTimeWindows
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetTimeWindowsFromValue(v string) (*WidgetTimeWindows, error) {
	ev := WidgetTimeWindows(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetTimeWindows: valid values are %v", v, allowedWidgetTimeWindowsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetTimeWindows) IsValid() bool {
	for _, existing := range allowedWidgetTimeWindowsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetTimeWindows value.
func (v WidgetTimeWindows) Ptr() *WidgetTimeWindows {
	return &v
}
