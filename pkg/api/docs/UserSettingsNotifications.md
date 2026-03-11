# UserSettingsNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationTypes** | Pointer to [**NotificationAgentTypes**](NotificationAgentTypes.md) |  | [optional] 
**EmailEnabled** | Pointer to **bool** |  | [optional] 
**PgpKey** | Pointer to **NullableString** |  | [optional] 
**DiscordEnabled** | Pointer to **bool** |  | [optional] 
**DiscordEnabledTypes** | Pointer to **NullableFloat32** |  | [optional] 
**DiscordId** | Pointer to **NullableString** |  | [optional] 
**PushbulletAccessToken** | Pointer to **NullableString** |  | [optional] 
**PushoverApplicationToken** | Pointer to **NullableString** |  | [optional] 
**PushoverUserKey** | Pointer to **NullableString** |  | [optional] 
**PushoverSound** | Pointer to **NullableString** |  | [optional] 
**TelegramEnabled** | Pointer to **bool** |  | [optional] 
**TelegramBotUsername** | Pointer to **NullableString** |  | [optional] 
**TelegramChatId** | Pointer to **NullableString** |  | [optional] 
**TelegramMessageThreadId** | Pointer to **NullableString** |  | [optional] 
**TelegramSendSilently** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUserSettingsNotifications

`func NewUserSettingsNotifications() *UserSettingsNotifications`

NewUserSettingsNotifications instantiates a new UserSettingsNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsNotificationsWithDefaults

`func NewUserSettingsNotificationsWithDefaults() *UserSettingsNotifications`

NewUserSettingsNotificationsWithDefaults instantiates a new UserSettingsNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationTypes

`func (o *UserSettingsNotifications) GetNotificationTypes() NotificationAgentTypes`

GetNotificationTypes returns the NotificationTypes field if non-nil, zero value otherwise.

### GetNotificationTypesOk

`func (o *UserSettingsNotifications) GetNotificationTypesOk() (*NotificationAgentTypes, bool)`

GetNotificationTypesOk returns a tuple with the NotificationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTypes

`func (o *UserSettingsNotifications) SetNotificationTypes(v NotificationAgentTypes)`

SetNotificationTypes sets NotificationTypes field to given value.

### HasNotificationTypes

`func (o *UserSettingsNotifications) HasNotificationTypes() bool`

HasNotificationTypes returns a boolean if a field has been set.

### GetEmailEnabled

`func (o *UserSettingsNotifications) GetEmailEnabled() bool`

GetEmailEnabled returns the EmailEnabled field if non-nil, zero value otherwise.

### GetEmailEnabledOk

`func (o *UserSettingsNotifications) GetEmailEnabledOk() (*bool, bool)`

GetEmailEnabledOk returns a tuple with the EmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailEnabled

`func (o *UserSettingsNotifications) SetEmailEnabled(v bool)`

SetEmailEnabled sets EmailEnabled field to given value.

### HasEmailEnabled

`func (o *UserSettingsNotifications) HasEmailEnabled() bool`

HasEmailEnabled returns a boolean if a field has been set.

### GetPgpKey

`func (o *UserSettingsNotifications) GetPgpKey() string`

GetPgpKey returns the PgpKey field if non-nil, zero value otherwise.

### GetPgpKeyOk

`func (o *UserSettingsNotifications) GetPgpKeyOk() (*string, bool)`

GetPgpKeyOk returns a tuple with the PgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpKey

`func (o *UserSettingsNotifications) SetPgpKey(v string)`

SetPgpKey sets PgpKey field to given value.

### HasPgpKey

`func (o *UserSettingsNotifications) HasPgpKey() bool`

HasPgpKey returns a boolean if a field has been set.

### SetPgpKeyNil

`func (o *UserSettingsNotifications) SetPgpKeyNil(b bool)`

 SetPgpKeyNil sets the value for PgpKey to be an explicit nil

### UnsetPgpKey
`func (o *UserSettingsNotifications) UnsetPgpKey()`

UnsetPgpKey ensures that no value is present for PgpKey, not even an explicit nil
### GetDiscordEnabled

`func (o *UserSettingsNotifications) GetDiscordEnabled() bool`

GetDiscordEnabled returns the DiscordEnabled field if non-nil, zero value otherwise.

### GetDiscordEnabledOk

`func (o *UserSettingsNotifications) GetDiscordEnabledOk() (*bool, bool)`

