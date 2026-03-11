# CreditCast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**OriginalLanguage** | Pointer to **string** |  | [optional] 
**EpisodeCount** | Pointer to **float32** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**OriginCountry** | Pointer to **[]string** |  | [optional] 
**OriginalName** | Pointer to **string** |  | [optional] 
**VoteCount** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**Popularity** | Pointer to **float32** |  | [optional] 
**CreditId** | Pointer to **string** |  | [optional] 
**BackdropPath** | Pointer to **string** |  | [optional] 
**FirstAirDate** | Pointer to **string** |  | [optional] 
**VoteAverage** | Pointer to **float32** |  | [optional] 
**GenreIds** | Pointer to **[]float32** |  | [optional] 
**PosterPath** | Pointer to **string** |  | [optional] 
**OriginalTitle** | Pointer to **string** |  | [optional] 
**Video** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Adult** | Pointer to **bool** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**Character** | Pointer to **string** |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 

## Methods

### NewCreditCast

`func NewCreditCast() *CreditCast`

NewCreditCast instantiates a new CreditCast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCastWithDefaults

`func NewCreditCastWithDefaults() *CreditCast`

NewCreditCastWithDefaults instantiates a new CreditCast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditCast) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditCast) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditCast) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *CreditCast) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOriginalLanguage

`func (o *CreditCast) GetOriginalLanguage() string`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *CreditCast) GetOriginalLanguageOk() (*string, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *CreditCast) SetOriginalLanguage(v string)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *CreditCast) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### GetEpisodeCount

`func (o *CreditCast) GetEpisodeCount() float32`

GetEpisodeCount returns the EpisodeCount field if non-nil, zero value otherwise.

### GetEpisodeCountOk

`func (o *CreditCast) GetEpisodeCountOk() (*float32, bool)`

GetEpisodeCountOk returns a tuple with the EpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeCount

`func (o *CreditCast) SetEpisodeCount(v float32)`

SetEpisodeCount sets EpisodeCount field to given value.

### HasEpisodeCount

`func (o *CreditCast) HasEpisodeCount() bool`

HasEpisodeCount returns a boolean if a field has been set.

### GetOverview

`func (o *CreditCast) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *CreditCast) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *CreditCast) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *CreditCast) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetOriginCountry

`func (o *CreditCast) GetOriginCountry() []string`

GetOriginCountry returns the OriginCountry field if non-nil, zero value otherwise.

### GetOriginCountryOk

`func (o *CreditCast) GetOriginCountryOk() (*[]string, bool)`

GetOriginCountryOk returns a tuple with the OriginCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountry

`func (o *CreditCast) SetOriginCountry(v []string)`

SetOriginCountry sets OriginCountry field to given value.

### HasOriginCountry

`func (o *CreditCast) HasOriginCountry() bool`

HasOriginCountry returns a boolean if a field has been set.

### GetOriginalName

`func (o *CreditCast) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *CreditCast) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *CreditCast) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *CreditCast) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetVoteCount

`func (o *CreditCast) GetVoteCount() float32`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *CreditCast) GetVoteCountOk() (*float32, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *CreditCast) SetVoteCount(v float32)`

SetVoteCount sets VoteCount field to given value.

### HasVoteCount

`func (o *CreditCast) HasVoteCount() bool`

HasVoteCount returns a boolean if a field has been set.

### GetName

`func (o *CreditCast) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreditCast) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreditCast) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreditCast) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMediaType

`func (o *CreditCast) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *CreditCast) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *CreditCast) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *CreditCast) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetPopularity

`func (o *CreditCast) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *CreditCast) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *CreditCast) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *CreditCast) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetCreditId

`func (o *CreditCast) GetCreditId() string`

GetCreditId returns the CreditId field if non-nil, zero value otherwise.

### GetCreditIdOk

`func (o *CreditCast) GetCreditIdOk() (*string, bool)`

GetCreditIdOk returns a tuple with the CreditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditId

`func (o *CreditCast) SetCreditId(v string)`

SetCreditId sets CreditId field to given value.

### HasCreditId

`func (o *CreditCast) HasCreditId() bool`

HasCreditId returns a boolean if a field has been set.

### GetBackdropPath

`func (o *CreditCast) GetBackdropPath() string`

GetBackdropPath returns the BackdropPath field if non-nil, zero value otherwise.

### GetBackdropPathOk

`func (o *CreditCast) GetBackdropPathOk() (*string, bool)`

