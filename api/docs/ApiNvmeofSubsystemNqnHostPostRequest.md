# ApiNvmeofSubsystemNqnHostPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GwGroup** | Pointer to **string** | NVMeoF gateway group | [optional] 
**HostNqn** | **string** | NVMeoF host NQN. Use \&quot;*\&quot; to allow any host. | 

## Methods

### NewApiNvmeofSubsystemNqnHostPostRequest

`func NewApiNvmeofSubsystemNqnHostPostRequest(hostNqn string, ) *ApiNvmeofSubsystemNqnHostPostRequest`

NewApiNvmeofSubsystemNqnHostPostRequest instantiates a new ApiNvmeofSubsystemNqnHostPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNvmeofSubsystemNqnHostPostRequestWithDefaults

`func NewApiNvmeofSubsystemNqnHostPostRequestWithDefaults() *ApiNvmeofSubsystemNqnHostPostRequest`

NewApiNvmeofSubsystemNqnHostPostRequestWithDefaults instantiates a new ApiNvmeofSubsystemNqnHostPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGwGroup

`func (o *ApiNvmeofSubsystemNqnHostPostRequest) GetGwGroup() string`

GetGwGroup returns the GwGroup field if non-nil, zero value otherwise.

### GetGwGroupOk

`func (o *ApiNvmeofSubsystemNqnHostPostRequest) GetGwGroupOk() (*string, bool)`

GetGwGroupOk returns a tuple with the GwGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwGroup

`func (o *ApiNvmeofSubsystemNqnHostPostRequest) SetGwGroup(v string)`

SetGwGroup sets GwGroup field to given value.

### HasGwGroup

`func (o *ApiNvmeofSubsystemNqnHostPostRequest) HasGwGroup() bool`

HasGwGroup returns a boolean if a field has been set.

### GetHostNqn

`func (o *ApiNvmeofSubsystemNqnHostPostRequest) GetHostNqn() string`

GetHostNqn returns the HostNqn field if non-nil, zero value otherwise.

### GetHostNqnOk

`func (o *ApiNvmeofSubsystemNqnHostPostRequest) GetHostNqnOk() (*string, bool)`

GetHostNqnOk returns a tuple with the HostNqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNqn

`func (o *ApiNvmeofSubsystemNqnHostPostRequest) SetHostNqn(v string)`

SetHostNqn sets HostNqn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


