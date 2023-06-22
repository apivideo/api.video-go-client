# AnalyticsPlaysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AnalyticsData**](AnalyticsData.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewAnalyticsPlaysResponse

`func NewAnalyticsPlaysResponse(data []AnalyticsData, pagination Pagination, ) *AnalyticsPlaysResponse`

NewAnalyticsPlaysResponse instantiates a new AnalyticsPlaysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsPlaysResponseWithDefaults

`func NewAnalyticsPlaysResponseWithDefaults() *AnalyticsPlaysResponse`

NewAnalyticsPlaysResponseWithDefaults instantiates a new AnalyticsPlaysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AnalyticsPlaysResponse) GetData() []AnalyticsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnalyticsPlaysResponse) GetDataOk() (*[]AnalyticsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnalyticsPlaysResponse) SetData(v []AnalyticsData)`

SetData sets Data field to given value.


### GetPagination

`func (o *AnalyticsPlaysResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AnalyticsPlaysResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AnalyticsPlaysResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


