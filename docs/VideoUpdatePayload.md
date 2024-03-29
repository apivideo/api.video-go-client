# VideoUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayerId** | Pointer to **NullableString** | The unique ID for the player you want to associate with your video. | [optional] 
**Title** | Pointer to **string** | The title you want to use for your video. | [optional] 
**Description** | Pointer to **string** | A brief description of the video. | [optional] 
**Public** | Pointer to **bool** | Whether the video is publicly available or not. False means it is set to private. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos/). | [optional] 
**Panoramic** | Pointer to **bool** | Whether the video is a 360 degree or immersive video. | [optional] 
**Mp4Support** | Pointer to **bool** | Whether the player supports the mp4 format. | [optional] 
**Tags** | Pointer to **[]string** | A list of terms or words you want to tag the video with. Make sure the list includes all the tags you want as whatever you send in this list will overwrite the existing list for the video. | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) | A list (array) of dictionaries where each dictionary contains a key value pair that describes the video. As with tags, you must send the complete list of metadata you want as whatever you send here will overwrite the existing metadata for the video. [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata/) allows you to define a key that allows any value pair. | [optional] 

## Methods

### NewVideoUpdatePayload

`func NewVideoUpdatePayload() *VideoUpdatePayload`

NewVideoUpdatePayload instantiates a new VideoUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoUpdatePayloadWithDefaults

`func NewVideoUpdatePayloadWithDefaults() *VideoUpdatePayload`

NewVideoUpdatePayloadWithDefaults instantiates a new VideoUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayerId

`func (o *VideoUpdatePayload) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *VideoUpdatePayload) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *VideoUpdatePayload) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *VideoUpdatePayload) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### SetPlayerIdNil

`func (o *VideoUpdatePayload) SetPlayerIdNil(b bool)`

 SetPlayerIdNil sets the value for PlayerId to be an explicit nil

### UnsetPlayerId
`func (o *VideoUpdatePayload) UnsetPlayerId()`

UnsetPlayerId ensures that no value is present for PlayerId, not even an explicit nil
### GetTitle

`func (o *VideoUpdatePayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VideoUpdatePayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VideoUpdatePayload) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VideoUpdatePayload) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *VideoUpdatePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoUpdatePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoUpdatePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoUpdatePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublic

`func (o *VideoUpdatePayload) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *VideoUpdatePayload) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *VideoUpdatePayload) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *VideoUpdatePayload) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPanoramic

`func (o *VideoUpdatePayload) GetPanoramic() bool`

GetPanoramic returns the Panoramic field if non-nil, zero value otherwise.

### GetPanoramicOk

`func (o *VideoUpdatePayload) GetPanoramicOk() (*bool, bool)`

GetPanoramicOk returns a tuple with the Panoramic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanoramic

`func (o *VideoUpdatePayload) SetPanoramic(v bool)`

SetPanoramic sets Panoramic field to given value.

### HasPanoramic

`func (o *VideoUpdatePayload) HasPanoramic() bool`

HasPanoramic returns a boolean if a field has been set.

### GetMp4Support

`func (o *VideoUpdatePayload) GetMp4Support() bool`

GetMp4Support returns the Mp4Support field if non-nil, zero value otherwise.

### GetMp4SupportOk

`func (o *VideoUpdatePayload) GetMp4SupportOk() (*bool, bool)`

GetMp4SupportOk returns a tuple with the Mp4Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4Support

`func (o *VideoUpdatePayload) SetMp4Support(v bool)`

SetMp4Support sets Mp4Support field to given value.

### HasMp4Support

`func (o *VideoUpdatePayload) HasMp4Support() bool`

HasMp4Support returns a boolean if a field has been set.

### GetTags

`func (o *VideoUpdatePayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VideoUpdatePayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VideoUpdatePayload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VideoUpdatePayload) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *VideoUpdatePayload) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VideoUpdatePayload) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VideoUpdatePayload) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VideoUpdatePayload) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


