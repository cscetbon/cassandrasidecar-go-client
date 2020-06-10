# CassandraStatusException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exception** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 

## Methods

### NewCassandraStatusException

`func NewCassandraStatusException() *CassandraStatusException`

NewCassandraStatusException instantiates a new CassandraStatusException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraStatusExceptionWithDefaults

`func NewCassandraStatusExceptionWithDefaults() *CassandraStatusException`

NewCassandraStatusExceptionWithDefaults instantiates a new CassandraStatusException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetException

`func (o *CassandraStatusException) GetException() map[string]interface{}`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *CassandraStatusException) GetExceptionOk() (*map[string]interface{}, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *CassandraStatusException) SetException(v map[string]interface{})`

SetException sets Exception field to given value.

### HasException

`func (o *CassandraStatusException) HasException() bool`

HasException returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


