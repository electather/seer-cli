# SonarrSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**SortTitle** | Pointer to **string** |  | [optional] 
**SeasonCount** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**AirTime** | Pointer to **string** |  | [optional] 
**Images** | Pointer to [**[]SonarrSeriesImagesInner**](SonarrSeriesImagesInner.md) |  | [optional] 
**RemotePoster** | Pointer to **string** |  | [optional] 
**Seasons** | Pointer to [**[]SonarrSeriesSeasonsInner**](SonarrSeriesSeasonsInner.md) |  | [optional] 
**Year** | Pointer to **float32** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **float32** |  | [optional] 
**LanguageProfileId** | Pointer to **float32** |  | [optional] 
**SeasonFolder** | Pointer to **bool** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**UseSceneNumbering** | Pointer to **bool** |  | [optional] 
**Runtime** | Pointer to **float32** |  | [optional] 
**TvdbId** | Pointer to **float32** |  | [optional] 
**TvRageId** | Pointer to **float32** |  | [optional] 
**TvMazeId** | Pointer to **float32** |  | [optional] 
**FirstAired** | Pointer to **string** |  | [optional] 
**LastInfoSync** | Pointer to **NullableString** |  | [optional] 
**SeriesType** | Pointer to **string** |  | [optional] 
**CleanTitle** | Pointer to **string** |  | [optional] 
**ImdbId** | Pointer to **string** |  | [optional] 
**TitleSlug** | Pointer to **string** |  | [optional] 
**Certification** | Pointer to **string** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Added** | Pointer to **string** |  | [optional] 
**Ratings** | Pointer to [**[]SonarrSeriesRatingsInner**](SonarrSeriesRatingsInner.md) |  | [optional] 
**QualityProfileId** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **NullableFloat32** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**AddOptions** | Pointer to [**[]SonarrSeriesAddOptionsInner**](SonarrSeriesAddOptionsInner.md) |  | [optional] 

## Methods

### NewSonarrSeries

`func NewSonarrSeries() *SonarrSeries`

NewSonarrSeries instantiates a new SonarrSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSonarrSeriesWithDefaults

`func NewSonarrSeriesWithDefaults() *SonarrSeries`

NewSonarrSeriesWithDefaults instantiates a new SonarrSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SonarrSeries) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SonarrSeries) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SonarrSeries) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SonarrSeries) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSortTitle

`func (o *SonarrSeries) GetSortTitle() string`

GetSortTitle returns the SortTitle field if non-nil, zero value otherwise.

### GetSortTitleOk

`func (o *SonarrSeries) GetSortTitleOk() (*string, bool)`

GetSortTitleOk returns a tuple with the SortTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortTitle

`func (o *SonarrSeries) SetSortTitle(v string)`

SetSortTitle sets SortTitle field to given value.

### HasSortTitle

`func (o *SonarrSeries) HasSortTitle() bool`

HasSortTitle returns a boolean if a field has been set.

### GetSeasonCount

`func (o *SonarrSeries) GetSeasonCount() float32`

GetSeasonCount returns the SeasonCount field if non-nil, zero value otherwise.

### GetSeasonCountOk

`func (o *SonarrSeries) GetSeasonCountOk() (*float32, bool)`

GetSeasonCountOk returns a tuple with the SeasonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonCount

`func (o *SonarrSeries) SetSeasonCount(v float32)`

SetSeasonCount sets SeasonCount field to given value.

### HasSeasonCount

`func (o *SonarrSeries) HasSeasonCount() bool`

HasSeasonCount returns a boolean if a field has been set.

### GetStatus

`func (o *SonarrSeries) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SonarrSeries) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SonarrSeries) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SonarrSeries) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOverview

`func (o *SonarrSeries) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *SonarrSeries) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *SonarrSeries) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *SonarrSeries) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetNetwork

