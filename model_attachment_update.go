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

// AttachmentUpdate struct for AttachmentUpdate
type AttachmentUpdate struct {
	Filename             *string        `json:"filename,omitempty"`
	Title                *string        `json:"title,omitempty"`
	Notes                NullableString `json:"notes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AttachmentUpdate AttachmentUpdate

// NewAttachmentUpdate instantiates a new AttachmentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentUpdate() *AttachmentUpdate {
	this := AttachmentUpdate{}
	return &this
}

// NewAttachmentUpdateWithDefaults instantiates a new AttachmentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentUpdateWithDefaults() *AttachmentUpdate {
	this := AttachmentUpdate{}
	return &this
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *AttachmentUpdate) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentUpdate) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *AttachmentUpdate) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *AttachmentUpdate) SetFilename(v string) {
	o.Filename = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AttachmentUpdate) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentUpdate) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AttachmentUpdate) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AttachmentUpdate) SetTitle(v string) {
	o.Title = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentUpdate) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentUpdate) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *AttachmentUpdate) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *AttachmentUpdate) SetNotes(v string) {
	o.Notes.Set(&v)
}

// SetNotesNil sets the value for Notes to be an explicit nil
func (o *AttachmentUpdate) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *AttachmentUpdate) UnsetNotes() {
	o.Notes.Unset()
}

func (o AttachmentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AttachmentUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varAttachmentUpdate := _AttachmentUpdate{}

	if err = json.Unmarshal(bytes, &varAttachmentUpdate); err == nil {
		*o = AttachmentUpdate(varAttachmentUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "filename")
		delete(additionalProperties, "title")
		delete(additionalProperties, "notes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttachmentUpdate struct {
	value *AttachmentUpdate
	isSet bool
}

func (v NullableAttachmentUpdate) Get() *AttachmentUpdate {
	return v.value
}

func (v *NullableAttachmentUpdate) Set(val *AttachmentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentUpdate(val *AttachmentUpdate) *NullableAttachmentUpdate {
	return &NullableAttachmentUpdate{value: val, isSet: true}
}

func (v NullableAttachmentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
