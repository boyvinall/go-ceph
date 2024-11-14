# ApiBlockImageImageSpecGet200Response

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
**Snapshots** | Pointer to [**[]ApiBlockImageImageSpecGet200ResponseSnapshotsInner**](ApiBlockImageImageSpecGet200ResponseSnapshotsInner.md) |  | [optional] 
**TotalDiskUsage** | Pointer to **float32** |  | [optional] 
**DiskUsage** | Pointer to **int32** |  | [optional] 
**Configuration** | Pointer to [**[]ApiBlockImageImageSpecGet200ResponseConfigurationInner**](ApiBlockImageImageSpecGet200ResponseConfigurationInner.md) |  | [optional] 

## Methods

### NewApiBlockImageImageSpecGet200Response

`func NewApiBlockImageImageSpecGet200Response() *ApiBlockImageImageSpecGet200Response`

NewApiBlockImageImageSpecGet200Response instantiates a new ApiBlockImageImageSpecGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBlockImageImageSpecGet200ResponseWithDefaults

`func NewApiBlockImageImageSpecGet200ResponseWithDefaults() *ApiBlockImageImageSpecGet200Response`

NewApiBlockImageImageSpecGet200ResponseWithDefaults instantiates a new ApiBlockImageImageSpecGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *ApiBlockImageImageSpecGet200Response) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApiBlockImageImageSpecGet200Response) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApiBlockImageImageSpecGet200Response) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ApiBlockImageImageSpecGet200Response) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetObjSize

`func (o *ApiBlockImageImageSpecGet200Response) GetObjSize() int32`

GetObjSize returns the ObjSize field if non-nil, zero value otherwise.

### GetObjSizeOk

`func (o *ApiBlockImageImageSpecGet200Response) GetObjSizeOk() (*int32, bool)`

GetObjSizeOk returns a tuple with the ObjSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSize

`func (o *ApiBlockImageImageSpecGet200Response) SetObjSize(v int32)`

SetObjSize sets ObjSize field to given value.

### HasObjSize

`func (o *ApiBlockImageImageSpecGet200Response) HasObjSize() bool`

HasObjSize returns a boolean if a field has been set.

### GetNumObjs

`func (o *ApiBlockImageImageSpecGet200Response) GetNumObjs() int32`

GetNumObjs returns the NumObjs field if non-nil, zero value otherwise.

### GetNumObjsOk

`func (o *ApiBlockImageImageSpecGet200Response) GetNumObjsOk() (*int32, bool)`

GetNumObjsOk returns a tuple with the NumObjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjs

`func (o *ApiBlockImageImageSpecGet200Response) SetNumObjs(v int32)`

SetNumObjs sets NumObjs field to given value.

### HasNumObjs

`func (o *ApiBlockImageImageSpecGet200Response) HasNumObjs() bool`

HasNumObjs returns a boolean if a field has been set.

### GetOrder

