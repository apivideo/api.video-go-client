<!--<documentation_excluded>-->
[![badge](https://img.shields.io/twitter/follow/api_video?style=social)](https://twitter.com/intent/follow?screen_name=api_video) &nbsp; [![badge](https://img.shields.io/github/stars/apivideo/api.video-go-client?style=social)](https://github.com/apivideo/api.video-go-client) &nbsp; [![badge](https://img.shields.io/discourse/topics?server=https%3A%2F%2Fcommunity.api.video)](https://community.api.video)
![](https://github.com/apivideo/.github/blob/main/assets/apivideo_banner.png)
<h1 align="center">api.video Go client</h1>

[api.video](https://api.video) is the video infrastructure for product builders. Lightning fast video APIs for integrating, scaling, and managing on-demand & low latency live streaming features in your app.

## Table of contents

- [Project description](#project-description)
- [Getting started](#getting-started)
  - [Installation](#installation)
  - [Code sample](#code-sample)
- [Documentation](#documentation)
  - [API Endpoints](#api-endpoints)
    - [Analytics](#analytics)
    - [Captions](#captions)
    - [Chapters](#chapters)
    - [LiveStreams](#livestreams)
    - [PlayerThemes](#playerthemes)
    - [UploadTokens](#uploadtokens)
    - [Videos](#videos)
    - [Watermarks](#watermarks)
    - [Webhooks](#webhooks)
  - [Models](#models)
- [Have you gotten use from this API client?](#have-you-gotten-use-from-this-api-client-)
- [Contribution](#contribution)
<!--</documentation_excluded>-->
<!--<documentation_only>
---
title: Go API client
meta: 
  description: The official Go API client for api.video. [api.video](https://api.video/) is the video infrastructure for product builders. Lightning fast video APIs for integrating, scaling, and managing on-demand & low latency live streaming features in your app.
---

# api.video Go API client

[api.video](https://api.video/) is the video infrastructure for product builders. Lightning fast video APIs for integrating, scaling, and managing on-demand & low latency live streaming features in your app.
</documentation_only>-->

## Project description

api.video's Go client streamlines the coding process. Chunking files is handled for you, as is pagination and refreshing your tokens.

## Getting started

### Installation
```bash
go get github.com/apivideo/api.video-go-client
```


### Code sample

For a more advanced usage you can checkout the rest of the documentation in the [docs directory](https://github.com/apivideo/api.video-go-client/blob/main/docs)

```golang
package main

import (
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
    //Connect to production environment
    client := apivideosdk.ClientBuilder("YOUR_API_KEY").Build()

    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_KEY").Build()


    //List Videos
    //First create the url options for searching
    opts := apivideosdk.VideosApiListRequest{}.
        CurrentPage(1).
        PageSize(25).
        SortBy("publishedAt").
        SortOrder("desc")

    //Then call the List endpoint with the options
    result, err := client.Videos.List(opts)

    if err != nil {
        fmt.Println(err)
    }

    for _, video := range result.Data {
        fmt.Printf("%s\n", video.VideoId)
        fmt.Printf("%s\n", *video.Title)
    }


    //Upload a video
    //First create a container
    create, err := client.Videos.Create(apivideosdk.VideoCreationPayload{Title: "My video title"})

    if err != nil {
        fmt.Println(err)
    }

    //Then open the video file
    videoFile, err := os.Open("path/to/video.mp4")

    if err != nil {
        fmt.Println(err)
    }

    //Finally upload your video to the container with the videoId
    uploadedVideo, err := client.Videos.UploadFile(create.VideoId, videoFile)

    if err != nil {
        fmt.Println(err)
    }


    //And get the assets
    fmt.Printf("%s\n", *uploadedVideo.Assets.Hls)
    fmt.Printf("%s\n", *uploadedVideo.Assets.Iframe)
}
```

## Documentation

### API Endpoints

All URIs are relative to *https://ws.api.video*


#### Analytics


##### Retrieve an instance of the Analytics API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
analyticsApi := client.Analytics
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAggregatedMetrics**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Analytics.md#GetAggregatedMetrics) | **Get** `/data/metrics/{metric}/{aggregation}` | Retrieve aggregated metrics
[**GetMetricsBreakdown**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Analytics.md#GetMetricsBreakdown) | **Get** `/data/buckets/{metric}/{breakdown}` | Retrieve metrics in a breakdown of dimensions
[**GetMetricsOverTime**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Analytics.md#GetMetricsOverTime) | **Get** `/data/timeseries/{metric}` | Retrieve metrics over time


#### Captions


##### Retrieve an instance of the Captions API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
captionsApi := client.Captions
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Captions.md#Upload) | **Post** `/videos/{videoId}/captions/{language}` | Upload a caption
[**Get**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Captions.md#Get) | **Get** `/videos/{videoId}/captions/{language}` | Retrieve a caption
[**Update**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Captions.md#Update) | **Patch** `/videos/{videoId}/captions/{language}` | Update a caption
[**Delete**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Captions.md#Delete) | **Delete** `/videos/{videoId}/captions/{language}` | Delete a caption
[**List**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Captions.md#List) | **Get** `/videos/{videoId}/captions` | List video captions


#### Chapters


##### Retrieve an instance of the Chapters API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
chaptersApi := client.Chapters
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Chapters.md#Upload) | **Post** `/videos/{videoId}/chapters/{language}` | Upload a chapter
[**Get**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Chapters.md#Get) | **Get** `/videos/{videoId}/chapters/{language}` | Retrieve a chapter
[**Delete**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Chapters.md#Delete) | **Delete** `/videos/{videoId}/chapters/{language}` | Delete a chapter
[**List**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Chapters.md#List) | **Get** `/videos/{videoId}/chapters` | List video chapters


#### LiveStreams


##### Retrieve an instance of the LiveStreams API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
liveStreamsApi := client.LiveStreams
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreams.md#Create) | **Post** `/live-streams` | Create live stream
[**Get**](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreams.md#Get) | **Get** `/live-streams/{liveStreamId}` | Retrieve live stream
[**Update**](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreams.md#Update) | **Patch** `/live-streams/{liveStreamId}` | Update a live stream
[**Delete**](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreams.md#Delete) | **Delete** `/live-streams/{liveStreamId}` | Delete a live stream
[**List**](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreams.md#List) | **Get** `/live-streams` | List all live streams
[**UploadThumbnail**](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreams.md#UploadThumbnail) | **Post** `/live-streams/{liveStreamId}/thumbnail` | Upload a thumbnail
[**DeleteThumbnail**](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreams.md#DeleteThumbnail) | **Delete** `/live-streams/{liveStreamId}/thumbnail` | Delete a thumbnail
[**Complete**](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreams.md#Complete) | **Put** `/live-streams/{liveStreamId}/complete` | Complete a live stream


#### PlayerThemes


##### Retrieve an instance of the PlayerThemes API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
playerThemesApi := client.PlayerThemes
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemes.md#Create) | **Post** `/players` | Create a player
[**Get**](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemes.md#Get) | **Get** `/players/{playerId}` | Retrieve a player
[**Update**](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemes.md#Update) | **Patch** `/players/{playerId}` | Update a player
[**Delete**](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemes.md#Delete) | **Delete** `/players/{playerId}` | Delete a player
[**List**](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemes.md#List) | **Get** `/players` | List all player themes
[**UploadLogo**](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemes.md#UploadLogo) | **Post** `/players/{playerId}/logo` | Upload a logo
[**DeleteLogo**](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemes.md#DeleteLogo) | **Delete** `/players/{playerId}/logo` | Delete logo


#### UploadTokens


##### Retrieve an instance of the UploadTokens API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
uploadTokensApi := client.UploadTokens
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](https://github.com/apivideo/api.video-go-client/blob/main/docs/UploadTokens.md#CreateToken) | **Post** `/upload-tokens` | Generate an upload token
[**GetToken**](https://github.com/apivideo/api.video-go-client/blob/main/docs/UploadTokens.md#GetToken) | **Get** `/upload-tokens/{uploadToken}` | Retrieve upload token
[**DeleteToken**](https://github.com/apivideo/api.video-go-client/blob/main/docs/UploadTokens.md#DeleteToken) | **Delete** `/upload-tokens/{uploadToken}` | Delete an upload token
[**List**](https://github.com/apivideo/api.video-go-client/blob/main/docs/UploadTokens.md#List) | **Get** `/upload-tokens` | List all active upload tokens


#### Videos


##### Retrieve an instance of the Videos API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
videosApi := client.Videos
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#Create) | **Post** `/videos` | Create a video object
[**Upload**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#Upload) | **Post** `/videos/{videoId}/source` | Upload a video
[**UploadWithUploadToken**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#UploadWithUploadToken) | **Post** `/upload` | Upload with an delegated upload token
[**Get**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#Get) | **Get** `/videos/{videoId}` | Retrieve a video object
[**Update**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#Update) | **Patch** `/videos/{videoId}` | Update a video object
[**Delete**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#Delete) | **Delete** `/videos/{videoId}` | Delete a video object
[**List**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#List) | **Get** `/videos` | List all video objects
[**UploadThumbnail**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#UploadThumbnail) | **Post** `/videos/{videoId}/thumbnail` | Upload a thumbnail
[**PickThumbnail**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#PickThumbnail) | **Patch** `/videos/{videoId}/thumbnail` | Set a thumbnail
[**GetDiscarded**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#GetDiscarded) | **Get** `/discarded/videos/{videoId}` | Retrieve a discarded video object
[**GetStatus**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#GetStatus) | **Get** `/videos/{videoId}/status` | Retrieve video status and details
[**ListDiscarded**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#ListDiscarded) | **Get** `/discarded/videos` | List all discarded video objects
[**UpdateDiscarded**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Videos.md#UpdateDiscarded) | **Patch** `/discarded/videos/{videoId}` | Update a discarded video object


#### Watermarks


##### Retrieve an instance of the Watermarks API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
watermarksApi := client.Watermarks
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Watermarks.md#Upload) | **Post** `/watermarks` | Upload a watermark
[**Delete**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Watermarks.md#Delete) | **Delete** `/watermarks/{watermarkId}` | Delete a watermark
[**List**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Watermarks.md#List) | **Get** `/watermarks` | List all watermarks


#### Webhooks


##### Retrieve an instance of the Webhooks API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
webhooksApi := client.Webhooks
```

##### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Webhooks.md#Create) | **Post** `/webhooks` | Create Webhook
[**Get**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Webhooks.md#Get) | **Get** `/webhooks/{webhookId}` | Retrieve Webhook details
[**Delete**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Webhooks.md#Delete) | **Delete** `/webhooks/{webhookId}` | Delete a Webhook
[**List**](https://github.com/apivideo/api.video-go-client/blob/main/docs/Webhooks.md#List) | **Get** `/webhooks` | List all webhooks




### Models

 - [AccessToken](https://github.com/apivideo/api.video-go-client/blob/main/docs/AccessToken.md)
 - [AdditionalBadRequestErrors](https://github.com/apivideo/api.video-go-client/blob/main/docs/AdditionalBadRequestErrors.md)
 - [AnalyticsAggregatedMetricsResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsAggregatedMetricsResponse.md)
 - [AnalyticsAggregatedMetricsResponseContext](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsAggregatedMetricsResponseContext.md)
 - [AnalyticsAggregatedMetricsResponseContextTimeframe](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsAggregatedMetricsResponseContextTimeframe.md)
 - [AnalyticsData](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsData.md)
 - [AnalyticsMetricsBreakdownResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsMetricsBreakdownResponse.md)
 - [AnalyticsMetricsBreakdownResponseContext](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsMetricsBreakdownResponseContext.md)
 - [AnalyticsMetricsBreakdownResponseData](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsMetricsBreakdownResponseData.md)
 - [AnalyticsMetricsOverTimeResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsMetricsOverTimeResponse.md)
 - [AnalyticsMetricsOverTimeResponseContext](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsMetricsOverTimeResponseContext.md)
 - [AnalyticsMetricsOverTimeResponseData](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsMetricsOverTimeResponseData.md)
 - [AnalyticsPlays400Error](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsPlays400Error.md)
 - [AnalyticsPlaysResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/AnalyticsPlaysResponse.md)
 - [AuthenticatePayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/AuthenticatePayload.md)
 - [BadRequest](https://github.com/apivideo/api.video-go-client/blob/main/docs/BadRequest.md)
 - [BytesRange](https://github.com/apivideo/api.video-go-client/blob/main/docs/BytesRange.md)
 - [Caption](https://github.com/apivideo/api.video-go-client/blob/main/docs/Caption.md)
 - [CaptionsListResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/CaptionsListResponse.md)
 - [CaptionsUpdatePayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/CaptionsUpdatePayload.md)
 - [Chapter](https://github.com/apivideo/api.video-go-client/blob/main/docs/Chapter.md)
 - [ChaptersListResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/ChaptersListResponse.md)
 - [DiscardedVideoUpdatePayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/DiscardedVideoUpdatePayload.md)
 - [FilterBy](https://github.com/apivideo/api.video-go-client/blob/main/docs/FilterBy.md)
 - [FilterBy1](https://github.com/apivideo/api.video-go-client/blob/main/docs/FilterBy1.md)
 - [FilterBy2](https://github.com/apivideo/api.video-go-client/blob/main/docs/FilterBy2.md)
 - [Link](https://github.com/apivideo/api.video-go-client/blob/main/docs/Link.md)
 - [LiveStream](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStream.md)
 - [LiveStreamAssets](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreamAssets.md)
 - [LiveStreamCreationPayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreamCreationPayload.md)
 - [LiveStreamListResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreamListResponse.md)
 - [LiveStreamUpdatePayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/LiveStreamUpdatePayload.md)
 - [Metadata](https://github.com/apivideo/api.video-go-client/blob/main/docs/Metadata.md)
 - [Model403ErrorSchema](https://github.com/apivideo/api.video-go-client/blob/main/docs/Model403ErrorSchema.md)
 - [NotFound](https://github.com/apivideo/api.video-go-client/blob/main/docs/NotFound.md)
 - [Pagination](https://github.com/apivideo/api.video-go-client/blob/main/docs/Pagination.md)
 - [PaginationLink](https://github.com/apivideo/api.video-go-client/blob/main/docs/PaginationLink.md)
 - [PlayerSessionEvent](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerSessionEvent.md)
 - [PlayerTheme](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerTheme.md)
 - [PlayerThemeAssets](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemeAssets.md)
 - [PlayerThemeCreationPayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemeCreationPayload.md)
 - [PlayerThemeUpdatePayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemeUpdatePayload.md)
 - [PlayerThemesListResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/PlayerThemesListResponse.md)
 - [Quality](https://github.com/apivideo/api.video-go-client/blob/main/docs/Quality.md)
 - [RefreshTokenPayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/RefreshTokenPayload.md)
 - [RestreamsRequestObject](https://github.com/apivideo/api.video-go-client/blob/main/docs/RestreamsRequestObject.md)
 - [RestreamsResponseObject](https://github.com/apivideo/api.video-go-client/blob/main/docs/RestreamsResponseObject.md)
 - [TokenCreationPayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/TokenCreationPayload.md)
 - [TokenListResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/TokenListResponse.md)
 - [TooManyRequests](https://github.com/apivideo/api.video-go-client/blob/main/docs/TooManyRequests.md)
 - [UnrecognizedRequestUrl](https://github.com/apivideo/api.video-go-client/blob/main/docs/UnrecognizedRequestUrl.md)
 - [UploadToken](https://github.com/apivideo/api.video-go-client/blob/main/docs/UploadToken.md)
 - [Video](https://github.com/apivideo/api.video-go-client/blob/main/docs/Video.md)
 - [VideoAssets](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoAssets.md)
 - [VideoClip](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoClip.md)
 - [VideoCreationPayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoCreationPayload.md)
 - [VideoSource](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoSource.md)
 - [VideoSourceLiveStream](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoSourceLiveStream.md)
 - [VideoSourceLiveStreamLink](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoSourceLiveStreamLink.md)
 - [VideoStatus](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoStatus.md)
 - [VideoStatusEncoding](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoStatusEncoding.md)
 - [VideoStatusEncodingMetadata](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoStatusEncodingMetadata.md)
 - [VideoStatusIngest](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoStatusIngest.md)
 - [VideoStatusIngestReceivedParts](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoStatusIngestReceivedParts.md)
 - [VideoThumbnailPickPayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoThumbnailPickPayload.md)
 - [VideoUpdatePayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoUpdatePayload.md)
 - [VideoWatermark](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideoWatermark.md)
 - [VideosListResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/VideosListResponse.md)
 - [Watermark](https://github.com/apivideo/api.video-go-client/blob/main/docs/Watermark.md)
 - [WatermarksListResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/WatermarksListResponse.md)
 - [Webhook](https://github.com/apivideo/api.video-go-client/blob/main/docs/Webhook.md)
 - [WebhooksCreationPayload](https://github.com/apivideo/api.video-go-client/blob/main/docs/WebhooksCreationPayload.md)
 - [WebhooksListResponse](https://github.com/apivideo/api.video-go-client/blob/main/docs/WebhooksListResponse.md)



## Have you gotten use from this API client?

Please take a moment to leave a star on the client ⭐

This helps other users to find the clients and also helps us understand which clients are most popular. Thank you!

## Contribution

Since this API client is generated from an OpenAPI description, we cannot accept pull requests made directly to the repository. If you want to contribute, you can open a pull request on the repository of our [client generator](https://github.com/apivideo/api-client-generator). Otherwise, you can also simply open an issue detailing your need on this repository.