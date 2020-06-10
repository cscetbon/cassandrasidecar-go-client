# ImportOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | has to be set to &#39;import&#39;  | 
**Id** | Pointer to **string** | unique identifier of an operation, a random id is assigned to each operation after a request is submitted, from caller&#39;s perspective, an id is sent back as a response to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint  | 
**State** | Pointer to **string** | state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively doing its job, COMPLETED - this operation has finished successfully, CANCELLED - this operation was interrupted while being run, FAILED - this operation has finished errorneously  | 
**Progress** | Pointer to **float32** | float from 0.0 to 1.0, 1.0 telling that operation is completed, either successfully or with errors.  | 
**CreationTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when this operation was created on Sidecar&#39;s side  | 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when this operation was started by Sidecar, if an operation is created, it does not necessarily mean that it will be started right away, in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that  | [optional] 
**CompletionTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when an operation has finished, irrelevant of its result, an operation can be failed and it would still have this field populated.  | [optional] 
**FailureCause** | Pointer to [**map[string]interface{}**](.md) | This field contains serialized java.lang.Throwable in case this operation has failed  | [optional] 
**Keyspace** | Pointer to **string** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**SourceDir** | Pointer to **string** |  | 
**KeepLevel** | Pointer to **bool** |  | [optional] 
**KeepRepaired** | Pointer to **bool** |  | [optional] 
**NoVerify** | Pointer to **bool** |  | [optional] 
**NoVerifyTokens** | Pointer to **bool** |  | [optional] 
**NoInvalidateCaches** | Pointer to **bool** |  | [optional] 
**Quick** | Pointer to **bool** | defaults to false, if true, noVerifyTokens, noInvalidateCaches and noVerify will be set to true automatically  | [optional] 
**ExtendedVerify** | Pointer to **bool** |  | [optional] 

## Methods

### NewImportOperationResponse

`func NewImportOperationResponse(type_ string, id string, state string, progress float32, creationTime time.Time, sourceDir string, ) *ImportOperationResponse`

NewImportOperationResponse instantiates a new ImportOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportOperationResponseWithDefaults

`func NewImportOperationResponseWithDefaults() *ImportOperationResponse`

NewImportOperationResponseWithDefaults instantiates a new ImportOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ImportOperationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportOperationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportOperationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ImportOperationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportOperationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportOperationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *ImportOperationResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ImportOperationResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ImportOperationResponse) SetState(v string)`

SetState sets State field to given value.


### GetProgress

`func (o *ImportOperationResponse) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ImportOperationResponse) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ImportOperationResponse) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetCreationTime

`func (o *ImportOperationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ImportOperationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ImportOperationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetStartTime

`func (o *ImportOperationResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ImportOperationResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ImportOperationResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ImportOperationResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *ImportOperationResponse) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *ImportOperationResponse) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *ImportOperationResponse) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *ImportOperationResponse) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetFailureCause

`func (o *ImportOperationResponse) GetFailureCause() map[string]interface{}`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *ImportOperationResponse) GetFailureCauseOk() (*map[string]interface{}, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *ImportOperationResponse) SetFailureCause(v map[string]interface{})`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *ImportOperationResponse) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetKeyspace

`func (o *ImportOperationResponse) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *ImportOperationResponse) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *ImportOperationResponse) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *ImportOperationResponse) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetTable

