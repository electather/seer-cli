# UserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DiscordId** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to **NullableString** |  | [optional] 
**DiscoverRegion** | Pointer to **NullableString** |  | [optional] 
**StreamingRegion** | Pointer to **NullableString** |  | [optional] 
**OriginalLanguage** | Pointer to **NullableString** |  | [optional] 
**MovieQuotaLimit** | Pointer to **NullableFloat32** | Maximum number of movie requests allowed | [optional] 
**MovieQuotaDays** | Pointer to **NullableFloat32** | Time period in days for movie quota | [optional] 
**TvQuotaLimit** | Pointer to **NullableFloat32** | Maximum number of TV requests allowed | [optional] 
**TvQuotaDays** | Pointer to **NullableFloat32** | Time period in days for TV quota | [optional] 
**GlobalMovieQuotaDays** | Pointer to **NullableFloat32** | Global movie quota days setting | [optional] 
**GlobalMovieQuotaLimit** | Pointer to **NullableFloat32** | Global movie quota limit setting | [optional] 
**GlobalTvQuotaLimit** | Pointer to **NullableFloat32** | Global TV quota limit setting | [optional] 
**GlobalTvQuotaDays** | Pointer to **NullableFloat32** | Global TV quota days setting | [optional] 
**WatchlistSyncMovies** | Pointer to **NullableBool** | Enable watchlist sync for movies | [optional] 
**WatchlistSyncTv** | Pointer to **NullableBool** | Enable watchlist sync for TV | [optional] 

## Methods

### NewUserSettings

`func NewUserSettings() *UserSettings`

NewUserSettings instantiates a new UserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsWithDefaults

`func NewUserSettingsWithDefaults() *UserSettings`

NewUserSettingsWithDefaults instantiates a new UserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserSettings) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSettings) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSettings) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserSettings) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UserSettings) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserSettings) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetEmail

`func (o *UserSettings) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSettings) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSettings) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserSettings) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDiscordId

`func (o *UserSettings) GetDiscordId() string`

GetDiscordId returns the DiscordId field if non-nil, zero value otherwise.

### GetDiscordIdOk

`func (o *UserSettings) GetDiscordIdOk() (*string, bool)`

GetDiscordIdOk returns a tuple with the DiscordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordId

`func (o *UserSettings) SetDiscordId(v string)`

SetDiscordId sets DiscordId field to given value.

### HasDiscordId

`func (o *UserSettings) HasDiscordId() bool`

HasDiscordId returns a boolean if a field has been set.

### SetDiscordIdNil

`func (o *UserSettings) SetDiscordIdNil(b bool)`

 SetDiscordIdNil sets the value for DiscordId to be an explicit nil

### UnsetDiscordId
`func (o *UserSettings) UnsetDiscordId()`

UnsetDiscordId ensures that no value is present for DiscordId, not even an explicit nil
### GetLocale

`func (o *UserSettings) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UserSettings) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UserSettings) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UserSettings) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *UserSettings) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *UserSettings) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetDiscoverRegion

`func (o *UserSettings) GetDiscoverRegion() string`

GetDiscoverRegion returns the DiscoverRegion field if non-nil, zero value otherwise.

### GetDiscoverRegionOk

`func (o *UserSettings) GetDiscoverRegionOk() (*string, bool)`

GetDiscoverRegionOk returns a tuple with the DiscoverRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverRegion

`func (o *UserSettings) SetDiscoverRegion(v string)`

SetDiscoverRegion sets DiscoverRegion field to given value.

### HasDiscoverRegion

`func (o *UserSettings) HasDiscoverRegion() bool`

HasDiscoverRegion returns a boolean if a field has been set.

### SetDiscoverRegionNil

`func (o *UserSettings) SetDiscoverRegionNil(b bool)`

 SetDiscoverRegionNil sets the value for DiscoverRegion to be an explicit nil

### UnsetDiscoverRegion
`func (o *UserSettings) UnsetDiscoverRegion()`

UnsetDiscoverRegion ensures that no value is present for DiscoverRegion, not even an explicit nil
### GetStreamingRegion

`func (o *UserSettings) GetStreamingRegion() string`

GetStreamingRegion returns the StreamingRegion field if non-nil, zero value otherwise.

### GetStreamingRegionOk

`func (o *UserSettings) GetStreamingRegionOk() (*string, bool)`

GetStreamingRegionOk returns a tuple with the StreamingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingRegion

`func (o *UserSettings) SetStreamingRegion(v string)`

SetStreamingRegion sets StreamingRegion field to given value.

### HasStreamingRegion

