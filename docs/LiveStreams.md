# \LiveStreams

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](LiveStreams.md#Create) | **Post** /live-streams | Create live stream
[**Get**](LiveStreams.md#Get) | **Get** /live-streams/{liveStreamId} | Retrieve live stream
[**Update**](LiveStreams.md#Update) | **Patch** /live-streams/{liveStreamId} | Update a live stream
[**Delete**](LiveStreams.md#Delete) | **Delete** /live-streams/{liveStreamId} | Delete a live stream
[**List**](LiveStreams.md#List) | **Get** /live-streams | List all live streams
[**UploadThumbnail**](LiveStreams.md#UploadThumbnail) | **Post** /live-streams/{liveStreamId}/thumbnail | Upload a thumbnail
[**DeleteThumbnail**](LiveStreams.md#DeleteThumbnail) | **Delete** /live-streams/{liveStreamId}/thumbnail | Delete a thumbnail
[**Complete**](LiveStreams.md#Complete) | **Put** /live-streams/{liveStreamId}/complete | Complete a live stream



## Create

> Create(liveStreamCreationPayload LiveStreamCreationPayload) (*LiveStream, error)

> CreateWithContext(ctx context.Context, liveStreamCreationPayload LiveStreamCreationPayload) (*LiveStream, error)


Create live stream



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
        
    liveStreamCreationPayload := *apivideosdk.NewLiveStreamCreationPayload("My Live Stream Video") // LiveStreamCreationPayload | 

    
    res, err := client.LiveStreams.Create(liveStreamCreationPayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreams.Create``: %v\n", err)
    }
    // response from `Create`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `LiveStreams.Create`: %v\n", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamCreationPayload** | [**LiveStreamCreationPayload**](LiveStreamCreationPayload.md) |  | 

### Return type

[**LiveStream**](LiveStream.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(liveStreamId string) (*LiveStream, error)

> GetWithContext(ctx context.Context, liveStreamId string) (*LiveStream, error)


Retrieve live stream



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
        
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique ID for the live stream you want to watch.

    
    res, err := client.LiveStreams.Get(liveStreamId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreams.Get``: %v\n", err)
    }
    // response from `Get`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `LiveStreams.Get`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamId** | **string** | The unique ID for the live stream you want to watch. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**LiveStream**](LiveStream.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(liveStreamId string, liveStreamUpdatePayload LiveStreamUpdatePayload) (*LiveStream, error)

> UpdateWithContext(ctx context.Context, liveStreamId string, liveStreamUpdatePayload LiveStreamUpdatePayload) (*LiveStream, error)


Update a live stream



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
        
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique ID for the live stream that you want to update information for such as player details.
    liveStreamUpdatePayload := *apivideosdk.NewLiveStreamUpdatePayload() // LiveStreamUpdatePayload | 

    
    res, err := client.LiveStreams.Update(liveStreamId, liveStreamUpdatePayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreams.Update``: %v\n", err)
    }
    // response from `Update`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `LiveStreams.Update`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamId** | **string** | The unique ID for the live stream that you want to update information for such as player details. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamUpdatePayload** | [**LiveStreamUpdatePayload**](LiveStreamUpdatePayload.md) |  | 

### Return type

[**LiveStream**](LiveStream.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(liveStreamId string) (error)

> DeleteWithContext(ctx context.Context, liveStreamId string) (error)


Delete a live stream



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
        
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique ID for the live stream that you want to remove.

    
    err := client.LiveStreams.Delete(liveStreamId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreams.Delete``: %v\n", err)
    }
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamId** | **string** | The unique ID for the live stream that you want to remove. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r LiveStreamsApiListRequest) (*LiveStreamListResponse, error)


> ListWithContext(ctx context.Context, r LiveStreamsApiListRequest) (*LiveStreamListResponse, error)



List all live streams



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
    req := apivideosdk.LiveStreamsApiListRequest{}
    
    req.StreamKey("dw-dew8-q6w9-k67w-1ws8") // string | The unique stream key that allows you to stream videos.
    req.Name("My Video") // string | You can filter live streams by their name or a part of their name.
    req.SortBy("createdAt") // string | Enables you to sort live stream results. Allowed attributes: `name`, `createdAt`, `updatedAt`. `name` - the name of the live stream. `createdAt` - the time a live stream was created. `updatedAt` - the time a live stream was last updated.  When using `createdAt` or `updatedAt`, the API sorts the results based on the ISO-8601 time format. 
    req.SortOrder("desc") // string | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones. For title, it is 0-9 and A-Z ascending and Z-A, 9-0 descending.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.LiveStreams.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreams.List``: %v\n", err)
    }
    // response from `List`: LiveStreamListResponse
    fmt.Fprintf(os.Stdout, "Response from `LiveStreams.List`: %v\n", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**streamKey** | **string** | The unique stream key that allows you to stream videos. | 
**name** | **string** | You can filter live streams by their name or a part of their name. | 
**sortBy** | **string** | Enables you to sort live stream results. Allowed attributes: &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;. &#x60;name&#x60; - the name of the live stream. &#x60;createdAt&#x60; - the time a live stream was created. &#x60;updatedAt&#x60; - the time a live stream was last updated.  When using &#x60;createdAt&#x60; or &#x60;updatedAt&#x60;, the API sorts the results based on the ISO-8601 time format.  | 
**sortOrder** | **string** | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones. For title, it is 0-9 and A-Z ascending and Z-A, 9-0 descending. | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**LiveStreamListResponse**](LiveStreamListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadThumbnail

> UploadThumbnailFile(liveStreamId string, file *os.File) (*LiveStream, error)
> UploadThumbnail(liveStreamId string, fileName string, fileReader io.Reader)
> UploadThumbnailFileWithContext(ctx context.Context, liveStreamId string, file *os.File) (*LiveStream, error)
> UploadThumbnailWithContext(ctx context.Context, liveStreamId string, fileName string, fileReader io.Reader)

Upload a thumbnail



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
        
    liveStreamId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique ID for the live stream you want to upload.
    file := os.NewFile(1234, "some_file") // *os.File | The image to be added as a thumbnail. The mime type should be image/jpeg, image/png or image/webp. The max allowed size is 8 MiB.

    
    res, err := client.LiveStreams.UploadThumbnailFile(liveStreamId, file)

    // you can also use a Reader instead of a File:
    // client.LiveStreams.UploadThumbnail(liveStreamId, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreams.UploadThumbnail``: %v\n", err)
    }
    // response from `UploadThumbnail`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `LiveStreams.UploadThumbnail`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamId** | **string** | The unique ID for the live stream you want to upload. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | The image to be added as a thumbnail. The mime type should be image/jpeg, image/png or image/webp. The max allowed size is 8 MiB. | 

### Return type

[**LiveStream**](LiveStream.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteThumbnail

> DeleteThumbnail(liveStreamId string) (*LiveStream, error)

> DeleteThumbnailWithContext(ctx context.Context, liveStreamId string) (*LiveStream, error)


Delete a thumbnail



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
        
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique identifier of the live stream whose thumbnail you want to delete.

    
    res, err := client.LiveStreams.DeleteThumbnail(liveStreamId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreams.DeleteThumbnail``: %v\n", err)
    }
    // response from `DeleteThumbnail`: LiveStream
    fmt.Fprintf(os.Stdout, "Response from `LiveStreams.DeleteThumbnail`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamId** | **string** | The unique identifier of the live stream whose thumbnail you want to delete. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**LiveStream**](LiveStream.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Complete

> Complete(liveStreamId string) (error)

> CompleteWithContext(ctx context.Context, liveStreamId string) (error)


Complete a live stream



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
        
    liveStreamId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique ID for the live stream you want to complete.

    
    err := client.LiveStreams.Complete(liveStreamId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveStreams.Complete``: %v\n", err)
    }
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamId** | **string** | The unique ID for the live stream you want to complete. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

