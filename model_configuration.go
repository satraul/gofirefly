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

// Configuration struct for Configuration
type Configuration struct {
	Title ConfigValueFilter `json:"title"`
	// Content of the configuration variable. Is very dynamic and can be anything from booleans to arrays.
	Value []string `json:"value"`
	// If this config variable can be edited by the user
	Editable             bool `json:"editable"`
	AdditionalProperties map[string]interface{}
}

type _Configuration Configuration

// NewConfiguration instantiates a new Configuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguration(title ConfigValueFilter, value []string, editable bool) *Configuration {
	this := Configuration{}
	this.Title = title
	this.Value = value
	this.Editable = editable
	return &this
}

// NewConfigurationWithDefaults instantiates a new Configuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationWithDefaults() *Configuration {
	this := Configuration{}
	return &this
}

// GetTitle returns the Title field value
func (o *Configuration) GetTitle() ConfigValueFilter {
	if o == nil {
		var ret ConfigValueFilter
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Configuration) GetTitleOk() (*ConfigValueFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Configuration) SetTitle(v ConfigValueFilter) {
	o.Title = v
}

// GetValue returns the Value field value
func (o *Configuration) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Configuration) GetValueOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Configuration) SetValue(v []string) {
	o.Value = v
}

// GetEditable returns the Editable field value
func (o *Configuration) GetEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Editable
}

// GetEditableOk returns a tuple with the Editable field value
// and a boolean to check if the value has been set.
func (o *Configuration) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editable, true
}

// SetEditable sets field value
func (o *Configuration) SetEditable(v bool) {
	o.Editable = v
}

func (o Configuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["editable"] = o.Editable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Configuration) UnmarshalJSON(bytes []byte) (err error) {
	varConfiguration := _Configuration{}

	if err = json.Unmarshal(bytes, &varConfiguration); err == nil {
		*o = Configuration(varConfiguration)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "value")
		delete(additionalProperties, "editable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfiguration struct {
	value *Configuration
	isSet bool
}

func (v NullableConfiguration) Get() *Configuration {
	return v.value
}

func (v *NullableConfiguration) Set(val *Configuration) {
	v.value = val
	v.isSet = true
}

func (v NullableConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfiguration(val *Configuration) *NullableConfiguration {
	return &NullableConfiguration{value: val, isSet: true}
}

func (v NullableConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
