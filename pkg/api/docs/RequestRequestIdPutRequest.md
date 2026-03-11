# RequestRequestIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaType** | **string** |  | 
**Seasons** | Pointer to **[]float32** |  | [optional] 
**Is4k** | Pointer to **bool** |  | [optional] 
**ServerId** | Pointer to **float32** |  | [optional] 
**ProfileId** | Pointer to **float32** |  | [optional] 
**RootFolder** | Pointer to **string** |  | [optional] 
**LanguageProfileId** | Pointer to **float32** |  | [optional] 
**UserId** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewRequestRequestIdPutRequest

`func NewRequestRequestIdPutRequest(mediaType string, ) *RequestRequestIdPutRequest`

NewRequestRequestIdPutRequest instantiates a new RequestRequestIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestRequestIdPutRequestWithDefaults

`func NewRequestRequestIdPutRequestWithDefaults() *RequestRequestIdPutRequest`

NewRequestRequestIdPutRequestWithDefaults instantiates a new RequestRequestIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaType

`func (o *RequestRequestIdPutRequest) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *RequestRequestIdPutRequest) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *RequestRequestIdPutRequest) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetSeasons

`func (o *RequestRequestIdPutRequest) GetSeasons() []float32`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *RequestRequestIdPutRequest) GetSeasonsOk() (*[]float32, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *RequestRequestIdPutRequest) SetSeasons(v []float32)`

SetSeasons sets Seasons field to given value.

### HasSeasons

`func (o *RequestRequestIdPutRequest) HasSeasons() bool`

HasSeasons returns a boolean if a field has been set.

### GetIs4k

`func (o *RequestRequestIdPutRequest) GetIs4k() bool`

GetIs4k returns the Is4k field if non-nil, zero value otherwise.

### GetIs4kOk

`func (o *RequestRequestIdPutRequest) GetIs4kOk() (*bool, bool)`

GetIs4kOk returns a tuple with the Is4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs4k

`func (o *RequestRequestIdPutRequest) SetIs4k(v bool)`

SetIs4k sets Is4k field to given value.

### HasIs4k

`func (o *RequestRequestIdPutRequest) HasIs4k() bool`

HasIs4k returns a boolean if a field has been set.

### GetServerId

`func (o *RequestRequestIdPutRequest) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *RequestRequestIdPutRequest) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *RequestRequestIdPutRequest) SetServerId(v float32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *RequestRequestIdPutRequest) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetProfileId

`func (o *RequestRequestIdPutRequest) GetProfileId() float32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *RequestRequestIdPutRequest) GetProfileIdOk() (*float32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *RequestRequestIdPutRequest) SetProfileId(v float32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *RequestRequestIdPutRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetRootFolder

`func (o *RequestRequestIdPutRequest) GetRootFolder() string`

GetRootFolder returns the RootFolder field if non-nil, zero value otherwise.

### GetRootFolderOk

`func (o *RequestRequestIdPutRequest) GetRootFolderOk() (*string, bool)`

GetRootFolderOk returns a tuple with the RootFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolder

`func (o *RequestRequestIdPutRequest) SetRootFolder(v string)`

SetRootFolder sets RootFolder field to given value.

### HasRootFolder

`func (o *RequestRequestIdPutRequest) HasRootFolder() bool`

HasRootFolder returns a boolean if a field has been set.

### GetLanguageProfileId

`func (o *RequestRequestIdPutRequest) GetLanguageProfileId() float32`

GetLanguageProfileId returns the LanguageProfileId field if non-nil, zero value otherwise.

### GetLanguageProfileIdOk

`func (o *RequestRequestIdPutRequest) GetLanguageProfileIdOk() (*float32, bool)`

GetLanguageProfileIdOk returns a tuple with the LanguageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageProfileId

`func (o *RequestRequestIdPutRequest) SetLanguageProfileId(v float32)`

SetLanguageProfileId sets LanguageProfileId field to given value.

### HasLanguageProfileId

`func (o *RequestRequestIdPutRequest) HasLanguageProfileId() bool`

HasLanguageProfileId returns a boolean if a field has been set.

### GetUserId

`func (o *RequestRequestIdPutRequest) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RequestRequestIdPutRequest) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RequestRequestIdPutRequest) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RequestRequestIdPutRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *RequestRequestIdPutRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *RequestRequestIdPutRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


