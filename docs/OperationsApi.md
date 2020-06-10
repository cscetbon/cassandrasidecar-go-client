# \OperationsApi

All URIs are relative to *http://localhost:4567*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperationsGet**](OperationsApi.md#OperationsGet) | **Get** /operations | All operations of Sidecar
[**OperationsOperationIdGet**](OperationsApi.md#OperationsOperationIdGet) | **Get** /operations/{operationId} | abc
[**OperationsPost**](OperationsApi.md#OperationsPost) | **Post** /operations | Submits an operation to this Sidecar



## OperationsGet

> OperationsGet(ctx).Type_(type_).Status(status).Execute()

All operations of Sidecar

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**[]string**](string.md) | type of operations to filter on | 
 **status** | [**[]string**](string.md) | status of operations to filter on | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsOperationIdGet

> OperationsOperationIdGet(ctx, operationId).Execute()

abc

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | [**string**](.md) | ID of operation to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperationsOperationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperationsPost

> OperationsPOST200 OperationsPost(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Submits an operation to this Sidecar

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) |  | 

### Return type

[**OperationsPOST200**](OperationsPOST200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

