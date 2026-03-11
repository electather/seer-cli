# BlocklistGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]BlocklistGet200ResponseResultsInner**](BlocklistGet200ResponseResultsInner.md) |  | [optional] 

## Methods

### NewBlocklistGet200Response

`func NewBlocklistGet200Response() *BlocklistGet200Response`

NewBlocklistGet200Response instantiates a new BlocklistGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlocklistGet200ResponseWithDefaults

`func NewBlocklistGet200ResponseWithDefaults() *BlocklistGet200Response`

NewBlocklistGet200ResponseWithDefaults instantiates a new BlocklistGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *BlocklistGet200Response) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *BlocklistGet200Response) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *BlocklistGet200Response) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *BlocklistGet200Response) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *BlocklistGet200Response) GetResults() []BlocklistGet200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BlocklistGet200Response) GetResultsOk() (*[]BlocklistGet200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BlocklistGet200Response) SetResults(v []BlocklistGet200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *BlocklistGet200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


