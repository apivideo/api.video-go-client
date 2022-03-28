# \RawStatistics

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLiveStreamSessions**](RawStatistics.md#ListLiveStreamSessions) | **Get** /analytics/live-streams/{liveStreamId} | List live stream player sessions
[**ListSessionEvents**](RawStatistics.md#ListSessionEvents) | **Get** /analytics/sessions/{sessionId}/events | List player session events
[**ListVideoSessions**](RawStatistics.md#ListVideoSessions) | **Get** /analytics/videos/{videoId} | List video player sessions



## ListLiveStreamSessions

> ListLiveStreamSessions(liveStreamId string, r RawStatisticsApiListLiveStreamSessionsRequest) (*RawStatisticsListLiveStreamAnalyticsResponse, error)


> ListLiveStreamSessionsWithContext(ctx context.Context, liveStreamId string, r RawStatisticsApiListLiveStreamSessionsRequest) (*RawStatisticsListLiveStreamAnalyticsResponse, error)



List live stream player sessions

### Example
```go
//install the Go API client
//go get github.com/apivideo/api.video-go-client
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
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
    req := apivideosdk.RawStatisticsApiListLiveStreamSessionsRequest{}
    
    req.LiveStreamId("vi4k0jvEUuaTdRAEjQ4Jfrgz") // string | The unique identifier for the live stream you want to retrieve analytics for.
    req.Period("2019-01-01") // string | Period must have one of the following formats:  - For a day : "2018-01-01", - For a week: "2018-W01", - For a month: "2018-01" - For a year: "2018"  For a range period: -  Date range: "2018-01-01/2018-01-15" 
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.RawStatistics.ListLiveStreamSessions(liveStreamId string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawStatistics.ListLiveStreamSessions``: %v\
", err)
    }
    // response from `ListLiveStreamSessions`: RawStatisticsListLiveStreamAnalyticsResponse
    fmt.Fprintf(os.Stdout, "Response from `RawStatistics.ListLiveStreamSessions`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**liveStreamId** | **string** | The unique identifier for the live stream you want to retrieve analytics for. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**period** | **string** | Period must have one of the following formats:  - For a day : \&quot;2018-01-01\&quot;, - For a week: \&quot;2018-W01\&quot;,  - For a month: \&quot;2018-01\&quot; - For a year: \&quot;2018\&quot; For a range period:  -  Date range: \&quot;2018-01-01/2018-01-15\&quot;  | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**RawStatisticsListLiveStreamAnalyticsResponse**](RawStatisticsListLiveStreamAnalyticsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionEvents

> ListSessionEvents(sessionId string, r RawStatisticsApiListSessionEventsRequest) (*RawStatisticsListPlayerSessionEventsResponse, error)


> ListSessionEventsWithContext(ctx context.Context, sessionId string, r RawStatisticsApiListSessionEventsRequest) (*RawStatisticsListPlayerSessionEventsResponse, error)



List player session events



### Example
```go
//install the Go API client
//go get github.com/apivideo/api.video-go-client
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
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
    req := apivideosdk.RawStatisticsApiListSessionEventsRequest{}
    
    req.SessionId("psEmFwGQUAXR2lFHj5nDOpy") // string | A unique identifier you can use to reference and track a session with.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.RawStatistics.ListSessionEvents(sessionId string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawStatistics.ListSessionEvents``: %v\
", err)
    }
    // response from `ListSessionEvents`: RawStatisticsListPlayerSessionEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `RawStatistics.ListSessionEvents`: %v\
", res)
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

[**RawStatisticsListPlayerSessionEventsResponse**](RawStatisticsListPlayerSessionEventsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVideoSessions

> ListVideoSessions(videoId string, r RawStatisticsApiListVideoSessionsRequest) (*RawStatisticsListSessionsResponse, error)


> ListVideoSessionsWithContext(ctx context.Context, videoId string, r RawStatisticsApiListVideoSessionsRequest) (*RawStatisticsListSessionsResponse, error)



List video player sessions



### Example
```go
//install the Go API client
//go get github.com/apivideo/api.video-go-client
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
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()
    req := apivideosdk.RawStatisticsApiListVideoSessionsRequest{}
    
    req.VideoId("vi4k0jvEUuaTdRAEjQ4Prklg") // string | The unique identifier for the video you want to retrieve session information for.
    req.Period("period_example") // string | Period must have one of the following formats:  - For a day : 2018-01-01, - For a week: 2018-W01, - For a month: 2018-01 - For a year: 2018  For a range period: -  Date range: 2018-01-01/2018-01-15 
    req.Metadata(map[string]string{"key": "Inner_example"}) // map[string]string | Metadata and Dynamic Metadata filter. Send an array of key value pairs you want to filter sessios with.
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.RawStatistics.ListVideoSessions(videoId string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawStatistics.ListVideoSessions``: %v\
", err)
    }
    // response from `ListVideoSessions`: RawStatisticsListSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `RawStatistics.ListVideoSessions`: %v\
", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**videoId** | **string** | The unique identifier for the video you want to retrieve session information for. | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**period** | **string** | Period must have one of the following formats:  - For a day : 2018-01-01, - For a week: 2018-W01,  - For a month: 2018-01 - For a year: 2018 For a range period:  -  Date range: 2018-01-01/2018-01-15  | 
**metadata** | [**map[string]**](.md) | Metadata and [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata) filter. Send an array of key value pairs you want to filter sessios with. | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**RawStatisticsListSessionsResponse**](RawStatisticsListSessionsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

