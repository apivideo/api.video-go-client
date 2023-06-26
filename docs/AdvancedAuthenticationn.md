# \AdvancedAuthenticationn

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Refresh**](AdvancedAuthenticationn.md#Refresh) | **Post** /auth/refresh | Refresh Bearer Token



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

    
    res, err := client.AdvancedAuthenticationn.Refresh(refreshTokenPayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdvancedAuthenticationn.Refresh``: %v\n", err)
    }
    // response from `Refresh`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `AdvancedAuthenticationn.Refresh`: %v\n", res)
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

