# TruncateOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Keyspace** | Pointer to **string** | keyspace to truncate | 
**Table** | Pointer to **string** | table to truncate | 

## Methods

### NewTruncateOperationRequest

`func NewTruncateOperationRequest(type_ string, keyspace string, table string, ) *TruncateOperationRequest`

NewTruncateOperationRequest instantiates a new TruncateOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruncateOperationRequestWithDefaults

`func NewTruncateOperationRequestWithDefaults() *TruncateOperationRequest`

NewTruncateOperationRequestWithDefaults instantiates a new TruncateOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TruncateOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TruncateOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TruncateOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetKeyspace

`func (o *TruncateOperationRequest) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *TruncateOperationRequest) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *TruncateOperationRequest) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.


### GetTable

`func (o *TruncateOperationRequest) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *TruncateOperationRequest) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *TruncateOperationRequest) SetTable(v string)`

SetTable sets Table field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


