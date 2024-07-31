# \Analytics

All URIs are relative to *https://ws.api.video*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregatedMetrics**](Analytics.md#GetAggregatedMetrics) | **Get** /data/metrics/{metric}/{aggregation} | Retrieve aggregated metrics
[**GetMetricsBreakdown**](Analytics.md#GetMetricsBreakdown) | **Get** /data/buckets/{metric}/{breakdown} | Retrieve metrics in a breakdown of dimensions
[**GetMetricsOverTime**](Analytics.md#GetMetricsOverTime) | **Get** /data/timeseries/{metric} | Retrieve metrics over time



## GetAggregatedMetrics

> GetAggregatedMetrics(metric string, aggregation string, r AnalyticsApiGetAggregatedMetricsRequest) (*AnalyticsAggregatedMetricsResponse, error)


> GetAggregatedMetricsWithContext(ctx context.Context, metric string, aggregation string, r AnalyticsApiGetAggregatedMetricsRequest) (*AnalyticsAggregatedMetricsResponse, error)



Retrieve aggregated metrics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_KEY").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_KEY").Build()
    req := apivideosdk.AnalyticsApiGetAggregatedMetricsRequest{}
    
    req.Metric("metric_example") // string | Use this path parameter to select a metric that you want analytics for.  - `play` is the number of times your content has been played. You can use the aggregations `count`, `rate`, and `total` with the `play` metric. - `start` is the number of times playback was started. You can use the aggregation `count` with this metric. - `end` is the number of times playback has ended with the content watch until the end. You can use the aggregation `count` with this metric. - `impression` is the number of times your content has been loaded and was ready for playback. You can use the aggregation `count` with this metric. - `impression-time` is the time in milliseconds that your content was loading for until the first video frame is displayed. You can use the aggregations `average` and `sum` with this metric. - `watch-time` is the cumulative time in seconds that the user has spent watching your content. You can use the aggregations `average` and `sum` with this metric. 
    req.Aggregation("aggregation_example") // string | Use this path parameter to define a way of collecting data for the metric that you want analytics for.  - `count` returns the overall number of events for the `play` metric. - `rate` returns the ratio that calculates the number of plays your content receives divided by its impressions. This aggregation can be used only with the `play` metric. - `total` calculates the total number of events for the `play` metric.  - `average` calculates an average value for the selected metric. - `sum` adds up the total value of the select metric. 
    req.From(time.Now()) // time.Time | Use this query parameter to define the starting date-time of the period you want analytics for.  - If you do not set a value for `from`, the default assigned value is 1 day ago, based on the `to` parameter. - The maximum value is 30 days ago. - The value you provide should follow the ATOM date-time format: `2024-02-05T00:00:00+01:00` - The API ignores this parameter when you call `/data/metrics/play/total`. 
    req.To(time.Now()) // time.Time | Use this query parameter to define the ending date-time of the period you want analytics for.  - If you do not set a value for `to`, the default assigned value is `now`. - The API ignores this parameter when you call `/data/metrics/play/total`. - The value for `to` is a non-inclusive value: the API returns data **before** the date-time that you set. 
    req.FilterBy(map[string][]apivideosdk.FilterBy2{"key": "TODO"}) // FilterBy2 | Use this parameter to filter the API's response based on different data dimensions. You can serialize filters in your query to receive more detailed breakdowns of your analytics.  - If you do not set a value for `filterBy`, the API returns the full dataset for your project. - The API only accepts the `mediaId` and `mediaType` filters when you call `/data/metrics/play/total` or `/data/buckets/play-total/media-id`.  These are the available breakdown dimensions:  - `mediaId`: Returns analytics based on the unique identifiers of a video or a live stream. - `mediaType`: Returns analytics based on the type of content. Possible values: `video` and `live-stream`.  - `continent`: Returns analytics based on the viewers' continent. The list of supported continents names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example `EU`. Possible values are: `AS`, `AF`, `NA`, `SA`, `AN`, `EU`, `AZ`.  - `country`: Returns analytics based on the viewers' country. The list of supported country names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example `FR`. - `deviceType`: Returns analytics based on the type of device used by the viewers. Response values can include: `computer`, `phone`, `tablet`, `tv`, `console`, `wearable`, `unknown`. - `operatingSystem`: Returns analytics based on the operating system used by the viewers. Response values can include `windows`, `mac osx`, `android`, `ios`, `linux`. - `browser`: Returns analytics based on the browser used by the viewers. Response values can include `chrome`, `firefox`, `edge`, `opera`. - `tag`: Returns analytics for videos using this tag. This filter only accepts a single value and is case sensitive. Read more about tagging your videos [here](https://docs.api.video/vod/tags-metadata). 

    res, err := client.Analytics.GetAggregatedMetrics(metric string, aggregation string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetAggregatedMetrics``: %v\n", err)
    }
    // response from `GetAggregatedMetrics`: AnalyticsAggregatedMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `Analytics.GetAggregatedMetrics`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**metric** | **string** | Use this path parameter to select a metric that you want analytics for.  - &#x60;play&#x60; is the number of times your content has been played. You can use the aggregations &#x60;count&#x60;, &#x60;rate&#x60;, and &#x60;total&#x60; with the &#x60;play&#x60; metric. - &#x60;start&#x60; is the number of times playback was started. You can use the aggregation &#x60;count&#x60; with this metric. - &#x60;end&#x60; is the number of times playback has ended with the content watch until the end. You can use the aggregation &#x60;count&#x60; with this metric. - &#x60;impression&#x60; is the number of times your content has been loaded and was ready for playback. You can use the aggregation &#x60;count&#x60; with this metric. - &#x60;impression-time&#x60; is the time in milliseconds that your content was loading for until the first video frame is displayed. You can use the aggregations &#x60;average&#x60; and &#x60;sum&#x60; with this metric. - &#x60;watch-time&#x60; is the cumulative time in seconds that the user has spent watching your content. You can use the aggregations &#x60;average&#x60; and &#x60;sum&#x60; with this metric.  | 
**aggregation** | **string** | Use this path parameter to define a way of collecting data for the metric that you want analytics for.  - &#x60;count&#x60; returns the overall number of events for the &#x60;play&#x60; metric. - &#x60;rate&#x60; returns the ratio that calculates the number of plays your content receives divided by its impressions. This aggregation can be used only with the &#x60;play&#x60; metric. - &#x60;total&#x60; calculates the total number of events for the &#x60;play&#x60; metric.  - &#x60;average&#x60; calculates an average value for the selected metric. - &#x60;sum&#x60; adds up the total value of the select metric.  | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**from** | **time.Time** | Use this query parameter to define the starting date-time of the period you want analytics for.  - If you do not set a value for &#x60;from&#x60;, the default assigned value is 1 day ago, based on the &#x60;to&#x60; parameter. - The maximum value is 30 days ago. - The value you provide should follow the ATOM date-time format: &#x60;2024-02-05T00:00:00+01:00&#x60; - The API ignores this parameter when you call &#x60;/data/metrics/play/total&#x60;.  | 
**to** | **time.Time** | Use this query parameter to define the ending date-time of the period you want analytics for.  - If you do not set a value for &#x60;to&#x60;, the default assigned value is &#x60;now&#x60;. - The API ignores this parameter when you call &#x60;/data/metrics/play/total&#x60;. - The value for &#x60;to&#x60; is a non-inclusive value: the API returns data **before** the date-time that you set.  | 
**filterBy** | [**FilterBy2**](FilterBy2.md) | Use this parameter to filter the API&#39;s response based on different data dimensions. You can serialize filters in your query to receive more detailed breakdowns of your analytics.  - If you do not set a value for &#x60;filterBy&#x60;, the API returns the full dataset for your project. - The API only accepts the &#x60;mediaId&#x60; and &#x60;mediaType&#x60; filters when you call &#x60;/data/metrics/play/total&#x60; or &#x60;/data/buckets/play-total/media-id&#x60;.  These are the available breakdown dimensions:  - &#x60;mediaId&#x60;: Returns analytics based on the unique identifiers of a video or a live stream. - &#x60;mediaType&#x60;: Returns analytics based on the type of content. Possible values: &#x60;video&#x60; and &#x60;live-stream&#x60;.  - &#x60;continent&#x60;: Returns analytics based on the viewers&#39; continent. The list of supported continents names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example &#x60;EU&#x60;. Possible values are: &#x60;AS&#x60;, &#x60;AF&#x60;, &#x60;NA&#x60;, &#x60;SA&#x60;, &#x60;AN&#x60;, &#x60;EU&#x60;, &#x60;AZ&#x60;.  - &#x60;country&#x60;: Returns analytics based on the viewers&#39; country. The list of supported country names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example &#x60;FR&#x60;. - &#x60;deviceType&#x60;: Returns analytics based on the type of device used by the viewers. Response values can include: &#x60;computer&#x60;, &#x60;phone&#x60;, &#x60;tablet&#x60;, &#x60;tv&#x60;, &#x60;console&#x60;, &#x60;wearable&#x60;, &#x60;unknown&#x60;. - &#x60;operatingSystem&#x60;: Returns analytics based on the operating system used by the viewers. Response values can include &#x60;windows&#x60;, &#x60;mac osx&#x60;, &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;linux&#x60;. - &#x60;browser&#x60;: Returns analytics based on the browser used by the viewers. Response values can include &#x60;chrome&#x60;, &#x60;firefox&#x60;, &#x60;edge&#x60;, &#x60;opera&#x60;. - &#x60;tag&#x60;: Returns analytics for videos using this tag. This filter only accepts a single value and is case sensitive. Read more about tagging your videos [here](https://docs.api.video/vod/tags-metadata).  | 

### Return type

[**AnalyticsAggregatedMetricsResponse**](AnalyticsAggregatedMetricsResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsBreakdown

> GetMetricsBreakdown(metric string, breakdown string, r AnalyticsApiGetMetricsBreakdownRequest) (*AnalyticsMetricsBreakdownResponse, error)


> GetMetricsBreakdownWithContext(ctx context.Context, metric string, breakdown string, r AnalyticsApiGetMetricsBreakdownRequest) (*AnalyticsMetricsBreakdownResponse, error)



Retrieve metrics in a breakdown of dimensions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_KEY").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_KEY").Build()
    req := apivideosdk.AnalyticsApiGetMetricsBreakdownRequest{}
    
    req.Metric("metric_example") // string | Use this path parameter to select a metric that you want analytics for.  - `play` is the number of times your content has been played. - `play-rate` is the ratio that calculates the number of plays your content receives divided by its impressions. - `play-total` is the total number of times a specific content has been played. You can only use the `media-id` breakdown with this metric. - `start` is the number of times playback was started. - `end` is the number of times playback has ended with the content watch until the end. - `impression` is the number of times your content has been loaded and was ready for playback. 
    req.Breakdown("breakdown_example") // string | Use this path parameter to define a dimension for segmenting analytics data. You must use `kebab-case` for path parameters.  These are the available dimensions:  - `media-id`: Returns analytics based on the unique identifiers of a video or a live stream. - `media-type`: Returns analytics based on the type of content. Possible values: `video` and `live-stream`.  - `continent`: Returns analytics based on the viewers' continent. The list of supported continents names are based on the [GeoNames public database](https://www.geonames.org/countries/). Possible values are: `AS`, `AF`, `NA`, `SA`, `AN`, `EU`, `AZ`.  - `country`: Returns analytics based on the viewers' country. The list of supported country names are based on the [GeoNames public database](https://www.geonames.org/countries/). - `device-type`: Returns analytics based on the type of device used by the viewers. Response values can include: `computer`, `phone`, `tablet`, `tv`, `console`, `wearable`, `unknown`. - `operating-system`: Returns analytics based on the operating system used by the viewers. Response values can include `windows`, `mac osx`, `android`, `ios`, `linux`. - `browser`: Returns analytics based on the browser used by the viewers. Response values can include `chrome`, `firefox`, `edge`, `opera`. 
    req.From(time.Now()) // time.Time | Use this query parameter to define the starting date-time of the period you want analytics for.  - If you do not set a value for `from`, the default assigned value is 1 day ago, based on the `to` parameter. - The maximum value is 30 days ago. - The value you provide should follow the ATOM date-time format: `2024-02-05T00:00:00+01:00` 
    req.To(time.Now()) // time.Time | Use this query parameter to define the ending date-time of the period you want analytics for.  - If you do not set a value for `to`, the default assigned value is `now`. - The value for `to` is a non-inclusive value: the API returns data **before** the date-time that you set. 
    req.FilterBy(map[string][]apivideosdk.FilterBy2{"key": "TODO"}) // FilterBy2 | Use this parameter to filter the API's response based on different data dimensions. You can serialize filters in your query to receive more detailed breakdowns of your analytics.  - If you do not set a value for `filterBy`, the API returns the full dataset for your project. - The API only accepts the `mediaId` and `mediaType` filters when you call `/data/metrics/play/total` or `/data/buckets/play-total/media-id`.  These are the available breakdown dimensions:  - `mediaId`: Returns analytics based on the unique identifiers of a video or a live stream. - `mediaType`: Returns analytics based on the type of content. Possible values: `video` and `live-stream`.  - `continent`: Returns analytics based on the viewers' continent. The list of supported continents names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example `EU`. Possible values are: `AS`, `AF`, `NA`, `SA`, `AN`, `EU`, `AZ`.  - `country`: Returns analytics based on the viewers' country. The list of supported country names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example `FR`. - `deviceType`: Returns analytics based on the type of device used by the viewers. Response values can include: `computer`, `phone`, `tablet`, `tv`, `console`, `wearable`, `unknown`. - `operatingSystem`: Returns analytics based on the operating system used by the viewers. Response values can include `windows`, `mac osx`, `android`, `ios`, `linux`. - `browser`: Returns analytics based on the browser used by the viewers. Response values can include `chrome`, `firefox`, `edge`, `opera`. - `tag`: Returns analytics for videos using this tag. This filter only accepts a single value and is case sensitive. Read more about tagging your videos [here](https://docs.api.video/vod/tags-metadata). 
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.Analytics.GetMetricsBreakdown(metric string, breakdown string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetMetricsBreakdown``: %v\n", err)
    }
    // response from `GetMetricsBreakdown`: AnalyticsMetricsBreakdownResponse
    fmt.Fprintf(os.Stdout, "Response from `Analytics.GetMetricsBreakdown`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**metric** | **string** | Use this path parameter to select a metric that you want analytics for.  - &#x60;play&#x60; is the number of times your content has been played. - &#x60;play-rate&#x60; is the ratio that calculates the number of plays your content receives divided by its impressions. - &#x60;play-total&#x60; is the total number of times a specific content has been played. You can only use the &#x60;media-id&#x60; breakdown with this metric. - &#x60;start&#x60; is the number of times playback was started. - &#x60;end&#x60; is the number of times playback has ended with the content watch until the end. - &#x60;impression&#x60; is the number of times your content has been loaded and was ready for playback.  | 
**breakdown** | **string** | Use this path parameter to define a dimension for segmenting analytics data. You must use &#x60;kebab-case&#x60; for path parameters.  These are the available dimensions:  - &#x60;media-id&#x60;: Returns analytics based on the unique identifiers of a video or a live stream. - &#x60;media-type&#x60;: Returns analytics based on the type of content. Possible values: &#x60;video&#x60; and &#x60;live-stream&#x60;.  - &#x60;continent&#x60;: Returns analytics based on the viewers&#39; continent. The list of supported continents names are based on the [GeoNames public database](https://www.geonames.org/countries/). Possible values are: &#x60;AS&#x60;, &#x60;AF&#x60;, &#x60;NA&#x60;, &#x60;SA&#x60;, &#x60;AN&#x60;, &#x60;EU&#x60;, &#x60;AZ&#x60;.  - &#x60;country&#x60;: Returns analytics based on the viewers&#39; country. The list of supported country names are based on the [GeoNames public database](https://www.geonames.org/countries/). - &#x60;device-type&#x60;: Returns analytics based on the type of device used by the viewers. Response values can include: &#x60;computer&#x60;, &#x60;phone&#x60;, &#x60;tablet&#x60;, &#x60;tv&#x60;, &#x60;console&#x60;, &#x60;wearable&#x60;, &#x60;unknown&#x60;. - &#x60;operating-system&#x60;: Returns analytics based on the operating system used by the viewers. Response values can include &#x60;windows&#x60;, &#x60;mac osx&#x60;, &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;linux&#x60;. - &#x60;browser&#x60;: Returns analytics based on the browser used by the viewers. Response values can include &#x60;chrome&#x60;, &#x60;firefox&#x60;, &#x60;edge&#x60;, &#x60;opera&#x60;.  | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**from** | **time.Time** | Use this query parameter to define the starting date-time of the period you want analytics for.  - If you do not set a value for &#x60;from&#x60;, the default assigned value is 1 day ago, based on the &#x60;to&#x60; parameter. - The maximum value is 30 days ago. - The value you provide should follow the ATOM date-time format: &#x60;2024-02-05T00:00:00+01:00&#x60;  | 
**to** | **time.Time** | Use this query parameter to define the ending date-time of the period you want analytics for.  - If you do not set a value for &#x60;to&#x60;, the default assigned value is &#x60;now&#x60;. - The value for &#x60;to&#x60; is a non-inclusive value: the API returns data **before** the date-time that you set.  | 
**filterBy** | [**FilterBy2**](FilterBy2.md) | Use this parameter to filter the API&#39;s response based on different data dimensions. You can serialize filters in your query to receive more detailed breakdowns of your analytics.  - If you do not set a value for &#x60;filterBy&#x60;, the API returns the full dataset for your project. - The API only accepts the &#x60;mediaId&#x60; and &#x60;mediaType&#x60; filters when you call &#x60;/data/metrics/play/total&#x60; or &#x60;/data/buckets/play-total/media-id&#x60;.  These are the available breakdown dimensions:  - &#x60;mediaId&#x60;: Returns analytics based on the unique identifiers of a video or a live stream. - &#x60;mediaType&#x60;: Returns analytics based on the type of content. Possible values: &#x60;video&#x60; and &#x60;live-stream&#x60;.  - &#x60;continent&#x60;: Returns analytics based on the viewers&#39; continent. The list of supported continents names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example &#x60;EU&#x60;. Possible values are: &#x60;AS&#x60;, &#x60;AF&#x60;, &#x60;NA&#x60;, &#x60;SA&#x60;, &#x60;AN&#x60;, &#x60;EU&#x60;, &#x60;AZ&#x60;.  - &#x60;country&#x60;: Returns analytics based on the viewers&#39; country. The list of supported country names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example &#x60;FR&#x60;. - &#x60;deviceType&#x60;: Returns analytics based on the type of device used by the viewers. Response values can include: &#x60;computer&#x60;, &#x60;phone&#x60;, &#x60;tablet&#x60;, &#x60;tv&#x60;, &#x60;console&#x60;, &#x60;wearable&#x60;, &#x60;unknown&#x60;. - &#x60;operatingSystem&#x60;: Returns analytics based on the operating system used by the viewers. Response values can include &#x60;windows&#x60;, &#x60;mac osx&#x60;, &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;linux&#x60;. - &#x60;browser&#x60;: Returns analytics based on the browser used by the viewers. Response values can include &#x60;chrome&#x60;, &#x60;firefox&#x60;, &#x60;edge&#x60;, &#x60;opera&#x60;. - &#x60;tag&#x60;: Returns analytics for videos using this tag. This filter only accepts a single value and is case sensitive. Read more about tagging your videos [here](https://docs.api.video/vod/tags-metadata).  | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**AnalyticsMetricsBreakdownResponse**](AnalyticsMetricsBreakdownResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsOverTime

> GetMetricsOverTime(metric string, r AnalyticsApiGetMetricsOverTimeRequest) (*AnalyticsMetricsOverTimeResponse, error)


> GetMetricsOverTimeWithContext(ctx context.Context, metric string, r AnalyticsApiGetMetricsOverTimeRequest) (*AnalyticsMetricsOverTimeResponse, error)



Retrieve metrics over time



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
    client := apivideosdk.ClientBuilder("YOUR_API_KEY").Build()
    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_KEY").Build()
    req := apivideosdk.AnalyticsApiGetMetricsOverTimeRequest{}
    
    req.Metric("metric_example") // string | Use this path parameter to select a metric that you want analytics for.  - `play` is the number of times your content has been played. - `play-rate` is the ratio that calculates the number of plays your content receives divided by its impressions. - `start` is the number of times playback was started. - `end` is the number of times playback has ended with the content watch until the end. - `impression` is the number of times your content has been loaded and was ready for playback. 
    req.From(time.Now()) // time.Time | Use this query parameter to define the starting date-time of the period you want analytics for.  - If you do not set a value for `from`, the default assigned value is 1 day ago, based on the `to` parameter. - The maximum value is 30 days ago. - The value you provide should follow the ATOM date-time format: `2024-02-05T00:00:00+01:00` 
    req.To(time.Now()) // time.Time | Use this query parameter to define the ending date-time of the period you want analytics for.  - If you do not set a value for `to`, the default assigned value is `now`. - The value for `to` is a non-inclusive value: the API returns data **before** the date-time that you set. 
    req.Interval("hour") // string | Use this query parameter to define how granularity of the data. Possible values: `hour`, `day`.  - Default: If no interval specified and the period (different between from and to) ≤ 2 days then hour, otherwise day.  - If you do not set a value for `interval`, and the period you set using the `from` and `to` parameters is less than or equals to 2 days, then the default assigned value is `hour`. Otherwise the API sets it to `day`. 
    req.FilterBy(map[string][]apivideosdk.FilterBy2{"key": "TODO"}) // FilterBy2 | Use this parameter to filter the API's response based on different data dimensions. You can serialize filters in your query to receive more detailed breakdowns of your analytics.  - If you do not set a value for `filterBy`, the API returns the full dataset for your project. - The API only accepts the `mediaId` and `mediaType` filters when you call `/data/metrics/play/total` or `/data/buckets/play-total/media-id`.  These are the available breakdown dimensions:  - `mediaId`: Returns analytics based on the unique identifiers of a video or a live stream. - `mediaType`: Returns analytics based on the type of content. Possible values: `video` and `live-stream`.  - `continent`: Returns analytics based on the viewers' continent. The list of supported continents names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example `EU`. Possible values are: `AS`, `AF`, `NA`, `SA`, `AN`, `EU`, `AZ`.  - `country`: Returns analytics based on the viewers' country. The list of supported country names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example `FR`. - `deviceType`: Returns analytics based on the type of device used by the viewers. Response values can include: `computer`, `phone`, `tablet`, `tv`, `console`, `wearable`, `unknown`. - `operatingSystem`: Returns analytics based on the operating system used by the viewers. Response values can include `windows`, `mac osx`, `android`, `ios`, `linux`. - `browser`: Returns analytics based on the browser used by the viewers. Response values can include `chrome`, `firefox`, `edge`, `opera`. - `tag`: Returns analytics for videos using this tag. This filter only accepts a single value and is case sensitive. Read more about tagging your videos [here](https://docs.api.video/vod/tags-metadata). 
    req.CurrentPage(int32(2)) // int32 | Choose the number of search results to return per page. Minimum value: 1 (default to 1)
    req.PageSize(int32(30)) // int32 | Results per page. Allowed values 1-100, default is 25. (default to 25)

    res, err := client.Analytics.GetMetricsOverTime(metric string, req)
    

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Analytics.GetMetricsOverTime``: %v\n", err)
    }
    // response from `GetMetricsOverTime`: AnalyticsMetricsOverTimeResponse
    fmt.Fprintf(os.Stdout, "Response from `Analytics.GetMetricsOverTime`: %v\n", res)
}
```
### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**metric** | **string** | Use this path parameter to select a metric that you want analytics for.  - &#x60;play&#x60; is the number of times your content has been played. - &#x60;play-rate&#x60; is the ratio that calculates the number of plays your content receives divided by its impressions. - &#x60;start&#x60; is the number of times playback was started. - &#x60;end&#x60; is the number of times playback has ended with the content watch until the end. - &#x60;impression&#x60; is the number of times your content has been loaded and was ready for playback.  | 

### Other Parameters



Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**from** | **time.Time** | Use this query parameter to define the starting date-time of the period you want analytics for.  - If you do not set a value for &#x60;from&#x60;, the default assigned value is 1 day ago, based on the &#x60;to&#x60; parameter. - The maximum value is 30 days ago. - The value you provide should follow the ATOM date-time format: &#x60;2024-02-05T00:00:00+01:00&#x60;  | 
**to** | **time.Time** | Use this query parameter to define the ending date-time of the period you want analytics for.  - If you do not set a value for &#x60;to&#x60;, the default assigned value is &#x60;now&#x60;. - The value for &#x60;to&#x60; is a non-inclusive value: the API returns data **before** the date-time that you set.  | 
**interval** | **string** | Use this query parameter to define how granularity of the data. Possible values: &#x60;hour&#x60;, &#x60;day&#x60;.  - Default: If no interval specified and the period (different between from and to) ≤ 2 days then hour, otherwise day.  - If you do not set a value for &#x60;interval&#x60;, and the period you set using the &#x60;from&#x60; and &#x60;to&#x60; parameters is less than or equals to 2 days, then the default assigned value is &#x60;hour&#x60;. Otherwise the API sets it to &#x60;day&#x60;.  | 
**filterBy** | [**FilterBy2**](FilterBy2.md) | Use this parameter to filter the API&#39;s response based on different data dimensions. You can serialize filters in your query to receive more detailed breakdowns of your analytics.  - If you do not set a value for &#x60;filterBy&#x60;, the API returns the full dataset for your project. - The API only accepts the &#x60;mediaId&#x60; and &#x60;mediaType&#x60; filters when you call &#x60;/data/metrics/play/total&#x60; or &#x60;/data/buckets/play-total/media-id&#x60;.  These are the available breakdown dimensions:  - &#x60;mediaId&#x60;: Returns analytics based on the unique identifiers of a video or a live stream. - &#x60;mediaType&#x60;: Returns analytics based on the type of content. Possible values: &#x60;video&#x60; and &#x60;live-stream&#x60;.  - &#x60;continent&#x60;: Returns analytics based on the viewers&#39; continent. The list of supported continents names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example &#x60;EU&#x60;. Possible values are: &#x60;AS&#x60;, &#x60;AF&#x60;, &#x60;NA&#x60;, &#x60;SA&#x60;, &#x60;AN&#x60;, &#x60;EU&#x60;, &#x60;AZ&#x60;.  - &#x60;country&#x60;: Returns analytics based on the viewers&#39; country. The list of supported country names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example &#x60;FR&#x60;. - &#x60;deviceType&#x60;: Returns analytics based on the type of device used by the viewers. Response values can include: &#x60;computer&#x60;, &#x60;phone&#x60;, &#x60;tablet&#x60;, &#x60;tv&#x60;, &#x60;console&#x60;, &#x60;wearable&#x60;, &#x60;unknown&#x60;. - &#x60;operatingSystem&#x60;: Returns analytics based on the operating system used by the viewers. Response values can include &#x60;windows&#x60;, &#x60;mac osx&#x60;, &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;linux&#x60;. - &#x60;browser&#x60;: Returns analytics based on the browser used by the viewers. Response values can include &#x60;chrome&#x60;, &#x60;firefox&#x60;, &#x60;edge&#x60;, &#x60;opera&#x60;. - &#x60;tag&#x60;: Returns analytics for videos using this tag. This filter only accepts a single value and is case sensitive. Read more about tagging your videos [here](https://docs.api.video/vod/tags-metadata).  | 
**currentPage** | **int32** | Choose the number of search results to return per page. Minimum value: 1 | [default to 1]
**pageSize** | **int32** | Results per page. Allowed values 1-100, default is 25. | [default to 25]

### Return type

[**AnalyticsMetricsOverTimeResponse**](AnalyticsMetricsOverTimeResponse.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

