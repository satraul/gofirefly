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

// ObjectGroupUpdate struct for ObjectGroupUpdate
type ObjectGroupUpdate struct {
	Title *string `json:"title,omitempty"`
	// Order of the object group
	Order                *int32 `json:"order,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ObjectGroupUpdate ObjectGroupUpdate

// NewObjectGroupUpdate instantiates a new ObjectGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectGroupUpdate() *ObjectGroupUpdate {
	this := ObjectGroupUpdate{}
	return &this
}

// NewObjectGroupUpdateWithDefaults instantiates a new ObjectGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectGroupUpdateWithDefaults() *ObjectGroupUpdate {
	this := ObjectGroupUpdate{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ObjectGroupUpdate) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectGroupUpdate) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ObjectGroupUpdate) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ObjectGroupUpdate) SetTitle(v string) {
	o.Title = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ObjectGroupUpdate) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectGroupUpdate) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ObjectGroupUpdate) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ObjectGroupUpdate) SetOrder(v int32) {
	o.Order = &v
}

func (o ObjectGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ObjectGroupUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varObjectGroupUpdate := _ObjectGroupUpdate{}

	if err = json.Unmarshal(bytes, &varObjectGroupUpdate); err == nil {
		*o = ObjectGroupUpdate(varObjectGroupUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "order")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObjectGroupUpdate struct {
	value *ObjectGroupUpdate
	isSet bool
}

func (v NullableObjectGroupUpdate) Get() *ObjectGroupUpdate {
	return v.value
}

func (v *NullableObjectGroupUpdate) Set(val *ObjectGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectGroupUpdate(val *ObjectGroupUpdate) *NullableObjectGroupUpdate {
	return &NullableObjectGroupUpdate{value: val, isSet: true}
}

func (v NullableObjectGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
