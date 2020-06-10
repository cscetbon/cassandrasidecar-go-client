# ImportOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | has to be set to &#39;import&#39;  | 
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

### NewImportOperationRequest

`func NewImportOperationRequest(type_ string, sourceDir string, ) *ImportOperationRequest`

NewImportOperationRequest instantiates a new ImportOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportOperationRequestWithDefaults

`func NewImportOperationRequestWithDefaults() *ImportOperationRequest`

NewImportOperationRequestWithDefaults instantiates a new ImportOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ImportOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetKeyspace

`func (o *ImportOperationRequest) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *ImportOperationRequest) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *ImportOperationRequest) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *ImportOperationRequest) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetTable

`func (o *ImportOperationRequest) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *ImportOperationRequest) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *ImportOperationRequest) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *ImportOperationRequest) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetSourceDir

`func (o *ImportOperationRequest) GetSourceDir() string`

GetSourceDir returns the SourceDir field if non-nil, zero value otherwise.

### GetSourceDirOk

`func (o *ImportOperationRequest) GetSourceDirOk() (*string, bool)`

GetSourceDirOk returns a tuple with the SourceDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDir

`func (o *ImportOperationRequest) SetSourceDir(v string)`

SetSourceDir sets SourceDir field to given value.


### GetKeepLevel

`func (o *ImportOperationRequest) GetKeepLevel() bool`

GetKeepLevel returns the KeepLevel field if non-nil, zero value otherwise.

### GetKeepLevelOk

`func (o *ImportOperationRequest) GetKeepLevelOk() (*bool, bool)`

GetKeepLevelOk returns a tuple with the KeepLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepLevel

`func (o *ImportOperationRequest) SetKeepLevel(v bool)`

SetKeepLevel sets KeepLevel field to given value.

### HasKeepLevel

`func (o *ImportOperationRequest) HasKeepLevel() bool`

HasKeepLevel returns a boolean if a field has been set.

### GetKeepRepaired

`func (o *ImportOperationRequest) GetKeepRepaired() bool`

GetKeepRepaired returns the KeepRepaired field if non-nil, zero value otherwise.

### GetKeepRepairedOk

`func (o *ImportOperationRequest) GetKeepRepairedOk() (*bool, bool)`

GetKeepRepairedOk returns a tuple with the KeepRepaired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepRepaired

`func (o *ImportOperationRequest) SetKeepRepaired(v bool)`

SetKeepRepaired sets KeepRepaired field to given value.

### HasKeepRepaired

`func (o *ImportOperationRequest) HasKeepRepaired() bool`

HasKeepRepaired returns a boolean if a field has been set.

### GetNoVerify

`func (o *ImportOperationRequest) GetNoVerify() bool`

GetNoVerify returns the NoVerify field if non-nil, zero value otherwise.

### GetNoVerifyOk

`func (o *ImportOperationRequest) GetNoVerifyOk() (*bool, bool)`

GetNoVerifyOk returns a tuple with the NoVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoVerify

`func (o *ImportOperationRequest) SetNoVerify(v bool)`

SetNoVerify sets NoVerify field to given value.

### HasNoVerify

`func (o *ImportOperationRequest) HasNoVerify() bool`

HasNoVerify returns a boolean if a field has been set.

### GetNoVerifyTokens

`func (o *ImportOperationRequest) GetNoVerifyTokens() bool`

GetNoVerifyTokens returns the NoVerifyTokens field if non-nil, zero value otherwise.

### GetNoVerifyTokensOk

`func (o *ImportOperationRequest) GetNoVerifyTokensOk() (*bool, bool)`

GetNoVerifyTokensOk returns a tuple with the NoVerifyTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoVerifyTokens

`func (o *ImportOperationRequest) SetNoVerifyTokens(v bool)`

SetNoVerifyTokens sets NoVerifyTokens field to given value.

### HasNoVerifyTokens

`func (o *ImportOperationRequest) HasNoVerifyTokens() bool`

HasNoVerifyTokens returns a boolean if a field has been set.

### GetNoInvalidateCaches

`func (o *ImportOperationRequest) GetNoInvalidateCaches() bool`

GetNoInvalidateCaches returns the NoInvalidateCaches field if non-nil, zero value otherwise.

### GetNoInvalidateCachesOk

`func (o *ImportOperationRequest) GetNoInvalidateCachesOk() (*bool, bool)`

GetNoInvalidateCachesOk returns a tuple with the NoInvalidateCaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoInvalidateCaches

`func (o *ImportOperationRequest) SetNoInvalidateCaches(v bool)`

SetNoInvalidateCaches sets NoInvalidateCaches field to given value.

### HasNoInvalidateCaches

`func (o *ImportOperationRequest) HasNoInvalidateCaches() bool`

HasNoInvalidateCaches returns a boolean if a field has been set.

### GetQuick

`func (o *ImportOperationRequest) GetQuick() bool`

GetQuick returns the Quick field if non-nil, zero value otherwise.

### GetQuickOk

`func (o *ImportOperationRequest) GetQuickOk() (*bool, bool)`

GetQuickOk returns a tuple with the Quick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuick

`func (o *ImportOperationRequest) SetQuick(v bool)`

SetQuick sets Quick field to given value.

### HasQuick

`func (o *ImportOperationRequest) HasQuick() bool`

HasQuick returns a boolean if a field has been set.

### GetExtendedVerify

`func (o *ImportOperationRequest) GetExtendedVerify() bool`

GetExtendedVerify returns the ExtendedVerify field if non-nil, zero value otherwise.

### GetExtendedVerifyOk

`func (o *ImportOperationRequest) GetExtendedVerifyOk() (*bool, bool)`

GetExtendedVerifyOk returns a tuple with the ExtendedVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedVerify

`func (o *ImportOperationRequest) SetExtendedVerify(v bool)`

SetExtendedVerify sets ExtendedVerify field to given value.

### HasExtendedVerify

`func (o *ImportOperationRequest) HasExtendedVerify() bool`

HasExtendedVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


