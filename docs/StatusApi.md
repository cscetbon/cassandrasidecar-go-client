# \StatusApi

All URIs are relative to *http://localhost:4567*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusGet**](StatusApi.md#StatusGet) | **Get** /status | returns a state of a Cassandra node



## StatusGet

> CassandraStatus StatusGet(ctx).Execute()

returns a state of a Cassandra node

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusGetRequest struct via the builder pattern


### Return type

[**CassandraStatus**](CassandraStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

