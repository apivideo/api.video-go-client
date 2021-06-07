# VideoStatusIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | There are three possible ingest statuses. missing - you are missing information required to ingest the video. uploading - the video is in the process of being uploaded. uploaded - the video is ready for use. | [optional] 
**Filesize** | Pointer to **int32** | The size of your file in bytes. | [optional] 
**ReceivedBytes** | Pointer to [**[]BytesRange**](BytesRange.md) | The total number of bytes received, listed for each chunk of the upload. | [optional] 

## Methods

### NewVideoStatusIngest

`func NewVideoStatusIngest() *VideoStatusIngest`

NewVideoStatusIngest instantiates a new VideoStatusIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoStatusIngestWithDefaults

`func NewVideoStatusIngestWithDefaults() *VideoStatusIngest`

NewVideoStatusIngestWithDefaults instantiates a new VideoStatusIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VideoStatusIngest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoStatusIngest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoStatusIngest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VideoStatusIngest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFilesize

`func (o *VideoStatusIngest) GetFilesize() int32`

GetFilesize returns the Filesize field if non-nil, zero value otherwise.

### GetFilesizeOk

`func (o *VideoStatusIngest) GetFilesizeOk() (*int32, bool)`

GetFilesizeOk returns a tuple with the Filesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesize

`func (o *VideoStatusIngest) SetFilesize(v int32)`

SetFilesize sets Filesize field to given value.

### HasFilesize

`func (o *VideoStatusIngest) HasFilesize() bool`

HasFilesize returns a boolean if a field has been set.

### GetReceivedBytes

`func (o *VideoStatusIngest) GetReceivedBytes() []BytesRange`

GetReceivedBytes returns the ReceivedBytes field if non-nil, zero value otherwise.

### GetReceivedBytesOk

`func (o *VideoStatusIngest) GetReceivedBytesOk() (*[]BytesRange, bool)`

GetReceivedBytesOk returns a tuple with the ReceivedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedBytes

`func (o *VideoStatusIngest) SetReceivedBytes(v []BytesRange)`

SetReceivedBytes sets ReceivedBytes field to given value.

### HasReceivedBytes

`func (o *VideoStatusIngest) HasReceivedBytes() bool`

HasReceivedBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


