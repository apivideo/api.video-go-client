# \LiveApi

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](LiveApi.md#Delete) | **Delete** /live-streams/{liveStreamId} | Delete a live stream
[**DeleteThumbnail**](LiveApi.md#DeleteThumbnail) | **Delete** /live-streams/{liveStreamId}/thumbnail | Delete a thumbnail
[**List**](LiveApi.md#List) | **Get** /live-streams | List all live streams
[**Get**](LiveApi.md#Get) | **Get** /live-streams/{liveStreamId} | Show live stream
[**Update**](LiveApi.md#Update) | **Patch** /live-streams/{liveStreamId} | Update a live stream
[**Create**](LiveApi.md#Create) | **Post** /live-streams | Create live stream
[**UploadThumbnail**](LiveApi.md#UploadThumbnail) | **Post** /live-streams/{liveStreamId}/thumbnail | Upload a thumbnail



## Delete

> Delete(ctx, liveStreamId).Execute()

Delete a live stream

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
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique ID for the live stream that you want to remove.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveApi.Delete(context.Background(), liveStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveStreamId** | **string** | The unique ID for the live stream that you want to remove. | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteThumbnail

> LiveStream DeleteThumbnail(ctx, liveStreamId).Execute()

Delete a thumbnail



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
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique identifier for the live stream you want to delete. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveApi.DeleteThumbnail(context.Background(), liveStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveApi.DeleteThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteThumbnail`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `LiveApi.DeleteThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveStreamId** | **string** | The unique identifier for the live stream you want to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LiveStream**](live-stream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> LiveStreamListResponse List(ctx).StreamKey(streamKey).Name(name).SortBy(sortBy).SortOrder(sortOrder).CurrentPage(currentPage).PageSize(pageSize).Execute()

List all live streams



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
    streamKey := "30087931-229e-42cf-b5f9-e91bcc1f7332" // string | The unique stream key that allows you to stream videos. (optional)
    name := "My Video" // string | You can filter live streams by their name or a part of their name. (optional)
    sortBy := "createdAt" // string | Allowed: createdAt, publishedAt, name. createdAt - the time a livestream was created using the specified streamKey. publishedAt - the time a livestream was published using the specified streamKey. name - the name of the livestream. If you choose one of the time based options, the time is presented in ISO-8601 format. (optional)
    sortOrder := "desc" // string | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones. For title, it is 0-9 and A-Z ascending and Z-A, 9-0 descending. (optional)
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (optional) (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveApi.List(context.Background()).StreamKey(streamKey).Name(name).SortBy(sortBy).SortOrder(sortOrder).CurrentPage(currentPage).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: LiveStreamListResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamKey** | **string** | The unique stream key that allows you to stream videos. | 
 **name** | **string** | You can filter live streams by their name or a part of their name. | 
 **sortBy** | **string** | Allowed: createdAt, publishedAt, name. createdAt - the time a livestream was created using the specified streamKey. publishedAt - the time a livestream was published using the specified streamKey. name - the name of the livestream. If you choose one of the time based options, the time is presented in ISO-8601 format. | 
 **sortOrder** | **string** | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones. For title, it is 0-9 and A-Z ascending and Z-A, 9-0 descending. | 
 **currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
 **pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**LiveStreamListResponse**](live-stream-list-response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> LiveStream Get(ctx, liveStreamId).Execute()

Show live stream



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
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique ID for the live stream you want to watch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveApi.Get(context.Background(), liveStreamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `LiveApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveStreamId** | **string** | The unique ID for the live stream you want to watch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LiveStream**](live-stream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> LiveStream Update(ctx, liveStreamId).LiveStreamUpdatePayload(liveStreamUpdatePayload).Execute()

Update a live stream



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
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique ID for the live stream that you want to update information for such as player details, or whether you want the recording on or off.
    liveStreamUpdatePayload := *openapiclient.NewLiveStreamUpdatePayload() // LiveStreamUpdatePayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveApi.Update(context.Background(), liveStreamId).LiveStreamUpdatePayload(liveStreamUpdatePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `LiveApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveStreamId** | **string** | The unique ID for the live stream that you want to update information for such as player details, or whether you want the recording on or off. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveStreamUpdatePayload** | [**LiveStreamUpdatePayload**](LiveStreamUpdatePayload.md) |  | 

### Return type

[**LiveStream**](live-stream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> LiveStream Create(ctx).LiveStreamCreatePayload(liveStreamCreatePayload).Execute()

Create live stream



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
    liveStreamCreatePayload := *openapiclient.NewLiveStreamCreatePayload("My Live Stream Video") // LiveStreamCreatePayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveApi.Create(context.Background()).LiveStreamCreatePayload(liveStreamCreatePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `LiveApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **liveStreamCreatePayload** | [**LiveStreamCreatePayload**](LiveStreamCreatePayload.md) |  | 

### Return type

[**LiveStream**](live-stream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadThumbnail

> LiveStream UploadThumbnail(ctx, liveStreamId).File(file).Execute()

Upload a thumbnail



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
    liveStreamId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique ID for the live stream you want to upload.
    file := os.NewFile(1234, "some_file") // *os.File | The image to be added as a thumbnail.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveApi.UploadThumbnail(context.Background(), liveStreamId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveApi.UploadThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadThumbnail`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `LiveApi.UploadThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveStreamId** | **string** | The unique ID for the live stream you want to upload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The image to be added as a thumbnail. | 

### Return type

[**LiveStream**](live-stream.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

