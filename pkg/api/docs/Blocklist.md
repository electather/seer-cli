# Blocklist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TmdbId** | Pointer to **float32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 
**UserId** | Pointer to **float32** |  | [optional] 

## Methods

### NewBlocklist

`func NewBlocklist() *Blocklist`

NewBlocklist instantiates a new Blocklist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlocklistWithDefaults

`func NewBlocklistWithDefaults() *Blocklist`

NewBlocklistWithDefaults instantiates a new Blocklist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTmdbId

`func (o *Blocklist) GetTmdbId() float32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *Blocklist) GetTmdbIdOk() (*float32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *Blocklist) SetTmdbId(v float32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *Blocklist) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetTitle

`func (o *Blocklist) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Blocklist) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Blocklist) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Blocklist) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMedia

`func (o *Blocklist) GetMedia() MediaInfo`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *Blocklist) GetMediaOk() (*MediaInfo, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *Blocklist) SetMedia(v MediaInfo)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *Blocklist) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetUserId

`func (o *Blocklist) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Blocklist) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Blocklist) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Blocklist) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


