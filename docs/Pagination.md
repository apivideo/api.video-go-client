# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemsTotal** | Pointer to **int32** | Total number of items that exist. | [optional] [readonly] 
**PagesTotal** | Pointer to **int32** | Number of items listed in the current page. | [optional] [readonly] 
**PageSize** | Pointer to **int32** | Maximum number of item per page. | [optional] [readonly] 
**CurrentPage** | Pointer to **int32** | The current page index. | [optional] [readonly] 
**CurrentPageItems** | Pointer to **int32** | The number of items on the current page. | [optional] [readonly] 
**Links** | [**[]PaginationLink**](PaginationLink.md) |  | 

## Methods

### NewPagination

`func NewPagination(links []PaginationLink, ) *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemsTotal

`func (o *Pagination) GetItemsTotal() int32`

GetItemsTotal returns the ItemsTotal field if non-nil, zero value otherwise.

### GetItemsTotalOk

`func (o *Pagination) GetItemsTotalOk() (*int32, bool)`

GetItemsTotalOk returns a tuple with the ItemsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTotal

`func (o *Pagination) SetItemsTotal(v int32)`

SetItemsTotal sets ItemsTotal field to given value.

### HasItemsTotal

`func (o *Pagination) HasItemsTotal() bool`

HasItemsTotal returns a boolean if a field has been set.

### GetPagesTotal

`func (o *Pagination) GetPagesTotal() int32`

GetPagesTotal returns the PagesTotal field if non-nil, zero value otherwise.

### GetPagesTotalOk

`func (o *Pagination) GetPagesTotalOk() (*int32, bool)`

GetPagesTotalOk returns a tuple with the PagesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesTotal

`func (o *Pagination) SetPagesTotal(v int32)`

SetPagesTotal sets PagesTotal field to given value.

### HasPagesTotal

`func (o *Pagination) HasPagesTotal() bool`

HasPagesTotal returns a boolean if a field has been set.

### GetPageSize

`func (o *Pagination) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Pagination) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Pagination) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Pagination) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetCurrentPage

`func (o *Pagination) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *Pagination) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *Pagination) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *Pagination) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentPageItems

`func (o *Pagination) GetCurrentPageItems() int32`

GetCurrentPageItems returns the CurrentPageItems field if non-nil, zero value otherwise.

### GetCurrentPageItemsOk

`func (o *Pagination) GetCurrentPageItemsOk() (*int32, bool)`

GetCurrentPageItemsOk returns a tuple with the CurrentPageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageItems

`func (o *Pagination) SetCurrentPageItems(v int32)`

SetCurrentPageItems sets CurrentPageItems field to given value.

### HasCurrentPageItems

`func (o *Pagination) HasCurrentPageItems() bool`

HasCurrentPageItems returns a boolean if a field has been set.

### GetLinks

`func (o *Pagination) GetLinks() []PaginationLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Pagination) GetLinksOk() (*[]PaginationLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Pagination) SetLinks(v []PaginationLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


