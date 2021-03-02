# VideoSessionSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** | The unique identifier for the session that you can use to track what happens during it. | [optional] 
**LoadedAt** | Pointer to **time.Time** | When the video session started, presented in ISO-8601 format. | [optional] 
**EndedAt** | Pointer to **time.Time** | When the video session ended, presented in ISO-8601 format. | [optional] 

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

`func (o *VideoSessionSession) GetLoadedAt() time.Time`

GetLoadedAt returns the LoadedAt field if non-nil, zero value otherwise.

### GetLoadedAtOk

`func (o *VideoSessionSession) GetLoadedAtOk() (*time.Time, bool)`

GetLoadedAtOk returns a tuple with the LoadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedAt

`func (o *VideoSessionSession) SetLoadedAt(v time.Time)`

SetLoadedAt sets LoadedAt field to given value.

### HasLoadedAt

`func (o *VideoSessionSession) HasLoadedAt() bool`

HasLoadedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *VideoSessionSession) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *VideoSessionSession) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *VideoSessionSession) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *VideoSessionSession) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


