# \Account

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](Account.md#Get) | **Get** /account | Show account



## Get

> Get() (*Account, error)


Show account



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

    client := apivideosdk.NewClient("YOUR_API_TOKEN")
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.NewSandboxClient("YOU_SANDBOX_API_TOKEN")

    res, err := client.Account.Get()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Account.Get``: %v\n", err)
    }
    // response from `Get`: Account
    fmt.Fprintf(os.Stdout, "Response from `Account.Get`: %v\n", res)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


### Return type

[**Account**](account.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