GetDiscordEnabledOk returns a tuple with the DiscordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordEnabled

`func (o *UserSettingsNotifications) SetDiscordEnabled(v bool)`

SetDiscordEnabled sets DiscordEnabled field to given value.

### HasDiscordEnabled

`func (o *UserSettingsNotifications) HasDiscordEnabled() bool`

HasDiscordEnabled returns a boolean if a field has been set.

### GetDiscordEnabledTypes

`func (o *UserSettingsNotifications) GetDiscordEnabledTypes() float32`

GetDiscordEnabledTypes returns the DiscordEnabledTypes field if non-nil, zero value otherwise.

### GetDiscordEnabledTypesOk

`func (o *UserSettingsNotifications) GetDiscordEnabledTypesOk() (*float32, bool)`

GetDiscordEnabledTypesOk returns a tuple with the DiscordEnabledTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordEnabledTypes

`func (o *UserSettingsNotifications) SetDiscordEnabledTypes(v float32)`

SetDiscordEnabledTypes sets DiscordEnabledTypes field to given value.

### HasDiscordEnabledTypes

`func (o *UserSettingsNotifications) HasDiscordEnabledTypes() bool`

HasDiscordEnabledTypes returns a boolean if a field has been set.

### SetDiscordEnabledTypesNil

`func (o *UserSettingsNotifications) SetDiscordEnabledTypesNil(b bool)`

 SetDiscordEnabledTypesNil sets the value for DiscordEnabledTypes to be an explicit nil

### UnsetDiscordEnabledTypes
`func (o *UserSettingsNotifications) UnsetDiscordEnabledTypes()`

UnsetDiscordEnabledTypes ensures that no value is present for DiscordEnabledTypes, not even an explicit nil
### GetDiscordId

`func (o *UserSettingsNotifications) GetDiscordId() string`

GetDiscordId returns the DiscordId field if non-nil, zero value otherwise.

### GetDiscordIdOk

`func (o *UserSettingsNotifications) GetDiscordIdOk() (*string, bool)`

GetDiscordIdOk returns a tuple with the DiscordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordId

`func (o *UserSettingsNotifications) SetDiscordId(v string)`

SetDiscordId sets DiscordId field to given value.

### HasDiscordId

`func (o *UserSettingsNotifications) HasDiscordId() bool`

HasDiscordId returns a boolean if a field has been set.

### SetDiscordIdNil

`func (o *UserSettingsNotifications) SetDiscordIdNil(b bool)`

 SetDiscordIdNil sets the value for DiscordId to be an explicit nil

### UnsetDiscordId
`func (o *UserSettingsNotifications) UnsetDiscordId()`

UnsetDiscordId ensures that no value is present for DiscordId, not even an explicit nil
### GetPushbulletAccessToken

`func (o *UserSettingsNotifications) GetPushbulletAccessToken() string`

GetPushbulletAccessToken returns the PushbulletAccessToken field if non-nil, zero value otherwise.

### GetPushbulletAccessTokenOk

`func (o *UserSettingsNotifications) GetPushbulletAccessTokenOk() (*string, bool)`

GetPushbulletAccessTokenOk returns a tuple with the PushbulletAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushbulletAccessToken

`func (o *UserSettingsNotifications) SetPushbulletAccessToken(v string)`

SetPushbulletAccessToken sets PushbulletAccessToken field to given value.

### HasPushbulletAccessToken

`func (o *UserSettingsNotifications) HasPushbulletAccessToken() bool`

HasPushbulletAccessToken returns a boolean if a field has been set.

### SetPushbulletAccessTokenNil

`func (o *UserSettingsNotifications) SetPushbulletAccessTokenNil(b bool)`

 SetPushbulletAccessTokenNil sets the value for PushbulletAccessToken to be an explicit nil

### UnsetPushbulletAccessToken
`func (o *UserSettingsNotifications) UnsetPushbulletAccessToken()`

UnsetPushbulletAccessToken ensures that no value is present for PushbulletAccessToken, not even an explicit nil
### GetPushoverApplicationToken

`func (o *UserSettingsNotifications) GetPushoverApplicationToken() string`

GetPushoverApplicationToken returns the PushoverApplicationToken field if non-nil, zero value otherwise.

### GetPushoverApplicationTokenOk

`func (o *UserSettingsNotifications) GetPushoverApplicationTokenOk() (*string, bool)`

