# RestreamsResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Returns the name of a restream destination. | [optional] 
**ServerUrl** | Pointer to **string** | Returns the server URL of a restream destination. | [optional] 
**StreamKey** | Pointer to **string** | Returns the unique key of the live stream that is set up for restreaming. | [optional] 

## Methods

### NewRestreamsResponseObject

`func NewRestreamsResponseObject() *RestreamsResponseObject`

NewRestreamsResponseObject instantiates a new RestreamsResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestreamsResponseObjectWithDefaults

`func NewRestreamsResponseObjectWithDefaults() *RestreamsResponseObject`

NewRestreamsResponseObjectWithDefaults instantiates a new RestreamsResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestreamsResponseObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestreamsResponseObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestreamsResponseObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RestreamsResponseObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerUrl

`func (o *RestreamsResponseObject) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *RestreamsResponseObject) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *RestreamsResponseObject) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.

### HasServerUrl

`func (o *RestreamsResponseObject) HasServerUrl() bool`

HasServerUrl returns a boolean if a field has been set.

### GetStreamKey

`func (o *RestreamsResponseObject) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *RestreamsResponseObject) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *RestreamsResponseObject) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.

### HasStreamKey

`func (o *RestreamsResponseObject) HasStreamKey() bool`

HasStreamKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


