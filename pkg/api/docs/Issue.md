# Issue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**IssueType** | Pointer to **float32** |  | [optional] 
**Media** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**ModifiedBy** | Pointer to [**User**](User.md) |  | [optional] 
**Comments** | Pointer to [**[]IssueComment**](IssueComment.md) |  | [optional] 

## Methods

### NewIssue

`func NewIssue() *Issue`

NewIssue instantiates a new Issue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWithDefaults

`func NewIssueWithDefaults() *Issue`

NewIssueWithDefaults instantiates a new Issue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Issue) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Issue) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Issue) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *Issue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueType

`func (o *Issue) GetIssueType() float32`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *Issue) GetIssueTypeOk() (*float32, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *Issue) SetIssueType(v float32)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *Issue) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetMedia

`func (o *Issue) GetMedia() MediaInfo`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *Issue) GetMediaOk() (*MediaInfo, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *Issue) SetMedia(v MediaInfo)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *Issue) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Issue) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Issue) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Issue) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Issue) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Issue) GetModifiedBy() User`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Issue) GetModifiedByOk() (*User, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Issue) SetModifiedBy(v User)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Issue) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetComments

`func (o *Issue) GetComments() []IssueComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Issue) GetCommentsOk() (*[]IssueComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Issue) SetComments(v []IssueComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Issue) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


