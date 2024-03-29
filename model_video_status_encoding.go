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

// VideoStatusEncoding struct for VideoStatusEncoding
type VideoStatusEncoding struct {
	// Whether the video is playable or not.
	Playable *bool `json:"playable,omitempty"`
	// Available qualities the video can be viewed in.
	Qualities *[]Quality                   `json:"qualities,omitempty"`
	Metadata  *VideoStatusEncodingMetadata `json:"metadata,omitempty"`
}

// NewVideoStatusEncoding instantiates a new VideoStatusEncoding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoStatusEncoding() *VideoStatusEncoding {
	this := VideoStatusEncoding{}
	return &this
}

// NewVideoStatusEncodingWithDefaults instantiates a new VideoStatusEncoding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoStatusEncodingWithDefaults() *VideoStatusEncoding {
	this := VideoStatusEncoding{}
	return &this
}

// GetPlayable returns the Playable field value if set, zero value otherwise.
func (o *VideoStatusEncoding) GetPlayable() bool {
	if o == nil || o.Playable == nil {
		var ret bool
		return ret
	}
	return *o.Playable
}

// GetPlayableOk returns a tuple with the Playable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStatusEncoding) GetPlayableOk() (*bool, bool) {
	if o == nil || o.Playable == nil {
		return nil, false
	}
	return o.Playable, true
}

// HasPlayable returns a boolean if a field has been set.
func (o *VideoStatusEncoding) HasPlayable() bool {
	if o != nil && o.Playable != nil {
		return true
	}

	return false
}

// SetPlayable gets a reference to the given bool and assigns it to the Playable field.
func (o *VideoStatusEncoding) SetPlayable(v bool) {
	o.Playable = &v
}

// GetQualities returns the Qualities field value if set, zero value otherwise.
func (o *VideoStatusEncoding) GetQualities() []Quality {
	if o == nil || o.Qualities == nil {
		var ret []Quality
		return ret
	}
	return *o.Qualities
}

// GetQualitiesOk returns a tuple with the Qualities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStatusEncoding) GetQualitiesOk() (*[]Quality, bool) {
	if o == nil || o.Qualities == nil {
		return nil, false
	}
	return o.Qualities, true
}

// HasQualities returns a boolean if a field has been set.
func (o *VideoStatusEncoding) HasQualities() bool {
	if o != nil && o.Qualities != nil {
		return true
	}

	return false
}

// SetQualities gets a reference to the given []Quality and assigns it to the Qualities field.
func (o *VideoStatusEncoding) SetQualities(v []Quality) {
	o.Qualities = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *VideoStatusEncoding) GetMetadata() VideoStatusEncodingMetadata {
	if o == nil || o.Metadata == nil {
		var ret VideoStatusEncodingMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStatusEncoding) GetMetadataOk() (*VideoStatusEncodingMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *VideoStatusEncoding) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given VideoStatusEncodingMetadata and assigns it to the Metadata field.
func (o *VideoStatusEncoding) SetMetadata(v VideoStatusEncodingMetadata) {
	o.Metadata = &v
}

type NullableVideoStatusEncoding struct {
	value *VideoStatusEncoding
	isSet bool
}

func (v NullableVideoStatusEncoding) Get() *VideoStatusEncoding {
	return v.value
}

func (v *NullableVideoStatusEncoding) Set(val *VideoStatusEncoding) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoStatusEncoding) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoStatusEncoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoStatusEncoding(val *VideoStatusEncoding) *NullableVideoStatusEncoding {
	return &NullableVideoStatusEncoding{value: val, isSet: true}
}
