# \Live

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](Live.md#Delete) | **Delete** /live-streams/{liveStreamId} | Delete a live stream
[**DeleteThumbnail**](Live.md#DeleteThumbnail) | **Delete** /live-streams/{liveStreamId}/thumbnail | Delete a thumbnail
[**List**](Live.md#List) | **Get** /live-streams | List all live streams
[**Get**](Live.md#Get) | **Get** /live-streams/{liveStreamId} | Show live stream
[**Update**](Live.md#Update) | **Patch** /live-streams/{liveStreamId} | Update a live stream
[**Create**](Live.md#Create) | **Post** /live-streams | Create live stream
[**UploadThumbnail**](Live.md#UploadThumbnail) | **Post** /live-streams/{liveStreamId}/thumbnail | Upload a thumbnail



## Delete

> Delete(liveStreamId string) (error)


Delete a live stream

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
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique ID for the live stream that you want to remove.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Live.Delete(liveStreamId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Live.Delete``: %v\n", err)
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

> DeleteThumbnail(liveStreamId string) (*LiveStream, error)


Delete a thumbnail



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
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique identifier for the live stream you want to delete. 

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Live.DeleteThumbnail(liveStreamId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Live.DeleteThumbnail``: %v\n", err)
    }
    // response from `DeleteThumbnail`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `Live.DeleteThumbnail`: %v\n", res)
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

> List(r LiveApiListRequest) (*LiveStreamListResponse, error)



List all live streams



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
    streamKey := "30087931-229e-42cf-b5f9-e91bcc1f7332" // string | The unique stream key that allows you to stream videos.
    name := "My Video" // string | You can filter live streams by their name or a part of their name.
    sortBy := "createdAt" // string | Allowed: createdAt, publishedAt, name. createdAt - the time a livestream was created using the specified streamKey. publishedAt - the time a livestream was published using the specified streamKey. name - the name of the livestream. If you choose one of the time based options, the time is presented in ISO-8601 format.
    sortOrder := "desc" // string | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones. For title, it is 0-9 and A-Z ascending and Z-A, 9-0 descending.
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Live.List()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Live.List``: %v\n", err)
    }
    // response from `List`: LiveStreamListResponse
    fmt.Fprintf(os.Stdout, "Response from `Live.List`: %v\n", res)
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

> Get(liveStreamId string) (*LiveStream, error)


Show live stream



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
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique ID for the live stream you want to watch.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Live.Get(liveStreamId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Live.Get``: %v\n", err)
    }
    // response from `Get`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `Live.Get`: %v\n", res)
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

> Update(liveStreamId string, liveStreamUpdatePayload LiveStreamUpdatePayload) (*LiveStream, error)


Update a live stream



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
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique ID for the live stream that you want to update information for such as player details, or whether you want the recording on or off.
    liveStreamUpdatePayload := *apivideosdk.NewLiveStreamUpdatePayload() // LiveStreamUpdatePayload | 

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Live.Update(liveStreamId, liveStreamUpdatePayload)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Live.Update``: %v\n", err)
    }
    // response from `Update`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `Live.Update`: %v\n", res)
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

> Create(liveStreamCreatePayload LiveStreamCreatePayload) (*LiveStream, error)


Create live stream



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
    liveStreamCreatePayload := *apivideosdk.NewLiveStreamCreatePayload("My Live Stream Video") // LiveStreamCreatePayload | 

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Live.Create(liveStreamCreatePayload)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Live.Create``: %v\n", err)
    }
    // response from `Create`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `Live.Create`: %v\n", res)
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

> UploadThumbnail(liveStreamId string, file *os.File) (*LiveStream, error)


Upload a thumbnail



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
    liveStreamId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique ID for the live stream you want to upload.
    file := os.NewFile(1234, "some_file") // *os.File | The image to be added as a thumbnail.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Live.UploadThumbnail(liveStreamId, file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Live.UploadThumbnail``: %v\n", err)
    }
    // response from `UploadThumbnail`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `Live.UploadThumbnail`: %v\n", res)
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

