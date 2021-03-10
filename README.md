![](https://raw.githubusercontent.com/apivideo/API_OAS_file/master/apivideo_banner.png)

# api.video Golang api client

api.video is an API that encodes on the go to facilitate immediate playback, enhancing viewer streaming experiences across multiple devices and platforms. You can stream live or on-demand online videos within minutes.
This documentation helps you use the corresponding Golang client.


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
    create, err := client.Videos.Create(apivideosdk.VideoCreatePayload{Title: "My video title"})

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


### Account API


#### Retrieve an instance of the Account API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
accountApi := client.Account
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
**(deprecated)** [**Get**](docs/Account.md#Get) | **Get** /account | Show account


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


### Live API


#### Retrieve an instance of the Live API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
liveApi := client.Live
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](docs/Live.md#Delete) | **Delete** /live-streams/{liveStreamId} | Delete a live stream
[**DeleteThumbnail**](docs/Live.md#DeleteThumbnail) | **Delete** /live-streams/{liveStreamId}/thumbnail | Delete a thumbnail
[**List**](docs/Live.md#List) | **Get** /live-streams | List all live streams
[**Get**](docs/Live.md#Get) | **Get** /live-streams/{liveStreamId} | Show live stream
[**Update**](docs/Live.md#Update) | **Patch** /live-streams/{liveStreamId} | Update a live stream
[**Create**](docs/Live.md#Create) | **Post** /live-streams | Create live stream
[**UploadThumbnail**](docs/Live.md#UploadThumbnail) | **Post** /live-streams/{liveStreamId}/thumbnail | Upload a thumbnail


### Players API


#### Retrieve an instance of the Players API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
playersApi := client.Players
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](docs/Players.md#Delete) | **Delete** /players/{playerId} | Delete a player
[**DeleteLogo**](docs/Players.md#DeleteLogo) | **Delete** /players/{playerId}/logo | Delete logo
[**List**](docs/Players.md#List) | **Get** /players | List all players
[**Get**](docs/Players.md#Get) | **Get** /players/{playerId} | Show a player
[**Update**](docs/Players.md#Update) | **Patch** /players/{playerId} | Update a player
[**Create**](docs/Players.md#Create) | **Post** /players | Create a player
[**UploadLogo**](docs/Players.md#UploadLogo) | **Post** /players/{playerId}/logo | Upload a logo


### RawStatistics API


#### Retrieve an instance of the RawStatistics API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
rawStatisticsApi := client.RawStatistics
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLiveStreamAnalytics**](docs/RawStatistics.md#GetLiveStreamAnalytics) | **Get** /analytics/live-streams/{liveStreamId} | List live stream player sessions
[**ListPlayerSessionEvents**](docs/RawStatistics.md#ListPlayerSessionEvents) | **Get** /analytics/sessions/{sessionId}/events | List player session events
[**ListSessions**](docs/RawStatistics.md#ListSessions) | **Get** /analytics/videos/{videoId} | List video player sessions


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
[**GetVideoStatus**](docs/Videos.md#GetVideoStatus) | **Get** /videos/{videoId}/status | Show video status
[**List**](docs/Videos.md#List) | **Get** /videos | List all videos
[**Update**](docs/Videos.md#Update) | **Patch** /videos/{videoId} | Update a video
[**PickThumbnail**](docs/Videos.md#PickThumbnail) | **Patch** /videos/{videoId}/thumbnail | Pick a thumbnail
[**Create**](docs/Videos.md#Create) | **Post** /videos | Create a video
[**Upload**](docs/Videos.md#Upload) | **Post** /videos/{videoId}/source | Upload a video
[**UploadThumbnail**](docs/Videos.md#UploadThumbnail) | **Post** /videos/{videoId}/thumbnail | Upload a thumbnail


### VideosDelegatedUpload API


#### Retrieve an instance of the VideosDelegatedUpload API:
```golang
client := apivideosdk.ClientBuilder("API_VIDEO_KEY").Build()
videosDelegatedUploadApi := client.VideosDelegatedUpload
```

#### Endpoints

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteToken**](docs/VideosDelegatedUpload.md#DeleteToken) | **Delete** /upload-tokens/{uploadToken} | Delete an upload token
[**ListTokens**](docs/VideosDelegatedUpload.md#ListTokens) | **Get** /upload-tokens | List all active upload tokens.
[**GetToken**](docs/VideosDelegatedUpload.md#GetToken) | **Get** /upload-tokens/{uploadToken} | Show upload token
[**Upload**](docs/VideosDelegatedUpload.md#Upload) | **Post** /upload | Upload with an upload token
[**CreateToken**](docs/VideosDelegatedUpload.md#CreateToken) | **Post** /upload-tokens | Generate an upload token


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
 - [CaptionsListResponse](docs/CaptionsListResponse.md)
 - [CaptionsUpdatePayload](docs/CaptionsUpdatePayload.md)
 - [Chapter](docs/Chapter.md)
 - [ChaptersListResponse](docs/ChaptersListResponse.md)
 - [Link](docs/Link.md)
 - [LiveStream](docs/LiveStream.md)
 - [LiveStreamAssets](docs/LiveStreamAssets.md)
 - [LiveStreamCreatePayload](docs/LiveStreamCreatePayload.md)
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
 - [Player](docs/Player.md)
 - [PlayerAllOf](docs/PlayerAllOf.md)
 - [PlayerAllOfAssets](docs/PlayerAllOfAssets.md)
 - [PlayerCreationPayload](docs/PlayerCreationPayload.md)
 - [PlayerSessionEvent](docs/PlayerSessionEvent.md)
 - [PlayerUpdatePayload](docs/PlayerUpdatePayload.md)
 - [Playerinput](docs/Playerinput.md)
 - [PlayersListResponse](docs/PlayersListResponse.md)
 - [Quality](docs/Quality.md)
 - [RawStatisticsListLiveStreamAnalyticsResponse](docs/RawStatisticsListLiveStreamAnalyticsResponse.md)
 - [RawStatisticsListPlayerSessionEventsResponse](docs/RawStatisticsListPlayerSessionEventsResponse.md)
 - [RawStatisticsListSessionsResponse](docs/RawStatisticsListSessionsResponse.md)
 - [RefreshTokenPayload](docs/RefreshTokenPayload.md)
 - [Subtitle](docs/Subtitle.md)
 - [TokenCreatePayload](docs/TokenCreatePayload.md)
 - [TokenListResponse](docs/TokenListResponse.md)
 - [UploadToken](docs/UploadToken.md)
 - [Video](docs/Video.md)
 - [VideoAssets](docs/VideoAssets.md)
 - [VideoCreatePayload](docs/VideoCreatePayload.md)
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
 - [VideoThumbnailPickPayload](docs/VideoThumbnailPickPayload.md)
 - [VideoUpdatePayload](docs/VideoUpdatePayload.md)
 - [VideosListResponse](docs/VideosListResponse.md)
 - [Videostatus](docs/Videostatus.md)
 - [VideostatusEncoding](docs/VideostatusEncoding.md)
 - [VideostatusEncodingMetadata](docs/VideostatusEncodingMetadata.md)
 - [VideostatusIngest](docs/VideostatusIngest.md)
 - [Webhook](docs/Webhook.md)
 - [WebhooksCreatePayload](docs/WebhooksCreatePayload.md)
 - [WebhooksListResponse](docs/WebhooksListResponse.md)


