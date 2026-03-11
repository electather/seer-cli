# WebhookSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Types** | Pointer to **float32** |  | [optional] 
**Options** | Pointer to [**WebhookSettingsOptions**](WebhookSettingsOptions.md) |  | [optional] 

## Methods

### NewWebhookSettings

`func NewWebhookSettings() *WebhookSettings`

NewWebhookSettings instantiates a new WebhookSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSettingsWithDefaults

`func NewWebhookSettingsWithDefaults() *WebhookSettings`

NewWebhookSettingsWithDefaults instantiates a new WebhookSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WebhookSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebhookSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebhookSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebhookSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTypes

`func (o *WebhookSettings) GetTypes() float32`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *WebhookSettings) GetTypesOk() (*float32, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *WebhookSettings) SetTypes(v float32)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *WebhookSettings) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetOptions

`func (o *WebhookSettings) GetOptions() WebhookSettingsOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *WebhookSettings) GetOptionsOk() (*WebhookSettingsOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *WebhookSettings) SetOptions(v WebhookSettingsOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *WebhookSettings) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


