# RawStatisticsListLiveStreamAnalyticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]LiveStreamSession**](LiveStreamSession.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](pagination.md) |  | [optional] 

## Methods

### NewRawStatisticsListLiveStreamAnalyticsResponse

`func NewRawStatisticsListLiveStreamAnalyticsResponse() *RawStatisticsListLiveStreamAnalyticsResponse`

NewRawStatisticsListLiveStreamAnalyticsResponse instantiates a new RawStatisticsListLiveStreamAnalyticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawStatisticsListLiveStreamAnalyticsResponseWithDefaults

`func NewRawStatisticsListLiveStreamAnalyticsResponseWithDefaults() *RawStatisticsListLiveStreamAnalyticsResponse`

NewRawStatisticsListLiveStreamAnalyticsResponseWithDefaults instantiates a new RawStatisticsListLiveStreamAnalyticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RawStatisticsListLiveStreamAnalyticsResponse) GetData() []LiveStreamSession`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RawStatisticsListLiveStreamAnalyticsResponse) GetDataOk() (*[]LiveStreamSession, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RawStatisticsListLiveStreamAnalyticsResponse) SetData(v []LiveStreamSession)`

SetData sets Data field to given value.

### HasData

`func (o *RawStatisticsListLiveStreamAnalyticsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *RawStatisticsListLiveStreamAnalyticsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RawStatisticsListLiveStreamAnalyticsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RawStatisticsListLiveStreamAnalyticsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RawStatisticsListLiveStreamAnalyticsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