`func (o *ApiBlockImageImageSpecGet200Response) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApiBlockImageImageSpecGet200Response) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApiBlockImageImageSpecGet200Response) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApiBlockImageImageSpecGet200Response) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetBlockNamePrefix

`func (o *ApiBlockImageImageSpecGet200Response) GetBlockNamePrefix() string`

GetBlockNamePrefix returns the BlockNamePrefix field if non-nil, zero value otherwise.

### GetBlockNamePrefixOk

`func (o *ApiBlockImageImageSpecGet200Response) GetBlockNamePrefixOk() (*string, bool)`

GetBlockNamePrefixOk returns a tuple with the BlockNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNamePrefix

`func (o *ApiBlockImageImageSpecGet200Response) SetBlockNamePrefix(v string)`

SetBlockNamePrefix sets BlockNamePrefix field to given value.

### HasBlockNamePrefix

`func (o *ApiBlockImageImageSpecGet200Response) HasBlockNamePrefix() bool`

HasBlockNamePrefix returns a boolean if a field has been set.

### GetName

`func (o *ApiBlockImageImageSpecGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiBlockImageImageSpecGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiBlockImageImageSpecGet200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiBlockImageImageSpecGet200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUniqueId

`func (o *ApiBlockImageImageSpecGet200Response) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ApiBlockImageImageSpecGet200Response) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ApiBlockImageImageSpecGet200Response) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ApiBlockImageImageSpecGet200Response) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetId

`func (o *ApiBlockImageImageSpecGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiBlockImageImageSpecGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiBlockImageImageSpecGet200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiBlockImageImageSpecGet200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageFormat

`func (o *ApiBlockImageImageSpecGet200Response) GetImageFormat() int32`

GetImageFormat returns the ImageFormat field if non-nil, zero value otherwise.

### GetImageFormatOk

`func (o *ApiBlockImageImageSpecGet200Response) GetImageFormatOk() (*int32, bool)`

GetImageFormatOk returns a tuple with the ImageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFormat

`func (o *ApiBlockImageImageSpecGet200Response) SetImageFormat(v int32)`

SetImageFormat sets ImageFormat field to given value.

### HasImageFormat

`func (o *ApiBlockImageImageSpecGet200Response) HasImageFormat() bool`

HasImageFormat returns a boolean if a field has been set.

### GetPoolName

`func (o *ApiBlockImageImageSpecGet200Response) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *ApiBlockImageImageSpecGet200Response) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *ApiBlockImageImageSpecGet200Response) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *ApiBlockImageImageSpecGet200Response) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetNamespace

`func (o *ApiBlockImageImageSpecGet200Response) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiBlockImageImageSpecGet200Response) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiBlockImageImageSpecGet200Response) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApiBlockImageImageSpecGet200Response) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetFeatures

`func (o *ApiBlockImageImageSpecGet200Response) GetFeatures() int32`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ApiBlockImageImageSpecGet200Response) GetFeaturesOk() (*int32, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ApiBlockImageImageSpecGet200Response) SetFeatures(v int32)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ApiBlockImageImageSpecGet200Response) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesName

`func (o *ApiBlockImageImageSpecGet200Response) GetFeaturesName() []string`

GetFeaturesName returns the FeaturesName field if non-nil, zero value otherwise.

### GetFeaturesNameOk

`func (o *ApiBlockImageImageSpecGet200Response) GetFeaturesNameOk() (*[]string, bool)`

GetFeaturesNameOk returns a tuple with the FeaturesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesName

`func (o *ApiBlockImageImageSpecGet200Response) SetFeaturesName(v []string)`

SetFeaturesName sets FeaturesName field to given value.

### HasFeaturesName

`func (o *ApiBlockImageImageSpecGet200Response) HasFeaturesName() bool`

HasFeaturesName returns a boolean if a field has been set.

### GetTimestamp

