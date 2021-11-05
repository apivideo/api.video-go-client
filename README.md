[![badge](https://img.shields.io/twitter/follow/api_video?style=social)](https://twitter.com/intent/follow?screen_name=api_video)

[![badge](https://img.shields.io/github/stars/apivideo/go-api-client?style=social)](https://github.com/apivideo/go-api-client)

[![badge](https://img.shields.io/discourse/topics?server=https%3A%2F%2Fcommunity.api.video)](https://community.api.video)

![](https://github.com/apivideo/API_OAS_file/blob/master/apivideo_banner.png)

[api.video](https://api.video) is an API that encodes on the go to facilitate immediate playback, enhancing viewer streaming experiences across multiple devices and platforms. You can stream live or on-demand online videos within minutes.

# api.video Golang api client

api.video is an API that encodes on the go to facilitate immediate playback, enhancing viewer streaming experiences across multiple devices and platforms. You can stream live or on-demand online videos within minutes.

## Warning

This API client is still in beta. Please feel free to report any issue you may encounter.

## Installation
```bash
go get github.com/apivideo/go-api-client
```

## Getting Started

For a more advanced usage you can checkout the rest of the documentation in the [docs directory](/docs)

```golang
package main

import (
    "fmt"
    "os"
    apivideosdk "github.com/apivideo/go-api-client"
)

func main() {
    //Connect to production environment
    client := apivideosdk.ClientBuilder("YOUR_API_TOKEN").Build()

    // if you rather like to use the sandbox environment:
    // client := apivideosdk.SandboxClientBuilder("YOU_SANDBOX_API_TOKEN").Build()


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
        fmt.Printf("%s\n", *video.VideoId)
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
    uploadedVideo, err := client.Videos.UploadFile(*create.VideoId, videoFile)

    if err != nil {
        fmt.Println(err)
    }


    //And get the assets
    fmt.Printf("%s\n", *uploadedVideo.Assets.Hls)
    fmt.Printf("%s\n", *uploadedVideo.Assets.Iframe)
}
```

## Documentation for API Endpoints

All URIs are relative to *https://ws.api.video*


### Captions API


#### Retrieve an instance of the Captions API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
captionsApi := client.Captions
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](docs/Captions.md#Delete) | **Delete** /videos/{videoId}/captions/{language} | Delete a caption
[**List**](docs/Captions.md#List) | **Get** /videos/{videoId}/captions | List video captions
[**Get**](docs/Captions.md#Get) | **Get** /videos/{videoId}/captions/{language} | Show a caption
[**Update**](docs/Captions.md#Update) | **Patch** /videos/{videoId}/captions/{language} | Update caption
[**Upload**](docs/Captions.md#Upload) | **Post** /videos/{videoId}/captions/{language} | Upload a caption


### Chapters API


#### Retrieve an instance of the Chapters API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
chaptersApi := client.Chapters
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](docs/Chapters.md#Delete) | **Delete** /videos/{videoId}/chapters/{language} | Delete a chapter
[**List**](docs/Chapters.md#List) | **Get** /videos/{videoId}/chapters | List video chapters
[**Get**](docs/Chapters.md#Get) | **Get** /videos/{videoId}/chapters/{language} | Show a chapter
[**Upload**](docs/Chapters.md#Upload) | **Post** /videos/{videoId}/chapters/{language} | Upload a chapter


### LiveStreams API


#### Retrieve an instance of the LiveStreams API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
liveStreamsApi := client.LiveStreams
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](docs/LiveStreams.md#Delete) | **Delete** /live-streams/{liveStreamId} | Delete a live stream
[**DeleteThumbnail**](docs/LiveStreams.md#DeleteThumbnail) | **Delete** /live-streams/{liveStreamId}/thumbnail | Delete a thumbnail
[**List**](docs/LiveStreams.md#List) | **Get** /live-streams | List all live streams
[**Get**](docs/LiveStreams.md#Get) | **Get** /live-streams/{liveStreamId} | Show live stream
[**Update**](docs/LiveStreams.md#Update) | **Patch** /live-streams/{liveStreamId} | Update a live stream
[**Create**](docs/LiveStreams.md#Create) | **Post** /live-streams | Create live stream
[**UploadThumbnail**](docs/LiveStreams.md#UploadThumbnail) | **Post** /live-streams/{liveStreamId}/thumbnail | Upload a thumbnail


### PlayerThemes API


#### Retrieve an instance of the PlayerThemes API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
playerThemesApi := client.PlayerThemes
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](docs/PlayerThemes.md#Delete) | **Delete** /players/{playerId} | Delete a player
[**DeleteLogo**](docs/PlayerThemes.md#DeleteLogo) | **Delete** /players/{playerId}/logo | Delete logo
[**List**](docs/PlayerThemes.md#List) | **Get** /players | List all players
[**Get**](docs/PlayerThemes.md#Get) | **Get** /players/{playerId} | Show a player
[**Update**](docs/PlayerThemes.md#Update) | **Patch** /players/{playerId} | Update a player
[**Create**](docs/PlayerThemes.md#Create) | **Post** /players | Create a player
[**UploadLogo**](docs/PlayerThemes.md#UploadLogo) | **Post** /players/{playerId}/logo | Upload a logo


### RawStatistics API


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


### UploadTokens API


#### Retrieve an instance of the UploadTokens API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
uploadTokensApi := client.UploadTokens
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteToken**](docs/UploadTokens.md#DeleteToken) | **Delete** /upload-tokens/{uploadToken} | Delete an upload token
[**List**](docs/UploadTokens.md#List) | **Get** /upload-tokens | List all active upload tokens.
[**GetToken**](docs/UploadTokens.md#GetToken) | **Get** /upload-tokens/{uploadToken} | Show upload token
[**CreateToken**](docs/UploadTokens.md#CreateToken) | **Post** /upload-tokens | Generate an upload token


### Videos API


#### Retrieve an instance of the Videos API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
videosApi := client.Videos
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](docs/Videos.md#Delete) | **Delete** /videos/{videoId} | Delete a video
[**Get**](docs/Videos.md#Get) | **Get** /videos/{videoId} | Show a video
[**GetStatus**](docs/Videos.md#GetStatus) | **Get** /videos/{videoId}/status | Show video status
[**List**](docs/Videos.md#List) | **Get** /videos | List all videos
[**Update**](docs/Videos.md#Update) | **Patch** /videos/{videoId} | Update a video
[**PickThumbnail**](docs/Videos.md#PickThumbnail) | **Patch** /videos/{videoId}/thumbnail | Pick a thumbnail
[**UploadWithUploadToken**](docs/Videos.md#UploadWithUploadToken) | **Post** /upload | Upload with an upload token
[**Create**](docs/Videos.md#Create) | **Post** /videos | Create a video
[**Upload**](docs/Videos.md#Upload) | **Post** /videos/{videoId}/source | Upload a video
[**UploadThumbnail**](docs/Videos.md#UploadThumbnail) | **Post** /videos/{videoId}/thumbnail | Upload a thumbnail


### Webhooks API


#### Retrieve an instance of the Webhooks API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
webhooksApi := client.Webhooks
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](docs/Webhooks.md#Delete) | **Delete** /webhooks/{webhookId} | Delete a Webhook
[**Get**](docs/Webhooks.md#Get) | **Get** /webhooks/{webhookId} | Show Webhook details
[**List**](docs/Webhooks.md#List) | **Get** /webhooks | List all webhooks
[**Create**](docs/Webhooks.md#Create) | **Post** /webhooks | Create Webhook




## Documentation For Models

 - [AccessToken](docs/AccessToken.md)
 - [Account](docs/Account.md)
 - [AccountQuota](docs/AccountQuota.md)
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
 - [VideoThumbnailPickPayload](docs/VideoThumbnailPickPayload.md)
 - [VideoUpdatePayload](docs/VideoUpdatePayload.md)
 - [VideosListResponse](docs/VideosListResponse.md)
 - [Webhook](docs/Webhook.md)
 - [WebhooksCreationPayload](docs/WebhooksCreationPayload.md)
 - [WebhooksListResponse](docs/WebhooksListResponse.md)



## Have you gotten use from this API client?

Please take a moment to leave a star on the client ⭐

This helps other users to find the clients and also helps us understand which clients are most popular. Thank you!
