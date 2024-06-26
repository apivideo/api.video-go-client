# \Captions

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](Captions.md#Upload) | **Post** /videos/{videoId}/captions/{language} | Upload a caption
[**Get**](Captions.md#Get) | **Get** /videos/{videoId}/captions/{language} | Retrieve a caption
[**Update**](Captions.md#Update) | **Patch** /videos/{videoId}/captions/{language} | Update a caption
[**Delete**](Captions.md#Delete) | **Delete** /videos/{videoId}/captions/{language} | Delete a caption
[**List**](Captions.md#List) | **Get** /videos/{videoId}/captions | List video captions



## Upload

> UploadFile(videoId string, language string, file *os.File) (*Caption, error)
> Upload(videoId string, language string, fileName string, fileReader io.Reader)
> UploadFileWithContext(ctx context.Context, videoId string, language string, file *os.File) (*Caption, error)
> UploadWithContext(ctx context.Context, videoId string, language string, fileName string, fileReader io.Reader)

Upload a caption



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want to add a caption to.
    language := "en" // string | A valid language identifier using IETF language tags. You can use primary subtags like `en` (English), extended subtags like `fr-CA` (French, Canada), or region subtags like `zh-Hans-CN` (Simplified Chinese used in the PRC).  - This parameter **only accepts dashes for separators**, for example `fr-CA`. If you use a different separator in your request, the API returns an error. - When the value in your request does not match any covered language, the API returns an error. - You can find the list of supported tags [here](https://docs.api.video/vod/add-captions#supported-caption-language-tags).
    file := os.NewFile(1234, "some_file") // *os.File | The video text track (VTT) you want to upload.

    
    res, err := client.Captions.UploadFile(videoId, language, file)

    // you can also use a Reader instead of a File:
    // client.Captions.Upload(videoId, language, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Captions.Upload``: %v\n", err)
    }
    // response from `Upload`: Caption
    fmt.Fprintf(os.Stdout, "Response from `Captions.Upload`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to add a caption to. | 
**language** | **string** | A valid language identifier using IETF language tags. You can use primary subtags like &#x60;en&#x60; (English), extended subtags like &#x60;fr-CA&#x60; (French, Canada), or region subtags like &#x60;zh-Hans-CN&#x60; (Simplified Chinese used in the PRC).  - This parameter **only accepts dashes for separators**, for example &#x60;fr-CA&#x60;. If you use a different separator in your request, the API returns an error. - When the value in your request does not match any covered language, the API returns an error. - You can find the list of supported tags [here](https://docs.api.video/vod/add-captions#supported-caption-language-tags). | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | The video text track (VTT) you want to upload. | 

### Return type

[**Caption**](Caption.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(videoId string, language string) (*Caption, error)

> GetWithContext(ctx context.Context, videoId string, language string) (*Caption, error)


Retrieve a caption



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want captions for.
    language := "en" // string | A valid language identifier using IETF language tags. You can use primary subtags like `en` (English), extended subtags like `fr-CA` (French, Canada), or region subtags like `zh-Hans-CN` (Simplified Chinese used in the PRC).  - This parameter **only accepts dashes for separators**, for example `fr-CA`. If you use a different separator in your request, the API returns an error. - When the value in your request does not match any covered language, the API returns an error. - You can find the list of supported tags [here](https://docs.api.video/vod/add-captions#supported-caption-language-tags).

    
    res, err := client.Captions.Get(videoId, language)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Captions.Get``: %v\n", err)
    }
    // response from `Get`: Caption
    fmt.Fprintf(os.Stdout, "Response from `Captions.Get`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want captions for. | 
**language** | **string** | A valid language identifier using IETF language tags. You can use primary subtags like &#x60;en&#x60; (English), extended subtags like &#x60;fr-CA&#x60; (French, Canada), or region subtags like &#x60;zh-Hans-CN&#x60; (Simplified Chinese used in the PRC).  - This parameter **only accepts dashes for separators**, for example &#x60;fr-CA&#x60;. If you use a different separator in your request, the API returns an error. - When the value in your request does not match any covered language, the API returns an error. - You can find the list of supported tags [here](https://docs.api.video/vod/add-captions#supported-caption-language-tags). | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Caption**](Caption.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(videoId string, language string, captionsUpdatePayload CaptionsUpdatePayload) (*Caption, error)

> UpdateWithContext(ctx context.Context, videoId string, language string, captionsUpdatePayload CaptionsUpdatePayload) (*Caption, error)


Update a caption



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want to have automatic captions for.
    language := "en" // string | A valid language identifier using IETF language tags. You can use primary subtags like `en` (English), extended subtags like `fr-CA` (French, Canada), or region subtags like `zh-Hans-CN` (Simplified Chinese used in the PRC).  - This parameter **only accepts dashes for separators**, for example `fr-CA`. If you use a different separator in your request, the API returns an error. - When the value in your request does not match any covered language, the API returns an error. - You can find the list of supported tags [here](https://docs.api.video/vod/add-captions#supported-caption-language-tags).
    captionsUpdatePayload := *apivideosdk.NewCaptionsUpdatePayload() // CaptionsUpdatePayload | 

    
    res, err := client.Captions.Update(videoId, language, captionsUpdatePayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Captions.Update``: %v\n", err)
    }
    // response from `Update`: Caption
    fmt.Fprintf(os.Stdout, "Response from `Captions.Update`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to have automatic captions for. | 
**language** | **string** | A valid language identifier using IETF language tags. You can use primary subtags like &#x60;en&#x60; (English), extended subtags like &#x60;fr-CA&#x60; (French, Canada), or region subtags like &#x60;zh-Hans-CN&#x60; (Simplified Chinese used in the PRC).  - This parameter **only accepts dashes for separators**, for example &#x60;fr-CA&#x60;. If you use a different separator in your request, the API returns an error. - When the value in your request does not match any covered language, the API returns an error. - You can find the list of supported tags [here](https://docs.api.video/vod/add-captions#supported-caption-language-tags). | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**captionsUpdatePayload** | [**CaptionsUpdatePayload**](CaptionsUpdatePayload.md) |  | 

### Return type

[**Caption**](Caption.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(videoId string, language string) (error)

> DeleteWithContext(ctx context.Context, videoId string, language string) (error)


Delete a caption



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklgc" // string | The unique identifier for the video you want to delete a caption from.
    language := "en" // string | A valid language identifier using IETF language tags. You can use primary subtags like `en` (English), extended subtags like `fr-CA` (French, Canada), or region subtags like `zh-Hans-CN` (Simplified Chinese used in the PRC).  - This parameter **only accepts dashes for separators**, for example `fr-CA`. If you use a different separator in your request, the API returns an error. - When the value in your request does not match any covered language, the API returns an error. - You can find the list of supported tags [here](https://docs.api.video/vod/add-captions#supported-caption-language-tags).

    
    err := client.Captions.Delete(videoId, language)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Captions.Delete``: %v\n", err)
    }
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to delete a caption from. | 
**language** | **string** | A valid language identifier using IETF language tags. You can use primary subtags like &#x60;en&#x60; (English), extended subtags like &#x60;fr-CA&#x60; (French, Canada), or region subtags like &#x60;zh-Hans-CN&#x60; (Simplified Chinese used in the PRC).  - This parameter **only accepts dashes for separators**, for example &#x60;fr-CA&#x60;. If you use a different separator in your request, the API returns an error. - When the value in your request does not match any covered language, the API returns an error. - You can find the list of supported tags [here](https://docs.api.video/vod/add-captions#supported-caption-language-tags). | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(videoId string, r CaptionsApiListRequest) (*CaptionsListResponse, error)


> ListWithContext(ctx context.Context, videoId string, r CaptionsApiListRequest) (*CaptionsListResponse, error)



List video captions



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
    req := apivideosdk.CaptionsApiListRequest{}
    
    req.VideoId("vi4k0jvEUuaTdRAEjQ4Prklg") // string | The unique identifier for the video you want to retrieve a list of captions for.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.Captions.List(videoId string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Captions.List``: %v\n", err)
    }
    // response from `List`: CaptionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `Captions.List`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to retrieve a list of captions for. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**CaptionsListResponse**](CaptionsListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

