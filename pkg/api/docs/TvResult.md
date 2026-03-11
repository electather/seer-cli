# TvResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**Popularity** | Pointer to **float32** |  | [optional] 
**PosterPath** | Pointer to **string** |  | [optional] 
**BackdropPath** | Pointer to **string** |  | [optional] 
**VoteCount** | Pointer to **float32** |  | [optional] 
**VoteAverage** | Pointer to **float32** |  | [optional] 
**GenreIds** | Pointer to **[]float32** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**OriginalLanguage** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OriginalName** | Pointer to **string** |  | [optional] 
**OriginCountry** | Pointer to **[]string** |  | [optional] 
**FirstAirDate** | Pointer to **string** |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 

## Methods

### NewTvResult

`func NewTvResult() *TvResult`

NewTvResult instantiates a new TvResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTvResultWithDefaults

`func NewTvResultWithDefaults() *TvResult`

NewTvResultWithDefaults instantiates a new TvResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TvResult) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TvResult) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TvResult) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *TvResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMediaType

`func (o *TvResult) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *TvResult) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *TvResult) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *TvResult) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetPopularity

`func (o *TvResult) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *TvResult) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *TvResult) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *TvResult) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetPosterPath

`func (o *TvResult) GetPosterPath() string`

GetPosterPath returns the PosterPath field if non-nil, zero value otherwise.

### GetPosterPathOk

`func (o *TvResult) GetPosterPathOk() (*string, bool)`

GetPosterPathOk returns a tuple with the PosterPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosterPath

`func (o *TvResult) SetPosterPath(v string)`

SetPosterPath sets PosterPath field to given value.

### HasPosterPath

`func (o *TvResult) HasPosterPath() bool`

HasPosterPath returns a boolean if a field has been set.

### GetBackdropPath

`func (o *TvResult) GetBackdropPath() string`

GetBackdropPath returns the BackdropPath field if non-nil, zero value otherwise.

### GetBackdropPathOk

`func (o *TvResult) GetBackdropPathOk() (*string, bool)`

GetBackdropPathOk returns a tuple with the BackdropPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropPath

`func (o *TvResult) SetBackdropPath(v string)`

SetBackdropPath sets BackdropPath field to given value.

### HasBackdropPath

`func (o *TvResult) HasBackdropPath() bool`

HasBackdropPath returns a boolean if a field has been set.

### GetVoteCount

`func (o *TvResult) GetVoteCount() float32`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *TvResult) GetVoteCountOk() (*float32, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *TvResult) SetVoteCount(v float32)`

SetVoteCount sets VoteCount field to given value.

### HasVoteCount

`func (o *TvResult) HasVoteCount() bool`

HasVoteCount returns a boolean if a field has been set.

### GetVoteAverage

`func (o *TvResult) GetVoteAverage() float32`

GetVoteAverage returns the VoteAverage field if non-nil, zero value otherwise.

### GetVoteAverageOk

`func (o *TvResult) GetVoteAverageOk() (*float32, bool)`

GetVoteAverageOk returns a tuple with the VoteAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteAverage

`func (o *TvResult) SetVoteAverage(v float32)`

SetVoteAverage sets VoteAverage field to given value.

### HasVoteAverage

`func (o *TvResult) HasVoteAverage() bool`

HasVoteAverage returns a boolean if a field has been set.

### GetGenreIds

`func (o *TvResult) GetGenreIds() []float32`

GetGenreIds returns the GenreIds field if non-nil, zero value otherwise.

### GetGenreIdsOk

`func (o *TvResult) GetGenreIdsOk() (*[]float32, bool)`

GetGenreIdsOk returns a tuple with the GenreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreIds

`func (o *TvResult) SetGenreIds(v []float32)`

SetGenreIds sets GenreIds field to given value.

### HasGenreIds

`func (o *TvResult) HasGenreIds() bool`

HasGenreIds returns a boolean if a field has been set.

### GetOverview

`func (o *TvResult) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *TvResult) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *TvResult) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *TvResult) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetOriginalLanguage

`func (o *TvResult) GetOriginalLanguage() string`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *TvResult) GetOriginalLanguageOk() (*string, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *TvResult) SetOriginalLanguage(v string)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *TvResult) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### GetName

`func (o *TvResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TvResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TvResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TvResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalName

`func (o *TvResult) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *TvResult) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *TvResult) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *TvResult) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetOriginCountry

`func (o *TvResult) GetOriginCountry() []string`

GetOriginCountry returns the OriginCountry field if non-nil, zero value otherwise.

### GetOriginCountryOk

`func (o *TvResult) GetOriginCountryOk() (*[]string, bool)`

GetOriginCountryOk returns a tuple with the OriginCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountry

`func (o *TvResult) SetOriginCountry(v []string)`

SetOriginCountry sets OriginCountry field to given value.

### HasOriginCountry

`func (o *TvResult) HasOriginCountry() bool`

HasOriginCountry returns a boolean if a field has been set.

### GetFirstAirDate

`func (o *TvResult) GetFirstAirDate() string`

GetFirstAirDate returns the FirstAirDate field if non-nil, zero value otherwise.

### GetFirstAirDateOk

`func (o *TvResult) GetFirstAirDateOk() (*string, bool)`

GetFirstAirDateOk returns a tuple with the FirstAirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAirDate

`func (o *TvResult) SetFirstAirDate(v string)`

SetFirstAirDate sets FirstAirDate field to given value.

### HasFirstAirDate

`func (o *TvResult) HasFirstAirDate() bool`

HasFirstAirDate returns a boolean if a field has been set.

### GetMediaInfo

`func (o *TvResult) GetMediaInfo() MediaInfo`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *TvResult) GetMediaInfoOk() (*MediaInfo, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *TvResult) SetMediaInfo(v MediaInfo)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *TvResult) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


