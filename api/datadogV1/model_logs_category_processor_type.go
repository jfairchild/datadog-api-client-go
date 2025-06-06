// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsCategoryProcessorType Type of logs category processor.
type LogsCategoryProcessorType string

// List of LogsCategoryProcessorType.
const (
	LOGSCATEGORYPROCESSORTYPE_CATEGORY_PROCESSOR LogsCategoryProcessorType = "category-processor"
)

var allowedLogsCategoryProcessorTypeEnumValues = []LogsCategoryProcessorType{
	LOGSCATEGORYPROCESSORTYPE_CATEGORY_PROCESSOR,
}

// GetAllowedValues returns the list of possible values.
func (v *LogsCategoryProcessorType) GetAllowedValues() []LogsCategoryProcessorType {
	return allowedLogsCategoryProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsCategoryProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsCategoryProcessorType(value)
	return nil
}

// NewLogsCategoryProcessorTypeFromValue returns a pointer to a valid LogsCategoryProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsCategoryProcessorTypeFromValue(v string) (*LogsCategoryProcessorType, error) {
	ev := LogsCategoryProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsCategoryProcessorType: valid values are %v", v, allowedLogsCategoryProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsCategoryProcessorType) IsValid() bool {
	for _, existing := range allowedLogsCategoryProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsCategoryProcessorType value.
func (v LogsCategoryProcessorType) Ptr() *LogsCategoryProcessorType {
	return &v
}
