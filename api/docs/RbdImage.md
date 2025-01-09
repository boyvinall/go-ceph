# RbdImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **float32** |  | [optional] 
**ObjSize** | Pointer to **int32** |  | [optional] 
**NumObjs** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**BlockNamePrefix** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImageFormat** | Pointer to **int32** |  | [optional] 
**PoolName** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **int32** |  | [optional] 
**FeaturesName** | Pointer to **[]string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**StripeCount** | Pointer to **int32** |  | [optional] 
**StripeUnit** | Pointer to **int32** |  | [optional] 
**DataPool** | Pointer to **NullableString** |  | [optional] 
**Parent** | Pointer to **NullableString** |  | [optional] 
**Snapshots** | Pointer to [**[]RbdImageSnapshotsInner**](RbdImageSnapshotsInner.md) |  | [optional] 
**TotalDiskUsage** | Pointer to **float32** |  | [optional] 
**DiskUsage** | Pointer to **int32** |  | [optional] 
**Configuration** | Pointer to [**[]RbdImageConfigurationInner**](RbdImageConfigurationInner.md) |  | [optional] 

## Methods

### NewRbdImage

`func NewRbdImage() *RbdImage`

NewRbdImage instantiates a new RbdImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbdImageWithDefaults

`func NewRbdImageWithDefaults() *RbdImage`

NewRbdImageWithDefaults instantiates a new RbdImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *RbdImage) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RbdImage) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RbdImage) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RbdImage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetObjSize

`func (o *RbdImage) GetObjSize() int32`

GetObjSize returns the ObjSize field if non-nil, zero value otherwise.

### GetObjSizeOk

`func (o *RbdImage) GetObjSizeOk() (*int32, bool)`

GetObjSizeOk returns a tuple with the ObjSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSize

`func (o *RbdImage) SetObjSize(v int32)`

SetObjSize sets ObjSize field to given value.

### HasObjSize

`func (o *RbdImage) HasObjSize() bool`

HasObjSize returns a boolean if a field has been set.

### GetNumObjs

`func (o *RbdImage) GetNumObjs() int32`

GetNumObjs returns the NumObjs field if non-nil, zero value otherwise.

### GetNumObjsOk

`func (o *RbdImage) GetNumObjsOk() (*int32, bool)`

GetNumObjsOk returns a tuple with the NumObjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjs

`func (o *RbdImage) SetNumObjs(v int32)`

SetNumObjs sets NumObjs field to given value.

### HasNumObjs

`func (o *RbdImage) HasNumObjs() bool`

HasNumObjs returns a boolean if a field has been set.

### GetOrder

`func (o *RbdImage) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RbdImage) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RbdImage) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RbdImage) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetBlockNamePrefix

`func (o *RbdImage) GetBlockNamePrefix() string`

GetBlockNamePrefix returns the BlockNamePrefix field if non-nil, zero value otherwise.

### GetBlockNamePrefixOk

`func (o *RbdImage) GetBlockNamePrefixOk() (*string, bool)`

GetBlockNamePrefixOk returns a tuple with the BlockNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNamePrefix

`func (o *RbdImage) SetBlockNamePrefix(v string)`

SetBlockNamePrefix sets BlockNamePrefix field to given value.

### HasBlockNamePrefix

`func (o *RbdImage) HasBlockNamePrefix() bool`

HasBlockNamePrefix returns a boolean if a field has been set.

### GetName

`func (o *RbdImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RbdImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RbdImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RbdImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUniqueId

`func (o *RbdImage) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *RbdImage) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *RbdImage) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *RbdImage) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetId

`func (o *RbdImage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RbdImage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RbdImage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RbdImage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageFormat

`func (o *RbdImage) GetImageFormat() int32`

GetImageFormat returns the ImageFormat field if non-nil, zero value otherwise.

### GetImageFormatOk

`func (o *RbdImage) GetImageFormatOk() (*int32, bool)`

GetImageFormatOk returns a tuple with the ImageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFormat

`func (o *RbdImage) SetImageFormat(v int32)`

SetImageFormat sets ImageFormat field to given value.

### HasImageFormat

`func (o *RbdImage) HasImageFormat() bool`

HasImageFormat returns a boolean if a field has been set.

### GetPoolName

`func (o *RbdImage) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *RbdImage) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *RbdImage) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *RbdImage) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetNamespace

`func (o *RbdImage) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RbdImage) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RbdImage) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RbdImage) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetFeatures

