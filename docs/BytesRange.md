# BytesRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int32** | The starting point for the range of bytes for a chunk of a video. | [optional] 
**To** | Pointer to **int32** | The ending point for the range of bytes for a chunk of a video. | [optional] 
**Total** | Pointer to **int32** | The total number of expected bytes. | [optional] 

## Methods

### NewBytesRange

`func NewBytesRange() *BytesRange`

NewBytesRange instantiates a new BytesRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBytesRangeWithDefaults

`func NewBytesRangeWithDefaults() *BytesRange`

NewBytesRangeWithDefaults instantiates a new BytesRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *BytesRange) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BytesRange) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BytesRange) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BytesRange) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *BytesRange) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BytesRange) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BytesRange) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *BytesRange) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotal

`func (o *BytesRange) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BytesRange) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BytesRange) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BytesRange) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


