# PersonResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**ProfilePath** | Pointer to **string** |  | [optional] 
**Adult** | Pointer to **bool** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] [default to "person"]
**KnownFor** | Pointer to [**[]PersonResultKnownForInner**](PersonResultKnownForInner.md) |  | [optional] 

## Methods

### NewPersonResult

`func NewPersonResult() *PersonResult`

NewPersonResult instantiates a new PersonResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonResultWithDefaults

`func NewPersonResultWithDefaults() *PersonResult`

NewPersonResultWithDefaults instantiates a new PersonResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonResult) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonResult) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonResult) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *PersonResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfilePath

`func (o *PersonResult) GetProfilePath() string`

GetProfilePath returns the ProfilePath field if non-nil, zero value otherwise.

### GetProfilePathOk

`func (o *PersonResult) GetProfilePathOk() (*string, bool)`

GetProfilePathOk returns a tuple with the ProfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePath

`func (o *PersonResult) SetProfilePath(v string)`

SetProfilePath sets ProfilePath field to given value.

### HasProfilePath

`func (o *PersonResult) HasProfilePath() bool`

HasProfilePath returns a boolean if a field has been set.

### GetAdult

`func (o *PersonResult) GetAdult() bool`

GetAdult returns the Adult field if non-nil, zero value otherwise.

### GetAdultOk

`func (o *PersonResult) GetAdultOk() (*bool, bool)`

GetAdultOk returns a tuple with the Adult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdult

`func (o *PersonResult) SetAdult(v bool)`

SetAdult sets Adult field to given value.

### HasAdult

`func (o *PersonResult) HasAdult() bool`

HasAdult returns a boolean if a field has been set.

### GetMediaType

`func (o *PersonResult) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *PersonResult) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *PersonResult) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *PersonResult) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetKnownFor

`func (o *PersonResult) GetKnownFor() []PersonResultKnownForInner`

GetKnownFor returns the KnownFor field if non-nil, zero value otherwise.

### GetKnownForOk

`func (o *PersonResult) GetKnownForOk() (*[]PersonResultKnownForInner, bool)`

GetKnownForOk returns a tuple with the KnownFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownFor

`func (o *PersonResult) SetKnownFor(v []PersonResultKnownForInner)`

SetKnownFor sets KnownFor field to given value.

### HasKnownFor

`func (o *PersonResult) HasKnownFor() bool`

HasKnownFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


