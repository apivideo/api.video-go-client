# \Authentication

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authenticate**](Authentication.md#Authenticate) | **Post** /auth/api-key | Authenticate
[**Refresh**](Authentication.md#Refresh) | **Post** /auth/refresh | Refresh token



## Authenticate

> Authenticate(authenticatePayload AuthenticatePayload) (*AccessToken, error)

> AuthenticateWithContext(ctx context.Context, authenticatePayload AuthenticatePayload) (*AccessToken, error)


Authenticate



### Example
```go
//With the api.video API clients, authentication is taken care of with each client created.
// You get to skip this step!
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**authenticatePayload** | [**AuthenticatePayload**](AuthenticatePayload.md) |  | 

### Return type

[**AccessToken**](AccessToken.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Refresh

> Refresh(refreshTokenPayload RefreshTokenPayload) (*AccessToken, error)

> RefreshWithContext(ctx context.Context, refreshTokenPayload RefreshTokenPayload) (*AccessToken, error)


Refresh token



### Example
```go
//With the api.video API clients, authentication is taken care of with each client created.
// You get to skip this step!
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**refreshTokenPayload** | [**RefreshTokenPayload**](RefreshTokenPayload.md) |  | 

### Return type

[**AccessToken**](AccessToken.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

