# PlayerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayerId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | When the player was created, presented in ISO-8601 format. | [optional] 
**UpdatedAt** | Pointer to **string** | When the player was last updated, presented in ISO-8601 format. | [optional] 
**ShapeMargin** | Pointer to **int32** | Deprecated | [optional] 
**ShapeRadius** | Pointer to **int32** | Deprecated | [optional] 
**ShapeAspect** | Pointer to **string** | Deprecated | [optional] 
**ShapeBackgroundTop** | Pointer to **string** | Deprecated | [optional] 
**ShapeBackgroundBottom** | Pointer to **string** | Deprecated | [optional] 
**LinkActive** | Pointer to **string** | Deprecated | [optional] 
**Assets** | Pointer to [**PlayerAllOfAssets**](player_allOf_assets.md) |  | [optional] 

## Methods

### NewPlayerAllOf

`func NewPlayerAllOf() *PlayerAllOf`

NewPlayerAllOf instantiates a new PlayerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerAllOfWithDefaults

`func NewPlayerAllOfWithDefaults() *PlayerAllOf`

NewPlayerAllOfWithDefaults instantiates a new PlayerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayerId

`func (o *PlayerAllOf) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *PlayerAllOf) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *PlayerAllOf) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *PlayerAllOf) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlayerAllOf) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlayerAllOf) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlayerAllOf) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlayerAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PlayerAllOf) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlayerAllOf) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlayerAllOf) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PlayerAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetShapeMargin

`func (o *PlayerAllOf) GetShapeMargin() int32`

GetShapeMargin returns the ShapeMargin field if non-nil, zero value otherwise.

### GetShapeMarginOk

`func (o *PlayerAllOf) GetShapeMarginOk() (*int32, bool)`

GetShapeMarginOk returns a tuple with the ShapeMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeMargin

`func (o *PlayerAllOf) SetShapeMargin(v int32)`

SetShapeMargin sets ShapeMargin field to given value.

### HasShapeMargin

`func (o *PlayerAllOf) HasShapeMargin() bool`

HasShapeMargin returns a boolean if a field has been set.

### GetShapeRadius

`func (o *PlayerAllOf) GetShapeRadius() int32`

GetShapeRadius returns the ShapeRadius field if non-nil, zero value otherwise.

### GetShapeRadiusOk

`func (o *PlayerAllOf) GetShapeRadiusOk() (*int32, bool)`

GetShapeRadiusOk returns a tuple with the ShapeRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeRadius

`func (o *PlayerAllOf) SetShapeRadius(v int32)`

SetShapeRadius sets ShapeRadius field to given value.

### HasShapeRadius

`func (o *PlayerAllOf) HasShapeRadius() bool`

HasShapeRadius returns a boolean if a field has been set.

### GetShapeAspect

`func (o *PlayerAllOf) GetShapeAspect() string`

GetShapeAspect returns the ShapeAspect field if non-nil, zero value otherwise.

### GetShapeAspectOk

`func (o *PlayerAllOf) GetShapeAspectOk() (*string, bool)`

GetShapeAspectOk returns a tuple with the ShapeAspect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeAspect

`func (o *PlayerAllOf) SetShapeAspect(v string)`

SetShapeAspect sets ShapeAspect field to given value.

### HasShapeAspect

`func (o *PlayerAllOf) HasShapeAspect() bool`

HasShapeAspect returns a boolean if a field has been set.

### GetShapeBackgroundTop

`func (o *PlayerAllOf) GetShapeBackgroundTop() string`

GetShapeBackgroundTop returns the ShapeBackgroundTop field if non-nil, zero value otherwise.

### GetShapeBackgroundTopOk

`func (o *PlayerAllOf) GetShapeBackgroundTopOk() (*string, bool)`

GetShapeBackgroundTopOk returns a tuple with the ShapeBackgroundTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBackgroundTop

`func (o *PlayerAllOf) SetShapeBackgroundTop(v string)`

SetShapeBackgroundTop sets ShapeBackgroundTop field to given value.

### HasShapeBackgroundTop

`func (o *PlayerAllOf) HasShapeBackgroundTop() bool`

HasShapeBackgroundTop returns a boolean if a field has been set.

### GetShapeBackgroundBottom

`func (o *PlayerAllOf) GetShapeBackgroundBottom() string`

GetShapeBackgroundBottom returns the ShapeBackgroundBottom field if non-nil, zero value otherwise.

### GetShapeBackgroundBottomOk

`func (o *PlayerAllOf) GetShapeBackgroundBottomOk() (*string, bool)`

GetShapeBackgroundBottomOk returns a tuple with the ShapeBackgroundBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShapeBackgroundBottom

`func (o *PlayerAllOf) SetShapeBackgroundBottom(v string)`

SetShapeBackgroundBottom sets ShapeBackgroundBottom field to given value.

### HasShapeBackgroundBottom

`func (o *PlayerAllOf) HasShapeBackgroundBottom() bool`

HasShapeBackgroundBottom returns a boolean if a field has been set.

### GetLinkActive

`func (o *PlayerAllOf) GetLinkActive() string`

GetLinkActive returns the LinkActive field if non-nil, zero value otherwise.

### GetLinkActiveOk

`func (o *PlayerAllOf) GetLinkActiveOk() (*string, bool)`

GetLinkActiveOk returns a tuple with the LinkActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkActive

`func (o *PlayerAllOf) SetLinkActive(v string)`

SetLinkActive sets LinkActive field to given value.

### HasLinkActive

`func (o *PlayerAllOf) HasLinkActive() bool`

HasLinkActive returns a boolean if a field has been set.

### GetAssets

`func (o *PlayerAllOf) GetAssets() PlayerAllOfAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *PlayerAllOf) GetAssetsOk() (*PlayerAllOfAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *PlayerAllOf) SetAssets(v PlayerAllOfAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *PlayerAllOf) HasAssets() bool`

HasAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


