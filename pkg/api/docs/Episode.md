# Episode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AirDate** | Pointer to **NullableString** |  | [optional] 
**EpisodeNumber** | Pointer to **float32** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**ProductionCode** | Pointer to **string** |  | [optional] 
**SeasonNumber** | Pointer to **float32** |  | [optional] 
**ShowId** | Pointer to **float32** |  | [optional] 
**StillPath** | Pointer to **NullableString** |  | [optional] 
**VoteAverage** | Pointer to **float32** |  | [optional] 
**VoteCount** | Pointer to **float32** |  | [optional] 

## Methods

### NewEpisode

`func NewEpisode() *Episode`

NewEpisode instantiates a new Episode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpisodeWithDefaults

`func NewEpisodeWithDefaults() *Episode`

NewEpisodeWithDefaults instantiates a new Episode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Episode) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Episode) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Episode) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *Episode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Episode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Episode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Episode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Episode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAirDate

`func (o *Episode) GetAirDate() string`

GetAirDate returns the AirDate field if non-nil, zero value otherwise.

### GetAirDateOk

`func (o *Episode) GetAirDateOk() (*string, bool)`

GetAirDateOk returns a tuple with the AirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDate

`func (o *Episode) SetAirDate(v string)`

SetAirDate sets AirDate field to given value.

### HasAirDate

`func (o *Episode) HasAirDate() bool`

HasAirDate returns a boolean if a field has been set.

### SetAirDateNil

`func (o *Episode) SetAirDateNil(b bool)`

 SetAirDateNil sets the value for AirDate to be an explicit nil

### UnsetAirDate
`func (o *Episode) UnsetAirDate()`

UnsetAirDate ensures that no value is present for AirDate, not even an explicit nil
### GetEpisodeNumber

`func (o *Episode) GetEpisodeNumber() float32`

GetEpisodeNumber returns the EpisodeNumber field if non-nil, zero value otherwise.

### GetEpisodeNumberOk

`func (o *Episode) GetEpisodeNumberOk() (*float32, bool)`

GetEpisodeNumberOk returns a tuple with the EpisodeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeNumber

`func (o *Episode) SetEpisodeNumber(v float32)`

SetEpisodeNumber sets EpisodeNumber field to given value.

### HasEpisodeNumber

`func (o *Episode) HasEpisodeNumber() bool`

HasEpisodeNumber returns a boolean if a field has been set.

### GetOverview

`func (o *Episode) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *Episode) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *Episode) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *Episode) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetProductionCode

`func (o *Episode) GetProductionCode() string`

GetProductionCode returns the ProductionCode field if non-nil, zero value otherwise.

### GetProductionCodeOk

`func (o *Episode) GetProductionCodeOk() (*string, bool)`

GetProductionCodeOk returns a tuple with the ProductionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionCode

`func (o *Episode) SetProductionCode(v string)`

SetProductionCode sets ProductionCode field to given value.

### HasProductionCode

`func (o *Episode) HasProductionCode() bool`

HasProductionCode returns a boolean if a field has been set.

### GetSeasonNumber

`func (o *Episode) GetSeasonNumber() float32`

GetSeasonNumber returns the SeasonNumber field if non-nil, zero value otherwise.

### GetSeasonNumberOk

`func (o *Episode) GetSeasonNumberOk() (*float32, bool)`

GetSeasonNumberOk returns a tuple with the SeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonNumber

`func (o *Episode) SetSeasonNumber(v float32)`

SetSeasonNumber sets SeasonNumber field to given value.

### HasSeasonNumber

`func (o *Episode) HasSeasonNumber() bool`

HasSeasonNumber returns a boolean if a field has been set.

### GetShowId

`func (o *Episode) GetShowId() float32`

GetShowId returns the ShowId field if non-nil, zero value otherwise.

### GetShowIdOk

`func (o *Episode) GetShowIdOk() (*float32, bool)`

GetShowIdOk returns a tuple with the ShowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowId

`func (o *Episode) SetShowId(v float32)`

SetShowId sets ShowId field to given value.

### HasShowId

`func (o *Episode) HasShowId() bool`

HasShowId returns a boolean if a field has been set.

### GetStillPath

`func (o *Episode) GetStillPath() string`

GetStillPath returns the StillPath field if non-nil, zero value otherwise.

### GetStillPathOk

`func (o *Episode) GetStillPathOk() (*string, bool)`

GetStillPathOk returns a tuple with the StillPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStillPath

`func (o *Episode) SetStillPath(v string)`

SetStillPath sets StillPath field to given value.

### HasStillPath

`func (o *Episode) HasStillPath() bool`

HasStillPath returns a boolean if a field has been set.

### SetStillPathNil

`func (o *Episode) SetStillPathNil(b bool)`

 SetStillPathNil sets the value for StillPath to be an explicit nil

### UnsetStillPath
`func (o *Episode) UnsetStillPath()`

UnsetStillPath ensures that no value is present for StillPath, not even an explicit nil
### GetVoteAverage

`func (o *Episode) GetVoteAverage() float32`

GetVoteAverage returns the VoteAverage field if non-nil, zero value otherwise.

### GetVoteAverageOk

`func (o *Episode) GetVoteAverageOk() (*float32, bool)`

GetVoteAverageOk returns a tuple with the VoteAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteAverage

`func (o *Episode) SetVoteAverage(v float32)`

SetVoteAverage sets VoteAverage field to given value.

### HasVoteAverage

`func (o *Episode) HasVoteAverage() bool`

HasVoteAverage returns a boolean if a field has been set.

### GetVoteCount

`func (o *Episode) GetVoteCount() float32`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *Episode) GetVoteCountOk() (*float32, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *Episode) SetVoteCount(v float32)`

SetVoteCount sets VoteCount field to given value.

### HasVoteCount

`func (o *Episode) HasVoteCount() bool`

HasVoteCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


