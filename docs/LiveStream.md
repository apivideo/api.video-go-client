# LiveStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LiveStreamId** | Pointer to **string** | The unique identifier for the live stream. Live stream IDs begin with \&quot;li.\&quot; | [optional] 
**Name** | Pointer to **string** | The name of your live stream. | [optional] 
**StreamKey** | Pointer to **string** | The unique, private stream key that you use to begin streaming. | [optional] 
**Record** | Pointer to **bool** | Whether you are recording or not. | [optional] 
**Public** | Pointer to **bool** | BETA FEATURE Please limit all public &#x3D; false (\&quot;private\&quot;) livestreams to 3,000 users. Whether your video can be viewed by everyone, or requires authentication to see it. A setting of false will require a unique token for each view. | [optional] 
**Assets** | Pointer to [**LiveStreamAssets**](live_stream_assets.md) |  | [optional] 
**PlayerId** | Pointer to **string** | The unique identifier for the player. | [optional] 
**Broadcasting** | Pointer to **bool** | Whether or not you are broadcasting the live video you recorded for others to see. True means you are broadcasting to viewers, false means you are not. | [optional] 

## Methods

### NewLiveStream

`func NewLiveStream() *LiveStream`

NewLiveStream instantiates a new LiveStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamWithDefaults

`func NewLiveStreamWithDefaults() *LiveStream`

NewLiveStreamWithDefaults instantiates a new LiveStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLiveStreamId

`func (o *LiveStream) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *LiveStream) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *LiveStream) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *LiveStream) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### GetName

`func (o *LiveStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LiveStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LiveStream) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LiveStream) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStreamKey

`func (o *LiveStream) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *LiveStream) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *LiveStream) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.

### HasStreamKey

`func (o *LiveStream) HasStreamKey() bool`

HasStreamKey returns a boolean if a field has been set.

### GetRecord

`func (o *LiveStream) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *LiveStream) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *LiveStream) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *LiveStream) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetPublic

`func (o *LiveStream) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *LiveStream) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *LiveStream) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *LiveStream) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetAssets

`func (o *LiveStream) GetAssets() LiveStreamAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *LiveStream) GetAssetsOk() (*LiveStreamAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *LiveStream) SetAssets(v LiveStreamAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *LiveStream) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPlayerId

`func (o *LiveStream) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *LiveStream) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *LiveStream) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *LiveStream) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetBroadcasting

`func (o *LiveStream) GetBroadcasting() bool`

GetBroadcasting returns the Broadcasting field if non-nil, zero value otherwise.

### GetBroadcastingOk

`func (o *LiveStream) GetBroadcastingOk() (*bool, bool)`

GetBroadcastingOk returns a tuple with the Broadcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasting

`func (o *LiveStream) SetBroadcasting(v bool)`

SetBroadcasting sets Broadcasting field to given value.

### HasBroadcasting

`func (o *LiveStream) HasBroadcasting() bool`

HasBroadcasting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


