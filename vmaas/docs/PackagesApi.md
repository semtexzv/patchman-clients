# \PackagesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPackagesHandlerGetGet**](PackagesApi.md#AppPackagesHandlerGetGet) | **Get** /v1/packages/{nevra} | 
[**AppPackagesHandlerPostPost**](PackagesApi.md#AppPackagesHandlerPostPost) | **Post** /v1/packages | 



## AppPackagesHandlerGetGet

> PackagesResponse AppPackagesHandlerGetGet(ctx, nevra)



Get details about packages.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**PackagesResponse**](PackagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPackagesHandlerPostPost

> PackagesResponse AppPackagesHandlerPostPost(ctx, optional)



Get details about packages. \"package_list\" must be a list of package NEVRAs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppPackagesHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPackagesHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packagesRequest** | [**optional.Interface of PackagesRequest**](PackagesRequest.md)|  | 

### Return type

[**PackagesResponse**](PackagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

