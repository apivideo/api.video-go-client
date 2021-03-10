# UploadToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The unique identifier for the token you will use to authenticate an upload. | [optional] 
**Ttl** | Pointer to **int32** | Time-to-live - how long the upload token is valid for. | [optional] 
**CreatedAt** | Pointer to **string** | When the token was created, displayed in ISO-8601 format. | [optional] 
**ExpiresAt** | Pointer to **string** | When the token expires, displayed in ISO-8601 format. | [optional] 

## Methods

### NewUploadToken

`func NewUploadToken() *UploadToken`

NewUploadToken instantiates a new UploadToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadTokenWithDefaults

`func NewUploadTokenWithDefaults() *UploadToken`

NewUploadTokenWithDefaults instantiates a new UploadToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *UploadToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UploadToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UploadToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UploadToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *UploadToken) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *UploadToken) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *UploadToken) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *UploadToken) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UploadToken) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UploadToken) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UploadToken) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UploadToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *UploadToken) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UploadToken) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UploadToken) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UploadToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


