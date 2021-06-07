# VideoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingest** | Pointer to [**VideoStatusIngest**](video-status-ingest.md) |  | [optional] 
**Encoding** | Pointer to [**VideoStatusEncoding**](video-status-encoding.md) |  | [optional] 

## Methods

### NewVideoStatus

`func NewVideoStatus() *VideoStatus`

NewVideoStatus instantiates a new VideoStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoStatusWithDefaults

`func NewVideoStatusWithDefaults() *VideoStatus`

NewVideoStatusWithDefaults instantiates a new VideoStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngest

`func (o *VideoStatus) GetIngest() VideoStatusIngest`

GetIngest returns the Ingest field if non-nil, zero value otherwise.

### GetIngestOk

`func (o *VideoStatus) GetIngestOk() (*VideoStatusIngest, bool)`

GetIngestOk returns a tuple with the Ingest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngest

`func (o *VideoStatus) SetIngest(v VideoStatusIngest)`

SetIngest sets Ingest field to given value.

### HasIngest

`func (o *VideoStatus) HasIngest() bool`

HasIngest returns a boolean if a field has been set.

### GetEncoding

`func (o *VideoStatus) GetEncoding() VideoStatusEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *VideoStatus) GetEncodingOk() (*VideoStatusEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *VideoStatus) SetEncoding(v VideoStatusEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *VideoStatus) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


