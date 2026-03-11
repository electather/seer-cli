# Season

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**AirDate** | Pointer to **NullableString** |  | [optional] 
**EpisodeCount** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**PosterPath** | Pointer to **string** |  | [optional] 
**SeasonNumber** | Pointer to **float32** |  | [optional] 
**Episodes** | Pointer to [**[]Episode**](Episode.md) |  | [optional] 

## Methods

### NewSeason

`func NewSeason() *Season`

NewSeason instantiates a new Season object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonWithDefaults

`func NewSeasonWithDefaults() *Season`

NewSeasonWithDefaults instantiates a new Season object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Season) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Season) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Season) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *Season) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAirDate

`func (o *Season) GetAirDate() string`

GetAirDate returns the AirDate field if non-nil, zero value otherwise.

### GetAirDateOk

`func (o *Season) GetAirDateOk() (*string, bool)`

GetAirDateOk returns a tuple with the AirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDate

`func (o *Season) SetAirDate(v string)`

SetAirDate sets AirDate field to given value.

### HasAirDate

`func (o *Season) HasAirDate() bool`

HasAirDate returns a boolean if a field has been set.

### SetAirDateNil

`func (o *Season) SetAirDateNil(b bool)`

 SetAirDateNil sets the value for AirDate to be an explicit nil

### UnsetAirDate
`func (o *Season) UnsetAirDate()`

UnsetAirDate ensures that no value is present for AirDate, not even an explicit nil
### GetEpisodeCount

`func (o *Season) GetEpisodeCount() float32`

GetEpisodeCount returns the EpisodeCount field if non-nil, zero value otherwise.

### GetEpisodeCountOk

`func (o *Season) GetEpisodeCountOk() (*float32, bool)`

GetEpisodeCountOk returns a tuple with the EpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeCount

`func (o *Season) SetEpisodeCount(v float32)`

SetEpisodeCount sets EpisodeCount field to given value.

### HasEpisodeCount

`func (o *Season) HasEpisodeCount() bool`

HasEpisodeCount returns a boolean if a field has been set.

### GetName

`func (o *Season) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Season) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Season) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Season) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverview

`func (o *Season) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *Season) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *Season) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *Season) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetPosterPath

`func (o *Season) GetPosterPath() string`

GetPosterPath returns the PosterPath field if non-nil, zero value otherwise.

### GetPosterPathOk

`func (o *Season) GetPosterPathOk() (*string, bool)`

GetPosterPathOk returns a tuple with the PosterPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosterPath

`func (o *Season) SetPosterPath(v string)`

SetPosterPath sets PosterPath field to given value.

### HasPosterPath

`func (o *Season) HasPosterPath() bool`

HasPosterPath returns a boolean if a field has been set.

### GetSeasonNumber

`func (o *Season) GetSeasonNumber() float32`

GetSeasonNumber returns the SeasonNumber field if non-nil, zero value otherwise.

### GetSeasonNumberOk

`func (o *Season) GetSeasonNumberOk() (*float32, bool)`

GetSeasonNumberOk returns a tuple with the SeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonNumber

`func (o *Season) SetSeasonNumber(v float32)`

SetSeasonNumber sets SeasonNumber field to given value.

### HasSeasonNumber

`func (o *Season) HasSeasonNumber() bool`

HasSeasonNumber returns a boolean if a field has been set.

### GetEpisodes

`func (o *Season) GetEpisodes() []Episode`

GetEpisodes returns the Episodes field if non-nil, zero value otherwise.

### GetEpisodesOk

`func (o *Season) GetEpisodesOk() (*[]Episode, bool)`

GetEpisodesOk returns a tuple with the Episodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodes

`func (o *Season) SetEpisodes(v []Episode)`

SetEpisodes sets Episodes field to given value.

### HasEpisodes

`func (o *Season) HasEpisodes() bool`

HasEpisodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