`func (o *SonarrSeries) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SonarrSeries) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SonarrSeries) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *SonarrSeries) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetAirTime

`func (o *SonarrSeries) GetAirTime() string`

GetAirTime returns the AirTime field if non-nil, zero value otherwise.

### GetAirTimeOk

`func (o *SonarrSeries) GetAirTimeOk() (*string, bool)`

GetAirTimeOk returns a tuple with the AirTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirTime

`func (o *SonarrSeries) SetAirTime(v string)`

SetAirTime sets AirTime field to given value.

### HasAirTime

`func (o *SonarrSeries) HasAirTime() bool`

HasAirTime returns a boolean if a field has been set.

### GetImages

`func (o *SonarrSeries) GetImages() []SonarrSeriesImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SonarrSeries) GetImagesOk() (*[]SonarrSeriesImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SonarrSeries) SetImages(v []SonarrSeriesImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *SonarrSeries) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetRemotePoster

`func (o *SonarrSeries) GetRemotePoster() string`

GetRemotePoster returns the RemotePoster field if non-nil, zero value otherwise.

### GetRemotePosterOk

`func (o *SonarrSeries) GetRemotePosterOk() (*string, bool)`

GetRemotePosterOk returns a tuple with the RemotePoster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePoster

`func (o *SonarrSeries) SetRemotePoster(v string)`

SetRemotePoster sets RemotePoster field to given value.

### HasRemotePoster

`func (o *SonarrSeries) HasRemotePoster() bool`

HasRemotePoster returns a boolean if a field has been set.

### GetSeasons

`func (o *SonarrSeries) GetSeasons() []SonarrSeriesSeasonsInner`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *SonarrSeries) GetSeasonsOk() (*[]SonarrSeriesSeasonsInner, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *SonarrSeries) SetSeasons(v []SonarrSeriesSeasonsInner)`

SetSeasons sets Seasons field to given value.

### HasSeasons

`func (o *SonarrSeries) HasSeasons() bool`

HasSeasons returns a boolean if a field has been set.

### GetYear

`func (o *SonarrSeries) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SonarrSeries) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SonarrSeries) SetYear(v float32)`

SetYear sets Year field to given value.

### HasYear

`func (o *SonarrSeries) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetPath

`func (o *SonarrSeries) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SonarrSeries) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SonarrSeries) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SonarrSeries) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProfileId

`func (o *SonarrSeries) GetProfileId() float32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *SonarrSeries) GetProfileIdOk() (*float32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *SonarrSeries) SetProfileId(v float32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *SonarrSeries) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetLanguageProfileId

`func (o *SonarrSeries) GetLanguageProfileId() float32`

GetLanguageProfileId returns the LanguageProfileId field if non-nil, zero value otherwise.

### GetLanguageProfileIdOk

`func (o *SonarrSeries) GetLanguageProfileIdOk() (*float32, bool)`

GetLanguageProfileIdOk returns a tuple with the LanguageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageProfileId

`func (o *SonarrSeries) SetLanguageProfileId(v float32)`

SetLanguageProfileId sets LanguageProfileId field to given value.

### HasLanguageProfileId

`func (o *SonarrSeries) HasLanguageProfileId() bool`

HasLanguageProfileId returns a boolean if a field has been set.

### GetSeasonFolder

`func (o *SonarrSeries) GetSeasonFolder() bool`

GetSeasonFolder returns the SeasonFolder field if non-nil, zero value otherwise.

### GetSeasonFolderOk

`func (o *SonarrSeries) GetSeasonFolderOk() (*bool, bool)`

GetSeasonFolderOk returns a tuple with the SeasonFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonFolder

`func (o *SonarrSeries) SetSeasonFolder(v bool)`

SetSeasonFolder sets SeasonFolder field to given value.

### HasSeasonFolder

`func (o *SonarrSeries) HasSeasonFolder() bool`

HasSeasonFolder returns a boolean if a field has been set.

### GetMonitored

