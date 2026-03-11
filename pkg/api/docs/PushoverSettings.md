# PushoverSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Types** | Pointer to **float32** |  | [optional] 
**Options** | Pointer to [**PushoverSettingsOptions**](PushoverSettingsOptions.md) |  | [optional] 

## Methods

### NewPushoverSettings

`func NewPushoverSettings() *PushoverSettings`

NewPushoverSettings instantiates a new PushoverSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushoverSettingsWithDefaults

`func NewPushoverSettingsWithDefaults() *PushoverSettings`

NewPushoverSettingsWithDefaults instantiates a new PushoverSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PushoverSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PushoverSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PushoverSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PushoverSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTypes

`func (o *PushoverSettings) GetTypes() float32`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *PushoverSettings) GetTypesOk() (*float32, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *PushoverSettings) SetTypes(v float32)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *PushoverSettings) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetOptions

`func (o *PushoverSettings) GetOptions() PushoverSettingsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PushoverSettings) GetOptionsOk() (*PushoverSettingsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PushoverSettings) SetOptions(v PushoverSettingsOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PushoverSettings) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


