# AnalyticsMetricsBreakdownResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**AnalyticsMetricsBreakdownResponseContext**](AnalyticsMetricsBreakdownResponseContext.md) |  | 
**Data** | [**[]AnalyticsMetricsBreakdownResponseData**](AnalyticsMetricsBreakdownResponseData.md) | Returns an array of dimensions and their respective metrics. | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewAnalyticsMetricsBreakdownResponse

`func NewAnalyticsMetricsBreakdownResponse(context AnalyticsMetricsBreakdownResponseContext, data []AnalyticsMetricsBreakdownResponseData, pagination Pagination, ) *AnalyticsMetricsBreakdownResponse`

NewAnalyticsMetricsBreakdownResponse instantiates a new AnalyticsMetricsBreakdownResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsMetricsBreakdownResponseWithDefaults

`func NewAnalyticsMetricsBreakdownResponseWithDefaults() *AnalyticsMetricsBreakdownResponse`

NewAnalyticsMetricsBreakdownResponseWithDefaults instantiates a new AnalyticsMetricsBreakdownResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *AnalyticsMetricsBreakdownResponse) GetContext() AnalyticsMetricsBreakdownResponseContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AnalyticsMetricsBreakdownResponse) GetContextOk() (*AnalyticsMetricsBreakdownResponseContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AnalyticsMetricsBreakdownResponse) SetContext(v AnalyticsMetricsBreakdownResponseContext)`

SetContext sets Context field to given value.


### GetData

`func (o *AnalyticsMetricsBreakdownResponse) GetData() []AnalyticsMetricsBreakdownResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnalyticsMetricsBreakdownResponse) GetDataOk() (*[]AnalyticsMetricsBreakdownResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnalyticsMetricsBreakdownResponse) SetData(v []AnalyticsMetricsBreakdownResponseData)`

SetData sets Data field to given value.


### GetPagination

`func (o *AnalyticsMetricsBreakdownResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AnalyticsMetricsBreakdownResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AnalyticsMetricsBreakdownResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


