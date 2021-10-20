/*
Firefly III API v1.5.4

This is the documentation of the Firefly III API. You can find accompanying documentation on the website of Firefly III itself (see below). Please report any bugs or issues. You may use the \"Authorize\" button to try the API below. This file was last generated on 2021-09-25T14:21:28+00:00

API version: 1.5.4
Contact: james@firefly-iii.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofirefly

import (
	"encoding/json"
)

// RuleRead struct for RuleRead
type RuleRead struct {
	// Immutable value
	Type                 string     `json:"type"`
	Id                   string     `json:"id"`
	Attributes           Rule       `json:"attributes"`
	Links                ObjectLink `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _RuleRead RuleRead

// NewRuleRead instantiates a new RuleRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleRead(type_ string, id string, attributes Rule, links ObjectLink) *RuleRead {
	this := RuleRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewRuleReadWithDefaults instantiates a new RuleRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleReadWithDefaults() *RuleRead {
	this := RuleRead{}
	return &this
}

// GetType returns the Type field value
func (o *RuleRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RuleRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RuleRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *RuleRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RuleRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RuleRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *RuleRead) GetAttributes() Rule {
	if o == nil {
		var ret Rule
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RuleRead) GetAttributesOk() (*Rule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *RuleRead) SetAttributes(v Rule) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *RuleRead) GetLinks() ObjectLink {
	if o == nil {
		var ret ObjectLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RuleRead) GetLinksOk() (*ObjectLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RuleRead) SetLinks(v ObjectLink) {
	o.Links = v
}

func (o RuleRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RuleRead) UnmarshalJSON(bytes []byte) (err error) {
	varRuleRead := _RuleRead{}

	if err = json.Unmarshal(bytes, &varRuleRead); err == nil {
		*o = RuleRead(varRuleRead)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleRead struct {
	value *RuleRead
	isSet bool
}

func (v NullableRuleRead) Get() *RuleRead {
	return v.value
}

func (v *NullableRuleRead) Set(val *RuleRead) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleRead) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleRead(val *RuleRead) *NullableRuleRead {
	return &NullableRuleRead{value: val, isSet: true}
}

func (v NullableRuleRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
