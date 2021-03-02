# LiveStreamSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | Pointer to [**LiveStreamSessionSession**](live_stream_session_session.md) |  | [optional] 
**Location** | Pointer to [**LiveStreamSessionLocation**](live_stream_session_location.md) |  | [optional] 
**Referrer** | Pointer to [**LiveStreamSessionReferrer**](live_stream_session_referrer.md) |  | [optional] 
**Device** | Pointer to [**LiveStreamSessionDevice**](live_stream_session_device.md) |  | [optional] 
**Os** | Pointer to [**VideoSessionOs**](video_session_os.md) |  | [optional] 
**Client** | Pointer to [**LiveStreamSessionClient**](live_stream_session_client.md) |  | [optional] 

## Methods

### NewLiveStreamSession

`func NewLiveStreamSession() *LiveStreamSession`

NewLiveStreamSession instantiates a new LiveStreamSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamSessionWithDefaults

`func NewLiveStreamSessionWithDefaults() *LiveStreamSession`

NewLiveStreamSessionWithDefaults instantiates a new LiveStreamSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *LiveStreamSession) GetSession() LiveStreamSessionSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *LiveStreamSession) GetSessionOk() (*LiveStreamSessionSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *LiveStreamSession) SetSession(v LiveStreamSessionSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *LiveStreamSession) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetLocation

`func (o *LiveStreamSession) GetLocation() LiveStreamSessionLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LiveStreamSession) GetLocationOk() (*LiveStreamSessionLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LiveStreamSession) SetLocation(v LiveStreamSessionLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LiveStreamSession) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetReferrer

`func (o *LiveStreamSession) GetReferrer() LiveStreamSessionReferrer`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *LiveStreamSession) GetReferrerOk() (*LiveStreamSessionReferrer, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *LiveStreamSession) SetReferrer(v LiveStreamSessionReferrer)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *LiveStreamSession) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetDevice

`func (o *LiveStreamSession) GetDevice() LiveStreamSessionDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *LiveStreamSession) GetDeviceOk() (*LiveStreamSessionDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *LiveStreamSession) SetDevice(v LiveStreamSessionDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *LiveStreamSession) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetOs

`func (o *LiveStreamSession) GetOs() VideoSessionOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *LiveStreamSession) GetOsOk() (*VideoSessionOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *LiveStreamSession) SetOs(v VideoSessionOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *LiveStreamSession) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetClient

`func (o *LiveStreamSession) GetClient() LiveStreamSessionClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *LiveStreamSession) GetClientOk() (*LiveStreamSessionClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *LiveStreamSession) SetClient(v LiveStreamSessionClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *LiveStreamSession) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


