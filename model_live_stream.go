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

// LiveStream struct for LiveStream
type LiveStream struct {
	// The unique identifier for the live stream. Live stream IDs begin with \"li.\"
	LiveStreamId string `json:"liveStreamId"`
	// The name of your live stream.
	Name *string `json:"name,omitempty"`
	// The unique, private stream key that you use to begin streaming.
	StreamKey *string `json:"streamKey,omitempty"`
	// Whether your video can be viewed by everyone, or requires authentication to see it. A setting of false will require a unique token for each view. Learn more about the Private Video feature [here](https://docs.api.video/delivery/video-privacy-access-management).
	Public *bool             `json:"public,omitempty"`
	Assets *LiveStreamAssets `json:"assets,omitempty"`
	// The unique identifier for the player.
	PlayerId *string `json:"playerId,omitempty"`
	// Whether or not you are broadcasting the live video you recorded for others to see. True means you are broadcasting to viewers, false means you are not.
	Broadcasting *bool `json:"broadcasting,omitempty"`
	// Returns the list of restream destinations.
	Restreams []RestreamsResponseObject `json:"restreams"`
	// When the player was created, presented in ATOM UTC format.
	CreatedAt *string `json:"createdAt,omitempty"`
	// When the player was last updated, presented in ATOM UTC format.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewLiveStream instantiates a new LiveStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveStream(liveStreamId string, restreams []RestreamsResponseObject) *LiveStream {
	this := LiveStream{}
	this.LiveStreamId = liveStreamId
	this.Restreams = restreams
	return &this
}

// NewLiveStreamWithDefaults instantiates a new LiveStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveStreamWithDefaults() *LiveStream {
	this := LiveStream{}
	return &this
}

// GetLiveStreamId returns the LiveStreamId field value
func (o *LiveStream) GetLiveStreamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LiveStreamId
}

// GetLiveStreamIdOk returns a tuple with the LiveStreamId field value
// and a boolean to check if the value has been set.
func (o *LiveStream) GetLiveStreamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LiveStreamId, true
}

// SetLiveStreamId sets field value
func (o *LiveStream) SetLiveStreamId(v string) {
	o.LiveStreamId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LiveStream) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LiveStream) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LiveStream) SetName(v string) {
	o.Name = &v
}

// GetStreamKey returns the StreamKey field value if set, zero value otherwise.
func (o *LiveStream) GetStreamKey() string {
	if o == nil || o.StreamKey == nil {
		var ret string
		return ret
	}
	return *o.StreamKey
}

// GetStreamKeyOk returns a tuple with the StreamKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetStreamKeyOk() (*string, bool) {
	if o == nil || o.StreamKey == nil {
		return nil, false
	}
	return o.StreamKey, true
}

// HasStreamKey returns a boolean if a field has been set.
func (o *LiveStream) HasStreamKey() bool {
	if o != nil && o.StreamKey != nil {
		return true
	}

	return false
}

// SetStreamKey gets a reference to the given string and assigns it to the StreamKey field.
func (o *LiveStream) SetStreamKey(v string) {
	o.StreamKey = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *LiveStream) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *LiveStream) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *LiveStream) SetPublic(v bool) {
	o.Public = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *LiveStream) GetAssets() LiveStreamAssets {
	if o == nil || o.Assets == nil {
		var ret LiveStreamAssets
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetAssetsOk() (*LiveStreamAssets, bool) {
	if o == nil || o.Assets == nil {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *LiveStream) HasAssets() bool {
	if o != nil && o.Assets != nil {
		return true
	}

	return false
}

// SetAssets gets a reference to the given LiveStreamAssets and assigns it to the Assets field.
func (o *LiveStream) SetAssets(v LiveStreamAssets) {
	o.Assets = &v
}

// GetPlayerId returns the PlayerId field value if set, zero value otherwise.
func (o *LiveStream) GetPlayerId() string {
	if o == nil || o.PlayerId == nil {
		var ret string
		return ret
	}
	return *o.PlayerId
}

// GetPlayerIdOk returns a tuple with the PlayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetPlayerIdOk() (*string, bool) {
	if o == nil || o.PlayerId == nil {
		return nil, false
	}
	return o.PlayerId, true
}

// HasPlayerId returns a boolean if a field has been set.
func (o *LiveStream) HasPlayerId() bool {
	if o != nil && o.PlayerId != nil {
		return true
	}

	return false
}

// SetPlayerId gets a reference to the given string and assigns it to the PlayerId field.
func (o *LiveStream) SetPlayerId(v string) {
	o.PlayerId = &v
}

// GetBroadcasting returns the Broadcasting field value if set, zero value otherwise.
func (o *LiveStream) GetBroadcasting() bool {
	if o == nil || o.Broadcasting == nil {
		var ret bool
		return ret
	}
	return *o.Broadcasting
}

// GetBroadcastingOk returns a tuple with the Broadcasting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetBroadcastingOk() (*bool, bool) {
	if o == nil || o.Broadcasting == nil {
		return nil, false
	}
	return o.Broadcasting, true
}

// HasBroadcasting returns a boolean if a field has been set.
func (o *LiveStream) HasBroadcasting() bool {
	if o != nil && o.Broadcasting != nil {
		return true
	}

	return false
}

// SetBroadcasting gets a reference to the given bool and assigns it to the Broadcasting field.
func (o *LiveStream) SetBroadcasting(v bool) {
	o.Broadcasting = &v
}

// GetRestreams returns the Restreams field value
func (o *LiveStream) GetRestreams() []RestreamsResponseObject {
	if o == nil {
		var ret []RestreamsResponseObject
		return ret
	}

	return o.Restreams
}

// GetRestreamsOk returns a tuple with the Restreams field value
// and a boolean to check if the value has been set.
func (o *LiveStream) GetRestreamsOk() (*[]RestreamsResponseObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Restreams, true
}

// SetRestreams sets field value
func (o *LiveStream) SetRestreams(v []RestreamsResponseObject) {
	o.Restreams = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LiveStream) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LiveStream) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LiveStream) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LiveStream) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LiveStream) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *LiveStream) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

type NullableLiveStream struct {
	value *LiveStream
	isSet bool
}

func (v NullableLiveStream) Get() *LiveStream {
	return v.value
}

func (v *NullableLiveStream) Set(val *LiveStream) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveStream) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveStream(val *LiveStream) *NullableLiveStream {
	return &NullableLiveStream{value: val, isSet: true}
}
