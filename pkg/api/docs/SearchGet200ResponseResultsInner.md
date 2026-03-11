# SearchGet200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**MediaType** | **string** |  | [default to "person"]
**Popularity** | Pointer to **float32** |  | [optional] 
**PosterPath** | Pointer to **string** |  | [optional] 
**BackdropPath** | Pointer to **string** |  | [optional] 
**VoteCount** | Pointer to **float32** |  | [optional] 
**VoteAverage** | Pointer to **float32** |  | [optional] 
**GenreIds** | Pointer to **[]float32** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**OriginalLanguage** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**OriginalTitle** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**Adult** | Pointer to **bool** |  | [optional] 
**Video** | Pointer to **bool** |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OriginalName** | Pointer to **string** |  | [optional] 
**OriginCountry** | Pointer to **[]string** |  | [optional] 
**FirstAirDate** | Pointer to **string** |  | [optional] 
**ProfilePath** | Pointer to **string** |  | [optional] 
**KnownFor** | Pointer to [**[]PersonResultKnownForInner**](PersonResultKnownForInner.md) |  | [optional] 

## Methods

### NewSearchGet200ResponseResultsInner

`func NewSearchGet200ResponseResultsInner(id float32, mediaType string, title string, ) *SearchGet200ResponseResultsInner`

NewSearchGet200ResponseResultsInner instantiates a new SearchGet200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchGet200ResponseResultsInnerWithDefaults

`func NewSearchGet200ResponseResultsInnerWithDefaults() *SearchGet200ResponseResultsInner`

NewSearchGet200ResponseResultsInnerWithDefaults instantiates a new SearchGet200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchGet200ResponseResultsInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchGet200ResponseResultsInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchGet200ResponseResultsInner) SetId(v float32)`

SetId sets Id field to given value.


### GetMediaType

`func (o *SearchGet200ResponseResultsInner) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *SearchGet200ResponseResultsInner) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *SearchGet200ResponseResultsInner) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetPopularity

`func (o *SearchGet200ResponseResultsInner) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *SearchGet200ResponseResultsInner) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *SearchGet200ResponseResultsInner) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *SearchGet200ResponseResultsInner) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetPosterPath

`func (o *SearchGet200ResponseResultsInner) GetPosterPath() string`

GetPosterPath returns the PosterPath field if non-nil, zero value otherwise.

### GetPosterPathOk

`func (o *SearchGet200ResponseResultsInner) GetPosterPathOk() (*string, bool)`

GetPosterPathOk returns a tuple with the PosterPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosterPath

`func (o *SearchGet200ResponseResultsInner) SetPosterPath(v string)`

SetPosterPath sets PosterPath field to given value.

### HasPosterPath

`func (o *SearchGet200ResponseResultsInner) HasPosterPath() bool`

HasPosterPath returns a boolean if a field has been set.

### GetBackdropPath

`func (o *SearchGet200ResponseResultsInner) GetBackdropPath() string`

GetBackdropPath returns the BackdropPath field if non-nil, zero value otherwise.

### GetBackdropPathOk

`func (o *SearchGet200ResponseResultsInner) GetBackdropPathOk() (*string, bool)`

GetBackdropPathOk returns a tuple with the BackdropPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropPath

`func (o *SearchGet200ResponseResultsInner) SetBackdropPath(v string)`

SetBackdropPath sets BackdropPath field to given value.

### HasBackdropPath

`func (o *SearchGet200ResponseResultsInner) HasBackdropPath() bool`

HasBackdropPath returns a boolean if a field has been set.

### GetVoteCount

`func (o *SearchGet200ResponseResultsInner) GetVoteCount() float32`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *SearchGet200ResponseResultsInner) GetVoteCountOk() (*float32, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *SearchGet200ResponseResultsInner) SetVoteCount(v float32)`

SetVoteCount sets VoteCount field to given value.

### HasVoteCount

`func (o *SearchGet200ResponseResultsInner) HasVoteCount() bool`

HasVoteCount returns a boolean if a field has been set.

### GetVoteAverage

`func (o *SearchGet200ResponseResultsInner) GetVoteAverage() float32`

GetVoteAverage returns the VoteAverage field if non-nil, zero value otherwise.

### GetVoteAverageOk

`func (o *SearchGet200ResponseResultsInner) GetVoteAverageOk() (*float32, bool)`

GetVoteAverageOk returns a tuple with the VoteAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteAverage

`func (o *SearchGet200ResponseResultsInner) SetVoteAverage(v float32)`

SetVoteAverage sets VoteAverage field to given value.

### HasVoteAverage

`func (o *SearchGet200ResponseResultsInner) HasVoteAverage() bool`

HasVoteAverage returns a boolean if a field has been set.

### GetGenreIds

`func (o *SearchGet200ResponseResultsInner) GetGenreIds() []float32`

GetGenreIds returns the GenreIds field if non-nil, zero value otherwise.

### GetGenreIdsOk

`func (o *SearchGet200ResponseResultsInner) GetGenreIdsOk() (*[]float32, bool)`

GetGenreIdsOk returns a tuple with the GenreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreIds

`func (o *SearchGet200ResponseResultsInner) SetGenreIds(v []float32)`

SetGenreIds sets GenreIds field to given value.

### HasGenreIds

`func (o *SearchGet200ResponseResultsInner) HasGenreIds() bool`

HasGenreIds returns a boolean if a field has been set.

### GetOverview