GetBackdropPathOk returns a tuple with the BackdropPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropPath

`func (o *CreditCast) SetBackdropPath(v string)`

SetBackdropPath sets BackdropPath field to given value.

### HasBackdropPath

`func (o *CreditCast) HasBackdropPath() bool`

HasBackdropPath returns a boolean if a field has been set.

### GetFirstAirDate

`func (o *CreditCast) GetFirstAirDate() string`

GetFirstAirDate returns the FirstAirDate field if non-nil, zero value otherwise.

### GetFirstAirDateOk

`func (o *CreditCast) GetFirstAirDateOk() (*string, bool)`

GetFirstAirDateOk returns a tuple with the FirstAirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAirDate

`func (o *CreditCast) SetFirstAirDate(v string)`

SetFirstAirDate sets FirstAirDate field to given value.

### HasFirstAirDate

`func (o *CreditCast) HasFirstAirDate() bool`

HasFirstAirDate returns a boolean if a field has been set.

### GetVoteAverage

`func (o *CreditCast) GetVoteAverage() float32`

GetVoteAverage returns the VoteAverage field if non-nil, zero value otherwise.

### GetVoteAverageOk

`func (o *CreditCast) GetVoteAverageOk() (*float32, bool)`

GetVoteAverageOk returns a tuple with the VoteAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteAverage

`func (o *CreditCast) SetVoteAverage(v float32)`

SetVoteAverage sets VoteAverage field to given value.

### HasVoteAverage

`func (o *CreditCast) HasVoteAverage() bool`

HasVoteAverage returns a boolean if a field has been set.

### GetGenreIds

`func (o *CreditCast) GetGenreIds() []float32`

GetGenreIds returns the GenreIds field if non-nil, zero value otherwise.

### GetGenreIdsOk

`func (o *CreditCast) GetGenreIdsOk() (*[]float32, bool)`

GetGenreIdsOk returns a tuple with the GenreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreIds

`func (o *CreditCast) SetGenreIds(v []float32)`

SetGenreIds sets GenreIds field to given value.

### HasGenreIds

`func (o *CreditCast) HasGenreIds() bool`

HasGenreIds returns a boolean if a field has been set.

### GetPosterPath

`func (o *CreditCast) GetPosterPath() string`

GetPosterPath returns the PosterPath field if non-nil, zero value otherwise.

### GetPosterPathOk

`func (o *CreditCast) GetPosterPathOk() (*string, bool)`

GetPosterPathOk returns a tuple with the PosterPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosterPath

`func (o *CreditCast) SetPosterPath(v string)`

SetPosterPath sets PosterPath field to given value.

### HasPosterPath

`func (o *CreditCast) HasPosterPath() bool`

HasPosterPath returns a boolean if a field has been set.

### GetOriginalTitle

`func (o *CreditCast) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *CreditCast) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *CreditCast) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *CreditCast) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### GetVideo

`func (o *CreditCast) GetVideo() bool`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CreditCast) GetVideoOk() (*bool, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CreditCast) SetVideo(v bool)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *CreditCast) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetTitle

`func (o *CreditCast) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreditCast) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreditCast) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreditCast) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAdult

`func (o *CreditCast) GetAdult() bool`

GetAdult returns the Adult field if non-nil, zero value otherwise.

### GetAdultOk

`func (o *CreditCast) GetAdultOk() (*bool, bool)`

GetAdultOk returns a tuple with the Adult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdult

`func (o *CreditCast) SetAdult(v bool)`

SetAdult sets Adult field to given value.

### HasAdult

`func (o *CreditCast) HasAdult() bool`

HasAdult returns a boolean if a field has been set.

### GetReleaseDate

`func (o *CreditCast) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *CreditCast) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *CreditCast) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *CreditCast) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetCharacter

`func (o *CreditCast) GetCharacter() string`

GetCharacter returns the Character field if non-nil, zero value otherwise.

### GetCharacterOk

`func (o *CreditCast) GetCharacterOk() (*string, bool)`

GetCharacterOk returns a tuple with the Character field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacter

`func (o *CreditCast) SetCharacter(v string)`

SetCharacter sets Character field to given value.

### HasCharacter

`func (o *CreditCast) HasCharacter() bool`

HasCharacter returns a boolean if a field has been set.

### GetMediaInfo

`func (o *CreditCast) GetMediaInfo() MediaInfo`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *CreditCast) GetMediaInfoOk() (*MediaInfo, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *CreditCast) SetMediaInfo(v MediaInfo)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *CreditCast) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