GetPushoverApplicationTokenOk returns a tuple with the PushoverApplicationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushoverApplicationToken

`func (o *UserSettingsNotifications) SetPushoverApplicationToken(v string)`

SetPushoverApplicationToken sets PushoverApplicationToken field to given value.

### HasPushoverApplicationToken

`func (o *UserSettingsNotifications) HasPushoverApplicationToken() bool`

HasPushoverApplicationToken returns a boolean if a field has been set.

### SetPushoverApplicationTokenNil

`func (o *UserSettingsNotifications) SetPushoverApplicationTokenNil(b bool)`

 SetPushoverApplicationTokenNil sets the value for PushoverApplicationToken to be an explicit nil

### UnsetPushoverApplicationToken
`func (o *UserSettingsNotifications) UnsetPushoverApplicationToken()`

UnsetPushoverApplicationToken ensures that no value is present for PushoverApplicationToken, not even an explicit nil
### GetPushoverUserKey

`func (o *UserSettingsNotifications) GetPushoverUserKey() string`

GetPushoverUserKey returns the PushoverUserKey field if non-nil, zero value otherwise.

### GetPushoverUserKeyOk

`func (o *UserSettingsNotifications) GetPushoverUserKeyOk() (*string, bool)`

GetPushoverUserKeyOk returns a tuple with the PushoverUserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushoverUserKey

`func (o *UserSettingsNotifications) SetPushoverUserKey(v string)`

SetPushoverUserKey sets PushoverUserKey field to given value.

### HasPushoverUserKey

`func (o *UserSettingsNotifications) HasPushoverUserKey() bool`

HasPushoverUserKey returns a boolean if a field has been set.

### SetPushoverUserKeyNil

`func (o *UserSettingsNotifications) SetPushoverUserKeyNil(b bool)`

 SetPushoverUserKeyNil sets the value for PushoverUserKey to be an explicit nil

### UnsetPushoverUserKey
`func (o *UserSettingsNotifications) UnsetPushoverUserKey()`

UnsetPushoverUserKey ensures that no value is present for PushoverUserKey, not even an explicit nil
### GetPushoverSound

`func (o *UserSettingsNotifications) GetPushoverSound() string`

GetPushoverSound returns the PushoverSound field if non-nil, zero value otherwise.

### GetPushoverSoundOk

`func (o *UserSettingsNotifications) GetPushoverSoundOk() (*string, bool)`

GetPushoverSoundOk returns a tuple with the PushoverSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushoverSound

`func (o *UserSettingsNotifications) SetPushoverSound(v string)`

SetPushoverSound sets PushoverSound field to given value.

### HasPushoverSound

`func (o *UserSettingsNotifications) HasPushoverSound() bool`

HasPushoverSound returns a boolean if a field has been set.

### SetPushoverSoundNil

`func (o *UserSettingsNotifications) SetPushoverSoundNil(b bool)`

 SetPushoverSoundNil sets the value for PushoverSound to be an explicit nil

### UnsetPushoverSound
`func (o *UserSettingsNotifications) UnsetPushoverSound()`

UnsetPushoverSound ensures that no value is present for PushoverSound, not even an explicit nil
### GetTelegramEnabled

`func (o *UserSettingsNotifications) GetTelegramEnabled() bool`

GetTelegramEnabled returns the TelegramEnabled field if non-nil, zero value otherwise.

### GetTelegramEnabledOk

`func (o *UserSettingsNotifications) GetTelegramEnabledOk() (*bool, bool)`

GetTelegramEnabledOk returns a tuple with the TelegramEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramEnabled

`func (o *UserSettingsNotifications) SetTelegramEnabled(v bool)`

SetTelegramEnabled sets TelegramEnabled field to given value.

### HasTelegramEnabled

`func (o *UserSettingsNotifications) HasTelegramEnabled() bool`

HasTelegramEnabled returns a boolean if a field has been set.

### GetTelegramBotUsername

`func (o *UserSettingsNotifications) GetTelegramBotUsername() string`

GetTelegramBotUsername returns the TelegramBotUsername field if non-nil, zero value otherwise.

### GetTelegramBotUsernameOk

`func (o *UserSettingsNotifications) GetTelegramBotUsernameOk() (*string, bool)`

GetTelegramBotUsernameOk returns a tuple with the TelegramBotUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramBotUsername

`func (o *UserSettingsNotifications) SetTelegramBotUsername(v string)`

