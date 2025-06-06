// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSearchResultsType Incident search result type.
type IncidentSearchResultsType string

// List of IncidentSearchResultsType.
const (
	INCIDENTSEARCHRESULTSTYPE_INCIDENTS_SEARCH_RESULTS IncidentSearchResultsType = "incidents_search_results"
)

var allowedIncidentSearchResultsTypeEnumValues = []IncidentSearchResultsType{
	INCIDENTSEARCHRESULTSTYPE_INCIDENTS_SEARCH_RESULTS,
}

// GetAllowedValues returns the list of possible values.
func (v *IncidentSearchResultsType) GetAllowedValues() []IncidentSearchResultsType {
	return allowedIncidentSearchResultsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentSearchResultsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentSearchResultsType(value)
	return nil
}

// NewIncidentSearchResultsTypeFromValue returns a pointer to a valid IncidentSearchResultsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentSearchResultsTypeFromValue(v string) (*IncidentSearchResultsType, error) {
	ev := IncidentSearchResultsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentSearchResultsType: valid values are %v", v, allowedIncidentSearchResultsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentSearchResultsType) IsValid() bool {
	for _, existing := range allowedIncidentSearchResultsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentSearchResultsType value.
func (v IncidentSearchResultsType) Ptr() *IncidentSearchResultsType {
	return &v
}
