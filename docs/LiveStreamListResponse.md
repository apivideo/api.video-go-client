# LiveStreamListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]LiveStream**](LiveStream.md) |  | 
**Pagination** | [**Pagination**](pagination.md) |  | 

## Methods

### NewLiveStreamListResponse

`func NewLiveStreamListResponse(data []LiveStream, pagination Pagination, ) *LiveStreamListResponse`

NewLiveStreamListResponse instantiates a new LiveStreamListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamListResponseWithDefaults

`func NewLiveStreamListResponseWithDefaults() *LiveStreamListResponse`

NewLiveStreamListResponseWithDefaults instantiates a new LiveStreamListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LiveStreamListResponse) GetData() []LiveStream`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LiveStreamListResponse) GetDataOk() (*[]LiveStream, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LiveStreamListResponse) SetData(v []LiveStream)`

SetData sets Data field to given value.


### GetPagination

`func (o *LiveStreamListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *LiveStreamListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *LiveStreamListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


