[![badge](https://img.shields.io/twitter/follow/api_video?style=social)](https://twitter.com/intent/follow?screen_name=api_video) &nbsp; [![badge](https://img.shields.io/github/stars/apivideo/api.video-go-client?style=social)](https://github.com/apivideo/api.video-go-client) &nbsp; [![badge](https://img.shields.io/discourse/topics?server=https%3A%2F%2Fcommunity.api.video)](https://community.api.video)
![](https://github.com/apivideo/.github/blob/main/assets/apivideo_banner.png)
<h1 align="center">api.video Go client</h1>

[api.video](https://api.video) is the video infrastructure for product builders. Lightning fast video APIs for integrating, scaling, and managing on-demand & low latency live streaming features in your app.

# Table of contents

- [Project description](#project-description)
- [Getting started](#getting-started)
  - [Installation](#installation)
  - [Code sample](#code-sample)
- [Documentation](#documentation)
  - [API Endpoints](#api-endpoints)
    - [Captions](#captions)
    - [Chapters](#chapters)
    - [LiveStreams](#livestreams)
    - [PlayerThemes](#playerthemes)
    - [RawStatistics](#rawstatistics)
    - [UploadTokens](#uploadtokens)
    - [Videos](#videos)
    - [Watermarks](#watermarks)
    - [Webhooks](#webhooks)
  - [Models](#models)
- [Have you gotten use from this API client?](#have-you-gotten-use-from-this-api-client-)
- [Contribution](#contribution)

# Project description

api.video's Java Go client streamlines the coding process. Chunking files is handled for you, as is pagination and refreshing your tokens.

# Getting started

## Installation
```bash
go get github.com/apivideo/api.video-go-client
```


## Code sample

For a more advanced usage you can checkout the rest of the documentation in the [docs directory](/docs)

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

# Documentation

## API Endpoints

All URIs are relative to *https://ws.api.video*


### Captions


#### Retrieve an instance of the Captions API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
captionsApi := client.Captions
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](docs/Captions.md#Upload) | **Post** /videos/{videoId}/captions/{language} | Upload a caption
[**Get**](docs/Captions.md#Get) | **Get** /videos/{videoId}/captions/{language} | Retrieve a caption
[**Update**](docs/Captions.md#Update) | **Patch** /videos/{videoId}/captions/{language} | Update a caption
[**Delete**](docs/Captions.md#Delete) | **Delete** /videos/{videoId}/captions/{language} | Delete a caption
[**List**](docs/Captions.md#List) | **Get** /videos/{videoId}/captions | List video captions


### Chapters


#### Retrieve an instance of the Chapters API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
chaptersApi := client.Chapters
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](docs/Chapters.md#Upload) | **Post** /videos/{videoId}/chapters/{language} | Upload a chapter
[**Get**](docs/Chapters.md#Get) | **Get** /videos/{videoId}/chapters/{language} | Retrieve a chapter
[**Delete**](docs/Chapters.md#Delete) | **Delete** /videos/{videoId}/chapters/{language} | Delete a chapter
[**List**](docs/Chapters.md#List) | **Get** /videos/{videoId}/chapters | List video chapters


### LiveStreams


#### Retrieve an instance of the LiveStreams API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
liveStreamsApi := client.LiveStreams
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](docs/LiveStreams.md#Create) | **Post** /live-streams | Create live stream
[**Get**](docs/LiveStreams.md#Get) | **Get** /live-streams/{liveStreamId} | Retrieve live stream
[**Update**](docs/LiveStreams.md#Update) | **Patch** /live-streams/{liveStreamId} | Update a live stream
[**Delete**](docs/LiveStreams.md#Delete) | **Delete** /live-streams/{liveStreamId} | Delete a live stream
[**List**](docs/LiveStreams.md#List) | **Get** /live-streams | List all live streams
[**UploadThumbnail**](docs/LiveStreams.md#UploadThumbnail) | **Post** /live-streams/{liveStreamId}/thumbnail | Upload a thumbnail
[**DeleteThumbnail**](docs/LiveStreams.md#DeleteThumbnail) | **Delete** /live-streams/{liveStreamId}/thumbnail | Delete a thumbnail


### PlayerThemes


#### Retrieve an instance of the PlayerThemes API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
playerThemesApi := client.PlayerThemes
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](docs/PlayerThemes.md#Create) | **Post** /players | Create a player
[**Get**](docs/PlayerThemes.md#Get) | **Get** /players/{playerId} | Retrieve a player
[**Update**](docs/PlayerThemes.md#Update) | **Patch** /players/{playerId} | Update a player
[**Delete**](docs/PlayerThemes.md#Delete) | **Delete** /players/{playerId} | Delete a player
[**List**](docs/PlayerThemes.md#List) | **Get** /players | List all player themes
[**UploadLogo**](docs/PlayerThemes.md#UploadLogo) | **Post** /players/{playerId}/logo | Upload a logo
[**DeleteLogo**](docs/PlayerThemes.md#DeleteLogo) | **Delete** /players/{playerId}/logo | Delete logo


### RawStatistics


#### Retrieve an instance of the RawStatistics API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
rawStatisticsApi := client.RawStatistics
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLiveStreamSessions**](docs/RawStatistics.md#ListLiveStreamSessions) | **Get** /analytics/live-streams/{liveStreamId} | List live stream player sessions
[**ListSessionEvents**](docs/RawStatistics.md#ListSessionEvents) | **Get** /analytics/sessions/{sessionId}/events | List player session events
[**ListVideoSessions**](docs/RawStatistics.md#ListVideoSessions) | **Get** /analytics/videos/{videoId} | List video player sessions


### UploadTokens


#### Retrieve an instance of the UploadTokens API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
uploadTokensApi := client.UploadTokens
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](docs/UploadTokens.md#CreateToken) | **Post** /upload-tokens | Generate an upload token
[**GetToken**](docs/UploadTokens.md#GetToken) | **Get** /upload-tokens/{uploadToken} | Retrieve upload token
[**DeleteToken**](docs/UploadTokens.md#DeleteToken) | **Delete** /upload-tokens/{uploadToken} | Delete an upload token
[**List**](docs/UploadTokens.md#List) | **Get** /upload-tokens | List all active upload tokens


### Videos


#### Retrieve an instance of the Videos API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
videosApi := client.Videos
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](docs/Videos.md#Create) | **Post** /videos | Create a video object
[**Upload**](docs/Videos.md#Upload) | **Post** /videos/{videoId}/source | Upload a video
[**UploadWithUploadToken**](docs/Videos.md#UploadWithUploadToken) | **Post** /upload | Upload with an delegated upload token
[**Get**](docs/Videos.md#Get) | **Get** /videos/{videoId} | Retrieve a video object
[**Update**](docs/Videos.md#Update) | **Patch** /videos/{videoId} | Update a video object
[**Delete**](docs/Videos.md#Delete) | **Delete** /videos/{videoId} | Delete a video object
[**List**](docs/Videos.md#List) | **Get** /videos | List all video objects
[**UploadThumbnail**](docs/Videos.md#UploadThumbnail) | **Post** /videos/{videoId}/thumbnail | Upload a thumbnail
[**PickThumbnail**](docs/Videos.md#PickThumbnail) | **Patch** /videos/{videoId}/thumbnail | Set a thumbnail
[**GetStatus**](docs/Videos.md#GetStatus) | **Get** /videos/{videoId}/status | Retrieve video status and details


### Watermarks


#### Retrieve an instance of the Watermarks API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
watermarksApi := client.Watermarks
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](docs/Watermarks.md#Upload) | **Post** /watermarks | Upload a watermark
[**Delete**](docs/Watermarks.md#Delete) | **Delete** /watermarks/{watermarkId} | Delete a watermark
[**List**](docs/Watermarks.md#List) | **Get** /watermarks | List all watermarks


### Webhooks


#### Retrieve an instance of the Webhooks API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
webhooksApi := client.Webhooks
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](docs/Webhooks.md#Create) | **Post** /webhooks | Create Webhook
[**Get**](docs/Webhooks.md#Get) | **Get** /webhooks/{webhookId} | Retrieve Webhook details
[**Delete**](docs/Webhooks.md#Delete) | **Delete** /webhooks/{webhookId} | Delete a Webhook
[**List**](docs/Webhooks.md#List) | **Get** /webhooks | List all webhooks




## Models

 - [AccessToken](docs/AccessToken.md)
 - [AuthenticatePayload](docs/AuthenticatePayload.md)
 - [BadRequest](docs/BadRequest.md)
 - [BytesRange](docs/BytesRange.md)
 - [Caption](docs/Caption.md)
 - [CaptionsListResponse](docs/CaptionsListResponse.md)
 - [CaptionsUpdatePayload](docs/CaptionsUpdatePayload.md)
 - [Chapter](docs/Chapter.md)
 - [ChaptersListResponse](docs/ChaptersListResponse.md)
 - [Link](docs/Link.md)
 - [LiveStream](docs/LiveStream.md)
 - [LiveStreamAssets](docs/LiveStreamAssets.md)
 - [LiveStreamCreationPayload](docs/LiveStreamCreationPayload.md)
 - [LiveStreamListResponse](docs/LiveStreamListResponse.md)
 - [LiveStreamSession](docs/LiveStreamSession.md)
 - [LiveStreamSessionClient](docs/LiveStreamSessionClient.md)
 - [LiveStreamSessionDevice](docs/LiveStreamSessionDevice.md)
 - [LiveStreamSessionLocation](docs/LiveStreamSessionLocation.md)
 - [LiveStreamSessionReferrer](docs/LiveStreamSessionReferrer.md)
 - [LiveStreamSessionSession](docs/LiveStreamSessionSession.md)
 - [LiveStreamUpdatePayload](docs/LiveStreamUpdatePayload.md)
 - [Metadata](docs/Metadata.md)
 - [NotFound](docs/NotFound.md)
 - [Pagination](docs/Pagination.md)
 - [PaginationLink](docs/PaginationLink.md)
 - [PlayerSessionEvent](docs/PlayerSessionEvent.md)
 - [PlayerTheme](docs/PlayerTheme.md)
 - [PlayerThemeAssets](docs/PlayerThemeAssets.md)
 - [PlayerThemeCreationPayload](docs/PlayerThemeCreationPayload.md)
 - [PlayerThemeUpdatePayload](docs/PlayerThemeUpdatePayload.md)
 - [PlayerThemesListResponse](docs/PlayerThemesListResponse.md)
 - [Quality](docs/Quality.md)
 - [RawStatisticsListLiveStreamAnalyticsResponse](docs/RawStatisticsListLiveStreamAnalyticsResponse.md)
 - [RawStatisticsListPlayerSessionEventsResponse](docs/RawStatisticsListPlayerSessionEventsResponse.md)
 - [RawStatisticsListSessionsResponse](docs/RawStatisticsListSessionsResponse.md)
 - [RefreshTokenPayload](docs/RefreshTokenPayload.md)
 - [TokenCreationPayload](docs/TokenCreationPayload.md)
 - [TokenListResponse](docs/TokenListResponse.md)
 - [UploadToken](docs/UploadToken.md)
 - [Video](docs/Video.md)
 - [VideoAssets](docs/VideoAssets.md)
 - [VideoClip](docs/VideoClip.md)
 - [VideoCreationPayload](docs/VideoCreationPayload.md)
 - [VideoSession](docs/VideoSession.md)
 - [VideoSessionClient](docs/VideoSessionClient.md)
 - [VideoSessionDevice](docs/VideoSessionDevice.md)
 - [VideoSessionLocation](docs/VideoSessionLocation.md)
 - [VideoSessionOs](docs/VideoSessionOs.md)
 - [VideoSessionReferrer](docs/VideoSessionReferrer.md)
 - [VideoSessionSession](docs/VideoSessionSession.md)
 - [VideoSource](docs/VideoSource.md)
 - [VideoSourceLiveStream](docs/VideoSourceLiveStream.md)
 - [VideoSourceLiveStreamLink](docs/VideoSourceLiveStreamLink.md)
 - [VideoStatus](docs/VideoStatus.md)
 - [VideoStatusEncoding](docs/VideoStatusEncoding.md)
 - [VideoStatusEncodingMetadata](docs/VideoStatusEncodingMetadata.md)
 - [VideoStatusIngest](docs/VideoStatusIngest.md)
 - [VideoStatusIngestReceivedParts](docs/VideoStatusIngestReceivedParts.md)
 - [VideoThumbnailPickPayload](docs/VideoThumbnailPickPayload.md)
 - [VideoUpdatePayload](docs/VideoUpdatePayload.md)
 - [VideoWatermark](docs/VideoWatermark.md)
 - [VideosListResponse](docs/VideosListResponse.md)
 - [Watermark](docs/Watermark.md)
 - [WatermarksListResponse](docs/WatermarksListResponse.md)
 - [Webhook](docs/Webhook.md)
 - [WebhooksCreationPayload](docs/WebhooksCreationPayload.md)
 - [WebhooksListResponse](docs/WebhooksListResponse.md)



# Have you gotten use from this API client?

Please take a moment to leave a star on the client ⭐

This helps other users to find the clients and also helps us understand which clients are most popular. Thank you!

# Contribution

Since this API client is generated from an OpenAPI description, we cannot accept pull requests made directly to the repository. If you want to contribute, you can open a pull request on the repository of our [client generator](https://github.com/apivideo/api-client-generator). Otherwise, you can also simply open an issue detailing your need on this repository.