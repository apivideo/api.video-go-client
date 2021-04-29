# Player

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
**EnableApi** | Pointer to **bool** | enable/disable player SDK access. Default: true | [optional] 
**EnableControls** | Pointer to **bool** | enable/disable player controls. Default: true | [optional] 
**ForceAutoplay** | Pointer to **bool** | enable/disable player autoplay. Default: false | [optional] 
**HideTitle** | Pointer to **bool** | enable/disable title. Default: false | [optional] 
**ForceLoop** | Pointer to **bool** | enable/disable looping. Default: false | [optional] 
**PlayerId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | When the player was created, presented in ISO-8601 format. | [optional] 
**UpdatedAt** | Pointer to **string** | When the player was last updated, presented in ISO-8601 format. | [optional] 
**ShapeMargin** | Pointer to **int32** | Deprecated | [optional] 
**ShapeRadius** | Pointer to **int32** | Deprecated | [optional] 
**ShapeAspect** | Pointer to **string** | Deprecated | [optional] 
**ShapeBackgroundTop** | Pointer to **string** | Deprecated | [optional] 
**ShapeBackgroundBottom** | Pointer to **string** | Deprecated | [optional] 
**LinkActive** | Pointer to **string** | Deprecated | [optional] 
**Assets** | Pointer to [**PlayerAssets**](player_assets.md) |  | [optional] 

## Methods

### NewPlayer

`func NewPlayer() *Player`

NewPlayer instantiates a new Player object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerWithDefaults

`func NewPlayerWithDefaults() *Player`

NewPlayerWithDefaults instantiates a new Player object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *Player) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Player) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Player) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Player) HasText() bool`

HasText returns a boolean if a field has been set.

### GetLink

`func (o *Player) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Player) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Player) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Player) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetLinkHover

`func (o *Player) GetLinkHover() string`

GetLinkHover returns the LinkHover field if non-nil, zero value otherwise.

### GetLinkHoverOk

`func (o *Player) GetLinkHoverOk() (*string, bool)`

GetLinkHoverOk returns a tuple with the LinkHover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkHover

`func (o *Player) SetLinkHover(v string)`

SetLinkHover sets LinkHover field to given value.

### HasLinkHover

`func (o *Player) HasLinkHover() bool`

HasLinkHover returns a boolean if a field has been set.

### GetTrackPlayed

`func (o *Player) GetTrackPlayed() string`

GetTrackPlayed returns the TrackPlayed field if non-nil, zero value otherwise.

### GetTrackPlayedOk

`func (o *Player) GetTrackPlayedOk() (*string, bool)`

GetTrackPlayedOk returns a tuple with the TrackPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackPlayed

`func (o *Player) SetTrackPlayed(v string)`

SetTrackPlayed sets TrackPlayed field to given value.

### HasTrackPlayed

`func (o *Player) HasTrackPlayed() bool`

HasTrackPlayed returns a boolean if a field has been set.

### GetTrackUnplayed

`func (o *Player) GetTrackUnplayed() string`

GetTrackUnplayed returns the TrackUnplayed field if non-nil, zero value otherwise.

### GetTrackUnplayedOk

`func (o *Player) GetTrackUnplayedOk() (*string, bool)`

GetTrackUnplayedOk returns a tuple with the TrackUnplayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackUnplayed

`func (o *Player) SetTrackUnplayed(v string)`

SetTrackUnplayed sets TrackUnplayed field to given value.

### HasTrackUnplayed

`func (o *Player) HasTrackUnplayed() bool`

HasTrackUnplayed returns a boolean if a field has been set.

### GetTrackBackground

`func (o *Player) GetTrackBackground() string`

GetTrackBackground returns the TrackBackground field if non-nil, zero value otherwise.

### GetTrackBackgroundOk

`func (o *Player) GetTrackBackgroundOk() (*string, bool)`

GetTrackBackgroundOk returns a tuple with the TrackBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackBackground

`func (o *Player) SetTrackBackground(v string)`

SetTrackBackground sets TrackBackground field to given value.

### HasTrackBackground

`func (o *Player) HasTrackBackground() bool`

HasTrackBackground returns a boolean if a field has been set.

### GetBackgroundTop

`func (o *Player) GetBackgroundTop() string`

GetBackgroundTop returns the BackgroundTop field if non-nil, zero value otherwise.

