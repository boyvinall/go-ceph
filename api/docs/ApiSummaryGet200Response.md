# ApiSummaryGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutingTasks** | **[]string** |  | 
**FinishedTasks** | [**[]ApiSummaryGet200ResponseFinishedTasksInner**](ApiSummaryGet200ResponseFinishedTasksInner.md) |  | 
**HaveMonConnection** | **string** |  | 
**HealthStatus** | **string** |  | 
**MgrHost** | **string** |  | 
**MgrId** | **string** |  | 
**RbdMirroring** | [**ApiSummaryGet200ResponseRbdMirroring**](ApiSummaryGet200ResponseRbdMirroring.md) |  | 
**Version** | **string** |  | 

## Methods

### NewApiSummaryGet200Response

`func NewApiSummaryGet200Response(executingTasks []string, finishedTasks []ApiSummaryGet200ResponseFinishedTasksInner, haveMonConnection string, healthStatus string, mgrHost string, mgrId string, rbdMirroring ApiSummaryGet200ResponseRbdMirroring, version string, ) *ApiSummaryGet200Response`

NewApiSummaryGet200Response instantiates a new ApiSummaryGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSummaryGet200ResponseWithDefaults

`func NewApiSummaryGet200ResponseWithDefaults() *ApiSummaryGet200Response`

NewApiSummaryGet200ResponseWithDefaults instantiates a new ApiSummaryGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutingTasks

`func (o *ApiSummaryGet200Response) GetExecutingTasks() []string`

GetExecutingTasks returns the ExecutingTasks field if non-nil, zero value otherwise.

### GetExecutingTasksOk

`func (o *ApiSummaryGet200Response) GetExecutingTasksOk() (*[]string, bool)`

GetExecutingTasksOk returns a tuple with the ExecutingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutingTasks

`func (o *ApiSummaryGet200Response) SetExecutingTasks(v []string)`

SetExecutingTasks sets ExecutingTasks field to given value.


### GetFinishedTasks

`func (o *ApiSummaryGet200Response) GetFinishedTasks() []ApiSummaryGet200ResponseFinishedTasksInner`

GetFinishedTasks returns the FinishedTasks field if non-nil, zero value otherwise.

### GetFinishedTasksOk

`func (o *ApiSummaryGet200Response) GetFinishedTasksOk() (*[]ApiSummaryGet200ResponseFinishedTasksInner, bool)`

GetFinishedTasksOk returns a tuple with the FinishedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTasks

`func (o *ApiSummaryGet200Response) SetFinishedTasks(v []ApiSummaryGet200ResponseFinishedTasksInner)`

SetFinishedTasks sets FinishedTasks field to given value.


### GetHaveMonConnection

`func (o *ApiSummaryGet200Response) GetHaveMonConnection() string`

GetHaveMonConnection returns the HaveMonConnection field if non-nil, zero value otherwise.

### GetHaveMonConnectionOk

`func (o *ApiSummaryGet200Response) GetHaveMonConnectionOk() (*string, bool)`

GetHaveMonConnectionOk returns a tuple with the HaveMonConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaveMonConnection

`func (o *ApiSummaryGet200Response) SetHaveMonConnection(v string)`

SetHaveMonConnection sets HaveMonConnection field to given value.


### GetHealthStatus

`func (o *ApiSummaryGet200Response) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *ApiSummaryGet200Response) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *ApiSummaryGet200Response) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.


### GetMgrHost

`func (o *ApiSummaryGet200Response) GetMgrHost() string`

GetMgrHost returns the MgrHost field if non-nil, zero value otherwise.

### GetMgrHostOk

`func (o *ApiSummaryGet200Response) GetMgrHostOk() (*string, bool)`

GetMgrHostOk returns a tuple with the MgrHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgrHost

`func (o *ApiSummaryGet200Response) SetMgrHost(v string)`

SetMgrHost sets MgrHost field to given value.


### GetMgrId

`func (o *ApiSummaryGet200Response) GetMgrId() string`

GetMgrId returns the MgrId field if non-nil, zero value otherwise.

### GetMgrIdOk

`func (o *ApiSummaryGet200Response) GetMgrIdOk() (*string, bool)`

GetMgrIdOk returns a tuple with the MgrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgrId

`func (o *ApiSummaryGet200Response) SetMgrId(v string)`

SetMgrId sets MgrId field to given value.


### GetRbdMirroring

`func (o *ApiSummaryGet200Response) GetRbdMirroring() ApiSummaryGet200ResponseRbdMirroring`

GetRbdMirroring returns the RbdMirroring field if non-nil, zero value otherwise.

### GetRbdMirroringOk

`func (o *ApiSummaryGet200Response) GetRbdMirroringOk() (*ApiSummaryGet200ResponseRbdMirroring, bool)`

GetRbdMirroringOk returns a tuple with the RbdMirroring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbdMirroring

`func (o *ApiSummaryGet200Response) SetRbdMirroring(v ApiSummaryGet200ResponseRbdMirroring)`

SetRbdMirroring sets RbdMirroring field to given value.


### GetVersion

`func (o *ApiSummaryGet200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiSummaryGet200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiSummaryGet200Response) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