`func (o *SonarrSeries) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *SonarrSeries) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *SonarrSeries) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *SonarrSeries) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetUseSceneNumbering

`func (o *SonarrSeries) GetUseSceneNumbering() bool`

GetUseSceneNumbering returns the UseSceneNumbering field if non-nil, zero value otherwise.

### GetUseSceneNumberingOk

`func (o *SonarrSeries) GetUseSceneNumberingOk() (*bool, bool)`

GetUseSceneNumberingOk returns a tuple with the UseSceneNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSceneNumbering

`func (o *SonarrSeries) SetUseSceneNumbering(v bool)`

SetUseSceneNumbering sets UseSceneNumbering field to given value.

### HasUseSceneNumbering

`func (o *SonarrSeries) HasUseSceneNumbering() bool`

HasUseSceneNumbering returns a boolean if a field has been set.

### GetRuntime

`func (o *SonarrSeries) GetRuntime() float32`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *SonarrSeries) GetRuntimeOk() (*float32, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *SonarrSeries) SetRuntime(v float32)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *SonarrSeries) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetTvdbId

`func (o *SonarrSeries) GetTvdbId() float32`

GetTvdbId returns the TvdbId field if non-nil, zero value otherwise.

### GetTvdbIdOk

`func (o *SonarrSeries) GetTvdbIdOk() (*float32, bool)`

GetTvdbIdOk returns a tuple with the TvdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvdbId

`func (o *SonarrSeries) SetTvdbId(v float32)`

SetTvdbId sets TvdbId field to given value.

### HasTvdbId

`func (o *SonarrSeries) HasTvdbId() bool`

HasTvdbId returns a boolean if a field has been set.

### GetTvRageId

`func (o *SonarrSeries) GetTvRageId() float32`

GetTvRageId returns the TvRageId field if non-nil, zero value otherwise.

### GetTvRageIdOk

`func (o *SonarrSeries) GetTvRageIdOk() (*float32, bool)`

GetTvRageIdOk returns a tuple with the TvRageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvRageId

`func (o *SonarrSeries) SetTvRageId(v float32)`

SetTvRageId sets TvRageId field to given value.

### HasTvRageId

`func (o *SonarrSeries) HasTvRageId() bool`

HasTvRageId returns a boolean if a field has been set.

### GetTvMazeId

`func (o *SonarrSeries) GetTvMazeId() float32`

GetTvMazeId returns the TvMazeId field if non-nil, zero value otherwise.

### GetTvMazeIdOk

`func (o *SonarrSeries) GetTvMazeIdOk() (*float32, bool)`

GetTvMazeIdOk returns a tuple with the TvMazeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvMazeId

`func (o *SonarrSeries) SetTvMazeId(v float32)`

SetTvMazeId sets TvMazeId field to given value.

### HasTvMazeId

`func (o *SonarrSeries) HasTvMazeId() bool`

HasTvMazeId returns a boolean if a field has been set.

### GetFirstAired

`func (o *SonarrSeries) GetFirstAired() string`

GetFirstAired returns the FirstAired field if non-nil, zero value otherwise.

### GetFirstAiredOk

`func (o *SonarrSeries) GetFirstAiredOk() (*string, bool)`

GetFirstAiredOk returns a tuple with the FirstAired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAired

`func (o *SonarrSeries) SetFirstAired(v string)`

SetFirstAired sets FirstAired field to given value.

### HasFirstAired

`func (o *SonarrSeries) HasFirstAired() bool`

HasFirstAired returns a boolean if a field has been set.

### GetLastInfoSync

`func (o *SonarrSeries) GetLastInfoSync() string`

GetLastInfoSync returns the LastInfoSync field if non-nil, zero value otherwise.

### GetLastInfoSyncOk

`func (o *SonarrSeries) GetLastInfoSyncOk() (*string, bool)`

GetLastInfoSyncOk returns a tuple with the LastInfoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInfoSync

`func (o *SonarrSeries) SetLastInfoSync(v string)`

SetLastInfoSync sets LastInfoSync field to given value.

### HasLastInfoSync

`func (o *SonarrSeries) HasLastInfoSync() bool`

HasLastInfoSync returns a boolean if a field has been set.

### SetLastInfoSyncNil

`func (o *SonarrSeries) SetLastInfoSyncNil(b bool)`

 SetLastInfoSyncNil sets the value for LastInfoSync to be an explicit nil

### UnsetLastInfoSync
`func (o *SonarrSeries) UnsetLastInfoSync()`

UnsetLastInfoSync ensures that no value is present for LastInfoSync, not even an explicit nil
### GetSeriesType

`func (o *SonarrSeries) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *SonarrSeries) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *SonarrSeries) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *SonarrSeries) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetCleanTitle

