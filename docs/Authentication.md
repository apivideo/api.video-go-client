# \Authentication

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authenticate**](Authentication.md#Authenticate) | **Post** /auth/api-key | Authenticate
[**Refresh**](Authentication.md#Refresh) | **Post** /auth/refresh | Refresh token



## Authenticate

> Authenticate(authenticatePayload AuthenticatePayload) (*AccessToken, error)


Authenticate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/go-api-client"
)

func main() {
    authenticatePayload := *apivideosdk.NewAuthenticatePayload("ApiKey_example") // AuthenticatePayload | 

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Authentication.Authenticate(authenticatePayload)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Authentication.Authenticate``: %v\n", err)
    }
    // response from `Authenticate`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `Authentication.Authenticate`: %v\n", res)
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

> Refresh(refreshTokenPayload RefreshTokenPayload) (*AccessToken, error)


Refresh token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/go-api-client"
)

func main() {
    refreshTokenPayload := *apivideosdk.NewRefreshTokenPayload("RefreshToken_example") // RefreshTokenPayload | 

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Authentication.Refresh(refreshTokenPayload)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Authentication.Refresh``: %v\n", err)
    }
    // response from `Refresh`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `Authentication.Refresh`: %v\n", res)
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

