# \DefaultApi

All URIs are relative to *https://auth.insolar.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetPassword**](DefaultApi.md#SetPassword) | **Post** /auth/set-password | Set password
[**Token**](DefaultApi.md#Token) | **Get** /auth/token | Get token



## SetPassword

> SetPassword(ctx, code, parametersSetPassword)

Set password

Sets the client's password for basic authentication.  The Observer node uses basic authentication to get a JSON Web Token (JWT) via the [token request](#operation/token). The JWT is required to authenticate every request from the Observer node to the Insolar network. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**| Unique registration code. Lives for 24 hours, expires upon first use, and allows to set the password for basic authentication. | 
**parametersSetPassword** | [**ParametersSetPassword**](ParametersSetPassword.md)|  | 

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


## Token

> ResponsesToken Token(ctx, authorizationBasic)

Get token

Gets an access token—JSON Web Token (JWT).  This request uses basic authentication through the Secure Sockets Layer (SSL).  The JWT is required to authenticate every request from the Observer node to the Insolar network.  JWT lives for 15 minutes. To continue accessing the Insolar network, the Observer node gets a new JWT upon the expiration of the old one. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationBasic** | **string**| HTTP basic authorization header with the Base64 encoding of login and password joined by a single colon &#x60;:&#x60;. | 

### Return type

[**ResponsesToken**](responses-token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

