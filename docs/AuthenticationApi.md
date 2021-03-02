# \AuthenticationApi

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authenticate**](AuthenticationApi.md#Authenticate) | **Post** /auth/api-key | Authenticate
[**Refresh**](AuthenticationApi.md#Refresh) | **Post** /auth/refresh | Refresh token



## Authenticate

> AccessToken Authenticate(ctx).AuthenticatePayload(authenticatePayload).Execute()

Authenticate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authenticatePayload := *openapiclient.NewAuthenticatePayload("ApiKey_example") // AuthenticatePayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.Authenticate(context.Background()).AuthenticatePayload(authenticatePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.Authenticate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Authenticate`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.Authenticate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticatePayload** | [**AuthenticatePayload**](AuthenticatePayload.md) |  | 

### Return type

[**AccessToken**](access-token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Refresh

> AccessToken Refresh(ctx).RefreshTokenPayload(refreshTokenPayload).Execute()

Refresh token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    refreshTokenPayload := *openapiclient.NewRefreshTokenPayload("RefreshToken_example") // RefreshTokenPayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.Refresh(context.Background()).RefreshTokenPayload(refreshTokenPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.Refresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Refresh`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.Refresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshTokenPayload** | [**RefreshTokenPayload**](RefreshTokenPayload.md) |  | 

### Return type

[**AccessToken**](access-token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

