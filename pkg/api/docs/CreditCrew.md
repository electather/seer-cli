# CreditCrew

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
**Department** | Pointer to **string** |  | [optional] 
**Job** | Pointer to **string** |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 

## Methods

### NewCreditCrew

`func NewCreditCrew() *CreditCrew`

NewCreditCrew instantiates a new CreditCrew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCrewWithDefaults

`func NewCreditCrewWithDefaults() *CreditCrew`

NewCreditCrewWithDefaults instantiates a new CreditCrew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditCrew) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditCrew) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditCrew) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *CreditCrew) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOriginalLanguage

`func (o *CreditCrew) GetOriginalLanguage() string`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *CreditCrew) GetOriginalLanguageOk() (*string, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *CreditCrew) SetOriginalLanguage(v string)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *CreditCrew) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### GetEpisodeCount

`func (o *CreditCrew) GetEpisodeCount() float32`

GetEpisodeCount returns the EpisodeCount field if non-nil, zero value otherwise.

### GetEpisodeCountOk

`func (o *CreditCrew) GetEpisodeCountOk() (*float32, bool)`

GetEpisodeCountOk returns a tuple with the EpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeCount

`func (o *CreditCrew) SetEpisodeCount(v float32)`

SetEpisodeCount sets EpisodeCount field to given value.

### HasEpisodeCount

`func (o *CreditCrew) HasEpisodeCount() bool`

HasEpisodeCount returns a boolean if a field has been set.

### GetOverview

`func (o *CreditCrew) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *CreditCrew) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *CreditCrew) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *CreditCrew) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetOriginCountry

`func (o *CreditCrew) GetOriginCountry() []string`

GetOriginCountry returns the OriginCountry field if non-nil, zero value otherwise.

### GetOriginCountryOk

`func (o *CreditCrew) GetOriginCountryOk() (*[]string, bool)`

GetOriginCountryOk returns a tuple with the OriginCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountry

`func (o *CreditCrew) SetOriginCountry(v []string)`

SetOriginCountry sets OriginCountry field to given value.

### HasOriginCountry

`func (o *CreditCrew) HasOriginCountry() bool`

HasOriginCountry returns a boolean if a field has been set.

### GetOriginalName

`func (o *CreditCrew) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *CreditCrew) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *CreditCrew) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *CreditCrew) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetVoteCount

`func (o *CreditCrew) GetVoteCount() float32`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *CreditCrew) GetVoteCountOk() (*float32, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *CreditCrew) SetVoteCount(v float32)`

SetVoteCount sets VoteCount field to given value.

### HasVoteCount

`func (o *CreditCrew) HasVoteCount() bool`

HasVoteCount returns a boolean if a field has been set.

### GetName

`func (o *CreditCrew) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreditCrew) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreditCrew) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreditCrew) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMediaType

`func (o *CreditCrew) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *CreditCrew) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *CreditCrew) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *CreditCrew) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetPopularity

`func (o *CreditCrew) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *CreditCrew) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *CreditCrew) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *CreditCrew) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetCreditId

`func (o *CreditCrew) GetCreditId() string`

GetCreditId returns the CreditId field if non-nil, zero value otherwise.

### GetCreditIdOk

`func (o *CreditCrew) GetCreditIdOk() (*string, bool)`

GetCreditIdOk returns a tuple with the CreditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditId

`func (o *CreditCrew) SetCreditId(v string)`

SetCreditId sets CreditId field to given value.

### HasCreditId

`func (o *CreditCrew) HasCreditId() bool`

HasCreditId returns a boolean if a field has been set.

### GetBackdropPath

`func (o *CreditCrew) GetBackdropPath() string`

GetBackdropPath returns the BackdropPath field if non-nil, zero value otherwise.

### GetBackdropPathOk

`func (o *CreditCrew) GetBackdropPathOk() (*string, bool)`

GetBackdropPathOk returns a tuple with the BackdropPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropPath

`func (o *CreditCrew) SetBackdropPath(v string)`

SetBackdropPath sets BackdropPath field to given value.

### HasBackdropPath

