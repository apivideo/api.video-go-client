# FilterBy2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaId** | Pointer to **[]string** | Returns analytics based on the unique identifiers of a video or a live stream. | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**Continent** | Pointer to **[]string** | Returns analytics based on the viewers&#39; continent. The list of supported continents names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example &#x60;EU&#x60;. | [optional] 
**Country** | Pointer to **[]string** | Returns analytics based on the viewers&#39; country. The list of supported country names are based on the [GeoNames public database](https://www.geonames.org/countries/). You must use the ISO-3166 alpha2 format, for example &#x60;FR&#x60;. | [optional] 
**DeviceType** | Pointer to **[]string** | Returns analytics based on the type of device used by the viewers. Response values can include: &#x60;computer&#x60;, &#x60;phone&#x60;, &#x60;tablet&#x60;, &#x60;tv&#x60;, &#x60;console&#x60;, &#x60;wearable&#x60;, &#x60;unknown&#x60;. | [optional] 
**OperatingSystem** | Pointer to **[]string** | Returns analytics based on the operating system used by the viewers. Response values can include &#x60;windows&#x60;, &#x60;mac osx&#x60;, &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;linux&#x60;. | [optional] 
**Browser** | Pointer to **[]string** | Returns analytics based on the browser used by the viewers. Response values can include &#x60;chrome&#x60;, &#x60;firefox&#x60;, &#x60;edge&#x60;, &#x60;opera&#x60;. | [optional] 
**Tag** | Pointer to **string** | Returns analytics for videos using this tag. This filter only accepts a single value and is case sensitive. Read more about tagging your videos [here](https://docs.api.video/vod/tags-metadata). | [optional] 

## Methods

### NewFilterBy2

`func NewFilterBy2() *FilterBy2`

NewFilterBy2 instantiates a new FilterBy2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterBy2WithDefaults

`func NewFilterBy2WithDefaults() *FilterBy2`

NewFilterBy2WithDefaults instantiates a new FilterBy2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaId

`func (o *FilterBy2) GetMediaId() []string`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *FilterBy2) GetMediaIdOk() (*[]string, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *FilterBy2) SetMediaId(v []string)`

SetMediaId sets MediaId field to given value.

### HasMediaId

`func (o *FilterBy2) HasMediaId() bool`

HasMediaId returns a boolean if a field has been set.

### GetMediaType

`func (o *FilterBy2) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *FilterBy2) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *FilterBy2) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *FilterBy2) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetContinent

`func (o *FilterBy2) GetContinent() []string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *FilterBy2) GetContinentOk() (*[]string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *FilterBy2) SetContinent(v []string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *FilterBy2) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### GetCountry

`func (o *FilterBy2) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FilterBy2) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FilterBy2) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *FilterBy2) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDeviceType

`func (o *FilterBy2) GetDeviceType() []string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *FilterBy2) GetDeviceTypeOk() (*[]string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *FilterBy2) SetDeviceType(v []string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *FilterBy2) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *FilterBy2) GetOperatingSystem() []string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *FilterBy2) GetOperatingSystemOk() (*[]string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *FilterBy2) SetOperatingSystem(v []string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *FilterBy2) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetBrowser

`func (o *FilterBy2) GetBrowser() []string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *FilterBy2) GetBrowserOk() (*[]string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *FilterBy2) SetBrowser(v []string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *FilterBy2) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetTag

`func (o *FilterBy2) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *FilterBy2) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *FilterBy2) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *FilterBy2) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


