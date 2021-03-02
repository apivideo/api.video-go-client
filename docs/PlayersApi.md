# \PlayersApi

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](PlayersApi.md#Delete) | **Delete** /players/{playerId} | Delete a player
[**DeleteLogo**](PlayersApi.md#DeleteLogo) | **Delete** /players/{playerId}/logo | Delete logo
[**List**](PlayersApi.md#List) | **Get** /players | List all players
[**Get**](PlayersApi.md#Get) | **Get** /players/{playerId} | Show a player
[**Update**](PlayersApi.md#Update) | **Patch** /players/{playerId} | Update a player
[**Create**](PlayersApi.md#Create) | **Post** /players | Create a player
[**UploadLogo**](PlayersApi.md#UploadLogo) | **Post** /players/{playerId}/logo | Upload a logo



## Delete

> Delete(ctx, playerId).Execute()

Delete a player



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
    playerId := "pl45d5vFFGrfdsdsd156dGhh" // string | The unique identifier for the player you want to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlayersApi.Delete(context.Background(), playerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | The unique identifier for the player you want to delete. | 

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


## DeleteLogo

> map[string]interface{} DeleteLogo(ctx, playerId).Execute()

Delete logo

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
    playerId := "pl14Db6oMJRH6SRVoOwORacK" // string | The unique identifier for the player.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlayersApi.DeleteLogo(context.Background(), playerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.DeleteLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.DeleteLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | The unique identifier for the player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> PlayersListResponse List(ctx).SortBy(sortBy).SortOrder(sortOrder).CurrentPage(currentPage).PageSize(pageSize).Execute()

List all players



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
    sortBy := "createdAt" // string | createdAt is the time the player was created. updatedAt is the time the player was last updated. The time is presented in ISO-8601 format. (optional)
    sortOrder := "asc" // string | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones. (optional)
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (optional) (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlayersApi.List(context.Background()).SortBy(sortBy).SortOrder(sortOrder).CurrentPage(currentPage).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: PlayersListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | createdAt is the time the player was created. updatedAt is the time the player was last updated. The time is presented in ISO-8601 format. | 
 **sortOrder** | **string** | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones. | 
 **currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
 **pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**PlayersListResponse**](players-list-response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Player Get(ctx, playerId).Execute()

Show a player



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
    playerId := "pl45d5vFFGrfdsdsd156dGhh" // string | The unique identifier for the player you want to retrieve. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlayersApi.Get(context.Background(), playerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Player
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | The unique identifier for the player you want to retrieve.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Player**](player.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Player Update(ctx, playerId).PlayerUpdatePayload(playerUpdatePayload).Execute()

Update a player



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
    playerId := "pl45d5vFFGrfdsdsd156dGhh" // string | The unique identifier for the player.
    playerUpdatePayload := *openapiclient.NewPlayerUpdatePayload() // PlayerUpdatePayload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlayersApi.Update(context.Background(), playerId).PlayerUpdatePayload(playerUpdatePayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Player
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | The unique identifier for the player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **playerUpdatePayload** | [**PlayerUpdatePayload**](PlayerUpdatePayload.md) |  | 

### Return type

[**Player**](player.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Player Create(ctx).PlayerCreationPayload(playerCreationPayload).Execute()

Create a player



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
    playerCreationPayload := *openapiclient.NewPlayerCreationPayload() // PlayerCreationPayload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlayersApi.Create(context.Background()).PlayerCreationPayload(playerCreationPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Player
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playerCreationPayload** | [**PlayerCreationPayload**](PlayerCreationPayload.md) |  | 

### Return type

[**Player**](player.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLogo

> Player UploadLogo(ctx, playerId).File(file).Link(link).Execute()

Upload a logo



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
    playerId := "pl14Db6oMJRH6SRVoOwORacK" // string | The unique identifier for the player.
    file := os.NewFile(1234, "some_file") // *os.File | The name of the file you want to use for your logo.
    link := "link_example" // string | The path to the file you want to upload and use as a logo.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlayersApi.UploadLogo(context.Background(), playerId).File(file).Link(link).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayersApi.UploadLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadLogo`: Player
    fmt.Fprintf(os.Stdout, "Response from `PlayersApi.UploadLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | The unique identifier for the player. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The name of the file you want to use for your logo. | 
 **link** | **string** | The path to the file you want to upload and use as a logo. | 

### Return type

[**Player**](player.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

