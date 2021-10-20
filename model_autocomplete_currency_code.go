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

// AutocompleteCurrencyCode struct for AutocompleteCurrencyCode
type AutocompleteCurrencyCode struct {
	Id string `json:"id"`
	// Currency name with the code between brackets.
	Name string `json:"name"`
	// Currency code.
	Code                 string `json:"code"`
	Symbol               string `json:"symbol"`
	DecimalPlaces        int32  `json:"decimal_places"`
	AdditionalProperties map[string]interface{}
}

type _AutocompleteCurrencyCode AutocompleteCurrencyCode

// NewAutocompleteCurrencyCode instantiates a new AutocompleteCurrencyCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteCurrencyCode(id string, name string, code string, symbol string, decimalPlaces int32) *AutocompleteCurrencyCode {
	this := AutocompleteCurrencyCode{}
	this.Id = id
	this.Name = name
	this.Code = code
	this.Symbol = symbol
	this.DecimalPlaces = decimalPlaces
	return &this
}

// NewAutocompleteCurrencyCodeWithDefaults instantiates a new AutocompleteCurrencyCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteCurrencyCodeWithDefaults() *AutocompleteCurrencyCode {
	this := AutocompleteCurrencyCode{}
	return &this
}

// GetId returns the Id field value
func (o *AutocompleteCurrencyCode) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutocompleteCurrencyCode) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutocompleteCurrencyCode) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AutocompleteCurrencyCode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutocompleteCurrencyCode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutocompleteCurrencyCode) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value
func (o *AutocompleteCurrencyCode) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AutocompleteCurrencyCode) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AutocompleteCurrencyCode) SetCode(v string) {
	o.Code = v
}

// GetSymbol returns the Symbol field value
func (o *AutocompleteCurrencyCode) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *AutocompleteCurrencyCode) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *AutocompleteCurrencyCode) SetSymbol(v string) {
	o.Symbol = v
}

// GetDecimalPlaces returns the DecimalPlaces field value
func (o *AutocompleteCurrencyCode) GetDecimalPlaces() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DecimalPlaces
}

// GetDecimalPlacesOk returns a tuple with the DecimalPlaces field value
// and a boolean to check if the value has been set.
func (o *AutocompleteCurrencyCode) GetDecimalPlacesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DecimalPlaces, true
}

// SetDecimalPlaces sets field value
func (o *AutocompleteCurrencyCode) SetDecimalPlaces(v int32) {
	o.DecimalPlaces = v
}

func (o AutocompleteCurrencyCode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["decimal_places"] = o.DecimalPlaces
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AutocompleteCurrencyCode) UnmarshalJSON(bytes []byte) (err error) {
	varAutocompleteCurrencyCode := _AutocompleteCurrencyCode{}

	if err = json.Unmarshal(bytes, &varAutocompleteCurrencyCode); err == nil {
		*o = AutocompleteCurrencyCode(varAutocompleteCurrencyCode)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "code")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "decimal_places")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutocompleteCurrencyCode struct {
	value *AutocompleteCurrencyCode
	isSet bool
}

func (v NullableAutocompleteCurrencyCode) Get() *AutocompleteCurrencyCode {
	return v.value
}

func (v *NullableAutocompleteCurrencyCode) Set(val *AutocompleteCurrencyCode) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteCurrencyCode) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteCurrencyCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteCurrencyCode(val *AutocompleteCurrencyCode) *NullableAutocompleteCurrencyCode {
	return &NullableAutocompleteCurrencyCode{value: val, isSet: true}
}

func (v NullableAutocompleteCurrencyCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteCurrencyCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
