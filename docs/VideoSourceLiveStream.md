# VideoSourceLiveStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LiveStreamId** | Pointer to **string** | The unique identifier for the live stream. | [optional] 
**Links** | Pointer to [**[]VideoSourceLiveStreamLink**](VideoSourceLiveStreamLink.md) |  | [optional] 

## Methods

### NewVideoSourceLiveStream

`func NewVideoSourceLiveStream() *VideoSourceLiveStream`

NewVideoSourceLiveStream instantiates a new VideoSourceLiveStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSourceLiveStreamWithDefaults

`func NewVideoSourceLiveStreamWithDefaults() *VideoSourceLiveStream`

NewVideoSourceLiveStreamWithDefaults instantiates a new VideoSourceLiveStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLiveStreamId

`func (o *VideoSourceLiveStream) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *VideoSourceLiveStream) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *VideoSourceLiveStream) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *VideoSourceLiveStream) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### GetLinks

`func (o *VideoSourceLiveStream) GetLinks() []VideoSourceLiveStreamLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VideoSourceLiveStream) GetLinksOk() (*[]VideoSourceLiveStreamLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VideoSourceLiveStream) SetLinks(v []VideoSourceLiveStreamLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VideoSourceLiveStream) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


