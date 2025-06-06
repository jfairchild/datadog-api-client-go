// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreatePageRequestData The main request body, including attributes and resource type.
type CreatePageRequestData struct {
	// Details about the On-Call Page you want to create.
	Attributes *CreatePageRequestDataAttributes `json:"attributes,omitempty"`
	// The type of resource used when creating an On-Call Page.
	Type CreatePageRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreatePageRequestData instantiates a new CreatePageRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreatePageRequestData(typeVar CreatePageRequestDataType) *CreatePageRequestData {
	this := CreatePageRequestData{}
	this.Type = typeVar
	return &this
}

// NewCreatePageRequestDataWithDefaults instantiates a new CreatePageRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreatePageRequestDataWithDefaults() *CreatePageRequestData {
	this := CreatePageRequestData{}
	var typeVar CreatePageRequestDataType = CREATEPAGEREQUESTDATATYPE_PAGES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreatePageRequestData) GetAttributes() CreatePageRequestDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret CreatePageRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePageRequestData) GetAttributesOk() (*CreatePageRequestDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreatePageRequestData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CreatePageRequestDataAttributes and assigns it to the Attributes field.
func (o *CreatePageRequestData) SetAttributes(v CreatePageRequestDataAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *CreatePageRequestData) GetType() CreatePageRequestDataType {
	if o == nil {
		var ret CreatePageRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreatePageRequestData) GetTypeOk() (*CreatePageRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreatePageRequestData) SetType(v CreatePageRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreatePageRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreatePageRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CreatePageRequestDataAttributes `json:"attributes,omitempty"`
		Type       *CreatePageRequestDataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