`func (o *CreditCrew) HasBackdropPath() bool`

HasBackdropPath returns a boolean if a field has been set.

### GetFirstAirDate

`func (o *CreditCrew) GetFirstAirDate() string`

GetFirstAirDate returns the FirstAirDate field if non-nil, zero value otherwise.

### GetFirstAirDateOk

`func (o *CreditCrew) GetFirstAirDateOk() (*string, bool)`

GetFirstAirDateOk returns a tuple with the FirstAirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAirDate

`func (o *CreditCrew) SetFirstAirDate(v string)`

SetFirstAirDate sets FirstAirDate field to given value.

### HasFirstAirDate

`func (o *CreditCrew) HasFirstAirDate() bool`

HasFirstAirDate returns a boolean if a field has been set.

### GetVoteAverage

`func (o *CreditCrew) GetVoteAverage() float32`

GetVoteAverage returns the VoteAverage field if non-nil, zero value otherwise.

### GetVoteAverageOk

`func (o *CreditCrew) GetVoteAverageOk() (*float32, bool)`

GetVoteAverageOk returns a tuple with the VoteAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteAverage

`func (o *CreditCrew) SetVoteAverage(v float32)`

SetVoteAverage sets VoteAverage field to given value.

### HasVoteAverage

`func (o *CreditCrew) HasVoteAverage() bool`

HasVoteAverage returns a boolean if a field has been set.

### GetGenreIds

`func (o *CreditCrew) GetGenreIds() []float32`

GetGenreIds returns the GenreIds field if non-nil, zero value otherwise.

### GetGenreIdsOk

`func (o *CreditCrew) GetGenreIdsOk() (*[]float32, bool)`

GetGenreIdsOk returns a tuple with the GenreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreIds

`func (o *CreditCrew) SetGenreIds(v []float32)`

SetGenreIds sets GenreIds field to given value.

### HasGenreIds

`func (o *CreditCrew) HasGenreIds() bool`

HasGenreIds returns a boolean if a field has been set.

### GetPosterPath

`func (o *CreditCrew) GetPosterPath() string`

GetPosterPath returns the PosterPath field if non-nil, zero value otherwise.

### GetPosterPathOk

`func (o *CreditCrew) GetPosterPathOk() (*string, bool)`

GetPosterPathOk returns a tuple with the PosterPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosterPath

`func (o *CreditCrew) SetPosterPath(v string)`

SetPosterPath sets PosterPath field to given value.

### HasPosterPath

`func (o *CreditCrew) HasPosterPath() bool`

HasPosterPath returns a boolean if a field has been set.

### GetOriginalTitle

`func (o *CreditCrew) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *CreditCrew) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *CreditCrew) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *CreditCrew) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### GetVideo

`func (o *CreditCrew) GetVideo() bool`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CreditCrew) GetVideoOk() (*bool, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CreditCrew) SetVideo(v bool)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *CreditCrew) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetTitle

`func (o *CreditCrew) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreditCrew) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreditCrew) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreditCrew) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAdult

`func (o *CreditCrew) GetAdult() bool`

GetAdult returns the Adult field if non-nil, zero value otherwise.

### GetAdultOk

`func (o *CreditCrew) GetAdultOk() (*bool, bool)`

GetAdultOk returns a tuple with the Adult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdult

`func (o *CreditCrew) SetAdult(v bool)`

SetAdult sets Adult field to given value.

### HasAdult

`func (o *CreditCrew) HasAdult() bool`

HasAdult returns a boolean if a field has been set.

### GetReleaseDate

`func (o *CreditCrew) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *CreditCrew) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *CreditCrew) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *CreditCrew) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetDepartment

`func (o *CreditCrew) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *CreditCrew) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *CreditCrew) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *CreditCrew) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetJob

`func (o *CreditCrew) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CreditCrew) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CreditCrew) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *CreditCrew) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetMediaInfo

`func (o *CreditCrew) GetMediaInfo() MediaInfo`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *CreditCrew) GetMediaInfoOk() (*MediaInfo, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *CreditCrew) SetMediaInfo(v MediaInfo)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *CreditCrew) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


