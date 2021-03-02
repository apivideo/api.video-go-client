# \CaptionsApi

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](CaptionsApi.md#Delete) | **Delete** /videos/{videoId}/captions/{language} | Delete a caption
[**List**](CaptionsApi.md#List) | **Get** /videos/{videoId}/captions | List video captions
[**Get**](CaptionsApi.md#Get) | **Get** /videos/{videoId}/captions/{language} | Show a caption
[**Update**](CaptionsApi.md#Update) | **Patch** /videos/{videoId}/captions/{language} | Update caption
[**Upload**](CaptionsApi.md#Upload) | **Post** /videos/{videoId}/captions/{language} | Upload a caption



## Delete

> Delete(ctx, videoId, language).Execute()

Delete a caption



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklgc" // string | The unique identifier for the video you want to delete a caption from.
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CaptionsApi.Delete(context.Background(), videoId, language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptionsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want to delete a caption from. | 
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

> CaptionsListResponse List(ctx, videoId).CurrentPage(currentPage).PageSize(pageSize).Execute()

List video captions



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want to retrieve a list of captions for.
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (optional) (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CaptionsApi.List(context.Background(), videoId).CurrentPage(currentPage).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptionsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: CaptionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `CaptionsApi.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want to retrieve a list of captions for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
 **pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**CaptionsListResponse**](captions-list-response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Subtitle Get(ctx, videoId, language).Execute()

Show a caption



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want captions for.
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CaptionsApi.Get(context.Background(), videoId, language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptionsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Subtitle
    fmt.Fprintf(os.Stdout, "Response from `CaptionsApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want captions for. | 
**language** | **string** | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Subtitle**](subtitle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Subtitle Update(ctx, videoId, language).CaptionsUpdatePayload(captionsUpdatePayload).Execute()

Update caption



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want to have automatic captions for. 
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.
    captionsUpdatePayload := *openapiclient.NewCaptionsUpdatePayload() // CaptionsUpdatePayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CaptionsApi.Update(context.Background(), videoId, language).CaptionsUpdatePayload(captionsUpdatePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptionsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Subtitle
    fmt.Fprintf(os.Stdout, "Response from `CaptionsApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want to have automatic captions for.  | 
**language** | **string** | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **captionsUpdatePayload** | [**CaptionsUpdatePayload**](CaptionsUpdatePayload.md) |  | 

### Return type

[**Subtitle**](subtitle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upload

> Subtitle Upload(ctx, videoId, language).File(file).Execute()

Upload a caption



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want to add a caption to.
    language := "en" // string | A valid BCP 47 language representation.
    file := os.NewFile(1234, "some_file") // *os.File | The video text track (VTT) you want to upload.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CaptionsApi.Upload(context.Background(), videoId, language).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptionsApi.Upload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Upload`: Subtitle
    fmt.Fprintf(os.Stdout, "Response from `CaptionsApi.Upload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want to add a caption to. | 
**language** | **string** | A valid BCP 47 language representation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The video text track (VTT) you want to upload. | 

### Return type

[**Subtitle**](subtitle.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

