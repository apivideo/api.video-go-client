# \Videos

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](Videos.md#Delete) | **Delete** /videos/{videoId} | Delete a video
[**Get**](Videos.md#Get) | **Get** /videos/{videoId} | Show a video
[**GetVideoStatus**](Videos.md#GetVideoStatus) | **Get** /videos/{videoId}/status | Show video status
[**List**](Videos.md#List) | **Get** /videos | List all videos
[**Update**](Videos.md#Update) | **Patch** /videos/{videoId} | Update a video
[**PickThumbnail**](Videos.md#PickThumbnail) | **Patch** /videos/{videoId}/thumbnail | Pick a thumbnail
[**Create**](Videos.md#Create) | **Post** /videos | Create a video
[**Upload**](Videos.md#Upload) | **Post** /videos/{videoId}/source | Upload a video
[**UploadThumbnail**](Videos.md#UploadThumbnail) | **Post** /videos/{videoId}/thumbnail | Upload a thumbnail



## Delete

> Delete(videoId string) (error)


Delete a video



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The video ID for the video you want to delete.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Videos.Delete(videoId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.Delete``: %v\n", err)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The video ID for the video you want to delete. | 

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


## Get

> Get(videoId string) (*Video, error)


Show a video



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
    videoId := "videoId_example" // string | The unique identifier for the video you want details about.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Videos.Get(videoId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.Get``: %v\n", err)
    }
    // response from `Get`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.Get`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want details about. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Video**](video.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoStatus

> GetVideoStatus(videoId string) (*Videostatus, error)


Show video status



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want the status for.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Videos.GetVideoStatus(videoId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.GetVideoStatus``: %v\n", err)
    }
    // response from `GetVideoStatus`: Videostatus
    fmt.Fprintf(os.Stdout, "Response from `Videos.GetVideoStatus`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want the status for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Videostatus**](videostatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r VideosApiListRequest) (*VideosListResponse, error)



List all videos



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
    title := "My Video.mp4" // string | The title of a specific video you want to find. The search will match exactly to what term you provide and return any videos that contain the same term as part of their titles.
    tags := []string{"Inner_example"} // []string | A tag is a category you create and apply to videos. You can search for videos with particular tags by listing one or more here. Only videos that have all the tags you list will be returned.
    metadata := []apivideosdk.Metadata{*apivideosdk.NewMetadata()} // []Metadata | Videos can be tagged with metadata tags in key:value pairs. You can search for videos with specific key value pairs using this parameter.
    description := "New Zealand" // string | If you described a video with a term or sentence, you can add it here to return videos containing this string.
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | If you know the ID for a live stream, you can retrieve the stream by adding the ID for it here.
    sortBy := "publishedAt" // string | Allowed: publishedAt, title. You can search by the time videos were published at, or by title.
    sortOrder := "asc" // string | Allowed: asc, desc. asc is ascending and sorts from A to Z. desc is descending and sorts from Z to A.
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Videos.List()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.List``: %v\n", err)
    }
    // response from `List`: VideosListResponse
    fmt.Fprintf(os.Stdout, "Response from `Videos.List`: %v\n", res)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** | The title of a specific video you want to find. The search will match exactly to what term you provide and return any videos that contain the same term as part of their titles. | 
 **tags** | **[]string** | A tag is a category you create and apply to videos. You can search for videos with particular tags by listing one or more here. Only videos that have all the tags you list will be returned. | 
 **metadata** | [**[]Metadata**](metadata.md) | Videos can be tagged with metadata tags in key:value pairs. You can search for videos with specific key value pairs using this parameter. | 
 **description** | **string** | If you described a video with a term or sentence, you can add it here to return videos containing this string. | 
 **liveStreamId** | **string** | If you know the ID for a live stream, you can retrieve the stream by adding the ID for it here. | 
 **sortBy** | **string** | Allowed: publishedAt, title. You can search by the time videos were published at, or by title. | 
 **sortOrder** | **string** | Allowed: asc, desc. asc is ascending and sorts from A to Z. desc is descending and sorts from Z to A. | 
 **currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
 **pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**VideosListResponse**](videos-list-response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(videoId string, videoUpdatePayload VideoUpdatePayload) (*Video, error)


Update a video



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The video ID for the video you want to delete.
    videoUpdatePayload := *apivideosdk.NewVideoUpdatePayload() // VideoUpdatePayload | 

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Videos.Update(videoId, videoUpdatePayload)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.Update``: %v\n", err)
    }
    // response from `Update`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.Update`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The video ID for the video you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **videoUpdatePayload** | [**VideoUpdatePayload**](VideoUpdatePayload.md) |  | 

### Return type

[**Video**](video.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PickThumbnail

> PickThumbnail(videoId string, videoThumbnailPickPayload VideoThumbnailPickPayload) (*Video, error)


Pick a thumbnail



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | Unique identifier of the video you want to add a thumbnail to, where you use a section of your video as the thumbnail.
    videoThumbnailPickPayload := *apivideosdk.NewVideoThumbnailPickPayload("Timecode_example") // VideoThumbnailPickPayload | 

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Videos.PickThumbnail(videoId, videoThumbnailPickPayload)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.PickThumbnail``: %v\n", err)
    }
    // response from `PickThumbnail`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.PickThumbnail`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Unique identifier of the video you want to add a thumbnail to, where you use a section of your video as the thumbnail. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPickThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **videoThumbnailPickPayload** | [**VideoThumbnailPickPayload**](VideoThumbnailPickPayload.md) |  | 

### Return type

[**Video**](video.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Create(videoCreatePayload VideoCreatePayload) (*Video, error)


Create a video



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
    videoCreatePayload := *apivideosdk.NewVideoCreatePayload("Maths video") // VideoCreatePayload | video to create

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Videos.Create(videoCreatePayload)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.Create``: %v\n", err)
    }
    // response from `Create`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.Create`: %v\n", res)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoCreatePayload** | [**VideoCreatePayload**](VideoCreatePayload.md) | video to create | 

### Return type

[**Video**](video.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upload

> Upload(videoId string, file *os.File) (*Video, error)


Upload a video



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | Enter the videoId you want to use to upload your video.
    file := os.NewFile(1234, "some_file") // *os.File | The path to the video you would like to upload. The path must be local. If you want to use a video from an online source, you must use the \\\"/videos\\\" endpoint and add the \\\"source\\\" parameter when you create a new video.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Videos.Upload(videoId, file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.Upload``: %v\n", err)
    }
    // response from `Upload`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.Upload`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Enter the videoId you want to use to upload your video. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The path to the video you would like to upload. The path must be local. If you want to use a video from an online source, you must use the \\\&quot;/videos\\\&quot; endpoint and add the \\\&quot;source\\\&quot; parameter when you create a new video. | 

### Return type

[**Video**](video.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadThumbnail

> UploadThumbnail(videoId string, file *os.File) (*Video, error)


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
    videoId := "videoId_example" // string | Unique identifier of the chosen video 
    file := os.NewFile(1234, "some_file") // *os.File | The image to be added as a thumbnail.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Videos.UploadThumbnail(videoId, file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.UploadThumbnail``: %v\n", err)
    }
    // response from `UploadThumbnail`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.UploadThumbnail`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Unique identifier of the chosen video  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The image to be added as a thumbnail. | 

### Return type

[**Video**](video.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

