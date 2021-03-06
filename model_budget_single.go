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

// BudgetSingle struct for BudgetSingle
type BudgetSingle struct {
	Data                 BudgetRead `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _BudgetSingle BudgetSingle

// NewBudgetSingle instantiates a new BudgetSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetSingle(data BudgetRead) *BudgetSingle {
	this := BudgetSingle{}
	this.Data = data
	return &this
}

// NewBudgetSingleWithDefaults instantiates a new BudgetSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetSingleWithDefaults() *BudgetSingle {
	this := BudgetSingle{}
	return &this
}

// GetData returns the Data field value
func (o *BudgetSingle) GetData() BudgetRead {
	if o == nil {
		var ret BudgetRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BudgetSingle) GetDataOk() (*BudgetRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BudgetSingle) SetData(v BudgetRead) {
	o.Data = v
}

func (o BudgetSingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BudgetSingle) UnmarshalJSON(bytes []byte) (err error) {
	varBudgetSingle := _BudgetSingle{}

	if err = json.Unmarshal(bytes, &varBudgetSingle); err == nil {
		*o = BudgetSingle(varBudgetSingle)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBudgetSingle struct {
	value *BudgetSingle
	isSet bool
}

func (v NullableBudgetSingle) Get() *BudgetSingle {
	return v.value
}

func (v *NullableBudgetSingle) Set(val *BudgetSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetSingle(val *BudgetSingle) *NullableBudgetSingle {
	return &NullableBudgetSingle{value: val, isSet: true}
}

func (v NullableBudgetSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
