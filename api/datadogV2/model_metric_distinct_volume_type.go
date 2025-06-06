// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricDistinctVolumeType The metric distinct volume type.
type MetricDistinctVolumeType string

// List of MetricDistinctVolumeType.
const (
	METRICDISTINCTVOLUMETYPE_DISTINCT_METRIC_VOLUMES MetricDistinctVolumeType = "distinct_metric_volumes"
)

var allowedMetricDistinctVolumeTypeEnumValues = []MetricDistinctVolumeType{
	METRICDISTINCTVOLUMETYPE_DISTINCT_METRIC_VOLUMES,
}

// GetAllowedValues returns the list of possible values.
func (v *MetricDistinctVolumeType) GetAllowedValues() []MetricDistinctVolumeType {
	return allowedMetricDistinctVolumeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricDistinctVolumeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricDistinctVolumeType(value)
	return nil
}

// NewMetricDistinctVolumeTypeFromValue returns a pointer to a valid MetricDistinctVolumeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricDistinctVolumeTypeFromValue(v string) (*MetricDistinctVolumeType, error) {
	ev := MetricDistinctVolumeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricDistinctVolumeType: valid values are %v", v, allowedMetricDistinctVolumeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricDistinctVolumeType) IsValid() bool {
	for _, existing := range allowedMetricDistinctVolumeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricDistinctVolumeType value.
func (v MetricDistinctVolumeType) Ptr() *MetricDistinctVolumeType {
	return &v
}
