# \Chapters

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](Chapters.md#Upload) | **Post** /videos/{videoId}/chapters/{language} | Upload a chapter
[**Get**](Chapters.md#Get) | **Get** /videos/{videoId}/chapters/{language} | Retrieve a chapter
[**Delete**](Chapters.md#Delete) | **Delete** /videos/{videoId}/chapters/{language} | Delete a chapter
[**List**](Chapters.md#List) | **Get** /videos/{videoId}/chapters | List video chapters



## Upload

> UploadFile(videoId string, language string, file *os.File) (*Chapter, error)
> Upload(videoId string, language string, fileName string, fileReader io.Reader)
> UploadFileWithContext(ctx context.Context, videoId string, language string, file *os.File) (*Chapter, error)
> UploadWithContext(ctx context.Context, videoId string, language string, fileName string, fileReader io.Reader)

Upload a chapter



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want to upload a chapter for.
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.
    file := os.NewFile(1234, "some_file") // *os.File | The VTT file describing the chapters you want to upload.

    
    res, err := client.Chapters.UploadFile(videoId, language, file)

    // you can also use a Reader instead of a File:
    // client.Chapters.Upload(videoId, language, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Chapters.Upload``: %v\
", err)
    }
    // response from `Upload`: Chapter
    fmt.Fprintf(os.Stdout, "Response from `Chapters.Upload`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to upload a chapter for. | 
**language** | **string** | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | The VTT file describing the chapters you want to upload. | 

### Return type

[**Chapter**](Chapter.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(videoId string, language string) (*Chapter, error)

> GetWithContext(ctx context.Context, videoId string, language string) (*Chapter, error)


Retrieve a chapter



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want to show a chapter for.
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.

    
    res, err := client.Chapters.Get(videoId, language)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Chapters.Get``: %v\
", err)
    }
    // response from `Get`: Chapter
    fmt.Fprintf(os.Stdout, "Response from `Chapters.Get`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to show a chapter for. | 
**language** | **string** | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Chapter**](Chapter.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(videoId string, language string) (error)

> DeleteWithContext(ctx context.Context, videoId string, language string) (error)


Delete a chapter



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want to delete a chapter from.
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.

    
    err := client.Chapters.Delete(videoId, language)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Chapters.Delete``: %v\
", err)
    }
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to delete a chapter from. | 
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

> List(videoId string, r ChaptersApiListRequest) (*ChaptersListResponse, error)


> ListWithContext(ctx context.Context, videoId string, r ChaptersApiListRequest) (*ChaptersListResponse, error)



List video chapters



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
        
    videoId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the video you want to show a chapter for.
    language := "en" // string | A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.

    
    res, err := client.Chapters.Get(videoId, language)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Chapters.Get``: %v\
", err)
    }
    // response from `Get`: Chapter
    fmt.Fprintf(os.Stdout, "Response from `Chapters.Get`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to retrieve a list of chapters for. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**ChaptersListResponse**](ChaptersListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

