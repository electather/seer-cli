# NtfySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Types** | Pointer to **float32** |  | [optional] 
**Options** | Pointer to [**NtfySettingsOptions**](NtfySettingsOptions.md) |  | [optional] 

## Methods

### NewNtfySettings

`func NewNtfySettings() *NtfySettings`

NewNtfySettings instantiates a new NtfySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtfySettingsWithDefaults

`func NewNtfySettingsWithDefaults() *NtfySettings`

NewNtfySettingsWithDefaults instantiates a new NtfySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NtfySettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NtfySettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NtfySettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NtfySettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTypes

`func (o *NtfySettings) GetTypes() float32`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *NtfySettings) GetTypesOk() (*float32, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *NtfySettings) SetTypes(v float32)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *NtfySettings) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetOptions

`func (o *NtfySettings) GetOptions() NtfySettingsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *NtfySettings) GetOptionsOk() (*NtfySettingsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *NtfySettings) SetOptions(v NtfySettingsOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *NtfySettings) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


