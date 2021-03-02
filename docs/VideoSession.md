# VideoSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | Pointer to [**VideoSessionSession**](video_session_session.md) |  | [optional] 
**Location** | Pointer to [**VideoSessionLocation**](video_session_location.md) |  | [optional] 
**Referrer** | Pointer to [**VideoSessionReferrer**](video_session_referrer.md) |  | [optional] 
**Device** | Pointer to [**VideoSessionDevice**](video_session_device.md) |  | [optional] 
**Os** | Pointer to [**VideoSessionOs**](video_session_os.md) |  | [optional] 
**Client** | Pointer to [**VideoSessionClient**](video_session_client.md) |  | [optional] 

## Methods

### NewVideoSession

`func NewVideoSession() *VideoSession`

NewVideoSession instantiates a new VideoSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSessionWithDefaults

`func NewVideoSessionWithDefaults() *VideoSession`

NewVideoSessionWithDefaults instantiates a new VideoSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *VideoSession) GetSession() VideoSessionSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *VideoSession) GetSessionOk() (*VideoSessionSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *VideoSession) SetSession(v VideoSessionSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *VideoSession) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetLocation

`func (o *VideoSession) GetLocation() VideoSessionLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VideoSession) GetLocationOk() (*VideoSessionLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VideoSession) SetLocation(v VideoSessionLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VideoSession) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetReferrer

`func (o *VideoSession) GetReferrer() VideoSessionReferrer`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *VideoSession) GetReferrerOk() (*VideoSessionReferrer, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *VideoSession) SetReferrer(v VideoSessionReferrer)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *VideoSession) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetDevice

`func (o *VideoSession) GetDevice() VideoSessionDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VideoSession) GetDeviceOk() (*VideoSessionDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VideoSession) SetDevice(v VideoSessionDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VideoSession) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetOs

`func (o *VideoSession) GetOs() VideoSessionOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *VideoSession) GetOsOk() (*VideoSessionOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *VideoSession) SetOs(v VideoSessionOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *VideoSession) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetClient

`func (o *VideoSession) GetClient() VideoSessionClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *VideoSession) GetClientOk() (*VideoSessionClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *VideoSession) SetClient(v VideoSessionClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *VideoSession) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


