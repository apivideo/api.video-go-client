# \VideosDelegatedUploadApi

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteToken**](VideosDelegatedUploadApi.md#DeleteToken) | **Delete** /upload-tokens/{uploadToken} | Delete an upload token
[**ListTokens**](VideosDelegatedUploadApi.md#ListTokens) | **Get** /upload-tokens | List all active upload tokens.
[**GetToken**](VideosDelegatedUploadApi.md#GetToken) | **Get** /upload-tokens/{uploadToken} | Show upload token
[**Upload**](VideosDelegatedUploadApi.md#Upload) | **Post** /upload | Upload with an upload token
[**CreateToken**](VideosDelegatedUploadApi.md#CreateToken) | **Post** /upload-tokens | Generate an upload token



## DeleteToken

> DeleteToken(ctx, uploadToken).Execute()

Delete an upload token



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
    uploadToken := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the upload token you want to delete. Deleting a token will make it so the token can no longer be used for authentication.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosDelegatedUploadApi.DeleteToken(context.Background(), uploadToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosDelegatedUploadApi.DeleteToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
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

> TokenListResponse ListTokens(ctx).SortBy(sortBy).SortOrder(sortOrder).CurrentPage(currentPage).PageSize(pageSize).Execute()

List all active upload tokens.



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
    sortBy := "ttl" // string | Allowed: createdAt, ttl. You can use these to sort by when a token was created, or how much longer the token will be active (ttl - time to live). Date and time is presented in ISO-8601 format. (optional)
    sortOrder := "asc" // string | Allowed: asc, desc. Ascending is 0-9 or A-Z. Descending is 9-0 or Z-A. (optional)
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (optional) (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosDelegatedUploadApi.ListTokens(context.Background()).SortBy(sortBy).SortOrder(sortOrder).CurrentPage(currentPage).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosDelegatedUploadApi.ListTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokens`: TokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `VideosDelegatedUploadApi.ListTokens`: %v\n", resp)
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

> UploadToken GetToken(ctx, uploadToken).Execute()

Show upload token



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
    uploadToken := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the token you want information about.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosDelegatedUploadApi.GetToken(context.Background(), uploadToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosDelegatedUploadApi.GetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetToken`: UploadToken
    fmt.Fprintf(os.Stdout, "Response from `VideosDelegatedUploadApi.GetToken`: %v\n", resp)
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

> Video Upload(ctx).Token(token).File(file).Execute()

Upload with an upload token



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
    token := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the token you want to use to upload a video.
    file := os.NewFile(1234, "some_file") // *os.File | The path to the video you want to upload.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosDelegatedUploadApi.Upload(context.Background()).Token(token).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosDelegatedUploadApi.Upload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Upload`: Video
    fmt.Fprintf(os.Stdout, "Response from `VideosDelegatedUploadApi.Upload`: %v\n", resp)
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

> UploadToken CreateToken(ctx).TokenCreatePayload(tokenCreatePayload).Execute()

Generate an upload token



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
    tokenCreatePayload := *openapiclient.NewTokenCreatePayload() // TokenCreatePayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosDelegatedUploadApi.CreateToken(context.Background()).TokenCreatePayload(tokenCreatePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosDelegatedUploadApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: UploadToken
    fmt.Fprintf(os.Stdout, "Response from `VideosDelegatedUploadApi.CreateToken`: %v\n", resp)
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

