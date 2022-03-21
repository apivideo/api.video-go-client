# VideoStatusEncodingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Width** | Pointer to **int32** | The width of the video in pixels. | [optional] 
**Height** | Pointer to **int32** | The height of the video in pixels. | [optional] 
**Bitrate** | Pointer to **float32** | The number of bits processed per second. | [optional] 
**Duration** | Pointer to **int32** | The length of the video. | [optional] 
**Framerate** | Pointer to **int32** | The frequency with which consecutive images or frames appear on a display. Shown in this API as frames per second (fps). | [optional] 
**Samplerate** | Pointer to **int32** | How many samples per second a digital audio system uses to record an audio signal. The higher the rate, the higher the frequencies that can be recorded. They are presented in this API using hertz. | [optional] 
**VideoCodec** | Pointer to **string** | The method used to compress and decompress digital video. API Video supports all codecs in the libavcodec library.  | [optional] 
**AudioCodec** | Pointer to **string** | The method used to compress and decompress digital audio for your video. | [optional] 
**AspectRatio** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVideoStatusEncodingMetadata

`func NewVideoStatusEncodingMetadata() *VideoStatusEncodingMetadata`

NewVideoStatusEncodingMetadata instantiates a new VideoStatusEncodingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoStatusEncodingMetadataWithDefaults

`func NewVideoStatusEncodingMetadataWithDefaults() *VideoStatusEncodingMetadata`

NewVideoStatusEncodingMetadataWithDefaults instantiates a new VideoStatusEncodingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidth

`func (o *VideoStatusEncodingMetadata) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VideoStatusEncodingMetadata) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VideoStatusEncodingMetadata) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VideoStatusEncodingMetadata) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *VideoStatusEncodingMetadata) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VideoStatusEncodingMetadata) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VideoStatusEncodingMetadata) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VideoStatusEncodingMetadata) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetBitrate

`func (o *VideoStatusEncodingMetadata) GetBitrate() float32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *VideoStatusEncodingMetadata) GetBitrateOk() (*float32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *VideoStatusEncodingMetadata) SetBitrate(v float32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *VideoStatusEncodingMetadata) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### GetDuration

`func (o *VideoStatusEncodingMetadata) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoStatusEncodingMetadata) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoStatusEncodingMetadata) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoStatusEncodingMetadata) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFramerate

`func (o *VideoStatusEncodingMetadata) GetFramerate() int32`

GetFramerate returns the Framerate field if non-nil, zero value otherwise.

### GetFramerateOk

`func (o *VideoStatusEncodingMetadata) GetFramerateOk() (*int32, bool)`

GetFramerateOk returns a tuple with the Framerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramerate

`func (o *VideoStatusEncodingMetadata) SetFramerate(v int32)`

SetFramerate sets Framerate field to given value.

### HasFramerate

`func (o *VideoStatusEncodingMetadata) HasFramerate() bool`

HasFramerate returns a boolean if a field has been set.

### GetSamplerate

`func (o *VideoStatusEncodingMetadata) GetSamplerate() int32`

GetSamplerate returns the Samplerate field if non-nil, zero value otherwise.

### GetSamplerateOk

`func (o *VideoStatusEncodingMetadata) GetSamplerateOk() (*int32, bool)`

GetSamplerateOk returns a tuple with the Samplerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplerate

`func (o *VideoStatusEncodingMetadata) SetSamplerate(v int32)`

SetSamplerate sets Samplerate field to given value.

### HasSamplerate

`func (o *VideoStatusEncodingMetadata) HasSamplerate() bool`

HasSamplerate returns a boolean if a field has been set.

### GetVideoCodec

`func (o *VideoStatusEncodingMetadata) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *VideoStatusEncodingMetadata) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *VideoStatusEncodingMetadata) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *VideoStatusEncodingMetadata) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetAudioCodec

`func (o *VideoStatusEncodingMetadata) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *VideoStatusEncodingMetadata) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *VideoStatusEncodingMetadata) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *VideoStatusEncodingMetadata) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetAspectRatio

`func (o *VideoStatusEncodingMetadata) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *VideoStatusEncodingMetadata) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *VideoStatusEncodingMetadata) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *VideoStatusEncodingMetadata) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *VideoStatusEncodingMetadata) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *VideoStatusEncodingMetadata) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


