# MovieDetailsReleasesResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iso31661** | Pointer to **string** |  | [optional] 
**Rating** | Pointer to **NullableString** |  | [optional] 
**ReleaseDates** | Pointer to [**[]MovieDetailsReleasesResultsInnerReleaseDatesInner**](MovieDetailsReleasesResultsInnerReleaseDatesInner.md) |  | [optional] 

## Methods

### NewMovieDetailsReleasesResultsInner

`func NewMovieDetailsReleasesResultsInner() *MovieDetailsReleasesResultsInner`

NewMovieDetailsReleasesResultsInner instantiates a new MovieDetailsReleasesResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieDetailsReleasesResultsInnerWithDefaults

`func NewMovieDetailsReleasesResultsInnerWithDefaults() *MovieDetailsReleasesResultsInner`

NewMovieDetailsReleasesResultsInnerWithDefaults instantiates a new MovieDetailsReleasesResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIso31661

`func (o *MovieDetailsReleasesResultsInner) GetIso31661() string`

GetIso31661 returns the Iso31661 field if non-nil, zero value otherwise.

### GetIso31661Ok

`func (o *MovieDetailsReleasesResultsInner) GetIso31661Ok() (*string, bool)`

GetIso31661Ok returns a tuple with the Iso31661 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso31661

`func (o *MovieDetailsReleasesResultsInner) SetIso31661(v string)`

SetIso31661 sets Iso31661 field to given value.

### HasIso31661

`func (o *MovieDetailsReleasesResultsInner) HasIso31661() bool`

HasIso31661 returns a boolean if a field has been set.

### GetRating

`func (o *MovieDetailsReleasesResultsInner) GetRating() string`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *MovieDetailsReleasesResultsInner) GetRatingOk() (*string, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *MovieDetailsReleasesResultsInner) SetRating(v string)`

SetRating sets Rating field to given value.

### HasRating

`func (o *MovieDetailsReleasesResultsInner) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *MovieDetailsReleasesResultsInner) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *MovieDetailsReleasesResultsInner) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetReleaseDates

`func (o *MovieDetailsReleasesResultsInner) GetReleaseDates() []MovieDetailsReleasesResultsInnerReleaseDatesInner`

GetReleaseDates returns the ReleaseDates field if non-nil, zero value otherwise.

### GetReleaseDatesOk

`func (o *MovieDetailsReleasesResultsInner) GetReleaseDatesOk() (*[]MovieDetailsReleasesResultsInnerReleaseDatesInner, bool)`

GetReleaseDatesOk returns a tuple with the ReleaseDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDates

`func (o *MovieDetailsReleasesResultsInner) SetReleaseDates(v []MovieDetailsReleasesResultsInnerReleaseDatesInner)`

SetReleaseDates sets ReleaseDates field to given value.

### HasReleaseDates

`func (o *MovieDetailsReleasesResultsInner) HasReleaseDates() bool`

HasReleaseDates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


