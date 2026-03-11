# SonarrSettings

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
**ActiveLanguageProfileId** | Pointer to **float32** |  | [optional] 
**ActiveAnimeProfileId** | Pointer to **NullableFloat32** |  | [optional] 
**ActiveAnimeLanguageProfileId** | Pointer to **NullableFloat32** |  | [optional] 
**ActiveAnimeProfileName** | Pointer to **NullableString** |  | [optional] 
**ActiveAnimeDirectory** | Pointer to **NullableString** |  | [optional] 
**Is4k** | **bool** |  | 
**EnableSeasonFolders** | **bool** |  | 
**IsDefault** | **bool** |  | 
**ExternalUrl** | Pointer to **string** |  | [optional] 
**SyncEnabled** | Pointer to **bool** |  | [optional] 
**PreventSearch** | Pointer to **bool** |  | [optional] 

## Methods

### NewSonarrSettings

`func NewSonarrSettings(name string, hostname string, port float32, apiKey string, useSsl bool, activeProfileId float32, activeProfileName string, activeDirectory string, is4k bool, enableSeasonFolders bool, isDefault bool, ) *SonarrSettings`

NewSonarrSettings instantiates a new SonarrSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSonarrSettingsWithDefaults

`func NewSonarrSettingsWithDefaults() *SonarrSettings`

NewSonarrSettingsWithDefaults instantiates a new SonarrSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SonarrSettings) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SonarrSettings) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SonarrSettings) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *SonarrSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SonarrSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SonarrSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SonarrSettings) SetName(v string)`

SetName sets Name field to given value.


### GetHostname

`func (o *SonarrSettings) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SonarrSettings) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SonarrSettings) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *SonarrSettings) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SonarrSettings) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SonarrSettings) SetPort(v float32)`

SetPort sets Port field to given value.


### GetApiKey

`func (o *SonarrSettings) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *SonarrSettings) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *SonarrSettings) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetUseSsl

`func (o *SonarrSettings) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *SonarrSettings) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *SonarrSettings) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.


### GetBaseUrl

`func (o *SonarrSettings) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *SonarrSettings) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *SonarrSettings) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *SonarrSettings) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetActiveProfileId

`func (o *SonarrSettings) GetActiveProfileId() float32`

GetActiveProfileId returns the ActiveProfileId field if non-nil, zero value otherwise.

### GetActiveProfileIdOk

`func (o *SonarrSettings) GetActiveProfileIdOk() (*float32, bool)`

GetActiveProfileIdOk returns a tuple with the ActiveProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfileId

`func (o *SonarrSettings) SetActiveProfileId(v float32)`

SetActiveProfileId sets ActiveProfileId field to given value.


### GetActiveProfileName

`func (o *SonarrSettings) GetActiveProfileName() string`

GetActiveProfileName returns the ActiveProfileName field if non-nil, zero value otherwise.

### GetActiveProfileNameOk

`func (o *SonarrSettings) GetActiveProfileNameOk() (*string, bool)`

GetActiveProfileNameOk returns a tuple with the ActiveProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProfileName

`func (o *SonarrSettings) SetActiveProfileName(v string)`

SetActiveProfileName sets ActiveProfileName field to given value.


### GetActiveDirectory

`func (o *SonarrSettings) GetActiveDirectory() string`

GetActiveDirectory returns the ActiveDirectory field if non-nil, zero value otherwise.

### GetActiveDirectoryOk

`func (o *SonarrSettings) GetActiveDirectoryOk() (*string, bool)`

GetActiveDirectoryOk returns a tuple with the ActiveDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectory

`func (o *SonarrSettings) SetActiveDirectory(v string)`

SetActiveDirectory sets ActiveDirectory field to given value.


### GetActiveLanguageProfileId

`func (o *SonarrSettings) GetActiveLanguageProfileId() float32`

GetActiveLanguageProfileId returns the ActiveLanguageProfileId field if non-nil, zero value otherwise.

### GetActiveLanguageProfileIdOk

`func (o *SonarrSettings) GetActiveLanguageProfileIdOk() (*float32, bool)`

GetActiveLanguageProfileIdOk returns a tuple with the ActiveLanguageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveLanguageProfileId

