// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricCustomTimeAggregation A time aggregation for use in query.
type MetricCustomTimeAggregation string

// List of MetricCustomTimeAggregation.
const (
	METRICCUSTOMTIMEAGGREGATION_AVG   MetricCustomTimeAggregation = "avg"
	METRICCUSTOMTIMEAGGREGATION_COUNT MetricCustomTimeAggregation = "count"
	METRICCUSTOMTIMEAGGREGATION_MAX   MetricCustomTimeAggregation = "max"
	METRICCUSTOMTIMEAGGREGATION_MIN   MetricCustomTimeAggregation = "min"
	METRICCUSTOMTIMEAGGREGATION_SUM   MetricCustomTimeAggregation = "sum"
)

var allowedMetricCustomTimeAggregationEnumValues = []MetricCustomTimeAggregation{
	METRICCUSTOMTIMEAGGREGATION_AVG,
	METRICCUSTOMTIMEAGGREGATION_COUNT,
	METRICCUSTOMTIMEAGGREGATION_MAX,
	METRICCUSTOMTIMEAGGREGATION_MIN,
	METRICCUSTOMTIMEAGGREGATION_SUM,
}

// GetAllowedValues returns the list of possible values.
func (v *MetricCustomTimeAggregation) GetAllowedValues() []MetricCustomTimeAggregation {
	return allowedMetricCustomTimeAggregationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricCustomTimeAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricCustomTimeAggregation(value)
	return nil
}

// NewMetricCustomTimeAggregationFromValue returns a pointer to a valid MetricCustomTimeAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricCustomTimeAggregationFromValue(v string) (*MetricCustomTimeAggregation, error) {
	ev := MetricCustomTimeAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricCustomTimeAggregation: valid values are %v", v, allowedMetricCustomTimeAggregationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricCustomTimeAggregation) IsValid() bool {
	for _, existing := range allowedMetricCustomTimeAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricCustomTimeAggregation value.
func (v MetricCustomTimeAggregation) Ptr() *MetricCustomTimeAggregation {
	return &v
}
