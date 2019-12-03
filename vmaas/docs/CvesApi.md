# \CvesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCVEHandlerGetGet**](CvesApi.md#AppCVEHandlerGetGet) | **Get** /v1/cves/{cve} | 
[**AppCVEHandlerPostPost**](CvesApi.md#AppCVEHandlerPostPost) | **Post** /v1/cves | 



## AppCVEHandlerGetGet

> CvesResponse AppCVEHandlerGetGet(ctx, cve)



Get details about CVEs. It is possible to use POSIX regular expression as a pattern for CVE names.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cve** | **string**| CVE name or POSIX regular expression pattern | 

### Return type

[**CvesResponse**](CvesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCVEHandlerPostPost

> CvesResponse AppCVEHandlerPostPost(ctx, optional)



Get details about CVEs with additional parameters. As a \"cve_list\" parameter a complete list of CVE names can be provided OR one POSIX regular expression.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppCVEHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCVEHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cvesRequest** | [**optional.Interface of CvesRequest**](CvesRequest.md)|  | 

### Return type

[**CvesResponse**](CvesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

