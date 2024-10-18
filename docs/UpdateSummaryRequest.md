# UpdateSummaryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A video title, based on the contents of the video. | [optional] 
**Abstract** | Pointer to **string** | A short outline of the contents of the video. | [optional] 
**Takeaways** | Pointer to **[]string** | A list of 3 key points from the video, in chronological order. | [optional] 

## Methods

### NewUpdateSummaryRequest

`func NewUpdateSummaryRequest() *UpdateSummaryRequest`

NewUpdateSummaryRequest instantiates a new UpdateSummaryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSummaryRequestWithDefaults

`func NewUpdateSummaryRequestWithDefaults() *UpdateSummaryRequest`

NewUpdateSummaryRequestWithDefaults instantiates a new UpdateSummaryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateSummaryRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateSummaryRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateSummaryRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateSummaryRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAbstract

`func (o *UpdateSummaryRequest) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *UpdateSummaryRequest) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *UpdateSummaryRequest) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.

### HasAbstract

`func (o *UpdateSummaryRequest) HasAbstract() bool`

HasAbstract returns a boolean if a field has been set.

### GetTakeaways

`func (o *UpdateSummaryRequest) GetTakeaways() []string`

GetTakeaways returns the Takeaways field if non-nil, zero value otherwise.

### GetTakeawaysOk

`func (o *UpdateSummaryRequest) GetTakeawaysOk() (*[]string, bool)`

GetTakeawaysOk returns a tuple with the Takeaways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeaways

`func (o *UpdateSummaryRequest) SetTakeaways(v []string)`

SetTakeaways sets Takeaways field to given value.

### HasTakeaways

`func (o *UpdateSummaryRequest) HasTakeaways() bool`

HasTakeaways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


