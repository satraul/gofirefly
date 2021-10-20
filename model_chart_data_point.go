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

// ChartDataPoint struct for ChartDataPoint
type ChartDataPoint struct {
	// The key is the label of the value, so for example: '2018-01-01' => 13 or 'Groceries' => -123.
	Key                  *string `json:"key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChartDataPoint ChartDataPoint

// NewChartDataPoint instantiates a new ChartDataPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartDataPoint() *ChartDataPoint {
	this := ChartDataPoint{}
	return &this
}

// NewChartDataPointWithDefaults instantiates a new ChartDataPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartDataPointWithDefaults() *ChartDataPoint {
	this := ChartDataPoint{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ChartDataPoint) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartDataPoint) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ChartDataPoint) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ChartDataPoint) SetKey(v string) {
	o.Key = &v
}

func (o ChartDataPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChartDataPoint) UnmarshalJSON(bytes []byte) (err error) {
	varChartDataPoint := _ChartDataPoint{}

	if err = json.Unmarshal(bytes, &varChartDataPoint); err == nil {
		*o = ChartDataPoint(varChartDataPoint)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChartDataPoint struct {
	value *ChartDataPoint
	isSet bool
}

func (v NullableChartDataPoint) Get() *ChartDataPoint {
	return v.value
}

func (v *NullableChartDataPoint) Set(val *ChartDataPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableChartDataPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableChartDataPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartDataPoint(val *ChartDataPoint) *NullableChartDataPoint {
	return &NullableChartDataPoint{value: val, isSet: true}
}

func (v NullableChartDataPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartDataPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}