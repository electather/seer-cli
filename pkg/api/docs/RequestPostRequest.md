# RequestPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaType** | **string** |  | 
**MediaId** | **float32** |  | 
**TvdbId** | Pointer to **float32** |  | [optional] 
**Seasons** | Pointer to [**RequestPostRequestSeasons**](RequestPostRequestSeasons.md) |  | [optional] 
**Is4k** | Pointer to **bool** |  | [optional] 
**ServerId** | Pointer to **float32** |  | [optional] 
**ProfileId** | Pointer to **float32** |  | [optional] 
**RootFolder** | Pointer to **string** |  | [optional] 
**LanguageProfileId** | Pointer to **float32** |  | [optional] 
**UserId** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewRequestPostRequest

`func NewRequestPostRequest(mediaType string, mediaId float32, ) *RequestPostRequest`

NewRequestPostRequest instantiates a new RequestPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPostRequestWithDefaults

`func NewRequestPostRequestWithDefaults() *RequestPostRequest`

NewRequestPostRequestWithDefaults instantiates a new RequestPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaType

`func (o *RequestPostRequest) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *RequestPostRequest) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *RequestPostRequest) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetMediaId

`func (o *RequestPostRequest) GetMediaId() float32`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *RequestPostRequest) GetMediaIdOk() (*float32, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *RequestPostRequest) SetMediaId(v float32)`

SetMediaId sets MediaId field to given value.


### GetTvdbId

`func (o *RequestPostRequest) GetTvdbId() float32`

GetTvdbId returns the TvdbId field if non-nil, zero value otherwise.

### GetTvdbIdOk

`func (o *RequestPostRequest) GetTvdbIdOk() (*float32, bool)`

GetTvdbIdOk returns a tuple with the TvdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvdbId

`func (o *RequestPostRequest) SetTvdbId(v float32)`

SetTvdbId sets TvdbId field to given value.

### HasTvdbId

`func (o *RequestPostRequest) HasTvdbId() bool`

HasTvdbId returns a boolean if a field has been set.

### GetSeasons

`func (o *RequestPostRequest) GetSeasons() RequestPostRequestSeasons`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *RequestPostRequest) GetSeasonsOk() (*RequestPostRequestSeasons, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *RequestPostRequest) SetSeasons(v RequestPostRequestSeasons)`

SetSeasons sets Seasons field to given value.

### HasSeasons

`func (o *RequestPostRequest) HasSeasons() bool`

HasSeasons returns a boolean if a field has been set.

### GetIs4k

`func (o *RequestPostRequest) GetIs4k() bool`

GetIs4k returns the Is4k field if non-nil, zero value otherwise.

### GetIs4kOk

`func (o *RequestPostRequest) GetIs4kOk() (*bool, bool)`

GetIs4kOk returns a tuple with the Is4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs4k

`func (o *RequestPostRequest) SetIs4k(v bool)`

SetIs4k sets Is4k field to given value.

### HasIs4k

`func (o *RequestPostRequest) HasIs4k() bool`

HasIs4k returns a boolean if a field has been set.

### GetServerId

`func (o *RequestPostRequest) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *RequestPostRequest) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *RequestPostRequest) SetServerId(v float32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *RequestPostRequest) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetProfileId

`func (o *RequestPostRequest) GetProfileId() float32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *RequestPostRequest) GetProfileIdOk() (*float32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *RequestPostRequest) SetProfileId(v float32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *RequestPostRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetRootFolder

`func (o *RequestPostRequest) GetRootFolder() string`

GetRootFolder returns the RootFolder field if non-nil, zero value otherwise.

### GetRootFolderOk

`func (o *RequestPostRequest) GetRootFolderOk() (*string, bool)`

GetRootFolderOk returns a tuple with the RootFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolder

`func (o *RequestPostRequest) SetRootFolder(v string)`

SetRootFolder sets RootFolder field to given value.

### HasRootFolder

`func (o *RequestPostRequest) HasRootFolder() bool`

HasRootFolder returns a boolean if a field has been set.

### GetLanguageProfileId

`func (o *RequestPostRequest) GetLanguageProfileId() float32`

GetLanguageProfileId returns the LanguageProfileId field if non-nil, zero value otherwise.

### GetLanguageProfileIdOk

`func (o *RequestPostRequest) GetLanguageProfileIdOk() (*float32, bool)`

GetLanguageProfileIdOk returns a tuple with the LanguageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageProfileId

`func (o *RequestPostRequest) SetLanguageProfileId(v float32)`

SetLanguageProfileId sets LanguageProfileId field to given value.

### HasLanguageProfileId

`func (o *RequestPostRequest) HasLanguageProfileId() bool`

HasLanguageProfileId returns a boolean if a field has been set.

### GetUserId

`func (o *RequestPostRequest) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RequestPostRequest) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RequestPostRequest) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RequestPostRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *RequestPostRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *RequestPostRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


