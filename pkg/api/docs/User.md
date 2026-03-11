# User

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

### NewUser

`func NewUser(id int32, email string, createdAt string, updatedAt string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPlexUsername

`func (o *User) GetPlexUsername() string`

GetPlexUsername returns the PlexUsername field if non-nil, zero value otherwise.

### GetPlexUsernameOk

`func (o *User) GetPlexUsernameOk() (*string, bool)`

GetPlexUsernameOk returns a tuple with the PlexUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexUsername

`func (o *User) SetPlexUsername(v string)`

SetPlexUsername sets PlexUsername field to given value.

### HasPlexUsername

`func (o *User) HasPlexUsername() bool`

HasPlexUsername returns a boolean if a field has been set.

### GetPlexToken

`func (o *User) GetPlexToken() string`

GetPlexToken returns the PlexToken field if non-nil, zero value otherwise.

### GetPlexTokenOk

`func (o *User) GetPlexTokenOk() (*string, bool)`

GetPlexTokenOk returns a tuple with the PlexToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexToken

`func (o *User) SetPlexToken(v string)`

SetPlexToken sets PlexToken field to given value.

### HasPlexToken

`func (o *User) HasPlexToken() bool`

HasPlexToken returns a boolean if a field has been set.

### GetJellyfinAuthToken

`func (o *User) GetJellyfinAuthToken() string`

GetJellyfinAuthToken returns the JellyfinAuthToken field if non-nil, zero value otherwise.

### GetJellyfinAuthTokenOk

`func (o *User) GetJellyfinAuthTokenOk() (*string, bool)`

GetJellyfinAuthTokenOk returns a tuple with the JellyfinAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJellyfinAuthToken

`func (o *User) SetJellyfinAuthToken(v string)`

SetJellyfinAuthToken sets JellyfinAuthToken field to given value.

### HasJellyfinAuthToken

`func (o *User) HasJellyfinAuthToken() bool`

HasJellyfinAuthToken returns a boolean if a field has been set.

### GetUserType

`func (o *User) GetUserType() int32`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *User) GetUserTypeOk() (*int32, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *User) SetUserType(v int32)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *User) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetPermissions

`func (o *User) GetPermissions() float32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *User) GetPermissionsOk() (*float32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *User) SetPermissions(v float32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *User) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetAvatar

`func (o *User) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *User) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *User) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *User) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *User) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRequestCount

`func (o *User) GetRequestCount() float32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *User) GetRequestCountOk() (*float32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *User) SetRequestCount(v float32)`

SetRequestCount sets RequestCount field to given value.

### HasRequestCount

`func (o *User) HasRequestCount() bool`

HasRequestCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


