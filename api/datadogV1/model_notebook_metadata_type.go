// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookMetadataType Metadata type of the notebook.
type NotebookMetadataType string

// List of NotebookMetadataType.
const (
	NOTEBOOKMETADATATYPE_POSTMORTEM    NotebookMetadataType = "postmortem"
	NOTEBOOKMETADATATYPE_RUNBOOK       NotebookMetadataType = "runbook"
	NOTEBOOKMETADATATYPE_INVESTIGATION NotebookMetadataType = "investigation"
	NOTEBOOKMETADATATYPE_DOCUMENTATION NotebookMetadataType = "documentation"
	NOTEBOOKMETADATATYPE_REPORT        NotebookMetadataType = "report"
)

var allowedNotebookMetadataTypeEnumValues = []NotebookMetadataType{
	NOTEBOOKMETADATATYPE_POSTMORTEM,
	NOTEBOOKMETADATATYPE_RUNBOOK,
	NOTEBOOKMETADATATYPE_INVESTIGATION,
	NOTEBOOKMETADATATYPE_DOCUMENTATION,
	NOTEBOOKMETADATATYPE_REPORT,
}

// GetAllowedValues returns the list of possible values.
func (v *NotebookMetadataType) GetAllowedValues() []NotebookMetadataType {
	return allowedNotebookMetadataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotebookMetadataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotebookMetadataType(value)
	return nil
}

// NewNotebookMetadataTypeFromValue returns a pointer to a valid NotebookMetadataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotebookMetadataTypeFromValue(v string) (*NotebookMetadataType, error) {
	ev := NotebookMetadataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotebookMetadataType: valid values are %v", v, allowedNotebookMetadataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotebookMetadataType) IsValid() bool {
	for _, existing := range allowedNotebookMetadataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotebookMetadataType value.
func (v NotebookMetadataType) Ptr() *NotebookMetadataType {
	return &v
}

// NullableNotebookMetadataType handles when a null is used for NotebookMetadataType.
type NullableNotebookMetadataType struct {
	value *NotebookMetadataType
	isSet bool
}

// Get returns the associated value.
func (v NullableNotebookMetadataType) Get() *NotebookMetadataType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableNotebookMetadataType) Set(val *NotebookMetadataType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableNotebookMetadataType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableNotebookMetadataType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNotebookMetadataType initializes the struct as if Set has been called.
func NewNullableNotebookMetadataType(val *NotebookMetadataType) *NullableNotebookMetadataType {
	return &NullableNotebookMetadataType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableNotebookMetadataType) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableNotebookMetadataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
