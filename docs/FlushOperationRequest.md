# FlushOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Keyspace** | Pointer to **string** | keyspace to flush  | 
**Tables** | Pointer to **[]string** | tables to flush, if not provided or empty, flush all tables in a keyspace  | [optional] 

## Methods

### NewFlushOperationRequest

`func NewFlushOperationRequest(type_ string, keyspace string, ) *FlushOperationRequest`

NewFlushOperationRequest instantiates a new FlushOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlushOperationRequestWithDefaults

`func NewFlushOperationRequestWithDefaults() *FlushOperationRequest`

NewFlushOperationRequestWithDefaults instantiates a new FlushOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlushOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlushOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlushOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetKeyspace

`func (o *FlushOperationRequest) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *FlushOperationRequest) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *FlushOperationRequest) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.


### GetTables

`func (o *FlushOperationRequest) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *FlushOperationRequest) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *FlushOperationRequest) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *FlushOperationRequest) HasTables() bool`

HasTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


