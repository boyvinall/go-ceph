# ApiBlockImageImageSpecCopyPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **string** |  | [optional] 
**DataPool** | Pointer to **string** |  | [optional] 
**DestImageName** | **string** |  | 
**DestNamespace** | **string** |  | 
**DestPoolName** | **string** |  | 
**Features** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**ObjSize** | Pointer to **int32** |  | [optional] 
**SnapshotName** | Pointer to **string** |  | [optional] 
**StripeCount** | Pointer to **int32** |  | [optional] 
**StripeUnit** | Pointer to **string** |  | [optional] 

## Methods

### NewApiBlockImageImageSpecCopyPostRequest

`func NewApiBlockImageImageSpecCopyPostRequest(destImageName string, destNamespace string, destPoolName string, ) *ApiBlockImageImageSpecCopyPostRequest`

NewApiBlockImageImageSpecCopyPostRequest instantiates a new ApiBlockImageImageSpecCopyPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBlockImageImageSpecCopyPostRequestWithDefaults

`func NewApiBlockImageImageSpecCopyPostRequestWithDefaults() *ApiBlockImageImageSpecCopyPostRequest`

NewApiBlockImageImageSpecCopyPostRequestWithDefaults instantiates a new ApiBlockImageImageSpecCopyPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ApiBlockImageImageSpecCopyPostRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDataPool

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetDataPool() string`

GetDataPool returns the DataPool field if non-nil, zero value otherwise.

### GetDataPoolOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetDataPoolOk() (*string, bool)`

GetDataPoolOk returns a tuple with the DataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPool

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetDataPool(v string)`

SetDataPool sets DataPool field to given value.

### HasDataPool

`func (o *ApiBlockImageImageSpecCopyPostRequest) HasDataPool() bool`

HasDataPool returns a boolean if a field has been set.

### GetDestImageName

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetDestImageName() string`

GetDestImageName returns the DestImageName field if non-nil, zero value otherwise.

### GetDestImageNameOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetDestImageNameOk() (*string, bool)`

GetDestImageNameOk returns a tuple with the DestImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestImageName

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetDestImageName(v string)`

SetDestImageName sets DestImageName field to given value.


### GetDestNamespace

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetDestNamespace() string`

GetDestNamespace returns the DestNamespace field if non-nil, zero value otherwise.

### GetDestNamespaceOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetDestNamespaceOk() (*string, bool)`

GetDestNamespaceOk returns a tuple with the DestNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestNamespace

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetDestNamespace(v string)`

SetDestNamespace sets DestNamespace field to given value.


### GetDestPoolName

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetDestPoolName() string`

GetDestPoolName returns the DestPoolName field if non-nil, zero value otherwise.

### GetDestPoolNameOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetDestPoolNameOk() (*string, bool)`

GetDestPoolNameOk returns a tuple with the DestPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPoolName

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetDestPoolName(v string)`

SetDestPoolName sets DestPoolName field to given value.


### GetFeatures

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ApiBlockImageImageSpecCopyPostRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetMetadata

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApiBlockImageImageSpecCopyPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetObjSize

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetObjSize() int32`

GetObjSize returns the ObjSize field if non-nil, zero value otherwise.

### GetObjSizeOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetObjSizeOk() (*int32, bool)`

GetObjSizeOk returns a tuple with the ObjSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSize

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetObjSize(v int32)`

SetObjSize sets ObjSize field to given value.

### HasObjSize

`func (o *ApiBlockImageImageSpecCopyPostRequest) HasObjSize() bool`

HasObjSize returns a boolean if a field has been set.

### GetSnapshotName

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetSnapshotName() string`

GetSnapshotName returns the SnapshotName field if non-nil, zero value otherwise.

### GetSnapshotNameOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetSnapshotNameOk() (*string, bool)`

GetSnapshotNameOk returns a tuple with the SnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotName

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetSnapshotName(v string)`

SetSnapshotName sets SnapshotName field to given value.

### HasSnapshotName

`func (o *ApiBlockImageImageSpecCopyPostRequest) HasSnapshotName() bool`

HasSnapshotName returns a boolean if a field has been set.

### GetStripeCount

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetStripeCount() int32`

GetStripeCount returns the StripeCount field if non-nil, zero value otherwise.

### GetStripeCountOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetStripeCountOk() (*int32, bool)`

GetStripeCountOk returns a tuple with the StripeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCount

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetStripeCount(v int32)`

SetStripeCount sets StripeCount field to given value.

### HasStripeCount

`func (o *ApiBlockImageImageSpecCopyPostRequest) HasStripeCount() bool`

HasStripeCount returns a boolean if a field has been set.

### GetStripeUnit

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetStripeUnit() string`

GetStripeUnit returns the StripeUnit field if non-nil, zero value otherwise.

### GetStripeUnitOk

`func (o *ApiBlockImageImageSpecCopyPostRequest) GetStripeUnitOk() (*string, bool)`

GetStripeUnitOk returns a tuple with the StripeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeUnit

`func (o *ApiBlockImageImageSpecCopyPostRequest) SetStripeUnit(v string)`

SetStripeUnit sets StripeUnit field to given value.

### HasStripeUnit

`func (o *ApiBlockImageImageSpecCopyPostRequest) HasStripeUnit() bool`

HasStripeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


