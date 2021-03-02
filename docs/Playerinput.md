# Playerinput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | RGBA color for timer text. Default: rgba(255, 255, 255, 1) | [optional] 
**Link** | Pointer to **string** | RGBA color for all controls. Default: rgba(255, 255, 255, 1) | [optional] 
**LinkHover** | Pointer to **string** | RGBA color for all controls when hovered. Default: rgba(255, 255, 255, 1) | [optional] 
**TrackPlayed** | Pointer to **string** | RGBA color playback bar: played content. Default: rgba(88, 131, 255, .95) | [optional] 
**TrackUnplayed** | Pointer to **string** | RGBA color playback bar: downloaded but unplayed (buffered) content. Default: rgba(255, 255, 255, .35) | [optional] 
**TrackBackground** | Pointer to **string** | RGBA color playback bar: background. Default: rgba(255, 255, 255, .2) | [optional] 
**BackgroundTop** | Pointer to **string** | RGBA color: top 50% of background. Default: rgba(0, 0, 0, .7) | [optional] 
**BackgroundBottom** | Pointer to **string** | RGBA color: bottom 50% of background. Default: rgba(0, 0, 0, .7) | [optional] 
**BackgroundText** | Pointer to **string** | RGBA color for title text. Default: rgba(255, 255, 255, 1) | [optional] 
**EnableApi** | Pointer to **bool** | enable/disable player SDK access. Default: true | [optional] [default to true]
**EnableControls** | Pointer to **bool** | enable/disable player controls. Default: true | [optional] [default to true]
**ForceAutoplay** | Pointer to **bool** | enable/disable player autoplay. Default: false | [optional] [default to false]
**HideTitle** | Pointer to **bool** | enable/disable title. Default: false | [optional] [default to false]
**ForceLoop** | Pointer to **bool** | enable/disable looping. Default: false | [optional] [default to false]

## Methods

### NewPlayerinput

`func NewPlayerinput() *Playerinput`

NewPlayerinput instantiates a new Playerinput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerinputWithDefaults

`func NewPlayerinputWithDefaults() *Playerinput`

NewPlayerinputWithDefaults instantiates a new Playerinput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *Playerinput) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Playerinput) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Playerinput) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Playerinput) HasText() bool`

HasText returns a boolean if a field has been set.

### GetLink

`func (o *Playerinput) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Playerinput) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Playerinput) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Playerinput) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetLinkHover

`func (o *Playerinput) GetLinkHover() string`

GetLinkHover returns the LinkHover field if non-nil, zero value otherwise.

### GetLinkHoverOk

`func (o *Playerinput) GetLinkHoverOk() (*string, bool)`

GetLinkHoverOk returns a tuple with the LinkHover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkHover

`func (o *Playerinput) SetLinkHover(v string)`

SetLinkHover sets LinkHover field to given value.

### HasLinkHover

`func (o *Playerinput) HasLinkHover() bool`

HasLinkHover returns a boolean if a field has been set.

### GetTrackPlayed

`func (o *Playerinput) GetTrackPlayed() string`

GetTrackPlayed returns the TrackPlayed field if non-nil, zero value otherwise.

### GetTrackPlayedOk

`func (o *Playerinput) GetTrackPlayedOk() (*string, bool)`

GetTrackPlayedOk returns a tuple with the TrackPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackPlayed

`func (o *Playerinput) SetTrackPlayed(v string)`

SetTrackPlayed sets TrackPlayed field to given value.

### HasTrackPlayed

`func (o *Playerinput) HasTrackPlayed() bool`

HasTrackPlayed returns a boolean if a field has been set.

### GetTrackUnplayed

`func (o *Playerinput) GetTrackUnplayed() string`

GetTrackUnplayed returns the TrackUnplayed field if non-nil, zero value otherwise.

### GetTrackUnplayedOk

`func (o *Playerinput) GetTrackUnplayedOk() (*string, bool)`

GetTrackUnplayedOk returns a tuple with the TrackUnplayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackUnplayed

`func (o *Playerinput) SetTrackUnplayed(v string)`

SetTrackUnplayed sets TrackUnplayed field to given value.

### HasTrackUnplayed

`func (o *Playerinput) HasTrackUnplayed() bool`

HasTrackUnplayed returns a boolean if a field has been set.

### GetTrackBackground

`func (o *Playerinput) GetTrackBackground() string`

GetTrackBackground returns the TrackBackground field if non-nil, zero value otherwise.

### GetTrackBackgroundOk

`func (o *Playerinput) GetTrackBackgroundOk() (*string, bool)`

GetTrackBackgroundOk returns a tuple with the TrackBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackBackground

`func (o *Playerinput) SetTrackBackground(v string)`

SetTrackBackground sets TrackBackground field to given value.

### HasTrackBackground

`func (o *Playerinput) HasTrackBackground() bool`

HasTrackBackground returns a boolean if a field has been set.

