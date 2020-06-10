# SidecarVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | textual representation of Sidecar version | [optional] 
**BuildTime** | Pointer to [**time.Time**](time.Time.md) | timestamp this Sidecar was built | [optional] 
**GitCommit** | Pointer to **string** | git commit hash this Sidecar was built from | [optional] 

## Methods

### NewSidecarVersion

`func NewSidecarVersion() *SidecarVersion`

NewSidecarVersion instantiates a new SidecarVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSidecarVersionWithDefaults

`func NewSidecarVersionWithDefaults() *SidecarVersion`

NewSidecarVersionWithDefaults instantiates a new SidecarVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SidecarVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SidecarVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SidecarVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SidecarVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBuildTime

`func (o *SidecarVersion) GetBuildTime() time.Time`

GetBuildTime returns the BuildTime field if non-nil, zero value otherwise.

### GetBuildTimeOk

`func (o *SidecarVersion) GetBuildTimeOk() (*time.Time, bool)`

GetBuildTimeOk returns a tuple with the BuildTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTime

`func (o *SidecarVersion) SetBuildTime(v time.Time)`

SetBuildTime sets BuildTime field to given value.

### HasBuildTime

`func (o *SidecarVersion) HasBuildTime() bool`

HasBuildTime returns a boolean if a field has been set.

### GetGitCommit

`func (o *SidecarVersion) GetGitCommit() string`

GetGitCommit returns the GitCommit field if non-nil, zero value otherwise.

### GetGitCommitOk

`func (o *SidecarVersion) GetGitCommitOk() (*string, bool)`

GetGitCommitOk returns a tuple with the GitCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommit

`func (o *SidecarVersion) SetGitCommit(v string)`

SetGitCommit sets GitCommit field to given value.

### HasGitCommit

`func (o *SidecarVersion) HasGitCommit() bool`

HasGitCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


