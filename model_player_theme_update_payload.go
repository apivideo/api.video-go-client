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

// PlayerThemeUpdatePayload struct for PlayerThemeUpdatePayload
type PlayerThemeUpdatePayload struct {
	// Add a name for your player theme here.
	Name *string `json:"name,omitempty"`
	// RGBA color for timer text. Default: rgba(255, 255, 255, 1)
	Text *string `json:"text,omitempty"`
	// RGBA color for all controls. Default: rgba(255, 255, 255, 1)
	Link *string `json:"link,omitempty"`
	// RGBA color for all controls when hovered. Default: rgba(255, 255, 255, 1)
	LinkHover *string `json:"linkHover,omitempty"`
	// RGBA color for the play button when hovered.
	LinkActive *string `json:"linkActive,omitempty"`
	// RGBA color playback bar: played content. Default: rgba(88, 131, 255, .95)
	TrackPlayed *string `json:"trackPlayed,omitempty"`
	// RGBA color playback bar: downloaded but unplayed (buffered) content. Default: rgba(255, 255, 255, .35)
	TrackUnplayed *string `json:"trackUnplayed,omitempty"`
	// RGBA color playback bar: background. Default: rgba(255, 255, 255, .2)
	TrackBackground *string `json:"trackBackground,omitempty"`
	// RGBA color: top 50% of background. Default: rgba(0, 0, 0, .7)
	BackgroundTop *string `json:"backgroundTop,omitempty"`
	// RGBA color: bottom 50% of background. Default: rgba(0, 0, 0, .7)
	BackgroundBottom *string `json:"backgroundBottom,omitempty"`
	// RGBA color for title text. Default: rgba(255, 255, 255, 1)
	BackgroundText *string `json:"backgroundText,omitempty"`
	// enable/disable player SDK access. Default: true
	EnableApi *bool `json:"enableApi,omitempty"`
	// enable/disable player controls. Default: true
	EnableControls *bool `json:"enableControls,omitempty"`
	// enable/disable player autoplay. Default: false
	ForceAutoplay *bool `json:"forceAutoplay,omitempty"`
	// enable/disable title. Default: false
	HideTitle *bool `json:"hideTitle,omitempty"`
	// enable/disable looping. Default: false
	ForceLoop *bool `json:"forceLoop,omitempty"`
}

// NewPlayerThemeUpdatePayload instantiates a new PlayerThemeUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerThemeUpdatePayload() *PlayerThemeUpdatePayload {
	this := PlayerThemeUpdatePayload{}
	return &this
}

// NewPlayerThemeUpdatePayloadWithDefaults instantiates a new PlayerThemeUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerThemeUpdatePayloadWithDefaults() *PlayerThemeUpdatePayload {
	this := PlayerThemeUpdatePayload{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlayerThemeUpdatePayload) SetName(v string) {
	o.Name = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PlayerThemeUpdatePayload) SetText(v string) {
	o.Text = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *PlayerThemeUpdatePayload) SetLink(v string) {
	o.Link = &v
}

// GetLinkHover returns the LinkHover field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetLinkHover() string {
	if o == nil || o.LinkHover == nil {
		var ret string
		return ret
	}
	return *o.LinkHover
}

// GetLinkHoverOk returns a tuple with the LinkHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetLinkHoverOk() (*string, bool) {
	if o == nil || o.LinkHover == nil {
		return nil, false
	}
	return o.LinkHover, true
}

// HasLinkHover returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasLinkHover() bool {
	if o != nil && o.LinkHover != nil {
		return true
	}

	return false
}

// SetLinkHover gets a reference to the given string and assigns it to the LinkHover field.
func (o *PlayerThemeUpdatePayload) SetLinkHover(v string) {
	o.LinkHover = &v
}

// GetLinkActive returns the LinkActive field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetLinkActive() string {
	if o == nil || o.LinkActive == nil {
		var ret string
		return ret
	}
	return *o.LinkActive
}

// GetLinkActiveOk returns a tuple with the LinkActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetLinkActiveOk() (*string, bool) {
	if o == nil || o.LinkActive == nil {
		return nil, false
	}
	return o.LinkActive, true
}

// HasLinkActive returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasLinkActive() bool {
	if o != nil && o.LinkActive != nil {
		return true
	}

	return false
}

// SetLinkActive gets a reference to the given string and assigns it to the LinkActive field.
func (o *PlayerThemeUpdatePayload) SetLinkActive(v string) {
	o.LinkActive = &v
}

// GetTrackPlayed returns the TrackPlayed field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetTrackPlayed() string {
	if o == nil || o.TrackPlayed == nil {
		var ret string
		return ret
	}
	return *o.TrackPlayed
}

