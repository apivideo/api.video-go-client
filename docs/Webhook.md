# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | Pointer to **string** | A unique identifier of the webhook you subscribed to. | [optional] 
**CreatedAt** | Pointer to **string** | The time and date when you created this webhook subscription, in ATOM UTC format. | [optional] 
**Events** | Pointer to **[]string** | A list of events that you subscribed to. When these events occur, the API triggers a webhook call to the URL you provided. | [optional] 
**Url** | Pointer to **string** | The URL where the API sends the webhook. | [optional] 
**SignatureSecret** | Pointer to **string** | A secret key for the webhook you subscribed to. You can use it to verify the origin of the webhook call that you receive. | [optional] 

## Methods

### NewWebhook

`func NewWebhook() *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *Webhook) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *Webhook) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *Webhook) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *Webhook) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Webhook) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Webhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvents

`func (o *Webhook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Webhook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Webhook) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Webhook) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSignatureSecret

`func (o *Webhook) GetSignatureSecret() string`

GetSignatureSecret returns the SignatureSecret field if non-nil, zero value otherwise.

### GetSignatureSecretOk

`func (o *Webhook) GetSignatureSecretOk() (*string, bool)`

GetSignatureSecretOk returns a tuple with the SignatureSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureSecret

`func (o *Webhook) SetSignatureSecret(v string)`

SetSignatureSecret sets SignatureSecret field to given value.

### HasSignatureSecret

`func (o *Webhook) HasSignatureSecret() bool`

HasSignatureSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


