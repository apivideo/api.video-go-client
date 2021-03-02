# \RawStatisticsApi

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLiveStreamAnalytics**](RawStatisticsApi.md#GetLiveStreamAnalytics) | **Get** /analytics/live-streams/{liveStreamId} | List live stream player sessions
[**ListPlayerSessionEvents**](RawStatisticsApi.md#ListPlayerSessionEvents) | **Get** /analytics/sessions/{sessionId}/events | List player session events
[**ListSessions**](RawStatisticsApi.md#ListSessions) | **Get** /analytics/videos/{videoId} | List video player sessions



## GetLiveStreamAnalytics

> RawStatisticsListLiveStreamAnalyticsResponse GetLiveStreamAnalytics(ctx, liveStreamId).Period(period).CurrentPage(currentPage).PageSize(pageSize).Execute()

List live stream player sessions

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
    liveStreamId := "vi4k0jvEUuaTdRAEjQ4Jfrgz" // string | The unique identifier for the live stream you want to retrieve analytics for.
    period := "2019-01-01" // string | Period must have one of the following formats:   - For a day : \"2018-01-01\", - For a week: \"2018-W01\",  - For a month: \"2018-01\" - For a year: \"2018\"  For a range period:  -  Date range: \"2018-01-01/2018-01-15\"  (optional)
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (optional) (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RawStatisticsApi.GetLiveStreamAnalytics(context.Background(), liveStreamId).Period(period).CurrentPage(currentPage).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawStatisticsApi.GetLiveStreamAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLiveStreamAnalytics`: RawStatisticsListLiveStreamAnalyticsResponse
    fmt.Fprintf(os.Stdout, "Response from `RawStatisticsApi.GetLiveStreamAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveStreamId** | **string** | The unique identifier for the live stream you want to retrieve analytics for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveStreamAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Period must have one of the following formats:   - For a day : \&quot;2018-01-01\&quot;, - For a week: \&quot;2018-W01\&quot;,  - For a month: \&quot;2018-01\&quot; - For a year: \&quot;2018\&quot;  For a range period:  -  Date range: \&quot;2018-01-01/2018-01-15\&quot;  | 
 **currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
 **pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**RawStatisticsListLiveStreamAnalyticsResponse**](raw-statistics-list-live-stream-analytics-response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlayerSessionEvents

> RawStatisticsListPlayerSessionEventsResponse ListPlayerSessionEvents(ctx, sessionId).CurrentPage(currentPage).PageSize(pageSize).Execute()

List player session events



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
    sessionId := "psEmFwGQUAXR2lFHj5nDOpy" // string | A unique identifier you can use to reference and track a session with.
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (optional) (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RawStatisticsApi.ListPlayerSessionEvents(context.Background(), sessionId).CurrentPage(currentPage).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawStatisticsApi.ListPlayerSessionEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlayerSessionEvents`: RawStatisticsListPlayerSessionEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `RawStatisticsApi.ListPlayerSessionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | A unique identifier you can use to reference and track a session with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPlayerSessionEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
 **pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**RawStatisticsListPlayerSessionEventsResponse**](raw-statistics-list-player-session-events-response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessions

> RawStatisticsListSessionsResponse ListSessions(ctx, videoId).Period(period).Metadata(metadata).CurrentPage(currentPage).PageSize(pageSize).Execute()

List video player sessions



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
    videoId := "vi4k0jvEUuaTdRAEjQ4Prklg" // string | The unique identifier for the video you want to retrieve session information for.
    period := "period_example" // string | Period must have one of the following formats:   - For a day : 2018-01-01, - For a week: 2018-W01,  - For a month: 2018-01 - For a year: 2018  For a range period:  -  Date range: 2018-01-01/2018-01-15  (optional)
    metadata := []string{"Inner_example"} // []string | Metadata and Dynamic Metadata filter. Send an array of key value pairs you want to filter sessios with. (optional)
    currentPage := int32(2) // int32 | Choose the number of search results to return per page. Minimum value: 1 (optional) (default to 1)
    pageSize := int32(30) // int32 | Results per page. Allowed values 1-100, default is 25. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RawStatisticsApi.ListSessions(context.Background(), videoId).Period(period).Metadata(metadata).CurrentPage(currentPage).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawStatisticsApi.ListSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessions`: RawStatisticsListSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RawStatisticsApi.ListSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | The unique identifier for the video you want to retrieve session information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Period must have one of the following formats:   - For a day : 2018-01-01, - For a week: 2018-W01,  - For a month: 2018-01 - For a year: 2018  For a range period:  -  Date range: 2018-01-01/2018-01-15  | 
 **metadata** | **[]string** | Metadata and Dynamic Metadata filter. Send an array of key value pairs you want to filter sessios with. | 
 **currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
 **pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**RawStatisticsListSessionsResponse**](raw-statistics-list-sessions-response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