`func (o *SonarrSettings) SetActiveLanguageProfileId(v float32)`

SetActiveLanguageProfileId sets ActiveLanguageProfileId field to given value.

### HasActiveLanguageProfileId

`func (o *SonarrSettings) HasActiveLanguageProfileId() bool`

HasActiveLanguageProfileId returns a boolean if a field has been set.

### GetActiveAnimeProfileId

`func (o *SonarrSettings) GetActiveAnimeProfileId() float32`

GetActiveAnimeProfileId returns the ActiveAnimeProfileId field if non-nil, zero value otherwise.

### GetActiveAnimeProfileIdOk

`func (o *SonarrSettings) GetActiveAnimeProfileIdOk() (*float32, bool)`

GetActiveAnimeProfileIdOk returns a tuple with the ActiveAnimeProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAnimeProfileId

`func (o *SonarrSettings) SetActiveAnimeProfileId(v float32)`

SetActiveAnimeProfileId sets ActiveAnimeProfileId field to given value.

### HasActiveAnimeProfileId

`func (o *SonarrSettings) HasActiveAnimeProfileId() bool`

HasActiveAnimeProfileId returns a boolean if a field has been set.

### SetActiveAnimeProfileIdNil

`func (o *SonarrSettings) SetActiveAnimeProfileIdNil(b bool)`

 SetActiveAnimeProfileIdNil sets the value for ActiveAnimeProfileId to be an explicit nil

### UnsetActiveAnimeProfileId
`func (o *SonarrSettings) UnsetActiveAnimeProfileId()`

UnsetActiveAnimeProfileId ensures that no value is present for ActiveAnimeProfileId, not even an explicit nil
### GetActiveAnimeLanguageProfileId

`func (o *SonarrSettings) GetActiveAnimeLanguageProfileId() float32`

GetActiveAnimeLanguageProfileId returns the ActiveAnimeLanguageProfileId field if non-nil, zero value otherwise.

### GetActiveAnimeLanguageProfileIdOk

`func (o *SonarrSettings) GetActiveAnimeLanguageProfileIdOk() (*float32, bool)`

GetActiveAnimeLanguageProfileIdOk returns a tuple with the ActiveAnimeLanguageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAnimeLanguageProfileId

`func (o *SonarrSettings) SetActiveAnimeLanguageProfileId(v float32)`

SetActiveAnimeLanguageProfileId sets ActiveAnimeLanguageProfileId field to given value.

### HasActiveAnimeLanguageProfileId

`func (o *SonarrSettings) HasActiveAnimeLanguageProfileId() bool`

HasActiveAnimeLanguageProfileId returns a boolean if a field has been set.

### SetActiveAnimeLanguageProfileIdNil

`func (o *SonarrSettings) SetActiveAnimeLanguageProfileIdNil(b bool)`

 SetActiveAnimeLanguageProfileIdNil sets the value for ActiveAnimeLanguageProfileId to be an explicit nil

### UnsetActiveAnimeLanguageProfileId
`func (o *SonarrSettings) UnsetActiveAnimeLanguageProfileId()`

UnsetActiveAnimeLanguageProfileId ensures that no value is present for ActiveAnimeLanguageProfileId, not even an explicit nil
### GetActiveAnimeProfileName

`func (o *SonarrSettings) GetActiveAnimeProfileName() string`

GetActiveAnimeProfileName returns the ActiveAnimeProfileName field if non-nil, zero value otherwise.

### GetActiveAnimeProfileNameOk

`func (o *SonarrSettings) GetActiveAnimeProfileNameOk() (*string, bool)`

GetActiveAnimeProfileNameOk returns a tuple with the ActiveAnimeProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAnimeProfileName

`func (o *SonarrSettings) SetActiveAnimeProfileName(v string)`

SetActiveAnimeProfileName sets ActiveAnimeProfileName field to given value.

### HasActiveAnimeProfileName

`func (o *SonarrSettings) HasActiveAnimeProfileName() bool`

HasActiveAnimeProfileName returns a boolean if a field has been set.

### SetActiveAnimeProfileNameNil

