# \Webhooks

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](Webhooks.md#Delete) | **Delete** /webhooks/{webhookId} | Delete a Webhook
[**Get**](Webhooks.md#Get) | **Get** /webhooks/{webhookId} | Show Webhook details
[**List**](Webhooks.md#List) | **Get** /webhooks | List all webhooks
[**Create**](Webhooks.md#Create) | **Post** /webhooks | Create Webhook



## Delete

> Delete(webhookId string) (error)

> DeleteWithContext(ctx context.Context, webhookId string) (error)


Delete a Webhook



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
        
    webhookId := "webhookId_example" // string | The webhook you wish to delete.

    
    err := client.Webhooks.Delete(webhookId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhooks.Delete``: %v\n", err)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**webhookId** | **string** | The webhook you wish to delete. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(webhookId string) (*Webhook, error)

> GetWithContext(ctx context.Context, webhookId string) (*Webhook, error)


Show Webhook details



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
        
    webhookId := "webhookId_example" // string | The unique webhook you wish to retreive details on.

    
    res, err := client.Webhooks.Get(webhookId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhooks.Get``: %v\n", err)
    }
    // response from `Get`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `Webhooks.Get`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**webhookId** | **string** | The unique webhook you wish to retreive details on. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Webhook**](webhook.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r WebhooksApiListRequest) (*WebhooksListResponse, error)


> ListWithContext(ctx context.Context, r WebhooksApiListRequest) (*WebhooksListResponse, error)



List all webhooks



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
    req := apivideosdk.WebhooksApiListRequest{}
    
    req.Events("video.encoding.quality.completed") // string | The webhook event that you wish to filter on.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.Webhooks.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhooks.List``: %v\n", err)
    }
    // response from `List`: WebhooksListResponse
    fmt.Fprintf(os.Stdout, "Response from `Webhooks.List`: %v\n", res)
}
```

### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**events** | **string** | The webhook event that you wish to filter on. | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**WebhooksListResponse**](webhooks-list-response.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Create(webhooksCreationPayload WebhooksCreationPayload) (*Webhook, error)

> CreateWithContext(ctx context.Context, webhooksCreationPayload WebhooksCreationPayload) (*Webhook, error)


Create Webhook



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
        
    webhooksCreationPayload := *apivideosdk.NewWebhooksCreationPayload([]string{"Events_example"}, "https://example.com/webhooks") // WebhooksCreationPayload | 

    
    res, err := client.Webhooks.Create(webhooksCreationPayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhooks.Create``: %v\n", err)
    }
    // response from `Create`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `Webhooks.Create`: %v\n", res)
}
```

### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**webhooksCreationPayload** | [**WebhooksCreationPayload**](WebhooksCreationPayload.md) |  | 

### Return type

[**Webhook**](webhook.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

