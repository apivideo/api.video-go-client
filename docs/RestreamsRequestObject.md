# RestreamsRequestObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Use this parameter to define a name for the restream destination. | 
**ServerUrl** | **string** | Use this parameter to set the &#x60;RTMPS&#x60; or &#x60;RTMP&#x60; server URL of the restream destination. | 
**StreamKey** | **string** | Use this parameter to provide the unique key of the live stream that you want to restream. | 

## Methods

### NewRestreamsRequestObject

`func NewRestreamsRequestObject(name string, serverUrl string, streamKey string, ) *RestreamsRequestObject`

NewRestreamsRequestObject instantiates a new RestreamsRequestObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestreamsRequestObjectWithDefaults

`func NewRestreamsRequestObjectWithDefaults() *RestreamsRequestObject`

NewRestreamsRequestObjectWithDefaults instantiates a new RestreamsRequestObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RestreamsRequestObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestreamsRequestObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestreamsRequestObject) SetName(v string)`

SetName sets Name field to given value.


### GetServerUrl

`func (o *RestreamsRequestObject) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *RestreamsRequestObject) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *RestreamsRequestObject) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.


### GetStreamKey

`func (o *RestreamsRequestObject) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *RestreamsRequestObject) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *RestreamsRequestObject) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


