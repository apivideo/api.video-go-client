# TokenListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]UploadToken**](UploadToken.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewTokenListResponse

`func NewTokenListResponse(data []UploadToken, pagination Pagination, ) *TokenListResponse`

NewTokenListResponse instantiates a new TokenListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenListResponseWithDefaults

`func NewTokenListResponseWithDefaults() *TokenListResponse`

NewTokenListResponseWithDefaults instantiates a new TokenListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokenListResponse) GetData() []UploadToken`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenListResponse) GetDataOk() (*[]UploadToken, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenListResponse) SetData(v []UploadToken)`

SetData sets Data field to given value.


### GetPagination

`func (o *TokenListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TokenListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TokenListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


