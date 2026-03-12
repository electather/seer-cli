# UserUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **float32** |  | [optional] 

## Methods

### NewUserUpdatePayload

`func NewUserUpdatePayload() *UserUpdatePayload`

NewUserUpdatePayload instantiates a new UserUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdatePayloadWithDefaults

`func NewUserUpdatePayloadWithDefaults() *UserUpdatePayload`

NewUserUpdatePayloadWithDefaults instantiates a new UserUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserUpdatePayload) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserUpdatePayload) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserUpdatePayload) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserUpdatePayload) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *UserUpdatePayload) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserUpdatePayload) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserUpdatePayload) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserUpdatePayload) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPermissions

`func (o *UserUpdatePayload) GetPermissions() float32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserUpdatePayload) GetPermissionsOk() (*float32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserUpdatePayload) SetPermissions(v float32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserUpdatePayload) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


