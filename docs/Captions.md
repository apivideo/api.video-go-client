# \Captions

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](Captions.md#Delete) | **Delete** /videos/{videoId}/captions/{language} | Delete a caption
[**List**](Captions.md#List) | **Get** /videos/{videoId}/captions | List video captions
[**Get**](Captions.md#Get) | **Get** /videos/{videoId}/captions/{language} | Show a caption
[**Update**](Captions.md#Update) | **Patch** /videos/{videoId}/captions/{language} | Update caption
[**Upload**](Captions.md#Upload) | **Post** /videos/{videoId}/captions/{language} | Upload a caption



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
    apivideosdk "github.com/apivideo/go-api-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklgc" // string | The unique identifier for the video you want to delete a caption from.
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.

    
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
**language** | **string** | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation. | 

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
    apivideosdk "github.com/apivideo/go-api-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
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

[**CaptionsListResponse**](captions-list-response.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(videoId string, language string) (*Subtitle, error)

> GetWithContext(ctx context.Context, videoId string, language string) (*Subtitle, error)


Show a caption



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want captions for.
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation

    
    res, err := client.Captions.Get(videoId, language)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Captions.Get``: %v\n", err)
    }
    // response from `Get`: Subtitle
    fmt.Fprintf(os.Stdout, "Response from `Captions.Get`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want captions for. | 
**language** | **string** | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Subtitle**](subtitle.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(videoId string, language string, captionsUpdatePayload CaptionsUpdatePayload) (*Subtitle, error)

> UpdateWithContext(ctx context.Context, videoId string, language string, captionsUpdatePayload CaptionsUpdatePayload) (*Subtitle, error)


Update caption



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want to have automatic captions for. 
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.
    captionsUpdatePayload := *apivideosdk.NewCaptionsUpdatePayload() // CaptionsUpdatePayload | 

    
    res, err := client.Captions.Update(videoId, language, captionsUpdatePayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Captions.Update``: %v\n", err)
    }
    // response from `Update`: Subtitle
    fmt.Fprintf(os.Stdout, "Response from `Captions.Update`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to have automatic captions for.  | 
**language** | **string** | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**captionsUpdatePayload** | [**CaptionsUpdatePayload**](CaptionsUpdatePayload.md) |  | 

### Return type

[**Subtitle**](subtitle.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upload

> UploadFile(videoId string, language string, file *os.File) (*Subtitle, error)
> Upload(videoId string, language string, fileName string, fileReader io.Reader)
> UploadFileWithContext(ctx context.Context, videoId string, language string, file *os.File) (*Subtitle, error)
> UploadWithContext(ctx context.Context, videoId string, language string, fileName string, fileReader io.Reader)

Upload a caption



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want to add a caption to.
    language := "en" // string | A valid BCP 47 language representation.
    file := os.NewFile(1234, "some_file") // *os.File | The video text track (VTT) you want to upload.

    
    res, err := client.Captions.UploadFile(videoId, language, file)

    // you can also use a Reader instead of a File:
    // client.Captions.Upload(videoId, language, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Captions.Upload``: %v\n", err)
    }
    // response from `Upload`: Subtitle
    fmt.Fprintf(os.Stdout, "Response from `Captions.Upload`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to add a caption to. | 
**language** | **string** | A valid BCP 47 language representation. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | The video text track (VTT) you want to upload. | 

### Return type

[**Subtitle**](subtitle.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

