# MediaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] [readonly] 
**TmdbId** | Pointer to **float32** |  | [optional] [readonly] 
**TvdbId** | Pointer to **NullableFloat32** |  | [optional] [readonly] 
**Status** | Pointer to **float32** | Availability of the media. 1 &#x3D; &#x60;UNKNOWN&#x60;, 2 &#x3D; &#x60;PENDING&#x60;, 3 &#x3D; &#x60;PROCESSING&#x60;, 4 &#x3D; &#x60;PARTIALLY_AVAILABLE&#x60;, 5 &#x3D; &#x60;AVAILABLE&#x60;, 6 &#x3D; &#x60;DELETED&#x60; | [optional] 
**Requests** | Pointer to [**[]MediaRequest**](MediaRequest.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMediaInfo

`func NewMediaInfo() *MediaInfo`

NewMediaInfo instantiates a new MediaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaInfoWithDefaults

`func NewMediaInfoWithDefaults() *MediaInfo`

NewMediaInfoWithDefaults instantiates a new MediaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaInfo) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaInfo) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaInfo) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *MediaInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTmdbId

`func (o *MediaInfo) GetTmdbId() float32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *MediaInfo) GetTmdbIdOk() (*float32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *MediaInfo) SetTmdbId(v float32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *MediaInfo) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetTvdbId

`func (o *MediaInfo) GetTvdbId() float32`

GetTvdbId returns the TvdbId field if non-nil, zero value otherwise.

### GetTvdbIdOk

`func (o *MediaInfo) GetTvdbIdOk() (*float32, bool)`

GetTvdbIdOk returns a tuple with the TvdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvdbId

`func (o *MediaInfo) SetTvdbId(v float32)`

SetTvdbId sets TvdbId field to given value.

### HasTvdbId

`func (o *MediaInfo) HasTvdbId() bool`

HasTvdbId returns a boolean if a field has been set.

### SetTvdbIdNil

`func (o *MediaInfo) SetTvdbIdNil(b bool)`

 SetTvdbIdNil sets the value for TvdbId to be an explicit nil

### UnsetTvdbId
`func (o *MediaInfo) UnsetTvdbId()`

UnsetTvdbId ensures that no value is present for TvdbId, not even an explicit nil
### GetStatus

`func (o *MediaInfo) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MediaInfo) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MediaInfo) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MediaInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequests

`func (o *MediaInfo) GetRequests() []MediaRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *MediaInfo) GetRequestsOk() (*[]MediaRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *MediaInfo) SetRequests(v []MediaRequest)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *MediaInfo) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MediaInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MediaInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MediaInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MediaInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MediaInfo) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MediaInfo) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MediaInfo) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MediaInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


