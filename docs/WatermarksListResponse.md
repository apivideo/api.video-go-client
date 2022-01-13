# WatermarksListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Watermark**](Watermark.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewWatermarksListResponse

`func NewWatermarksListResponse(data []Watermark, pagination Pagination, ) *WatermarksListResponse`

NewWatermarksListResponse instantiates a new WatermarksListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarksListResponseWithDefaults

`func NewWatermarksListResponseWithDefaults() *WatermarksListResponse`

NewWatermarksListResponseWithDefaults instantiates a new WatermarksListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WatermarksListResponse) GetData() []Watermark`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WatermarksListResponse) GetDataOk() (*[]Watermark, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WatermarksListResponse) SetData(v []Watermark)`

SetData sets Data field to given value.


### GetPagination

`func (o *WatermarksListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *WatermarksListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *WatermarksListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


