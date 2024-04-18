# TooManyRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A link to the error documentation. | [optional] 
**Title** | Pointer to **string** | A description of the error that occurred. | [optional] 
**Status** | Pointer to **int32** | The HTTP status code. | [optional] 

## Methods

### NewTooManyRequests

`func NewTooManyRequests() *TooManyRequests`

NewTooManyRequests instantiates a new TooManyRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTooManyRequestsWithDefaults

`func NewTooManyRequestsWithDefaults() *TooManyRequests`

NewTooManyRequestsWithDefaults instantiates a new TooManyRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TooManyRequests) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TooManyRequests) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TooManyRequests) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TooManyRequests) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *TooManyRequests) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TooManyRequests) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TooManyRequests) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TooManyRequests) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *TooManyRequests) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TooManyRequests) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TooManyRequests) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TooManyRequests) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


