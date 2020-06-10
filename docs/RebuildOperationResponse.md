# RebuildOperationResponse

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
**Keyspace** | Pointer to **string** | specific keyspace to rebuild, if not specified, all keyspaces are rebuilt  | [optional] 
**SourceDC** | Pointer to **string** | name of DC from which to select sources for streaming, by default, pick any DC  | [optional] 
**SpecificTokens** | Pointer to [**[]TokenRange**](TokenRange.md) | rebuild specific token ranges  | [optional] 
**SpecificSources** | Pointer to **[]string** | specify hosts that this node should stream from when specificTokens are used  | [optional] 

## Methods

### NewRebuildOperationResponse

`func NewRebuildOperationResponse(type_ string, id string, state string, progress float32, creationTime time.Time, ) *RebuildOperationResponse`

NewRebuildOperationResponse instantiates a new RebuildOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebuildOperationResponseWithDefaults

`func NewRebuildOperationResponseWithDefaults() *RebuildOperationResponse`

NewRebuildOperationResponseWithDefaults instantiates a new RebuildOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RebuildOperationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RebuildOperationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RebuildOperationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *RebuildOperationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RebuildOperationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RebuildOperationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *RebuildOperationResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RebuildOperationResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RebuildOperationResponse) SetState(v string)`

SetState sets State field to given value.


### GetProgress

`func (o *RebuildOperationResponse) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RebuildOperationResponse) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RebuildOperationResponse) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetCreationTime

`func (o *RebuildOperationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RebuildOperationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RebuildOperationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetStartTime

`func (o *RebuildOperationResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RebuildOperationResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RebuildOperationResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RebuildOperationResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *RebuildOperationResponse) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *RebuildOperationResponse) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *RebuildOperationResponse) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *RebuildOperationResponse) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetFailureCause

`func (o *RebuildOperationResponse) GetFailureCause() map[string]interface{}`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *RebuildOperationResponse) GetFailureCauseOk() (*map[string]interface{}, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *RebuildOperationResponse) SetFailureCause(v map[string]interface{})`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *RebuildOperationResponse) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetKeyspace

`func (o *RebuildOperationResponse) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *RebuildOperationResponse) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *RebuildOperationResponse) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *RebuildOperationResponse) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetSourceDC

`func (o *RebuildOperationResponse) GetSourceDC() string`

GetSourceDC returns the SourceDC field if non-nil, zero value otherwise.

### GetSourceDCOk

`func (o *RebuildOperationResponse) GetSourceDCOk() (*string, bool)`

GetSourceDCOk returns a tuple with the SourceDC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDC

`func (o *RebuildOperationResponse) SetSourceDC(v string)`

SetSourceDC sets SourceDC field to given value.

### HasSourceDC

`func (o *RebuildOperationResponse) HasSourceDC() bool`

HasSourceDC returns a boolean if a field has been set.

### GetSpecificTokens

`func (o *RebuildOperationResponse) GetSpecificTokens() []TokenRange`

GetSpecificTokens returns the SpecificTokens field if non-nil, zero value otherwise.

### GetSpecificTokensOk

`func (o *RebuildOperationResponse) GetSpecificTokensOk() (*[]TokenRange, bool)`

GetSpecificTokensOk returns a tuple with the SpecificTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificTokens

`func (o *RebuildOperationResponse) SetSpecificTokens(v []TokenRange)`

SetSpecificTokens sets SpecificTokens field to given value.

### HasSpecificTokens

`func (o *RebuildOperationResponse) HasSpecificTokens() bool`

HasSpecificTokens returns a boolean if a field has been set.

### GetSpecificSources

`func (o *RebuildOperationResponse) GetSpecificSources() []string`

GetSpecificSources returns the SpecificSources field if non-nil, zero value otherwise.

### GetSpecificSourcesOk

`func (o *RebuildOperationResponse) GetSpecificSourcesOk() (*[]string, bool)`

GetSpecificSourcesOk returns a tuple with the SpecificSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificSources

`func (o *RebuildOperationResponse) SetSpecificSources(v []string)`

SetSpecificSources sets SpecificSources field to given value.

### HasSpecificSources

`func (o *RebuildOperationResponse) HasSpecificSources() bool`

HasSpecificSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


