# \Summaries

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Summaries.md#Create) | **Post** /summaries | Generate video summary
[**Update**](Summaries.md#Update) | **Patch** /summaries/{summaryId}/source | Update summary details
[**Delete**](Summaries.md#Delete) | **Delete** /summaries/{summaryId} | Delete video summary
[**List**](Summaries.md#List) | **Get** /summaries | List summaries
[**GetSummarySource**](Summaries.md#GetSummarySource) | **Get** /summaries/{summaryId}/source | Get summary details



## Create

> Create(summaryCreationPayload SummaryCreationPayload) (*Summary, error)

> CreateWithContext(ctx context.Context, summaryCreationPayload SummaryCreationPayload) (*Summary, error)


Generate video summary



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
        
    summaryCreationPayload := *apivideosdk.NewSummaryCreationPayload("vi4k0jvEUuaTdRAEjQ4Jfrgz") // SummaryCreationPayload | 

    
    res, err := client.Summaries.Create(summaryCreationPayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Summaries.Create``: %v\n", err)
    }
    // response from `Create`: Summary
    fmt.Fprintf(os.Stdout, "Response from `Summaries.Create`: %v\n", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**summaryCreationPayload** | [**SummaryCreationPayload**](SummaryCreationPayload.md) |  | 

### Return type

[**Summary**](Summary.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(summaryId string, summaryUpdatePayload SummaryUpdatePayload) (*SummarySource, error)

> UpdateWithContext(ctx context.Context, summaryId string, summaryUpdatePayload SummaryUpdatePayload) (*SummarySource, error)


Update summary details



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
        
    summaryId := "summary_1CGHWuXjhxmeH4WiZ51234" // string | The unique identifier of the summary source you want to update.
    summaryUpdatePayload := *apivideosdk.NewSummaryUpdatePayload() // SummaryUpdatePayload | 

    
    res, err := client.Summaries.Update(summaryId, summaryUpdatePayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Summaries.Update``: %v\n", err)
    }
    // response from `Update`: SummarySource
    fmt.Fprintf(os.Stdout, "Response from `Summaries.Update`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**summaryId** | **string** | The unique identifier of the summary source you want to update. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**summaryUpdatePayload** | [**SummaryUpdatePayload**](SummaryUpdatePayload.md) |  | 

### Return type

[**SummarySource**](SummarySource.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(summaryId string) (error)

> DeleteWithContext(ctx context.Context, summaryId string) (error)


Delete video summary



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
        
    summaryId := "summary_1CGHWuXjhxmeH4WiZ51234" // string | The unique identifier of the summary you want to delete.

    
    err := client.Summaries.Delete(summaryId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Summaries.Delete``: %v\n", err)
    }
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**summaryId** | **string** | The unique identifier of the summary you want to delete. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r SummariesApiListRequest) (*SummariesListResponse, error)


> ListWithContext(ctx context.Context, r SummariesApiListRequest) (*SummariesListResponse, error)



List summaries



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
    req := apivideosdk.SummariesApiListRequest{}
    
    req.VideoId("vilkR8K3N7yrRcxcMt91234") // string | Use this parameter to filter for a summary that belongs to a specific video.
    req.Origin("auto") // string | Use this parameter to filter for summaries based on the way they were created: automatically or manually via the API.
    req.SourceStatus("auto") // string | Use this parameter to filter for summaries based on the current status of the summary source.  These are the available statuses:  `missing`: the input for a summary is not present. `waiting` : the input video is being processed and a summary will be generated. `failed`: a technical issue prevented summary generation. `completed`: the summary is generated. `unprocessable`: the API rules the source video to be unsuitable for summary generation. An example for this is an input video that has no audio. 
    req.SortBy("createdAt") // string | Use this parameter to choose which field the API will use to sort the response data. The default is `value`.  These are the available fields to sort by:  - `createdAt`: Sorts the results based on date and timestamps when summaries were created. - `updatedAt`: Sorts the results based on date and timestamps when summaries were last updated. - `videoId`: Sorts the results based on video IDs. 
    req.SortOrder("asc") // string | Use this parameter to sort results. `asc` is ascending and sorts from A to Z. `desc` is descending and sorts from Z to A.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.Summaries.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Summaries.List``: %v\n", err)
    }
    // response from `List`: SummariesListResponse
    fmt.Fprintf(os.Stdout, "Response from `Summaries.List`: %v\n", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | Use this parameter to filter for a summary that belongs to a specific video. | 
**origin** | **string** | Use this parameter to filter for summaries based on the way they were created: automatically or manually via the API. | 
**sourceStatus** | **string** | Use this parameter to filter for summaries based on the current status of the summary source.  These are the available statuses:  &#x60;missing&#x60;: the input for a summary is not present. &#x60;waiting&#x60; : the input video is being processed and a summary will be generated. &#x60;failed&#x60;: a technical issue prevented summary generation. &#x60;completed&#x60;: the summary is generated. &#x60;unprocessable&#x60;: the API rules the source video to be unsuitable for summary generation. An example for this is an input video that has no audio.  | 
**sortBy** | **string** | Use this parameter to choose which field the API will use to sort the response data. The default is &#x60;value&#x60;.  These are the available fields to sort by:  - &#x60;createdAt&#x60;: Sorts the results based on date and timestamps when summaries were created. - &#x60;updatedAt&#x60;: Sorts the results based on date and timestamps when summaries were last updated. - &#x60;videoId&#x60;: Sorts the results based on video IDs.  | 
**sortOrder** | **string** | Use this parameter to sort results. &#x60;asc&#x60; is ascending and sorts from A to Z. &#x60;desc&#x60; is descending and sorts from Z to A. | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**SummariesListResponse**](SummariesListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSummarySource

> GetSummarySource(summaryId string) (*SummarySource, error)

> GetSummarySourceWithContext(ctx context.Context, summaryId string) (*SummarySource, error)


Get summary details



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
        
    summaryId := "summary_1CGHWuXjhxmeH4WiZ51234" // string | The unique identifier of the summary source you want to retrieve.

    
    res, err := client.Summaries.GetSummarySource(summaryId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Summaries.GetSummarySource``: %v\n", err)
    }
    // response from `GetSummarySource`: SummarySource
    fmt.Fprintf(os.Stdout, "Response from `Summaries.GetSummarySource`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**summaryId** | **string** | The unique identifier of the summary source you want to retrieve. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**SummarySource**](SummarySource.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

