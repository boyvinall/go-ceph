# ApiNvmeofSubsystemPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableHa** | **bool** | Enable high availability | 
**GwGroup** | Pointer to **string** | NVMeoF gateway group | [optional] 
**MaxNamespaces** | Pointer to **int32** | Maximum number of namespaces | [optional] [default to 1024]
**Nqn** | **string** | NVMeoF subsystem NQN | 

## Methods

### NewApiNvmeofSubsystemPostRequest

`func NewApiNvmeofSubsystemPostRequest(enableHa bool, nqn string, ) *ApiNvmeofSubsystemPostRequest`

NewApiNvmeofSubsystemPostRequest instantiates a new ApiNvmeofSubsystemPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNvmeofSubsystemPostRequestWithDefaults

`func NewApiNvmeofSubsystemPostRequestWithDefaults() *ApiNvmeofSubsystemPostRequest`

NewApiNvmeofSubsystemPostRequestWithDefaults instantiates a new ApiNvmeofSubsystemPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableHa

`func (o *ApiNvmeofSubsystemPostRequest) GetEnableHa() bool`

GetEnableHa returns the EnableHa field if non-nil, zero value otherwise.

### GetEnableHaOk

`func (o *ApiNvmeofSubsystemPostRequest) GetEnableHaOk() (*bool, bool)`

GetEnableHaOk returns a tuple with the EnableHa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHa

`func (o *ApiNvmeofSubsystemPostRequest) SetEnableHa(v bool)`

SetEnableHa sets EnableHa field to given value.


### GetGwGroup

`func (o *ApiNvmeofSubsystemPostRequest) GetGwGroup() string`

GetGwGroup returns the GwGroup field if non-nil, zero value otherwise.

### GetGwGroupOk

`func (o *ApiNvmeofSubsystemPostRequest) GetGwGroupOk() (*string, bool)`

GetGwGroupOk returns a tuple with the GwGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwGroup

`func (o *ApiNvmeofSubsystemPostRequest) SetGwGroup(v string)`

SetGwGroup sets GwGroup field to given value.

### HasGwGroup

`func (o *ApiNvmeofSubsystemPostRequest) HasGwGroup() bool`

HasGwGroup returns a boolean if a field has been set.

### GetMaxNamespaces

`func (o *ApiNvmeofSubsystemPostRequest) GetMaxNamespaces() int32`

GetMaxNamespaces returns the MaxNamespaces field if non-nil, zero value otherwise.

### GetMaxNamespacesOk

`func (o *ApiNvmeofSubsystemPostRequest) GetMaxNamespacesOk() (*int32, bool)`

GetMaxNamespacesOk returns a tuple with the MaxNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNamespaces

`func (o *ApiNvmeofSubsystemPostRequest) SetMaxNamespaces(v int32)`

SetMaxNamespaces sets MaxNamespaces field to given value.

### HasMaxNamespaces

`func (o *ApiNvmeofSubsystemPostRequest) HasMaxNamespaces() bool`

HasMaxNamespaces returns a boolean if a field has been set.

### GetNqn

`func (o *ApiNvmeofSubsystemPostRequest) GetNqn() string`

GetNqn returns the Nqn field if non-nil, zero value otherwise.

### GetNqnOk

`func (o *ApiNvmeofSubsystemPostRequest) GetNqnOk() (*string, bool)`

GetNqnOk returns a tuple with the Nqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNqn

`func (o *ApiNvmeofSubsystemPostRequest) SetNqn(v string)`

SetNqn sets Nqn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


