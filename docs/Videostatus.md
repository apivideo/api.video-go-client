# Videostatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingest** | Pointer to [**VideostatusIngest**](videostatus_ingest.md) |  | [optional] 
**Encoding** | Pointer to [**VideostatusEncoding**](videostatus_encoding.md) |  | [optional] 

## Methods

### NewVideostatus

`func NewVideostatus() *Videostatus`

NewVideostatus instantiates a new Videostatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideostatusWithDefaults

`func NewVideostatusWithDefaults() *Videostatus`

NewVideostatusWithDefaults instantiates a new Videostatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngest

`func (o *Videostatus) GetIngest() VideostatusIngest`

GetIngest returns the Ingest field if non-nil, zero value otherwise.

### GetIngestOk

`func (o *Videostatus) GetIngestOk() (*VideostatusIngest, bool)`

GetIngestOk returns a tuple with the Ingest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngest

`func (o *Videostatus) SetIngest(v VideostatusIngest)`

SetIngest sets Ingest field to given value.

### HasIngest

`func (o *Videostatus) HasIngest() bool`

HasIngest returns a boolean if a field has been set.

### GetEncoding

`func (o *Videostatus) GetEncoding() VideostatusEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *Videostatus) GetEncodingOk() (*VideostatusEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *Videostatus) SetEncoding(v VideostatusEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *Videostatus) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


