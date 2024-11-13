# ApiNvmeofSubsystemNqnNamespacePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSize** | Pointer to **int32** | NVMeoF namespace block size | [optional] [default to 512]
**CreateImage** | Pointer to **bool** | Create RBD image | [optional] [default to true]
**GwGroup** | Pointer to **string** | NVMeoF gateway group | [optional] 
**LoadBalancingGroup** | Pointer to **int32** | Load balancing group | [optional] 
**RbdImageName** | **string** | RBD image name | 
**RbdPool** | Pointer to **string** | RBD pool name | [optional] [default to "rbd"]
**Size** | Pointer to **int32** | RBD image size | [optional] [default to 1024]

## Methods

### NewApiNvmeofSubsystemNqnNamespacePostRequest

`func NewApiNvmeofSubsystemNqnNamespacePostRequest(rbdImageName string, ) *ApiNvmeofSubsystemNqnNamespacePostRequest`

NewApiNvmeofSubsystemNqnNamespacePostRequest instantiates a new ApiNvmeofSubsystemNqnNamespacePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNvmeofSubsystemNqnNamespacePostRequestWithDefaults

`func NewApiNvmeofSubsystemNqnNamespacePostRequestWithDefaults() *ApiNvmeofSubsystemNqnNamespacePostRequest`

NewApiNvmeofSubsystemNqnNamespacePostRequestWithDefaults instantiates a new ApiNvmeofSubsystemNqnNamespacePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSize

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetBlockSize() int32`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetBlockSizeOk() (*int32, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) SetBlockSize(v int32)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetCreateImage

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetCreateImage() bool`

GetCreateImage returns the CreateImage field if non-nil, zero value otherwise.

### GetCreateImageOk

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetCreateImageOk() (*bool, bool)`

GetCreateImageOk returns a tuple with the CreateImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateImage

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) SetCreateImage(v bool)`

SetCreateImage sets CreateImage field to given value.

### HasCreateImage

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) HasCreateImage() bool`

HasCreateImage returns a boolean if a field has been set.

### GetGwGroup

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetGwGroup() string`

GetGwGroup returns the GwGroup field if non-nil, zero value otherwise.

### GetGwGroupOk

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetGwGroupOk() (*string, bool)`

GetGwGroupOk returns a tuple with the GwGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwGroup

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) SetGwGroup(v string)`

SetGwGroup sets GwGroup field to given value.

### HasGwGroup

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) HasGwGroup() bool`

HasGwGroup returns a boolean if a field has been set.

### GetLoadBalancingGroup

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetLoadBalancingGroup() int32`

GetLoadBalancingGroup returns the LoadBalancingGroup field if non-nil, zero value otherwise.

### GetLoadBalancingGroupOk

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetLoadBalancingGroupOk() (*int32, bool)`

GetLoadBalancingGroupOk returns a tuple with the LoadBalancingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingGroup

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) SetLoadBalancingGroup(v int32)`

SetLoadBalancingGroup sets LoadBalancingGroup field to given value.

### HasLoadBalancingGroup

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) HasLoadBalancingGroup() bool`

HasLoadBalancingGroup returns a boolean if a field has been set.

### GetRbdImageName

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetRbdImageName() string`

GetRbdImageName returns the RbdImageName field if non-nil, zero value otherwise.

### GetRbdImageNameOk

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetRbdImageNameOk() (*string, bool)`

GetRbdImageNameOk returns a tuple with the RbdImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbdImageName

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) SetRbdImageName(v string)`

SetRbdImageName sets RbdImageName field to given value.


### GetRbdPool

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetRbdPool() string`

GetRbdPool returns the RbdPool field if non-nil, zero value otherwise.

### GetRbdPoolOk

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetRbdPoolOk() (*string, bool)`

GetRbdPoolOk returns a tuple with the RbdPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbdPool

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) SetRbdPool(v string)`

SetRbdPool sets RbdPool field to given value.

### HasRbdPool

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) HasRbdPool() bool`

HasRbdPool returns a boolean if a field has been set.

### GetSize

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ApiNvmeofSubsystemNqnNamespacePostRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


