# AuthenticatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | Your account API key. You can use your sandbox API key, or you can use your production API key. | 

## Methods

### NewAuthenticatePayload

`func NewAuthenticatePayload(apiKey string, ) *AuthenticatePayload`

NewAuthenticatePayload instantiates a new AuthenticatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatePayloadWithDefaults

`func NewAuthenticatePayloadWithDefaults() *AuthenticatePayload`

NewAuthenticatePayloadWithDefaults instantiates a new AuthenticatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *AuthenticatePayload) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AuthenticatePayload) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AuthenticatePayload) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


