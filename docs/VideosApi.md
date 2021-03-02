# \VideosApi

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](VideosApi.md#Delete) | **Delete** /videos/{videoId} | Delete a video
[**Get**](VideosApi.md#Get) | **Get** /videos/{videoId} | Show a video
[**GetVideoStatus**](VideosApi.md#GetVideoStatus) | **Get** /videos/{videoId}/status | Show video status
[**List**](VideosApi.md#List) | **Get** /videos | List all videos
[**Update**](VideosApi.md#Update) | **Patch** /videos/{videoId} | Update a video
[**PickThumbnail**](VideosApi.md#PickThumbnail) | **Patch** /videos/{videoId}/thumbnail | Pick a thumbnail
[**Create**](VideosApi.md#Create) | **Post** /videos | Create a video
[**Upload**](VideosApi.md#Upload) | **Post** /videos/{videoId}/source | Upload a video
[**UploadThumbnail**](VideosApi.md#UploadThumbnail) | **Post** /videos/{videoId}/thumbnail | Upload a thumbnail



## Delete

> Delete(ctx, videoId).Execute()

Delete a video



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The video ID for the video you want to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosApi.Delete(context.Background(), videoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
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

> Video Get(ctx, videoId).Execute()

Show a video



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
    videoId := "videoId_example" // string | The unique identifier for the video you want details about.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosApi.Get(context.Background(), videoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Video
    fmt.Fprintf(os.Stdout, "Response from `VideosApi.Get`: %v\n", resp)
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

> Videostatus GetVideoStatus(ctx, videoId).Execute()

Show video status



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want the status for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosApi.GetVideoStatus(context.Background(), videoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosApi.GetVideoStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVideoStatus`: Videostatus
    fmt.Fprintf(os.Stdout, "Response from `VideosApi.GetVideoStatus`: %v\n", resp)
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

> VideosListResponse List(ctx).Title(title).Tags(tags).Metadata(metadata).Description(description).LiveStreamId(liveStreamId).SortBy(sortBy).SortOrder(sortOrder).CurrentPage(currentPage).PageSize(pageSize).Execute()

List all videos



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
    title := "My Video.mp4" // string | The title of a specific video you want to find. The search will match exactly to what term you provide and return any videos that contain the same term as part of their titles. (optional)
    tags := []string{"Inner_example"} // []string | A tag is a category you create and apply to videos. You can search for videos with particular tags by listing one or more here. Only videos that have all the tags you list will be returned. (optional)
    metadata := []string{"Inner_example"} // []string | Videos can be tagged with metadata tags in key:value pairs. You can search for videos with specific key value pairs using this parameter. (optional)
    description := "New Zealand" // string | If you described a video with a term or sentence, you can add it here to return videos containing this string. (optional)
    liveStreamId := "li400mYKSgQ6xs7taUeSaEKr" // string | If you know the ID for a live stream, you can retrieve the stream by adding the ID for it here. (optional)
    sortBy := "publishedAt" // string | Allowed: publishedAt, title. You can search by the time videos were published at, or by title. (optional)
    sortOrder := "asc" // string | Allowed: asc, desc. asc is ascending and sorts from A to Z. desc is descending and sorts from Z to A. (optional)
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (optional) (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosApi.List(context.Background()).Title(title).Tags(tags).Metadata(metadata).Description(description).LiveStreamId(liveStreamId).SortBy(sortBy).SortOrder(sortOrder).CurrentPage(currentPage).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: VideosListResponse
    fmt.Fprintf(os.Stdout, "Response from `VideosApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** | The title of a specific video you want to find. The search will match exactly to what term you provide and return any videos that contain the same term as part of their titles. | 
 **tags** | **[]string** | A tag is a category you create and apply to videos. You can search for videos with particular tags by listing one or more here. Only videos that have all the tags you list will be returned. | 
 **metadata** | **[]string** | Videos can be tagged with metadata tags in key:value pairs. You can search for videos with specific key value pairs using this parameter. | 
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

> Video Update(ctx, videoId).VideoUpdatePayload(videoUpdatePayload).Execute()

Update a video



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The video ID for the video you want to delete.
    videoUpdatePayload := *openapiclient.NewVideoUpdatePayload() // VideoUpdatePayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosApi.Update(context.Background(), videoId).VideoUpdatePayload(videoUpdatePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Video
    fmt.Fprintf(os.Stdout, "Response from `VideosApi.Update`: %v\n", resp)
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

> Video PickThumbnail(ctx, videoId).VideoThumbnailPickPayload(videoThumbnailPickPayload).Execute()

Pick a thumbnail



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | Unique identifier of the video you want to add a thumbnail to, where you use a section of your video as the thumbnail.
    videoThumbnailPickPayload := *openapiclient.NewVideoThumbnailPickPayload("Timecode_example") // VideoThumbnailPickPayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosApi.PickThumbnail(context.Background(), videoId).VideoThumbnailPickPayload(videoThumbnailPickPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosApi.PickThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PickThumbnail`: Video
    fmt.Fprintf(os.Stdout, "Response from `VideosApi.PickThumbnail`: %v\n", resp)
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

> Video Create(ctx).VideoCreatePayload(videoCreatePayload).Execute()

Create a video



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
    videoCreatePayload := *openapiclient.NewVideoCreatePayload("Maths video") // VideoCreatePayload | video to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosApi.Create(context.Background()).VideoCreatePayload(videoCreatePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Video
    fmt.Fprintf(os.Stdout, "Response from `VideosApi.Create`: %v\n", resp)
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

> Video Upload(ctx, videoId).File(file).Execute()

Upload a video



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | Enter the videoId you want to use to upload your video.
    file := os.NewFile(1234, "some_file") // *os.File | The path to the video you would like to upload. The path must be local. If you want to use a video from an online source, you must use the \\\"/videos\\\" endpoint and add the \\\"source\\\" parameter when you create a new video.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosApi.Upload(context.Background(), videoId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosApi.Upload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Upload`: Video
    fmt.Fprintf(os.Stdout, "Response from `VideosApi.Upload`: %v\n", resp)
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

> Video UploadThumbnail(ctx, videoId).File(file).Execute()

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
    videoId := "videoId_example" // string | Unique identifier of the chosen video 
    file := os.NewFile(1234, "some_file") // *os.File | The image to be added as a thumbnail.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VideosApi.UploadThumbnail(context.Background(), videoId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideosApi.UploadThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadThumbnail`: Video
    fmt.Fprintf(os.Stdout, "Response from `VideosApi.UploadThumbnail`: %v\n", resp)
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

