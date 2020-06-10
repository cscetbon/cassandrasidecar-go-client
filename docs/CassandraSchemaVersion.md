# CassandraSchemaVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** | UUID representing a version of Cassandra node as reported by its StorageServiceMBean#getSchemaVersion | [optional] 

## Methods

### NewCassandraSchemaVersion

`func NewCassandraSchemaVersion() *CassandraSchemaVersion`

NewCassandraSchemaVersion instantiates a new CassandraSchemaVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraSchemaVersionWithDefaults

`func NewCassandraSchemaVersionWithDefaults() *CassandraSchemaVersion`

NewCassandraSchemaVersionWithDefaults instantiates a new CassandraSchemaVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *CassandraSchemaVersion) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CassandraSchemaVersion) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CassandraSchemaVersion) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CassandraSchemaVersion) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


