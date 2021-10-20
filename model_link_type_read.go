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

// LinkTypeRead struct for LinkTypeRead
type LinkTypeRead struct {
	// Immutable value
	Type                 string     `json:"type"`
	Id                   string     `json:"id"`
	Attributes           LinkType   `json:"attributes"`
	Links                ObjectLink `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _LinkTypeRead LinkTypeRead

// NewLinkTypeRead instantiates a new LinkTypeRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTypeRead(type_ string, id string, attributes LinkType, links ObjectLink) *LinkTypeRead {
	this := LinkTypeRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewLinkTypeReadWithDefaults instantiates a new LinkTypeRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTypeReadWithDefaults() *LinkTypeRead {
	this := LinkTypeRead{}
	return &this
}

// GetType returns the Type field value
func (o *LinkTypeRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LinkTypeRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LinkTypeRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *LinkTypeRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LinkTypeRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LinkTypeRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *LinkTypeRead) GetAttributes() LinkType {
	if o == nil {
		var ret LinkType
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LinkTypeRead) GetAttributesOk() (*LinkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *LinkTypeRead) SetAttributes(v LinkType) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *LinkTypeRead) GetLinks() ObjectLink {
	if o == nil {
		var ret ObjectLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *LinkTypeRead) GetLinksOk() (*ObjectLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *LinkTypeRead) SetLinks(v ObjectLink) {
	o.Links = v
}

func (o LinkTypeRead) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTypeRead) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTypeRead := _LinkTypeRead{}

	if err = json.Unmarshal(bytes, &varLinkTypeRead); err == nil {
		*o = LinkTypeRead(varLinkTypeRead)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTypeRead struct {
	value *LinkTypeRead
	isSet bool
}

func (v NullableLinkTypeRead) Get() *LinkTypeRead {
	return v.value
}

func (v *NullableLinkTypeRead) Set(val *LinkTypeRead) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTypeRead) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTypeRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTypeRead(val *LinkTypeRead) *NullableLinkTypeRead {
	return &NullableLinkTypeRead{value: val, isSet: true}
}

func (v NullableLinkTypeRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTypeRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}