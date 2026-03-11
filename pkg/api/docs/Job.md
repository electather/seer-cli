# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextExecutionTime** | Pointer to **string** |  | [optional] 
**Running** | Pointer to **bool** |  | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Job) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Job) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Job) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInterval

`func (o *Job) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Job) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Job) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Job) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Job) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Job) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextExecutionTime

`func (o *Job) GetNextExecutionTime() string`

GetNextExecutionTime returns the NextExecutionTime field if non-nil, zero value otherwise.

### GetNextExecutionTimeOk

`func (o *Job) GetNextExecutionTimeOk() (*string, bool)`

GetNextExecutionTimeOk returns a tuple with the NextExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecutionTime

`func (o *Job) SetNextExecutionTime(v string)`

SetNextExecutionTime sets NextExecutionTime field to given value.

### HasNextExecutionTime

`func (o *Job) HasNextExecutionTime() bool`

HasNextExecutionTime returns a boolean if a field has been set.

### GetRunning

`func (o *Job) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *Job) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *Job) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *Job) HasRunning() bool`

HasRunning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