`func (o *SearchGet200ResponseResultsInner) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *SearchGet200ResponseResultsInner) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *SearchGet200ResponseResultsInner) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *SearchGet200ResponseResultsInner) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetOriginalLanguage

`func (o *SearchGet200ResponseResultsInner) GetOriginalLanguage() string`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *SearchGet200ResponseResultsInner) GetOriginalLanguageOk() (*string, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *SearchGet200ResponseResultsInner) SetOriginalLanguage(v string)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *SearchGet200ResponseResultsInner) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### GetTitle

`func (o *SearchGet200ResponseResultsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SearchGet200ResponseResultsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SearchGet200ResponseResultsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetOriginalTitle

`func (o *SearchGet200ResponseResultsInner) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *SearchGet200ResponseResultsInner) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *SearchGet200ResponseResultsInner) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *SearchGet200ResponseResultsInner) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### GetReleaseDate

`func (o *SearchGet200ResponseResultsInner) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SearchGet200ResponseResultsInner) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SearchGet200ResponseResultsInner) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *SearchGet200ResponseResultsInner) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetAdult

`func (o *SearchGet200ResponseResultsInner) GetAdult() bool`

GetAdult returns the Adult field if non-nil, zero value otherwise.

### GetAdultOk

`func (o *SearchGet200ResponseResultsInner) GetAdultOk() (*bool, bool)`

GetAdultOk returns a tuple with the Adult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdult

`func (o *SearchGet200ResponseResultsInner) SetAdult(v bool)`

SetAdult sets Adult field to given value.

### HasAdult

`func (o *SearchGet200ResponseResultsInner) HasAdult() bool`

HasAdult returns a boolean if a field has been set.

### GetVideo

`func (o *SearchGet200ResponseResultsInner) GetVideo() bool`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *SearchGet200ResponseResultsInner) GetVideoOk() (*bool, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *SearchGet200ResponseResultsInner) SetVideo(v bool)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *SearchGet200ResponseResultsInner) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetMediaInfo

`func (o *SearchGet200ResponseResultsInner) GetMediaInfo() MediaInfo`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *SearchGet200ResponseResultsInner) GetMediaInfoOk() (*MediaInfo, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *SearchGet200ResponseResultsInner) SetMediaInfo(v MediaInfo)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *SearchGet200ResponseResultsInner) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetName

`func (o *SearchGet200ResponseResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchGet200ResponseResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchGet200ResponseResultsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchGet200ResponseResultsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalName

`func (o *SearchGet200ResponseResultsInner) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *SearchGet200ResponseResultsInner) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *SearchGet200ResponseResultsInner) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *SearchGet200ResponseResultsInner) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetOriginCountry

`func (o *SearchGet200ResponseResultsInner) GetOriginCountry() []string`

GetOriginCountry returns the OriginCountry field if non-nil, zero value otherwise.

### GetOriginCountryOk

`func (o *SearchGet200ResponseResultsInner) GetOriginCountryOk() (*[]string, bool)`

GetOriginCountryOk returns a tuple with the OriginCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountry

`func (o *SearchGet200ResponseResultsInner) SetOriginCountry(v []string)`

SetOriginCountry sets OriginCountry field to given value.

### HasOriginCountry

`func (o *SearchGet200ResponseResultsInner) HasOriginCountry() bool`

HasOriginCountry returns a boolean if a field has been set.

### GetFirstAirDate

`func (o *SearchGet200ResponseResultsInner) GetFirstAirDate() string`

GetFirstAirDate returns the FirstAirDate field if non-nil, zero value otherwise.

### GetFirstAirDateOk

`func (o *SearchGet200ResponseResultsInner) GetFirstAirDateOk() (*string, bool)`

GetFirstAirDateOk returns a tuple with the FirstAirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAirDate

`func (o *SearchGet200ResponseResultsInner) SetFirstAirDate(v string)`

SetFirstAirDate sets FirstAirDate field to given value.

### HasFirstAirDate

`func (o *SearchGet200ResponseResultsInner) HasFirstAirDate() bool`

HasFirstAirDate returns a boolean if a field has been set.

### GetProfilePath

`func (o *SearchGet200ResponseResultsInner) GetProfilePath() string`

GetProfilePath returns the ProfilePath field if non-nil, zero value otherwise.

### GetProfilePathOk

`func (o *SearchGet200ResponseResultsInner) GetProfilePathOk() (*string, bool)`

GetProfilePathOk returns a tuple with the ProfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePath

`func (o *SearchGet200ResponseResultsInner) SetProfilePath(v string)`

SetProfilePath sets ProfilePath field to given value.

### HasProfilePath

`func (o *SearchGet200ResponseResultsInner) HasProfilePath() bool`

HasProfilePath returns a boolean if a field has been set.

### GetKnownFor

`func (o *SearchGet200ResponseResultsInner) GetKnownFor() []PersonResultKnownForInner`

GetKnownFor returns the KnownFor field if non-nil, zero value otherwise.

### GetKnownForOk

`func (o *SearchGet200ResponseResultsInner) GetKnownForOk() (*[]PersonResultKnownForInner, bool)`

GetKnownForOk returns a tuple with the KnownFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownFor

`func (o *SearchGet200ResponseResultsInner) SetKnownFor(v []PersonResultKnownForInner)`

SetKnownFor sets KnownFor field to given value.

### HasKnownFor

`func (o *SearchGet200ResponseResultsInner) HasKnownFor() bool`

HasKnownFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


