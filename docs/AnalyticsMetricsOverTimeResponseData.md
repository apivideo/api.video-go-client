# AnalyticsMetricsOverTimeResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmittedAt** | Pointer to **string** | Returns the timestamp of the event that belongs to a specific metric in ATOM date-time format. For example, if you set &#x60;play&#x60; with an &#x60;hour&#x60; interval in your request, then &#x60;emittedAt&#x60; returns the hourly timestamps of every play event within the timeframe you defined. | [optional] 
**MetricValue** | Pointer to **float32** | Returns the data for a specific metric value. | [optional] 

## Methods

### NewAnalyticsMetricsOverTimeResponseData

`func NewAnalyticsMetricsOverTimeResponseData() *AnalyticsMetricsOverTimeResponseData`

NewAnalyticsMetricsOverTimeResponseData instantiates a new AnalyticsMetricsOverTimeResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsMetricsOverTimeResponseDataWithDefaults

`func NewAnalyticsMetricsOverTimeResponseDataWithDefaults() *AnalyticsMetricsOverTimeResponseData`

NewAnalyticsMetricsOverTimeResponseDataWithDefaults instantiates a new AnalyticsMetricsOverTimeResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmittedAt

`func (o *AnalyticsMetricsOverTimeResponseData) GetEmittedAt() string`

GetEmittedAt returns the EmittedAt field if non-nil, zero value otherwise.

### GetEmittedAtOk

`func (o *AnalyticsMetricsOverTimeResponseData) GetEmittedAtOk() (*string, bool)`

GetEmittedAtOk returns a tuple with the EmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmittedAt

`func (o *AnalyticsMetricsOverTimeResponseData) SetEmittedAt(v string)`

SetEmittedAt sets EmittedAt field to given value.

### HasEmittedAt

`func (o *AnalyticsMetricsOverTimeResponseData) HasEmittedAt() bool`

HasEmittedAt returns a boolean if a field has been set.

### GetMetricValue

`func (o *AnalyticsMetricsOverTimeResponseData) GetMetricValue() float32`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *AnalyticsMetricsOverTimeResponseData) GetMetricValueOk() (*float32, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *AnalyticsMetricsOverTimeResponseData) SetMetricValue(v float32)`

SetMetricValue sets MetricValue field to given value.

### HasMetricValue

`func (o *AnalyticsMetricsOverTimeResponseData) HasMetricValue() bool`

HasMetricValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


