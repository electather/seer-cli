# Watchlist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**TmdbId** | Pointer to **float32** |  | [optional] 
**RatingKey** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**RequestedBy** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewWatchlist

`func NewWatchlist() *Watchlist`

NewWatchlist instantiates a new Watchlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchlistWithDefaults

`func NewWatchlistWithDefaults() *Watchlist`

NewWatchlistWithDefaults instantiates a new Watchlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Watchlist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Watchlist) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Watchlist) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Watchlist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTmdbId

`func (o *Watchlist) GetTmdbId() float32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *Watchlist) GetTmdbIdOk() (*float32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *Watchlist) SetTmdbId(v float32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *Watchlist) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetRatingKey

`func (o *Watchlist) GetRatingKey() string`

GetRatingKey returns the RatingKey field if non-nil, zero value otherwise.

### GetRatingKeyOk

`func (o *Watchlist) GetRatingKeyOk() (*string, bool)`

GetRatingKeyOk returns a tuple with the RatingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingKey

`func (o *Watchlist) SetRatingKey(v string)`

SetRatingKey sets RatingKey field to given value.

### HasRatingKey

`func (o *Watchlist) HasRatingKey() bool`

HasRatingKey returns a boolean if a field has been set.

### GetType

`func (o *Watchlist) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Watchlist) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Watchlist) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Watchlist) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *Watchlist) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Watchlist) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Watchlist) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Watchlist) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMedia

`func (o *Watchlist) GetMedia() MediaInfo`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *Watchlist) GetMediaOk() (*MediaInfo, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *Watchlist) SetMedia(v MediaInfo)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *Watchlist) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Watchlist) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Watchlist) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Watchlist) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Watchlist) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Watchlist) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Watchlist) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Watchlist) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Watchlist) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRequestedBy

`func (o *Watchlist) GetRequestedBy() User`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *Watchlist) GetRequestedByOk() (*User, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *Watchlist) SetRequestedBy(v User)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *Watchlist) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


