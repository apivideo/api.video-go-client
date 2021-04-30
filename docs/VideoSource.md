# VideoSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **string** | The URL where the video is stored. | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**LiveStream** | Pointer to [**VideoSourceLiveStream**](video-source-live-stream.md) |  | [optional] 

## Methods

### NewVideoSource

`func NewVideoSource() *VideoSource`

NewVideoSource instantiates a new VideoSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSourceWithDefaults

`func NewVideoSourceWithDefaults() *VideoSource`

NewVideoSourceWithDefaults instantiates a new VideoSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *VideoSource) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *VideoSource) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *VideoSource) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *VideoSource) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetType

`func (o *VideoSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VideoSource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLiveStream

`func (o *VideoSource) GetLiveStream() VideoSourceLiveStream`

GetLiveStream returns the LiveStream field if non-nil, zero value otherwise.

### GetLiveStreamOk

`func (o *VideoSource) GetLiveStreamOk() (*VideoSourceLiveStream, bool)`

GetLiveStreamOk returns a tuple with the LiveStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStream

`func (o *VideoSource) SetLiveStream(v VideoSourceLiveStream)`

SetLiveStream sets LiveStream field to given value.

### HasLiveStream

`func (o *VideoSource) HasLiveStream() bool`

HasLiveStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


