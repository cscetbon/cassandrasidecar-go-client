# CleanupOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Keyspace** | Pointer to **string** | keyspace to cleanup  | 
**Tables** | Pointer to **[]string** | tables to cleanup, when not specified, all tables in a keyspace will be cleaned up  | [optional] 
**Jobs** | Pointer to **int32** | number of jobs to use, never uses more that concurrent_compactor threads  | [optional] 

## Methods

### NewCleanupOperationRequest

`func NewCleanupOperationRequest(type_ string, keyspace string, ) *CleanupOperationRequest`

NewCleanupOperationRequest instantiates a new CleanupOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanupOperationRequestWithDefaults

`func NewCleanupOperationRequestWithDefaults() *CleanupOperationRequest`

NewCleanupOperationRequestWithDefaults instantiates a new CleanupOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CleanupOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CleanupOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CleanupOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetKeyspace

`func (o *CleanupOperationRequest) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *CleanupOperationRequest) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *CleanupOperationRequest) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.


### GetTables

`func (o *CleanupOperationRequest) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *CleanupOperationRequest) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *CleanupOperationRequest) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *CleanupOperationRequest) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetJobs

`func (o *CleanupOperationRequest) GetJobs() int32`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *CleanupOperationRequest) GetJobsOk() (*int32, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *CleanupOperationRequest) SetJobs(v int32)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *CleanupOperationRequest) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


