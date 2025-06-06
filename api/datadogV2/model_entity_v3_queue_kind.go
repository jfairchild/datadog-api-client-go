// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3QueueKind The definition of Entity V3 Queue Kind object.
type EntityV3QueueKind string

// List of EntityV3QueueKind.
const (
	ENTITYV3QUEUEKIND_QUEUE EntityV3QueueKind = "queue"
)

var allowedEntityV3QueueKindEnumValues = []EntityV3QueueKind{
	ENTITYV3QUEUEKIND_QUEUE,
}

// GetAllowedValues returns the list of possible values.
func (v *EntityV3QueueKind) GetAllowedValues() []EntityV3QueueKind {
	return allowedEntityV3QueueKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EntityV3QueueKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EntityV3QueueKind(value)
	return nil
}

// NewEntityV3QueueKindFromValue returns a pointer to a valid EntityV3QueueKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEntityV3QueueKindFromValue(v string) (*EntityV3QueueKind, error) {
	ev := EntityV3QueueKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EntityV3QueueKind: valid values are %v", v, allowedEntityV3QueueKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EntityV3QueueKind) IsValid() bool {
	for _, existing := range allowedEntityV3QueueKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityV3QueueKind value.
func (v EntityV3QueueKind) Ptr() *EntityV3QueueKind {
	return &v
}
