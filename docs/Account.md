# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quota** | Pointer to [**AccountQuota**](account_quota.md) |  | [optional] 
**Features** | Pointer to **[]string** | Deprecated. What features are enabled for your account. Choices include: app.dynamic_metadata - the ability to dynamically tag videos to better segment and understand your audiences, app.event_log - the ability to create and retrieve a log detailing how your videos were interacted with, player.white_label - the ability to customise your player, stats.player_events - the ability to see statistics about how your player is being used, transcode.mp4_support - the ability to reformat content into mp4 using the H264 codec. | [optional] 
**Environment** | Pointer to **string** | Deprecated. Whether you are using your production or sandbox API key will impact what environment is displayed here, as well as stats and features information. If you use your sandbox key, the environment is \&quot;sandbox.\&quot; If you use your production key, the environment is \&quot;production.\&quot; | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuota

`func (o *Account) GetQuota() AccountQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Account) GetQuotaOk() (*AccountQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Account) SetQuota(v AccountQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *Account) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetFeatures

`func (o *Account) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Account) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Account) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Account) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetEnvironment

`func (o *Account) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Account) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Account) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Account) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


