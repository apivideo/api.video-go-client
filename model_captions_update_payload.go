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

// CaptionsUpdatePayload struct for CaptionsUpdatePayload
type CaptionsUpdatePayload struct {
	// Set this parameter to `true` to define a caption as the default for a video.
	Default *bool `json:"default,omitempty"`
}

// NewCaptionsUpdatePayload instantiates a new CaptionsUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptionsUpdatePayload() *CaptionsUpdatePayload {
	this := CaptionsUpdatePayload{}
	return &this
}

// NewCaptionsUpdatePayloadWithDefaults instantiates a new CaptionsUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptionsUpdatePayloadWithDefaults() *CaptionsUpdatePayload {
	this := CaptionsUpdatePayload{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CaptionsUpdatePayload) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptionsUpdatePayload) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CaptionsUpdatePayload) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CaptionsUpdatePayload) SetDefault(v bool) {
	o.Default = &v
}

type NullableCaptionsUpdatePayload struct {
	value *CaptionsUpdatePayload
	isSet bool
}

func (v NullableCaptionsUpdatePayload) Get() *CaptionsUpdatePayload {
	return v.value
}

func (v *NullableCaptionsUpdatePayload) Set(val *CaptionsUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptionsUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptionsUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptionsUpdatePayload(val *CaptionsUpdatePayload) *NullableCaptionsUpdatePayload {
	return &NullableCaptionsUpdatePayload{value: val, isSet: true}
}
