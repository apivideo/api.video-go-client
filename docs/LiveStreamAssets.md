# LiveStreamAssets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hls** | Pointer to **string** | The http live streaming (HLS) link for your live video stream. | [optional] 
**Iframe** | Pointer to **string** | The embed code for the iframe containing your live video stream. | [optional] 
**Player** | Pointer to **string** | A link to the video player that is playing your live stream. | [optional] 
**Thumbnail** | Pointer to **string** | A link to the thumbnail for your video. | [optional] 

## Methods

### NewLiveStreamAssets

`func NewLiveStreamAssets() *LiveStreamAssets`

NewLiveStreamAssets instantiates a new LiveStreamAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamAssetsWithDefaults

`func NewLiveStreamAssetsWithDefaults() *LiveStreamAssets`

NewLiveStreamAssetsWithDefaults instantiates a new LiveStreamAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHls

`func (o *LiveStreamAssets) GetHls() string`

GetHls returns the Hls field if non-nil, zero value otherwise.

### GetHlsOk

`func (o *LiveStreamAssets) GetHlsOk() (*string, bool)`

GetHlsOk returns a tuple with the Hls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHls

`func (o *LiveStreamAssets) SetHls(v string)`

SetHls sets Hls field to given value.

### HasHls

`func (o *LiveStreamAssets) HasHls() bool`

HasHls returns a boolean if a field has been set.

### GetIframe

`func (o *LiveStreamAssets) GetIframe() string`

GetIframe returns the Iframe field if non-nil, zero value otherwise.

### GetIframeOk

`func (o *LiveStreamAssets) GetIframeOk() (*string, bool)`

GetIframeOk returns a tuple with the Iframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIframe

`func (o *LiveStreamAssets) SetIframe(v string)`

SetIframe sets Iframe field to given value.

### HasIframe

`func (o *LiveStreamAssets) HasIframe() bool`

HasIframe returns a boolean if a field has been set.

### GetPlayer

`func (o *LiveStreamAssets) GetPlayer() string`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *LiveStreamAssets) GetPlayerOk() (*string, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *LiveStreamAssets) SetPlayer(v string)`

SetPlayer sets Player field to given value.

### HasPlayer

`func (o *LiveStreamAssets) HasPlayer() bool`

HasPlayer returns a boolean if a field has been set.

### GetThumbnail

`func (o *LiveStreamAssets) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *LiveStreamAssets) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *LiveStreamAssets) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *LiveStreamAssets) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