// GetTrackPlayedOk returns a tuple with the TrackPlayed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetTrackPlayedOk() (*string, bool) {
	if o == nil || o.TrackPlayed == nil {
		return nil, false
	}
	return o.TrackPlayed, true
}

// HasTrackPlayed returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasTrackPlayed() bool {
	if o != nil && o.TrackPlayed != nil {
		return true
	}

	return false
}

// SetTrackPlayed gets a reference to the given string and assigns it to the TrackPlayed field.
func (o *PlayerThemeUpdatePayload) SetTrackPlayed(v string) {
	o.TrackPlayed = &v
}

// GetTrackUnplayed returns the TrackUnplayed field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetTrackUnplayed() string {
	if o == nil || o.TrackUnplayed == nil {
		var ret string
		return ret
	}
	return *o.TrackUnplayed
}

// GetTrackUnplayedOk returns a tuple with the TrackUnplayed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetTrackUnplayedOk() (*string, bool) {
	if o == nil || o.TrackUnplayed == nil {
		return nil, false
	}
	return o.TrackUnplayed, true
}

// HasTrackUnplayed returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasTrackUnplayed() bool {
	if o != nil && o.TrackUnplayed != nil {
		return true
	}

	return false
}

// SetTrackUnplayed gets a reference to the given string and assigns it to the TrackUnplayed field.
func (o *PlayerThemeUpdatePayload) SetTrackUnplayed(v string) {
	o.TrackUnplayed = &v
}

// GetTrackBackground returns the TrackBackground field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetTrackBackground() string {
	if o == nil || o.TrackBackground == nil {
		var ret string
		return ret
	}
	return *o.TrackBackground
}

// GetTrackBackgroundOk returns a tuple with the TrackBackground field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetTrackBackgroundOk() (*string, bool) {
	if o == nil || o.TrackBackground == nil {
		return nil, false
	}
	return o.TrackBackground, true
}

// HasTrackBackground returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasTrackBackground() bool {
	if o != nil && o.TrackBackground != nil {
		return true
	}

	return false
}

// SetTrackBackground gets a reference to the given string and assigns it to the TrackBackground field.
func (o *PlayerThemeUpdatePayload) SetTrackBackground(v string) {
	o.TrackBackground = &v
}

// GetBackgroundTop returns the BackgroundTop field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetBackgroundTop() string {
	if o == nil || o.BackgroundTop == nil {
		var ret string
		return ret
	}
	return *o.BackgroundTop
}

// GetBackgroundTopOk returns a tuple with the BackgroundTop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetBackgroundTopOk() (*string, bool) {
	if o == nil || o.BackgroundTop == nil {
		return nil, false
	}
	return o.BackgroundTop, true
}

// HasBackgroundTop returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasBackgroundTop() bool {
	if o != nil && o.BackgroundTop != nil {
		return true
	}

	return false
}

// SetBackgroundTop gets a reference to the given string and assigns it to the BackgroundTop field.
func (o *PlayerThemeUpdatePayload) SetBackgroundTop(v string) {
	o.BackgroundTop = &v
}

// GetBackgroundBottom returns the BackgroundBottom field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetBackgroundBottom() string {
	if o == nil || o.BackgroundBottom == nil {
		var ret string
		return ret
	}
	return *o.BackgroundBottom
}

// GetBackgroundBottomOk returns a tuple with the BackgroundBottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetBackgroundBottomOk() (*string, bool) {
	if o == nil || o.BackgroundBottom == nil {
		return nil, false
	}
	return o.BackgroundBottom, true
}

// HasBackgroundBottom returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasBackgroundBottom() bool {
	if o != nil && o.BackgroundBottom != nil {
		return true
	}

	return false
}

// SetBackgroundBottom gets a reference to the given string and assigns it to the BackgroundBottom field.
func (o *PlayerThemeUpdatePayload) SetBackgroundBottom(v string) {
	o.BackgroundBottom = &v
}

// GetBackgroundText returns the BackgroundText field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetBackgroundText() string {
	if o == nil || o.BackgroundText == nil {
		var ret string
		return ret
	}
	return *o.BackgroundText
}

// GetBackgroundTextOk returns a tuple with the BackgroundText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetBackgroundTextOk() (*string, bool) {
	if o == nil || o.BackgroundText == nil {
		return nil, false
	}
	return o.BackgroundText, true
}

// HasBackgroundText returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasBackgroundText() bool {
	if o != nil && o.BackgroundText != nil {
		return true
	}

	return false
}

// SetBackgroundText gets a reference to the given string and assigns it to the BackgroundText field.
func (o *PlayerThemeUpdatePayload) SetBackgroundText(v string) {
	o.BackgroundText = &v
}

