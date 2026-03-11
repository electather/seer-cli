# SettingsCacheGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageCache** | Pointer to [**SettingsCacheGet200ResponseImageCache**](SettingsCacheGet200ResponseImageCache.md) |  | [optional] 
**DnsCache** | Pointer to [**SettingsCacheGet200ResponseDnsCache**](SettingsCacheGet200ResponseDnsCache.md) |  | [optional] 
**ApiCaches** | Pointer to [**[]SettingsCacheGet200ResponseApiCachesInner**](SettingsCacheGet200ResponseApiCachesInner.md) |  | [optional] 

## Methods

### NewSettingsCacheGet200Response

`func NewSettingsCacheGet200Response() *SettingsCacheGet200Response`

NewSettingsCacheGet200Response instantiates a new SettingsCacheGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsCacheGet200ResponseWithDefaults

`func NewSettingsCacheGet200ResponseWithDefaults() *SettingsCacheGet200Response`

NewSettingsCacheGet200ResponseWithDefaults instantiates a new SettingsCacheGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageCache

`func (o *SettingsCacheGet200Response) GetImageCache() SettingsCacheGet200ResponseImageCache`

GetImageCache returns the ImageCache field if non-nil, zero value otherwise.

### GetImageCacheOk

`func (o *SettingsCacheGet200Response) GetImageCacheOk() (*SettingsCacheGet200ResponseImageCache, bool)`

GetImageCacheOk returns a tuple with the ImageCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCache

`func (o *SettingsCacheGet200Response) SetImageCache(v SettingsCacheGet200ResponseImageCache)`

SetImageCache sets ImageCache field to given value.

### HasImageCache

`func (o *SettingsCacheGet200Response) HasImageCache() bool`

HasImageCache returns a boolean if a field has been set.

### GetDnsCache

`func (o *SettingsCacheGet200Response) GetDnsCache() SettingsCacheGet200ResponseDnsCache`

GetDnsCache returns the DnsCache field if non-nil, zero value otherwise.

### GetDnsCacheOk

`func (o *SettingsCacheGet200Response) GetDnsCacheOk() (*SettingsCacheGet200ResponseDnsCache, bool)`

GetDnsCacheOk returns a tuple with the DnsCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCache

`func (o *SettingsCacheGet200Response) SetDnsCache(v SettingsCacheGet200ResponseDnsCache)`

SetDnsCache sets DnsCache field to given value.

### HasDnsCache

`func (o *SettingsCacheGet200Response) HasDnsCache() bool`

HasDnsCache returns a boolean if a field has been set.

### GetApiCaches

`func (o *SettingsCacheGet200Response) GetApiCaches() []SettingsCacheGet200ResponseApiCachesInner`

GetApiCaches returns the ApiCaches field if non-nil, zero value otherwise.

### GetApiCachesOk

`func (o *SettingsCacheGet200Response) GetApiCachesOk() (*[]SettingsCacheGet200ResponseApiCachesInner, bool)`

GetApiCachesOk returns a tuple with the ApiCaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCaches

`func (o *SettingsCacheGet200Response) SetApiCaches(v []SettingsCacheGet200ResponseApiCachesInner)`

SetApiCaches sets ApiCaches field to given value.

### HasApiCaches

`func (o *SettingsCacheGet200Response) HasApiCaches() bool`

HasApiCaches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


