# AutocompleteTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of a transaction journal (basically a single split). | 
**TransactionGroupId** | Pointer to **string** | The ID of the underlying transaction group. | [optional] 
**Name** | **string** | Transaction description | 
**Description** | **string** | Transaction description | 

## Methods

### NewAutocompleteTransaction

`func NewAutocompleteTransaction(id string, name string, description string, ) *AutocompleteTransaction`

NewAutocompleteTransaction instantiates a new AutocompleteTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteTransactionWithDefaults

`func NewAutocompleteTransactionWithDefaults() *AutocompleteTransaction`

NewAutocompleteTransactionWithDefaults instantiates a new AutocompleteTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutocompleteTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutocompleteTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutocompleteTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetTransactionGroupId

`func (o *AutocompleteTransaction) GetTransactionGroupId() string`

GetTransactionGroupId returns the TransactionGroupId field if non-nil, zero value otherwise.

### GetTransactionGroupIdOk

`func (o *AutocompleteTransaction) GetTransactionGroupIdOk() (*string, bool)`

GetTransactionGroupIdOk returns a tuple with the TransactionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionGroupId

`func (o *AutocompleteTransaction) SetTransactionGroupId(v string)`

SetTransactionGroupId sets TransactionGroupId field to given value.

### HasTransactionGroupId

`func (o *AutocompleteTransaction) HasTransactionGroupId() bool`

HasTransactionGroupId returns a boolean if a field has been set.

### GetName

`func (o *AutocompleteTransaction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutocompleteTransaction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutocompleteTransaction) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AutocompleteTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutocompleteTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutocompleteTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


