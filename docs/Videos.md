# \Videos

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Videos.md#Create) | **Post** /videos | Create a video object
[**Upload**](Videos.md#Upload) | **Post** /videos/{videoId}/source | Upload a video
[**UploadWithUploadToken**](Videos.md#UploadWithUploadToken) | **Post** /upload | Upload with an delegated upload token
[**Get**](Videos.md#Get) | **Get** /videos/{videoId} | Retrieve a video object
[**Update**](Videos.md#Update) | **Patch** /videos/{videoId} | Update a video object
[**Delete**](Videos.md#Delete) | **Delete** /videos/{videoId} | Delete a video object
[**List**](Videos.md#List) | **Get** /videos | List all video objects
[**UploadThumbnail**](Videos.md#UploadThumbnail) | **Post** /videos/{videoId}/thumbnail | Upload a thumbnail
[**PickThumbnail**](Videos.md#PickThumbnail) | **Patch** /videos/{videoId}/thumbnail | Set a thumbnail
[**GetStatus**](Videos.md#GetStatus) | **Get** /videos/{videoId}/status | Retrieve video status and details



## Create

> Create(videoCreationPayload VideoCreationPayload) (*Video, error)

> CreateWithContext(ctx context.Context, videoCreationPayload VideoCreationPayload) (*Video, error)


Create a video object



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
    // client := apivideosdk.SandboxClientBuilder("YOUR_SANDBOX_API_KEY").Build()
        
    videoCreationPayload := *apivideosdk.NewVideoCreationPayload("Maths video") // VideoCreationPayload | video to create

    
    res, err := client.Videos.Create(videoCreationPayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.Create``: %v\
", err)
    }
    // response from `Create`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.Create`: %v\
", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoCreationPayload** | [**VideoCreationPayload**](VideoCreationPayload.md) | video to create | 

### Return type

[**Video**](Video.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upload

> UploadFile(videoId string, file *os.File) (*Video, error)
> Upload(videoId string, fileName string, fileReader io.Reader, fileSize int64)
> UploadFileWithContext(ctx context.Context, videoId string, file *os.File) (*Video, error)
> UploadWithContext(ctx context.Context, videoId string, fileName string, fileReader io.Reader, fileSize int64)

Upload a video



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
    // client := apivideosdk.SandboxClientBuilder("YOUR_SANDBOX_API_KEY").Build()

      videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" 
    // string | Enter the videoId you want to use to upload your video.
      file := os.NewFile(1234, "some_file") 
    // *os.File | The path to the video you would like to upload. The path must be local. If you want to use a video from an online source, you must use the "/videos" endpoint and add the "source" parameter when you create a new video.
    
    
      res, err := client.Videos.UploadFile(videoId, file)
    
      // you can also use a Reader instead of a File:
      // client.Videos.Upload(videoId, fileName, fileReader, fileSize)
    
      if err != nil {
          fmt.Fprintf(os.Stderr, "Error when calling `Videos.Upload``: %v\
", err)
      }
      // response from `Upload`: Video
      fmt.Fprintf(os.Stdout, "Response from `Videos.Upload`: %v\
", res)
    }
      }
```
### Progressive uploads

Progressive uploads make it possible to upload a video source "progressively," i.e., before knowing the total size of the video. This is done by sending chunks of the video source file sequentially.
The last chunk is sent by calling a different method, so api.video knows that it is time to reassemble the different chunks that were received.


```go

videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | Enter the videoId you want to use to upload your video.

part1, err := os.Open("10m.mp4.part.a")
part2, err := os.Open("10m.mp4.part.b")
part3, err := os.Open("10m.mp4.part.c")

stream = client.Videos.CreateUploadStream(videoId)
_, err = stream.UploadPartFile(part1)
_, err = stream.UploadPartFile(part2)
res, err := stream.UploadLastPartFile(part3)

err = part1.Close()
err = part2.Close()
err = part3.Close()

```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | Enter the videoId you want to use to upload your video. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | The path to the video you would like to upload. The path must be local. If you want to use a video from an online source, you must use the \\\&quot;/videos\\\&quot; endpoint and add the \\\&quot;source\\\&quot; parameter when you create a new video. | 

### Return type

[**Video**](Video.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadWithUploadToken

> UploadWithUploadTokenFile(token string, file *os.File) (*Video, error)
> UploadWithUploadToken(token string, fileName string, fileReader io.Reader, fileSize int64)
> UploadWithUploadTokenFileWithContext(ctx context.Context, token string, file *os.File) (*Video, error)
> UploadWithUploadTokenWithContext(ctx context.Context, token string, fileName string, fileReader io.Reader, fileSize int64)

Upload with an delegated upload token



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
        
    token := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the token you want to use to upload a video.
    file := os.NewFile(1234, "some_file") // *os.File | The path to the video you want to upload.

    
    res, err := client.Videos.UploadWithUploadTokenFile(token, file)

    // you can also use a Reader instead of a File:
    // client.Videos.UploadWithUploadToken(token, fileName, fileReader, fileSize)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.UploadWithUploadToken``: %v\n", err)
    }
    // response from `UploadWithUploadToken`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.UploadWithUploadToken`: %v\n", res)
}
```
### Progressive uploads

Progressive uploads make it possible to upload a video source "progressively," i.e., before knowing the total size of the video. This is done by sending chunks of the video source file sequentially.
The last chunk is sent by calling a different method, so api.video knows that it is time to reassemble the different chunks that were received.


```go

token := "to1tcmSFHeYY5KzyhOqVKMKb" // string | The unique identifier for the token you want to use to upload a video.

part1, err := os.Open("10m.mp4.part.a")
part2, err := os.Open("10m.mp4.part.b")
part3, err := os.Open("10m.mp4.part.c")

stream = client.Videos.CreateUploadWithUploadTokenStream(token, nil)
// or, if you want to upload to an existing video:
// stream = client.Videos.CreateUploadWithUploadTokenStream(token, videoId)
_, err = stream.UploadPartFile(part1)
_, err = stream.UploadPartFile(part2)
res, err := stream.UploadLastPartFile(part3)

err = part1.Close()
err = part2.Close()
err = part3.Close()

```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**token** | **string** | The unique identifier for the token you want to use to upload a video. | 
**file** | ***os.File** | The path to the video you want to upload. | 

### Return type

[**Video**](Video.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(videoId string) (*Video, error)

> GetWithContext(ctx context.Context, videoId string) (*Video, error)


Retrieve a video object



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
    // client := apivideosdk.SandboxClientBuilder("YOUR_SANDBOX_API_KEY").Build()
        
    videoId := "videoId_example" // string | The unique identifier for the video you want details about.

    
    res, err := client.Videos.Get(videoId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.Get``: %v\
", err)
    }
    // response from `Get`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.Get`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want details about. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Video**](Video.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(videoId string, videoUpdatePayload VideoUpdatePayload) (*Video, error)

> UpdateWithContext(ctx context.Context, videoId string, videoUpdatePayload VideoUpdatePayload) (*Video, error)


Update a video object



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
    // client := apivideosdk.SandboxClientBuilder("YOUR_SANDBOX_API_KEY").Build()
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The video ID for the video you want to delete.
    videoUpdatePayload := *apivideosdk.NewVideoUpdatePayload() // VideoUpdatePayload | 

    
    res, err := client.Videos.Update(videoId, videoUpdatePayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.Update``: %v\
", err)
    }
    // response from `Update`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.Update`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The video ID for the video you want to update. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoUpdatePayload** | [**VideoUpdatePayload**](VideoUpdatePayload.md) |  | 

### Return type

[**Video**](Video.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(videoId string) (error)

> DeleteWithContext(ctx context.Context, videoId string) (error)


Delete a video object



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
    // client := apivideosdk.SandboxClientBuilder("YOUR_SANDBOX_API_KEY").Build()
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The video ID for the video you want to delete.
    err := client.Videos.Delete(videoId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.Delete``: %v\
", err)
    }
}  
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The video ID for the video you want to delete. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r VideosApiListRequest) (*VideosListResponse, error)


> ListWithContext(ctx context.Context, r VideosApiListRequest) (*VideosListResponse, error)



List all video objects



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
    req := apivideosdk.VideosApiListRequest{}
    
    req.Title("My Video.mp4") // string | The title of a specific video you want to find. The search will match exactly to what term you provide and return any videos that contain the same term as part of their titles.
    req.Tags([]string{"Inner_example"}) // []string | A tag is a category you create and apply to videos. You can search for videos with particular tags by listing one or more here. Only videos that have all the tags you list will be returned.
    req.Metadata(map[string]string{"key": "TODO"}) // map[string]string | Videos can be tagged with metadata tags in key:value pairs. You can search for videos with specific key value pairs using this parameter. [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata) allows you to define a key that allows any value pair.
    req.Description("New Zealand") // string | Retrieve video objects by `description`.
    req.LiveStreamId("li400mYKSgQ6xs7taUeSaEKr") // string | Retrieve video objects that were recorded from a live stream by `liveStreamId`.
    req.SortBy("publishedAt") // string | Use this parameter to sort videos by the their created time, published time, updated time, or by title.
    req.SortOrder("asc") // string | Use this parameter to sort results. `asc` is ascending and sorts from A to Z. `desc` is descending and sorts from Z to A.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.Videos.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.List``: %v\n", err)
    }
    // response from `List`: VideosListResponse
    fmt.Fprintf(os.Stdout, "Response from `Videos.List`: %v\n", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**title** | **string** | The title of a specific video you want to find. The search will match exactly to what term you provide and return any videos that contain the same term as part of their titles. | 
**tags** | **[]string** | A tag is a category you create and apply to videos. You can search for videos with particular tags by listing one or more here. Only videos that have all the tags you list will be returned. | 
**metadata** | [**map[string]**](.md) | Videos can be tagged with metadata tags in key:value pairs. You can search for videos with specific key value pairs using this parameter. [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata) allows you to define a key that allows any value pair. | 
**description** | **string** | Retrieve video objects by &#x60;description&#x60;. | 
**liveStreamId** | **string** | Retrieve video objects that were recorded from a live stream by &#x60;liveStreamId&#x60;. | 
**sortBy** | **string** | Use this parameter to sort videos by the their created time, published time, updated time, or by title. | 
**sortOrder** | **string** | Use this parameter to sort results. &#x60;asc&#x60; is ascending and sorts from A to Z. &#x60;desc&#x60; is descending and sorts from Z to A. | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**VideosListResponse**](VideosListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadThumbnail

> UploadThumbnailFile(videoId string, file *os.File) (*Video, error)
> UploadThumbnail(videoId string, fileName string, fileReader io.Reader)
> UploadThumbnailFileWithContext(ctx context.Context, videoId string, file *os.File) (*Video, error)
> UploadThumbnailWithContext(ctx context.Context, videoId string, fileName string, fileReader io.Reader)

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
    // client := apivideosdk.SandboxClientBuilder("YOUR_SANDBOX_API_KEY").Build()

    videoId := "videoId_example" // string | Unique identifier of the chosen video 
    file := os.NewFile(1234, "some_file") // *os.File | The image to be added as a thumbnail.


    res, err := client.Videos.UploadThumbnailFile(videoId, file)

    // you can also use a Reader instead of a File:
    // client.Videos.UploadThumbnail(videoId, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.UploadThumbnail``: %v\
", err)
    }
    // response from `UploadThumbnail`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.UploadThumbnail`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | Unique identifier of the chosen video  | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | The image to be added as a thumbnail. The mime type should be image/jpeg, image/png or image/webp. The max allowed size is 8 MiB. | 

### Return type

[**Video**](Video.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PickThumbnail

> PickThumbnail(videoId string, videoThumbnailPickPayload VideoThumbnailPickPayload) (*Video, error)

> PickThumbnailWithContext(ctx context.Context, videoId string, videoThumbnailPickPayload VideoThumbnailPickPayload) (*Video, error)


Set a thumbnail



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
    // client := apivideosdk.SandboxClientBuilder("YOUR_SANDBOX_API_KEY").Build()
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | Unique identifier of the video you want to add a thumbnail to, where you use a section of your video as the thumbnail.
    videoThumbnailPickPayload := *apivideosdk.NewVideoThumbnailPickPayload("Timecode_example") // VideoThumbnailPickPayload | 

    
    res, err := client.Videos.PickThumbnail(videoId, videoThumbnailPickPayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.PickThumbnail``: %v\
", err)
    }
    // response from `PickThumbnail`: Video
    fmt.Fprintf(os.Stdout, "Response from `Videos.PickThumbnail`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | Unique identifier of the video you want to add a thumbnail to, where you use a section of your video as the thumbnail. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoThumbnailPickPayload** | [**VideoThumbnailPickPayload**](VideoThumbnailPickPayload.md) |  | 

### Return type

[**Video**](Video.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> GetStatus(videoId string) (*VideoStatus, error)

> GetStatusWithContext(ctx context.Context, videoId string) (*VideoStatus, error)


Retrieve video status and details



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
    // client := apivideosdk.SandboxClientBuilder("YOUR_SANDBOX_API_KEY").Build()
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want the status for.

    
    res, err := client.Videos.GetStatus(videoId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Videos.GetStatus``: %v\
", err)
    }
    // response from `GetStatus`: VideoStatus
    fmt.Fprintf(os.Stdout, "Response from `Videos.GetStatus`: %v\
", res)
}             
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want the status for. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**VideoStatus**](VideoStatus.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

