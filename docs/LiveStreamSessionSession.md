# LiveStreamSessionSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** | A unique identifier for your session. You can use this to track what happens during a specific session. | [optional] 
**LoadedAt** | Pointer to **string** | When the session started, with the date and time presented in ISO-8601 format. | [optional] 
**EndedAt** | Pointer to **string** | When the session ended, with the date and time presented in ISO-8601 format. | [optional] 

## Methods

### NewLiveStreamSessionSession

`func NewLiveStreamSessionSession() *LiveStreamSessionSession`

NewLiveStreamSessionSession instantiates a new LiveStreamSessionSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamSessionSessionWithDefaults

`func NewLiveStreamSessionSessionWithDefaults() *LiveStreamSessionSession`

NewLiveStreamSessionSessionWithDefaults instantiates a new LiveStreamSessionSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *LiveStreamSessionSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *LiveStreamSessionSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *LiveStreamSessionSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *LiveStreamSessionSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetLoadedAt

`func (o *LiveStreamSessionSession) GetLoadedAt() string`

GetLoadedAt returns the LoadedAt field if non-nil, zero value otherwise.

### GetLoadedAtOk

`func (o *LiveStreamSessionSession) GetLoadedAtOk() (*string, bool)`

GetLoadedAtOk returns a tuple with the LoadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedAt

`func (o *LiveStreamSessionSession) SetLoadedAt(v string)`

SetLoadedAt sets LoadedAt field to given value.

### HasLoadedAt

`func (o *LiveStreamSessionSession) HasLoadedAt() bool`

HasLoadedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *LiveStreamSessionSession) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *LiveStreamSessionSession) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *LiveStreamSessionSession) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *LiveStreamSessionSession) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


