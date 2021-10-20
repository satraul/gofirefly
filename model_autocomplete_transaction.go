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

// AutocompleteTransaction struct for AutocompleteTransaction
type AutocompleteTransaction struct {
	// The ID of a transaction journal (basically a single split).
	Id string `json:"id"`
	// The ID of the underlying transaction group.
	TransactionGroupId *string `json:"transaction_group_id,omitempty"`
	// Transaction description
	Name string `json:"name"`
	// Transaction description
	Description          string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _AutocompleteTransaction AutocompleteTransaction

// NewAutocompleteTransaction instantiates a new AutocompleteTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteTransaction(id string, name string, description string) *AutocompleteTransaction {
	this := AutocompleteTransaction{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewAutocompleteTransactionWithDefaults instantiates a new AutocompleteTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteTransactionWithDefaults() *AutocompleteTransaction {
	this := AutocompleteTransaction{}
	return &this
}

// GetId returns the Id field value
func (o *AutocompleteTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutocompleteTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutocompleteTransaction) SetId(v string) {
	o.Id = v
}

// GetTransactionGroupId returns the TransactionGroupId field value if set, zero value otherwise.
func (o *AutocompleteTransaction) GetTransactionGroupId() string {
	if o == nil || o.TransactionGroupId == nil {
		var ret string
		return ret
	}
	return *o.TransactionGroupId
}

// GetTransactionGroupIdOk returns a tuple with the TransactionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompleteTransaction) GetTransactionGroupIdOk() (*string, bool) {
	if o == nil || o.TransactionGroupId == nil {
		return nil, false
	}
	return o.TransactionGroupId, true
}

// HasTransactionGroupId returns a boolean if a field has been set.
func (o *AutocompleteTransaction) HasTransactionGroupId() bool {
	if o != nil && o.TransactionGroupId != nil {
		return true
	}

	return false
}

// SetTransactionGroupId gets a reference to the given string and assigns it to the TransactionGroupId field.
func (o *AutocompleteTransaction) SetTransactionGroupId(v string) {
	o.TransactionGroupId = &v
}

// GetName returns the Name field value
func (o *AutocompleteTransaction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutocompleteTransaction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutocompleteTransaction) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AutocompleteTransaction) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AutocompleteTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AutocompleteTransaction) SetDescription(v string) {
	o.Description = v
}

func (o AutocompleteTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.TransactionGroupId != nil {
		toSerialize["transaction_group_id"] = o.TransactionGroupId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AutocompleteTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varAutocompleteTransaction := _AutocompleteTransaction{}

	if err = json.Unmarshal(bytes, &varAutocompleteTransaction); err == nil {
		*o = AutocompleteTransaction(varAutocompleteTransaction)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "transaction_group_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutocompleteTransaction struct {
	value *AutocompleteTransaction
	isSet bool
}

func (v NullableAutocompleteTransaction) Get() *AutocompleteTransaction {
	return v.value
}

func (v *NullableAutocompleteTransaction) Set(val *AutocompleteTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteTransaction(val *AutocompleteTransaction) *NullableAutocompleteTransaction {
	return &NullableAutocompleteTransaction{value: val, isSet: true}
}

func (v NullableAutocompleteTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}