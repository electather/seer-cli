# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**LogoPath** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewCompany

`func NewCompany() *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Company) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Company) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Company) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *Company) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogoPath

`func (o *Company) GetLogoPath() string`

GetLogoPath returns the LogoPath field if non-nil, zero value otherwise.

### GetLogoPathOk

`func (o *Company) GetLogoPathOk() (*string, bool)`

GetLogoPathOk returns a tuple with the LogoPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoPath

`func (o *Company) SetLogoPath(v string)`

SetLogoPath sets LogoPath field to given value.

### HasLogoPath

`func (o *Company) HasLogoPath() bool`

HasLogoPath returns a boolean if a field has been set.

### SetLogoPathNil

`func (o *Company) SetLogoPathNil(b bool)`

 SetLogoPathNil sets the value for LogoPath to be an explicit nil

### UnsetLogoPath
`func (o *Company) UnsetLogoPath()`

UnsetLogoPath ensures that no value is present for LogoPath, not even an explicit nil
### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Company) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


