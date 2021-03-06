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

// CategoryArray struct for CategoryArray
type CategoryArray struct {
	Data                 []CategoryRead `json:"data"`
	Meta                 Meta           `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _CategoryArray CategoryArray

// NewCategoryArray instantiates a new CategoryArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryArray(data []CategoryRead, meta Meta) *CategoryArray {
	this := CategoryArray{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewCategoryArrayWithDefaults instantiates a new CategoryArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryArrayWithDefaults() *CategoryArray {
	this := CategoryArray{}
	return &this
}

// GetData returns the Data field value
func (o *CategoryArray) GetData() []CategoryRead {
	if o == nil {
		var ret []CategoryRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CategoryArray) GetDataOk() (*[]CategoryRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CategoryArray) SetData(v []CategoryRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *CategoryArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *CategoryArray) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *CategoryArray) SetMeta(v Meta) {
	o.Meta = v
}

func (o CategoryArray) MarshalJSON() ([]byte, error) {
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

func (o *CategoryArray) UnmarshalJSON(bytes []byte) (err error) {
	varCategoryArray := _CategoryArray{}

	if err = json.Unmarshal(bytes, &varCategoryArray); err == nil {
		*o = CategoryArray(varCategoryArray)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCategoryArray struct {
	value *CategoryArray
	isSet bool
}

func (v NullableCategoryArray) Get() *CategoryArray {
	return v.value
}

func (v *NullableCategoryArray) Set(val *CategoryArray) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryArray) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryArray(val *CategoryArray) *NullableCategoryArray {
	return &NullableCategoryArray{value: val, isSet: true}
}

func (v NullableCategoryArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
