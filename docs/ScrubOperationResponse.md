# ScrubOperationResponse

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
**Keyspace** | Pointer to **string** | keyspace to scrub  | 
**Tables** | Pointer to **[]string** | tables to scrub, empty or not provided will scrub all tables in respective keyspace  | [optional] 
**Jobs** | Pointer to **int32** | number of sstables to scrub simultanously, set to 0 to use all available compaction threads  | [optional] 
**DisableSnapshot** | Pointer to **bool** | scrubbed CFs will be snapshotted first, defaults to false  | [optional] 
**SkipCorrupted** | Pointer to **bool** | skip corrupted partitions even when scrubbing counter tables, defaults to false  | [optional] 
**NoValidate** | Pointer to **bool** | do not validate columns using column validator, defaults to false  | [optional] 
**ReinsertOverflowedTTL** | Pointer to **bool** | Rewrites rows with overflowed expiration date affected by CASSANDRA-14092 with the maximum supported expiration date of 2038-01-19T03:14:06+00:00. The rows are rewritten with the original timestamp incremented by one millisecond to override/supersede any potential tombstone that may have been generated during compaction of the affected rows.  | [optional] 

## Methods

### NewScrubOperationResponse

`func NewScrubOperationResponse(type_ string, id string, state string, progress float32, creationTime time.Time, keyspace string, ) *ScrubOperationResponse`

NewScrubOperationResponse instantiates a new ScrubOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScrubOperationResponseWithDefaults

`func NewScrubOperationResponseWithDefaults() *ScrubOperationResponse`

NewScrubOperationResponseWithDefaults instantiates a new ScrubOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScrubOperationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScrubOperationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScrubOperationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ScrubOperationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScrubOperationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScrubOperationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *ScrubOperationResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ScrubOperationResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ScrubOperationResponse) SetState(v string)`

SetState sets State field to given value.


### GetProgress

`func (o *ScrubOperationResponse) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ScrubOperationResponse) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ScrubOperationResponse) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetCreationTime

`func (o *ScrubOperationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ScrubOperationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ScrubOperationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetStartTime

`func (o *ScrubOperationResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScrubOperationResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScrubOperationResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ScrubOperationResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *ScrubOperationResponse) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *ScrubOperationResponse) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *ScrubOperationResponse) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *ScrubOperationResponse) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetFailureCause

`func (o *ScrubOperationResponse) GetFailureCause() map[string]interface{}`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *ScrubOperationResponse) GetFailureCauseOk() (*map[string]interface{}, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *ScrubOperationResponse) SetFailureCause(v map[string]interface{})`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *ScrubOperationResponse) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetKeyspace

`func (o *ScrubOperationResponse) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *ScrubOperationResponse) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *ScrubOperationResponse) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.


### GetTables

`func (o *ScrubOperationResponse) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *ScrubOperationResponse) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *ScrubOperationResponse) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *ScrubOperationResponse) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetJobs

`func (o *ScrubOperationResponse) GetJobs() int32`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ScrubOperationResponse) GetJobsOk() (*int32, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ScrubOperationResponse) SetJobs(v int32)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ScrubOperationResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetDisableSnapshot

`func (o *ScrubOperationResponse) GetDisableSnapshot() bool`

GetDisableSnapshot returns the DisableSnapshot field if non-nil, zero value otherwise.

### GetDisableSnapshotOk

`func (o *ScrubOperationResponse) GetDisableSnapshotOk() (*bool, bool)`

GetDisableSnapshotOk returns a tuple with the DisableSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSnapshot

`func (o *ScrubOperationResponse) SetDisableSnapshot(v bool)`

SetDisableSnapshot sets DisableSnapshot field to given value.

### HasDisableSnapshot

`func (o *ScrubOperationResponse) HasDisableSnapshot() bool`

HasDisableSnapshot returns a boolean if a field has been set.

### GetSkipCorrupted

`func (o *ScrubOperationResponse) GetSkipCorrupted() bool`

GetSkipCorrupted returns the SkipCorrupted field if non-nil, zero value otherwise.

### GetSkipCorruptedOk

`func (o *ScrubOperationResponse) GetSkipCorruptedOk() (*bool, bool)`

GetSkipCorruptedOk returns a tuple with the SkipCorrupted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCorrupted

`func (o *ScrubOperationResponse) SetSkipCorrupted(v bool)`

SetSkipCorrupted sets SkipCorrupted field to given value.

### HasSkipCorrupted

`func (o *ScrubOperationResponse) HasSkipCorrupted() bool`

HasSkipCorrupted returns a boolean if a field has been set.

### GetNoValidate

`func (o *ScrubOperationResponse) GetNoValidate() bool`

GetNoValidate returns the NoValidate field if non-nil, zero value otherwise.

### GetNoValidateOk

`func (o *ScrubOperationResponse) GetNoValidateOk() (*bool, bool)`

GetNoValidateOk returns a tuple with the NoValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoValidate

`func (o *ScrubOperationResponse) SetNoValidate(v bool)`

SetNoValidate sets NoValidate field to given value.

### HasNoValidate

`func (o *ScrubOperationResponse) HasNoValidate() bool`

HasNoValidate returns a boolean if a field has been set.

### GetReinsertOverflowedTTL

`func (o *ScrubOperationResponse) GetReinsertOverflowedTTL() bool`

GetReinsertOverflowedTTL returns the ReinsertOverflowedTTL field if non-nil, zero value otherwise.

### GetReinsertOverflowedTTLOk

`func (o *ScrubOperationResponse) GetReinsertOverflowedTTLOk() (*bool, bool)`

GetReinsertOverflowedTTLOk returns a tuple with the ReinsertOverflowedTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReinsertOverflowedTTL

`func (o *ScrubOperationResponse) SetReinsertOverflowedTTL(v bool)`

SetReinsertOverflowedTTL sets ReinsertOverflowedTTL field to given value.

### HasReinsertOverflowedTTL

`func (o *ScrubOperationResponse) HasReinsertOverflowedTTL() bool`

HasReinsertOverflowedTTL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


