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

// CategoryEarned struct for CategoryEarned
type CategoryEarned struct {
	CurrencyId     *string `json:"currency_id,omitempty"`
	CurrencyCode   *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	// Number of decimals supported by the currency
	CurrencyDecimalPlaces *int32 `json:"currency_decimal_places,omitempty"`
	// The amount earned.
	Sum                  *string `json:"sum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CategoryEarned CategoryEarned

// NewCategoryEarned instantiates a new CategoryEarned object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryEarned() *CategoryEarned {
	this := CategoryEarned{}
	return &this
}

// NewCategoryEarnedWithDefaults instantiates a new CategoryEarned object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryEarnedWithDefaults() *CategoryEarned {
	this := CategoryEarned{}
	return &this
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *CategoryEarned) GetCurrencyId() string {
	if o == nil || o.CurrencyId == nil {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryEarned) GetCurrencyIdOk() (*string, bool) {
	if o == nil || o.CurrencyId == nil {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *CategoryEarned) HasCurrencyId() bool {
	if o != nil && o.CurrencyId != nil {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *CategoryEarned) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *CategoryEarned) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryEarned) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *CategoryEarned) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *CategoryEarned) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *CategoryEarned) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryEarned) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *CategoryEarned) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *CategoryEarned) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *CategoryEarned) GetCurrencyDecimalPlaces() int32 {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryEarned) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *CategoryEarned) HasCurrencyDecimalPlaces() bool {
	if o != nil && o.CurrencyDecimalPlaces != nil {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *CategoryEarned) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

// GetSum returns the Sum field value if set, zero value otherwise.
func (o *CategoryEarned) GetSum() string {
	if o == nil || o.Sum == nil {
		var ret string
		return ret
	}
	return *o.Sum
}

// GetSumOk returns a tuple with the Sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryEarned) GetSumOk() (*string, bool) {
	if o == nil || o.Sum == nil {
		return nil, false
	}
	return o.Sum, true
}

// HasSum returns a boolean if a field has been set.
func (o *CategoryEarned) HasSum() bool {
	if o != nil && o.Sum != nil {
		return true
	}

	return false
}

// SetSum gets a reference to the given string and assigns it to the Sum field.
func (o *CategoryEarned) SetSum(v string) {
	o.Sum = &v
}

func (o CategoryEarned) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencyId != nil {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.CurrencySymbol != nil {
		toSerialize["currency_symbol"] = o.CurrencySymbol
	}
	if o.CurrencyDecimalPlaces != nil {
		toSerialize["currency_decimal_places"] = o.CurrencyDecimalPlaces
	}
	if o.Sum != nil {
		toSerialize["sum"] = o.Sum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CategoryEarned) UnmarshalJSON(bytes []byte) (err error) {
	varCategoryEarned := _CategoryEarned{}

	if err = json.Unmarshal(bytes, &varCategoryEarned); err == nil {
		*o = CategoryEarned(varCategoryEarned)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "currency_id")
		delete(additionalProperties, "currency_code")
		delete(additionalProperties, "currency_symbol")
		delete(additionalProperties, "currency_decimal_places")
		delete(additionalProperties, "sum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCategoryEarned struct {
	value *CategoryEarned
	isSet bool
}

func (v NullableCategoryEarned) Get() *CategoryEarned {
	return v.value
}

func (v *NullableCategoryEarned) Set(val *CategoryEarned) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryEarned) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryEarned) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryEarned(val *CategoryEarned) *NullableCategoryEarned {
	return &NullableCategoryEarned{value: val, isSet: true}
}

func (v NullableCategoryEarned) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryEarned) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}