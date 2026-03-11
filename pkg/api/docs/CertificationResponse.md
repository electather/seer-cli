# CertificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certifications** | Pointer to [**map[string][]Certification**](array.md) |  | [optional] 

## Methods

### NewCertificationResponse

`func NewCertificationResponse() *CertificationResponse`

NewCertificationResponse instantiates a new CertificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificationResponseWithDefaults

`func NewCertificationResponseWithDefaults() *CertificationResponse`

NewCertificationResponseWithDefaults instantiates a new CertificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertifications

`func (o *CertificationResponse) GetCertifications() map[string][]Certification`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *CertificationResponse) GetCertificationsOk() (*map[string][]Certification, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *CertificationResponse) SetCertifications(v map[string][]Certification)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *CertificationResponse) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