// GetEnableApi returns the EnableApi field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetEnableApi() bool {
	if o == nil || o.EnableApi == nil {
		var ret bool
		return ret
	}
	return *o.EnableApi
}

// GetEnableApiOk returns a tuple with the EnableApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetEnableApiOk() (*bool, bool) {
	if o == nil || o.EnableApi == nil {
		return nil, false
	}
	return o.EnableApi, true
}

// HasEnableApi returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasEnableApi() bool {
	if o != nil && o.EnableApi != nil {
		return true
	}

	return false
}

// SetEnableApi gets a reference to the given bool and assigns it to the EnableApi field.
func (o *PlayerThemeUpdatePayload) SetEnableApi(v bool) {
	o.EnableApi = &v
}

// GetEnableControls returns the EnableControls field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetEnableControls() bool {
	if o == nil || o.EnableControls == nil {
		var ret bool
		return ret
	}
	return *o.EnableControls
}

// GetEnableControlsOk returns a tuple with the EnableControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetEnableControlsOk() (*bool, bool) {
	if o == nil || o.EnableControls == nil {
		return nil, false
	}
	return o.EnableControls, true
}

// HasEnableControls returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasEnableControls() bool {
	if o != nil && o.EnableControls != nil {
		return true
	}

	return false
}

// SetEnableControls gets a reference to the given bool and assigns it to the EnableControls field.
func (o *PlayerThemeUpdatePayload) SetEnableControls(v bool) {
	o.EnableControls = &v
}

// GetForceAutoplay returns the ForceAutoplay field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetForceAutoplay() bool {
	if o == nil || o.ForceAutoplay == nil {
		var ret bool
		return ret
	}
	return *o.ForceAutoplay
}

// GetForceAutoplayOk returns a tuple with the ForceAutoplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetForceAutoplayOk() (*bool, bool) {
	if o == nil || o.ForceAutoplay == nil {
		return nil, false
	}
	return o.ForceAutoplay, true
}

// HasForceAutoplay returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasForceAutoplay() bool {
	if o != nil && o.ForceAutoplay != nil {
		return true
	}

	return false
}

// SetForceAutoplay gets a reference to the given bool and assigns it to the ForceAutoplay field.
func (o *PlayerThemeUpdatePayload) SetForceAutoplay(v bool) {
	o.ForceAutoplay = &v
}

// GetHideTitle returns the HideTitle field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetHideTitle() bool {
	if o == nil || o.HideTitle == nil {
		var ret bool
		return ret
	}
	return *o.HideTitle
}

// GetHideTitleOk returns a tuple with the HideTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetHideTitleOk() (*bool, bool) {
	if o == nil || o.HideTitle == nil {
		return nil, false
	}
	return o.HideTitle, true
}

// HasHideTitle returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasHideTitle() bool {
	if o != nil && o.HideTitle != nil {
		return true
	}

	return false
}

// SetHideTitle gets a reference to the given bool and assigns it to the HideTitle field.
func (o *PlayerThemeUpdatePayload) SetHideTitle(v bool) {
	o.HideTitle = &v
}

// GetForceLoop returns the ForceLoop field value if set, zero value otherwise.
func (o *PlayerThemeUpdatePayload) GetForceLoop() bool {
	if o == nil || o.ForceLoop == nil {
		var ret bool
		return ret
	}
	return *o.ForceLoop
}

// GetForceLoopOk returns a tuple with the ForceLoop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerThemeUpdatePayload) GetForceLoopOk() (*bool, bool) {
	if o == nil || o.ForceLoop == nil {
		return nil, false
	}
	return o.ForceLoop, true
}

// HasForceLoop returns a boolean if a field has been set.
func (o *PlayerThemeUpdatePayload) HasForceLoop() bool {
	if o != nil && o.ForceLoop != nil {
		return true
	}

	return false
}

// SetForceLoop gets a reference to the given bool and assigns it to the ForceLoop field.
func (o *PlayerThemeUpdatePayload) SetForceLoop(v bool) {
	o.ForceLoop = &v
}

type NullablePlayerThemeUpdatePayload struct {
	value *PlayerThemeUpdatePayload
	isSet bool
}

func (v NullablePlayerThemeUpdatePayload) Get() *PlayerThemeUpdatePayload {
	return v.value
}

func (v *NullablePlayerThemeUpdatePayload) Set(val *PlayerThemeUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerThemeUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerThemeUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerThemeUpdatePayload(val *PlayerThemeUpdatePayload) *NullablePlayerThemeUpdatePayload {
	return &NullablePlayerThemeUpdatePayload{value: val, isSet: true}
}
