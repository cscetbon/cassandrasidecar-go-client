# RebuildOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | 
**Keyspace** | Pointer to **string** | specific keyspace to rebuild, if not specified, all keyspaces are rebuilt  | [optional] 
**SourceDC** | Pointer to **string** | name of DC from which to select sources for streaming, by default, pick any DC  | [optional] 
**SpecificTokens** | Pointer to [**[]TokenRange**](TokenRange.md) | rebuild specific token ranges  | [optional] 
**SpecificSources** | Pointer to **[]string** | specify hosts that this node should stream from when specificTokens are used  | [optional] 

## Methods

### NewRebuildOperationRequest

`func NewRebuildOperationRequest(type_ string, ) *RebuildOperationRequest`

NewRebuildOperationRequest instantiates a new RebuildOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebuildOperationRequestWithDefaults

`func NewRebuildOperationRequestWithDefaults() *RebuildOperationRequest`

NewRebuildOperationRequestWithDefaults instantiates a new RebuildOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RebuildOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RebuildOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RebuildOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetKeyspace

`func (o *RebuildOperationRequest) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *RebuildOperationRequest) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *RebuildOperationRequest) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *RebuildOperationRequest) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetSourceDC

`func (o *RebuildOperationRequest) GetSourceDC() string`

GetSourceDC returns the SourceDC field if non-nil, zero value otherwise.

### GetSourceDCOk

`func (o *RebuildOperationRequest) GetSourceDCOk() (*string, bool)`

GetSourceDCOk returns a tuple with the SourceDC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDC

`func (o *RebuildOperationRequest) SetSourceDC(v string)`

SetSourceDC sets SourceDC field to given value.

### HasSourceDC

`func (o *RebuildOperationRequest) HasSourceDC() bool`

HasSourceDC returns a boolean if a field has been set.

### GetSpecificTokens

`func (o *RebuildOperationRequest) GetSpecificTokens() []TokenRange`

GetSpecificTokens returns the SpecificTokens field if non-nil, zero value otherwise.

### GetSpecificTokensOk

`func (o *RebuildOperationRequest) GetSpecificTokensOk() (*[]TokenRange, bool)`

GetSpecificTokensOk returns a tuple with the SpecificTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificTokens

`func (o *RebuildOperationRequest) SetSpecificTokens(v []TokenRange)`

SetSpecificTokens sets SpecificTokens field to given value.

### HasSpecificTokens

`func (o *RebuildOperationRequest) HasSpecificTokens() bool`

HasSpecificTokens returns a boolean if a field has been set.

### GetSpecificSources

`func (o *RebuildOperationRequest) GetSpecificSources() []string`

GetSpecificSources returns the SpecificSources field if non-nil, zero value otherwise.

### GetSpecificSourcesOk

`func (o *RebuildOperationRequest) GetSpecificSourcesOk() (*[]string, bool)`

GetSpecificSourcesOk returns a tuple with the SpecificSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificSources

`func (o *RebuildOperationRequest) SetSpecificSources(v []string)`

SetSpecificSources sets SpecificSources field to given value.

### HasSpecificSources

`func (o *RebuildOperationRequest) HasSpecificSources() bool`

HasSpecificSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


