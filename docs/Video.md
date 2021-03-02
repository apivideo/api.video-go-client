# Video

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoId** | Pointer to **string** | The unique identifier of the video object. | [optional] 
**Title** | Pointer to **string** | The title of the video content.  | [optional] 
**Description** | Pointer to **string** | A description for the video content.  | [optional] 
**PublishedAt** | Pointer to **string** | The date and time the API created the video. Date and time are provided using ISO-8601 UTC format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time the video was updated. Date and time are provided using ISO-8601 UTC format. | [optional] 
**Tags** | Pointer to **[]interface{}** | One array of tags (each tag is a string) in order to categorize a video. Tags may include spaces.   | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) | Metadata you can use to categorise and filter videos. Metadata is a list of dictionaries, where each dictionary represents a key value pair for categorising a video.   | [optional] 
**Source** | Pointer to [**VideoSource**](videoSource.md) |  | [optional] 
**Assets** | Pointer to [**VideoAssets**](videoAssets.md) |  | [optional] 
**PlayerId** | Pointer to **string** | The id of the player that will be applied on the video.  | [optional] 
**Public** | Pointer to **bool** | Defines if the content is publicly reachable or if a unique token is needed for each play session.  | [optional] 
**Panoramic** | Pointer to **bool** | Defines if video is panoramic.  | [optional] 
**Mp4Support** | Pointer to **bool** | This lets you know whether mp4 is supported. If enabled, an mp4 URL will be provided in the response for the video.  | [optional] 

## Methods

### NewVideo

`func NewVideo() *Video`

NewVideo instantiates a new Video object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithDefaults

`func NewVideoWithDefaults() *Video`

NewVideoWithDefaults instantiates a new Video object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoId

`func (o *Video) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *Video) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *Video) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *Video) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetTitle

`func (o *Video) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Video) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Video) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Video) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Video) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Video) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Video) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Video) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublishedAt

`func (o *Video) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *Video) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *Video) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *Video) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Video) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Video) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Video) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Video) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *Video) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Video) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Video) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Video) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *Video) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Video) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Video) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Video) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSource

`func (o *Video) GetSource() VideoSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Video) GetSourceOk() (*VideoSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Video) SetSource(v VideoSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Video) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAssets

`func (o *Video) GetAssets() VideoAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Video) GetAssetsOk() (*VideoAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Video) SetAssets(v VideoAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *Video) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPlayerId

`func (o *Video) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *Video) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *Video) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *Video) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetPublic

`func (o *Video) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Video) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Video) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *Video) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPanoramic

`func (o *Video) GetPanoramic() bool`

GetPanoramic returns the Panoramic field if non-nil, zero value otherwise.

### GetPanoramicOk

`func (o *Video) GetPanoramicOk() (*bool, bool)`

GetPanoramicOk returns a tuple with the Panoramic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanoramic

`func (o *Video) SetPanoramic(v bool)`

SetPanoramic sets Panoramic field to given value.

### HasPanoramic

`func (o *Video) HasPanoramic() bool`

HasPanoramic returns a boolean if a field has been set.

### GetMp4Support

`func (o *Video) GetMp4Support() bool`

GetMp4Support returns the Mp4Support field if non-nil, zero value otherwise.

### GetMp4SupportOk

`func (o *Video) GetMp4SupportOk() (*bool, bool)`

GetMp4SupportOk returns a tuple with the Mp4Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4Support

`func (o *Video) SetMp4Support(v bool)`

SetMp4Support sets Mp4Support field to given value.

### HasMp4Support

`func (o *Video) HasMp4Support() bool`

HasMp4Support returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


