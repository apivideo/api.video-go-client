# VideoThumbnailPickPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timecode** | **string** | Frame in video to be used as a placeholder before the video plays.  Example: &#39;\&quot;00:01:00.000\&quot; for 1 minute into the video.&#39; Valid Patterns:  \&quot;hh:mm:ss.ms\&quot; \&quot;hh:mm:ss:frameNumber\&quot; \&quot;124\&quot; (integer value is reported as seconds)  If selection is out of range, \&quot;00:00:00.00\&quot; will be chosen. | 

## Methods

### NewVideoThumbnailPickPayload

`func NewVideoThumbnailPickPayload(timecode string, ) *VideoThumbnailPickPayload`

NewVideoThumbnailPickPayload instantiates a new VideoThumbnailPickPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoThumbnailPickPayloadWithDefaults

`func NewVideoThumbnailPickPayloadWithDefaults() *VideoThumbnailPickPayload`

NewVideoThumbnailPickPayloadWithDefaults instantiates a new VideoThumbnailPickPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimecode

`func (o *VideoThumbnailPickPayload) GetTimecode() string`

GetTimecode returns the Timecode field if non-nil, zero value otherwise.

### GetTimecodeOk

`func (o *VideoThumbnailPickPayload) GetTimecodeOk() (*string, bool)`

GetTimecodeOk returns a tuple with the Timecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimecode

`func (o *VideoThumbnailPickPayload) SetTimecode(v string)`

SetTimecode sets Timecode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


