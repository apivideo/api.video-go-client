# TokenCreationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | Pointer to **int32** | Time in seconds that the token will be active. A value of 0 means that the token has no exipration date. The default is to have no expiration. | [optional] [default to 0]

## Methods

### NewTokenCreationPayload

`func NewTokenCreationPayload() *TokenCreationPayload`

NewTokenCreationPayload instantiates a new TokenCreationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreationPayloadWithDefaults

`func NewTokenCreationPayloadWithDefaults() *TokenCreationPayload`

NewTokenCreationPayloadWithDefaults instantiates a new TokenCreationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *TokenCreationPayload) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *TokenCreationPayload) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *TokenCreationPayload) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *TokenCreationPayload) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


