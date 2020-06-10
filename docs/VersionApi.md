# \VersionApi

All URIs are relative to *http://localhost:4567*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VersionCassandraGet**](VersionApi.md#VersionCassandraGet) | **Get** /version/cassandra | returns version of Cassandra node
[**VersionGet**](VersionApi.md#VersionGet) | **Get** /version | returns version of Cassandra Sidecar itself
[**VersionSchemaGet**](VersionApi.md#VersionSchemaGet) | **Get** /version/schema | returns schema version this Cassandra node is on, same as calling StorageServiceMBean#getSchemaVersion
[**VersionSidecarGet**](VersionApi.md#VersionSidecarGet) | **Get** /version/sidecar | alias for /version endpoint, returns version of Cassandra Sidecar itself



## VersionCassandraGet

> CassandraVersion VersionCassandraGet(ctx).Execute()

returns version of Cassandra node

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionCassandraGetRequest struct via the builder pattern


### Return type

[**CassandraVersion**](CassandraVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionGet

> SidecarVersion VersionGet(ctx).Execute()

returns version of Cassandra Sidecar itself

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionGetRequest struct via the builder pattern


### Return type

[**SidecarVersion**](SidecarVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionSchemaGet

> CassandraSchemaVersion VersionSchemaGet(ctx).Execute()

returns schema version this Cassandra node is on, same as calling StorageServiceMBean#getSchemaVersion

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionSchemaGetRequest struct via the builder pattern


### Return type

[**CassandraSchemaVersion**](CassandraSchemaVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionSidecarGet

> SidecarVersion VersionSidecarGet(ctx).Execute()

alias for /version endpoint, returns version of Cassandra Sidecar itself

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionSidecarGetRequest struct via the builder pattern


### Return type

[**SidecarVersion**](SidecarVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

