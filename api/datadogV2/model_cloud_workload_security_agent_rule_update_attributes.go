// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleUpdateAttributes Update an existing Cloud Workload Security Agent rule
type CloudWorkloadSecurityAgentRuleUpdateAttributes struct {
	// The description of the Agent rule
	Description *string `json:"description,omitempty"`
	// Whether the Agent rule is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The SECL expression of the Agent rule
	Expression *string `json:"expression,omitempty"`
	// The ID of the policy where the Agent rule is saved
	PolicyId *string `json:"policy_id,omitempty"`
	// The list of product tags associated with the rule
	ProductTags []string `json:"product_tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleUpdateAttributes instantiates a new CloudWorkloadSecurityAgentRuleUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleUpdateAttributes() *CloudWorkloadSecurityAgentRuleUpdateAttributes {
	this := CloudWorkloadSecurityAgentRuleUpdateAttributes{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleUpdateAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleUpdateAttributesWithDefaults() *CloudWorkloadSecurityAgentRuleUpdateAttributes {
	this := CloudWorkloadSecurityAgentRuleUpdateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasExpression() bool {
	return o != nil && o.Expression != nil
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetExpression(v string) {
	o.Expression = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasPolicyId() bool {
	return o != nil && o.PolicyId != nil
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetProductTags returns the ProductTags field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetProductTags() []string {
	if o == nil || o.ProductTags == nil {
		var ret []string
		return ret
	}
	return o.ProductTags
}

// GetProductTagsOk returns a tuple with the ProductTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetProductTagsOk() (*[]string, bool) {
	if o == nil || o.ProductTags == nil {
		return nil, false
	}
	return &o.ProductTags, true
}

// HasProductTags returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasProductTags() bool {
	return o != nil && o.ProductTags != nil
}

// SetProductTags gets a reference to the given []string and assigns it to the ProductTags field.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetProductTags(v []string) {
	o.ProductTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.PolicyId != nil {
		toSerialize["policy_id"] = o.PolicyId
	}
	if o.ProductTags != nil {
		toSerialize["product_tags"] = o.ProductTags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string  `json:"description,omitempty"`
		Enabled     *bool    `json:"enabled,omitempty"`
		Expression  *string  `json:"expression,omitempty"`
		PolicyId    *string  `json:"policy_id,omitempty"`
		ProductTags []string `json:"product_tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "enabled", "expression", "policy_id", "product_tags"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Enabled = all.Enabled
	o.Expression = all.Expression
	o.PolicyId = all.PolicyId
	o.ProductTags = all.ProductTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
