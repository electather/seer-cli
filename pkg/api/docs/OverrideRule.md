# OverrideRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Users** | Pointer to **[]float32** |  | [optional] 
**Genre** | Pointer to **NullableFloat32** |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 
**Keywords** | Pointer to **[]float32** |  | [optional] 
**ProfileId** | Pointer to **NullableFloat32** |  | [optional] 
**RootFolder** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]float32** |  | [optional] 
**RadarrServiceId** | Pointer to **NullableFloat32** |  | [optional] 
**SonarrServiceId** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOverrideRule

`func NewOverrideRule() *OverrideRule`

NewOverrideRule instantiates a new OverrideRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideRuleWithDefaults

`func NewOverrideRuleWithDefaults() *OverrideRule`

NewOverrideRuleWithDefaults instantiates a new OverrideRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OverrideRule) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OverrideRule) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OverrideRule) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *OverrideRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsers

`func (o *OverrideRule) GetUsers() []float32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *OverrideRule) GetUsersOk() (*[]float32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *OverrideRule) SetUsers(v []float32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *OverrideRule) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *OverrideRule) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *OverrideRule) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetGenre

`func (o *OverrideRule) GetGenre() float32`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *OverrideRule) GetGenreOk() (*float32, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *OverrideRule) SetGenre(v float32)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *OverrideRule) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### SetGenreNil

`func (o *OverrideRule) SetGenreNil(b bool)`

 SetGenreNil sets the value for Genre to be an explicit nil

### UnsetGenre
`func (o *OverrideRule) UnsetGenre()`

UnsetGenre ensures that no value is present for Genre, not even an explicit nil
### GetLanguage

`func (o *OverrideRule) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *OverrideRule) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *OverrideRule) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *OverrideRule) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *OverrideRule) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *OverrideRule) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetKeywords

`func (o *OverrideRule) GetKeywords() []float32`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *OverrideRule) GetKeywordsOk() (*[]float32, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *OverrideRule) SetKeywords(v []float32)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *OverrideRule) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### SetKeywordsNil

`func (o *OverrideRule) SetKeywordsNil(b bool)`

 SetKeywordsNil sets the value for Keywords to be an explicit nil

### UnsetKeywords
`func (o *OverrideRule) UnsetKeywords()`

UnsetKeywords ensures that no value is present for Keywords, not even an explicit nil
### GetProfileId

`func (o *OverrideRule) GetProfileId() float32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *OverrideRule) GetProfileIdOk() (*float32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *OverrideRule) SetProfileId(v float32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *OverrideRule) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *OverrideRule) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *OverrideRule) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetRootFolder

`func (o *OverrideRule) GetRootFolder() string`

GetRootFolder returns the RootFolder field if non-nil, zero value otherwise.

### GetRootFolderOk

`func (o *OverrideRule) GetRootFolderOk() (*string, bool)`

GetRootFolderOk returns a tuple with the RootFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolder

`func (o *OverrideRule) SetRootFolder(v string)`

SetRootFolder sets RootFolder field to given value.

### HasRootFolder

`func (o *OverrideRule) HasRootFolder() bool`

HasRootFolder returns a boolean if a field has been set.

### SetRootFolderNil

`func (o *OverrideRule) SetRootFolderNil(b bool)`

 SetRootFolderNil sets the value for RootFolder to be an explicit nil

### UnsetRootFolder
`func (o *OverrideRule) UnsetRootFolder()`

UnsetRootFolder ensures that no value is present for RootFolder, not even an explicit nil
### GetTags

`func (o *OverrideRule) GetTags() []float32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OverrideRule) GetTagsOk() (*[]float32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OverrideRule) SetTags(v []float32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OverrideRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *OverrideRule) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *OverrideRule) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetRadarrServiceId

`func (o *OverrideRule) GetRadarrServiceId() float32`

GetRadarrServiceId returns the RadarrServiceId field if non-nil, zero value otherwise.

### GetRadarrServiceIdOk

`func (o *OverrideRule) GetRadarrServiceIdOk() (*float32, bool)`

GetRadarrServiceIdOk returns a tuple with the RadarrServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadarrServiceId

`func (o *OverrideRule) SetRadarrServiceId(v float32)`

SetRadarrServiceId sets RadarrServiceId field to given value.

### HasRadarrServiceId

`func (o *OverrideRule) HasRadarrServiceId() bool`

HasRadarrServiceId returns a boolean if a field has been set.

### SetRadarrServiceIdNil

`func (o *OverrideRule) SetRadarrServiceIdNil(b bool)`

 SetRadarrServiceIdNil sets the value for RadarrServiceId to be an explicit nil

### UnsetRadarrServiceId
`func (o *OverrideRule) UnsetRadarrServiceId()`

UnsetRadarrServiceId ensures that no value is present for RadarrServiceId, not even an explicit nil
### GetSonarrServiceId

`func (o *OverrideRule) GetSonarrServiceId() float32`

GetSonarrServiceId returns the SonarrServiceId field if non-nil, zero value otherwise.

### GetSonarrServiceIdOk

`func (o *OverrideRule) GetSonarrServiceIdOk() (*float32, bool)`

GetSonarrServiceIdOk returns a tuple with the SonarrServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonarrServiceId

`func (o *OverrideRule) SetSonarrServiceId(v float32)`

SetSonarrServiceId sets SonarrServiceId field to given value.

### HasSonarrServiceId

`func (o *OverrideRule) HasSonarrServiceId() bool`

HasSonarrServiceId returns a boolean if a field has been set.

### SetSonarrServiceIdNil

`func (o *OverrideRule) SetSonarrServiceIdNil(b bool)`

 SetSonarrServiceIdNil sets the value for SonarrServiceId to be an explicit nil

### UnsetSonarrServiceId
`func (o *OverrideRule) UnsetSonarrServiceId()`

UnsetSonarrServiceId ensures that no value is present for SonarrServiceId, not even an explicit nil
### GetCreatedAt

`func (o *OverrideRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OverrideRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OverrideRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OverrideRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OverrideRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OverrideRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OverrideRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OverrideRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


