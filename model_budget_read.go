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

// BudgetRead struct for BudgetRead
type BudgetRead struct {
	// Immutable value
	Type                 string `json:"type"`
	Id                   string `json:"id"`
	Attributes           Budget `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _BudgetRead BudgetRead

// NewBudgetRead instantiates a new BudgetRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRead(type_ string, id string, attributes Budget) *BudgetRead {
	this := BudgetRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewBudgetReadWithDefaults instantiates a new BudgetRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetReadWithDefaults() *BudgetRead {
	this := BudgetRead{}
	return &this
}

// GetType returns the Type field value
func (o *BudgetRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BudgetRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BudgetRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BudgetRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BudgetRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BudgetRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *BudgetRead) GetAttributes() Budget {
	if o == nil {
		var ret Budget
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BudgetRead) GetAttributesOk() (*Budget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BudgetRead) SetAttributes(v Budget) {
	o.Attributes = v
}

func (o BudgetRead) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BudgetRead) UnmarshalJSON(bytes []byte) (err error) {
	varBudgetRead := _BudgetRead{}

	if err = json.Unmarshal(bytes, &varBudgetRead); err == nil {
		*o = BudgetRead(varBudgetRead)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBudgetRead struct {
	value *BudgetRead
	isSet bool
}

func (v NullableBudgetRead) Get() *BudgetRead {
	return v.value
}

func (v *NullableBudgetRead) Set(val *BudgetRead) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRead) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRead(val *BudgetRead) *NullableBudgetRead {
	return &NullableBudgetRead{value: val, isSet: true}
}

func (v NullableBudgetRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
