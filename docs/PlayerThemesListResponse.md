# PlayerThemesListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PlayerTheme**](PlayerTheme.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](pagination.md) |  | [optional] 

## Methods

### NewPlayerThemesListResponse

`func NewPlayerThemesListResponse() *PlayerThemesListResponse`

NewPlayerThemesListResponse instantiates a new PlayerThemesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerThemesListResponseWithDefaults

`func NewPlayerThemesListResponseWithDefaults() *PlayerThemesListResponse`

NewPlayerThemesListResponseWithDefaults instantiates a new PlayerThemesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PlayerThemesListResponse) GetData() []PlayerTheme`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PlayerThemesListResponse) GetDataOk() (*[]PlayerTheme, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PlayerThemesListResponse) SetData(v []PlayerTheme)`

SetData sets Data field to given value.

### HasData

`func (o *PlayerThemesListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *PlayerThemesListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PlayerThemesListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PlayerThemesListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PlayerThemesListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


