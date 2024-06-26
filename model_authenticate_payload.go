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

// AuthenticatePayload struct for AuthenticatePayload
type AuthenticatePayload struct {
	// Your account API key. You can use your sandbox API key, or you can use your production API key.
	ApiKey string `json:"apiKey"`
}

// NewAuthenticatePayload instantiates a new AuthenticatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatePayload(apiKey string) *AuthenticatePayload {
	this := AuthenticatePayload{}
	this.ApiKey = apiKey
	return &this
}

// NewAuthenticatePayloadWithDefaults instantiates a new AuthenticatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatePayloadWithDefaults() *AuthenticatePayload {
	this := AuthenticatePayload{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *AuthenticatePayload) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *AuthenticatePayload) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *AuthenticatePayload) SetApiKey(v string) {
	o.ApiKey = v
}

type NullableAuthenticatePayload struct {
	value *AuthenticatePayload
	isSet bool
}

func (v NullableAuthenticatePayload) Get() *AuthenticatePayload {
	return v.value
}

func (v *NullableAuthenticatePayload) Set(val *AuthenticatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatePayload(val *AuthenticatePayload) *NullableAuthenticatePayload {
	return &NullableAuthenticatePayload{value: val, isSet: true}
}
