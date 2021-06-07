# ChaptersListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Chapter**](Chapter.md) |  | 
**Pagination** | [**Pagination**](pagination.md) |  | 

## Methods

### NewChaptersListResponse

`func NewChaptersListResponse(data []Chapter, pagination Pagination, ) *ChaptersListResponse`

NewChaptersListResponse instantiates a new ChaptersListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChaptersListResponseWithDefaults

`func NewChaptersListResponseWithDefaults() *ChaptersListResponse`

NewChaptersListResponseWithDefaults instantiates a new ChaptersListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ChaptersListResponse) GetData() []Chapter`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChaptersListResponse) GetDataOk() (*[]Chapter, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChaptersListResponse) SetData(v []Chapter)`

SetData sets Data field to given value.


### GetPagination

`func (o *ChaptersListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ChaptersListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ChaptersListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