`func (o *UserSettings) HasStreamingRegion() bool`

HasStreamingRegion returns a boolean if a field has been set.

### SetStreamingRegionNil

`func (o *UserSettings) SetStreamingRegionNil(b bool)`

 SetStreamingRegionNil sets the value for StreamingRegion to be an explicit nil

### UnsetStreamingRegion
`func (o *UserSettings) UnsetStreamingRegion()`

UnsetStreamingRegion ensures that no value is present for StreamingRegion, not even an explicit nil
### GetOriginalLanguage

`func (o *UserSettings) GetOriginalLanguage() string`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *UserSettings) GetOriginalLanguageOk() (*string, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *UserSettings) SetOriginalLanguage(v string)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *UserSettings) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### SetOriginalLanguageNil

`func (o *UserSettings) SetOriginalLanguageNil(b bool)`

 SetOriginalLanguageNil sets the value for OriginalLanguage to be an explicit nil

### UnsetOriginalLanguage
`func (o *UserSettings) UnsetOriginalLanguage()`

UnsetOriginalLanguage ensures that no value is present for OriginalLanguage, not even an explicit nil
### GetMovieQuotaLimit

`func (o *UserSettings) GetMovieQuotaLimit() float32`

GetMovieQuotaLimit returns the MovieQuotaLimit field if non-nil, zero value otherwise.

### GetMovieQuotaLimitOk

`func (o *UserSettings) GetMovieQuotaLimitOk() (*float32, bool)`

GetMovieQuotaLimitOk returns a tuple with the MovieQuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieQuotaLimit

`func (o *UserSettings) SetMovieQuotaLimit(v float32)`

SetMovieQuotaLimit sets MovieQuotaLimit field to given value.

### HasMovieQuotaLimit

`func (o *UserSettings) HasMovieQuotaLimit() bool`

HasMovieQuotaLimit returns a boolean if a field has been set.

### SetMovieQuotaLimitNil

`func (o *UserSettings) SetMovieQuotaLimitNil(b bool)`

 SetMovieQuotaLimitNil sets the value for MovieQuotaLimit to be an explicit nil

### UnsetMovieQuotaLimit
`func (o *UserSettings) UnsetMovieQuotaLimit()`

UnsetMovieQuotaLimit ensures that no value is present for MovieQuotaLimit, not even an explicit nil
### GetMovieQuotaDays

`func (o *UserSettings) GetMovieQuotaDays() float32`

GetMovieQuotaDays returns the MovieQuotaDays field if non-nil, zero value otherwise.

### GetMovieQuotaDaysOk

`func (o *UserSettings) GetMovieQuotaDaysOk() (*float32, bool)`

GetMovieQuotaDaysOk returns a tuple with the MovieQuotaDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieQuotaDays

`func (o *UserSettings) SetMovieQuotaDays(v float32)`

SetMovieQuotaDays sets MovieQuotaDays field to given value.

### HasMovieQuotaDays

`func (o *UserSettings) HasMovieQuotaDays() bool`

HasMovieQuotaDays returns a boolean if a field has been set.

### SetMovieQuotaDaysNil

`func (o *UserSettings) SetMovieQuotaDaysNil(b bool)`

 SetMovieQuotaDaysNil sets the value for MovieQuotaDays to be an explicit nil

### UnsetMovieQuotaDays
`func (o *UserSettings) UnsetMovieQuotaDays()`

UnsetMovieQuotaDays ensures that no value is present for MovieQuotaDays, not even an explicit nil
### GetTvQuotaLimit

`func (o *UserSettings) GetTvQuotaLimit() float32`

GetTvQuotaLimit returns the TvQuotaLimit field if non-nil, zero value otherwise.

### GetTvQuotaLimitOk

`func (o *UserSettings) GetTvQuotaLimitOk() (*float32, bool)`

GetTvQuotaLimitOk returns a tuple with the TvQuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvQuotaLimit

`func (o *UserSettings) SetTvQuotaLimit(v float32)`

SetTvQuotaLimit sets TvQuotaLimit field to given value.

### HasTvQuotaLimit

`func (o *UserSettings) HasTvQuotaLimit() bool`

HasTvQuotaLimit returns a boolean if a field has been set.

### SetTvQuotaLimitNil

`func (o *UserSettings) SetTvQuotaLimitNil(b bool)`

 SetTvQuotaLimitNil sets the value for TvQuotaLimit to be an explicit nil

### UnsetTvQuotaLimit
`func (o *UserSettings) UnsetTvQuotaLimit()`

