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
package main

import (
    "context"
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/go-api-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    authenticatePayload := *apivideosdk.NewAuthenticatePayload("ApiKey_example") // AuthenticatePayload | 

    
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



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**authenticatePayload** | [**AuthenticatePayload**](AuthenticatePayload.md) |  | 

### Return type

[**AccessToken**](access-token.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Refresh

> Refresh(refreshTokenPayload RefreshTokenPayload) (*AccessToken, error)

> RefreshWithContext(ctx context.Context, refreshTokenPayload RefreshTokenPayload) (*AccessToken, error)


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
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    refreshTokenPayload := *apivideosdk.NewRefreshTokenPayload("RefreshToken_example") // RefreshTokenPayload | 

    
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



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**refreshTokenPayload** | [**RefreshTokenPayload**](RefreshTokenPayload.md) |  | 

### Return type

[**AccessToken**](access-token.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

