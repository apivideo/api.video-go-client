# CaptionsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Caption**](Caption.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](pagination.md) |  | [optional] 

## Methods

### NewCaptionsListResponse

`func NewCaptionsListResponse() *CaptionsListResponse`

NewCaptionsListResponse instantiates a new CaptionsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptionsListResponseWithDefaults

`func NewCaptionsListResponseWithDefaults() *CaptionsListResponse`

NewCaptionsListResponseWithDefaults instantiates a new CaptionsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CaptionsListResponse) GetData() []Caption`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CaptionsListResponse) GetDataOk() (*[]Caption, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CaptionsListResponse) SetData(v []Caption)`

SetData sets Data field to given value.

### HasData

`func (o *CaptionsListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *CaptionsListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CaptionsListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CaptionsListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *CaptionsListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


