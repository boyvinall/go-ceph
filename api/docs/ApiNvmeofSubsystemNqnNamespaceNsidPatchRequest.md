# ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GwGroup** | Pointer to **string** | NVMeoF gateway group | [optional] 
**LoadBalancingGroup** | Pointer to **int32** | Load balancing group | [optional] 
**RMbytesPerSecond** | Pointer to **int32** | Read MB/s | [optional] 
**RbdImageSize** | Pointer to **int32** | RBD image size | [optional] 
**RwIosPerSecond** | Pointer to **int32** | Read/Write IOPS | [optional] 
**RwMbytesPerSecond** | Pointer to **int32** | Read/Write MB/s | [optional] 
**WMbytesPerSecond** | Pointer to **int32** | Write MB/s | [optional] 

## Methods

### NewApiNvmeofSubsystemNqnNamespaceNsidPatchRequest

`func NewApiNvmeofSubsystemNqnNamespaceNsidPatchRequest() *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest`

NewApiNvmeofSubsystemNqnNamespaceNsidPatchRequest instantiates a new ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNvmeofSubsystemNqnNamespaceNsidPatchRequestWithDefaults

`func NewApiNvmeofSubsystemNqnNamespaceNsidPatchRequestWithDefaults() *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest`

NewApiNvmeofSubsystemNqnNamespaceNsidPatchRequestWithDefaults instantiates a new ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGwGroup

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetGwGroup() string`

GetGwGroup returns the GwGroup field if non-nil, zero value otherwise.

### GetGwGroupOk

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetGwGroupOk() (*string, bool)`

GetGwGroupOk returns a tuple with the GwGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwGroup

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) SetGwGroup(v string)`

SetGwGroup sets GwGroup field to given value.

### HasGwGroup

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) HasGwGroup() bool`

HasGwGroup returns a boolean if a field has been set.

### GetLoadBalancingGroup

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetLoadBalancingGroup() int32`

GetLoadBalancingGroup returns the LoadBalancingGroup field if non-nil, zero value otherwise.

### GetLoadBalancingGroupOk

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetLoadBalancingGroupOk() (*int32, bool)`

GetLoadBalancingGroupOk returns a tuple with the LoadBalancingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingGroup

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) SetLoadBalancingGroup(v int32)`

SetLoadBalancingGroup sets LoadBalancingGroup field to given value.

### HasLoadBalancingGroup

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) HasLoadBalancingGroup() bool`

HasLoadBalancingGroup returns a boolean if a field has been set.

### GetRMbytesPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetRMbytesPerSecond() int32`

GetRMbytesPerSecond returns the RMbytesPerSecond field if non-nil, zero value otherwise.

### GetRMbytesPerSecondOk

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetRMbytesPerSecondOk() (*int32, bool)`

GetRMbytesPerSecondOk returns a tuple with the RMbytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRMbytesPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) SetRMbytesPerSecond(v int32)`

SetRMbytesPerSecond sets RMbytesPerSecond field to given value.

### HasRMbytesPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) HasRMbytesPerSecond() bool`

HasRMbytesPerSecond returns a boolean if a field has been set.

### GetRbdImageSize

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetRbdImageSize() int32`

GetRbdImageSize returns the RbdImageSize field if non-nil, zero value otherwise.

### GetRbdImageSizeOk

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetRbdImageSizeOk() (*int32, bool)`

GetRbdImageSizeOk returns a tuple with the RbdImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbdImageSize

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) SetRbdImageSize(v int32)`

SetRbdImageSize sets RbdImageSize field to given value.

### HasRbdImageSize

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) HasRbdImageSize() bool`

HasRbdImageSize returns a boolean if a field has been set.

### GetRwIosPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetRwIosPerSecond() int32`

GetRwIosPerSecond returns the RwIosPerSecond field if non-nil, zero value otherwise.

### GetRwIosPerSecondOk

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetRwIosPerSecondOk() (*int32, bool)`

GetRwIosPerSecondOk returns a tuple with the RwIosPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRwIosPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) SetRwIosPerSecond(v int32)`

SetRwIosPerSecond sets RwIosPerSecond field to given value.

### HasRwIosPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) HasRwIosPerSecond() bool`

HasRwIosPerSecond returns a boolean if a field has been set.

### GetRwMbytesPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetRwMbytesPerSecond() int32`

GetRwMbytesPerSecond returns the RwMbytesPerSecond field if non-nil, zero value otherwise.

### GetRwMbytesPerSecondOk

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetRwMbytesPerSecondOk() (*int32, bool)`

GetRwMbytesPerSecondOk returns a tuple with the RwMbytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRwMbytesPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) SetRwMbytesPerSecond(v int32)`

SetRwMbytesPerSecond sets RwMbytesPerSecond field to given value.

### HasRwMbytesPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) HasRwMbytesPerSecond() bool`

HasRwMbytesPerSecond returns a boolean if a field has been set.

### GetWMbytesPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetWMbytesPerSecond() int32`

GetWMbytesPerSecond returns the WMbytesPerSecond field if non-nil, zero value otherwise.

### GetWMbytesPerSecondOk

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) GetWMbytesPerSecondOk() (*int32, bool)`

GetWMbytesPerSecondOk returns a tuple with the WMbytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWMbytesPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) SetWMbytesPerSecond(v int32)`

SetWMbytesPerSecond sets WMbytesPerSecond field to given value.

### HasWMbytesPerSecond

`func (o *ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest) HasWMbytesPerSecond() bool`

HasWMbytesPerSecond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


