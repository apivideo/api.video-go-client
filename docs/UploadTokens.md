# \UploadTokens

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteToken**](UploadTokens.md#DeleteToken) | **Delete** /upload-tokens/{uploadToken} | Delete an upload token
[**List**](UploadTokens.md#List) | **Get** /upload-tokens | List all active upload tokens.
[**GetToken**](UploadTokens.md#GetToken) | **Get** /upload-tokens/{uploadToken} | Show upload token
[**CreateToken**](UploadTokens.md#CreateToken) | **Post** /upload-tokens | Generate an upload token



## DeleteToken

> DeleteToken(uploadToken string) (error)

> DeleteTokenWithContext(ctx context.Context, uploadToken string) (error)


Delete an upload token



### Example
```go
//install the Go API client
//go get github.com/apivideo/api.video-go-client
package main

import (
    "context"
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    uploadToken := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the upload token you want to delete. Deleting a token will make it so the token can no longer be used for authentication.

    
    err := client.UploadTokens.DeleteToken(uploadToken)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadTokens.DeleteToken``: %v\
", err)
    }
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uploadToken** | **string** | The unique identifier for the upload token you want to delete. Deleting a token will make it so the token can no longer be used for authentication. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r UploadTokensApiListRequest) (*TokenListResponse, error)


> ListWithContext(ctx context.Context, r UploadTokensApiListRequest) (*TokenListResponse, error)



List all active upload tokens.



### Example
```go
//install the Go API client
//go get github.com/apivideo/api.video-go-client
package main
import (
    "context"
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    uploadToken := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the token you want information about.

    
    res, err := client.UploadTokens.GetToken(uploadToken)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadTokens.GetToken``: %v\
", err)
    }
    // response from `GetToken`: UploadToken
    fmt.Fprintf(os.Stdout, "Response from `UploadTokens.GetToken`: %v\
", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sortBy** | **string** | Allowed: createdAt, ttl. You can use these to sort by when a token was created, or how much longer the token will be active (ttl - time to live). Date and time is presented in ISO-8601 format. | 
**sortOrder** | **string** | Allowed: asc, desc. Ascending is 0-9 or A-Z. Descending is 9-0 or Z-A. | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**TokenListResponse**](TokenListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> GetToken(uploadToken string) (*UploadToken, error)

> GetTokenWithContext(ctx context.Context, uploadToken string) (*UploadToken, error)


Show upload token



### Example
```go
//install the Go API client
//go get github.com/apivideo/api.video-go-client
package main

import (
    "context"
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    uploadToken := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the token you want information about.

    
    res, err := client.UploadTokens.GetToken(uploadToken)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadTokens.GetToken``: %v\
", err)
    }
    // response from `GetToken`: UploadToken
    fmt.Fprintf(os.Stdout, "Response from `UploadTokens.GetToken`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**uploadToken** | **string** | The unique identifier for the token you want information about. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**UploadToken**](UploadToken.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> CreateToken(tokenCreationPayload TokenCreationPayload) (*UploadToken, error)

> CreateTokenWithContext(ctx context.Context, tokenCreationPayload TokenCreationPayload) (*UploadToken, error)


Generate an upload token



### Example
```go
//install the Go API client
//go get github.com/apivideo/api.video-go-client
package main

import (
    "context"
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    tokenCreationPayload := *apivideosdk.NewTokenCreationPayload() // TokenCreationPayload | 

    
    res, err := client.UploadTokens.CreateToken(tokenCreationPayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadTokens.CreateToken``: %v\
", err)
    }
    // response from `CreateToken`: UploadToken
    fmt.Fprintf(os.Stdout, "Response from `UploadTokens.CreateToken`: %v\
", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**tokenCreationPayload** | [**TokenCreationPayload**](TokenCreationPayload.md) |  | 

### Return type

[**UploadToken**](UploadToken.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

