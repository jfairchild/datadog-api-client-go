// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRulesListResponse Response object that includes a list of Agent rule
type CloudWorkloadSecurityAgentRulesListResponse struct {
	// A list of Agent rules objects
	Data []CloudWorkloadSecurityAgentRuleData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRulesListResponse instantiates a new CloudWorkloadSecurityAgentRulesListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRulesListResponse() *CloudWorkloadSecurityAgentRulesListResponse {
	this := CloudWorkloadSecurityAgentRulesListResponse{}
	return &this
}

// NewCloudWorkloadSecurityAgentRulesListResponseWithDefaults instantiates a new CloudWorkloadSecurityAgentRulesListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRulesListResponseWithDefaults() *CloudWorkloadSecurityAgentRulesListResponse {
	this := CloudWorkloadSecurityAgentRulesListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRulesListResponse) GetData() []CloudWorkloadSecurityAgentRuleData {
	if o == nil || o.Data == nil {
		var ret []CloudWorkloadSecurityAgentRuleData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRulesListResponse) GetDataOk() (*[]CloudWorkloadSecurityAgentRuleData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRulesListResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []CloudWorkloadSecurityAgentRuleData and assigns it to the Data field.
func (o *CloudWorkloadSecurityAgentRulesListResponse) SetData(v []CloudWorkloadSecurityAgentRuleData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRulesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentRulesListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []CloudWorkloadSecurityAgentRuleData `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
