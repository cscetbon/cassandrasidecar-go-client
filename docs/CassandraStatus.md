# CassandraStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeState** | Pointer to **string** | State of a Cassandra node as reported from StorageServiceMBean#getOperationMode | [optional] 

## Methods

### NewCassandraStatus

`func NewCassandraStatus() *CassandraStatus`

NewCassandraStatus instantiates a new CassandraStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraStatusWithDefaults

`func NewCassandraStatusWithDefaults() *CassandraStatus`

NewCassandraStatusWithDefaults instantiates a new CassandraStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeState

`func (o *CassandraStatus) GetNodeState() string`

GetNodeState returns the NodeState field if non-nil, zero value otherwise.

### GetNodeStateOk

`func (o *CassandraStatus) GetNodeStateOk() (*string, bool)`

GetNodeStateOk returns a tuple with the NodeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeState

`func (o *CassandraStatus) SetNodeState(v string)`

SetNodeState sets NodeState field to given value.

### HasNodeState

`func (o *CassandraStatus) HasNodeState() bool`

HasNodeState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


