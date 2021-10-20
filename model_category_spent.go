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

// CategorySpent struct for CategorySpent
type CategorySpent struct {
	CurrencyId     *string `json:"currency_id,omitempty"`
	CurrencyCode   *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	// Number of decimals supported by the currency
	CurrencyDecimalPlaces *int32 `json:"currency_decimal_places,omitempty"`
	// The amount spent.
	Sum                  *string `json:"sum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CategorySpent CategorySpent

// NewCategorySpent instantiates a new CategorySpent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategorySpent() *CategorySpent {
	this := CategorySpent{}
	return &this
}

// NewCategorySpentWithDefaults instantiates a new CategorySpent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategorySpentWithDefaults() *CategorySpent {
	this := CategorySpent{}
	return &this
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *CategorySpent) GetCurrencyId() string {
	if o == nil || o.CurrencyId == nil {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySpent) GetCurrencyIdOk() (*string, bool) {
	if o == nil || o.CurrencyId == nil {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *CategorySpent) HasCurrencyId() bool {
	if o != nil && o.CurrencyId != nil {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *CategorySpent) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *CategorySpent) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySpent) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *CategorySpent) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *CategorySpent) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *CategorySpent) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySpent) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *CategorySpent) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *CategorySpent) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *CategorySpent) GetCurrencyDecimalPlaces() int32 {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySpent) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *CategorySpent) HasCurrencyDecimalPlaces() bool {
	if o != nil && o.CurrencyDecimalPlaces != nil {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *CategorySpent) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

// GetSum returns the Sum field value if set, zero value otherwise.
func (o *CategorySpent) GetSum() string {
	if o == nil || o.Sum == nil {
		var ret string
		return ret
	}
	return *o.Sum
}

// GetSumOk returns a tuple with the Sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategorySpent) GetSumOk() (*string, bool) {
	if o == nil || o.Sum == nil {
		return nil, false
	}
	return o.Sum, true
}

// HasSum returns a boolean if a field has been set.
func (o *CategorySpent) HasSum() bool {
	if o != nil && o.Sum != nil {
		return true
	}

	return false
}

// SetSum gets a reference to the given string and assigns it to the Sum field.
func (o *CategorySpent) SetSum(v string) {
	o.Sum = &v
}

func (o CategorySpent) MarshalJSON() ([]byte, error) {
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

func (o *CategorySpent) UnmarshalJSON(bytes []byte) (err error) {
	varCategorySpent := _CategorySpent{}

	if err = json.Unmarshal(bytes, &varCategorySpent); err == nil {
		*o = CategorySpent(varCategorySpent)
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

type NullableCategorySpent struct {
	value *CategorySpent
	isSet bool
}

func (v NullableCategorySpent) Get() *CategorySpent {
	return v.value
}

func (v *NullableCategorySpent) Set(val *CategorySpent) {
	v.value = val
	v.isSet = true
}

func (v NullableCategorySpent) IsSet() bool {
	return v.isSet
}

func (v *NullableCategorySpent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategorySpent(val *CategorySpent) *NullableCategorySpent {
	return &NullableCategorySpent{value: val, isSet: true}
}

func (v NullableCategorySpent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategorySpent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
