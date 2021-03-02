# PlayersListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Player**](Player.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](pagination.md) |  | [optional] 

## Methods

### NewPlayersListResponse

`func NewPlayersListResponse() *PlayersListResponse`

NewPlayersListResponse instantiates a new PlayersListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayersListResponseWithDefaults

`func NewPlayersListResponseWithDefaults() *PlayersListResponse`

NewPlayersListResponseWithDefaults instantiates a new PlayersListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PlayersListResponse) GetData() []Player`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PlayersListResponse) GetDataOk() (*[]Player, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PlayersListResponse) SetData(v []Player)`

SetData sets Data field to given value.

### HasData

`func (o *PlayersListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *PlayersListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PlayersListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PlayersListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PlayersListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


