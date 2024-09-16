# PlayerSessionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Possible values are: ready, play, pause, resume, seek.backward, seek.forward, end | [optional] 
**EmittedAt** | Pointer to **string** | When an event occurred, presented in ATOM UTC format. | [optional] 
**At** | Pointer to **int32** |  | [optional] 
**From** | Pointer to **int32** |  | [optional] 
**To** | Pointer to **int32** |  | [optional] 

## Methods

### NewPlayerSessionEvent

`func NewPlayerSessionEvent() *PlayerSessionEvent`

NewPlayerSessionEvent instantiates a new PlayerSessionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerSessionEventWithDefaults

`func NewPlayerSessionEventWithDefaults() *PlayerSessionEvent`

NewPlayerSessionEventWithDefaults instantiates a new PlayerSessionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PlayerSessionEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlayerSessionEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlayerSessionEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlayerSessionEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmittedAt

`func (o *PlayerSessionEvent) GetEmittedAt() string`

GetEmittedAt returns the EmittedAt field if non-nil, zero value otherwise.

### GetEmittedAtOk

`func (o *PlayerSessionEvent) GetEmittedAtOk() (*string, bool)`

GetEmittedAtOk returns a tuple with the EmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmittedAt

`func (o *PlayerSessionEvent) SetEmittedAt(v string)`

SetEmittedAt sets EmittedAt field to given value.

### HasEmittedAt

`func (o *PlayerSessionEvent) HasEmittedAt() bool`

HasEmittedAt returns a boolean if a field has been set.

### GetAt

`func (o *PlayerSessionEvent) GetAt() int32`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *PlayerSessionEvent) GetAtOk() (*int32, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *PlayerSessionEvent) SetAt(v int32)`

SetAt sets At field to given value.

### HasAt

`func (o *PlayerSessionEvent) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetFrom

`func (o *PlayerSessionEvent) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PlayerSessionEvent) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PlayerSessionEvent) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PlayerSessionEvent) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *PlayerSessionEvent) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PlayerSessionEvent) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PlayerSessionEvent) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *PlayerSessionEvent) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


