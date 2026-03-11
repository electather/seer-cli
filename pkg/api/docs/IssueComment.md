# IssueComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewIssueComment

`func NewIssueComment() *IssueComment`

NewIssueComment instantiates a new IssueComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCommentWithDefaults

`func NewIssueCommentWithDefaults() *IssueComment`

NewIssueCommentWithDefaults instantiates a new IssueComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueComment) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueComment) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueComment) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *IssueComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUser

`func (o *IssueComment) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IssueComment) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IssueComment) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *IssueComment) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetMessage

`func (o *IssueComment) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IssueComment) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IssueComment) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IssueComment) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


