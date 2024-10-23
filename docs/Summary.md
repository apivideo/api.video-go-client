# Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryId** | Pointer to **string** | The unique identifier of the summary object. | [optional] 
**CreatedAt** | Pointer to **string** | Returns the date and time when the summary was created in ATOM date-time format. | [optional] 
**UpdatedAt** | Pointer to **string** | Returns the date and time when the summary was last updated in ATOM date-time format. | [optional] 
**VideoId** | Pointer to **string** | The unique identifier of the video object. | [optional] 
**Origin** | Pointer to **string** | Returns the origin of how the summary was created.  - &#x60;api&#x60; means that no summary was generated automatically. You can add summary manually using the &#x60;PATCH /summaries/{summaryId}/source&#x60; endpoint operation. Until this happens, &#x60;sourceStatus&#x60; returns &#x60;missing&#x60;. - &#x60;auto&#x60; means that the API generated the summary automatically. | [optional] 
**SourceStatus** | Pointer to **string** | Returns the current status of summary generation.  &#x60;missing&#x60;: the input for a summary is not present. &#x60;waiting&#x60; : the input video is being processed and a summary will be generated. &#x60;failed&#x60;: a technical issue prevented summary generation. &#x60;completed&#x60;: the summary is generated. &#x60;unprocessable&#x60;: the API rules the source video to be unsuitable for summary generation. An example for this is an input video that has no audio. | [optional] 

## Methods

### NewSummary

`func NewSummary() *Summary`

NewSummary instantiates a new Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryWithDefaults

`func NewSummaryWithDefaults() *Summary`

NewSummaryWithDefaults instantiates a new Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryId

`func (o *Summary) GetSummaryId() string`

GetSummaryId returns the SummaryId field if non-nil, zero value otherwise.

### GetSummaryIdOk

`func (o *Summary) GetSummaryIdOk() (*string, bool)`

GetSummaryIdOk returns a tuple with the SummaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryId

`func (o *Summary) SetSummaryId(v string)`

SetSummaryId sets SummaryId field to given value.

### HasSummaryId

`func (o *Summary) HasSummaryId() bool`

HasSummaryId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Summary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Summary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Summary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Summary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Summary) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Summary) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Summary) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Summary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVideoId

`func (o *Summary) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *Summary) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *Summary) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *Summary) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetOrigin

`func (o *Summary) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Summary) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Summary) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Summary) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetSourceStatus

`func (o *Summary) GetSourceStatus() string`

GetSourceStatus returns the SourceStatus field if non-nil, zero value otherwise.

### GetSourceStatusOk

`func (o *Summary) GetSourceStatusOk() (*string, bool)`

GetSourceStatusOk returns a tuple with the SourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStatus

`func (o *Summary) SetSourceStatus(v string)`

SetSourceStatus sets SourceStatus field to given value.

### HasSourceStatus

`func (o *Summary) HasSourceStatus() bool`

HasSourceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


