# UserPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **float32** |  | [optional] 

## Methods

### NewUserPostRequest

`func NewUserPostRequest() *UserPostRequest`

NewUserPostRequest instantiates a new UserPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPostRequestWithDefaults

`func NewUserPostRequestWithDefaults() *UserPostRequest`

NewUserPostRequestWithDefaults instantiates a new UserPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *UserPostRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserPostRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserPostRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserPostRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPermissions

`func (o *UserPostRequest) GetPermissions() float32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserPostRequest) GetPermissionsOk() (*float32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserPostRequest) SetPermissions(v float32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserPostRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


