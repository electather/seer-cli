# MediaMediaIdStatusPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Is4k** | Pointer to **bool** | When true, updates the 4K status field (status4k). When false or not provided, updates the regular status field (status). This applies to all status values (available, partial, processing, pending, unknown).  | [optional] 

## Methods

### NewMediaMediaIdStatusPostRequest

`func NewMediaMediaIdStatusPostRequest() *MediaMediaIdStatusPostRequest`

NewMediaMediaIdStatusPostRequest instantiates a new MediaMediaIdStatusPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaMediaIdStatusPostRequestWithDefaults

`func NewMediaMediaIdStatusPostRequestWithDefaults() *MediaMediaIdStatusPostRequest`

NewMediaMediaIdStatusPostRequestWithDefaults instantiates a new MediaMediaIdStatusPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIs4k

`func (o *MediaMediaIdStatusPostRequest) GetIs4k() bool`

GetIs4k returns the Is4k field if non-nil, zero value otherwise.

### GetIs4kOk

`func (o *MediaMediaIdStatusPostRequest) GetIs4kOk() (*bool, bool)`

GetIs4kOk returns a tuple with the Is4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs4k

`func (o *MediaMediaIdStatusPostRequest) SetIs4k(v bool)`

SetIs4k sets Is4k field to given value.

### HasIs4k

`func (o *MediaMediaIdStatusPostRequest) HasIs4k() bool`

HasIs4k returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


