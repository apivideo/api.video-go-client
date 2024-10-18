# SummaryObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryId** | Pointer to **string** | The unique identifier of the summary object. | [optional] 
**CreatedAt** | Pointer to **string** | Returns the date and time when the summary was created in ATOM date-time format. | [optional] 
**UpdatedAt** | Pointer to **string** | Returns the date and time when the summary was last updated in ATOM date-time format. | [optional] 
**VideoId** | Pointer to **string** | The unique identifier of the video object. | [optional] 
**Origin** | Pointer to **interface{}** | Returns the origin of how the summary was created.  - &#x60;api&#x60; means that no summary was generated automatically. You can add summary manually using the &#x60;PATCH /summaries/{summaryId}/source&#x60; endpoint operation. Until this happens, &#x60;sourceStatus&#x60; returns &#x60;missing&#x60;. - &#x60;auto&#x60; means that the API generated the summary automatically. | [optional] 
**SourceStatus** | Pointer to **string** | Returns the current status of summary generation.  &#x60;missing&#x60;: the input for a summary is not present. &#x60;waiting&#x60; : the input video is being processed and a summary will be generated. &#x60;failed&#x60;: a technical issue prevented summary generation. &#x60;completed&#x60;: the summary is generated. &#x60;unprocessable&#x60;: the API rules the source video to be unsuitable for summary generation. An example for this is an input video that has no audio. | [optional] 

## Methods

### NewSummaryObject

`func NewSummaryObject() *SummaryObject`

NewSummaryObject instantiates a new SummaryObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryObjectWithDefaults

`func NewSummaryObjectWithDefaults() *SummaryObject`

NewSummaryObjectWithDefaults instantiates a new SummaryObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryId

`func (o *SummaryObject) GetSummaryId() string`

GetSummaryId returns the SummaryId field if non-nil, zero value otherwise.

### GetSummaryIdOk

`func (o *SummaryObject) GetSummaryIdOk() (*string, bool)`

GetSummaryIdOk returns a tuple with the SummaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryId

`func (o *SummaryObject) SetSummaryId(v string)`

SetSummaryId sets SummaryId field to given value.

### HasSummaryId

`func (o *SummaryObject) HasSummaryId() bool`

HasSummaryId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SummaryObject) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SummaryObject) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SummaryObject) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SummaryObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SummaryObject) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SummaryObject) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SummaryObject) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SummaryObject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVideoId

`func (o *SummaryObject) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *SummaryObject) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *SummaryObject) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *SummaryObject) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetOrigin

`func (o *SummaryObject) GetOrigin() interface{}`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SummaryObject) GetOriginOk() (*interface{}, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SummaryObject) SetOrigin(v interface{})`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *SummaryObject) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *SummaryObject) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *SummaryObject) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetSourceStatus

`func (o *SummaryObject) GetSourceStatus() string`

GetSourceStatus returns the SourceStatus field if non-nil, zero value otherwise.

### GetSourceStatusOk

`func (o *SummaryObject) GetSourceStatusOk() (*string, bool)`

GetSourceStatusOk returns a tuple with the SourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStatus

`func (o *SummaryObject) SetSourceStatus(v string)`

SetSourceStatus sets SourceStatus field to given value.

### HasSourceStatus

`func (o *SummaryObject) HasSourceStatus() bool`

HasSourceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


