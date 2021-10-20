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

// BudgetLimitArray struct for BudgetLimitArray
type BudgetLimitArray struct {
	Data                 []BudgetLimitRead `json:"data"`
	Meta                 Meta              `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _BudgetLimitArray BudgetLimitArray

// NewBudgetLimitArray instantiates a new BudgetLimitArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetLimitArray(data []BudgetLimitRead, meta Meta) *BudgetLimitArray {
	this := BudgetLimitArray{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewBudgetLimitArrayWithDefaults instantiates a new BudgetLimitArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetLimitArrayWithDefaults() *BudgetLimitArray {
	this := BudgetLimitArray{}
	return &this
}

// GetData returns the Data field value
func (o *BudgetLimitArray) GetData() []BudgetLimitRead {
	if o == nil {
		var ret []BudgetLimitRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BudgetLimitArray) GetDataOk() (*[]BudgetLimitRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BudgetLimitArray) SetData(v []BudgetLimitRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *BudgetLimitArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *BudgetLimitArray) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *BudgetLimitArray) SetMeta(v Meta) {
	o.Meta = v
}

func (o BudgetLimitArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BudgetLimitArray) UnmarshalJSON(bytes []byte) (err error) {
	varBudgetLimitArray := _BudgetLimitArray{}

	if err = json.Unmarshal(bytes, &varBudgetLimitArray); err == nil {
		*o = BudgetLimitArray(varBudgetLimitArray)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBudgetLimitArray struct {
	value *BudgetLimitArray
	isSet bool
}

func (v NullableBudgetLimitArray) Get() *BudgetLimitArray {
	return v.value
}

func (v *NullableBudgetLimitArray) Set(val *BudgetLimitArray) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetLimitArray) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetLimitArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetLimitArray(val *BudgetLimitArray) *NullableBudgetLimitArray {
	return &NullableBudgetLimitArray{value: val, isSet: true}
}

func (v NullableBudgetLimitArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetLimitArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