`func (o *ApiBlockImageImageSpecGet200Response) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiBlockImageImageSpecGet200Response) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiBlockImageImageSpecGet200Response) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiBlockImageImageSpecGet200Response) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStripeCount

`func (o *ApiBlockImageImageSpecGet200Response) GetStripeCount() int32`

GetStripeCount returns the StripeCount field if non-nil, zero value otherwise.

### GetStripeCountOk

`func (o *ApiBlockImageImageSpecGet200Response) GetStripeCountOk() (*int32, bool)`

GetStripeCountOk returns a tuple with the StripeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCount

`func (o *ApiBlockImageImageSpecGet200Response) SetStripeCount(v int32)`

SetStripeCount sets StripeCount field to given value.

### HasStripeCount

`func (o *ApiBlockImageImageSpecGet200Response) HasStripeCount() bool`

HasStripeCount returns a boolean if a field has been set.

### GetStripeUnit

`func (o *ApiBlockImageImageSpecGet200Response) GetStripeUnit() int32`

GetStripeUnit returns the StripeUnit field if non-nil, zero value otherwise.

### GetStripeUnitOk

`func (o *ApiBlockImageImageSpecGet200Response) GetStripeUnitOk() (*int32, bool)`

GetStripeUnitOk returns a tuple with the StripeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeUnit

`func (o *ApiBlockImageImageSpecGet200Response) SetStripeUnit(v int32)`

SetStripeUnit sets StripeUnit field to given value.

### HasStripeUnit

`func (o *ApiBlockImageImageSpecGet200Response) HasStripeUnit() bool`

HasStripeUnit returns a boolean if a field has been set.

### GetDataPool

`func (o *ApiBlockImageImageSpecGet200Response) GetDataPool() string`

GetDataPool returns the DataPool field if non-nil, zero value otherwise.

### GetDataPoolOk

`func (o *ApiBlockImageImageSpecGet200Response) GetDataPoolOk() (*string, bool)`

GetDataPoolOk returns a tuple with the DataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPool

`func (o *ApiBlockImageImageSpecGet200Response) SetDataPool(v string)`

SetDataPool sets DataPool field to given value.

### HasDataPool

`func (o *ApiBlockImageImageSpecGet200Response) HasDataPool() bool`

HasDataPool returns a boolean if a field has been set.

### SetDataPoolNil

`func (o *ApiBlockImageImageSpecGet200Response) SetDataPoolNil(b bool)`

 SetDataPoolNil sets the value for DataPool to be an explicit nil

### UnsetDataPool
`func (o *ApiBlockImageImageSpecGet200Response) UnsetDataPool()`

UnsetDataPool ensures that no value is present for DataPool, not even an explicit nil
### GetParent

`func (o *ApiBlockImageImageSpecGet200Response) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ApiBlockImageImageSpecGet200Response) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ApiBlockImageImageSpecGet200Response) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ApiBlockImageImageSpecGet200Response) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *ApiBlockImageImageSpecGet200Response) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *ApiBlockImageImageSpecGet200Response) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetSnapshots

`func (o *ApiBlockImageImageSpecGet200Response) GetSnapshots() []ApiBlockImageImageSpecGet200ResponseSnapshotsInner`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ApiBlockImageImageSpecGet200Response) GetSnapshotsOk() (*[]ApiBlockImageImageSpecGet200ResponseSnapshotsInner, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ApiBlockImageImageSpecGet200Response) SetSnapshots(v []ApiBlockImageImageSpecGet200ResponseSnapshotsInner)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *ApiBlockImageImageSpecGet200Response) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetTotalDiskUsage

`func (o *ApiBlockImageImageSpecGet200Response) GetTotalDiskUsage() float32`

GetTotalDiskUsage returns the TotalDiskUsage field if non-nil, zero value otherwise.

### GetTotalDiskUsageOk

`func (o *ApiBlockImageImageSpecGet200Response) GetTotalDiskUsageOk() (*float32, bool)`

GetTotalDiskUsageOk returns a tuple with the TotalDiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiskUsage

`func (o *ApiBlockImageImageSpecGet200Response) SetTotalDiskUsage(v float32)`

SetTotalDiskUsage sets TotalDiskUsage field to given value.

### HasTotalDiskUsage

`func (o *ApiBlockImageImageSpecGet200Response) HasTotalDiskUsage() bool`

HasTotalDiskUsage returns a boolean if a field has been set.

### GetDiskUsage

`func (o *ApiBlockImageImageSpecGet200Response) GetDiskUsage() int32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *ApiBlockImageImageSpecGet200Response) GetDiskUsageOk() (*int32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *ApiBlockImageImageSpecGet200Response) SetDiskUsage(v int32)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *ApiBlockImageImageSpecGet200Response) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.

### GetConfiguration

`func (o *ApiBlockImageImageSpecGet200Response) GetConfiguration() []ApiBlockImageImageSpecGet200ResponseConfigurationInner`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ApiBlockImageImageSpecGet200Response) GetConfigurationOk() (*[]ApiBlockImageImageSpecGet200ResponseConfigurationInner, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ApiBlockImageImageSpecGet200Response) SetConfiguration(v []ApiBlockImageImageSpecGet200ResponseConfigurationInner)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ApiBlockImageImageSpecGet200Response) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


