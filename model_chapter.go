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

// Chapter struct for Chapter
type Chapter struct {
	Uri *string `json:"uri,omitempty"`
	// The link to your VTT file, which contains your chapters information for the video.
	Src      *string `json:"src,omitempty"`
	Language *string `json:"language,omitempty"`
}

// NewChapter instantiates a new Chapter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChapter() *Chapter {
	this := Chapter{}
	return &this
}

// NewChapterWithDefaults instantiates a new Chapter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChapterWithDefaults() *Chapter {
	this := Chapter{}
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Chapter) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chapter) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Chapter) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Chapter) SetUri(v string) {
	o.Uri = &v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *Chapter) GetSrc() string {
	if o == nil || o.Src == nil {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chapter) GetSrcOk() (*string, bool) {
	if o == nil || o.Src == nil {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *Chapter) HasSrc() bool {
	if o != nil && o.Src != nil {
		return true
	}

	return false
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *Chapter) SetSrc(v string) {
	o.Src = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Chapter) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chapter) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Chapter) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Chapter) SetLanguage(v string) {
	o.Language = &v
}

type NullableChapter struct {
	value *Chapter
	isSet bool
}

func (v NullableChapter) Get() *Chapter {
	return v.value
}

func (v *NullableChapter) Set(val *Chapter) {
	v.value = val
	v.isSet = true
}

func (v NullableChapter) IsSet() bool {
	return v.isSet
}

func (v *NullableChapter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChapter(val *Chapter) *NullableChapter {
	return &NullableChapter{value: val, isSet: true}
}
