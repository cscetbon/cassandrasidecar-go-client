# DrainOperationResponse

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

## Methods

### NewDrainOperationResponse

`func NewDrainOperationResponse(type_ string, id string, state string, progress float32, creationTime time.Time, ) *DrainOperationResponse`

NewDrainOperationResponse instantiates a new DrainOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrainOperationResponseWithDefaults

`func NewDrainOperationResponseWithDefaults() *DrainOperationResponse`

NewDrainOperationResponseWithDefaults instantiates a new DrainOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DrainOperationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DrainOperationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DrainOperationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *DrainOperationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DrainOperationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DrainOperationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *DrainOperationResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DrainOperationResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DrainOperationResponse) SetState(v string)`

SetState sets State field to given value.


### GetProgress

`func (o *DrainOperationResponse) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DrainOperationResponse) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DrainOperationResponse) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetCreationTime

`func (o *DrainOperationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DrainOperationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DrainOperationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetStartTime

`func (o *DrainOperationResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DrainOperationResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DrainOperationResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DrainOperationResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *DrainOperationResponse) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *DrainOperationResponse) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *DrainOperationResponse) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *DrainOperationResponse) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetFailureCause

`func (o *DrainOperationResponse) GetFailureCause() map[string]interface{}`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *DrainOperationResponse) GetFailureCauseOk() (*map[string]interface{}, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *DrainOperationResponse) SetFailureCause(v map[string]interface{})`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *DrainOperationResponse) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


