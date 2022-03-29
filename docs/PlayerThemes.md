# \PlayerThemes

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](PlayerThemes.md#Delete) | **Delete** /players/{playerId} | Delete a player
[**DeleteLogo**](PlayerThemes.md#DeleteLogo) | **Delete** /players/{playerId}/logo | Delete logo
[**List**](PlayerThemes.md#List) | **Get** /players | List all player themes
[**Get**](PlayerThemes.md#Get) | **Get** /players/{playerId} | Retrieve a player
[**Update**](PlayerThemes.md#Update) | **Patch** /players/{playerId} | Update a player
[**Create**](PlayerThemes.md#Create) | **Post** /players | Create a player
[**UploadLogo**](PlayerThemes.md#UploadLogo) | **Post** /players/{playerId}/logo | Upload a logo



## Delete

> Delete(playerId string) (error)

> DeleteWithContext(ctx context.Context, playerId string) (error)


Delete a player



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
        
    playerId := "pl45d5vFFGrfdsdsd156dGhh" // string | The unique identifier for the player you want to delete.

    
    err := client.PlayerThemes.Delete(playerId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerThemes.Delete``: %v\
", err)
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

> DeleteLogoWithContext(ctx context.Context, playerId string) (error)


Delete logo



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
        
    playerId := "pl14Db6oMJRH6SRVoOwORacK" // string | The unique identifier for the player.

    
    err := client.PlayerThemes.DeleteLogo(playerId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerThemes.DeleteLogo``: %v\
", err)
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

> List(r PlayerThemesApiListRequest) (*PlayerThemesListResponse, error)


> ListWithContext(ctx context.Context, r PlayerThemesApiListRequest) (*PlayerThemesListResponse, error)



List all player themes



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
    req := apivideosdk.PlayerThemesApiListRequest{}
    
    req.SortBy("createdAt") // string | createdAt is the time the player was created. updatedAt is the time the player was last updated. The time is presented in ISO-8601 format.
    req.SortOrder("asc") // string | Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.PlayerThemes.List(req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerThemes.List``: %v\
", err)
    }
    // response from `List`: PlayerThemesListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlayerThemes.List`: %v\
", res)
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

[**PlayerThemesListResponse**](PlayerThemesListResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Get(playerId string) (*PlayerTheme, error)

> GetWithContext(ctx context.Context, playerId string) (*PlayerTheme, error)


Retrieve a player



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
        
    playerId := "pl45d5vFFGrfdsdsd156dGhh" // string | The unique identifier for the player you want to retrieve. 

    
    res, err := client.PlayerThemes.Get(playerId)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerThemes.Get``: %v\
", err)
    }
    // response from `Get`: PlayerTheme
    fmt.Fprintf(os.Stdout, "Response from `PlayerThemes.Get`: %v\
", res)
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

[**PlayerTheme**](PlayerTheme.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(playerId string, playerThemeUpdatePayload PlayerThemeUpdatePayload) (*PlayerTheme, error)

> UpdateWithContext(ctx context.Context, playerId string, playerThemeUpdatePayload PlayerThemeUpdatePayload) (*PlayerTheme, error)


Update a player



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
        
    playerId := "pl45d5vFFGrfdsdsd156dGhh" // string | The unique identifier for the player.
    playerThemeUpdatePayload := *apivideosdk.NewPlayerThemeUpdatePayload() // PlayerThemeUpdatePayload | 

    
    res, err := client.PlayerThemes.Update(playerId, playerThemeUpdatePayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerThemes.Update``: %v\
", err)
    }
    // response from `Update`: PlayerTheme
    fmt.Fprintf(os.Stdout, "Response from `PlayerThemes.Update`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerId** | **string** | The unique identifier for the player. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerThemeUpdatePayload** | [**PlayerThemeUpdatePayload**](PlayerThemeUpdatePayload.md) |  | 

### Return type

[**PlayerTheme**](PlayerTheme.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Create(playerThemeCreationPayload PlayerThemeCreationPayload) (*PlayerTheme, error)

> CreateWithContext(ctx context.Context, playerThemeCreationPayload PlayerThemeCreationPayload) (*PlayerTheme, error)


Create a player



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
        
    playerThemeCreationPayload := *apivideosdk.NewPlayerThemeCreationPayload() // PlayerThemeCreationPayload | 

    
    res, err := client.PlayerThemes.Create(playerThemeCreationPayload)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerThemes.Create``: %v\
", err)
    }
    // response from `Create`: PlayerTheme
    fmt.Fprintf(os.Stdout, "Response from `PlayerThemes.Create`: %v\
", res)
}
```
### Path Parameters



### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**playerThemeCreationPayload** | [**PlayerThemeCreationPayload**](PlayerThemeCreationPayload.md) |  | 

### Return type

[**PlayerTheme**](PlayerTheme.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLogo

> UploadLogoFile(playerId string, file *os.File) (*PlayerTheme, error)
> UploadLogo(playerId string, fileName string, fileReader io.Reader)
> UploadLogoFileWithContext(ctx context.Context, playerId string, file *os.File) (*PlayerTheme, error)
> UploadLogoWithContext(ctx context.Context, playerId string, fileName string, fileReader io.Reader)

Upload a logo



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
        
    playerId := "pl14Db6oMJRH6SRVoOwORacK" // string | The unique identifier for the player.
    file := os.NewFile(1234, "some_file") // *os.File | The name of the file you want to use for your logo.
    link := "link_example" // string | A public link that you want to advertise in your player. For example, you could add a link to your company. When a viewer clicks on your logo, they will be taken to this address.

    
    res, err := client.PlayerThemes.UploadLogoFile(playerId, file)

    // you can also use a Reader instead of a File:
    // client.PlayerThemes.UploadLogo(playerId, fileName, fileReader)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerThemes.UploadLogo``: %v\
", err)
    }
    // response from `UploadLogo`: PlayerTheme
    fmt.Fprintf(os.Stdout, "Response from `PlayerThemes.UploadLogo`: %v\
", res)
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
**link** | **string** | A public link that you want to advertise in your player. For example, you could add a link to your company. When a viewer clicks on your logo, they will be taken to this address. | 

### Return type

[**PlayerTheme**](PlayerTheme.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

