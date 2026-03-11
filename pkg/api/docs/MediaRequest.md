# MediaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | [readonly] 
**Status** | **float32** | Status of the request. 1 &#x3D; PENDING APPROVAL, 2 &#x3D; APPROVED, 3 &#x3D; DECLINED | [readonly] 
**Media** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**RequestedBy** | Pointer to [**User**](User.md) |  | [optional] 
**ModifiedBy** | Pointer to [**MediaRequestModifiedBy**](MediaRequestModifiedBy.md) |  | [optional] 
**Is4k** | Pointer to **bool** |  | [optional] 
**ServerId** | Pointer to **float32** |  | [optional] 
**ProfileId** | Pointer to **float32** |  | [optional] 
**RootFolder** | Pointer to **string** |  | [optional] 

## Methods

### NewMediaRequest

`func NewMediaRequest(id float32, status float32, ) *MediaRequest`

NewMediaRequest instantiates a new MediaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaRequestWithDefaults

`func NewMediaRequestWithDefaults() *MediaRequest`

NewMediaRequestWithDefaults instantiates a new MediaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaRequest) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaRequest) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaRequest) SetId(v float32)`

SetId sets Id field to given value.


### GetStatus

`func (o *MediaRequest) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MediaRequest) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MediaRequest) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetMedia

`func (o *MediaRequest) GetMedia() MediaInfo`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *MediaRequest) GetMediaOk() (*MediaInfo, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *MediaRequest) SetMedia(v MediaInfo)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *MediaRequest) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MediaRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MediaRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MediaRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MediaRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MediaRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MediaRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MediaRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MediaRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRequestedBy

`func (o *MediaRequest) GetRequestedBy() User`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *MediaRequest) GetRequestedByOk() (*User, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *MediaRequest) SetRequestedBy(v User)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *MediaRequest) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *MediaRequest) GetModifiedBy() MediaRequestModifiedBy`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *MediaRequest) GetModifiedByOk() (*MediaRequestModifiedBy, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *MediaRequest) SetModifiedBy(v MediaRequestModifiedBy)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *MediaRequest) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetIs4k

`func (o *MediaRequest) GetIs4k() bool`

GetIs4k returns the Is4k field if non-nil, zero value otherwise.

### GetIs4kOk

`func (o *MediaRequest) GetIs4kOk() (*bool, bool)`

GetIs4kOk returns a tuple with the Is4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs4k

`func (o *MediaRequest) SetIs4k(v bool)`

SetIs4k sets Is4k field to given value.

### HasIs4k

`func (o *MediaRequest) HasIs4k() bool`

HasIs4k returns a boolean if a field has been set.

### GetServerId

`func (o *MediaRequest) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *MediaRequest) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *MediaRequest) SetServerId(v float32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *MediaRequest) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetProfileId

`func (o *MediaRequest) GetProfileId() float32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *MediaRequest) GetProfileIdOk() (*float32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *MediaRequest) SetProfileId(v float32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *MediaRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetRootFolder

`func (o *MediaRequest) GetRootFolder() string`

GetRootFolder returns the RootFolder field if non-nil, zero value otherwise.

### GetRootFolderOk

`func (o *MediaRequest) GetRootFolderOk() (*string, bool)`

GetRootFolderOk returns a tuple with the RootFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolder

`func (o *MediaRequest) SetRootFolder(v string)`

SetRootFolder sets RootFolder field to given value.

### HasRootFolder

`func (o *MediaRequest) HasRootFolder() bool`

HasRootFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


