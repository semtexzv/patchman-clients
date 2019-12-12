# \HostsApi

All URIs are relative to *http://localhost/api/inventory/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiHostAddHostList**](HostsApi.md#ApiHostAddHostList) | **Post** /hosts | Create/update multiple host and add them to the host list
[**ApiHostDeleteById**](HostsApi.md#ApiHostDeleteById) | **Delete** /hosts/{host_id_list} | Delete hosts by IDs
[**ApiHostGetHostById**](HostsApi.md#ApiHostGetHostById) | **Get** /hosts/{host_id_list} | Find hosts by their IDs
[**ApiHostGetHostList**](HostsApi.md#ApiHostGetHostList) | **Get** /hosts | Read the entire list of hosts
[**ApiHostGetHostSystemProfileById**](HostsApi.md#ApiHostGetHostSystemProfileById) | **Get** /hosts/{host_id_list}/system_profile | Return one or more hosts system profile
[**ApiHostGetHostTagCount**](HostsApi.md#ApiHostGetHostTagCount) | **Get** /hosts/{host_id_list}/tags/count | Get the number of tags on a host
[**ApiHostGetHostTags**](HostsApi.md#ApiHostGetHostTags) | **Get** /hosts/{host_id_list}/tags | Get the tags on a host
[**ApiHostMergeFacts**](HostsApi.md#ApiHostMergeFacts) | **Patch** /hosts/{host_id_list}/facts/{namespace} | Merge facts under a namespace
[**ApiHostPatchById**](HostsApi.md#ApiHostPatchById) | **Patch** /hosts/{host_id_list} | Update a host
[**ApiHostReplaceFacts**](HostsApi.md#ApiHostReplaceFacts) | **Put** /hosts/{host_id_list}/facts/{namespace} | Replace facts under a namespace



## ApiHostAddHostList

> BulkHostOut ApiHostAddHostList(ctx, createHostIn)

Create/update multiple host and add them to the host list

Create a new host and add it to the host list or update an existing hosts. A host is updated if there is already one with the same canonicals facts and belonging to the same account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createHostIn** | [**[]CreateHostIn**](CreateHostIn.md)| A list of host objects to be added to the host list | 

### Return type

[**BulkHostOut**](BulkHostOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostDeleteById

> ApiHostDeleteById(ctx, hostIdList, optional)

Delete hosts by IDs

Delete hosts by IDs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
 **optional** | ***ApiHostDeleteByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostDeleteByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branchId** | **optional.String**| Filter by branch_id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostById

> HostQueryOutput ApiHostGetHostById(ctx, hostIdList, optional)

Find hosts by their IDs

Find one or more hosts by their ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
 **optional** | ***ApiHostGetHostByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostGetHostByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branchId** | **optional.String**| Filter by branch_id | 
 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **orderBy** | **optional.String**| Ordering field name | 
 **orderHow** | **optional.String**| Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 

### Return type

[**HostQueryOutput**](HostQueryOutput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostList

> HostQueryOutput ApiHostGetHostList(ctx, optional)

Read the entire list of hosts

Read the entire list of all hosts available to the account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiHostGetHostListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostGetHostListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayName** | **optional.String**| A part of a searched host’s display name. | 
 **fqdn** | **optional.String**| Filter by a host&#39;s FQDN | 
 **hostnameOrId** | **optional.String**| Search for a host by display_name, fqdn, id | 
 **insightsId** | [**optional.Interface of string**](.md)| Search for a host by insights_id | 
 **tags** | [**optional.Interface of []string**](string.md)| search for a host by the tags on the system | 
 **branchId** | **optional.String**| Filter by branch_id | 
 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **orderBy** | **optional.String**| Ordering field name | 
 **orderHow** | **optional.String**| Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 
 **staleness** | [**optional.Interface of []string**](string.md)| Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]

### Return type

[**HostQueryOutput**](HostQueryOutput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostSystemProfileById

> SystemProfileByHostOut ApiHostGetHostSystemProfileById(ctx, hostIdList, optional)

Return one or more hosts system profile

Find one or more hosts by their ID and return the id and system profile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
 **optional** | ***ApiHostGetHostSystemProfileByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostGetHostSystemProfileByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **orderBy** | **optional.String**| Ordering field name | 
 **orderHow** | **optional.String**| Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 
 **branchId** | **optional.String**| Filter by branch_id | 

### Return type

[**SystemProfileByHostOut**](SystemProfileByHostOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostTagCount

> TagCountOut ApiHostGetHostTagCount(ctx, hostIdList, optional)

Get the number of tags on a host

Get the number of tags on a host

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
 **optional** | ***ApiHostGetHostTagCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostGetHostTagCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **orderBy** | **optional.String**| Ordering field name | 
 **orderHow** | **optional.String**| Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 

### Return type

[**TagCountOut**](TagCountOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostTags

> TagsOut ApiHostGetHostTags(ctx, hostIdList, optional)

Get the tags on a host

Get the tags on a host

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
 **optional** | ***ApiHostGetHostTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostGetHostTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **orderBy** | **optional.String**| Ordering field name | 
 **orderHow** | **optional.String**| Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 

### Return type

[**TagsOut**](TagsOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostMergeFacts

> ApiHostMergeFacts(ctx, hostIdList, namespace, body, optional)

Merge facts under a namespace

Merge one or multiple hosts facts under a namespace.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
**namespace** | **string**| A namespace of the merged facts. | 
**body** | **map[string]interface{}**| A dictionary with the new facts to merge with the original ones. | 
 **optional** | ***ApiHostMergeFactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostMergeFactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **branchId** | **optional.String**| Filter by branch_id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostPatchById

> ApiHostPatchById(ctx, hostIdList, patchHostIn, optional)

Update a host

Update a host

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
**patchHostIn** | [**PatchHostIn**](PatchHostIn.md)| A group of fields to be updated on the host | 
 **optional** | ***ApiHostPatchByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostPatchByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branchId** | **optional.String**| Filter by branch_id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostReplaceFacts

> ApiHostReplaceFacts(ctx, hostIdList, namespace, body, optional)

Replace facts under a namespace

Replace facts under a namespace

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
**namespace** | **string**| A namespace of the merged facts. | 
**body** | **map[string]interface{}**| A dictionary with the new facts to replace the original ones. | 
 **optional** | ***ApiHostReplaceFactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostReplaceFactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **branchId** | **optional.String**| Filter by branch_id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

