# JellyfinSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Hostname** | Pointer to **string** |  | [optional] 
**ExternalHostname** | Pointer to **string** |  | [optional] 
**JellyfinForgotPasswordUrl** | Pointer to **string** |  | [optional] 
**AdminUser** | Pointer to **string** |  | [optional] 
**AdminPass** | Pointer to **string** |  | [optional] 
**Libraries** | Pointer to [**[]JellyfinLibrary**](JellyfinLibrary.md) |  | [optional] [readonly] 
**ServerID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewJellyfinSettings

`func NewJellyfinSettings() *JellyfinSettings`

NewJellyfinSettings instantiates a new JellyfinSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJellyfinSettingsWithDefaults

`func NewJellyfinSettingsWithDefaults() *JellyfinSettings`

NewJellyfinSettingsWithDefaults instantiates a new JellyfinSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JellyfinSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JellyfinSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JellyfinSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JellyfinSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostname

`func (o *JellyfinSettings) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *JellyfinSettings) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *JellyfinSettings) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *JellyfinSettings) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetExternalHostname

`func (o *JellyfinSettings) GetExternalHostname() string`

GetExternalHostname returns the ExternalHostname field if non-nil, zero value otherwise.

### GetExternalHostnameOk

`func (o *JellyfinSettings) GetExternalHostnameOk() (*string, bool)`

GetExternalHostnameOk returns a tuple with the ExternalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHostname

`func (o *JellyfinSettings) SetExternalHostname(v string)`

SetExternalHostname sets ExternalHostname field to given value.

### HasExternalHostname

`func (o *JellyfinSettings) HasExternalHostname() bool`

HasExternalHostname returns a boolean if a field has been set.

### GetJellyfinForgotPasswordUrl

`func (o *JellyfinSettings) GetJellyfinForgotPasswordUrl() string`

GetJellyfinForgotPasswordUrl returns the JellyfinForgotPasswordUrl field if non-nil, zero value otherwise.

### GetJellyfinForgotPasswordUrlOk

`func (o *JellyfinSettings) GetJellyfinForgotPasswordUrlOk() (*string, bool)`

GetJellyfinForgotPasswordUrlOk returns a tuple with the JellyfinForgotPasswordUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJellyfinForgotPasswordUrl

`func (o *JellyfinSettings) SetJellyfinForgotPasswordUrl(v string)`

SetJellyfinForgotPasswordUrl sets JellyfinForgotPasswordUrl field to given value.

### HasJellyfinForgotPasswordUrl

`func (o *JellyfinSettings) HasJellyfinForgotPasswordUrl() bool`

HasJellyfinForgotPasswordUrl returns a boolean if a field has been set.

### GetAdminUser

`func (o *JellyfinSettings) GetAdminUser() string`

GetAdminUser returns the AdminUser field if non-nil, zero value otherwise.

### GetAdminUserOk

`func (o *JellyfinSettings) GetAdminUserOk() (*string, bool)`

GetAdminUserOk returns a tuple with the AdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUser

`func (o *JellyfinSettings) SetAdminUser(v string)`

SetAdminUser sets AdminUser field to given value.

### HasAdminUser

`func (o *JellyfinSettings) HasAdminUser() bool`

HasAdminUser returns a boolean if a field has been set.

### GetAdminPass

`func (o *JellyfinSettings) GetAdminPass() string`

GetAdminPass returns the AdminPass field if non-nil, zero value otherwise.

### GetAdminPassOk

`func (o *JellyfinSettings) GetAdminPassOk() (*string, bool)`

GetAdminPassOk returns a tuple with the AdminPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPass

`func (o *JellyfinSettings) SetAdminPass(v string)`

SetAdminPass sets AdminPass field to given value.

### HasAdminPass

`func (o *JellyfinSettings) HasAdminPass() bool`

HasAdminPass returns a boolean if a field has been set.

### GetLibraries

`func (o *JellyfinSettings) GetLibraries() []JellyfinLibrary`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *JellyfinSettings) GetLibrariesOk() (*[]JellyfinLibrary, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *JellyfinSettings) SetLibraries(v []JellyfinLibrary)`

SetLibraries sets Libraries field to given value.

### HasLibraries

`func (o *JellyfinSettings) HasLibraries() bool`

HasLibraries returns a boolean if a field has been set.

### GetServerID

`func (o *JellyfinSettings) GetServerID() string`

GetServerID returns the ServerID field if non-nil, zero value otherwise.

### GetServerIDOk

`func (o *JellyfinSettings) GetServerIDOk() (*string, bool)`

GetServerIDOk returns a tuple with the ServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerID

`func (o *JellyfinSettings) SetServerID(v string)`

SetServerID sets ServerID field to given value.

### HasServerID

`func (o *JellyfinSettings) HasServerID() bool`

HasServerID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


