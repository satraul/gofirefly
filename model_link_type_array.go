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

// LinkTypeArray struct for LinkTypeArray
type LinkTypeArray struct {
	Data                 []LinkTypeRead `json:"data"`
	Meta                 Meta           `json:"meta"`
	Links                PageLink       `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _LinkTypeArray LinkTypeArray

// NewLinkTypeArray instantiates a new LinkTypeArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTypeArray(data []LinkTypeRead, meta Meta, links PageLink) *LinkTypeArray {
	this := LinkTypeArray{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewLinkTypeArrayWithDefaults instantiates a new LinkTypeArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTypeArrayWithDefaults() *LinkTypeArray {
	this := LinkTypeArray{}
	return &this
}

// GetData returns the Data field value
func (o *LinkTypeArray) GetData() []LinkTypeRead {
	if o == nil {
		var ret []LinkTypeRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LinkTypeArray) GetDataOk() (*[]LinkTypeRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LinkTypeArray) SetData(v []LinkTypeRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *LinkTypeArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *LinkTypeArray) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *LinkTypeArray) SetMeta(v Meta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *LinkTypeArray) GetLinks() PageLink {
	if o == nil {
		var ret PageLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *LinkTypeArray) GetLinksOk() (*PageLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *LinkTypeArray) SetLinks(v PageLink) {
	o.Links = v
}

func (o LinkTypeArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	if true {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTypeArray) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTypeArray := _LinkTypeArray{}

	if err = json.Unmarshal(bytes, &varLinkTypeArray); err == nil {
		*o = LinkTypeArray(varLinkTypeArray)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTypeArray struct {
	value *LinkTypeArray
	isSet bool
}

func (v NullableLinkTypeArray) Get() *LinkTypeArray {
	return v.value
}

func (v *NullableLinkTypeArray) Set(val *LinkTypeArray) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTypeArray) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTypeArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTypeArray(val *LinkTypeArray) *NullableLinkTypeArray {
	return &NullableLinkTypeArray{value: val, isSet: true}
}

func (v NullableLinkTypeArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTypeArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
