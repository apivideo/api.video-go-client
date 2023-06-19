# AdditionalBadRequestErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A link to the error documentation. | [optional] 
**Title** | Pointer to **string** | A description of the error that occurred. | [optional] 
**Name** | Pointer to **string** | The name of the parameter that caused the error. | [optional] 
**Status** | Pointer to **int32** | The HTTP status code. | [optional] 

## Methods

### NewAdditionalBadRequestErrors

`func NewAdditionalBadRequestErrors() *AdditionalBadRequestErrors`

NewAdditionalBadRequestErrors instantiates a new AdditionalBadRequestErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalBadRequestErrorsWithDefaults

`func NewAdditionalBadRequestErrorsWithDefaults() *AdditionalBadRequestErrors`

NewAdditionalBadRequestErrorsWithDefaults instantiates a new AdditionalBadRequestErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdditionalBadRequestErrors) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdditionalBadRequestErrors) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdditionalBadRequestErrors) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdditionalBadRequestErrors) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *AdditionalBadRequestErrors) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdditionalBadRequestErrors) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdditionalBadRequestErrors) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdditionalBadRequestErrors) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetName

`func (o *AdditionalBadRequestErrors) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdditionalBadRequestErrors) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdditionalBadRequestErrors) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdditionalBadRequestErrors) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AdditionalBadRequestErrors) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdditionalBadRequestErrors) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdditionalBadRequestErrors) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdditionalBadRequestErrors) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


