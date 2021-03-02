# \VideosDelegatedUpload

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteToken**](VideosDelegatedUpload.md#DeleteToken) | **Delete** /upload-tokens/{uploadToken} | Delete an upload token
[**ListTokens**](VideosDelegatedUpload.md#ListTokens) | **Get** /upload-tokens | List all active upload tokens.
[**GetToken**](VideosDelegatedUpload.md#GetToken) | **Get** /upload-tokens/{uploadToken} | Show upload token
[**Upload**](VideosDelegatedUpload.md#Upload) | **Post** /upload | Upload with an upload token
[**CreateToken**](VideosDelegatedUpload.md#CreateToken) | **Post** /upload-tokens | Generate an upload token



## DeleteToken

> DeleteToken(uploadToken string) (error)


Delete an upload token



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
    uploadToken := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the upload token you want to delete. Deleting a token will make it so the token can no longer be used for authentication.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.VideosDelegatedUpload.DeleteToken(uploadToken)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosDelegatedUpload.DeleteToken``: %v\n", err)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadToken** | **string** | The unique identifier for the upload token you want to delete. Deleting a token will make it so the token can no longer be used for authentication. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokens

> ListTokens(r VideosDelegatedUploadApiListTokensRequest) (*TokenListResponse, error)



List all active upload tokens.



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
    sortBy := "ttl" // string | Allowed: createdAt, ttl. You can use these to sort by when a token was created, or how much longer the token will be active (ttl - time to live). Date and time is presented in ISO-8601 format.
    sortOrder := "asc" // string | Allowed: asc, desc. Ascending is 0-9 or A-Z. Descending is 9-0 or Z-A.
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.VideosDelegatedUpload.ListTokens()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosDelegatedUpload.ListTokens``: %v\n", err)
    }
    // response from `ListTokens`: TokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `VideosDelegatedUpload.ListTokens`: %v\n", res)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Allowed: createdAt, ttl. You can use these to sort by when a token was created, or how much longer the token will be active (ttl - time to live). Date and time is presented in ISO-8601 format. | 
 **sortOrder** | **string** | Allowed: asc, desc. Ascending is 0-9 or A-Z. Descending is 9-0 or Z-A. | 
 **currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
 **pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**TokenListResponse**](token-list-response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> GetToken(uploadToken string) (*UploadToken, error)


Show upload token



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
    uploadToken := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the token you want information about.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.VideosDelegatedUpload.GetToken(uploadToken)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosDelegatedUpload.GetToken``: %v\n", err)
    }
    // response from `GetToken`: UploadToken
    fmt.Fprintf(os.Stdout, "Response from `VideosDelegatedUpload.GetToken`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadToken** | **string** | The unique identifier for the token you want information about. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UploadToken**](upload-token.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upload

> Upload(token string, file *os.File) (*Video, error)


Upload with an upload token



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
    token := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the token you want to use to upload a video.
    file := os.NewFile(1234, "some_file") // *os.File | The path to the video you want to upload.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.VideosDelegatedUpload.Upload(token, file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosDelegatedUpload.Upload``: %v\n", err)
    }
    // response from `Upload`: Video
    fmt.Fprintf(os.Stdout, "Response from `VideosDelegatedUpload.Upload`: %v\n", res)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The unique identifier for the token you want to use to upload a video. | 
 **file** | ***os.File** | The path to the video you want to upload. | 

### Return type

[**Video**](video.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> CreateToken(tokenCreatePayload TokenCreatePayload) (*UploadToken, error)


Generate an upload token



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
    tokenCreatePayload := *apivideosdk.NewTokenCreatePayload() // TokenCreatePayload | 

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.VideosDelegatedUpload.CreateToken(tokenCreatePayload)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosDelegatedUpload.CreateToken``: %v\n", err)
    }
    // response from `CreateToken`: UploadToken
    fmt.Fprintf(os.Stdout, "Response from `VideosDelegatedUpload.CreateToken`: %v\n", res)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenCreatePayload** | [**TokenCreatePayload**](TokenCreatePayload.md) |  | 

### Return type

[**UploadToken**](upload-token.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

