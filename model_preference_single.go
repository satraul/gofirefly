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

// PreferenceSingle struct for PreferenceSingle
type PreferenceSingle struct {
	Data                 PreferenceRead `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _PreferenceSingle PreferenceSingle

// NewPreferenceSingle instantiates a new PreferenceSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferenceSingle(data PreferenceRead) *PreferenceSingle {
	this := PreferenceSingle{}
	this.Data = data
	return &this
}

// NewPreferenceSingleWithDefaults instantiates a new PreferenceSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferenceSingleWithDefaults() *PreferenceSingle {
	this := PreferenceSingle{}
	return &this
}

// GetData returns the Data field value
func (o *PreferenceSingle) GetData() PreferenceRead {
	if o == nil {
		var ret PreferenceRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PreferenceSingle) GetDataOk() (*PreferenceRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PreferenceSingle) SetData(v PreferenceRead) {
	o.Data = v
}

func (o PreferenceSingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PreferenceSingle) UnmarshalJSON(bytes []byte) (err error) {
	varPreferenceSingle := _PreferenceSingle{}

	if err = json.Unmarshal(bytes, &varPreferenceSingle); err == nil {
		*o = PreferenceSingle(varPreferenceSingle)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreferenceSingle struct {
	value *PreferenceSingle
	isSet bool
}

func (v NullablePreferenceSingle) Get() *PreferenceSingle {
	return v.value
}

func (v *NullablePreferenceSingle) Set(val *PreferenceSingle) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferenceSingle) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferenceSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferenceSingle(val *PreferenceSingle) *NullablePreferenceSingle {
	return &NullablePreferenceSingle{value: val, isSet: true}
}

func (v NullablePreferenceSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferenceSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
