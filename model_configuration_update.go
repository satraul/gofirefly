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

// ConfigurationUpdate struct for ConfigurationUpdate
type ConfigurationUpdate struct {
	// Can be a number, a string, boolean or object. This depends on the actual configuration value.
	Value                []string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationUpdate ConfigurationUpdate

// NewConfigurationUpdate instantiates a new ConfigurationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationUpdate(value []string) *ConfigurationUpdate {
	this := ConfigurationUpdate{}
	this.Value = value
	return &this
}

// NewConfigurationUpdateWithDefaults instantiates a new ConfigurationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationUpdateWithDefaults() *ConfigurationUpdate {
	this := ConfigurationUpdate{}
	return &this
}

// GetValue returns the Value field value
func (o *ConfigurationUpdate) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ConfigurationUpdate) GetValueOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ConfigurationUpdate) SetValue(v []string) {
	o.Value = v
}

func (o ConfigurationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConfigurationUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varConfigurationUpdate := _ConfigurationUpdate{}

	if err = json.Unmarshal(bytes, &varConfigurationUpdate); err == nil {
		*o = ConfigurationUpdate(varConfigurationUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigurationUpdate struct {
	value *ConfigurationUpdate
	isSet bool
}

func (v NullableConfigurationUpdate) Get() *ConfigurationUpdate {
	return v.value
}

func (v *NullableConfigurationUpdate) Set(val *ConfigurationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationUpdate(val *ConfigurationUpdate) *NullableConfigurationUpdate {
	return &NullableConfigurationUpdate{value: val, isSet: true}
}

func (v NullableConfigurationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
