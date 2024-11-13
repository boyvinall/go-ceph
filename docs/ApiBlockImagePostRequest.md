# ApiBlockImagePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **string** |  | [optional] 
**DataPool** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**MirrorMode** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 
**ObjSize** | Pointer to **int32** |  | [optional] 
**PoolName** | **string** |  | 
**ScheduleInterval** | Pointer to **string** |  | [optional] [default to ""]
**Size** | **int32** |  | 
**StripeCount** | Pointer to **int32** |  | [optional] 
**StripeUnit** | Pointer to **string** |  | [optional] 

## Methods

### NewApiBlockImagePostRequest

`func NewApiBlockImagePostRequest(name string, poolName string, size int32, ) *ApiBlockImagePostRequest`

NewApiBlockImagePostRequest instantiates a new ApiBlockImagePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBlockImagePostRequestWithDefaults

`func NewApiBlockImagePostRequestWithDefaults() *ApiBlockImagePostRequest`

NewApiBlockImagePostRequestWithDefaults instantiates a new ApiBlockImagePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *ApiBlockImagePostRequest) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ApiBlockImagePostRequest) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ApiBlockImagePostRequest) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ApiBlockImagePostRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDataPool

`func (o *ApiBlockImagePostRequest) GetDataPool() string`

GetDataPool returns the DataPool field if non-nil, zero value otherwise.

### GetDataPoolOk

`func (o *ApiBlockImagePostRequest) GetDataPoolOk() (*string, bool)`

GetDataPoolOk returns a tuple with the DataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPool

`func (o *ApiBlockImagePostRequest) SetDataPool(v string)`

SetDataPool sets DataPool field to given value.

### HasDataPool

`func (o *ApiBlockImagePostRequest) HasDataPool() bool`

HasDataPool returns a boolean if a field has been set.

### GetFeatures

`func (o *ApiBlockImagePostRequest) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ApiBlockImagePostRequest) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ApiBlockImagePostRequest) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ApiBlockImagePostRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetMetadata

`func (o *ApiBlockImagePostRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApiBlockImagePostRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApiBlockImagePostRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApiBlockImagePostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMirrorMode

`func (o *ApiBlockImagePostRequest) GetMirrorMode() string`

GetMirrorMode returns the MirrorMode field if non-nil, zero value otherwise.

### GetMirrorModeOk

`func (o *ApiBlockImagePostRequest) GetMirrorModeOk() (*string, bool)`

GetMirrorModeOk returns a tuple with the MirrorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorMode

`func (o *ApiBlockImagePostRequest) SetMirrorMode(v string)`

SetMirrorMode sets MirrorMode field to given value.

### HasMirrorMode

`func (o *ApiBlockImagePostRequest) HasMirrorMode() bool`

HasMirrorMode returns a boolean if a field has been set.

### GetName

`func (o *ApiBlockImagePostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiBlockImagePostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiBlockImagePostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *ApiBlockImagePostRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiBlockImagePostRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiBlockImagePostRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApiBlockImagePostRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetObjSize

`func (o *ApiBlockImagePostRequest) GetObjSize() int32`

GetObjSize returns the ObjSize field if non-nil, zero value otherwise.

### GetObjSizeOk

`func (o *ApiBlockImagePostRequest) GetObjSizeOk() (*int32, bool)`

GetObjSizeOk returns a tuple with the ObjSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSize

`func (o *ApiBlockImagePostRequest) SetObjSize(v int32)`

SetObjSize sets ObjSize field to given value.

### HasObjSize

`func (o *ApiBlockImagePostRequest) HasObjSize() bool`

HasObjSize returns a boolean if a field has been set.

### GetPoolName

`func (o *ApiBlockImagePostRequest) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *ApiBlockImagePostRequest) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *ApiBlockImagePostRequest) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.


### GetScheduleInterval

`func (o *ApiBlockImagePostRequest) GetScheduleInterval() string`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *ApiBlockImagePostRequest) GetScheduleIntervalOk() (*string, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *ApiBlockImagePostRequest) SetScheduleInterval(v string)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *ApiBlockImagePostRequest) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetSize

`func (o *ApiBlockImagePostRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApiBlockImagePostRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApiBlockImagePostRequest) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStripeCount

`func (o *ApiBlockImagePostRequest) GetStripeCount() int32`

GetStripeCount returns the StripeCount field if non-nil, zero value otherwise.

### GetStripeCountOk

`func (o *ApiBlockImagePostRequest) GetStripeCountOk() (*int32, bool)`

GetStripeCountOk returns a tuple with the StripeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCount

`func (o *ApiBlockImagePostRequest) SetStripeCount(v int32)`

SetStripeCount sets StripeCount field to given value.

### HasStripeCount

`func (o *ApiBlockImagePostRequest) HasStripeCount() bool`

HasStripeCount returns a boolean if a field has been set.

### GetStripeUnit

`func (o *ApiBlockImagePostRequest) GetStripeUnit() string`

GetStripeUnit returns the StripeUnit field if non-nil, zero value otherwise.

### GetStripeUnitOk

`func (o *ApiBlockImagePostRequest) GetStripeUnitOk() (*string, bool)`

GetStripeUnitOk returns a tuple with the StripeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeUnit

`func (o *ApiBlockImagePostRequest) SetStripeUnit(v string)`

SetStripeUnit sets StripeUnit field to given value.

### HasStripeUnit

`func (o *ApiBlockImagePostRequest) HasStripeUnit() bool`

HasStripeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


