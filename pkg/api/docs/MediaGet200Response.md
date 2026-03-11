# MediaGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]MediaInfo**](MediaInfo.md) |  | [optional] 

## Methods

### NewMediaGet200Response

`func NewMediaGet200Response() *MediaGet200Response`

NewMediaGet200Response instantiates a new MediaGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaGet200ResponseWithDefaults

`func NewMediaGet200ResponseWithDefaults() *MediaGet200Response`

NewMediaGet200ResponseWithDefaults instantiates a new MediaGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *MediaGet200Response) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *MediaGet200Response) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *MediaGet200Response) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *MediaGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *MediaGet200Response) GetResults() []MediaInfo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MediaGet200Response) GetResultsOk() (*[]MediaInfo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MediaGet200Response) SetResults(v []MediaInfo)`

SetResults sets Results field to given value.

### HasResults

`func (o *MediaGet200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


