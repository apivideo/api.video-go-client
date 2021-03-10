# \RawStatistics

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLiveStreamAnalytics**](RawStatistics.md#GetLiveStreamAnalytics) | **Get** /analytics/live-streams/{liveStreamId} | List live stream player sessions
[**ListPlayerSessionEvents**](RawStatistics.md#ListPlayerSessionEvents) | **Get** /analytics/sessions/{sessionId}/events | List player session events
[**ListSessions**](RawStatistics.md#ListSessions) | **Get** /analytics/videos/{videoId} | List video player sessions



## GetLiveStreamAnalytics

> GetLiveStreamAnalytics(liveStreamId string, r RawStatisticsApiGetLiveStreamAnalyticsRequest) (*RawStatisticsListLiveStreamAnalyticsResponse, error)



List live stream player sessions

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
    req := apivideosdk.RawStatisticsApiGetLiveStreamAnalyticsRequest{}
    
    req.LiveStreamId("vi4k0jvEUuaTdRAEjQ4Jfrgz") // string | The unique identifier for the live stream you want to retrieve analytics for.
    req.Period("2019-01-01") // string | Period must have one of the following formats:   - For a day : \"2018-01-01\", - For a week: \"2018-W01\",  - For a month: \"2018-01\" - For a year: \"2018\"  For a range period:  -  Date range: \"2018-01-01/2018-01-15\" 
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.RawStatistics.GetLiveStreamAnalytics(liveStreamId string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawStatistics.GetLiveStreamAnalytics``: %v\n", err)
    }
    // response from `GetLiveStreamAnalytics`: RawStatisticsListLiveStreamAnalyticsResponse
    fmt.Fprintf(os.Stdout, "Response from `RawStatistics.GetLiveStreamAnalytics`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamId** | **string** | The unique identifier for the live stream you want to retrieve analytics for. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**period** | **string** | Period must have one of the following formats:   - For a day : \&quot;2018-01-01\&quot;, - For a week: \&quot;2018-W01\&quot;,  - For a month: \&quot;2018-01\&quot; - For a year: \&quot;2018\&quot;  For a range period:  -  Date range: \&quot;2018-01-01/2018-01-15\&quot;  | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**RawStatisticsListLiveStreamAnalyticsResponse**](raw-statistics-list-live-stream-analytics-response.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlayerSessionEvents

> ListPlayerSessionEvents(sessionId string, r RawStatisticsApiListPlayerSessionEventsRequest) (*RawStatisticsListPlayerSessionEventsResponse, error)



List player session events



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
    req := apivideosdk.RawStatisticsApiListPlayerSessionEventsRequest{}
    
    req.SessionId("psEmFwGQUAXR2lFHj5nDOpy") // string | A unique identifier you can use to reference and track a session with.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.RawStatistics.ListPlayerSessionEvents(sessionId string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawStatistics.ListPlayerSessionEvents``: %v\n", err)
    }
    // response from `ListPlayerSessionEvents`: RawStatisticsListPlayerSessionEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `RawStatistics.ListPlayerSessionEvents`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sessionId** | **string** | A unique identifier you can use to reference and track a session with. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**RawStatisticsListPlayerSessionEventsResponse**](raw-statistics-list-player-session-events-response.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessions

> ListSessions(videoId string, r RawStatisticsApiListSessionsRequest) (*RawStatisticsListSessionsResponse, error)



List video player sessions



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
    req := apivideosdk.RawStatisticsApiListSessionsRequest{}
    
    req.VideoId("vi4k0jvEUuaTdRAEjQ4Prklg") // string | The unique identifier for the video you want to retrieve session information for.
    req.Period("period_example") // string | Period must have one of the following formats:   - For a day : 2018-01-01, - For a week: 2018-W01,  - For a month: 2018-01 - For a year: 2018  For a range period:  -  Date range: 2018-01-01/2018-01-15 
    req.Metadata([]string{"Inner_example"}) // []string | Metadata and Dynamic Metadata filter. Send an array of key value pairs you want to filter sessios with.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.RawStatistics.ListSessions(videoId string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawStatistics.ListSessions``: %v\n", err)
    }
    // response from `ListSessions`: RawStatisticsListSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RawStatistics.ListSessions`: %v\n", res)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to retrieve session information for. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**period** | **string** | Period must have one of the following formats:   - For a day : 2018-01-01, - For a week: 2018-W01,  - For a month: 2018-01 - For a year: 2018  For a range period:  -  Date range: 2018-01-01/2018-01-15  | 
**metadata** | **[]string** | Metadata and Dynamic Metadata filter. Send an array of key value pairs you want to filter sessios with. | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**RawStatisticsListSessionsResponse**](raw-statistics-list-sessions-response.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

