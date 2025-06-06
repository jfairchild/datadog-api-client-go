// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsGrokParserType Type of logs grok parser.
type LogsGrokParserType string

// List of LogsGrokParserType.
const (
	LOGSGROKPARSERTYPE_GROK_PARSER LogsGrokParserType = "grok-parser"
)

var allowedLogsGrokParserTypeEnumValues = []LogsGrokParserType{
	LOGSGROKPARSERTYPE_GROK_PARSER,
}

// GetAllowedValues returns the list of possible values.
func (v *LogsGrokParserType) GetAllowedValues() []LogsGrokParserType {
	return allowedLogsGrokParserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsGrokParserType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsGrokParserType(value)
	return nil
}

// NewLogsGrokParserTypeFromValue returns a pointer to a valid LogsGrokParserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsGrokParserTypeFromValue(v string) (*LogsGrokParserType, error) {
	ev := LogsGrokParserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsGrokParserType: valid values are %v", v, allowedLogsGrokParserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsGrokParserType) IsValid() bool {
	for _, existing := range allowedLogsGrokParserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsGrokParserType value.
func (v LogsGrokParserType) Ptr() *LogsGrokParserType {
	return &v
}
