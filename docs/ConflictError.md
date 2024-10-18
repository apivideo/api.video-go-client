# ConflictError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A link to the error documentation. | [optional] 
**Title** | Pointer to **string** | A description of the error that occurred. | [optional] 
**Name** | Pointer to **string** | The name of the parameter that caused the error. | [optional] 
**Status** | Pointer to **int32** | The HTTP status code. | [optional] 
**Detail** | Pointer to **string** | A solution for the error. | [optional] 

## Methods

### NewConflictError

`func NewConflictError() *ConflictError`

NewConflictError instantiates a new ConflictError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConflictErrorWithDefaults

`func NewConflictErrorWithDefaults() *ConflictError`

NewConflictErrorWithDefaults instantiates a new ConflictError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConflictError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConflictError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConflictError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConflictError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *ConflictError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConflictError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConflictError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConflictError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetName

`func (o *ConflictError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConflictError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConflictError) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConflictError) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ConflictError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConflictError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConflictError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConflictError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *ConflictError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ConflictError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ConflictError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ConflictError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


