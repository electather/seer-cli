# NotificationEmailSettingsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailFrom** | Pointer to **string** |  | [optional] 
**SenderName** | Pointer to **string** |  | [optional] 
**SmtpHost** | Pointer to **string** |  | [optional] 
**SmtpPort** | Pointer to **float32** |  | [optional] 
**Secure** | Pointer to **bool** |  | [optional] 
**IgnoreTls** | Pointer to **bool** |  | [optional] 
**RequireTls** | Pointer to **bool** |  | [optional] 
**AuthUser** | Pointer to **NullableString** |  | [optional] 
**AuthPass** | Pointer to **NullableString** |  | [optional] 
**AllowSelfSigned** | Pointer to **bool** |  | [optional] 

## Methods

### NewNotificationEmailSettingsOptions

`func NewNotificationEmailSettingsOptions() *NotificationEmailSettingsOptions`

NewNotificationEmailSettingsOptions instantiates a new NotificationEmailSettingsOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEmailSettingsOptionsWithDefaults

`func NewNotificationEmailSettingsOptionsWithDefaults() *NotificationEmailSettingsOptions`

NewNotificationEmailSettingsOptionsWithDefaults instantiates a new NotificationEmailSettingsOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailFrom

`func (o *NotificationEmailSettingsOptions) GetEmailFrom() string`

GetEmailFrom returns the EmailFrom field if non-nil, zero value otherwise.

### GetEmailFromOk

`func (o *NotificationEmailSettingsOptions) GetEmailFromOk() (*string, bool)`

GetEmailFromOk returns a tuple with the EmailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFrom

`func (o *NotificationEmailSettingsOptions) SetEmailFrom(v string)`

SetEmailFrom sets EmailFrom field to given value.

### HasEmailFrom

`func (o *NotificationEmailSettingsOptions) HasEmailFrom() bool`

HasEmailFrom returns a boolean if a field has been set.

### GetSenderName

`func (o *NotificationEmailSettingsOptions) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *NotificationEmailSettingsOptions) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *NotificationEmailSettingsOptions) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *NotificationEmailSettingsOptions) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### GetSmtpHost

`func (o *NotificationEmailSettingsOptions) GetSmtpHost() string`

GetSmtpHost returns the SmtpHost field if non-nil, zero value otherwise.

### GetSmtpHostOk

`func (o *NotificationEmailSettingsOptions) GetSmtpHostOk() (*string, bool)`

GetSmtpHostOk returns a tuple with the SmtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpHost

`func (o *NotificationEmailSettingsOptions) SetSmtpHost(v string)`

SetSmtpHost sets SmtpHost field to given value.

### HasSmtpHost

`func (o *NotificationEmailSettingsOptions) HasSmtpHost() bool`

HasSmtpHost returns a boolean if a field has been set.

### GetSmtpPort

`func (o *NotificationEmailSettingsOptions) GetSmtpPort() float32`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *NotificationEmailSettingsOptions) GetSmtpPortOk() (*float32, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *NotificationEmailSettingsOptions) SetSmtpPort(v float32)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *NotificationEmailSettingsOptions) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSecure

`func (o *NotificationEmailSettingsOptions) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *NotificationEmailSettingsOptions) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *NotificationEmailSettingsOptions) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *NotificationEmailSettingsOptions) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetIgnoreTls

`func (o *NotificationEmailSettingsOptions) GetIgnoreTls() bool`

GetIgnoreTls returns the IgnoreTls field if non-nil, zero value otherwise.

### GetIgnoreTlsOk

`func (o *NotificationEmailSettingsOptions) GetIgnoreTlsOk() (*bool, bool)`

GetIgnoreTlsOk returns a tuple with the IgnoreTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTls

`func (o *NotificationEmailSettingsOptions) SetIgnoreTls(v bool)`

SetIgnoreTls sets IgnoreTls field to given value.

### HasIgnoreTls

`func (o *NotificationEmailSettingsOptions) HasIgnoreTls() bool`

HasIgnoreTls returns a boolean if a field has been set.

### GetRequireTls

`func (o *NotificationEmailSettingsOptions) GetRequireTls() bool`

GetRequireTls returns the RequireTls field if non-nil, zero value otherwise.

### GetRequireTlsOk

`func (o *NotificationEmailSettingsOptions) GetRequireTlsOk() (*bool, bool)`

GetRequireTlsOk returns a tuple with the RequireTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTls

`func (o *NotificationEmailSettingsOptions) SetRequireTls(v bool)`

SetRequireTls sets RequireTls field to given value.

### HasRequireTls

`func (o *NotificationEmailSettingsOptions) HasRequireTls() bool`

HasRequireTls returns a boolean if a field has been set.

### GetAuthUser

`func (o *NotificationEmailSettingsOptions) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *NotificationEmailSettingsOptions) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *NotificationEmailSettingsOptions) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *NotificationEmailSettingsOptions) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

### SetAuthUserNil

`func (o *NotificationEmailSettingsOptions) SetAuthUserNil(b bool)`

 SetAuthUserNil sets the value for AuthUser to be an explicit nil

### UnsetAuthUser
`func (o *NotificationEmailSettingsOptions) UnsetAuthUser()`

UnsetAuthUser ensures that no value is present for AuthUser, not even an explicit nil
### GetAuthPass

`func (o *NotificationEmailSettingsOptions) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *NotificationEmailSettingsOptions) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *NotificationEmailSettingsOptions) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.

### HasAuthPass

`func (o *NotificationEmailSettingsOptions) HasAuthPass() bool`

HasAuthPass returns a boolean if a field has been set.

### SetAuthPassNil

`func (o *NotificationEmailSettingsOptions) SetAuthPassNil(b bool)`

 SetAuthPassNil sets the value for AuthPass to be an explicit nil

### UnsetAuthPass
`func (o *NotificationEmailSettingsOptions) UnsetAuthPass()`

UnsetAuthPass ensures that no value is present for AuthPass, not even an explicit nil
### GetAllowSelfSigned

`func (o *NotificationEmailSettingsOptions) GetAllowSelfSigned() bool`

GetAllowSelfSigned returns the AllowSelfSigned field if non-nil, zero value otherwise.

### GetAllowSelfSignedOk

`func (o *NotificationEmailSettingsOptions) GetAllowSelfSignedOk() (*bool, bool)`

GetAllowSelfSignedOk returns a tuple with the AllowSelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSelfSigned

`func (o *NotificationEmailSettingsOptions) SetAllowSelfSigned(v bool)`

SetAllowSelfSigned sets AllowSelfSigned field to given value.

### HasAllowSelfSigned

`func (o *NotificationEmailSettingsOptions) HasAllowSelfSigned() bool`

HasAllowSelfSigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


