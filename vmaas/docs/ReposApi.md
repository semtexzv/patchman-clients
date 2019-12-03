# \ReposApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppReposHandlerGetGet**](ReposApi.md#AppReposHandlerGetGet) | **Get** /v1/repos/{repo} | 
[**AppReposHandlerPostPost**](ReposApi.md#AppReposHandlerPostPost) | **Post** /v1/repos | 



## AppReposHandlerGetGet

> ReposResponse AppReposHandlerGetGet(ctx, repo)



Get details about a repository or repository-expression. It is allowed to use POSIX regular expression as a pattern for repository names.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string**| Repository name or POSIX regular expression pattern | 

### Return type

[**ReposResponse**](ReposResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppReposHandlerPostPost

> ReposResponse AppReposHandlerPostPost(ctx, optional)



Get details about list of repositories. \"repository_list\" can be either a list of repository names, OR a single POSIX regular expression.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppReposHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppReposHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reposRequest** | [**optional.Interface of ReposRequest**](ReposRequest.md)|  | 

### Return type

[**ReposResponse**](ReposResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

