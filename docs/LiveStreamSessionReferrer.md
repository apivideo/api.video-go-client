# LiveStreamSessionReferrer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The website the viewer of the live stream was referred to in order to view the live stream. | [optional] 
**Medium** | Pointer to **string** | The type of search that brought the viewer to the live stream. Organic would be they found it on their own, paid would be they found it via an advertisement. | [optional] 
**Source** | Pointer to **string** | Where the viewer came from to see the live stream (usually where they searched from). | [optional] 
**SearchTerm** | Pointer to **string** | What term they searched for that led them to the live stream. | [optional] 

## Methods

### NewLiveStreamSessionReferrer

`func NewLiveStreamSessionReferrer() *LiveStreamSessionReferrer`

NewLiveStreamSessionReferrer instantiates a new LiveStreamSessionReferrer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamSessionReferrerWithDefaults

`func NewLiveStreamSessionReferrerWithDefaults() *LiveStreamSessionReferrer`

NewLiveStreamSessionReferrerWithDefaults instantiates a new LiveStreamSessionReferrer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *LiveStreamSessionReferrer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LiveStreamSessionReferrer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LiveStreamSessionReferrer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LiveStreamSessionReferrer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMedium

`func (o *LiveStreamSessionReferrer) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *LiveStreamSessionReferrer) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *LiveStreamSessionReferrer) SetMedium(v string)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *LiveStreamSessionReferrer) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetSource

`func (o *LiveStreamSessionReferrer) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LiveStreamSessionReferrer) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LiveStreamSessionReferrer) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LiveStreamSessionReferrer) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSearchTerm

`func (o *LiveStreamSessionReferrer) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *LiveStreamSessionReferrer) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *LiveStreamSessionReferrer) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *LiveStreamSessionReferrer) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


