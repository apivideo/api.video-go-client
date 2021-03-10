# \Players

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](Players.md#Delete) | **Delete** /players/{playerId} | Delete a player
[**DeleteLogo**](Players.md#DeleteLogo) | **Delete** /players/{playerId}/logo | Delete logo
[**List**](Players.md#List) | **Get** /players | List all players
[**Get**](Players.md#Get) | **Get** /players/{playerId} | Show a player
[**Update**](Players.md#Update) | **Patch** /players/{playerId} | Update a player
[**Create**](Players.md#Create) | **Post** /players | Create a player
[**UploadLogo**](Players.md#UploadLogo) | **Post** /players/{playerId}/logo | Upload a logo



## Delete

> Delete(playerId string) (error)


Delete a player



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
        
    playerId := "pl45d5vFFGrfdsdsd156dGhh" // string | The unique identifier for the player you want to delete.

    
    err := client.Players.Delete(playerId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Delete``: %v\n", err)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerId** | **string** | The unique identifier for the player you want to delete. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogo

> DeleteLogo(playerId string) (error)


Delete logo

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
        
    playerId := "pl14Db6oMJRH6SRVoOwORacK" // string | The unique identifier for the player.

    
    err := client.Players.DeleteLogo(playerId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.DeleteLogo``: %v\n", err)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerId** | **string** | The unique identifier for the player. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> List(r PlayersApiListRequest) (*PlayersListResponse, error)



List all players



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
    req := apivideosdk.PlayersApiListRequest{}
    
    req.SortBy("createdAt") // string | createdAt is the time the player was created. updatedAt is the time the player was last updated. The time is presented in ISO-8601 format.
    req.SortOrder("asc") // string | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.Players.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.List``: %v\n", err)
    }
    // response from `List`: PlayersListResponse
    fmt.Fprintf(os.Stdout, "Response from `Players.List`: %v\n", res)
}
```

### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sortBy** | **string** | createdAt is the time the player was created. updatedAt is the time the player was last updated. The time is presented in ISO-8601 format. | 
**sortOrder** | **string** | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones. | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**PlayersListResponse**](players-list-response.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(playerId string) (*Player, error)


Show a player



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
        
    playerId := "pl45d5vFFGrfdsdsd156dGhh" // string | The unique identifier for the player you want to retrieve. 

    
    res, err := client.Players.Get(playerId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Get``: %v\n", err)
    }
    // response from `Get`: Player
    fmt.Fprintf(os.Stdout, "Response from `Players.Get`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerId** | **string** | The unique identifier for the player you want to retrieve.  | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Player**](player.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(playerId string, playerUpdatePayload PlayerUpdatePayload) (*Player, error)


Update a player



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
        
    playerId := "pl45d5vFFGrfdsdsd156dGhh" // string | The unique identifier for the player.
    playerUpdatePayload := *apivideosdk.NewPlayerUpdatePayload() // PlayerUpdatePayload | 

    
    res, err := client.Players.Update(playerId, playerUpdatePayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Update``: %v\n", err)
    }
    // response from `Update`: Player
    fmt.Fprintf(os.Stdout, "Response from `Players.Update`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerId** | **string** | The unique identifier for the player. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerUpdatePayload** | [**PlayerUpdatePayload**](PlayerUpdatePayload.md) |  | 

### Return type

[**Player**](player.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Create(playerCreationPayload PlayerCreationPayload) (*Player, error)


Create a player



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
        
    playerCreationPayload := *apivideosdk.NewPlayerCreationPayload() // PlayerCreationPayload | 

    
    res, err := client.Players.Create(playerCreationPayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.Create``: %v\n", err)
    }
    // response from `Create`: Player
    fmt.Fprintf(os.Stdout, "Response from `Players.Create`: %v\n", res)
}
```

### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerCreationPayload** | [**PlayerCreationPayload**](PlayerCreationPayload.md) |  | 

### Return type

[**Player**](player.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLogo

> UploadLogoFile(playerId string, file *os.File, link string) (*Player, error)

> UploadLogo(playerId string, link string, fileName string, fileReader io.Reader)

Upload a logo



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
        
    playerId := "pl14Db6oMJRH6SRVoOwORacK" // string | The unique identifier for the player.
    file := os.NewFile(1234, "some_file") // *os.File | The name of the file you want to use for your logo.
    link := "link_example" // string | The path to the file you want to upload and use as a logo.

    
    res, err := client.Players.UploadLogoFile(playerId, file, link)

    // you can also use a Reader instead of a File:
    // client.Players.UploadLogo(playerId, link, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Players.UploadLogo``: %v\n", err)
    }
    // response from `UploadLogo`: Player
    fmt.Fprintf(os.Stdout, "Response from `Players.UploadLogo`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerId** | **string** | The unique identifier for the player. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | The name of the file you want to use for your logo. | 
**link** | **string** | The path to the file you want to upload and use as a logo. | 

### Return type

[**Player**](player.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

