# ScrubOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Keyspace** | Pointer to **string** | keyspace to scrub  | 
**Tables** | Pointer to **[]string** | tables to scrub, empty or not provided will scrub all tables in respective keyspace  | [optional] 
**Jobs** | Pointer to **int32** | number of sstables to scrub simultanously, set to 0 to use all available compaction threads  | [optional] 
**DisableSnapshot** | Pointer to **bool** | scrubbed CFs will be snapshotted first, defaults to false  | [optional] 
**SkipCorrupted** | Pointer to **bool** | skip corrupted partitions even when scrubbing counter tables, defaults to false  | [optional] 
**NoValidate** | Pointer to **bool** | do not validate columns using column validator, defaults to false  | [optional] 
**ReinsertOverflowedTTL** | Pointer to **bool** | Rewrites rows with overflowed expiration date affected by CASSANDRA-14092 with the maximum supported expiration date of 2038-01-19T03:14:06+00:00. The rows are rewritten with the original timestamp incremented by one millisecond to override/supersede any potential tombstone that may have been generated during compaction of the affected rows.  | [optional] 

## Methods

### NewScrubOperationRequest

`func NewScrubOperationRequest(type_ string, keyspace string, ) *ScrubOperationRequest`

NewScrubOperationRequest instantiates a new ScrubOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScrubOperationRequestWithDefaults

`func NewScrubOperationRequestWithDefaults() *ScrubOperationRequest`

NewScrubOperationRequestWithDefaults instantiates a new ScrubOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScrubOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScrubOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScrubOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetKeyspace

`func (o *ScrubOperationRequest) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *ScrubOperationRequest) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *ScrubOperationRequest) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.


### GetTables

`func (o *ScrubOperationRequest) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *ScrubOperationRequest) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *ScrubOperationRequest) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *ScrubOperationRequest) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetJobs

`func (o *ScrubOperationRequest) GetJobs() int32`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ScrubOperationRequest) GetJobsOk() (*int32, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ScrubOperationRequest) SetJobs(v int32)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ScrubOperationRequest) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetDisableSnapshot

`func (o *ScrubOperationRequest) GetDisableSnapshot() bool`

GetDisableSnapshot returns the DisableSnapshot field if non-nil, zero value otherwise.

### GetDisableSnapshotOk

`func (o *ScrubOperationRequest) GetDisableSnapshotOk() (*bool, bool)`

GetDisableSnapshotOk returns a tuple with the DisableSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSnapshot

`func (o *ScrubOperationRequest) SetDisableSnapshot(v bool)`

SetDisableSnapshot sets DisableSnapshot field to given value.

### HasDisableSnapshot

`func (o *ScrubOperationRequest) HasDisableSnapshot() bool`

HasDisableSnapshot returns a boolean if a field has been set.

### GetSkipCorrupted

`func (o *ScrubOperationRequest) GetSkipCorrupted() bool`

GetSkipCorrupted returns the SkipCorrupted field if non-nil, zero value otherwise.

### GetSkipCorruptedOk

`func (o *ScrubOperationRequest) GetSkipCorruptedOk() (*bool, bool)`

GetSkipCorruptedOk returns a tuple with the SkipCorrupted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCorrupted

`func (o *ScrubOperationRequest) SetSkipCorrupted(v bool)`

SetSkipCorrupted sets SkipCorrupted field to given value.

### HasSkipCorrupted

`func (o *ScrubOperationRequest) HasSkipCorrupted() bool`

HasSkipCorrupted returns a boolean if a field has been set.

### GetNoValidate

`func (o *ScrubOperationRequest) GetNoValidate() bool`

GetNoValidate returns the NoValidate field if non-nil, zero value otherwise.

### GetNoValidateOk

`func (o *ScrubOperationRequest) GetNoValidateOk() (*bool, bool)`

GetNoValidateOk returns a tuple with the NoValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoValidate

`func (o *ScrubOperationRequest) SetNoValidate(v bool)`

SetNoValidate sets NoValidate field to given value.

### HasNoValidate

`func (o *ScrubOperationRequest) HasNoValidate() bool`

HasNoValidate returns a boolean if a field has been set.

### GetReinsertOverflowedTTL

`func (o *ScrubOperationRequest) GetReinsertOverflowedTTL() bool`

GetReinsertOverflowedTTL returns the ReinsertOverflowedTTL field if non-nil, zero value otherwise.

### GetReinsertOverflowedTTLOk

`func (o *ScrubOperationRequest) GetReinsertOverflowedTTLOk() (*bool, bool)`

GetReinsertOverflowedTTLOk returns a tuple with the ReinsertOverflowedTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReinsertOverflowedTTL

`func (o *ScrubOperationRequest) SetReinsertOverflowedTTL(v bool)`

SetReinsertOverflowedTTL sets ReinsertOverflowedTTL field to given value.

### HasReinsertOverflowedTTL

`func (o *ScrubOperationRequest) HasReinsertOverflowedTTL() bool`

HasReinsertOverflowedTTL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