`func (o *RbdImage) GetFeatures() int32`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *RbdImage) GetFeaturesOk() (*int32, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *RbdImage) SetFeatures(v int32)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *RbdImage) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesName

`func (o *RbdImage) GetFeaturesName() []string`

GetFeaturesName returns the FeaturesName field if non-nil, zero value otherwise.

### GetFeaturesNameOk

`func (o *RbdImage) GetFeaturesNameOk() (*[]string, bool)`

GetFeaturesNameOk returns a tuple with the FeaturesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesName

`func (o *RbdImage) SetFeaturesName(v []string)`

SetFeaturesName sets FeaturesName field to given value.

### HasFeaturesName

`func (o *RbdImage) HasFeaturesName() bool`

HasFeaturesName returns a boolean if a field has been set.

### GetTimestamp

`func (o *RbdImage) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RbdImage) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RbdImage) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RbdImage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStripeCount

`func (o *RbdImage) GetStripeCount() int32`

GetStripeCount returns the StripeCount field if non-nil, zero value otherwise.

### GetStripeCountOk

`func (o *RbdImage) GetStripeCountOk() (*int32, bool)`

GetStripeCountOk returns a tuple with the StripeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCount

`func (o *RbdImage) SetStripeCount(v int32)`

SetStripeCount sets StripeCount field to given value.

### HasStripeCount

`func (o *RbdImage) HasStripeCount() bool`

HasStripeCount returns a boolean if a field has been set.

### GetStripeUnit

`func (o *RbdImage) GetStripeUnit() int32`

GetStripeUnit returns the StripeUnit field if non-nil, zero value otherwise.

### GetStripeUnitOk

`func (o *RbdImage) GetStripeUnitOk() (*int32, bool)`

GetStripeUnitOk returns a tuple with the StripeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeUnit

`func (o *RbdImage) SetStripeUnit(v int32)`

SetStripeUnit sets StripeUnit field to given value.

### HasStripeUnit

`func (o *RbdImage) HasStripeUnit() bool`

HasStripeUnit returns a boolean if a field has been set.

### GetDataPool

`func (o *RbdImage) GetDataPool() string`

GetDataPool returns the DataPool field if non-nil, zero value otherwise.

### GetDataPoolOk

`func (o *RbdImage) GetDataPoolOk() (*string, bool)`

GetDataPoolOk returns a tuple with the DataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPool

`func (o *RbdImage) SetDataPool(v string)`

SetDataPool sets DataPool field to given value.

### HasDataPool

`func (o *RbdImage) HasDataPool() bool`

HasDataPool returns a boolean if a field has been set.

### SetDataPoolNil

`func (o *RbdImage) SetDataPoolNil(b bool)`

 SetDataPoolNil sets the value for DataPool to be an explicit nil

### UnsetDataPool
`func (o *RbdImage) UnsetDataPool()`

UnsetDataPool ensures that no value is present for DataPool, not even an explicit nil
### GetParent

`func (o *RbdImage) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RbdImage) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RbdImage) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RbdImage) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *RbdImage) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *RbdImage) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetSnapshots

`func (o *RbdImage) GetSnapshots() []RbdImageSnapshotsInner`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RbdImage) GetSnapshotsOk() (*[]RbdImageSnapshotsInner, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RbdImage) SetSnapshots(v []RbdImageSnapshotsInner)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *RbdImage) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetTotalDiskUsage

`func (o *RbdImage) GetTotalDiskUsage() float32`

GetTotalDiskUsage returns the TotalDiskUsage field if non-nil, zero value otherwise.

### GetTotalDiskUsageOk

`func (o *RbdImage) GetTotalDiskUsageOk() (*float32, bool)`

GetTotalDiskUsageOk returns a tuple with the TotalDiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiskUsage

`func (o *RbdImage) SetTotalDiskUsage(v float32)`

SetTotalDiskUsage sets TotalDiskUsage field to given value.

### HasTotalDiskUsage

`func (o *RbdImage) HasTotalDiskUsage() bool`

HasTotalDiskUsage returns a boolean if a field has been set.

### GetDiskUsage

`func (o *RbdImage) GetDiskUsage() int32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *RbdImage) GetDiskUsageOk() (*int32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *RbdImage) SetDiskUsage(v int32)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *RbdImage) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.

### GetConfiguration

`func (o *RbdImage) GetConfiguration() []RbdImageConfigurationInner`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RbdImage) GetConfigurationOk() (*[]RbdImageConfigurationInner, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RbdImage) SetConfiguration(v []RbdImageConfigurationInner)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RbdImage) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


