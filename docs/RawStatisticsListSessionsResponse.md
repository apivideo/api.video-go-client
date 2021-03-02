# RawStatisticsListSessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]VideoSession**](VideoSession.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](pagination.md) |  | [optional] 

## Methods

### NewRawStatisticsListSessionsResponse

`func NewRawStatisticsListSessionsResponse() *RawStatisticsListSessionsResponse`

NewRawStatisticsListSessionsResponse instantiates a new RawStatisticsListSessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawStatisticsListSessionsResponseWithDefaults

`func NewRawStatisticsListSessionsResponseWithDefaults() *RawStatisticsListSessionsResponse`

NewRawStatisticsListSessionsResponseWithDefaults instantiates a new RawStatisticsListSessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RawStatisticsListSessionsResponse) GetData() []VideoSession`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RawStatisticsListSessionsResponse) GetDataOk() (*[]VideoSession, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RawStatisticsListSessionsResponse) SetData(v []VideoSession)`

SetData sets Data field to given value.

### HasData

`func (o *RawStatisticsListSessionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *RawStatisticsListSessionsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RawStatisticsListSessionsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RawStatisticsListSessionsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RawStatisticsListSessionsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


