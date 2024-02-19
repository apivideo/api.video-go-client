# VideoStatusIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | There are four possible statuses depending on how you provide a video file: - &#x60;uploading&#x60; - the API is gathering the video source file from an upload. - &#x60;uploaded&#x60; - the video file is fully uploaded. - &#x60;ingesting&#x60; - the API is gathering the video source file from either a URL, or from cloning. - &#x60;ingested&#x60; - the video file is fully stored.  | [optional] 
**Filesize** | Pointer to **NullableInt32** | The size of your file in bytes. | [optional] 
**ReceivedBytes** | Pointer to [**[]BytesRange**](BytesRange.md) | The total number of bytes received, listed for each chunk of the upload. | [optional] 
**ReceivedParts** | Pointer to [**VideoStatusIngestReceivedParts**](VideoStatusIngestReceivedParts.md) |  | [optional] 

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

### SetFilesizeNil

`func (o *VideoStatusIngest) SetFilesizeNil(b bool)`

 SetFilesizeNil sets the value for Filesize to be an explicit nil

### UnsetFilesize
`func (o *VideoStatusIngest) UnsetFilesize()`

UnsetFilesize ensures that no value is present for Filesize, not even an explicit nil
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

### GetReceivedParts

`func (o *VideoStatusIngest) GetReceivedParts() VideoStatusIngestReceivedParts`

GetReceivedParts returns the ReceivedParts field if non-nil, zero value otherwise.

### GetReceivedPartsOk

`func (o *VideoStatusIngest) GetReceivedPartsOk() (*VideoStatusIngestReceivedParts, bool)`

GetReceivedPartsOk returns a tuple with the ReceivedParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedParts

`func (o *VideoStatusIngest) SetReceivedParts(v VideoStatusIngestReceivedParts)`

SetReceivedParts sets ReceivedParts field to given value.

### HasReceivedParts

`func (o *VideoStatusIngest) HasReceivedParts() bool`

HasReceivedParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


