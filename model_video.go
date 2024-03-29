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

// Video struct for Video
type Video struct {
	// The unique identifier of the video object.
	VideoId string `json:"videoId"`
	// When a video was created, presented in ISO-8601 format.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The title of the video content.
	Title *string `json:"title,omitempty"`
	// A description for the video content.
	Description *string `json:"description,omitempty"`
	// The date and time the API created the video. Date and time are provided using ISO-8601 UTC format.
	PublishedAt *string `json:"publishedAt,omitempty"`
	// The date and time the video was updated. Date and time are provided using ISO-8601 UTC format.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// One array of tags (each tag is a string) in order to categorize a video. Tags may include spaces.
	Tags *[]string `json:"tags,omitempty"`
	// Metadata you can use to categorise and filter videos. Metadata is a list of dictionaries, where each dictionary represents a key value pair for categorising a video. [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata/) allows you to define a key that allows any value pair.
	Metadata *[]Metadata  `json:"metadata,omitempty"`
	Source   *VideoSource `json:"source,omitempty"`
	Assets   *VideoAssets `json:"assets,omitempty"`
	// The id of the player that will be applied on the video.
	PlayerId *string `json:"playerId,omitempty"`
	// Defines if the content is publicly reachable or if a unique token is needed for each play session. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos/).
	Public *bool `json:"public,omitempty"`
	// Defines if video is panoramic.
	Panoramic *bool `json:"panoramic,omitempty"`
	// This lets you know whether mp4 is supported. If enabled, an mp4 URL will be provided in the response for the video.
	Mp4Support *bool `json:"mp4Support,omitempty"`
}

// NewVideo instantiates a new Video object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideo(videoId string) *Video {
	this := Video{}
	this.VideoId = videoId
	return &this
}

// NewVideoWithDefaults instantiates a new Video object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoWithDefaults() *Video {
	this := Video{}
	return &this
}

// GetVideoId returns the VideoId field value
func (o *Video) GetVideoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VideoId
}

// GetVideoIdOk returns a tuple with the VideoId field value
// and a boolean to check if the value has been set.
func (o *Video) GetVideoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VideoId, true
}

// SetVideoId sets field value
func (o *Video) SetVideoId(v string) {
	o.VideoId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Video) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Video) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Video) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Video) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Video) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Video) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Video) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Video) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Video) SetDescription(v string) {
	o.Description = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *Video) GetPublishedAt() string {
	if o == nil || o.PublishedAt == nil {
		var ret string
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetPublishedAtOk() (*string, bool) {
	if o == nil || o.PublishedAt == nil {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *Video) HasPublishedAt() bool {
	if o != nil && o.PublishedAt != nil {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given string and assigns it to the PublishedAt field.
func (o *Video) SetPublishedAt(v string) {
	o.PublishedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Video) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Video) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Video) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Video) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Video) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Video) SetTags(v []string) {
	o.Tags = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Video) GetMetadata() []Metadata {
	if o == nil || o.Metadata == nil {
		var ret []Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetMetadataOk() (*[]Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Video) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []Metadata and assigns it to the Metadata field.
func (o *Video) SetMetadata(v []Metadata) {
	o.Metadata = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Video) GetSource() VideoSource {
	if o == nil || o.Source == nil {
		var ret VideoSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetSourceOk() (*VideoSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Video) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given VideoSource and assigns it to the Source field.
func (o *Video) SetSource(v VideoSource) {
	o.Source = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *Video) GetAssets() VideoAssets {
	if o == nil || o.Assets == nil {
		var ret VideoAssets
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetAssetsOk() (*VideoAssets, bool) {
	if o == nil || o.Assets == nil {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *Video) HasAssets() bool {
	if o != nil && o.Assets != nil {
		return true
	}

	return false
}

// SetAssets gets a reference to the given VideoAssets and assigns it to the Assets field.
func (o *Video) SetAssets(v VideoAssets) {
	o.Assets = &v
}

// GetPlayerId returns the PlayerId field value if set, zero value otherwise.
func (o *Video) GetPlayerId() string {
	if o == nil || o.PlayerId == nil {
		var ret string
		return ret
	}
	return *o.PlayerId
}

// GetPlayerIdOk returns a tuple with the PlayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetPlayerIdOk() (*string, bool) {
	if o == nil || o.PlayerId == nil {
		return nil, false
	}
	return o.PlayerId, true
}

// HasPlayerId returns a boolean if a field has been set.
func (o *Video) HasPlayerId() bool {
	if o != nil && o.PlayerId != nil {
		return true
	}

	return false
}

// SetPlayerId gets a reference to the given string and assigns it to the PlayerId field.
func (o *Video) SetPlayerId(v string) {
	o.PlayerId = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *Video) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *Video) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *Video) SetPublic(v bool) {
	o.Public = &v
}

// GetPanoramic returns the Panoramic field value if set, zero value otherwise.
func (o *Video) GetPanoramic() bool {
	if o == nil || o.Panoramic == nil {
		var ret bool
		return ret
	}
	return *o.Panoramic
}

// GetPanoramicOk returns a tuple with the Panoramic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetPanoramicOk() (*bool, bool) {
	if o == nil || o.Panoramic == nil {
		return nil, false
	}
	return o.Panoramic, true
}

// HasPanoramic returns a boolean if a field has been set.
func (o *Video) HasPanoramic() bool {
	if o != nil && o.Panoramic != nil {
		return true
	}

	return false
}

// SetPanoramic gets a reference to the given bool and assigns it to the Panoramic field.
func (o *Video) SetPanoramic(v bool) {
	o.Panoramic = &v
}

// GetMp4Support returns the Mp4Support field value if set, zero value otherwise.
func (o *Video) GetMp4Support() bool {
	if o == nil || o.Mp4Support == nil {
		var ret bool
		return ret
	}
	return *o.Mp4Support
}

// GetMp4SupportOk returns a tuple with the Mp4Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetMp4SupportOk() (*bool, bool) {
	if o == nil || o.Mp4Support == nil {
		return nil, false
	}
	return o.Mp4Support, true
}

// HasMp4Support returns a boolean if a field has been set.
func (o *Video) HasMp4Support() bool {
	if o != nil && o.Mp4Support != nil {
		return true
	}

	return false
}

// SetMp4Support gets a reference to the given bool and assigns it to the Mp4Support field.
func (o *Video) SetMp4Support(v bool) {
	o.Mp4Support = &v
}

type NullableVideo struct {
	value *Video
	isSet bool
}

func (v NullableVideo) Get() *Video {
	return v.value
}

func (v *NullableVideo) Set(val *Video) {
	v.value = val
	v.isSet = true
}

func (v NullableVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideo(val *Video) *NullableVideo {
	return &NullableVideo{value: val, isSet: true}
}