`func (o *SonarrSettings) SetActiveAnimeProfileNameNil(b bool)`

 SetActiveAnimeProfileNameNil sets the value for ActiveAnimeProfileName to be an explicit nil

### UnsetActiveAnimeProfileName
`func (o *SonarrSettings) UnsetActiveAnimeProfileName()`

UnsetActiveAnimeProfileName ensures that no value is present for ActiveAnimeProfileName, not even an explicit nil
### GetActiveAnimeDirectory

`func (o *SonarrSettings) GetActiveAnimeDirectory() string`

GetActiveAnimeDirectory returns the ActiveAnimeDirectory field if non-nil, zero value otherwise.

### GetActiveAnimeDirectoryOk

`func (o *SonarrSettings) GetActiveAnimeDirectoryOk() (*string, bool)`

GetActiveAnimeDirectoryOk returns a tuple with the ActiveAnimeDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAnimeDirectory

`func (o *SonarrSettings) SetActiveAnimeDirectory(v string)`

SetActiveAnimeDirectory sets ActiveAnimeDirectory field to given value.

### HasActiveAnimeDirectory

`func (o *SonarrSettings) HasActiveAnimeDirectory() bool`

HasActiveAnimeDirectory returns a boolean if a field has been set.

### SetActiveAnimeDirectoryNil

`func (o *SonarrSettings) SetActiveAnimeDirectoryNil(b bool)`

 SetActiveAnimeDirectoryNil sets the value for ActiveAnimeDirectory to be an explicit nil

### UnsetActiveAnimeDirectory
`func (o *SonarrSettings) UnsetActiveAnimeDirectory()`

UnsetActiveAnimeDirectory ensures that no value is present for ActiveAnimeDirectory, not even an explicit nil
### GetIs4k

`func (o *SonarrSettings) GetIs4k() bool`

GetIs4k returns the Is4k field if non-nil, zero value otherwise.

### GetIs4kOk

`func (o *SonarrSettings) GetIs4kOk() (*bool, bool)`

GetIs4kOk returns a tuple with the Is4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs4k

`func (o *SonarrSettings) SetIs4k(v bool)`

SetIs4k sets Is4k field to given value.


### GetEnableSeasonFolders

`func (o *SonarrSettings) GetEnableSeasonFolders() bool`

GetEnableSeasonFolders returns the EnableSeasonFolders field if non-nil, zero value otherwise.

### GetEnableSeasonFoldersOk

`func (o *SonarrSettings) GetEnableSeasonFoldersOk() (*bool, bool)`

GetEnableSeasonFoldersOk returns a tuple with the EnableSeasonFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSeasonFolders

`func (o *SonarrSettings) SetEnableSeasonFolders(v bool)`

SetEnableSeasonFolders sets EnableSeasonFolders field to given value.


### GetIsDefault

`func (o *SonarrSettings) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SonarrSettings) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SonarrSettings) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetExternalUrl

`func (o *SonarrSettings) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *SonarrSettings) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *SonarrSettings) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *SonarrSettings) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetSyncEnabled

`func (o *SonarrSettings) GetSyncEnabled() bool`

GetSyncEnabled returns the SyncEnabled field if non-nil, zero value otherwise.

### GetSyncEnabledOk

`func (o *SonarrSettings) GetSyncEnabledOk() (*bool, bool)`

GetSyncEnabledOk returns a tuple with the SyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncEnabled

`func (o *SonarrSettings) SetSyncEnabled(v bool)`

SetSyncEnabled sets SyncEnabled field to given value.

### HasSyncEnabled

`func (o *SonarrSettings) HasSyncEnabled() bool`

HasSyncEnabled returns a boolean if a field has been set.

### GetPreventSearch

`func (o *SonarrSettings) GetPreventSearch() bool`

GetPreventSearch returns the PreventSearch field if non-nil, zero value otherwise.

### GetPreventSearchOk

`func (o *SonarrSettings) GetPreventSearchOk() (*bool, bool)`

GetPreventSearchOk returns a tuple with the PreventSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventSearch

`func (o *SonarrSettings) SetPreventSearch(v bool)`

SetPreventSearch sets PreventSearch field to given value.

### HasPreventSearch

`func (o *SonarrSettings) HasPreventSearch() bool`

HasPreventSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


