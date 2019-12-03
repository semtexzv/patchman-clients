# \ErrataApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppErrataHandlerGetGet**](ErrataApi.md#AppErrataHandlerGetGet) | **Get** /v1/errata/{erratum} | 
[**AppErrataHandlerPostPost**](ErrataApi.md#AppErrataHandlerPostPost) | **Post** /v1/errata | 



## AppErrataHandlerGetGet

> ErrataResponse AppErrataHandlerGetGet(ctx, erratum)



Get details about errata. It is possible to use POSIX regular expression as a pattern for errata names.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**erratum** | **string**| Errata advisory name or POSIX regular expression pattern | 

### Return type

[**ErrataResponse**](ErrataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppErrataHandlerPostPost

> ErrataResponse AppErrataHandlerPostPost(ctx, optional)



Get details about errata with additional parameters. \"errata_list\" parameter can be either a list of errata names OR a single POSIX regular expression.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppErrataHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppErrataHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **errataRequest** | [**optional.Interface of ErrataRequest**](ErrataRequest.md)|  | 

### Return type

[**ErrataResponse**](ErrataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

