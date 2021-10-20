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
	"fmt"
)

// AccountTypeFilter the model 'AccountTypeFilter'
type AccountTypeFilter string

// List of AccountTypeFilter
const (
	ACCOUNT_ALL                     AccountTypeFilter = "all"
	ACCOUNT_ASSET                   AccountTypeFilter = "asset"
	ACCOUNT_CASH                    AccountTypeFilter = "cash"
	ACCOUNT_EXPENSE                 AccountTypeFilter = "expense"
	ACCOUNT_REVENUE                 AccountTypeFilter = "revenue"
	ACCOUNT_SPECIAL                 AccountTypeFilter = "special"
	ACCOUNT_HIDDEN                  AccountTypeFilter = "hidden"
	ACCOUNT_LIABILITY               AccountTypeFilter = "liability"
	ACCOUNT_LIABILITIES             AccountTypeFilter = "liabilities"
	ACCOUNT_DEFAULT_ACCOUNT         AccountTypeFilter = "Default account"
	ACCOUNT_CASH_ACCOUNT            AccountTypeFilter = "Cash account"
	ACCOUNT_ASSET_ACCOUNT           AccountTypeFilter = "Asset account"
	ACCOUNT_EXPENSE_ACCOUNT         AccountTypeFilter = "Expense account"
	ACCOUNT_REVENUE_ACCOUNT         AccountTypeFilter = "Revenue account"
	ACCOUNT_INITIAL_BALANCE_ACCOUNT AccountTypeFilter = "Initial balance account"
	ACCOUNT_BENEFICIARY_ACCOUNT     AccountTypeFilter = "Beneficiary account"
	ACCOUNT_IMPORT_ACCOUNT          AccountTypeFilter = "Import account"
	ACCOUNT_RECONCILIATION_ACCOUNT  AccountTypeFilter = "Reconciliation account"
	ACCOUNT_LOAN                    AccountTypeFilter = "Loan"
	ACCOUNT_DEBT                    AccountTypeFilter = "Debt"
	ACCOUNT_MORTGAGE                AccountTypeFilter = "Mortgage"
)

var allowedAccountTypeFilterEnumValues = []AccountTypeFilter{
	"all",
	"asset",
	"cash",
	"expense",
	"revenue",
	"special",
	"hidden",
	"liability",
	"liabilities",
	"Default account",
	"Cash account",
	"Asset account",
	"Expense account",
	"Revenue account",
	"Initial balance account",
	"Beneficiary account",
	"Import account",
	"Reconciliation account",
	"Loan",
	"Debt",
	"Mortgage",
}

func (v *AccountTypeFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountTypeFilter(value)
	for _, existing := range allowedAccountTypeFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountTypeFilter", value)
}

// NewAccountTypeFilterFromValue returns a pointer to a valid AccountTypeFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountTypeFilterFromValue(v string) (*AccountTypeFilter, error) {
	ev := AccountTypeFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountTypeFilter: valid values are %v", v, allowedAccountTypeFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountTypeFilter) IsValid() bool {
	for _, existing := range allowedAccountTypeFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountTypeFilter value
func (v AccountTypeFilter) Ptr() *AccountTypeFilter {
	return &v
}

type NullableAccountTypeFilter struct {
	value *AccountTypeFilter
	isSet bool
}

func (v NullableAccountTypeFilter) Get() *AccountTypeFilter {
	return v.value
}

func (v *NullableAccountTypeFilter) Set(val *AccountTypeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTypeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTypeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTypeFilter(val *AccountTypeFilter) *NullableAccountTypeFilter {
	return &NullableAccountTypeFilter{value: val, isSet: true}
}

func (v NullableAccountTypeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTypeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}