# VideoSessionClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the browser used to view the video session. | [optional] 
**Version** | Pointer to **string** | The version of the browser used to view the video session. | [optional] 
**Type** | Pointer to **string** | The type of client used to view the video session. | [optional] 

## Methods

### NewVideoSessionClient

`func NewVideoSessionClient() *VideoSessionClient`

NewVideoSessionClient instantiates a new VideoSessionClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoSessionClientWithDefaults

`func NewVideoSessionClientWithDefaults() *VideoSessionClient`

NewVideoSessionClientWithDefaults instantiates a new VideoSessionClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VideoSessionClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoSessionClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoSessionClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VideoSessionClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *VideoSessionClient) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VideoSessionClient) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VideoSessionClient) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VideoSessionClient) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *VideoSessionClient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoSessionClient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoSessionClient) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VideoSessionClient) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


