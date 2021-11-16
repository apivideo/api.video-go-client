# VideosListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Video**](Video.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewVideosListResponse

`func NewVideosListResponse(data []Video, pagination Pagination, ) *VideosListResponse`

NewVideosListResponse instantiates a new VideosListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideosListResponseWithDefaults

`func NewVideosListResponseWithDefaults() *VideosListResponse`

NewVideosListResponseWithDefaults instantiates a new VideosListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VideosListResponse) GetData() []Video`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VideosListResponse) GetDataOk() (*[]Video, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VideosListResponse) SetData(v []Video)`

SetData sets Data field to given value.


### GetPagination

`func (o *VideosListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VideosListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VideosListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


