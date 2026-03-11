# TelegramSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Types** | Pointer to **float32** |  | [optional] 
**Options** | Pointer to [**TelegramSettingsOptions**](TelegramSettingsOptions.md) |  | [optional] 

## Methods

### NewTelegramSettings

`func NewTelegramSettings() *TelegramSettings`

NewTelegramSettings instantiates a new TelegramSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelegramSettingsWithDefaults

`func NewTelegramSettingsWithDefaults() *TelegramSettings`

NewTelegramSettingsWithDefaults instantiates a new TelegramSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TelegramSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TelegramSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TelegramSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TelegramSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTypes

`func (o *TelegramSettings) GetTypes() float32`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *TelegramSettings) GetTypesOk() (*float32, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *TelegramSettings) SetTypes(v float32)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *TelegramSettings) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetOptions

`func (o *TelegramSettings) GetOptions() TelegramSettingsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TelegramSettings) GetOptionsOk() (*TelegramSettingsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TelegramSettings) SetOptions(v TelegramSettingsOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *TelegramSettings) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


