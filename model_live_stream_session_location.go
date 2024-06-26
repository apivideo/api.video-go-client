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

// LiveStreamSessionLocation The location of the viewer of the live stream.
type LiveStreamSessionLocation struct {
	// The country of the viewer of the live stream.
	Country *string `json:"country,omitempty"`
	// The city of the viewer of the live stream.
	City *string `json:"city,omitempty"`
}

// NewLiveStreamSessionLocation instantiates a new LiveStreamSessionLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveStreamSessionLocation() *LiveStreamSessionLocation {
	this := LiveStreamSessionLocation{}
	return &this
}

// NewLiveStreamSessionLocationWithDefaults instantiates a new LiveStreamSessionLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveStreamSessionLocationWithDefaults() *LiveStreamSessionLocation {
	this := LiveStreamSessionLocation{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *LiveStreamSessionLocation) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamSessionLocation) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *LiveStreamSessionLocation) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *LiveStreamSessionLocation) SetCountry(v string) {
	o.Country = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *LiveStreamSessionLocation) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamSessionLocation) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *LiveStreamSessionLocation) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *LiveStreamSessionLocation) SetCity(v string) {
	o.City = &v
}

type NullableLiveStreamSessionLocation struct {
	value *LiveStreamSessionLocation
	isSet bool
}

func (v NullableLiveStreamSessionLocation) Get() *LiveStreamSessionLocation {
	return v.value
}

func (v *NullableLiveStreamSessionLocation) Set(val *LiveStreamSessionLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveStreamSessionLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveStreamSessionLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveStreamSessionLocation(val *LiveStreamSessionLocation) *NullableLiveStreamSessionLocation {
	return &NullableLiveStreamSessionLocation{value: val, isSet: true}
}
