# Caption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **string** |  | [optional] 
**Src** | Pointer to **string** |  | [optional] 
**Srclang** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** | Whether you will have subtitles or not. True for yes you will have subtitles, false for no you will not have subtitles. | [optional] [default to false]

## Methods

### NewCaption

`func NewCaption() *Caption`

NewCaption instantiates a new Caption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptionWithDefaults

`func NewCaptionWithDefaults() *Caption`

NewCaptionWithDefaults instantiates a new Caption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *Caption) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Caption) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Caption) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Caption) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetSrc

`func (o *Caption) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *Caption) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *Caption) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *Caption) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetSrclang

`func (o *Caption) GetSrclang() string`

GetSrclang returns the Srclang field if non-nil, zero value otherwise.

### GetSrclangOk

`func (o *Caption) GetSrclangOk() (*string, bool)`

GetSrclangOk returns a tuple with the Srclang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrclang

`func (o *Caption) SetSrclang(v string)`

SetSrclang sets Srclang field to given value.

### HasSrclang

`func (o *Caption) HasSrclang() bool`

HasSrclang returns a boolean if a field has been set.

### GetDefault

`func (o *Caption) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Caption) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Caption) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Caption) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