### GetBackgroundTop

`func (o *Playerinput) GetBackgroundTop() string`

GetBackgroundTop returns the BackgroundTop field if non-nil, zero value otherwise.

### GetBackgroundTopOk

`func (o *Playerinput) GetBackgroundTopOk() (*string, bool)`

GetBackgroundTopOk returns a tuple with the BackgroundTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundTop

`func (o *Playerinput) SetBackgroundTop(v string)`

SetBackgroundTop sets BackgroundTop field to given value.

### HasBackgroundTop

`func (o *Playerinput) HasBackgroundTop() bool`

HasBackgroundTop returns a boolean if a field has been set.

### GetBackgroundBottom

`func (o *Playerinput) GetBackgroundBottom() string`

GetBackgroundBottom returns the BackgroundBottom field if non-nil, zero value otherwise.

### GetBackgroundBottomOk

`func (o *Playerinput) GetBackgroundBottomOk() (*string, bool)`

GetBackgroundBottomOk returns a tuple with the BackgroundBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundBottom

`func (o *Playerinput) SetBackgroundBottom(v string)`

SetBackgroundBottom sets BackgroundBottom field to given value.

### HasBackgroundBottom

`func (o *Playerinput) HasBackgroundBottom() bool`

HasBackgroundBottom returns a boolean if a field has been set.

### GetBackgroundText

`func (o *Playerinput) GetBackgroundText() string`

GetBackgroundText returns the BackgroundText field if non-nil, zero value otherwise.

### GetBackgroundTextOk

`func (o *Playerinput) GetBackgroundTextOk() (*string, bool)`

GetBackgroundTextOk returns a tuple with the BackgroundText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundText

`func (o *Playerinput) SetBackgroundText(v string)`

SetBackgroundText sets BackgroundText field to given value.

### HasBackgroundText

`func (o *Playerinput) HasBackgroundText() bool`

HasBackgroundText returns a boolean if a field has been set.

### GetEnableApi

`func (o *Playerinput) GetEnableApi() bool`

GetEnableApi returns the EnableApi field if non-nil, zero value otherwise.

### GetEnableApiOk

`func (o *Playerinput) GetEnableApiOk() (*bool, bool)`

GetEnableApiOk returns a tuple with the EnableApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApi

`func (o *Playerinput) SetEnableApi(v bool)`

SetEnableApi sets EnableApi field to given value.

### HasEnableApi

`func (o *Playerinput) HasEnableApi() bool`

HasEnableApi returns a boolean if a field has been set.

### GetEnableControls

`func (o *Playerinput) GetEnableControls() bool`

GetEnableControls returns the EnableControls field if non-nil, zero value otherwise.

### GetEnableControlsOk

`func (o *Playerinput) GetEnableControlsOk() (*bool, bool)`

GetEnableControlsOk returns a tuple with the EnableControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableControls

`func (o *Playerinput) SetEnableControls(v bool)`

SetEnableControls sets EnableControls field to given value.

### HasEnableControls

`func (o *Playerinput) HasEnableControls() bool`

HasEnableControls returns a boolean if a field has been set.

### GetForceAutoplay

`func (o *Playerinput) GetForceAutoplay() bool`

GetForceAutoplay returns the ForceAutoplay field if non-nil, zero value otherwise.

### GetForceAutoplayOk

`func (o *Playerinput) GetForceAutoplayOk() (*bool, bool)`

GetForceAutoplayOk returns a tuple with the ForceAutoplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAutoplay

`func (o *Playerinput) SetForceAutoplay(v bool)`

SetForceAutoplay sets ForceAutoplay field to given value.

### HasForceAutoplay

`func (o *Playerinput) HasForceAutoplay() bool`

HasForceAutoplay returns a boolean if a field has been set.

### GetHideTitle

`func (o *Playerinput) GetHideTitle() bool`

GetHideTitle returns the HideTitle field if non-nil, zero value otherwise.

### GetHideTitleOk

`func (o *Playerinput) GetHideTitleOk() (*bool, bool)`

GetHideTitleOk returns a tuple with the HideTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTitle

`func (o *Playerinput) SetHideTitle(v bool)`

SetHideTitle sets HideTitle field to given value.

### HasHideTitle

`func (o *Playerinput) HasHideTitle() bool`

HasHideTitle returns a boolean if a field has been set.

### GetForceLoop

`func (o *Playerinput) GetForceLoop() bool`

GetForceLoop returns the ForceLoop field if non-nil, zero value otherwise.

### GetForceLoopOk

`func (o *Playerinput) GetForceLoopOk() (*bool, bool)`

GetForceLoopOk returns a tuple with the ForceLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceLoop

`func (o *Playerinput) SetForceLoop(v bool)`

SetForceLoop sets ForceLoop field to given value.

### HasForceLoop

`func (o *Playerinput) HasForceLoop() bool`

HasForceLoop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


