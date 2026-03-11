# MainSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] [readonly] 
**AppLanguage** | Pointer to **string** |  | [optional] 
**ApplicationTitle** | Pointer to **string** |  | [optional] 
**ApplicationUrl** | Pointer to **string** |  | [optional] 
**HideAvailable** | Pointer to **bool** |  | [optional] 
**PartialRequestsEnabled** | Pointer to **bool** |  | [optional] 
**LocalLogin** | Pointer to **bool** |  | [optional] 
**MediaServerType** | Pointer to **float32** |  | [optional] 
**NewPlexLogin** | Pointer to **bool** |  | [optional] 
**DefaultPermissions** | Pointer to **float32** |  | [optional] 
**EnableSpecialEpisodes** | Pointer to **bool** |  | [optional] 

## Methods

### NewMainSettings

`func NewMainSettings() *MainSettings`

NewMainSettings instantiates a new MainSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainSettingsWithDefaults

`func NewMainSettingsWithDefaults() *MainSettings`

NewMainSettingsWithDefaults instantiates a new MainSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *MainSettings) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *MainSettings) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *MainSettings) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *MainSettings) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAppLanguage

`func (o *MainSettings) GetAppLanguage() string`

GetAppLanguage returns the AppLanguage field if non-nil, zero value otherwise.

### GetAppLanguageOk

`func (o *MainSettings) GetAppLanguageOk() (*string, bool)`

GetAppLanguageOk returns a tuple with the AppLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLanguage

`func (o *MainSettings) SetAppLanguage(v string)`

SetAppLanguage sets AppLanguage field to given value.

### HasAppLanguage

`func (o *MainSettings) HasAppLanguage() bool`

HasAppLanguage returns a boolean if a field has been set.

### GetApplicationTitle

`func (o *MainSettings) GetApplicationTitle() string`

GetApplicationTitle returns the ApplicationTitle field if non-nil, zero value otherwise.

### GetApplicationTitleOk

`func (o *MainSettings) GetApplicationTitleOk() (*string, bool)`

GetApplicationTitleOk returns a tuple with the ApplicationTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTitle

`func (o *MainSettings) SetApplicationTitle(v string)`

SetApplicationTitle sets ApplicationTitle field to given value.

### HasApplicationTitle

`func (o *MainSettings) HasApplicationTitle() bool`

HasApplicationTitle returns a boolean if a field has been set.

### GetApplicationUrl

`func (o *MainSettings) GetApplicationUrl() string`

GetApplicationUrl returns the ApplicationUrl field if non-nil, zero value otherwise.

### GetApplicationUrlOk

`func (o *MainSettings) GetApplicationUrlOk() (*string, bool)`

GetApplicationUrlOk returns a tuple with the ApplicationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUrl

`func (o *MainSettings) SetApplicationUrl(v string)`

SetApplicationUrl sets ApplicationUrl field to given value.

### HasApplicationUrl

`func (o *MainSettings) HasApplicationUrl() bool`

HasApplicationUrl returns a boolean if a field has been set.

### GetHideAvailable

`func (o *MainSettings) GetHideAvailable() bool`

GetHideAvailable returns the HideAvailable field if non-nil, zero value otherwise.

### GetHideAvailableOk

`func (o *MainSettings) GetHideAvailableOk() (*bool, bool)`

GetHideAvailableOk returns a tuple with the HideAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAvailable

`func (o *MainSettings) SetHideAvailable(v bool)`

SetHideAvailable sets HideAvailable field to given value.

### HasHideAvailable

`func (o *MainSettings) HasHideAvailable() bool`

HasHideAvailable returns a boolean if a field has been set.

### GetPartialRequestsEnabled

`func (o *MainSettings) GetPartialRequestsEnabled() bool`

GetPartialRequestsEnabled returns the PartialRequestsEnabled field if non-nil, zero value otherwise.

### GetPartialRequestsEnabledOk