`func (o *SonarrSeries) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *SonarrSeries) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *SonarrSeries) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *SonarrSeries) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### GetImdbId

`func (o *SonarrSeries) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *SonarrSeries) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *SonarrSeries) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *SonarrSeries) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### GetTitleSlug

`func (o *SonarrSeries) GetTitleSlug() string`

GetTitleSlug returns the TitleSlug field if non-nil, zero value otherwise.

### GetTitleSlugOk

`func (o *SonarrSeries) GetTitleSlugOk() (*string, bool)`

GetTitleSlugOk returns a tuple with the TitleSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSlug

`func (o *SonarrSeries) SetTitleSlug(v string)`

SetTitleSlug sets TitleSlug field to given value.

### HasTitleSlug

`func (o *SonarrSeries) HasTitleSlug() bool`

HasTitleSlug returns a boolean if a field has been set.

### GetCertification

`func (o *SonarrSeries) GetCertification() string`

GetCertification returns the Certification field if non-nil, zero value otherwise.

### GetCertificationOk

`func (o *SonarrSeries) GetCertificationOk() (*string, bool)`

GetCertificationOk returns a tuple with the Certification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertification

`func (o *SonarrSeries) SetCertification(v string)`

SetCertification sets Certification field to given value.

### HasCertification

`func (o *SonarrSeries) HasCertification() bool`

HasCertification returns a boolean if a field has been set.

### GetGenres

`func (o *SonarrSeries) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *SonarrSeries) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *SonarrSeries) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *SonarrSeries) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetTags

`func (o *SonarrSeries) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SonarrSeries) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SonarrSeries) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SonarrSeries) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAdded

`func (o *SonarrSeries) GetAdded() string`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *SonarrSeries) GetAddedOk() (*string, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *SonarrSeries) SetAdded(v string)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *SonarrSeries) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetRatings

`func (o *SonarrSeries) GetRatings() []SonarrSeriesRatingsInner`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *SonarrSeries) GetRatingsOk() (*[]SonarrSeriesRatingsInner, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *SonarrSeries) SetRatings(v []SonarrSeriesRatingsInner)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *SonarrSeries) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetQualityProfileId

`func (o *SonarrSeries) GetQualityProfileId() float32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *SonarrSeries) GetQualityProfileIdOk() (*float32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *SonarrSeries) SetQualityProfileId(v float32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *SonarrSeries) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### GetId

`func (o *SonarrSeries) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SonarrSeries) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SonarrSeries) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *SonarrSeries) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SonarrSeries) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SonarrSeries) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRootFolderPath

`func (o *SonarrSeries) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *SonarrSeries) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *SonarrSeries) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *SonarrSeries) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *SonarrSeries) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *SonarrSeries) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetAddOptions

`func (o *SonarrSeries) GetAddOptions() []SonarrSeriesAddOptionsInner`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *SonarrSeries) GetAddOptionsOk() (*[]SonarrSeriesAddOptionsInner, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *SonarrSeries) SetAddOptions(v []SonarrSeriesAddOptionsInner)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *SonarrSeries) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


