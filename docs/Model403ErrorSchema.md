# Model403ErrorSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A link to the error documentation. | [optional] 
**Title** | Pointer to **string** | A description of the error that occurred. | [optional] 
**Name** | Pointer to **NullableString** | The name of the parameter that caused the error. | [optional] 
**Status** | Pointer to **int32** | The HTTP status code. | [optional] 

## Methods

### NewModel403ErrorSchema

`func NewModel403ErrorSchema() *Model403ErrorSchema`

NewModel403ErrorSchema instantiates a new Model403ErrorSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel403ErrorSchemaWithDefaults

`func NewModel403ErrorSchemaWithDefaults() *Model403ErrorSchema`

NewModel403ErrorSchemaWithDefaults instantiates a new Model403ErrorSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Model403ErrorSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Model403ErrorSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Model403ErrorSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Model403ErrorSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *Model403ErrorSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Model403ErrorSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Model403ErrorSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Model403ErrorSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetName

`func (o *Model403ErrorSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Model403ErrorSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Model403ErrorSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Model403ErrorSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Model403ErrorSchema) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Model403ErrorSchema) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *Model403ErrorSchema) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Model403ErrorSchema) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Model403ErrorSchema) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Model403ErrorSchema) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


