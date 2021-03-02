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
    webhookId := "webhookId_example" // string | The webhook you wish to delete.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Webhooks.Delete(webhookId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhooks.Delete``: %v\n", err)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The webhook you wish to delete. | 

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

> Get(webhookId string) (*Webhook, error)


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
    webhookId := "webhookId_example" // string | The unique webhook you wish to retreive details on.

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

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
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The unique webhook you wish to retreive details on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](webhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r WebhooksApiListRequest) (*WebhooksListResponse, error)



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
    events := "video.encoding.quality.completed" // string | The webhook event that you wish to filter on.
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Webhooks.List()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhooks.List``: %v\n", err)
    }
    // response from `List`: WebhooksListResponse
    fmt.Fprintf(os.Stdout, "Response from `Webhooks.List`: %v\n", res)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **events** | **string** | The webhook event that you wish to filter on. | 
 **currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
 **pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**WebhooksListResponse**](webhooks-list-response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Create(webhooksCreatePayload WebhooksCreatePayload) (*Webhook, error)


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
    webhooksCreatePayload := *apivideosdk.NewWebhooksCreatePayload([]string{"Events_example"}, "https://example.com/webhooks") // WebhooksCreatePayload | 

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Webhooks.Create(webhooksCreatePayload)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Webhooks.Create``: %v\n", err)
    }
    // response from `Create`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `Webhooks.Create`: %v\n", res)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhooksCreatePayload** | [**WebhooksCreatePayload**](WebhooksCreatePayload.md) |  | 

### Return type

[**Webhook**](webhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

