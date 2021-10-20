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

// TransactionStore struct for TransactionStore
type TransactionStore struct {
	// Break if the submitted transaction exists already.
	ErrorIfDuplicateHash *bool `json:"error_if_duplicate_hash,omitempty"`
	// Whether or not to apply rules when submitting transaction.
	ApplyRules *bool `json:"apply_rules,omitempty"`
	// Whether or not to fire the webhooks that are related to this event.
	FireWebhooks *bool `json:"fire_webhooks,omitempty"`
	// Title of the transaction if it has been split in more than one piece. Empty otherwise.
	GroupTitle           NullableString          `json:"group_title,omitempty"`
	Transactions         []TransactionSplitStore `json:"transactions"`
	AdditionalProperties map[string]interface{}
}

type _TransactionStore TransactionStore

// NewTransactionStore instantiates a new TransactionStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionStore(transactions []TransactionSplitStore) *TransactionStore {
	this := TransactionStore{}
	var fireWebhooks bool = true
	this.FireWebhooks = &fireWebhooks
	this.Transactions = transactions
	return &this
}

// NewTransactionStoreWithDefaults instantiates a new TransactionStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionStoreWithDefaults() *TransactionStore {
	this := TransactionStore{}
	var fireWebhooks bool = true
	this.FireWebhooks = &fireWebhooks
	return &this
}

// GetErrorIfDuplicateHash returns the ErrorIfDuplicateHash field value if set, zero value otherwise.
func (o *TransactionStore) GetErrorIfDuplicateHash() bool {
	if o == nil || o.ErrorIfDuplicateHash == nil {
		var ret bool
		return ret
	}
	return *o.ErrorIfDuplicateHash
}

// GetErrorIfDuplicateHashOk returns a tuple with the ErrorIfDuplicateHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStore) GetErrorIfDuplicateHashOk() (*bool, bool) {
	if o == nil || o.ErrorIfDuplicateHash == nil {
		return nil, false
	}
	return o.ErrorIfDuplicateHash, true
}

// HasErrorIfDuplicateHash returns a boolean if a field has been set.
func (o *TransactionStore) HasErrorIfDuplicateHash() bool {
	if o != nil && o.ErrorIfDuplicateHash != nil {
		return true
	}

	return false
}

// SetErrorIfDuplicateHash gets a reference to the given bool and assigns it to the ErrorIfDuplicateHash field.
func (o *TransactionStore) SetErrorIfDuplicateHash(v bool) {
	o.ErrorIfDuplicateHash = &v
}

// GetApplyRules returns the ApplyRules field value if set, zero value otherwise.
func (o *TransactionStore) GetApplyRules() bool {
	if o == nil || o.ApplyRules == nil {
		var ret bool
		return ret
	}
	return *o.ApplyRules
}

// GetApplyRulesOk returns a tuple with the ApplyRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStore) GetApplyRulesOk() (*bool, bool) {
	if o == nil || o.ApplyRules == nil {
		return nil, false
	}
	return o.ApplyRules, true
}

// HasApplyRules returns a boolean if a field has been set.
func (o *TransactionStore) HasApplyRules() bool {
	if o != nil && o.ApplyRules != nil {
		return true
	}

	return false
}

// SetApplyRules gets a reference to the given bool and assigns it to the ApplyRules field.
func (o *TransactionStore) SetApplyRules(v bool) {
	o.ApplyRules = &v
}

// GetFireWebhooks returns the FireWebhooks field value if set, zero value otherwise.
func (o *TransactionStore) GetFireWebhooks() bool {
	if o == nil || o.FireWebhooks == nil {
		var ret bool
		return ret
	}
	return *o.FireWebhooks
}

// GetFireWebhooksOk returns a tuple with the FireWebhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStore) GetFireWebhooksOk() (*bool, bool) {
	if o == nil || o.FireWebhooks == nil {
		return nil, false
	}
	return o.FireWebhooks, true
}

// HasFireWebhooks returns a boolean if a field has been set.
func (o *TransactionStore) HasFireWebhooks() bool {
	if o != nil && o.FireWebhooks != nil {
		return true
	}

	return false
}

// SetFireWebhooks gets a reference to the given bool and assigns it to the FireWebhooks field.
func (o *TransactionStore) SetFireWebhooks(v bool) {
	o.FireWebhooks = &v
}

// GetGroupTitle returns the GroupTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionStore) GetGroupTitle() string {
	if o == nil || o.GroupTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.GroupTitle.Get()
}

// GetGroupTitleOk returns a tuple with the GroupTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionStore) GetGroupTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupTitle.Get(), o.GroupTitle.IsSet()
}

// HasGroupTitle returns a boolean if a field has been set.
func (o *TransactionStore) HasGroupTitle() bool {
	if o != nil && o.GroupTitle.IsSet() {
		return true
	}

	return false
}

// SetGroupTitle gets a reference to the given NullableString and assigns it to the GroupTitle field.
func (o *TransactionStore) SetGroupTitle(v string) {
	o.GroupTitle.Set(&v)
}

// SetGroupTitleNil sets the value for GroupTitle to be an explicit nil
func (o *TransactionStore) SetGroupTitleNil() {
	o.GroupTitle.Set(nil)
}

// UnsetGroupTitle ensures that no value is present for GroupTitle, not even an explicit nil
func (o *TransactionStore) UnsetGroupTitle() {
	o.GroupTitle.Unset()
}

// GetTransactions returns the Transactions field value
func (o *TransactionStore) GetTransactions() []TransactionSplitStore {
	if o == nil {
		var ret []TransactionSplitStore
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *TransactionStore) GetTransactionsOk() (*[]TransactionSplitStore, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *TransactionStore) SetTransactions(v []TransactionSplitStore) {
	o.Transactions = v
}

func (o TransactionStore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorIfDuplicateHash != nil {
		toSerialize["error_if_duplicate_hash"] = o.ErrorIfDuplicateHash
	}
	if o.ApplyRules != nil {
		toSerialize["apply_rules"] = o.ApplyRules
	}
	if o.FireWebhooks != nil {
		toSerialize["fire_webhooks"] = o.FireWebhooks
	}
	if o.GroupTitle.IsSet() {
		toSerialize["group_title"] = o.GroupTitle.Get()
	}
	if true {
		toSerialize["transactions"] = o.Transactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionStore) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionStore := _TransactionStore{}

	if err = json.Unmarshal(bytes, &varTransactionStore); err == nil {
		*o = TransactionStore(varTransactionStore)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error_if_duplicate_hash")
		delete(additionalProperties, "apply_rules")
		delete(additionalProperties, "fire_webhooks")
		delete(additionalProperties, "group_title")
		delete(additionalProperties, "transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionStore struct {
	value *TransactionStore
	isSet bool
}

func (v NullableTransactionStore) Get() *TransactionStore {
	return v.value
}

func (v *NullableTransactionStore) Set(val *TransactionStore) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStore) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStore(val *TransactionStore) *NullableTransactionStore {
	return &NullableTransactionStore{value: val, isSet: true}
}

func (v NullableTransactionStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
