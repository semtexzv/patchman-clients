# \VulnerabilitiesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppVulnerabilitiesHandlerGetGet**](VulnerabilitiesApi.md#AppVulnerabilitiesHandlerGetGet) | **Get** /v1/vulnerabilities/{nevra} | 
[**AppVulnerabilitiesHandlerPostPost**](VulnerabilitiesApi.md#AppVulnerabilitiesHandlerPostPost) | **Post** /v1/vulnerabilities | 



## AppVulnerabilitiesHandlerGetGet

> VulnerabilitiesResponse AppVulnerabilitiesHandlerGetGet(ctx, nevra)



List of applicable CVEs for a single package NEVRA

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**VulnerabilitiesResponse**](VulnerabilitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVulnerabilitiesHandlerPostPost

> VulnerabilitiesResponse AppVulnerabilitiesHandlerPostPost(ctx, optional)



List of applicable CVEs to a package list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppVulnerabilitiesHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppVulnerabilitiesHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vulnerabilitiesRequest** | [**optional.Interface of VulnerabilitiesRequest**](VulnerabilitiesRequest.md)|  | 

### Return type

[**VulnerabilitiesResponse**](VulnerabilitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

