# VideoWatermark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of the watermark | [optional] 
**Top** | Pointer to **string** | Distance expressed in px or % between the top-border of the video and the watermark-image. | [optional] 
**Left** | Pointer to **string** | Distance expressed in px or % between the left-border of the video and the watermark-image. | [optional] 
**Bottom** | Pointer to **string** | Distance expressed in px or % between the bottom-border of the video and the watermark-image. | [optional] 
**Right** | Pointer to **string** | Distance expressed in px or % between the right-border of the video and the watermark-image. | [optional] 
**Width** | Pointer to **string** | Width of the watermark-image relative to the video if expressed in %. Otherwise a fixed width. NOTE: To keep intrinsic watermark-image width use initial | [optional] 
**Height** | Pointer to **string** | Width of the watermark-image relative to the video if expressed in %. Otherwise a fixed height. NOTE: To keep intrinsic watermark-image height use initial | [optional] 
**Opacity** | Pointer to **string** | Opacity expressed in % only to specify the degree of the watermark-image transparency with the video. | [optional] 

## Methods

### NewVideoWatermark

`func NewVideoWatermark() *VideoWatermark`

NewVideoWatermark instantiates a new VideoWatermark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWatermarkWithDefaults

`func NewVideoWatermarkWithDefaults() *VideoWatermark`

NewVideoWatermarkWithDefaults instantiates a new VideoWatermark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoWatermark) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoWatermark) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoWatermark) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VideoWatermark) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTop

`func (o *VideoWatermark) GetTop() string`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *VideoWatermark) GetTopOk() (*string, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *VideoWatermark) SetTop(v string)`

SetTop sets Top field to given value.

### HasTop

`func (o *VideoWatermark) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetLeft

`func (o *VideoWatermark) GetLeft() string`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *VideoWatermark) GetLeftOk() (*string, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *VideoWatermark) SetLeft(v string)`

SetLeft sets Left field to given value.

### HasLeft

`func (o *VideoWatermark) HasLeft() bool`

HasLeft returns a boolean if a field has been set.

### GetBottom

`func (o *VideoWatermark) GetBottom() string`

GetBottom returns the Bottom field if non-nil, zero value otherwise.

### GetBottomOk

`func (o *VideoWatermark) GetBottomOk() (*string, bool)`

GetBottomOk returns a tuple with the Bottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottom

`func (o *VideoWatermark) SetBottom(v string)`

SetBottom sets Bottom field to given value.

### HasBottom

`func (o *VideoWatermark) HasBottom() bool`

HasBottom returns a boolean if a field has been set.

### GetRight

`func (o *VideoWatermark) GetRight() string`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *VideoWatermark) GetRightOk() (*string, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *VideoWatermark) SetRight(v string)`

SetRight sets Right field to given value.

### HasRight

`func (o *VideoWatermark) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetWidth

`func (o *VideoWatermark) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VideoWatermark) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VideoWatermark) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VideoWatermark) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *VideoWatermark) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VideoWatermark) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VideoWatermark) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VideoWatermark) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetOpacity

`func (o *VideoWatermark) GetOpacity() string`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *VideoWatermark) GetOpacityOk() (*string, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *VideoWatermark) SetOpacity(v string)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *VideoWatermark) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


