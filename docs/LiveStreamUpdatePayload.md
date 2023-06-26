# LiveStreamUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name you want to use for your live stream. | [optional] 
**Public** | Pointer to **bool** | Whether your video can be viewed by everyone, or requires authentication to see it. A setting of false will require a unique token for each view. Learn more about the Private Video feature [here](https://docs.api.video/docs/private-videos). | [optional] 
**Record** | Pointer to **bool** | Use this to indicate whether you want the recording on or off. On is true, off is false. | [optional] 
**PlayerId** | Pointer to **string** | The unique ID for the player associated with a live stream that you want to update. | [optional] 
**Restreams** | Pointer to [**[]RestreamsRequestObject**](RestreamsRequestObject.md) | Use this parameter to add, edit, or remove RTMP services where you want to restream a live stream. The list can only contain up to 5 destinations. This operation updates all restream destinations in the same request. If you do not want to modify an existing restream destination, you need to include it in your request, otherwise it is removed. | [optional] 

## Methods

### NewLiveStreamUpdatePayload

`func NewLiveStreamUpdatePayload() *LiveStreamUpdatePayload`

NewLiveStreamUpdatePayload instantiates a new LiveStreamUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamUpdatePayloadWithDefaults

`func NewLiveStreamUpdatePayloadWithDefaults() *LiveStreamUpdatePayload`

NewLiveStreamUpdatePayloadWithDefaults instantiates a new LiveStreamUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LiveStreamUpdatePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LiveStreamUpdatePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LiveStreamUpdatePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LiveStreamUpdatePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublic

`func (o *LiveStreamUpdatePayload) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *LiveStreamUpdatePayload) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *LiveStreamUpdatePayload) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *LiveStreamUpdatePayload) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRecord

`func (o *LiveStreamUpdatePayload) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *LiveStreamUpdatePayload) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *LiveStreamUpdatePayload) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *LiveStreamUpdatePayload) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetPlayerId

`func (o *LiveStreamUpdatePayload) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *LiveStreamUpdatePayload) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *LiveStreamUpdatePayload) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *LiveStreamUpdatePayload) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetRestreams

`func (o *LiveStreamUpdatePayload) GetRestreams() []RestreamsRequestObject`

GetRestreams returns the Restreams field if non-nil, zero value otherwise.

### GetRestreamsOk

`func (o *LiveStreamUpdatePayload) GetRestreamsOk() (*[]RestreamsRequestObject, bool)`

GetRestreamsOk returns a tuple with the Restreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestreams

`func (o *LiveStreamUpdatePayload) SetRestreams(v []RestreamsRequestObject)`

SetRestreams sets Restreams field to given value.

### HasRestreams

`func (o *LiveStreamUpdatePayload) HasRestreams() bool`

HasRestreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


