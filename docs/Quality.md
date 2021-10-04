# Quality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of video (hls or mp4). | [optional] 
**Quality** | Pointer to **string** | The quality of the video you have, in pixels. Choices include 360p, 480p, 720p, 1080p, and 2160p. | [optional] 
**Status** | Pointer to **string** | The status of your video. Statuses include waiting - the video is waiting to be encoded. encoding - the video is in the process of being encoded. encoded - the video was successfully encoded. failed - the video failed to be encoded. | [optional] 

## Methods

### NewQuality

`func NewQuality() *Quality`

NewQuality instantiates a new Quality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityWithDefaults

`func NewQualityWithDefaults() *Quality`

NewQualityWithDefaults instantiates a new Quality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Quality) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Quality) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Quality) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Quality) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuality

`func (o *Quality) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *Quality) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *Quality) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *Quality) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetStatus

`func (o *Quality) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Quality) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Quality) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Quality) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


