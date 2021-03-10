![](https://raw.githubusercontent.com/apivideo/API_OAS_file/master/apivideo_banner.png)

# api.video Golang api client

api.video is an API that encodes on the go to facilitate immediate playback, enhancing viewer streaming experiences across multiple devices and platforms. You can stream live or on-demand online videos within minutes.
This documentation helps you use the corresponding Golang client.


## Installation
```bash
go get github.com/apivideo/go-api-client
```


## Quick Start

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
    client := apivideosdk.NewClient(os.Getenv("API_VIDEO_KEY"))

    //Alternatively, connect to the sandbox environment for testing
    client := apivideosdk.NewSandboxClient(os.Getenv("API_VIDEO_SANDBOX_KEY"))

    //List Videos
    //First create the url options for searching
    opts := &apivideosdk.VideoOpts{
        CurrentPage: 1,
        PageSize: 25,
        SortBy:    "publishedAt",
        SortOrder: "desc",
    }

    //Then call the List endpoint with the options
    result, err := client.Videos.List(opts)

    if err != nil {
        fmt.Println(err)
    }

    for _, video := range result.Data {
        fmt.Printf("%s\n", video.VideoID)
        fmt.Printf("%s\n", video.Title)
    }

    //Upload a video
    //First create a container
    videoRequest := &apivideosdk.VideoRequest{
        Title: "My video title",
    }
    newVideo, err := client.Videos.Create(videoRequest)
    if err != nil {
        fmt.Println(err)
    }

    //Then upload your video to the container with the videoID
    uploadedVideo, err := client.Videos.Upload(newVideo.VideoID, "path/to/video.mp4")
    if err != nil {
        fmt.Println(err)
    }

    //And get the assets
    fmt.Printf("%s\n", uploadedVideo.Assets.Hls)
    fmt.Printf("%s\n", uploadedVideo.Assets.Iframe)
}
```

## Documentation for API Endpoints

All URIs are relative to *https://ws.api.video*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Account* | [**Get**](docs/Account.md#getaccount) | **Get** /account | Show account
*Authentication* | [**Authenticate**](docs/Authentication.md#postauthapikey) | **Post** /auth/api-key | Authenticate
*Authentication* | [**Refresh**](docs/Authentication.md#postauthrefresh) | **Post** /auth/refresh | Refresh token
*Captions* | [**Delete**](docs/Captions.md#deletevideosvideoidcaptionslanguage) | **Delete** /videos/{videoId}/captions/{language} | Delete a caption
*Captions* | [**List**](docs/Captions.md#getvideosvideoidcaptions) | **Get** /videos/{videoId}/captions | List video captions
*Captions* | [**Get**](docs/Captions.md#getvideosvideoidcaptionslanguage) | **Get** /videos/{videoId}/captions/{language} | Show a caption
*Captions* | [**Update**](docs/Captions.md#patchvideosvideoidcaptionslanguage) | **Patch** /videos/{videoId}/captions/{language} | Update caption
*Captions* | [**Upload**](docs/Captions.md#postvideosvideoidcaptionslanguage) | **Post** /videos/{videoId}/captions/{language} | Upload a caption
*Chapters* | [**Delete**](docs/Chapters.md#deletevideosvideoidchapterslanguage) | **Delete** /videos/{videoId}/chapters/{language} | Delete a chapter
*Chapters* | [**List**](docs/Chapters.md#getvideosvideoidchapters) | **Get** /videos/{videoId}/chapters | List video chapters
*Chapters* | [**Get**](docs/Chapters.md#getvideosvideoidchapterslanguage) | **Get** /videos/{videoId}/chapters/{language} | Show a chapter
*Chapters* | [**Upload**](docs/Chapters.md#postvideosvideoidchapterslanguage) | **Post** /videos/{videoId}/chapters/{language} | Upload a chapter
*Live* | [**Delete**](docs/Live.md#deletelivestreamslivestreamid) | **Delete** /live-streams/{liveStreamId} | Delete a live stream
*Live* | [**DeleteThumbnail**](docs/Live.md#deletelivestreamslivestreamidthumbnail) | **Delete** /live-streams/{liveStreamId}/thumbnail | Delete a thumbnail
*Live* | [**List**](docs/Live.md#getlivestreams) | **Get** /live-streams | List all live streams
*Live* | [**Get**](docs/Live.md#getlivestreamslivestreamid) | **Get** /live-streams/{liveStreamId} | Show live stream
*Live* | [**Update**](docs/Live.md#patchlivestreamslivestreamid) | **Patch** /live-streams/{liveStreamId} | Update a live stream
*Live* | [**Create**](docs/Live.md#postlivestreams) | **Post** /live-streams | Create live stream
*Live* | [**UploadThumbnail**](docs/Live.md#postlivestreamslivestreamidthumbnail) | **Post** /live-streams/{liveStreamId}/thumbnail | Upload a thumbnail
*Players* | [**Delete**](docs/Players.md#deleteplayersplayerid) | **Delete** /players/{playerId} | Delete a player
*Players* | [**DeleteLogo**](docs/Players.md#deleteplayersplayeridlogo) | **Delete** /players/{playerId}/logo | Delete logo
*Players* | [**List**](docs/Players.md#getplayers) | **Get** /players | List all players
*Players* | [**Get**](docs/Players.md#getplayersplayerid) | **Get** /players/{playerId} | Show a player
*Players* | [**Update**](docs/Players.md#patchplayersplayerid) | **Patch** /players/{playerId} | Update a player
*Players* | [**Create**](docs/Players.md#postplayers) | **Post** /players | Create a player
*Players* | [**UploadLogo**](docs/Players.md#postplayersplayeridlogo) | **Post** /players/{playerId}/logo | Upload a logo
*RawStatistics* | [**GetLiveStreamAnalytics**](docs/RawStatistics.md#getanalyticslivestreamslivestreamid) | **Get** /analytics/live-streams/{liveStreamId} | List live stream player sessions
*RawStatistics* | [**ListPlayerSessionEvents**](docs/RawStatistics.md#getanalyticssessionssessionidevents) | **Get** /analytics/sessions/{sessionId}/events | List player session events
*RawStatistics* | [**ListSessions**](docs/RawStatistics.md#getanalyticsvideosvideoid) | **Get** /analytics/videos/{videoId} | List video player sessions
*Videos* | [**Delete**](docs/Videos.md#deletevideo) | **Delete** /videos/{videoId} | Delete a video
*Videos* | [**Get**](docs/Videos.md#getvideo) | **Get** /videos/{videoId} | Show a video
*Videos* | [**GetVideoStatus**](docs/Videos.md#getvideostatus) | **Get** /videos/{videoId}/status | Show video status
*Videos* | [**List**](docs/Videos.md#listvideos) | **Get** /videos | List all videos
*Videos* | [**Update**](docs/Videos.md#patchvideo) | **Patch** /videos/{videoId} | Update a video
*Videos* | [**PickThumbnail**](docs/Videos.md#patchvideosvideoidthumbnail) | **Patch** /videos/{videoId}/thumbnail | Pick a thumbnail
*Videos* | [**Create**](docs/Videos.md#postvideo) | **Post** /videos | Create a video
*Videos* | [**Upload**](docs/Videos.md#postvideosvideoidsource) | **Post** /videos/{videoId}/source | Upload a video
*Videos* | [**UploadThumbnail**](docs/Videos.md#postvideosvideoidthumbnail) | **Post** /videos/{videoId}/thumbnail | Upload a thumbnail
*VideosDelegatedUpload* | [**DeleteToken**](docs/VideosDelegatedUpload.md#deleteuploadtokensuploadtoken) | **Delete** /upload-tokens/{uploadToken} | Delete an upload token
*VideosDelegatedUpload* | [**ListTokens**](docs/VideosDelegatedUpload.md#getuploadtokens) | **Get** /upload-tokens | List all active upload tokens.
*VideosDelegatedUpload* | [**GetToken**](docs/VideosDelegatedUpload.md#getuploadtokensuploadtoken) | **Get** /upload-tokens/{uploadToken} | Show upload token
*VideosDelegatedUpload* | [**Upload**](docs/VideosDelegatedUpload.md#postupload) | **Post** /upload | Upload with an upload token
*VideosDelegatedUpload* | [**CreateToken**](docs/VideosDelegatedUpload.md#postuploadtokens) | **Post** /upload-tokens | Generate an upload token
*Webhooks* | [**Delete**](docs/Webhooks.md#deletewebhook) | **Delete** /webhooks/{webhookId} | Delete a Webhook
*Webhooks* | [**Get**](docs/Webhooks.md#getwebhook) | **Get** /webhooks/{webhookId} | Show Webhook details
*Webhooks* | [**List**](docs/Webhooks.md#listwebhooks) | **Get** /webhooks | List all webhooks
*Webhooks* | [**Create**](docs/Webhooks.md#postwebhooks) | **Post** /webhooks | Create Webhook


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