UnsetTvQuotaLimit ensures that no value is present for TvQuotaLimit, not even an explicit nil
### GetTvQuotaDays

`func (o *UserSettings) GetTvQuotaDays() float32`

GetTvQuotaDays returns the TvQuotaDays field if non-nil, zero value otherwise.

### GetTvQuotaDaysOk

`func (o *UserSettings) GetTvQuotaDaysOk() (*float32, bool)`

GetTvQuotaDaysOk returns a tuple with the TvQuotaDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvQuotaDays

`func (o *UserSettings) SetTvQuotaDays(v float32)`

SetTvQuotaDays sets TvQuotaDays field to given value.

### HasTvQuotaDays

`func (o *UserSettings) HasTvQuotaDays() bool`

HasTvQuotaDays returns a boolean if a field has been set.

### SetTvQuotaDaysNil

`func (o *UserSettings) SetTvQuotaDaysNil(b bool)`

 SetTvQuotaDaysNil sets the value for TvQuotaDays to be an explicit nil

### UnsetTvQuotaDays
`func (o *UserSettings) UnsetTvQuotaDays()`

UnsetTvQuotaDays ensures that no value is present for TvQuotaDays, not even an explicit nil
### GetGlobalMovieQuotaDays

`func (o *UserSettings) GetGlobalMovieQuotaDays() float32`

GetGlobalMovieQuotaDays returns the GlobalMovieQuotaDays field if non-nil, zero value otherwise.

### GetGlobalMovieQuotaDaysOk

`func (o *UserSettings) GetGlobalMovieQuotaDaysOk() (*float32, bool)`

GetGlobalMovieQuotaDaysOk returns a tuple with the GlobalMovieQuotaDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalMovieQuotaDays

`func (o *UserSettings) SetGlobalMovieQuotaDays(v float32)`

SetGlobalMovieQuotaDays sets GlobalMovieQuotaDays field to given value.

### HasGlobalMovieQuotaDays

`func (o *UserSettings) HasGlobalMovieQuotaDays() bool`

HasGlobalMovieQuotaDays returns a boolean if a field has been set.

### SetGlobalMovieQuotaDaysNil

`func (o *UserSettings) SetGlobalMovieQuotaDaysNil(b bool)`

 SetGlobalMovieQuotaDaysNil sets the value for GlobalMovieQuotaDays to be an explicit nil

### UnsetGlobalMovieQuotaDays
`func (o *UserSettings) UnsetGlobalMovieQuotaDays()`

UnsetGlobalMovieQuotaDays ensures that no value is present for GlobalMovieQuotaDays, not even an explicit nil
### GetGlobalMovieQuotaLimit

`func (o *UserSettings) GetGlobalMovieQuotaLimit() float32`

GetGlobalMovieQuotaLimit returns the GlobalMovieQuotaLimit field if non-nil, zero value otherwise.

### GetGlobalMovieQuotaLimitOk

`func (o *UserSettings) GetGlobalMovieQuotaLimitOk() (*float32, bool)`

GetGlobalMovieQuotaLimitOk returns a tuple with the GlobalMovieQuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalMovieQuotaLimit

`func (o *UserSettings) SetGlobalMovieQuotaLimit(v float32)`

SetGlobalMovieQuotaLimit sets GlobalMovieQuotaLimit field to given value.

### HasGlobalMovieQuotaLimit

`func (o *UserSettings) HasGlobalMovieQuotaLimit() bool`

HasGlobalMovieQuotaLimit returns a boolean if a field has been set.

### SetGlobalMovieQuotaLimitNil

`func (o *UserSettings) SetGlobalMovieQuotaLimitNil(b bool)`

 SetGlobalMovieQuotaLimitNil sets the value for GlobalMovieQuotaLimit to be an explicit nil

### UnsetGlobalMovieQuotaLimit
`func (o *UserSettings) UnsetGlobalMovieQuotaLimit()`

UnsetGlobalMovieQuotaLimit ensures that no value is present for GlobalMovieQuotaLimit, not even an explicit nil
### GetGlobalTvQuotaLimit

`func (o *UserSettings) GetGlobalTvQuotaLimit() float32`

GetGlobalTvQuotaLimit returns the GlobalTvQuotaLimit field if non-nil, zero value otherwise.

### GetGlobalTvQuotaLimitOk

`func (o *UserSettings) GetGlobalTvQuotaLimitOk() (*float32, bool)`

GetGlobalTvQuotaLimitOk returns a tuple with the GlobalTvQuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTvQuotaLimit

`func (o *UserSettings) SetGlobalTvQuotaLimit(v float32)`

