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

// BudgetUpdate struct for BudgetUpdate
type BudgetUpdate struct {
	Name   *string `json:"name,omitempty"`
	Active *bool   `json:"active,omitempty"`
	Order  *int32  `json:"order,omitempty"`
	// The type of auto-budget that Firefly III must create.
	AutoBudgetType NullableString `json:"auto_budget_type,omitempty"`
	// Use either currency_id or currency_code. Defaults to the user's default currency.
	AutoBudgetCurrencyId NullableString `json:"auto_budget_currency_id,omitempty"`
	// Use either currency_id or currency_code. Defaults to the user's default currency.
	AutoBudgetCurrencyCode NullableString `json:"auto_budget_currency_code,omitempty"`
	AutoBudgetAmount       NullableString `json:"auto_budget_amount,omitempty"`
	// Period for the auto budget
	AutoBudgetPeriod     NullableString `json:"auto_budget_period,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BudgetUpdate BudgetUpdate

// NewBudgetUpdate instantiates a new BudgetUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetUpdate() *BudgetUpdate {
	this := BudgetUpdate{}
	return &this
}

// NewBudgetUpdateWithDefaults instantiates a new BudgetUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetUpdateWithDefaults() *BudgetUpdate {
	this := BudgetUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BudgetUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BudgetUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BudgetUpdate) SetName(v string) {
	o.Name = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *BudgetUpdate) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUpdate) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *BudgetUpdate) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *BudgetUpdate) SetActive(v bool) {
	o.Active = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *BudgetUpdate) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUpdate) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *BudgetUpdate) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *BudgetUpdate) SetOrder(v int32) {
	o.Order = &v
}

// GetAutoBudgetType returns the AutoBudgetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetUpdate) GetAutoBudgetType() string {
	if o == nil || o.AutoBudgetType.Get() == nil {
		var ret string
		return ret
	}
	return *o.AutoBudgetType.Get()
}

// GetAutoBudgetTypeOk returns a tuple with the AutoBudgetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BudgetUpdate) GetAutoBudgetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoBudgetType.Get(), o.AutoBudgetType.IsSet()
}

// HasAutoBudgetType returns a boolean if a field has been set.
func (o *BudgetUpdate) HasAutoBudgetType() bool {
	if o != nil && o.AutoBudgetType.IsSet() {
		return true
	}

	return false
}

// SetAutoBudgetType gets a reference to the given NullableString and assigns it to the AutoBudgetType field.
func (o *BudgetUpdate) SetAutoBudgetType(v string) {
	o.AutoBudgetType.Set(&v)
}

// SetAutoBudgetTypeNil sets the value for AutoBudgetType to be an explicit nil
func (o *BudgetUpdate) SetAutoBudgetTypeNil() {
	o.AutoBudgetType.Set(nil)
}

// UnsetAutoBudgetType ensures that no value is present for AutoBudgetType, not even an explicit nil
func (o *BudgetUpdate) UnsetAutoBudgetType() {
	o.AutoBudgetType.Unset()
}

// GetAutoBudgetCurrencyId returns the AutoBudgetCurrencyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetUpdate) GetAutoBudgetCurrencyId() string {
	if o == nil || o.AutoBudgetCurrencyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AutoBudgetCurrencyId.Get()
}

// GetAutoBudgetCurrencyIdOk returns a tuple with the AutoBudgetCurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BudgetUpdate) GetAutoBudgetCurrencyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoBudgetCurrencyId.Get(), o.AutoBudgetCurrencyId.IsSet()
}

// HasAutoBudgetCurrencyId returns a boolean if a field has been set.
func (o *BudgetUpdate) HasAutoBudgetCurrencyId() bool {
	if o != nil && o.AutoBudgetCurrencyId.IsSet() {
		return true
	}

	return false
}

// SetAutoBudgetCurrencyId gets a reference to the given NullableString and assigns it to the AutoBudgetCurrencyId field.
func (o *BudgetUpdate) SetAutoBudgetCurrencyId(v string) {
	o.AutoBudgetCurrencyId.Set(&v)
}

// SetAutoBudgetCurrencyIdNil sets the value for AutoBudgetCurrencyId to be an explicit nil
func (o *BudgetUpdate) SetAutoBudgetCurrencyIdNil() {
	o.AutoBudgetCurrencyId.Set(nil)
}

// UnsetAutoBudgetCurrencyId ensures that no value is present for AutoBudgetCurrencyId, not even an explicit nil
func (o *BudgetUpdate) UnsetAutoBudgetCurrencyId() {
	o.AutoBudgetCurrencyId.Unset()
}

// GetAutoBudgetCurrencyCode returns the AutoBudgetCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetUpdate) GetAutoBudgetCurrencyCode() string {
	if o == nil || o.AutoBudgetCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.AutoBudgetCurrencyCode.Get()
}

// GetAutoBudgetCurrencyCodeOk returns a tuple with the AutoBudgetCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BudgetUpdate) GetAutoBudgetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoBudgetCurrencyCode.Get(), o.AutoBudgetCurrencyCode.IsSet()
}

// HasAutoBudgetCurrencyCode returns a boolean if a field has been set.
func (o *BudgetUpdate) HasAutoBudgetCurrencyCode() bool {
	if o != nil && o.AutoBudgetCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetAutoBudgetCurrencyCode gets a reference to the given NullableString and assigns it to the AutoBudgetCurrencyCode field.
func (o *BudgetUpdate) SetAutoBudgetCurrencyCode(v string) {
	o.AutoBudgetCurrencyCode.Set(&v)
}

// SetAutoBudgetCurrencyCodeNil sets the value for AutoBudgetCurrencyCode to be an explicit nil
func (o *BudgetUpdate) SetAutoBudgetCurrencyCodeNil() {
	o.AutoBudgetCurrencyCode.Set(nil)
}

// UnsetAutoBudgetCurrencyCode ensures that no value is present for AutoBudgetCurrencyCode, not even an explicit nil
func (o *BudgetUpdate) UnsetAutoBudgetCurrencyCode() {
	o.AutoBudgetCurrencyCode.Unset()
}

// GetAutoBudgetAmount returns the AutoBudgetAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetUpdate) GetAutoBudgetAmount() string {
	if o == nil || o.AutoBudgetAmount.Get() == nil {
		var ret string
		return ret
	}
	return *o.AutoBudgetAmount.Get()
}

// GetAutoBudgetAmountOk returns a tuple with the AutoBudgetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BudgetUpdate) GetAutoBudgetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoBudgetAmount.Get(), o.AutoBudgetAmount.IsSet()
}

// HasAutoBudgetAmount returns a boolean if a field has been set.
func (o *BudgetUpdate) HasAutoBudgetAmount() bool {
	if o != nil && o.AutoBudgetAmount.IsSet() {
		return true
	}

	return false
}

// SetAutoBudgetAmount gets a reference to the given NullableString and assigns it to the AutoBudgetAmount field.
func (o *BudgetUpdate) SetAutoBudgetAmount(v string) {
	o.AutoBudgetAmount.Set(&v)
}

// SetAutoBudgetAmountNil sets the value for AutoBudgetAmount to be an explicit nil
func (o *BudgetUpdate) SetAutoBudgetAmountNil() {
	o.AutoBudgetAmount.Set(nil)
}

// UnsetAutoBudgetAmount ensures that no value is present for AutoBudgetAmount, not even an explicit nil
func (o *BudgetUpdate) UnsetAutoBudgetAmount() {
	o.AutoBudgetAmount.Unset()
}

// GetAutoBudgetPeriod returns the AutoBudgetPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetUpdate) GetAutoBudgetPeriod() string {
	if o == nil || o.AutoBudgetPeriod.Get() == nil {
		var ret string
		return ret
	}
	return *o.AutoBudgetPeriod.Get()
}

// GetAutoBudgetPeriodOk returns a tuple with the AutoBudgetPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BudgetUpdate) GetAutoBudgetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoBudgetPeriod.Get(), o.AutoBudgetPeriod.IsSet()
}

// HasAutoBudgetPeriod returns a boolean if a field has been set.
func (o *BudgetUpdate) HasAutoBudgetPeriod() bool {
	if o != nil && o.AutoBudgetPeriod.IsSet() {
		return true
	}

	return false
}

// SetAutoBudgetPeriod gets a reference to the given NullableString and assigns it to the AutoBudgetPeriod field.
func (o *BudgetUpdate) SetAutoBudgetPeriod(v string) {
	o.AutoBudgetPeriod.Set(&v)
}

// SetAutoBudgetPeriodNil sets the value for AutoBudgetPeriod to be an explicit nil
func (o *BudgetUpdate) SetAutoBudgetPeriodNil() {
	o.AutoBudgetPeriod.Set(nil)
}

// UnsetAutoBudgetPeriod ensures that no value is present for AutoBudgetPeriod, not even an explicit nil
func (o *BudgetUpdate) UnsetAutoBudgetPeriod() {
	o.AutoBudgetPeriod.Unset()
}

func (o BudgetUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.AutoBudgetType.IsSet() {
		toSerialize["auto_budget_type"] = o.AutoBudgetType.Get()
	}
	if o.AutoBudgetCurrencyId.IsSet() {
		toSerialize["auto_budget_currency_id"] = o.AutoBudgetCurrencyId.Get()
	}
	if o.AutoBudgetCurrencyCode.IsSet() {
		toSerialize["auto_budget_currency_code"] = o.AutoBudgetCurrencyCode.Get()
	}
	if o.AutoBudgetAmount.IsSet() {
		toSerialize["auto_budget_amount"] = o.AutoBudgetAmount.Get()
	}
	if o.AutoBudgetPeriod.IsSet() {
		toSerialize["auto_budget_period"] = o.AutoBudgetPeriod.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BudgetUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varBudgetUpdate := _BudgetUpdate{}

	if err = json.Unmarshal(bytes, &varBudgetUpdate); err == nil {
		*o = BudgetUpdate(varBudgetUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "active")
		delete(additionalProperties, "order")
		delete(additionalProperties, "auto_budget_type")
		delete(additionalProperties, "auto_budget_currency_id")
		delete(additionalProperties, "auto_budget_currency_code")
		delete(additionalProperties, "auto_budget_amount")
		delete(additionalProperties, "auto_budget_period")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBudgetUpdate struct {
	value *BudgetUpdate
	isSet bool
}

func (v NullableBudgetUpdate) Get() *BudgetUpdate {
	return v.value
}

func (v *NullableBudgetUpdate) Set(val *BudgetUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetUpdate(val *BudgetUpdate) *NullableBudgetUpdate {
	return &NullableBudgetUpdate{value: val, isSet: true}
}

func (v NullableBudgetUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
