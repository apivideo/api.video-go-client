/*
 * api.video
 *
 * api.video is an API that encodes on the go to facilitate immediate playback, enhancing viewer streaming experiences across multiple devices and platforms. You can stream live or on-demand online videos within minutes.
 *
 * API version: 1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apivideosdk

import (
//"encoding/json"
)

// AnalyticsMetricsBreakdownResponse struct for AnalyticsMetricsBreakdownResponse
type AnalyticsMetricsBreakdownResponse struct {
	Context AnalyticsMetricsBreakdownResponseContext `json:"context"`
	// Returns an array of dimensions and their respective metrics.
	Data       []AnalyticsMetricsBreakdownResponseData `json:"data"`
	Pagination Pagination                              `json:"pagination"`
}

// NewAnalyticsMetricsBreakdownResponse instantiates a new AnalyticsMetricsBreakdownResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsMetricsBreakdownResponse(context AnalyticsMetricsBreakdownResponseContext, data []AnalyticsMetricsBreakdownResponseData, pagination Pagination) *AnalyticsMetricsBreakdownResponse {
	this := AnalyticsMetricsBreakdownResponse{}
	this.Context = context
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewAnalyticsMetricsBreakdownResponseWithDefaults instantiates a new AnalyticsMetricsBreakdownResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsMetricsBreakdownResponseWithDefaults() *AnalyticsMetricsBreakdownResponse {
	this := AnalyticsMetricsBreakdownResponse{}
	return &this
}

// GetContext returns the Context field value
func (o *AnalyticsMetricsBreakdownResponse) GetContext() AnalyticsMetricsBreakdownResponseContext {
	if o == nil {
		var ret AnalyticsMetricsBreakdownResponseContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *AnalyticsMetricsBreakdownResponse) GetContextOk() (*AnalyticsMetricsBreakdownResponseContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *AnalyticsMetricsBreakdownResponse) SetContext(v AnalyticsMetricsBreakdownResponseContext) {
	o.Context = v
}

// GetData returns the Data field value
func (o *AnalyticsMetricsBreakdownResponse) GetData() []AnalyticsMetricsBreakdownResponseData {
	if o == nil {
		var ret []AnalyticsMetricsBreakdownResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AnalyticsMetricsBreakdownResponse) GetDataOk() (*[]AnalyticsMetricsBreakdownResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AnalyticsMetricsBreakdownResponse) SetData(v []AnalyticsMetricsBreakdownResponseData) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *AnalyticsMetricsBreakdownResponse) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *AnalyticsMetricsBreakdownResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *AnalyticsMetricsBreakdownResponse) SetPagination(v Pagination) {
	o.Pagination = v
}

type NullableAnalyticsMetricsBreakdownResponse struct {
	value *AnalyticsMetricsBreakdownResponse
	isSet bool
}

func (v NullableAnalyticsMetricsBreakdownResponse) Get() *AnalyticsMetricsBreakdownResponse {
	return v.value
}

func (v *NullableAnalyticsMetricsBreakdownResponse) Set(val *AnalyticsMetricsBreakdownResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsMetricsBreakdownResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsMetricsBreakdownResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsMetricsBreakdownResponse(val *AnalyticsMetricsBreakdownResponse) *NullableAnalyticsMetricsBreakdownResponse {
	return &NullableAnalyticsMetricsBreakdownResponse{value: val, isSet: true}
}
