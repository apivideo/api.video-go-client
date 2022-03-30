# \Watermarks

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](Watermarks.md#Upload) | **Post** /watermarks | Upload a watermark
[**Delete**](Watermarks.md#Delete) | **Delete** /watermarks/{watermarkId} | Delete a watermark
[**List**](Watermarks.md#List) | **Get** /watermarks | List all watermarks



## Upload

> UploadFile(file *os.File) (*Watermark, error)
> Upload(fileName string, fileReader io.Reader)
> UploadFileWithContext(ctx context.Context, file *os.File) (*Watermark, error)
> UploadWithContext(ctx context.Context, fileName string, fileReader io.Reader)

Upload a watermark



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

    file, _ := os.Open("./watermark.jpg")
    
    res, err := client.Watermarks.UploadFile(file)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Watermarks.UploadFile``: %v", err)
    }
    // response from `UploadFile`: Watermark
    fmt.Fprintf(os.Stdout, "Response from `Watermarks.UploadFile`: %v", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | The &#x60;.jpg&#x60; or &#x60;.png&#x60; image to be added as a watermark. | 

### Return type

[**Watermark**](Watermark.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(watermarkId string) (error)

> DeleteWithContext(ctx context.Context, watermarkId string) (error)


Delete a watermark



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
        
    watermarkId := "watermark_1BWr2L5MTQwxGkuxKjzh6i" // string | The watermark ID for the watermark you want to delete.

    
    err := client.Watermarks.Delete(watermarkId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Watermarks.Delete``: %v\n", err)
    }
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**watermarkId** | **string** | The watermark ID for the watermark you want to delete. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r WatermarksApiListRequest) (*WatermarksListResponse, error)


> ListWithContext(ctx context.Context, r WatermarksApiListRequest) (*WatermarksListResponse, error)



List all watermarks



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
    req := apivideosdk.WatermarksApiListRequest{}
    
    req.SortBy("createdAt") // string | Allowed: createdAt. You can search by the time watermark were created at.
    req.SortOrder("asc") // string | Allowed: asc, desc. asc is ascending and sorts from A to Z. desc is descending and sorts from Z to A.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.Watermarks.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Watermarks.List``: %v\n", err)
    }
    // response from `List`: WatermarksListResponse
    fmt.Fprintf(os.Stdout, "Response from `Watermarks.List`: %v\n", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sortBy** | **string** | Allowed: createdAt. You can search by the time watermark were created at. | 
**sortOrder** | **string** | Allowed: asc, desc. asc is ascending and sorts from A to Z. desc is descending and sorts from Z to A. | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**WatermarksListResponse**](WatermarksListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

