# \ChaptersApi

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](ChaptersApi.md#Delete) | **Delete** /videos/{videoId}/chapters/{language} | Delete a chapter
[**List**](ChaptersApi.md#List) | **Get** /videos/{videoId}/chapters | List video chapters
[**Get**](ChaptersApi.md#Get) | **Get** /videos/{videoId}/chapters/{language} | Show a chapter
[**Upload**](ChaptersApi.md#Upload) | **Post** /videos/{videoId}/chapters/{language} | Upload a chapter



## Delete

> Delete(ctx, videoId, language).Execute()

Delete a chapter

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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want to delete a chapter from. 
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChaptersApi.Delete(context.Background(), videoId, language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChaptersApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want to delete a chapter from.  | 
**language** | **string** | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


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


## List

> ChaptersListResponse List(ctx, videoId).CurrentPage(currentPage).PageSize(pageSize).Execute()

List video chapters



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want to retrieve a list of chapters for.
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (optional) (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChaptersApi.List(context.Background(), videoId).CurrentPage(currentPage).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChaptersApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ChaptersListResponse
    fmt.Fprintf(os.Stdout, "Response from `ChaptersApi.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want to retrieve a list of chapters for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
 **pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**ChaptersListResponse**](chapters-list-response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Chapter Get(ctx, videoId, language).Execute()

Show a chapter

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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want to show a chapter for.
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChaptersApi.Get(context.Background(), videoId, language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChaptersApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Chapter
    fmt.Fprintf(os.Stdout, "Response from `ChaptersApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want to show a chapter for. | 
**language** | **string** | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Chapter**](chapter.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upload

> Chapter Upload(ctx, videoId, language).File(file).Execute()

Upload a chapter



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want to upload a chapter for.
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.
    file := os.NewFile(1234, "some_file") // *os.File | The VTT file describing the chapters you want to upload.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChaptersApi.Upload(context.Background(), videoId, language).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChaptersApi.Upload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Upload`: Chapter
    fmt.Fprintf(os.Stdout, "Response from `ChaptersApi.Upload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want to upload a chapter for. | 
**language** | **string** | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The VTT file describing the chapters you want to upload. | 

### Return type

[**Chapter**](chapter.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