`func (o *ImportOperationResponse) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *ImportOperationResponse) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *ImportOperationResponse) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *ImportOperationResponse) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetSourceDir

`func (o *ImportOperationResponse) GetSourceDir() string`

GetSourceDir returns the SourceDir field if non-nil, zero value otherwise.

### GetSourceDirOk

`func (o *ImportOperationResponse) GetSourceDirOk() (*string, bool)`

GetSourceDirOk returns a tuple with the SourceDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDir

`func (o *ImportOperationResponse) SetSourceDir(v string)`

SetSourceDir sets SourceDir field to given value.


### GetKeepLevel

`func (o *ImportOperationResponse) GetKeepLevel() bool`

GetKeepLevel returns the KeepLevel field if non-nil, zero value otherwise.

### GetKeepLevelOk

`func (o *ImportOperationResponse) GetKeepLevelOk() (*bool, bool)`

GetKeepLevelOk returns a tuple with the KeepLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepLevel

`func (o *ImportOperationResponse) SetKeepLevel(v bool)`

SetKeepLevel sets KeepLevel field to given value.

### HasKeepLevel

`func (o *ImportOperationResponse) HasKeepLevel() bool`

HasKeepLevel returns a boolean if a field has been set.

### GetKeepRepaired

`func (o *ImportOperationResponse) GetKeepRepaired() bool`

GetKeepRepaired returns the KeepRepaired field if non-nil, zero value otherwise.

### GetKeepRepairedOk

`func (o *ImportOperationResponse) GetKeepRepairedOk() (*bool, bool)`

GetKeepRepairedOk returns a tuple with the KeepRepaired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepRepaired

`func (o *ImportOperationResponse) SetKeepRepaired(v bool)`

SetKeepRepaired sets KeepRepaired field to given value.

### HasKeepRepaired

`func (o *ImportOperationResponse) HasKeepRepaired() bool`

HasKeepRepaired returns a boolean if a field has been set.

### GetNoVerify

`func (o *ImportOperationResponse) GetNoVerify() bool`

GetNoVerify returns the NoVerify field if non-nil, zero value otherwise.

### GetNoVerifyOk

`func (o *ImportOperationResponse) GetNoVerifyOk() (*bool, bool)`

GetNoVerifyOk returns a tuple with the NoVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoVerify

`func (o *ImportOperationResponse) SetNoVerify(v bool)`

SetNoVerify sets NoVerify field to given value.

### HasNoVerify

`func (o *ImportOperationResponse) HasNoVerify() bool`

HasNoVerify returns a boolean if a field has been set.

### GetNoVerifyTokens

`func (o *ImportOperationResponse) GetNoVerifyTokens() bool`

GetNoVerifyTokens returns the NoVerifyTokens field if non-nil, zero value otherwise.

### GetNoVerifyTokensOk

`func (o *ImportOperationResponse) GetNoVerifyTokensOk() (*bool, bool)`

GetNoVerifyTokensOk returns a tuple with the NoVerifyTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoVerifyTokens

`func (o *ImportOperationResponse) SetNoVerifyTokens(v bool)`

SetNoVerifyTokens sets NoVerifyTokens field to given value.

### HasNoVerifyTokens

`func (o *ImportOperationResponse) HasNoVerifyTokens() bool`

HasNoVerifyTokens returns a boolean if a field has been set.

### GetNoInvalidateCaches

`func (o *ImportOperationResponse) GetNoInvalidateCaches() bool`

GetNoInvalidateCaches returns the NoInvalidateCaches field if non-nil, zero value otherwise.

### GetNoInvalidateCachesOk

`func (o *ImportOperationResponse) GetNoInvalidateCachesOk() (*bool, bool)`

GetNoInvalidateCachesOk returns a tuple with the NoInvalidateCaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoInvalidateCaches

`func (o *ImportOperationResponse) SetNoInvalidateCaches(v bool)`

SetNoInvalidateCaches sets NoInvalidateCaches field to given value.

### HasNoInvalidateCaches

`func (o *ImportOperationResponse) HasNoInvalidateCaches() bool`

HasNoInvalidateCaches returns a boolean if a field has been set.

### GetQuick

`func (o *ImportOperationResponse) GetQuick() bool`

GetQuick returns the Quick field if non-nil, zero value otherwise.

### GetQuickOk

`func (o *ImportOperationResponse) GetQuickOk() (*bool, bool)`

GetQuickOk returns a tuple with the Quick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuick

`func (o *ImportOperationResponse) SetQuick(v bool)`

SetQuick sets Quick field to given value.

### HasQuick

`func (o *ImportOperationResponse) HasQuick() bool`

HasQuick returns a boolean if a field has been set.

### GetExtendedVerify

`func (o *ImportOperationResponse) GetExtendedVerify() bool`

GetExtendedVerify returns the ExtendedVerify field if non-nil, zero value otherwise.

### GetExtendedVerifyOk

`func (o *ImportOperationResponse) GetExtendedVerifyOk() (*bool, bool)`

GetExtendedVerifyOk returns a tuple with the ExtendedVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedVerify

`func (o *ImportOperationResponse) SetExtendedVerify(v bool)`

SetExtendedVerify sets ExtendedVerify field to given value.

### HasExtendedVerify

`func (o *ImportOperationResponse) HasExtendedVerify() bool`

HasExtendedVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


