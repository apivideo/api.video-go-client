# BadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Problems** | Pointer to [**[]BadRequest**](BadRequest.md) |  | [optional] 

## Methods

### NewBadRequest

`func NewBadRequest() *BadRequest`

NewBadRequest instantiates a new BadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestWithDefaults

`func NewBadRequestWithDefaults() *BadRequest`

NewBadRequestWithDefaults instantiates a new BadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BadRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BadRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BadRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BadRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *BadRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BadRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BadRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BadRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetName

`func (o *BadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BadRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BadRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *BadRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BadRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BadRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BadRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProblems

`func (o *BadRequest) GetProblems() []BadRequest`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *BadRequest) GetProblemsOk() (*[]BadRequest, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *BadRequest) SetProblems(v []BadRequest)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *BadRequest) HasProblems() bool`

HasProblems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


