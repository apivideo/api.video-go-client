# \AdvancedAuthentication

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authenticate**](AdvancedAuthentication.md#Authenticate) | **Post** /auth/api-key | Get Bearer Token
[**Refresh**](AdvancedAuthentication.md#Refresh) | **Post** /auth/refresh | Refresh Bearer Token



## Authenticate

> Authenticate(authenticatePayload AuthenticatePayload) (*AccessToken, error)

> AuthenticateWithContext(ctx context.Context, authenticatePayload AuthenticatePayload) (*AccessToken, error)


Get Bearer Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_KEY").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_KEY").Build()
        
    authenticatePayload := *apivideosdk.NewAuthenticatePayload("ApiKey_example") // AuthenticatePayload | 

    
    res, err := client.AdvancedAuthentication.Authenticate(authenticatePayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAuthentication.Authenticate``: %v\n", err)
    }
    // response from `Authenticate`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `AdvancedAuthentication.Authenticate`: %v\n", res)
}
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


Refresh Bearer Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_KEY").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_KEY").Build()
        
    refreshTokenPayload := *apivideosdk.NewRefreshTokenPayload("RefreshToken_example") // RefreshTokenPayload | 

    
    res, err := client.AdvancedAuthentication.Refresh(refreshTokenPayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAuthentication.Refresh``: %v\n", err)
    }
    // response from `Refresh`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `AdvancedAuthentication.Refresh`: %v\n", res)
}
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