### GetBackgroundTopOk

`func (o *Player) GetBackgroundTopOk() (*string, bool)`

GetBackgroundTopOk returns a tuple with the BackgroundTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundTop

`func (o *Player) SetBackgroundTop(v string)`

SetBackgroundTop sets BackgroundTop field to given value.

### HasBackgroundTop

`func (o *Player) HasBackgroundTop() bool`

HasBackgroundTop returns a boolean if a field has been set.

### GetBackgroundBottom

`func (o *Player) GetBackgroundBottom() string`

GetBackgroundBottom returns the BackgroundBottom field if non-nil, zero value otherwise.

### GetBackgroundBottomOk

`func (o *Player) GetBackgroundBottomOk() (*string, bool)`

GetBackgroundBottomOk returns a tuple with the BackgroundBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundBottom

`func (o *Player) SetBackgroundBottom(v string)`

SetBackgroundBottom sets BackgroundBottom field to given value.

### HasBackgroundBottom

`func (o *Player) HasBackgroundBottom() bool`

HasBackgroundBottom returns a boolean if a field has been set.

### GetBackgroundText

`func (o *Player) GetBackgroundText() string`

GetBackgroundText returns the BackgroundText field if non-nil, zero value otherwise.

### GetBackgroundTextOk

`func (o *Player) GetBackgroundTextOk() (*string, bool)`

GetBackgroundTextOk returns a tuple with the BackgroundText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundText

`func (o *Player) SetBackgroundText(v string)`

SetBackgroundText sets BackgroundText field to given value.

### HasBackgroundText

`func (o *Player) HasBackgroundText() bool`

HasBackgroundText returns a boolean if a field has been set.

### GetEnableApi

`func (o *Player) GetEnableApi() bool`

GetEnableApi returns the EnableApi field if non-nil, zero value otherwise.

### GetEnableApiOk

`func (o *Player) GetEnableApiOk() (*bool, bool)`

GetEnableApiOk returns a tuple with the EnableApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApi

`func (o *Player) SetEnableApi(v bool)`

SetEnableApi sets EnableApi field to given value.

### HasEnableApi

`func (o *Player) HasEnableApi() bool`

HasEnableApi returns a boolean if a field has been set.

### GetEnableControls

`func (o *Player) GetEnableControls() bool`

GetEnableControls returns the EnableControls field if non-nil, zero value otherwise.

### GetEnableControlsOk

`func (o *Player) GetEnableControlsOk() (*bool, bool)`

GetEnableControlsOk returns a tuple with the EnableControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableControls

`func (o *Player) SetEnableControls(v bool)`

SetEnableControls sets EnableControls field to given value.

### HasEnableControls

`func (o *Player) HasEnableControls() bool`

HasEnableControls returns a boolean if a field has been set.

### GetForceAutoplay

`func (o *Player) GetForceAutoplay() bool`

GetForceAutoplay returns the ForceAutoplay field if non-nil, zero value otherwise.

### GetForceAutoplayOk

`func (o *Player) GetForceAutoplayOk() (*bool, bool)`

GetForceAutoplayOk returns a tuple with the ForceAutoplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAutoplay

`func (o *Player) SetForceAutoplay(v bool)`

SetForceAutoplay sets ForceAutoplay field to given value.

### HasForceAutoplay

`func (o *Player) HasForceAutoplay() bool`

HasForceAutoplay returns a boolean if a field has been set.

### GetHideTitle

`func (o *Player) GetHideTitle() bool`

GetHideTitle returns the HideTitle field if non-nil, zero value otherwise.

### GetHideTitleOk

`func (o *Player) GetHideTitleOk() (*bool, bool)`

GetHideTitleOk returns a tuple with the HideTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTitle

`func (o *Player) SetHideTitle(v bool)`

SetHideTitle sets HideTitle field to given value.

### HasHideTitle

`func (o *Player) HasHideTitle() bool`

HasHideTitle returns a boolean if a field has been set.

### GetForceLoop

`func (o *Player) GetForceLoop() bool`

GetForceLoop returns the ForceLoop field if non-nil, zero value otherwise.

### GetForceLoopOk

`func (o *Player) GetForceLoopOk() (*bool, bool)`

GetForceLoopOk returns a tuple with the ForceLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceLoop

`func (o *Player) SetForceLoop(v bool)`

SetForceLoop sets ForceLoop field to given value.

