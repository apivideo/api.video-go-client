# LiveStreamSessionLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The country of the viewer of the live stream. | [optional] 
**City** | Pointer to **string** | The city of the viewer of the live stream. | [optional] 

## Methods

### NewLiveStreamSessionLocation

`func NewLiveStreamSessionLocation() *LiveStreamSessionLocation`

NewLiveStreamSessionLocation instantiates a new LiveStreamSessionLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamSessionLocationWithDefaults

`func NewLiveStreamSessionLocationWithDefaults() *LiveStreamSessionLocation`

NewLiveStreamSessionLocationWithDefaults instantiates a new LiveStreamSessionLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *LiveStreamSessionLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LiveStreamSessionLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LiveStreamSessionLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LiveStreamSessionLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *LiveStreamSessionLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LiveStreamSessionLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LiveStreamSessionLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *LiveStreamSessionLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


