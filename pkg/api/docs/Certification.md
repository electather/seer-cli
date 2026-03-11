# Certification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certification** | **string** |  | 
**Meaning** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewCertification

`func NewCertification(certification string, ) *Certification`

NewCertification instantiates a new Certification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificationWithDefaults

`func NewCertificationWithDefaults() *Certification`

NewCertificationWithDefaults instantiates a new Certification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertification

`func (o *Certification) GetCertification() string`

GetCertification returns the Certification field if non-nil, zero value otherwise.

### GetCertificationOk

`func (o *Certification) GetCertificationOk() (*string, bool)`

GetCertificationOk returns a tuple with the Certification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertification

`func (o *Certification) SetCertification(v string)`

SetCertification sets Certification field to given value.


### GetMeaning

`func (o *Certification) GetMeaning() string`

GetMeaning returns the Meaning field if non-nil, zero value otherwise.

### GetMeaningOk

`func (o *Certification) GetMeaningOk() (*string, bool)`

GetMeaningOk returns a tuple with the Meaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeaning

`func (o *Certification) SetMeaning(v string)`

SetMeaning sets Meaning field to given value.

### HasMeaning

`func (o *Certification) HasMeaning() bool`

HasMeaning returns a boolean if a field has been set.

### SetMeaningNil

`func (o *Certification) SetMeaningNil(b bool)`

 SetMeaningNil sets the value for Meaning to be an explicit nil

### UnsetMeaning
`func (o *Certification) UnsetMeaning()`

UnsetMeaning ensures that no value is present for Meaning, not even an explicit nil
### GetOrder

`func (o *Certification) GetOrder() float32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Certification) GetOrderOk() (*float32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Certification) SetOrder(v float32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Certification) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *Certification) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *Certification) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


