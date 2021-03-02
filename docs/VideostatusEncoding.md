# VideostatusEncoding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Playable** | Pointer to **bool** | Whether the video is playable or not. | [optional] 
**Qualities** | Pointer to [**[]Quality**](Quality.md) | Available qualities the video can be viewed in. | [optional] 
**Metadata** | Pointer to [**VideostatusEncodingMetadata**](videostatus_encoding_metadata.md) |  | [optional] 

## Methods

### NewVideostatusEncoding

`func NewVideostatusEncoding() *VideostatusEncoding`

NewVideostatusEncoding instantiates a new VideostatusEncoding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideostatusEncodingWithDefaults

`func NewVideostatusEncodingWithDefaults() *VideostatusEncoding`

NewVideostatusEncodingWithDefaults instantiates a new VideostatusEncoding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayable

`func (o *VideostatusEncoding) GetPlayable() bool`

GetPlayable returns the Playable field if non-nil, zero value otherwise.

### GetPlayableOk

`func (o *VideostatusEncoding) GetPlayableOk() (*bool, bool)`

GetPlayableOk returns a tuple with the Playable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayable

`func (o *VideostatusEncoding) SetPlayable(v bool)`

SetPlayable sets Playable field to given value.

### HasPlayable

`func (o *VideostatusEncoding) HasPlayable() bool`

HasPlayable returns a boolean if a field has been set.

### GetQualities

`func (o *VideostatusEncoding) GetQualities() []Quality`

GetQualities returns the Qualities field if non-nil, zero value otherwise.

### GetQualitiesOk

`func (o *VideostatusEncoding) GetQualitiesOk() (*[]Quality, bool)`

GetQualitiesOk returns a tuple with the Qualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualities

`func (o *VideostatusEncoding) SetQualities(v []Quality)`

SetQualities sets Qualities field to given value.

### HasQualities

`func (o *VideostatusEncoding) HasQualities() bool`

HasQualities returns a boolean if a field has been set.

### GetMetadata

`func (o *VideostatusEncoding) GetMetadata() VideostatusEncodingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VideostatusEncoding) GetMetadataOk() (*VideostatusEncodingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VideostatusEncoding) SetMetadata(v VideostatusEncodingMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VideostatusEncoding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


