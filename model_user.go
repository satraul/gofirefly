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

// User struct for User
type User struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The new users email address.
	Email string `json:"email"`
	// Boolean to indicate if the user is blocked.
	Blocked *bool `json:"blocked,omitempty"`
	// If you say the user must be blocked, this will be the reason code.
	BlockedCode *string `json:"blocked_code,omitempty"`
	// Role for the new user. Can be empty or omitted.
	Role                 *string `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(email string) *User {
	this := User{}
	this.Email = email
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *User) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *User) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *User) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *User) GetBlocked() bool {
	if o == nil || o.Blocked == nil {
		var ret bool
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBlockedOk() (*bool, bool) {
	if o == nil || o.Blocked == nil {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *User) HasBlocked() bool {
	if o != nil && o.Blocked != nil {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given bool and assigns it to the Blocked field.
func (o *User) SetBlocked(v bool) {
	o.Blocked = &v
}

// GetBlockedCode returns the BlockedCode field value if set, zero value otherwise.
func (o *User) GetBlockedCode() string {
	if o == nil || o.BlockedCode == nil {
		var ret string
		return ret
	}
	return *o.BlockedCode
}

// GetBlockedCodeOk returns a tuple with the BlockedCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBlockedCodeOk() (*string, bool) {
	if o == nil || o.BlockedCode == nil {
		return nil, false
	}
	return o.BlockedCode, true
}

// HasBlockedCode returns a boolean if a field has been set.
func (o *User) HasBlockedCode() bool {
	if o != nil && o.BlockedCode != nil {
		return true
	}

	return false
}

// SetBlockedCode gets a reference to the given string and assigns it to the BlockedCode field.
func (o *User) SetBlockedCode(v string) {
	o.BlockedCode = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *User) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *User) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *User) SetRole(v string) {
	o.Role = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Blocked != nil {
		toSerialize["blocked"] = o.Blocked
	}
	if o.BlockedCode != nil {
		toSerialize["blocked_code"] = o.BlockedCode
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *User) UnmarshalJSON(bytes []byte) (err error) {
	varUser := _User{}

	if err = json.Unmarshal(bytes, &varUser); err == nil {
		*o = User(varUser)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "email")
		delete(additionalProperties, "blocked")
		delete(additionalProperties, "blocked_code")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}