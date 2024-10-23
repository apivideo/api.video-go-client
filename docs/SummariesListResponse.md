# SummariesListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Summary**](Summary.md) | An array of summary objects. | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewSummariesListResponse

`func NewSummariesListResponse(data []Summary, pagination Pagination, ) *SummariesListResponse`

NewSummariesListResponse instantiates a new SummariesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummariesListResponseWithDefaults

`func NewSummariesListResponseWithDefaults() *SummariesListResponse`

NewSummariesListResponseWithDefaults instantiates a new SummariesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SummariesListResponse) GetData() []Summary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SummariesListResponse) GetDataOk() (*[]Summary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SummariesListResponse) SetData(v []Summary)`

SetData sets Data field to given value.


### GetPagination

`func (o *SummariesListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SummariesListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SummariesListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


