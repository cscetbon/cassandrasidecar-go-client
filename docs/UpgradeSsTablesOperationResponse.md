# UpgradeSSTablesOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Id** | Pointer to **string** | unique identifier of an operation, a random id is assigned to each operation after a request is submitted, from caller&#39;s perspective, an id is sent back as a response to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint  | 
**State** | Pointer to **string** | state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively doing its job, COMPLETED - this operation has finished successfully, CANCELLED - this operation was interrupted while being run, FAILED - this operation has finished errorneously  | 
**Progress** | Pointer to **float32** | float from 0.0 to 1.0, 1.0 telling that operation is completed, either successfully or with errors.  | 
**CreationTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when this operation was created on Sidecar&#39;s side  | 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when this operation was started by Sidecar, if an operation is created, it does not necessarily mean that it will be started right away, in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that  | [optional] 
**CompletionTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when an operation has finished, irrelevant of its result, an operation can be failed and it would still have this field populated.  | [optional] 
**FailureCause** | Pointer to [**map[string]interface{}**](.md) | This field contains serialized java.lang.Throwable in case this operation has failed  | [optional] 
**Keyspace** | Pointer to **string** | keyspace to upgrade SSTables of  | 
**Tables** | Pointer to **[]string** | an array of tables to upgrade SSTables of, empty or not provided array will default to upgrading of SSTables of all tables in respective keyspace  | [optional] 
**Jobs** | Pointer to **int32** | the number of threads to use - 0 means use all available, it never uses more than concurrent_compactor threads  | [optional] 
**IncludeAllSStables** | Pointer to **bool** | include all sstables, even those already on the current version, defaults to false | [optional] 

## Methods

### NewUpgradeSSTablesOperationResponse

`func NewUpgradeSSTablesOperationResponse(type_ string, id string, state string, progress float32, creationTime time.Time, keyspace string, ) *UpgradeSSTablesOperationResponse`

NewUpgradeSSTablesOperationResponse instantiates a new UpgradeSSTablesOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeSSTablesOperationResponseWithDefaults

`func NewUpgradeSSTablesOperationResponseWithDefaults() *UpgradeSSTablesOperationResponse`

NewUpgradeSSTablesOperationResponseWithDefaults instantiates a new UpgradeSSTablesOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpgradeSSTablesOperationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpgradeSSTablesOperationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpgradeSSTablesOperationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *UpgradeSSTablesOperationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpgradeSSTablesOperationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpgradeSSTablesOperationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *UpgradeSSTablesOperationResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpgradeSSTablesOperationResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpgradeSSTablesOperationResponse) SetState(v string)`

SetState sets State field to given value.


### GetProgress

`func (o *UpgradeSSTablesOperationResponse) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *UpgradeSSTablesOperationResponse) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *UpgradeSSTablesOperationResponse) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetCreationTime

`func (o *UpgradeSSTablesOperationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *UpgradeSSTablesOperationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *UpgradeSSTablesOperationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetStartTime

`func (o *UpgradeSSTablesOperationResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpgradeSSTablesOperationResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpgradeSSTablesOperationResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpgradeSSTablesOperationResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *UpgradeSSTablesOperationResponse) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *UpgradeSSTablesOperationResponse) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *UpgradeSSTablesOperationResponse) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *UpgradeSSTablesOperationResponse) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetFailureCause

`func (o *UpgradeSSTablesOperationResponse) GetFailureCause() map[string]interface{}`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *UpgradeSSTablesOperationResponse) GetFailureCauseOk() (*map[string]interface{}, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *UpgradeSSTablesOperationResponse) SetFailureCause(v map[string]interface{})`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *UpgradeSSTablesOperationResponse) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetKeyspace

`func (o *UpgradeSSTablesOperationResponse) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *UpgradeSSTablesOperationResponse) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *UpgradeSSTablesOperationResponse) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.


### GetTables

`func (o *UpgradeSSTablesOperationResponse) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *UpgradeSSTablesOperationResponse) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *UpgradeSSTablesOperationResponse) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *UpgradeSSTablesOperationResponse) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetJobs

`func (o *UpgradeSSTablesOperationResponse) GetJobs() int32`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *UpgradeSSTablesOperationResponse) GetJobsOk() (*int32, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *UpgradeSSTablesOperationResponse) SetJobs(v int32)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *UpgradeSSTablesOperationResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetIncludeAllSStables

`func (o *UpgradeSSTablesOperationResponse) GetIncludeAllSStables() bool`

GetIncludeAllSStables returns the IncludeAllSStables field if non-nil, zero value otherwise.

### GetIncludeAllSStablesOk

`func (o *UpgradeSSTablesOperationResponse) GetIncludeAllSStablesOk() (*bool, bool)`

GetIncludeAllSStablesOk returns a tuple with the IncludeAllSStables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllSStables

`func (o *UpgradeSSTablesOperationResponse) SetIncludeAllSStables(v bool)`

SetIncludeAllSStables sets IncludeAllSStables field to given value.

### HasIncludeAllSStables

`func (o *UpgradeSSTablesOperationResponse) HasIncludeAllSStables() bool`

HasIncludeAllSStables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


