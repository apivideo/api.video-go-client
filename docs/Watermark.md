# Watermark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WatermarkId** | Pointer to **string** | The unique identifier of the watermark. | [optional] 
**CreatedAt** | Pointer to **string** | When the watermark was created, presented in ISO-8601 format. | [optional] 

## Methods

### NewWatermark

`func NewWatermark() *Watermark`

NewWatermark instantiates a new Watermark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkWithDefaults

`func NewWatermarkWithDefaults() *Watermark`

NewWatermarkWithDefaults instantiates a new Watermark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWatermarkId

`func (o *Watermark) GetWatermarkId() string`

GetWatermarkId returns the WatermarkId field if non-nil, zero value otherwise.

### GetWatermarkIdOk

`func (o *Watermark) GetWatermarkIdOk() (*string, bool)`

GetWatermarkIdOk returns a tuple with the WatermarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkId

`func (o *Watermark) SetWatermarkId(v string)`

SetWatermarkId sets WatermarkId field to given value.

### HasWatermarkId

`func (o *Watermark) HasWatermarkId() bool`

HasWatermarkId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Watermark) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Watermark) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Watermark) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Watermark) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


