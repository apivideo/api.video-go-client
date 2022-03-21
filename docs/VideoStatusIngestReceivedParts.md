# VideoStatusIngestReceivedParts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parts** | Pointer to **[]int32** | The parts that have been uploaded, ordered. For example, if part 2 was sent before part 1, and both have been uploaded, the output will be [1, 2]. | [optional] 
**Total** | Pointer to **NullableInt32** | Contains the number of expected parts. The total will be listed as \&quot;null\&quot; until the total number of parts is known. | [optional] 

## Methods

### NewVideoStatusIngestReceivedParts

`func NewVideoStatusIngestReceivedParts() *VideoStatusIngestReceivedParts`

NewVideoStatusIngestReceivedParts instantiates a new VideoStatusIngestReceivedParts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoStatusIngestReceivedPartsWithDefaults

`func NewVideoStatusIngestReceivedPartsWithDefaults() *VideoStatusIngestReceivedParts`

NewVideoStatusIngestReceivedPartsWithDefaults instantiates a new VideoStatusIngestReceivedParts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParts

`func (o *VideoStatusIngestReceivedParts) GetParts() []int32`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *VideoStatusIngestReceivedParts) GetPartsOk() (*[]int32, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *VideoStatusIngestReceivedParts) SetParts(v []int32)`

SetParts sets Parts field to given value.

### HasParts

`func (o *VideoStatusIngestReceivedParts) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetTotal

`func (o *VideoStatusIngestReceivedParts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VideoStatusIngestReceivedParts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VideoStatusIngestReceivedParts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *VideoStatusIngestReceivedParts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *VideoStatusIngestReceivedParts) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *VideoStatusIngestReceivedParts) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


