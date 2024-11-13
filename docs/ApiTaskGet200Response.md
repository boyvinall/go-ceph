# ApiTaskGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutingTasks** | **string** | ongoing executing tasks | 
**FinishedTasks** | [**[]ApiTaskGet200ResponseFinishedTasksInner**](ApiTaskGet200ResponseFinishedTasksInner.md) |  | 

## Methods

### NewApiTaskGet200Response

`func NewApiTaskGet200Response(executingTasks string, finishedTasks []ApiTaskGet200ResponseFinishedTasksInner, ) *ApiTaskGet200Response`

NewApiTaskGet200Response instantiates a new ApiTaskGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTaskGet200ResponseWithDefaults

`func NewApiTaskGet200ResponseWithDefaults() *ApiTaskGet200Response`

NewApiTaskGet200ResponseWithDefaults instantiates a new ApiTaskGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutingTasks

`func (o *ApiTaskGet200Response) GetExecutingTasks() string`

GetExecutingTasks returns the ExecutingTasks field if non-nil, zero value otherwise.

### GetExecutingTasksOk

`func (o *ApiTaskGet200Response) GetExecutingTasksOk() (*string, bool)`

GetExecutingTasksOk returns a tuple with the ExecutingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutingTasks

`func (o *ApiTaskGet200Response) SetExecutingTasks(v string)`

SetExecutingTasks sets ExecutingTasks field to given value.


### GetFinishedTasks

`func (o *ApiTaskGet200Response) GetFinishedTasks() []ApiTaskGet200ResponseFinishedTasksInner`

GetFinishedTasks returns the FinishedTasks field if non-nil, zero value otherwise.

### GetFinishedTasksOk

`func (o *ApiTaskGet200Response) GetFinishedTasksOk() (*[]ApiTaskGet200ResponseFinishedTasksInner, bool)`

GetFinishedTasksOk returns a tuple with the FinishedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTasks

`func (o *ApiTaskGet200Response) SetFinishedTasks(v []ApiTaskGet200ResponseFinishedTasksInner)`

SetFinishedTasks sets FinishedTasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


