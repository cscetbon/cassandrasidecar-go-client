# UpgradeSSTablesOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Keyspace** | Pointer to **string** | keyspace to upgrade SSTables of  | 
**Tables** | Pointer to **[]string** | an array of tables to upgrade SSTables of, empty or not provided array will default to upgrading of SSTables of all tables in respective keyspace  | [optional] 
**Jobs** | Pointer to **int32** | the number of threads to use - 0 means use all available, it never uses more than concurrent_compactor threads  | [optional] 
**IncludeAllSStables** | Pointer to **bool** | include all sstables, even those already on the current version, defaults to false | [optional] 

## Methods

### NewUpgradeSSTablesOperationRequest

`func NewUpgradeSSTablesOperationRequest(type_ string, keyspace string, ) *UpgradeSSTablesOperationRequest`

NewUpgradeSSTablesOperationRequest instantiates a new UpgradeSSTablesOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeSSTablesOperationRequestWithDefaults

`func NewUpgradeSSTablesOperationRequestWithDefaults() *UpgradeSSTablesOperationRequest`

NewUpgradeSSTablesOperationRequestWithDefaults instantiates a new UpgradeSSTablesOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpgradeSSTablesOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpgradeSSTablesOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpgradeSSTablesOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetKeyspace

`func (o *UpgradeSSTablesOperationRequest) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *UpgradeSSTablesOperationRequest) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *UpgradeSSTablesOperationRequest) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.


### GetTables

`func (o *UpgradeSSTablesOperationRequest) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *UpgradeSSTablesOperationRequest) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *UpgradeSSTablesOperationRequest) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *UpgradeSSTablesOperationRequest) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetJobs

`func (o *UpgradeSSTablesOperationRequest) GetJobs() int32`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *UpgradeSSTablesOperationRequest) GetJobsOk() (*int32, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *UpgradeSSTablesOperationRequest) SetJobs(v int32)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *UpgradeSSTablesOperationRequest) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetIncludeAllSStables

`func (o *UpgradeSSTablesOperationRequest) GetIncludeAllSStables() bool`

GetIncludeAllSStables returns the IncludeAllSStables field if non-nil, zero value otherwise.

### GetIncludeAllSStablesOk

`func (o *UpgradeSSTablesOperationRequest) GetIncludeAllSStablesOk() (*bool, bool)`

GetIncludeAllSStablesOk returns a tuple with the IncludeAllSStables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllSStables

`func (o *UpgradeSSTablesOperationRequest) SetIncludeAllSStables(v bool)`

SetIncludeAllSStables sets IncludeAllSStables field to given value.

### HasIncludeAllSStables

`func (o *UpgradeSSTablesOperationRequest) HasIncludeAllSStables() bool`

HasIncludeAllSStables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


