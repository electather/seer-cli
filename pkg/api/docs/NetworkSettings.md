# NetworkSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfProtection** | Pointer to **bool** |  | [optional] 
**ForceIpv4First** | Pointer to **bool** |  | [optional] 
**TrustProxy** | Pointer to **bool** |  | [optional] 
**Proxy** | Pointer to [**NetworkSettingsProxy**](NetworkSettingsProxy.md) |  | [optional] 
**DnsCache** | Pointer to [**NetworkSettingsDnsCache**](NetworkSettingsDnsCache.md) |  | [optional] 

## Methods

### NewNetworkSettings

`func NewNetworkSettings() *NetworkSettings`

NewNetworkSettings instantiates a new NetworkSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSettingsWithDefaults

`func NewNetworkSettingsWithDefaults() *NetworkSettings`

NewNetworkSettingsWithDefaults instantiates a new NetworkSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfProtection

`func (o *NetworkSettings) GetCsrfProtection() bool`

GetCsrfProtection returns the CsrfProtection field if non-nil, zero value otherwise.

### GetCsrfProtectionOk

`func (o *NetworkSettings) GetCsrfProtectionOk() (*bool, bool)`

GetCsrfProtectionOk returns a tuple with the CsrfProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfProtection

`func (o *NetworkSettings) SetCsrfProtection(v bool)`

SetCsrfProtection sets CsrfProtection field to given value.

### HasCsrfProtection

`func (o *NetworkSettings) HasCsrfProtection() bool`

HasCsrfProtection returns a boolean if a field has been set.

### GetForceIpv4First

`func (o *NetworkSettings) GetForceIpv4First() bool`

GetForceIpv4First returns the ForceIpv4First field if non-nil, zero value otherwise.

### GetForceIpv4FirstOk

`func (o *NetworkSettings) GetForceIpv4FirstOk() (*bool, bool)`

GetForceIpv4FirstOk returns a tuple with the ForceIpv4First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceIpv4First

`func (o *NetworkSettings) SetForceIpv4First(v bool)`

SetForceIpv4First sets ForceIpv4First field to given value.

### HasForceIpv4First

`func (o *NetworkSettings) HasForceIpv4First() bool`

HasForceIpv4First returns a boolean if a field has been set.

### GetTrustProxy

`func (o *NetworkSettings) GetTrustProxy() bool`

GetTrustProxy returns the TrustProxy field if non-nil, zero value otherwise.

### GetTrustProxyOk

`func (o *NetworkSettings) GetTrustProxyOk() (*bool, bool)`

GetTrustProxyOk returns a tuple with the TrustProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustProxy

`func (o *NetworkSettings) SetTrustProxy(v bool)`

SetTrustProxy sets TrustProxy field to given value.

### HasTrustProxy

`func (o *NetworkSettings) HasTrustProxy() bool`

HasTrustProxy returns a boolean if a field has been set.

### GetProxy

`func (o *NetworkSettings) GetProxy() NetworkSettingsProxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *NetworkSettings) GetProxyOk() (*NetworkSettingsProxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *NetworkSettings) SetProxy(v NetworkSettingsProxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *NetworkSettings) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetDnsCache

`func (o *NetworkSettings) GetDnsCache() NetworkSettingsDnsCache`

GetDnsCache returns the DnsCache field if non-nil, zero value otherwise.

### GetDnsCacheOk

`func (o *NetworkSettings) GetDnsCacheOk() (*NetworkSettingsDnsCache, bool)`

GetDnsCacheOk returns a tuple with the DnsCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCache

`func (o *NetworkSettings) SetDnsCache(v NetworkSettingsDnsCache)`

SetDnsCache sets DnsCache field to given value.

### HasDnsCache

`func (o *NetworkSettings) HasDnsCache() bool`

HasDnsCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