SetTelegramBotUsername sets TelegramBotUsername field to given value.

### HasTelegramBotUsername

`func (o *UserSettingsNotifications) HasTelegramBotUsername() bool`

HasTelegramBotUsername returns a boolean if a field has been set.

### SetTelegramBotUsernameNil

`func (o *UserSettingsNotifications) SetTelegramBotUsernameNil(b bool)`

 SetTelegramBotUsernameNil sets the value for TelegramBotUsername to be an explicit nil

### UnsetTelegramBotUsername
`func (o *UserSettingsNotifications) UnsetTelegramBotUsername()`

UnsetTelegramBotUsername ensures that no value is present for TelegramBotUsername, not even an explicit nil
### GetTelegramChatId

`func (o *UserSettingsNotifications) GetTelegramChatId() string`

GetTelegramChatId returns the TelegramChatId field if non-nil, zero value otherwise.

### GetTelegramChatIdOk

`func (o *UserSettingsNotifications) GetTelegramChatIdOk() (*string, bool)`

GetTelegramChatIdOk returns a tuple with the TelegramChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramChatId

`func (o *UserSettingsNotifications) SetTelegramChatId(v string)`

SetTelegramChatId sets TelegramChatId field to given value.

### HasTelegramChatId

`func (o *UserSettingsNotifications) HasTelegramChatId() bool`

HasTelegramChatId returns a boolean if a field has been set.

### SetTelegramChatIdNil

`func (o *UserSettingsNotifications) SetTelegramChatIdNil(b bool)`

 SetTelegramChatIdNil sets the value for TelegramChatId to be an explicit nil

### UnsetTelegramChatId
`func (o *UserSettingsNotifications) UnsetTelegramChatId()`

UnsetTelegramChatId ensures that no value is present for TelegramChatId, not even an explicit nil
### GetTelegramMessageThreadId

`func (o *UserSettingsNotifications) GetTelegramMessageThreadId() string`

GetTelegramMessageThreadId returns the TelegramMessageThreadId field if non-nil, zero value otherwise.

### GetTelegramMessageThreadIdOk

`func (o *UserSettingsNotifications) GetTelegramMessageThreadIdOk() (*string, bool)`

GetTelegramMessageThreadIdOk returns a tuple with the TelegramMessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramMessageThreadId

`func (o *UserSettingsNotifications) SetTelegramMessageThreadId(v string)`

SetTelegramMessageThreadId sets TelegramMessageThreadId field to given value.

### HasTelegramMessageThreadId

`func (o *UserSettingsNotifications) HasTelegramMessageThreadId() bool`

HasTelegramMessageThreadId returns a boolean if a field has been set.

### SetTelegramMessageThreadIdNil

`func (o *UserSettingsNotifications) SetTelegramMessageThreadIdNil(b bool)`

 SetTelegramMessageThreadIdNil sets the value for TelegramMessageThreadId to be an explicit nil

### UnsetTelegramMessageThreadId
`func (o *UserSettingsNotifications) UnsetTelegramMessageThreadId()`

UnsetTelegramMessageThreadId ensures that no value is present for TelegramMessageThreadId, not even an explicit nil
### GetTelegramSendSilently

`func (o *UserSettingsNotifications) GetTelegramSendSilently() bool`

GetTelegramSendSilently returns the TelegramSendSilently field if non-nil, zero value otherwise.

### GetTelegramSendSilentlyOk

`func (o *UserSettingsNotifications) GetTelegramSendSilentlyOk() (*bool, bool)`

GetTelegramSendSilentlyOk returns a tuple with the TelegramSendSilently field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramSendSilently

`func (o *UserSettingsNotifications) SetTelegramSendSilently(v bool)`

SetTelegramSendSilently sets TelegramSendSilently field to given value.

### HasTelegramSendSilently

`func (o *UserSettingsNotifications) HasTelegramSendSilently() bool`

HasTelegramSendSilently returns a boolean if a field has been set.

### SetTelegramSendSilentlyNil

`func (o *UserSettingsNotifications) SetTelegramSendSilentlyNil(b bool)`

 SetTelegramSendSilentlyNil sets the value for TelegramSendSilently to be an explicit nil

### UnsetTelegramSendSilently
`func (o *UserSettingsNotifications) UnsetTelegramSendSilently()`

UnsetTelegramSendSilently ensures that no value is present for TelegramSendSilently, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


