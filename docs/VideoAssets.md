# VideoAssets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hls** | Pointer to **string** | This is the manifest URL. For HTTP Live Streaming (HLS), when a HLS video stream is initiated, the first file to download is the manifest. This file has the extension M3U8, and provides the video player with information about the various bitrates available for streaming. | [optional] 
**Iframe** | Pointer to **string** | Code to use video from a third party website | [optional] 
**Player** | Pointer to **string** | Raw url of the player. | [optional] 
**Thumbnail** | Pointer to **string** | Poster of the video. | [optional] 
**Mp4** | Pointer to **string** | Available only if mp4Support is enabled. Raw mp4 url. | [optional] 

## Methods

### NewVideoAssets

`func NewVideoAssets() *VideoAssets`

NewVideoAssets instantiates a new VideoAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoAssetsWithDefaults

`func NewVideoAssetsWithDefaults() *VideoAssets`

NewVideoAssetsWithDefaults instantiates a new VideoAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHls

`func (o *VideoAssets) GetHls() string`

GetHls returns the Hls field if non-nil, zero value otherwise.

### GetHlsOk

`func (o *VideoAssets) GetHlsOk() (*string, bool)`

GetHlsOk returns a tuple with the Hls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHls

`func (o *VideoAssets) SetHls(v string)`

SetHls sets Hls field to given value.

### HasHls

`func (o *VideoAssets) HasHls() bool`

HasHls returns a boolean if a field has been set.

### GetIframe

`func (o *VideoAssets) GetIframe() string`

GetIframe returns the Iframe field if non-nil, zero value otherwise.

### GetIframeOk

`func (o *VideoAssets) GetIframeOk() (*string, bool)`

GetIframeOk returns a tuple with the Iframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIframe

`func (o *VideoAssets) SetIframe(v string)`

SetIframe sets Iframe field to given value.

### HasIframe

`func (o *VideoAssets) HasIframe() bool`

HasIframe returns a boolean if a field has been set.

### GetPlayer

`func (o *VideoAssets) GetPlayer() string`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *VideoAssets) GetPlayerOk() (*string, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *VideoAssets) SetPlayer(v string)`

SetPlayer sets Player field to given value.

### HasPlayer

`func (o *VideoAssets) HasPlayer() bool`

HasPlayer returns a boolean if a field has been set.

### GetThumbnail

`func (o *VideoAssets) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *VideoAssets) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *VideoAssets) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *VideoAssets) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetMp4

`func (o *VideoAssets) GetMp4() string`

GetMp4 returns the Mp4 field if non-nil, zero value otherwise.

### GetMp4Ok

`func (o *VideoAssets) GetMp4Ok() (*string, bool)`

GetMp4Ok returns a tuple with the Mp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4

`func (o *VideoAssets) SetMp4(v string)`

SetMp4 sets Mp4 field to given value.

### HasMp4

`func (o *VideoAssets) HasMp4() bool`

HasMp4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


