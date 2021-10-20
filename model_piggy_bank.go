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

// PiggyBank struct for PiggyBank
type PiggyBank struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The ID of the asset account this piggy bank is connected to.
	AccountId string `json:"account_id"`
	// The name of the asset account this piggy bank is connected to.
	AccountName    *string `json:"account_name,omitempty"`
	Name           string  `json:"name"`
	CurrencyId     *string `json:"currency_id,omitempty"`
	CurrencyCode   *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	// Number of decimals supported by the currency
	CurrencyDecimalPlaces *int32   `json:"currency_decimal_places,omitempty"`
	TargetAmount          string   `json:"target_amount"`
	Percentage            *float32 `json:"percentage,omitempty"`
	CurrentAmount         *string  `json:"current_amount,omitempty"`
	LeftToSave            *string  `json:"left_to_save,omitempty"`
	SavePerMonth          *string  `json:"save_per_month,omitempty"`
	// The date you started with this piggy bank.
	StartDate *time.Time `json:"start_date,omitempty"`
	// The date you intend to finish saving money.
	TargetDate NullableTime   `json:"target_date,omitempty"`
	Order      *int32         `json:"order,omitempty"`
	Active     *bool          `json:"active,omitempty"`
	Notes      NullableString `json:"notes,omitempty"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupId NullableString `json:"object_group_id,omitempty"`
	// The order of the group. At least 1, for the highest sorting.
	ObjectGroupOrder NullableInt32 `json:"object_group_order,omitempty"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle     NullableString `json:"object_group_title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PiggyBank PiggyBank

// NewPiggyBank instantiates a new PiggyBank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPiggyBank(accountId string, name string, targetAmount string) *PiggyBank {
	this := PiggyBank{}
	this.AccountId = accountId
	this.Name = name
	this.TargetAmount = targetAmount
	return &this
}

// NewPiggyBankWithDefaults instantiates a new PiggyBank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPiggyBankWithDefaults() *PiggyBank {
	this := PiggyBank{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PiggyBank) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PiggyBank) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PiggyBank) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PiggyBank) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PiggyBank) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PiggyBank) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAccountId returns the AccountId field value
func (o *PiggyBank) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *PiggyBank) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *PiggyBank) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *PiggyBank) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *PiggyBank) SetAccountName(v string) {
	o.AccountName = &v
}

// GetName returns the Name field value
func (o *PiggyBank) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PiggyBank) SetName(v string) {
	o.Name = v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *PiggyBank) GetCurrencyId() string {
	if o == nil || o.CurrencyId == nil {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetCurrencyIdOk() (*string, bool) {
	if o == nil || o.CurrencyId == nil {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *PiggyBank) HasCurrencyId() bool {
	if o != nil && o.CurrencyId != nil {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *PiggyBank) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PiggyBank) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PiggyBank) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PiggyBank) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *PiggyBank) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *PiggyBank) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *PiggyBank) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *PiggyBank) GetCurrencyDecimalPlaces() int32 {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *PiggyBank) HasCurrencyDecimalPlaces() bool {
	if o != nil && o.CurrencyDecimalPlaces != nil {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *PiggyBank) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

// GetTargetAmount returns the TargetAmount field value
func (o *PiggyBank) GetTargetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAmount
}

// GetTargetAmountOk returns a tuple with the TargetAmount field value
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetTargetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAmount, true
}

// SetTargetAmount sets field value
func (o *PiggyBank) SetTargetAmount(v string) {
	o.TargetAmount = v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *PiggyBank) GetPercentage() float32 {
	if o == nil || o.Percentage == nil {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetPercentageOk() (*float32, bool) {
	if o == nil || o.Percentage == nil {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *PiggyBank) HasPercentage() bool {
	if o != nil && o.Percentage != nil {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *PiggyBank) SetPercentage(v float32) {
	o.Percentage = &v
}

// GetCurrentAmount returns the CurrentAmount field value if set, zero value otherwise.
func (o *PiggyBank) GetCurrentAmount() string {
	if o == nil || o.CurrentAmount == nil {
		var ret string
		return ret
	}
	return *o.CurrentAmount
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetCurrentAmountOk() (*string, bool) {
	if o == nil || o.CurrentAmount == nil {
		return nil, false
	}
	return o.CurrentAmount, true
}

// HasCurrentAmount returns a boolean if a field has been set.
func (o *PiggyBank) HasCurrentAmount() bool {
	if o != nil && o.CurrentAmount != nil {
		return true
	}

	return false
}

// SetCurrentAmount gets a reference to the given string and assigns it to the CurrentAmount field.
func (o *PiggyBank) SetCurrentAmount(v string) {
	o.CurrentAmount = &v
}

// GetLeftToSave returns the LeftToSave field value if set, zero value otherwise.
func (o *PiggyBank) GetLeftToSave() string {
	if o == nil || o.LeftToSave == nil {
		var ret string
		return ret
	}
	return *o.LeftToSave
}

// GetLeftToSaveOk returns a tuple with the LeftToSave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetLeftToSaveOk() (*string, bool) {
	if o == nil || o.LeftToSave == nil {
		return nil, false
	}
	return o.LeftToSave, true
}

// HasLeftToSave returns a boolean if a field has been set.
func (o *PiggyBank) HasLeftToSave() bool {
	if o != nil && o.LeftToSave != nil {
		return true
	}

	return false
}

// SetLeftToSave gets a reference to the given string and assigns it to the LeftToSave field.
func (o *PiggyBank) SetLeftToSave(v string) {
	o.LeftToSave = &v
}

// GetSavePerMonth returns the SavePerMonth field value if set, zero value otherwise.
func (o *PiggyBank) GetSavePerMonth() string {
	if o == nil || o.SavePerMonth == nil {
		var ret string
		return ret
	}
	return *o.SavePerMonth
}

// GetSavePerMonthOk returns a tuple with the SavePerMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetSavePerMonthOk() (*string, bool) {
	if o == nil || o.SavePerMonth == nil {
		return nil, false
	}
	return o.SavePerMonth, true
}

// HasSavePerMonth returns a boolean if a field has been set.
func (o *PiggyBank) HasSavePerMonth() bool {
	if o != nil && o.SavePerMonth != nil {
		return true
	}

	return false
}

// SetSavePerMonth gets a reference to the given string and assigns it to the SavePerMonth field.
func (o *PiggyBank) SetSavePerMonth(v string) {
	o.SavePerMonth = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *PiggyBank) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *PiggyBank) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *PiggyBank) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetTargetDate returns the TargetDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PiggyBank) GetTargetDate() time.Time {
	if o == nil || o.TargetDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.TargetDate.Get()
}

// GetTargetDateOk returns a tuple with the TargetDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PiggyBank) GetTargetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetDate.Get(), o.TargetDate.IsSet()
}

// HasTargetDate returns a boolean if a field has been set.
func (o *PiggyBank) HasTargetDate() bool {
	if o != nil && o.TargetDate.IsSet() {
		return true
	}

	return false
}

// SetTargetDate gets a reference to the given NullableTime and assigns it to the TargetDate field.
func (o *PiggyBank) SetTargetDate(v time.Time) {
	o.TargetDate.Set(&v)
}

// SetTargetDateNil sets the value for TargetDate to be an explicit nil
func (o *PiggyBank) SetTargetDateNil() {
	o.TargetDate.Set(nil)
}

// UnsetTargetDate ensures that no value is present for TargetDate, not even an explicit nil
func (o *PiggyBank) UnsetTargetDate() {
	o.TargetDate.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PiggyBank) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PiggyBank) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *PiggyBank) SetOrder(v int32) {
	o.Order = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PiggyBank) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PiggyBank) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PiggyBank) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PiggyBank) SetActive(v bool) {
	o.Active = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PiggyBank) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PiggyBank) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *PiggyBank) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *PiggyBank) SetNotes(v string) {
	o.Notes.Set(&v)
}

// SetNotesNil sets the value for Notes to be an explicit nil
func (o *PiggyBank) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *PiggyBank) UnsetNotes() {
	o.Notes.Unset()
}

// GetObjectGroupId returns the ObjectGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PiggyBank) GetObjectGroupId() string {
	if o == nil || o.ObjectGroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectGroupId.Get()
}

// GetObjectGroupIdOk returns a tuple with the ObjectGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PiggyBank) GetObjectGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectGroupId.Get(), o.ObjectGroupId.IsSet()
}

// HasObjectGroupId returns a boolean if a field has been set.
func (o *PiggyBank) HasObjectGroupId() bool {
	if o != nil && o.ObjectGroupId.IsSet() {
		return true
	}

	return false
}

// SetObjectGroupId gets a reference to the given NullableString and assigns it to the ObjectGroupId field.
func (o *PiggyBank) SetObjectGroupId(v string) {
	o.ObjectGroupId.Set(&v)
}

// SetObjectGroupIdNil sets the value for ObjectGroupId to be an explicit nil
func (o *PiggyBank) SetObjectGroupIdNil() {
	o.ObjectGroupId.Set(nil)
}

// UnsetObjectGroupId ensures that no value is present for ObjectGroupId, not even an explicit nil
func (o *PiggyBank) UnsetObjectGroupId() {
	o.ObjectGroupId.Unset()
}

// GetObjectGroupOrder returns the ObjectGroupOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PiggyBank) GetObjectGroupOrder() int32 {
	if o == nil || o.ObjectGroupOrder.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ObjectGroupOrder.Get()
}

// GetObjectGroupOrderOk returns a tuple with the ObjectGroupOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PiggyBank) GetObjectGroupOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectGroupOrder.Get(), o.ObjectGroupOrder.IsSet()
}

// HasObjectGroupOrder returns a boolean if a field has been set.
func (o *PiggyBank) HasObjectGroupOrder() bool {
	if o != nil && o.ObjectGroupOrder.IsSet() {
		return true
	}

	return false
}

// SetObjectGroupOrder gets a reference to the given NullableInt32 and assigns it to the ObjectGroupOrder field.
func (o *PiggyBank) SetObjectGroupOrder(v int32) {
	o.ObjectGroupOrder.Set(&v)
}

// SetObjectGroupOrderNil sets the value for ObjectGroupOrder to be an explicit nil
func (o *PiggyBank) SetObjectGroupOrderNil() {
	o.ObjectGroupOrder.Set(nil)
}

// UnsetObjectGroupOrder ensures that no value is present for ObjectGroupOrder, not even an explicit nil
func (o *PiggyBank) UnsetObjectGroupOrder() {
	o.ObjectGroupOrder.Unset()
}

// GetObjectGroupTitle returns the ObjectGroupTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PiggyBank) GetObjectGroupTitle() string {
	if o == nil || o.ObjectGroupTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectGroupTitle.Get()
}

// GetObjectGroupTitleOk returns a tuple with the ObjectGroupTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PiggyBank) GetObjectGroupTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectGroupTitle.Get(), o.ObjectGroupTitle.IsSet()
}

// HasObjectGroupTitle returns a boolean if a field has been set.
func (o *PiggyBank) HasObjectGroupTitle() bool {
	if o != nil && o.ObjectGroupTitle.IsSet() {
		return true
	}

	return false
}

// SetObjectGroupTitle gets a reference to the given NullableString and assigns it to the ObjectGroupTitle field.
func (o *PiggyBank) SetObjectGroupTitle(v string) {
	o.ObjectGroupTitle.Set(&v)
}

// SetObjectGroupTitleNil sets the value for ObjectGroupTitle to be an explicit nil
func (o *PiggyBank) SetObjectGroupTitleNil() {
	o.ObjectGroupTitle.Set(nil)
}

// UnsetObjectGroupTitle ensures that no value is present for ObjectGroupTitle, not even an explicit nil
func (o *PiggyBank) UnsetObjectGroupTitle() {
	o.ObjectGroupTitle.Unset()
}

func (o PiggyBank) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if true {
		toSerialize["name"] = o.Name
	}
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
	if true {
		toSerialize["target_amount"] = o.TargetAmount
	}
	if o.Percentage != nil {
		toSerialize["percentage"] = o.Percentage
	}
	if o.CurrentAmount != nil {
		toSerialize["current_amount"] = o.CurrentAmount
	}
	if o.LeftToSave != nil {
		toSerialize["left_to_save"] = o.LeftToSave
	}
	if o.SavePerMonth != nil {
		toSerialize["save_per_month"] = o.SavePerMonth
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.TargetDate.IsSet() {
		toSerialize["target_date"] = o.TargetDate.Get()
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.ObjectGroupId.IsSet() {
		toSerialize["object_group_id"] = o.ObjectGroupId.Get()
	}
	if o.ObjectGroupOrder.IsSet() {
		toSerialize["object_group_order"] = o.ObjectGroupOrder.Get()
	}
	if o.ObjectGroupTitle.IsSet() {
		toSerialize["object_group_title"] = o.ObjectGroupTitle.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PiggyBank) UnmarshalJSON(bytes []byte) (err error) {
	varPiggyBank := _PiggyBank{}

	if err = json.Unmarshal(bytes, &varPiggyBank); err == nil {
		*o = PiggyBank(varPiggyBank)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "account_name")
		delete(additionalProperties, "name")
		delete(additionalProperties, "currency_id")
		delete(additionalProperties, "currency_code")
		delete(additionalProperties, "currency_symbol")
		delete(additionalProperties, "currency_decimal_places")
		delete(additionalProperties, "target_amount")
		delete(additionalProperties, "percentage")
		delete(additionalProperties, "current_amount")
		delete(additionalProperties, "left_to_save")
		delete(additionalProperties, "save_per_month")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "target_date")
		delete(additionalProperties, "order")
		delete(additionalProperties, "active")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "object_group_id")
		delete(additionalProperties, "object_group_order")
		delete(additionalProperties, "object_group_title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePiggyBank struct {
	value *PiggyBank
	isSet bool
}

func (v NullablePiggyBank) Get() *PiggyBank {
	return v.value
}

func (v *NullablePiggyBank) Set(val *PiggyBank) {
	v.value = val
	v.isSet = true
}

func (v NullablePiggyBank) IsSet() bool {
	return v.isSet
}

func (v *NullablePiggyBank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePiggyBank(val *PiggyBank) *NullablePiggyBank {
	return &NullablePiggyBank{value: val, isSet: true}
}

func (v NullablePiggyBank) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePiggyBank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
