# LiveStreamSessionDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | What the type is like desktop, laptop, mobile. | [optional] 
**Vendor** | Pointer to **string** | If known, what the brand of the device is, like Apple, Dell, etc. | [optional] 
**Model** | Pointer to **string** | The specific model of the device, if known. | [optional] 

## Methods

### NewLiveStreamSessionDevice

`func NewLiveStreamSessionDevice() *LiveStreamSessionDevice`

NewLiveStreamSessionDevice instantiates a new LiveStreamSessionDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamSessionDeviceWithDefaults

`func NewLiveStreamSessionDeviceWithDefaults() *LiveStreamSessionDevice`

NewLiveStreamSessionDeviceWithDefaults instantiates a new LiveStreamSessionDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LiveStreamSessionDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LiveStreamSessionDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LiveStreamSessionDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LiveStreamSessionDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *LiveStreamSessionDevice) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *LiveStreamSessionDevice) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *LiveStreamSessionDevice) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *LiveStreamSessionDevice) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetModel

`func (o *LiveStreamSessionDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *LiveStreamSessionDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *LiveStreamSessionDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *LiveStreamSessionDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


