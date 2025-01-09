# RbdImageSnapshotsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**IsProtected** | Pointer to **bool** |  | [optional] 
**UsedBytes** | Pointer to **NullableInt32** |  | [optional] 
**Children** | Pointer to **[]int32** |  | [optional] 
**DiskUsage** | Pointer to **float32** |  | [optional] 

## Methods

### NewRbdImageSnapshotsInner

`func NewRbdImageSnapshotsInner() *RbdImageSnapshotsInner`

NewRbdImageSnapshotsInner instantiates a new RbdImageSnapshotsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbdImageSnapshotsInnerWithDefaults

`func NewRbdImageSnapshotsInnerWithDefaults() *RbdImageSnapshotsInner`

NewRbdImageSnapshotsInnerWithDefaults instantiates a new RbdImageSnapshotsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RbdImageSnapshotsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RbdImageSnapshotsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RbdImageSnapshotsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RbdImageSnapshotsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSize

`func (o *RbdImageSnapshotsInner) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RbdImageSnapshotsInner) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RbdImageSnapshotsInner) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RbdImageSnapshotsInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetName

`func (o *RbdImageSnapshotsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RbdImageSnapshotsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RbdImageSnapshotsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RbdImageSnapshotsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *RbdImageSnapshotsInner) GetNamespace() int32`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RbdImageSnapshotsInner) GetNamespaceOk() (*int32, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RbdImageSnapshotsInner) SetNamespace(v int32)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RbdImageSnapshotsInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTimestamp

`func (o *RbdImageSnapshotsInner) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RbdImageSnapshotsInner) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RbdImageSnapshotsInner) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RbdImageSnapshotsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetIsProtected

`func (o *RbdImageSnapshotsInner) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *RbdImageSnapshotsInner) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *RbdImageSnapshotsInner) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *RbdImageSnapshotsInner) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetUsedBytes

`func (o *RbdImageSnapshotsInner) GetUsedBytes() int32`

GetUsedBytes returns the UsedBytes field if non-nil, zero value otherwise.

### GetUsedBytesOk

`func (o *RbdImageSnapshotsInner) GetUsedBytesOk() (*int32, bool)`

GetUsedBytesOk returns a tuple with the UsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBytes

`func (o *RbdImageSnapshotsInner) SetUsedBytes(v int32)`

SetUsedBytes sets UsedBytes field to given value.

### HasUsedBytes

`func (o *RbdImageSnapshotsInner) HasUsedBytes() bool`

HasUsedBytes returns a boolean if a field has been set.

### SetUsedBytesNil

`func (o *RbdImageSnapshotsInner) SetUsedBytesNil(b bool)`

 SetUsedBytesNil sets the value for UsedBytes to be an explicit nil

### UnsetUsedBytes
`func (o *RbdImageSnapshotsInner) UnsetUsedBytes()`

UnsetUsedBytes ensures that no value is present for UsedBytes, not even an explicit nil
### GetChildren

`func (o *RbdImageSnapshotsInner) GetChildren() []int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *RbdImageSnapshotsInner) GetChildrenOk() (*[]int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *RbdImageSnapshotsInner) SetChildren(v []int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *RbdImageSnapshotsInner) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetDiskUsage

`func (o *RbdImageSnapshotsInner) GetDiskUsage() float32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *RbdImageSnapshotsInner) GetDiskUsageOk() (*float32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *RbdImageSnapshotsInner) SetDiskUsage(v float32)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *RbdImageSnapshotsInner) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


