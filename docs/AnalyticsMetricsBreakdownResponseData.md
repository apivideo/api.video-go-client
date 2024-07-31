# AnalyticsMetricsBreakdownResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DimensionValue** | Pointer to **string** | Returns a specific value for the dimension you selected, based on the data. For example if you select &#x60;continent&#x60; as a dimension, then &#x60;dimensionValue&#x60; returns values like &#x60;EU&#x60; or \&quot;AZ\&quot;. | [optional] 
**MetricValue** | Pointer to **float32** | Returns the data for a specific dimension value. | [optional] 

## Methods

### NewAnalyticsMetricsBreakdownResponseData

`func NewAnalyticsMetricsBreakdownResponseData() *AnalyticsMetricsBreakdownResponseData`

NewAnalyticsMetricsBreakdownResponseData instantiates a new AnalyticsMetricsBreakdownResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsMetricsBreakdownResponseDataWithDefaults

`func NewAnalyticsMetricsBreakdownResponseDataWithDefaults() *AnalyticsMetricsBreakdownResponseData`

NewAnalyticsMetricsBreakdownResponseDataWithDefaults instantiates a new AnalyticsMetricsBreakdownResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensionValue

`func (o *AnalyticsMetricsBreakdownResponseData) GetDimensionValue() string`

GetDimensionValue returns the DimensionValue field if non-nil, zero value otherwise.

### GetDimensionValueOk

`func (o *AnalyticsMetricsBreakdownResponseData) GetDimensionValueOk() (*string, bool)`

GetDimensionValueOk returns a tuple with the DimensionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionValue

`func (o *AnalyticsMetricsBreakdownResponseData) SetDimensionValue(v string)`

SetDimensionValue sets DimensionValue field to given value.

### HasDimensionValue

`func (o *AnalyticsMetricsBreakdownResponseData) HasDimensionValue() bool`

HasDimensionValue returns a boolean if a field has been set.

### GetMetricValue

`func (o *AnalyticsMetricsBreakdownResponseData) GetMetricValue() float32`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *AnalyticsMetricsBreakdownResponseData) GetMetricValueOk() (*float32, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *AnalyticsMetricsBreakdownResponseData) SetMetricValue(v float32)`

SetMetricValue sets MetricValue field to given value.

### HasMetricValue

`func (o *AnalyticsMetricsBreakdownResponseData) HasMetricValue() bool`

HasMetricValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


