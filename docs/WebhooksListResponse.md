# WebhooksListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Webhook**](Webhook.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewWebhooksListResponse

`func NewWebhooksListResponse(data []Webhook, pagination Pagination, ) *WebhooksListResponse`

NewWebhooksListResponse instantiates a new WebhooksListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksListResponseWithDefaults

`func NewWebhooksListResponseWithDefaults() *WebhooksListResponse`

NewWebhooksListResponseWithDefaults instantiates a new WebhooksListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebhooksListResponse) GetData() []Webhook`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhooksListResponse) GetDataOk() (*[]Webhook, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhooksListResponse) SetData(v []Webhook)`

SetData sets Data field to given value.


### GetPagination

`func (o *WebhooksListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *WebhooksListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *WebhooksListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