SetGlobalTvQuotaLimit sets GlobalTvQuotaLimit field to given value.

### HasGlobalTvQuotaLimit

`func (o *UserSettings) HasGlobalTvQuotaLimit() bool`

HasGlobalTvQuotaLimit returns a boolean if a field has been set.

### SetGlobalTvQuotaLimitNil

`func (o *UserSettings) SetGlobalTvQuotaLimitNil(b bool)`

 SetGlobalTvQuotaLimitNil sets the value for GlobalTvQuotaLimit to be an explicit nil

### UnsetGlobalTvQuotaLimit
`func (o *UserSettings) UnsetGlobalTvQuotaLimit()`

UnsetGlobalTvQuotaLimit ensures that no value is present for GlobalTvQuotaLimit, not even an explicit nil
### GetGlobalTvQuotaDays

`func (o *UserSettings) GetGlobalTvQuotaDays() float32`

GetGlobalTvQuotaDays returns the GlobalTvQuotaDays field if non-nil, zero value otherwise.

### GetGlobalTvQuotaDaysOk

`func (o *UserSettings) GetGlobalTvQuotaDaysOk() (*float32, bool)`

GetGlobalTvQuotaDaysOk returns a tuple with the GlobalTvQuotaDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTvQuotaDays

`func (o *UserSettings) SetGlobalTvQuotaDays(v float32)`

SetGlobalTvQuotaDays sets GlobalTvQuotaDays field to given value.

### HasGlobalTvQuotaDays

`func (o *UserSettings) HasGlobalTvQuotaDays() bool`

HasGlobalTvQuotaDays returns a boolean if a field has been set.

### SetGlobalTvQuotaDaysNil

`func (o *UserSettings) SetGlobalTvQuotaDaysNil(b bool)`

 SetGlobalTvQuotaDaysNil sets the value for GlobalTvQuotaDays to be an explicit nil

### UnsetGlobalTvQuotaDays
`func (o *UserSettings) UnsetGlobalTvQuotaDays()`

UnsetGlobalTvQuotaDays ensures that no value is present for GlobalTvQuotaDays, not even an explicit nil
### GetWatchlistSyncMovies

`func (o *UserSettings) GetWatchlistSyncMovies() bool`

GetWatchlistSyncMovies returns the WatchlistSyncMovies field if non-nil, zero value otherwise.

### GetWatchlistSyncMoviesOk

`func (o *UserSettings) GetWatchlistSyncMoviesOk() (*bool, bool)`

GetWatchlistSyncMoviesOk returns a tuple with the WatchlistSyncMovies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistSyncMovies

`func (o *UserSettings) SetWatchlistSyncMovies(v bool)`

SetWatchlistSyncMovies sets WatchlistSyncMovies field to given value.

### HasWatchlistSyncMovies

`func (o *UserSettings) HasWatchlistSyncMovies() bool`

HasWatchlistSyncMovies returns a boolean if a field has been set.

### SetWatchlistSyncMoviesNil

`func (o *UserSettings) SetWatchlistSyncMoviesNil(b bool)`

 SetWatchlistSyncMoviesNil sets the value for WatchlistSyncMovies to be an explicit nil

### UnsetWatchlistSyncMovies
`func (o *UserSettings) UnsetWatchlistSyncMovies()`

UnsetWatchlistSyncMovies ensures that no value is present for WatchlistSyncMovies, not even an explicit nil
### GetWatchlistSyncTv

`func (o *UserSettings) GetWatchlistSyncTv() bool`

GetWatchlistSyncTv returns the WatchlistSyncTv field if non-nil, zero value otherwise.

### GetWatchlistSyncTvOk

`func (o *UserSettings) GetWatchlistSyncTvOk() (*bool, bool)`

GetWatchlistSyncTvOk returns a tuple with the WatchlistSyncTv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistSyncTv

`func (o *UserSettings) SetWatchlistSyncTv(v bool)`

SetWatchlistSyncTv sets WatchlistSyncTv field to given value.

### HasWatchlistSyncTv

`func (o *UserSettings) HasWatchlistSyncTv() bool`

HasWatchlistSyncTv returns a boolean if a field has been set.

### SetWatchlistSyncTvNil

`func (o *UserSettings) SetWatchlistSyncTvNil(b bool)`

 SetWatchlistSyncTvNil sets the value for WatchlistSyncTv to be an explicit nil

### UnsetWatchlistSyncTv
`func (o *UserSettings) UnsetWatchlistSyncTv()`

UnsetWatchlistSyncTv ensures that no value is present for WatchlistSyncTv, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


