# \LiveStreams

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](LiveStreams.md#Delete) | **Delete** /live-streams/{liveStreamId} | Delete a live stream
[**DeleteThumbnail**](LiveStreams.md#DeleteThumbnail) | **Delete** /live-streams/{liveStreamId}/thumbnail | Delete a thumbnail
[**List**](LiveStreams.md#List) | **Get** /live-streams | List all live streams
[**Get**](LiveStreams.md#Get) | **Get** /live-streams/{liveStreamId} | Show live stream
[**Update**](LiveStreams.md#Update) | **Patch** /live-streams/{liveStreamId} | Update a live stream
[**Create**](LiveStreams.md#Create) | **Post** /live-streams | Create live stream
[**UploadThumbnail**](LiveStreams.md#UploadThumbnail) | **Post** /live-streams/{liveStreamId}/thumbnail | Upload a thumbnail



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
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
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
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique identifier for the live stream you want to delete. 

    
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
**liveStreamId** | **string** | The unique identifier for the live stream you want to delete.  | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**LiveStream**](live-stream.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r LiveStreamsApiListRequest) (*LiveStreamListResponse, error)



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
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
    req := apivideosdk.LiveStreamsApiListRequest{}
    
    req.StreamKey("30087931-229e-42cf-b5f9-e91bcc1f7332") // string | The unique stream key that allows you to stream videos.
    req.Name("My Video") // string | You can filter live streams by their name or a part of their name.
    req.SortBy("createdAt") // string | Allowed: createdAt, publishedAt, name. createdAt - the time a livestream was created using the specified streamKey. publishedAt - the time a livestream was published using the specified streamKey. name - the name of the livestream. If you choose one of the time based options, the time is presented in ISO-8601 format.
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
**sortBy** | **string** | Allowed: createdAt, publishedAt, name. createdAt - the time a livestream was created using the specified streamKey. publishedAt - the time a livestream was published using the specified streamKey. name - the name of the livestream. If you choose one of the time based options, the time is presented in ISO-8601 format. | 
**sortOrder** | **string** | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones. For title, it is 0-9 and A-Z ascending and Z-A, 9-0 descending. | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**LiveStreamListResponse**](live-stream-list-response.md)

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
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
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

[**LiveStream**](live-stream.md)

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
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | The unique ID for the live stream that you want to update information for such as player details, or whether you want the recording on or off.
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
**liveStreamId** | **string** | The unique ID for the live stream that you want to update information for such as player details, or whether you want the recording on or off. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamUpdatePayload** | [**LiveStreamUpdatePayload**](LiveStreamUpdatePayload.md) |  | 

### Return type

[**LiveStream**](live-stream.md)

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
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    liveStreamCreatePayload := *apivideosdk.NewLiveStreamCreatePayload("My Live Stream Video") // LiveStreamCreatePayload | 

    
    res, err := client.LiveStreams.Create(liveStreamCreatePayload)

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
**liveStreamCreatePayload** | [**LiveStreamCreatePayload**](LiveStreamCreatePayload.md) |  | 

### Return type

[**LiveStream**](live-stream.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadThumbnail

> UploadThumbnailFile(liveStreamId string, file *os.File) (*LiveStream, error)

> UploadThumbnail(liveStreamId string, fileName string, fileReader io.Reader)

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
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    liveStreamId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique ID for the live stream you want to upload.
    file := os.NewFile(1234, "some_file") // *os.File | The image to be added as a thumbnail.

    
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
**file** | ***os.File** | The image to be added as a thumbnail. | 

### Return type

[**LiveStream**](live-stream.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

