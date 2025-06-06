// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListTeamsSort Specifies the order of the returned teams
type ListTeamsSort string

// List of ListTeamsSort.
const (
	LISTTEAMSSORT_NAME        ListTeamsSort = "name"
	LISTTEAMSSORT__NAME       ListTeamsSort = "-name"
	LISTTEAMSSORT_USER_COUNT  ListTeamsSort = "user_count"
	LISTTEAMSSORT__USER_COUNT ListTeamsSort = "-user_count"
)

var allowedListTeamsSortEnumValues = []ListTeamsSort{
	LISTTEAMSSORT_NAME,
	LISTTEAMSSORT__NAME,
	LISTTEAMSSORT_USER_COUNT,
	LISTTEAMSSORT__USER_COUNT,
}

// GetAllowedValues returns the list of possible values.
func (v *ListTeamsSort) GetAllowedValues() []ListTeamsSort {
	return allowedListTeamsSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListTeamsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListTeamsSort(value)
	return nil
}

// NewListTeamsSortFromValue returns a pointer to a valid ListTeamsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListTeamsSortFromValue(v string) (*ListTeamsSort, error) {
	ev := ListTeamsSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListTeamsSort: valid values are %v", v, allowedListTeamsSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListTeamsSort) IsValid() bool {
	for _, existing := range allowedListTeamsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListTeamsSort value.
func (v ListTeamsSort) Ptr() *ListTeamsSort {
	return &v
}
