# \DefaultApi

All URIs are relative to *https://auth.insolar.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateClient**](DefaultApi.md#ActivateClient) | **Put** /clients | Activate or deactivate client
[**AddClient**](DefaultApi.md#AddClient) | **Post** /clients | Add client
[**ClientArray**](DefaultApi.md#ClientArray) | **Get** /clients | Get clients
[**GetClient**](DefaultApi.md#GetClient) | **Get** /clients/{login} | Get client



## ActivateClient

> ActivateClient(ctx, schemasClientActivate)

Activate or deactivate client

Activate or deactivate access token issuance to the client registered in the auth-service.  To deactivate a client, specify in the request body:  * Client's unique login * State: `\"active\": false` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemasClientActivate** | [**SchemasClientActivate**](SchemasClientActivate.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddClient

> ResponsesNewClient AddClient(ctx, schemasClient)

Add client

Adds a new client to the auth-service.  To add a client, provide a unique login in the request body. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemasClient** | [**SchemasClient**](SchemasClient.md)|  | 

### Return type

[**ResponsesNewClient**](responses-new-client.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientArray

> ResponsesClientsYaml ClientArray(ctx, )

Get clients

Gets an array of client objects registered in the auth-service.  Each object contains:  * Unique login * State: active or inactive * Timestamp of the latest modification  To get information on a specific client, provide the unique client login in path: `/clients/{login}`. 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ResponsesClientsYaml**](responses-clients-yaml.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClient

> ResponsesClientYaml GetClient(ctx, login)

Get client

Gets the following information on a client registered in the authorization service:  * Unique login * State: active or inactive * Timestamp of the latest modification 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string**| Client&#39;s login—unique identificator of the client to get information on. | 

### Return type

[**ResponsesClientYaml**](responses-client-yaml.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

