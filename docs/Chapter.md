# Chapter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **string** |  | [optional] 
**Src** | Pointer to **string** | The link to your VTT file, which contains your chapters information for the video. | [optional] 
**Language** | Pointer to **string** |  | [optional] 

## Methods

### NewChapter

`func NewChapter() *Chapter`

NewChapter instantiates a new Chapter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChapterWithDefaults

`func NewChapterWithDefaults() *Chapter`

NewChapterWithDefaults instantiates a new Chapter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *Chapter) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Chapter) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Chapter) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Chapter) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetSrc

`func (o *Chapter) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *Chapter) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *Chapter) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *Chapter) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetLanguage

`func (o *Chapter) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Chapter) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Chapter) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Chapter) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


