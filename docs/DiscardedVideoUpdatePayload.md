# DiscardedVideoUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discarded** | Pointer to **bool** | Use this parameter to restore a discarded video when you have the Video Restore feature enabled. This parameter only accepts &#x60;false&#x60; as a value! | [optional] 

## Methods

### NewDiscardedVideoUpdatePayload

`func NewDiscardedVideoUpdatePayload() *DiscardedVideoUpdatePayload`

NewDiscardedVideoUpdatePayload instantiates a new DiscardedVideoUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscardedVideoUpdatePayloadWithDefaults

`func NewDiscardedVideoUpdatePayloadWithDefaults() *DiscardedVideoUpdatePayload`

NewDiscardedVideoUpdatePayloadWithDefaults instantiates a new DiscardedVideoUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscarded

`func (o *DiscardedVideoUpdatePayload) GetDiscarded() bool`

GetDiscarded returns the Discarded field if non-nil, zero value otherwise.

### GetDiscardedOk

`func (o *DiscardedVideoUpdatePayload) GetDiscardedOk() (*bool, bool)`

GetDiscardedOk returns a tuple with the Discarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscarded

`func (o *DiscardedVideoUpdatePayload) SetDiscarded(v bool)`

SetDiscarded sets Discarded field to given value.

### HasDiscarded

`func (o *DiscardedVideoUpdatePayload) HasDiscarded() bool`

HasDiscarded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


