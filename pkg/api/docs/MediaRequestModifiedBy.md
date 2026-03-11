# MediaRequestModifiedBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Email** | **string** |  | [readonly] 
**Username** | Pointer to **string** |  | [optional] 
**PlexUsername** | Pointer to **string** |  | [optional] [readonly] 
**PlexToken** | Pointer to **string** |  | [optional] [readonly] 
**JellyfinAuthToken** | Pointer to **string** |  | [optional] [readonly] 
**UserType** | Pointer to **int32** |  | [optional] [readonly] 
**Permissions** | Pointer to **float32** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | **string** |  | [readonly] 
**UpdatedAt** | **string** |  | [readonly] 
**RequestCount** | Pointer to **float32** |  | [optional] [readonly] 

## Methods

### NewMediaRequestModifiedBy

`func NewMediaRequestModifiedBy(id int32, email string, createdAt string, updatedAt string, ) *MediaRequestModifiedBy`

NewMediaRequestModifiedBy instantiates a new MediaRequestModifiedBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaRequestModifiedByWithDefaults

`func NewMediaRequestModifiedByWithDefaults() *MediaRequestModifiedBy`

NewMediaRequestModifiedByWithDefaults instantiates a new MediaRequestModifiedBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaRequestModifiedBy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaRequestModifiedBy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaRequestModifiedBy) SetId(v int32)`

SetId sets Id field to given value.


### GetEmail

`func (o *MediaRequestModifiedBy) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MediaRequestModifiedBy) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MediaRequestModifiedBy) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUsername

`func (o *MediaRequestModifiedBy) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MediaRequestModifiedBy) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MediaRequestModifiedBy) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MediaRequestModifiedBy) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPlexUsername

`func (o *MediaRequestModifiedBy) GetPlexUsername() string`

GetPlexUsername returns the PlexUsername field if non-nil, zero value otherwise.

### GetPlexUsernameOk

`func (o *MediaRequestModifiedBy) GetPlexUsernameOk() (*string, bool)`

GetPlexUsernameOk returns a tuple with the PlexUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexUsername

`func (o *MediaRequestModifiedBy) SetPlexUsername(v string)`

SetPlexUsername sets PlexUsername field to given value.

### HasPlexUsername

`func (o *MediaRequestModifiedBy) HasPlexUsername() bool`

HasPlexUsername returns a boolean if a field has been set.

### GetPlexToken

`func (o *MediaRequestModifiedBy) GetPlexToken() string`

GetPlexToken returns the PlexToken field if non-nil, zero value otherwise.

### GetPlexTokenOk

`func (o *MediaRequestModifiedBy) GetPlexTokenOk() (*string, bool)`

GetPlexTokenOk returns a tuple with the PlexToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexToken

`func (o *MediaRequestModifiedBy) SetPlexToken(v string)`

SetPlexToken sets PlexToken field to given value.

### HasPlexToken

`func (o *MediaRequestModifiedBy) HasPlexToken() bool`

HasPlexToken returns a boolean if a field has been set.

### GetJellyfinAuthToken

`func (o *MediaRequestModifiedBy) GetJellyfinAuthToken() string`

GetJellyfinAuthToken returns the JellyfinAuthToken field if non-nil, zero value otherwise.

### GetJellyfinAuthTokenOk

`func (o *MediaRequestModifiedBy) GetJellyfinAuthTokenOk() (*string, bool)`

GetJellyfinAuthTokenOk returns a tuple with the JellyfinAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJellyfinAuthToken

`func (o *MediaRequestModifiedBy) SetJellyfinAuthToken(v string)`

SetJellyfinAuthToken sets JellyfinAuthToken field to given value.

### HasJellyfinAuthToken

`func (o *MediaRequestModifiedBy) HasJellyfinAuthToken() bool`

HasJellyfinAuthToken returns a boolean if a field has been set.

### GetUserType

`func (o *MediaRequestModifiedBy) GetUserType() int32`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *MediaRequestModifiedBy) GetUserTypeOk() (*int32, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *MediaRequestModifiedBy) SetUserType(v int32)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *MediaRequestModifiedBy) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetPermissions

`func (o *MediaRequestModifiedBy) GetPermissions() float32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MediaRequestModifiedBy) GetPermissionsOk() (*float32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MediaRequestModifiedBy) SetPermissions(v float32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *MediaRequestModifiedBy) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetAvatar

`func (o *MediaRequestModifiedBy) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *MediaRequestModifiedBy) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *MediaRequestModifiedBy) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *MediaRequestModifiedBy) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MediaRequestModifiedBy) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MediaRequestModifiedBy) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MediaRequestModifiedBy) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MediaRequestModifiedBy) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MediaRequestModifiedBy) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MediaRequestModifiedBy) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRequestCount

`func (o *MediaRequestModifiedBy) GetRequestCount() float32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *MediaRequestModifiedBy) GetRequestCountOk() (*float32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *MediaRequestModifiedBy) SetRequestCount(v float32)`

SetRequestCount sets RequestCount field to given value.

### HasRequestCount

`func (o *MediaRequestModifiedBy) HasRequestCount() bool`

HasRequestCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