`func (o *MainSettings) GetPartialRequestsEnabledOk() (*bool, bool)`

GetPartialRequestsEnabledOk returns a tuple with the PartialRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialRequestsEnabled

`func (o *MainSettings) SetPartialRequestsEnabled(v bool)`

SetPartialRequestsEnabled sets PartialRequestsEnabled field to given value.

### HasPartialRequestsEnabled

`func (o *MainSettings) HasPartialRequestsEnabled() bool`

HasPartialRequestsEnabled returns a boolean if a field has been set.

### GetLocalLogin

`func (o *MainSettings) GetLocalLogin() bool`

GetLocalLogin returns the LocalLogin field if non-nil, zero value otherwise.

### GetLocalLoginOk

`func (o *MainSettings) GetLocalLoginOk() (*bool, bool)`

GetLocalLoginOk returns a tuple with the LocalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalLogin

`func (o *MainSettings) SetLocalLogin(v bool)`

SetLocalLogin sets LocalLogin field to given value.

### HasLocalLogin

`func (o *MainSettings) HasLocalLogin() bool`

HasLocalLogin returns a boolean if a field has been set.

### GetMediaServerType

`func (o *MainSettings) GetMediaServerType() float32`

GetMediaServerType returns the MediaServerType field if non-nil, zero value otherwise.

### GetMediaServerTypeOk

`func (o *MainSettings) GetMediaServerTypeOk() (*float32, bool)`

GetMediaServerTypeOk returns a tuple with the MediaServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaServerType

`func (o *MainSettings) SetMediaServerType(v float32)`

SetMediaServerType sets MediaServerType field to given value.

### HasMediaServerType

`func (o *MainSettings) HasMediaServerType() bool`

HasMediaServerType returns a boolean if a field has been set.

### GetNewPlexLogin

`func (o *MainSettings) GetNewPlexLogin() bool`

GetNewPlexLogin returns the NewPlexLogin field if non-nil, zero value otherwise.

### GetNewPlexLoginOk

`func (o *MainSettings) GetNewPlexLoginOk() (*bool, bool)`

GetNewPlexLoginOk returns a tuple with the NewPlexLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlexLogin

`func (o *MainSettings) SetNewPlexLogin(v bool)`

SetNewPlexLogin sets NewPlexLogin field to given value.

### HasNewPlexLogin

`func (o *MainSettings) HasNewPlexLogin() bool`

HasNewPlexLogin returns a boolean if a field has been set.

### GetDefaultPermissions

`func (o *MainSettings) GetDefaultPermissions() float32`

GetDefaultPermissions returns the DefaultPermissions field if non-nil, zero value otherwise.

### GetDefaultPermissionsOk

`func (o *MainSettings) GetDefaultPermissionsOk() (*float32, bool)`

GetDefaultPermissionsOk returns a tuple with the DefaultPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPermissions

`func (o *MainSettings) SetDefaultPermissions(v float32)`

SetDefaultPermissions sets DefaultPermissions field to given value.

### HasDefaultPermissions

`func (o *MainSettings) HasDefaultPermissions() bool`

HasDefaultPermissions returns a boolean if a field has been set.

### GetEnableSpecialEpisodes

`func (o *MainSettings) GetEnableSpecialEpisodes() bool`

GetEnableSpecialEpisodes returns the EnableSpecialEpisodes field if non-nil, zero value otherwise.

### GetEnableSpecialEpisodesOk

`func (o *MainSettings) GetEnableSpecialEpisodesOk() (*bool, bool)`

GetEnableSpecialEpisodesOk returns a tuple with the EnableSpecialEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpecialEpisodes

`func (o *MainSettings) SetEnableSpecialEpisodes(v bool)`

SetEnableSpecialEpisodes sets EnableSpecialEpisodes field to given value.

### HasEnableSpecialEpisodes

`func (o *MainSettings) HasEnableSpecialEpisodes() bool`

HasEnableSpecialEpisodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


