# RefreshTokenPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshToken** | **string** | The refresh token is either the first refresh token you received when you authenticated with the auth/api-key endpoint, or it&#39;s the refresh token from the last time you used the auth/refresh endpoint. Place this in the body of your request to obtain a new access token (which is valid for an hour) and a new refresh token.  | 

## Methods

### NewRefreshTokenPayload

`func NewRefreshTokenPayload(refreshToken string, ) *RefreshTokenPayload`

NewRefreshTokenPayload instantiates a new RefreshTokenPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenPayloadWithDefaults

`func NewRefreshTokenPayloadWithDefaults() *RefreshTokenPayload`

NewRefreshTokenPayloadWithDefaults instantiates a new RefreshTokenPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefreshToken

`func (o *RefreshTokenPayload) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *RefreshTokenPayload) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *RefreshTokenPayload) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


