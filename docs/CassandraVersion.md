# CassandraVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Major** | Pointer to **int32** |  | [optional] 
**Minor** | Pointer to **int32** |  | [optional] 
**Patch** | Pointer to **int32** |  | [optional] 
**DsePatch** | Pointer to **int32** |  | [optional] 
**PreReleaseLabels** | Pointer to **[]string** | Labels for Cassandra version, e.g. SNAPSHOT | [optional] 

## Methods

### NewCassandraVersion

`func NewCassandraVersion() *CassandraVersion`

NewCassandraVersion instantiates a new CassandraVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraVersionWithDefaults

`func NewCassandraVersionWithDefaults() *CassandraVersion`

NewCassandraVersionWithDefaults instantiates a new CassandraVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMajor

`func (o *CassandraVersion) GetMajor() int32`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *CassandraVersion) GetMajorOk() (*int32, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *CassandraVersion) SetMajor(v int32)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *CassandraVersion) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinor

`func (o *CassandraVersion) GetMinor() int32`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *CassandraVersion) GetMinorOk() (*int32, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *CassandraVersion) SetMinor(v int32)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *CassandraVersion) HasMinor() bool`

HasMinor returns a boolean if a field has been set.

### GetPatch

`func (o *CassandraVersion) GetPatch() int32`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *CassandraVersion) GetPatchOk() (*int32, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *CassandraVersion) SetPatch(v int32)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *CassandraVersion) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetDsePatch

`func (o *CassandraVersion) GetDsePatch() int32`

GetDsePatch returns the DsePatch field if non-nil, zero value otherwise.

### GetDsePatchOk

`func (o *CassandraVersion) GetDsePatchOk() (*int32, bool)`

GetDsePatchOk returns a tuple with the DsePatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsePatch

`func (o *CassandraVersion) SetDsePatch(v int32)`

SetDsePatch sets DsePatch field to given value.

### HasDsePatch

`func (o *CassandraVersion) HasDsePatch() bool`

HasDsePatch returns a boolean if a field has been set.

### GetPreReleaseLabels

`func (o *CassandraVersion) GetPreReleaseLabels() []string`

GetPreReleaseLabels returns the PreReleaseLabels field if non-nil, zero value otherwise.

### GetPreReleaseLabelsOk

`func (o *CassandraVersion) GetPreReleaseLabelsOk() (*[]string, bool)`

GetPreReleaseLabelsOk returns a tuple with the PreReleaseLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreReleaseLabels

`func (o *CassandraVersion) SetPreReleaseLabels(v []string)`

SetPreReleaseLabels sets PreReleaseLabels field to given value.

### HasPreReleaseLabels

`func (o *CassandraVersion) HasPreReleaseLabels() bool`

HasPreReleaseLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


