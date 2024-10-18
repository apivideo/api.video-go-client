# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoId** | Pointer to **string** | Create a summary of a video using the video ID. | [optional] 
**Origin** | Pointer to **string** | Use this parameter to define how the API generates the summary. The only allowed value is &#x60;auto&#x60;, which means that the API generates a summary automatically.  If you do not set this parameter, **the API will not generate a summary automatically**.  In this case, &#x60;sourceStatus&#x60; will return &#x60;missing&#x60;, and you have to manually add a summary using the &#x60;PATCH /summaries/{summaryId}/source&#x60; endpoint operation. | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoId

`func (o *InlineObject) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *InlineObject) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *InlineObject) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *InlineObject) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetOrigin

`func (o *InlineObject) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *InlineObject) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *InlineObject) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *InlineObject) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


