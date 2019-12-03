# \UpdatesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppUpdatesHandlerGetGet**](UpdatesApi.md#AppUpdatesHandlerGetGet) | **Get** /v1/updates/{nevra} | 
[**AppUpdatesHandlerPostPost**](UpdatesApi.md#AppUpdatesHandlerPostPost) | **Post** /v1/updates | 
[**AppUpdatesHandlerV2GetGet**](UpdatesApi.md#AppUpdatesHandlerV2GetGet) | **Get** /v2/updates/{nevra} | 
[**AppUpdatesHandlerV2PostPost**](UpdatesApi.md#AppUpdatesHandlerV2PostPost) | **Post** /v2/updates | 
[**AppUpdatesHandlerV3GetGet**](UpdatesApi.md#AppUpdatesHandlerV3GetGet) | **Get** /v3/updates/{nevra} | 
[**AppUpdatesHandlerV3PostPost**](UpdatesApi.md#AppUpdatesHandlerV3PostPost) | **Post** /v3/updates | 



## AppUpdatesHandlerGetGet

> UpdatesResponse AppUpdatesHandlerGetGet(ctx, nevra)



List security updates for single package NEVRA

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**UpdatesResponse**](UpdatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerPostPost

> UpdatesResponse AppUpdatesHandlerPostPost(ctx, optional)



List security updates for list of package NEVRAs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppUpdatesHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppUpdatesHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatesRequest** | [**optional.Interface of UpdatesRequest**](UpdatesRequest.md)| Input json | 

### Return type

[**UpdatesResponse**](UpdatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerV2GetGet

> UpdatesV2Response AppUpdatesHandlerV2GetGet(ctx, nevra)



List security updates for single package NEVRA

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**UpdatesV2Response**](UpdatesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerV2PostPost

> UpdatesV2Response AppUpdatesHandlerV2PostPost(ctx, optional)



List security updates for list of package NEVRAs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppUpdatesHandlerV2PostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppUpdatesHandlerV2PostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatesRequest** | [**optional.Interface of UpdatesRequest**](UpdatesRequest.md)| Input json | 

### Return type

[**UpdatesV2Response**](UpdatesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerV3GetGet

> UpdatesV2Response AppUpdatesHandlerV3GetGet(ctx, nevra)



List all updates for single package NEVRA

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**UpdatesV2Response**](UpdatesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerV3PostPost

> UpdatesV2Response AppUpdatesHandlerV3PostPost(ctx, optional)



List all updates for list of package NEVRAs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppUpdatesHandlerV3PostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppUpdatesHandlerV3PostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatesV3Request** | [**optional.Interface of UpdatesV3Request**](UpdatesV3Request.md)| Input json | 

### Return type

[**UpdatesV2Response**](UpdatesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

