# PlexSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [readonly] 
**MachineId** | **string** |  | [readonly] 
**Ip** | **string** |  | 
**Port** | **float32** |  | 
**UseSsl** | Pointer to **NullableBool** |  | [optional] 
**Libraries** | Pointer to [**[]PlexLibrary**](PlexLibrary.md) |  | [optional] [readonly] 
**WebAppUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPlexSettings

`func NewPlexSettings(name string, machineId string, ip string, port float32, ) *PlexSettings`

NewPlexSettings instantiates a new PlexSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlexSettingsWithDefaults

`func NewPlexSettingsWithDefaults() *PlexSettings`

NewPlexSettingsWithDefaults instantiates a new PlexSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlexSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlexSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlexSettings) SetName(v string)`

SetName sets Name field to given value.


### GetMachineId

`func (o *PlexSettings) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *PlexSettings) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *PlexSettings) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.


### GetIp

`func (o *PlexSettings) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PlexSettings) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PlexSettings) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPort

`func (o *PlexSettings) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PlexSettings) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PlexSettings) SetPort(v float32)`

SetPort sets Port field to given value.


### GetUseSsl

`func (o *PlexSettings) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *PlexSettings) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *PlexSettings) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *PlexSettings) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### SetUseSslNil

`func (o *PlexSettings) SetUseSslNil(b bool)`

 SetUseSslNil sets the value for UseSsl to be an explicit nil

### UnsetUseSsl
`func (o *PlexSettings) UnsetUseSsl()`

UnsetUseSsl ensures that no value is present for UseSsl, not even an explicit nil
### GetLibraries

`func (o *PlexSettings) GetLibraries() []PlexLibrary`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *PlexSettings) GetLibrariesOk() (*[]PlexLibrary, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *PlexSettings) SetLibraries(v []PlexLibrary)`

SetLibraries sets Libraries field to given value.

### HasLibraries

`func (o *PlexSettings) HasLibraries() bool`

HasLibraries returns a boolean if a field has been set.

### GetWebAppUrl

`func (o *PlexSettings) GetWebAppUrl() string`

GetWebAppUrl returns the WebAppUrl field if non-nil, zero value otherwise.

### GetWebAppUrlOk

`func (o *PlexSettings) GetWebAppUrlOk() (*string, bool)`

GetWebAppUrlOk returns a tuple with the WebAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAppUrl

`func (o *PlexSettings) SetWebAppUrl(v string)`

SetWebAppUrl sets WebAppUrl field to given value.

### HasWebAppUrl

`func (o *PlexSettings) HasWebAppUrl() bool`

HasWebAppUrl returns a boolean if a field has been set.

### SetWebAppUrlNil

`func (o *PlexSettings) SetWebAppUrlNil(b bool)`

 SetWebAppUrlNil sets the value for WebAppUrl to be an explicit nil

### UnsetWebAppUrl
`func (o *PlexSettings) UnsetWebAppUrl()`

UnsetWebAppUrl ensures that no value is present for WebAppUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


