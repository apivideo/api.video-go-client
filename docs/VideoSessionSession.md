# VideoSessionSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** | The unique identifier for the session that you can use to track what happens during it. | [optional] 
**LoadedAt** | Pointer to **string** | When the video session started, presented in ISO-8601 format. | [optional] 
**EndedAt** | Pointer to **string** | When the video session ended, presented in ISO-8601 format. | [optional] 
**Metadata** | Pointer to [**[]Metadata**](Metadata.md) | A list of key value pairs that you use to provide metadata for your video. These pairs can be made dynamic, allowing you to segment your audience. You can also just use the pairs as another way to tag and categorize your videos. | [optional] 

## Methods

### NewVideoSessionSession

`func NewVideoSessionSession() *VideoSessionSession`

NewVideoSessionSession instantiates a new VideoSessionSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSessionSessionWithDefaults

`func NewVideoSessionSessionWithDefaults() *VideoSessionSession`

NewVideoSessionSessionWithDefaults instantiates a new VideoSessionSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *VideoSessionSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *VideoSessionSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *VideoSessionSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *VideoSessionSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetLoadedAt

`func (o *VideoSessionSession) GetLoadedAt() string`

GetLoadedAt returns the LoadedAt field if non-nil, zero value otherwise.

### GetLoadedAtOk

`func (o *VideoSessionSession) GetLoadedAtOk() (*string, bool)`

GetLoadedAtOk returns a tuple with the LoadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedAt

`func (o *VideoSessionSession) SetLoadedAt(v string)`

SetLoadedAt sets LoadedAt field to given value.

### HasLoadedAt

`func (o *VideoSessionSession) HasLoadedAt() bool`

HasLoadedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *VideoSessionSession) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *VideoSessionSession) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *VideoSessionSession) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *VideoSessionSession) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *VideoSessionSession) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VideoSessionSession) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VideoSessionSession) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VideoSessionSession) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


