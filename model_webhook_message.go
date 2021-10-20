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
	"time"
)

// WebhookMessage struct for WebhookMessage
type WebhookMessage struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// If this message is sent yet.
	Sent *bool `json:"sent,omitempty"`
	// If this message has errored out.
	Errored *bool `json:"errored,omitempty"`
	// The ID of the webhook this message belongs to.
	WebhookId *string `json:"webhook_id,omitempty"`
	// Long UUID string for identification of this webhook message.
	Uuid *string `json:"uuid,omitempty"`
	// The actual message that is sent or will be sent as JSON string.
	String               NullableString `json:"string,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebhookMessage WebhookMessage

// NewWebhookMessage instantiates a new WebhookMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookMessage() *WebhookMessage {
	this := WebhookMessage{}
	return &this
}

// NewWebhookMessageWithDefaults instantiates a new WebhookMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookMessageWithDefaults() *WebhookMessage {
	this := WebhookMessage{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WebhookMessage) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookMessage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WebhookMessage) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WebhookMessage) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WebhookMessage) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookMessage) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WebhookMessage) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WebhookMessage) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *WebhookMessage) GetSent() bool {
	if o == nil || o.Sent == nil {
		var ret bool
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookMessage) GetSentOk() (*bool, bool) {
	if o == nil || o.Sent == nil {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *WebhookMessage) HasSent() bool {
	if o != nil && o.Sent != nil {
		return true
	}

	return false
}

// SetSent gets a reference to the given bool and assigns it to the Sent field.
func (o *WebhookMessage) SetSent(v bool) {
	o.Sent = &v
}

// GetErrored returns the Errored field value if set, zero value otherwise.
func (o *WebhookMessage) GetErrored() bool {
	if o == nil || o.Errored == nil {
		var ret bool
		return ret
	}
	return *o.Errored
}

// GetErroredOk returns a tuple with the Errored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookMessage) GetErroredOk() (*bool, bool) {
	if o == nil || o.Errored == nil {
		return nil, false
	}
	return o.Errored, true
}

// HasErrored returns a boolean if a field has been set.
func (o *WebhookMessage) HasErrored() bool {
	if o != nil && o.Errored != nil {
		return true
	}

	return false
}

// SetErrored gets a reference to the given bool and assigns it to the Errored field.
func (o *WebhookMessage) SetErrored(v bool) {
	o.Errored = &v
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *WebhookMessage) GetWebhookId() string {
	if o == nil || o.WebhookId == nil {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookMessage) GetWebhookIdOk() (*string, bool) {
	if o == nil || o.WebhookId == nil {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *WebhookMessage) HasWebhookId() bool {
	if o != nil && o.WebhookId != nil {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *WebhookMessage) SetWebhookId(v string) {
	o.WebhookId = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *WebhookMessage) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookMessage) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *WebhookMessage) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *WebhookMessage) SetUuid(v string) {
	o.Uuid = &v
}

// GetString returns the String field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookMessage) GetString() string {
	if o == nil || o.String.Get() == nil {
		var ret string
		return ret
	}
	return *o.String.Get()
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookMessage) GetStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.String.Get(), o.String.IsSet()
}

// HasString returns a boolean if a field has been set.
func (o *WebhookMessage) HasString() bool {
	if o != nil && o.String.IsSet() {
		return true
	}

	return false
}

// SetString gets a reference to the given NullableString and assigns it to the String field.
func (o *WebhookMessage) SetString(v string) {
	o.String.Set(&v)
}

// SetStringNil sets the value for String to be an explicit nil
func (o *WebhookMessage) SetStringNil() {
	o.String.Set(nil)
}

// UnsetString ensures that no value is present for String, not even an explicit nil
func (o *WebhookMessage) UnsetString() {
	o.String.Unset()
}

func (o WebhookMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Sent != nil {
		toSerialize["sent"] = o.Sent
	}
	if o.Errored != nil {
		toSerialize["errored"] = o.Errored
	}
	if o.WebhookId != nil {
		toSerialize["webhook_id"] = o.WebhookId
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.String.IsSet() {
		toSerialize["string"] = o.String.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WebhookMessage) UnmarshalJSON(bytes []byte) (err error) {
	varWebhookMessage := _WebhookMessage{}

	if err = json.Unmarshal(bytes, &varWebhookMessage); err == nil {
		*o = WebhookMessage(varWebhookMessage)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "sent")
		delete(additionalProperties, "errored")
		delete(additionalProperties, "webhook_id")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "string")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookMessage struct {
	value *WebhookMessage
	isSet bool
}

func (v NullableWebhookMessage) Get() *WebhookMessage {
	return v.value
}

func (v *NullableWebhookMessage) Set(val *WebhookMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookMessage(val *WebhookMessage) *NullableWebhookMessage {
	return &NullableWebhookMessage{value: val, isSet: true}
}

func (v NullableWebhookMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
