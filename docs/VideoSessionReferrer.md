# VideoSessionReferrer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** | The link the viewer used to reach the video session. | [optional] 
**Medium** | Pointer to **string** | How they arrived at the site, for example organic or paid. Organic meaning they found it themselves and paid meaning they followed a link from an advertisement. | [optional] 
**Source** | Pointer to **string** | The source the referrer came from to the video session. For example if they searched through google to find the stream. | [optional] 
**SearchTerm** | Pointer to **string** | The search term they typed to arrive at the video session. | [optional] 

## Methods

### NewVideoSessionReferrer

`func NewVideoSessionReferrer() *VideoSessionReferrer`

NewVideoSessionReferrer instantiates a new VideoSessionReferrer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSessionReferrerWithDefaults

`func NewVideoSessionReferrerWithDefaults() *VideoSessionReferrer`

NewVideoSessionReferrerWithDefaults instantiates a new VideoSessionReferrer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *VideoSessionReferrer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoSessionReferrer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoSessionReferrer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoSessionReferrer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VideoSessionReferrer) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VideoSessionReferrer) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetMedium

`func (o *VideoSessionReferrer) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *VideoSessionReferrer) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *VideoSessionReferrer) SetMedium(v string)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *VideoSessionReferrer) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetSource

`func (o *VideoSessionReferrer) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VideoSessionReferrer) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VideoSessionReferrer) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VideoSessionReferrer) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSearchTerm

`func (o *VideoSessionReferrer) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *VideoSessionReferrer) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *VideoSessionReferrer) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *VideoSessionReferrer) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