### HasForceLoop

`func (o *Player) HasForceLoop() bool`

HasForceLoop returns a boolean if a field has been set.

### GetPlayerId

`func (o *Player) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *Player) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *Player) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *Player) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Player) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Player) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Player) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Player) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Player) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Player) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Player) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Player) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetShapeMargin

`func (o *Player) GetShapeMargin() int32`

GetShapeMargin returns the ShapeMargin field if non-nil, zero value otherwise.

### GetShapeMarginOk

`func (o *Player) GetShapeMarginOk() (*int32, bool)`

GetShapeMarginOk returns a tuple with the ShapeMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeMargin

`func (o *Player) SetShapeMargin(v int32)`

SetShapeMargin sets ShapeMargin field to given value.

### HasShapeMargin

`func (o *Player) HasShapeMargin() bool`

HasShapeMargin returns a boolean if a field has been set.

### GetShapeRadius

`func (o *Player) GetShapeRadius() int32`

GetShapeRadius returns the ShapeRadius field if non-nil, zero value otherwise.

### GetShapeRadiusOk

`func (o *Player) GetShapeRadiusOk() (*int32, bool)`

GetShapeRadiusOk returns a tuple with the ShapeRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeRadius

`func (o *Player) SetShapeRadius(v int32)`

SetShapeRadius sets ShapeRadius field to given value.

### HasShapeRadius

`func (o *Player) HasShapeRadius() bool`

HasShapeRadius returns a boolean if a field has been set.

### GetShapeAspect

`func (o *Player) GetShapeAspect() string`

GetShapeAspect returns the ShapeAspect field if non-nil, zero value otherwise.

### GetShapeAspectOk

`func (o *Player) GetShapeAspectOk() (*string, bool)`

GetShapeAspectOk returns a tuple with the ShapeAspect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeAspect

`func (o *Player) SetShapeAspect(v string)`

SetShapeAspect sets ShapeAspect field to given value.

### HasShapeAspect

`func (o *Player) HasShapeAspect() bool`

HasShapeAspect returns a boolean if a field has been set.

### GetShapeBackgroundTop

`func (o *Player) GetShapeBackgroundTop() string`

GetShapeBackgroundTop returns the ShapeBackgroundTop field if non-nil, zero value otherwise.

### GetShapeBackgroundTopOk

`func (o *Player) GetShapeBackgroundTopOk() (*string, bool)`

GetShapeBackgroundTopOk returns a tuple with the ShapeBackgroundTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBackgroundTop

`func (o *Player) SetShapeBackgroundTop(v string)`

SetShapeBackgroundTop sets ShapeBackgroundTop field to given value.

### HasShapeBackgroundTop

`func (o *Player) HasShapeBackgroundTop() bool`

HasShapeBackgroundTop returns a boolean if a field has been set.

### GetShapeBackgroundBottom

`func (o *Player) GetShapeBackgroundBottom() string`

GetShapeBackgroundBottom returns the ShapeBackgroundBottom field if non-nil, zero value otherwise.

### GetShapeBackgroundBottomOk

`func (o *Player) GetShapeBackgroundBottomOk() (*string, bool)`

GetShapeBackgroundBottomOk returns a tuple with the ShapeBackgroundBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBackgroundBottom

`func (o *Player) SetShapeBackgroundBottom(v string)`

SetShapeBackgroundBottom sets ShapeBackgroundBottom field to given value.

### HasShapeBackgroundBottom

`func (o *Player) HasShapeBackgroundBottom() bool`

HasShapeBackgroundBottom returns a boolean if a field has been set.

### GetLinkActive

`func (o *Player) GetLinkActive() string`

GetLinkActive returns the LinkActive field if non-nil, zero value otherwise.

### GetLinkActiveOk

`func (o *Player) GetLinkActiveOk() (*string, bool)`

GetLinkActiveOk returns a tuple with the LinkActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkActive

`func (o *Player) SetLinkActive(v string)`

SetLinkActive sets LinkActive field to given value.

### HasLinkActive

`func (o *Player) HasLinkActive() bool`

HasLinkActive returns a boolean if a field has been set.

### GetAssets

`func (o *Player) GetAssets() PlayerAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Player) GetAssetsOk() (*PlayerAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Player) SetAssets(v PlayerAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *Player) HasAssets() bool`

HasAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


