# \Summaries

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Summaries.md#Create) | **Post** /summaries | Generate video summary
[**Get**](Summaries.md#Get) | **Get** /summaries/{summaryId}/source | Get summary details
[**Update**](Summaries.md#Update) | **Patch** /summaries/{summaryId}/source | Update summary details
[**Delete**](Summaries.md#Delete) | **Delete** /summaries/{summaryId} | Delete video summary
[**List**](Summaries.md#List) | **Get** /summaries | List summaries



## Create

> Create(inlineObject InlineObject) (*SummaryObject, error)

> CreateWithContext(ctx context.Context, inlineObject InlineObject) (*SummaryObject, error)


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
        
    inlineObject := *apivideosdk.NewInlineObject() // InlineObject | 

    
    res, err := client.Summaries.Create(inlineObject)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Summaries.Create``: %v\n", err)
    }
    // response from `Create`: SummaryObject
    fmt.Fprintf(os.Stdout, "Response from `Summaries.Create`: %v\n", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**SummaryObject**](SummaryObject.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(summaryId string) (*SummarySource, error)

> GetWithContext(ctx context.Context, summaryId string) (*SummarySource, error)


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

    
    res, err := client.Summaries.Get(summaryId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Summaries.Get``: %v\n", err)
    }
    // response from `Get`: SummarySource
    fmt.Fprintf(os.Stdout, "Response from `Summaries.Get`: %v\n", res)
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


## Update

> Update(summaryId string, updateSummaryRequest UpdateSummaryRequest) (*SummarySource, error)

> UpdateWithContext(ctx context.Context, summaryId string, updateSummaryRequest UpdateSummaryRequest) (*SummarySource, error)


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
    updateSummaryRequest := *apivideosdk.NewUpdateSummaryRequest() // UpdateSummaryRequest | 

    
    res, err := client.Summaries.Update(summaryId, updateSummaryRequest)

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
**updateSummaryRequest** | [**UpdateSummaryRequest**](UpdateSummaryRequest.md) |  | 

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

> List() (*GetSummaries, error)

> ListWithContext(ctx context.Context, ) (*GetSummaries, error)


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
        

    
    res, err := client.Summaries.List()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Summaries.List``: %v\n", err)
    }
    // response from `List`: GetSummaries
    fmt.Fprintf(os.Stdout, "Response from `Summaries.List`: %v\n", res)
}
```
### Path Parameters

This endpoint does not need any parameter.

### Other Parameters



### Return type

[**GetSummaries**](GetSummaries.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

