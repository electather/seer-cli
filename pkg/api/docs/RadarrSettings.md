# RadarrSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Hostname** | **string** |  | 
**Port** | **float32** |  | 
**ApiKey** | **string** |  | 
**UseSsl** | **bool** |  | 
**BaseUrl** | Pointer to **string** |  | [optional] 
**ActiveProfileId** | **float32** |  | 
**ActiveProfileName** | **string** |  | 
**ActiveDirectory** | **string** |  | 
**Is4k** | **bool** |  | 
**MinimumAvailability** | **string** |  | 
**IsDefault** | **bool** |  | 
**ExternalUrl** | Pointer to **string** |  | [optional] 
**SyncEnabled** | Pointer to **bool** |  | [optional] 
**PreventSearch** | Pointer to **bool** |  | [optional] 

## Methods

### NewRadarrSettings

`func NewRadarrSettings(name string, hostname string, port float32, apiKey string, useSsl bool, activeProfileId float32, activeProfileName string, activeDirectory string, is4k bool, minimumAvailability string, isDefault bool, ) *RadarrSettings`

NewRadarrSettings instantiates a new RadarrSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadarrSettingsWithDefaults

`func NewRadarrSettingsWithDefaults() *RadarrSettings`

NewRadarrSettingsWithDefaults instantiates a new RadarrSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RadarrSettings) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RadarrSettings) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RadarrSettings) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *RadarrSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RadarrSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RadarrSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RadarrSettings) SetName(v string)`

SetName sets Name field to given value.


### GetHostname

`func (o *RadarrSettings) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RadarrSettings) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RadarrSettings) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *RadarrSettings) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RadarrSettings) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RadarrSettings) SetPort(v float32)`

SetPort sets Port field to given value.


### GetApiKey

`func (o *RadarrSettings) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *RadarrSettings) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *RadarrSettings) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetUseSsl

`func (o *RadarrSettings) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *RadarrSettings) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *RadarrSettings) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.


### GetBaseUrl

`func (o *RadarrSettings) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *RadarrSettings) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *RadarrSettings) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *RadarrSettings) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetActiveProfileId

`func (o *RadarrSettings) GetActiveProfileId() float32`

GetActiveProfileId returns the ActiveProfileId field if non-nil, zero value otherwise.

### GetActiveProfileIdOk

`func (o *RadarrSettings) GetActiveProfileIdOk() (*float32, bool)`

GetActiveProfileIdOk returns a tuple with the ActiveProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfileId

`func (o *RadarrSettings) SetActiveProfileId(v float32)`

SetActiveProfileId sets ActiveProfileId field to given value.


### GetActiveProfileName

`func (o *RadarrSettings) GetActiveProfileName() string`

GetActiveProfileName returns the ActiveProfileName field if non-nil, zero value otherwise.

### GetActiveProfileNameOk

`func (o *RadarrSettings) GetActiveProfileNameOk() (*string, bool)`

GetActiveProfileNameOk returns a tuple with the ActiveProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfileName

`func (o *RadarrSettings) SetActiveProfileName(v string)`

SetActiveProfileName sets ActiveProfileName field to given value.


### GetActiveDirectory

`func (o *RadarrSettings) GetActiveDirectory() string`

GetActiveDirectory returns the ActiveDirectory field if non-nil, zero value otherwise.

### GetActiveDirectoryOk

`func (o *RadarrSettings) GetActiveDirectoryOk() (*string, bool)`

GetActiveDirectoryOk returns a tuple with the ActiveDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectory

`func (o *RadarrSettings) SetActiveDirectory(v string)`

SetActiveDirectory sets ActiveDirectory field to given value.


### GetIs4k

`func (o *RadarrSettings) GetIs4k() bool`

GetIs4k returns the Is4k field if non-nil, zero value otherwise.

### GetIs4kOk

`func (o *RadarrSettings) GetIs4kOk() (*bool, bool)`

GetIs4kOk returns a tuple with the Is4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs4k

`func (o *RadarrSettings) SetIs4k(v bool)`

SetIs4k sets Is4k field to given value.


### GetMinimumAvailability

`func (o *RadarrSettings) GetMinimumAvailability() string`

GetMinimumAvailability returns the MinimumAvailability field if non-nil, zero value otherwise.

### GetMinimumAvailabilityOk

`func (o *RadarrSettings) GetMinimumAvailabilityOk() (*string, bool)`

GetMinimumAvailabilityOk returns a tuple with the MinimumAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAvailability

`func (o *RadarrSettings) SetMinimumAvailability(v string)`

SetMinimumAvailability sets MinimumAvailability field to given value.


### GetIsDefault

`func (o *RadarrSettings) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *RadarrSettings) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *RadarrSettings) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetExternalUrl

`func (o *RadarrSettings) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *RadarrSettings) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *RadarrSettings) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *RadarrSettings) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetSyncEnabled

`func (o *RadarrSettings) GetSyncEnabled() bool`

GetSyncEnabled returns the SyncEnabled field if non-nil, zero value otherwise.

### GetSyncEnabledOk

`func (o *RadarrSettings) GetSyncEnabledOk() (*bool, bool)`

GetSyncEnabledOk returns a tuple with the SyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncEnabled

`func (o *RadarrSettings) SetSyncEnabled(v bool)`

SetSyncEnabled sets SyncEnabled field to given value.

### HasSyncEnabled

`func (o *RadarrSettings) HasSyncEnabled() bool`

HasSyncEnabled returns a boolean if a field has been set.

### GetPreventSearch

`func (o *RadarrSettings) GetPreventSearch() bool`

GetPreventSearch returns the PreventSearch field if non-nil, zero value otherwise.

### GetPreventSearchOk

`func (o *RadarrSettings) GetPreventSearchOk() (*bool, bool)`

GetPreventSearchOk returns a tuple with the PreventSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventSearch

`func (o *RadarrSettings) SetPreventSearch(v bool)`

SetPreventSearch sets PreventSearch field to given value.

### HasPreventSearch

`func (o *RadarrSettings) HasPreventSearch() bool`

HasPreventSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


