# SettingsPlexSyncGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Running** | Pointer to **bool** |  | [optional] 
**Progress** | Pointer to **float32** |  | [optional] 
**Total** | Pointer to **float32** |  | [optional] 
**CurrentLibrary** | Pointer to [**PlexLibrary**](PlexLibrary.md) |  | [optional] 
**Libraries** | Pointer to [**[]PlexLibrary**](PlexLibrary.md) |  | [optional] 

## Methods

### NewSettingsPlexSyncGet200Response

`func NewSettingsPlexSyncGet200Response() *SettingsPlexSyncGet200Response`

NewSettingsPlexSyncGet200Response instantiates a new SettingsPlexSyncGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsPlexSyncGet200ResponseWithDefaults

`func NewSettingsPlexSyncGet200ResponseWithDefaults() *SettingsPlexSyncGet200Response`

NewSettingsPlexSyncGet200ResponseWithDefaults instantiates a new SettingsPlexSyncGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunning

`func (o *SettingsPlexSyncGet200Response) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *SettingsPlexSyncGet200Response) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *SettingsPlexSyncGet200Response) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *SettingsPlexSyncGet200Response) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetProgress

`func (o *SettingsPlexSyncGet200Response) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SettingsPlexSyncGet200Response) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SettingsPlexSyncGet200Response) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SettingsPlexSyncGet200Response) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetTotal

`func (o *SettingsPlexSyncGet200Response) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SettingsPlexSyncGet200Response) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SettingsPlexSyncGet200Response) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SettingsPlexSyncGet200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCurrentLibrary

`func (o *SettingsPlexSyncGet200Response) GetCurrentLibrary() PlexLibrary`

GetCurrentLibrary returns the CurrentLibrary field if non-nil, zero value otherwise.

### GetCurrentLibraryOk

`func (o *SettingsPlexSyncGet200Response) GetCurrentLibraryOk() (*PlexLibrary, bool)`

GetCurrentLibraryOk returns a tuple with the CurrentLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLibrary

`func (o *SettingsPlexSyncGet200Response) SetCurrentLibrary(v PlexLibrary)`

SetCurrentLibrary sets CurrentLibrary field to given value.

### HasCurrentLibrary

`func (o *SettingsPlexSyncGet200Response) HasCurrentLibrary() bool`

HasCurrentLibrary returns a boolean if a field has been set.

### GetLibraries

`func (o *SettingsPlexSyncGet200Response) GetLibraries() []PlexLibrary`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *SettingsPlexSyncGet200Response) GetLibrariesOk() (*[]PlexLibrary, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *SettingsPlexSyncGet200Response) SetLibraries(v []PlexLibrary)`

SetLibraries sets Libraries field to given value.

### HasLibraries

`func (o *SettingsPlexSyncGet200Response) HasLibraries() bool`

HasLibraries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


