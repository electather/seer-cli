# DiscordSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Types** | Pointer to **float32** |  | [optional] 
**Options** | Pointer to [**DiscordSettingsOptions**](DiscordSettingsOptions.md) |  | [optional] 

## Methods

### NewDiscordSettings

`func NewDiscordSettings() *DiscordSettings`

NewDiscordSettings instantiates a new DiscordSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscordSettingsWithDefaults

`func NewDiscordSettingsWithDefaults() *DiscordSettings`

NewDiscordSettingsWithDefaults instantiates a new DiscordSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DiscordSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DiscordSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DiscordSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DiscordSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTypes

`func (o *DiscordSettings) GetTypes() float32`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *DiscordSettings) GetTypesOk() (*float32, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *DiscordSettings) SetTypes(v float32)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *DiscordSettings) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetOptions

`func (o *DiscordSettings) GetOptions() DiscordSettingsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DiscordSettings) GetOptionsOk() (*DiscordSettingsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DiscordSettings) SetOptions(v DiscordSettingsOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DiscordSettings) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


