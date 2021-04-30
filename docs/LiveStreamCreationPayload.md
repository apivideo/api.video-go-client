# LiveStreamCreationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Add a name for your live stream here. | 
**Record** | Pointer to **bool** | Whether you are recording or not. True for record, false for not record. | [optional] [default to false]
**Public** | Pointer to **bool** | BETA FEATURE Please limit all public &#x3D; false (\&quot;private\&quot;) livestreams to 3,000 users. Whether your video can be viewed by everyone, or requires authentication to see it. A setting of false will require a unique token for each view. | [optional] 
**PlayerId** | Pointer to **string** | The unique identifier for the player. | [optional] 

## Methods

### NewLiveStreamCreationPayload

`func NewLiveStreamCreationPayload(name string, ) *LiveStreamCreationPayload`

NewLiveStreamCreationPayload instantiates a new LiveStreamCreationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamCreationPayloadWithDefaults

`func NewLiveStreamCreationPayloadWithDefaults() *LiveStreamCreationPayload`

NewLiveStreamCreationPayloadWithDefaults instantiates a new LiveStreamCreationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LiveStreamCreationPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LiveStreamCreationPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LiveStreamCreationPayload) SetName(v string)`

SetName sets Name field to given value.


### GetRecord

`func (o *LiveStreamCreationPayload) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *LiveStreamCreationPayload) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *LiveStreamCreationPayload) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *LiveStreamCreationPayload) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetPublic

`func (o *LiveStreamCreationPayload) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *LiveStreamCreationPayload) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *LiveStreamCreationPayload) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *LiveStreamCreationPayload) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPlayerId

`func (o *LiveStreamCreationPayload) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *LiveStreamCreationPayload) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *LiveStreamCreationPayload) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *LiveStreamCreationPayload) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


