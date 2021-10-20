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

// AttachmentRead struct for AttachmentRead
type AttachmentRead struct {
	// Immutable value
	Type                 string     `json:"type"`
	Id                   string     `json:"id"`
	Attributes           Attachment `json:"attributes"`
	Links                ObjectLink `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _AttachmentRead AttachmentRead

// NewAttachmentRead instantiates a new AttachmentRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentRead(type_ string, id string, attributes Attachment, links ObjectLink) *AttachmentRead {
	this := AttachmentRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewAttachmentReadWithDefaults instantiates a new AttachmentRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentReadWithDefaults() *AttachmentRead {
	this := AttachmentRead{}
	return &this
}

// GetType returns the Type field value
func (o *AttachmentRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachmentRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttachmentRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AttachmentRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachmentRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *AttachmentRead) GetAttributes() Attachment {
	if o == nil {
		var ret Attachment
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AttachmentRead) GetAttributesOk() (*Attachment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AttachmentRead) SetAttributes(v Attachment) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *AttachmentRead) GetLinks() ObjectLink {
	if o == nil {
		var ret ObjectLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AttachmentRead) GetLinksOk() (*ObjectLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AttachmentRead) SetLinks(v ObjectLink) {
	o.Links = v
}

func (o AttachmentRead) MarshalJSON() ([]byte, error) {
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

func (o *AttachmentRead) UnmarshalJSON(bytes []byte) (err error) {
	varAttachmentRead := _AttachmentRead{}

	if err = json.Unmarshal(bytes, &varAttachmentRead); err == nil {
		*o = AttachmentRead(varAttachmentRead)
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

type NullableAttachmentRead struct {
	value *AttachmentRead
	isSet bool
}

func (v NullableAttachmentRead) Get() *AttachmentRead {
	return v.value
}

func (v *NullableAttachmentRead) Set(val *AttachmentRead) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentRead) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentRead(val *AttachmentRead) *NullableAttachmentRead {
	return &NullableAttachmentRead{value: val, isSet: true}
}

func (v NullableAttachmentRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
