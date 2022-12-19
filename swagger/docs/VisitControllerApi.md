# \VisitControllerApi

All URIs are relative to *https://localhost:8082*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST**](VisitControllerApi.md#CreateUsingPOST) | **Post** /visit | create
[**GetByUserIdUsingGET**](VisitControllerApi.md#GetByUserIdUsingGET) | **Get** /visit/get-all/{userId} | getByUserId


# **CreateUsingPOST**
> VisitDtoRes CreateUsingPOST(ctx, visitDto)
create

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **visitDto** | [**VisitDtoReq**](VisitDtoReq.md)| visitDto | 

### Return type

[**VisitDtoRes**](VisitDtoRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetByUserIdUsingGET**
> []VisitDtoRes GetByUserIdUsingGET(ctx, userId)
getByUserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int64**| userId | 

### Return type

[**[]VisitDtoRes**](VisitDtoRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

