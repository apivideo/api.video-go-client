# VideostatusEncodingMetadata

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
**AspectRatio** | Pointer to **string** |  | [optional] 

## Methods

### NewVideostatusEncodingMetadata

`func NewVideostatusEncodingMetadata() *VideostatusEncodingMetadata`

NewVideostatusEncodingMetadata instantiates a new VideostatusEncodingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideostatusEncodingMetadataWithDefaults

`func NewVideostatusEncodingMetadataWithDefaults() *VideostatusEncodingMetadata`

NewVideostatusEncodingMetadataWithDefaults instantiates a new VideostatusEncodingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidth

`func (o *VideostatusEncodingMetadata) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VideostatusEncodingMetadata) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VideostatusEncodingMetadata) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VideostatusEncodingMetadata) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *VideostatusEncodingMetadata) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VideostatusEncodingMetadata) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VideostatusEncodingMetadata) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VideostatusEncodingMetadata) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetBitrate

`func (o *VideostatusEncodingMetadata) GetBitrate() float32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *VideostatusEncodingMetadata) GetBitrateOk() (*float32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *VideostatusEncodingMetadata) SetBitrate(v float32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *VideostatusEncodingMetadata) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### GetDuration

`func (o *VideostatusEncodingMetadata) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideostatusEncodingMetadata) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideostatusEncodingMetadata) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideostatusEncodingMetadata) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFramerate

`func (o *VideostatusEncodingMetadata) GetFramerate() int32`

GetFramerate returns the Framerate field if non-nil, zero value otherwise.

### GetFramerateOk

`func (o *VideostatusEncodingMetadata) GetFramerateOk() (*int32, bool)`

GetFramerateOk returns a tuple with the Framerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramerate

`func (o *VideostatusEncodingMetadata) SetFramerate(v int32)`

SetFramerate sets Framerate field to given value.

### HasFramerate

`func (o *VideostatusEncodingMetadata) HasFramerate() bool`

HasFramerate returns a boolean if a field has been set.

### GetSamplerate

`func (o *VideostatusEncodingMetadata) GetSamplerate() int32`

GetSamplerate returns the Samplerate field if non-nil, zero value otherwise.

### GetSamplerateOk

`func (o *VideostatusEncodingMetadata) GetSamplerateOk() (*int32, bool)`

GetSamplerateOk returns a tuple with the Samplerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplerate

`func (o *VideostatusEncodingMetadata) SetSamplerate(v int32)`

SetSamplerate sets Samplerate field to given value.

### HasSamplerate

`func (o *VideostatusEncodingMetadata) HasSamplerate() bool`

HasSamplerate returns a boolean if a field has been set.

### GetVideoCodec

`func (o *VideostatusEncodingMetadata) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *VideostatusEncodingMetadata) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *VideostatusEncodingMetadata) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *VideostatusEncodingMetadata) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetAudioCodec

`func (o *VideostatusEncodingMetadata) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *VideostatusEncodingMetadata) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *VideostatusEncodingMetadata) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *VideostatusEncodingMetadata) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetAspectRatio

`func (o *VideostatusEncodingMetadata) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *VideostatusEncodingMetadata) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *VideostatusEncodingMetadata) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *VideostatusEncodingMetadata) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


