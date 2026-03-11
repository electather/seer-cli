# PlexConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** |  | 
**Address** | **string** |  | 
**Port** | **float32** |  | 
**Uri** | **string** |  | 
**Local** | **bool** |  | 
**Status** | Pointer to **float32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewPlexConnection

`func NewPlexConnection(protocol string, address string, port float32, uri string, local bool, ) *PlexConnection`

NewPlexConnection instantiates a new PlexConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlexConnectionWithDefaults

`func NewPlexConnectionWithDefaults() *PlexConnection`

NewPlexConnectionWithDefaults instantiates a new PlexConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *PlexConnection) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PlexConnection) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PlexConnection) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetAddress

`func (o *PlexConnection) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PlexConnection) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PlexConnection) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPort

`func (o *PlexConnection) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PlexConnection) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PlexConnection) SetPort(v float32)`

SetPort sets Port field to given value.


### GetUri

`func (o *PlexConnection) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PlexConnection) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PlexConnection) SetUri(v string)`

SetUri sets Uri field to given value.


### GetLocal

`func (o *PlexConnection) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *PlexConnection) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *PlexConnection) SetLocal(v bool)`

SetLocal sets Local field to given value.


### GetStatus

`func (o *PlexConnection) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlexConnection) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlexConnection) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlexConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *PlexConnection) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PlexConnection) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PlexConnection) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PlexConnection) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


