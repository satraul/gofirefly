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

// PreferenceRead struct for PreferenceRead
type PreferenceRead struct {
	// Immutable value
	Type                 string     `json:"type"`
	Id                   string     `json:"id"`
	Attributes           Preference `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _PreferenceRead PreferenceRead

// NewPreferenceRead instantiates a new PreferenceRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferenceRead(type_ string, id string, attributes Preference) *PreferenceRead {
	this := PreferenceRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPreferenceReadWithDefaults instantiates a new PreferenceRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferenceReadWithDefaults() *PreferenceRead {
	this := PreferenceRead{}
	return &this
}

// GetType returns the Type field value
func (o *PreferenceRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PreferenceRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PreferenceRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PreferenceRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PreferenceRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PreferenceRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PreferenceRead) GetAttributes() Preference {
	if o == nil {
		var ret Preference
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PreferenceRead) GetAttributesOk() (*Preference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PreferenceRead) SetAttributes(v Preference) {
	o.Attributes = v
}

func (o PreferenceRead) MarshalJSON() ([]byte, error) {
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

func (o *PreferenceRead) UnmarshalJSON(bytes []byte) (err error) {
	varPreferenceRead := _PreferenceRead{}

	if err = json.Unmarshal(bytes, &varPreferenceRead); err == nil {
		*o = PreferenceRead(varPreferenceRead)
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

type NullablePreferenceRead struct {
	value *PreferenceRead
	isSet bool
}

func (v NullablePreferenceRead) Get() *PreferenceRead {
	return v.value
}

func (v *NullablePreferenceRead) Set(val *PreferenceRead) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferenceRead) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferenceRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferenceRead(val *PreferenceRead) *NullablePreferenceRead {
	return &NullablePreferenceRead{value: val, isSet: true}
}

func (v NullablePreferenceRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferenceRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}