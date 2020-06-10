# RefreshOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Keyspace** | Pointer to **string** | keyspace to refresh  | 
**Table** | Pointer to **string** | table to refresh  | 

## Methods

### NewRefreshOperationRequest

`func NewRefreshOperationRequest(type_ string, keyspace string, table string, ) *RefreshOperationRequest`

NewRefreshOperationRequest instantiates a new RefreshOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshOperationRequestWithDefaults

`func NewRefreshOperationRequestWithDefaults() *RefreshOperationRequest`

NewRefreshOperationRequestWithDefaults instantiates a new RefreshOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RefreshOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RefreshOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RefreshOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetKeyspace

`func (o *RefreshOperationRequest) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *RefreshOperationRequest) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *RefreshOperationRequest) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.


### GetTable

`func (o *RefreshOperationRequest) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *RefreshOperationRequest) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *RefreshOperationRequest) SetTable(v string)`

SetTable sets Table field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


