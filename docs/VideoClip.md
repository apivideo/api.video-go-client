# VideoClip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimecode** | Pointer to **string** | The timestamp that defines the beginning of the video clip you want to create. The value must follow the &#x60;HH:MM:SS&#x60; format. | [optional] 
**EndTimecode** | Pointer to **string** | The timestamp that defines the end of the video clip you want to create. The value must follow the &#x60;HH:MM:SS&#x60; format. | [optional] 

## Methods

### NewVideoClip

`func NewVideoClip() *VideoClip`

NewVideoClip instantiates a new VideoClip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoClipWithDefaults

`func NewVideoClipWithDefaults() *VideoClip`

NewVideoClipWithDefaults instantiates a new VideoClip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimecode

`func (o *VideoClip) GetStartTimecode() string`

GetStartTimecode returns the StartTimecode field if non-nil, zero value otherwise.

### GetStartTimecodeOk

`func (o *VideoClip) GetStartTimecodeOk() (*string, bool)`

GetStartTimecodeOk returns a tuple with the StartTimecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimecode

`func (o *VideoClip) SetStartTimecode(v string)`

SetStartTimecode sets StartTimecode field to given value.

### HasStartTimecode

`func (o *VideoClip) HasStartTimecode() bool`

HasStartTimecode returns a boolean if a field has been set.

### GetEndTimecode

`func (o *VideoClip) GetEndTimecode() string`

GetEndTimecode returns the EndTimecode field if non-nil, zero value otherwise.

### GetEndTimecodeOk

`func (o *VideoClip) GetEndTimecodeOk() (*string, bool)`

GetEndTimecodeOk returns a tuple with the EndTimecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimecode

`func (o *VideoClip) SetEndTimecode(v string)`

SetEndTimecode sets EndTimecode field to given value.

### HasEndTimecode

`func (o *VideoClip) HasEndTimecode() bool`

HasEndTimecode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


