# ApiTelemetryReportGet200ResponseReportCrush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketAlgs** | [**ApiTelemetryReportGet200ResponseReportCrushBucketAlgs**](ApiTelemetryReportGet200ResponseReportCrushBucketAlgs.md) |  | 
**BucketSizes** | [**ApiTelemetryReportGet200ResponseReportCrushBucketSizes**](ApiTelemetryReportGet200ResponseReportCrushBucketSizes.md) |  | 
**BucketTypes** | [**ApiTelemetryReportGet200ResponseReportCrushBucketTypes**](ApiTelemetryReportGet200ResponseReportCrushBucketTypes.md) |  | 
**CompatWeightSet** | **bool** |  | 
**DeviceClasses** | **[]int32** |  | 
**NumBuckets** | **int32** |  | 
**NumDevices** | **int32** |  | 
**NumRules** | **int32** |  | 
**NumTypes** | **int32** |  | 
**NumWeightSets** | **int32** |  | 
**Tunables** | [**ApiTelemetryReportGet200ResponseReportCrushTunables**](ApiTelemetryReportGet200ResponseReportCrushTunables.md) |  | 

## Methods

### NewApiTelemetryReportGet200ResponseReportCrush

`func NewApiTelemetryReportGet200ResponseReportCrush(bucketAlgs ApiTelemetryReportGet200ResponseReportCrushBucketAlgs, bucketSizes ApiTelemetryReportGet200ResponseReportCrushBucketSizes, bucketTypes ApiTelemetryReportGet200ResponseReportCrushBucketTypes, compatWeightSet bool, deviceClasses []int32, numBuckets int32, numDevices int32, numRules int32, numTypes int32, numWeightSets int32, tunables ApiTelemetryReportGet200ResponseReportCrushTunables, ) *ApiTelemetryReportGet200ResponseReportCrush`

NewApiTelemetryReportGet200ResponseReportCrush instantiates a new ApiTelemetryReportGet200ResponseReportCrush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTelemetryReportGet200ResponseReportCrushWithDefaults

`func NewApiTelemetryReportGet200ResponseReportCrushWithDefaults() *ApiTelemetryReportGet200ResponseReportCrush`

NewApiTelemetryReportGet200ResponseReportCrushWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportCrush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketAlgs

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketAlgs() ApiTelemetryReportGet200ResponseReportCrushBucketAlgs`

GetBucketAlgs returns the BucketAlgs field if non-nil, zero value otherwise.

### GetBucketAlgsOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketAlgsOk() (*ApiTelemetryReportGet200ResponseReportCrushBucketAlgs, bool)`

GetBucketAlgsOk returns a tuple with the BucketAlgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketAlgs

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetBucketAlgs(v ApiTelemetryReportGet200ResponseReportCrushBucketAlgs)`

SetBucketAlgs sets BucketAlgs field to given value.


### GetBucketSizes

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketSizes() ApiTelemetryReportGet200ResponseReportCrushBucketSizes`

GetBucketSizes returns the BucketSizes field if non-nil, zero value otherwise.

### GetBucketSizesOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketSizesOk() (*ApiTelemetryReportGet200ResponseReportCrushBucketSizes, bool)`

GetBucketSizesOk returns a tuple with the BucketSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSizes

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetBucketSizes(v ApiTelemetryReportGet200ResponseReportCrushBucketSizes)`

SetBucketSizes sets BucketSizes field to given value.


### GetBucketTypes

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketTypes() ApiTelemetryReportGet200ResponseReportCrushBucketTypes`

GetBucketTypes returns the BucketTypes field if non-nil, zero value otherwise.

### GetBucketTypesOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketTypesOk() (*ApiTelemetryReportGet200ResponseReportCrushBucketTypes, bool)`

GetBucketTypesOk returns a tuple with the BucketTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketTypes

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetBucketTypes(v ApiTelemetryReportGet200ResponseReportCrushBucketTypes)`

SetBucketTypes sets BucketTypes field to given value.


### GetCompatWeightSet

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetCompatWeightSet() bool`

GetCompatWeightSet returns the CompatWeightSet field if non-nil, zero value otherwise.

### GetCompatWeightSetOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetCompatWeightSetOk() (*bool, bool)`

GetCompatWeightSetOk returns a tuple with the CompatWeightSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatWeightSet

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetCompatWeightSet(v bool)`

SetCompatWeightSet sets CompatWeightSet field to given value.


### GetDeviceClasses

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetDeviceClasses() []int32`

GetDeviceClasses returns the DeviceClasses field if non-nil, zero value otherwise.

### GetDeviceClassesOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetDeviceClassesOk() (*[]int32, bool)`

GetDeviceClassesOk returns a tuple with the DeviceClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClasses

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetDeviceClasses(v []int32)`

SetDeviceClasses sets DeviceClasses field to given value.


### GetNumBuckets

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumBuckets() int32`

GetNumBuckets returns the NumBuckets field if non-nil, zero value otherwise.

### GetNumBucketsOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumBucketsOk() (*int32, bool)`

GetNumBucketsOk returns a tuple with the NumBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBuckets

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetNumBuckets(v int32)`

SetNumBuckets sets NumBuckets field to given value.


### GetNumDevices

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumDevices() int32`

GetNumDevices returns the NumDevices field if non-nil, zero value otherwise.

### GetNumDevicesOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumDevicesOk() (*int32, bool)`

GetNumDevicesOk returns a tuple with the NumDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDevices

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetNumDevices(v int32)`

SetNumDevices sets NumDevices field to given value.


### GetNumRules

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumRules() int32`

GetNumRules returns the NumRules field if non-nil, zero value otherwise.

### GetNumRulesOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumRulesOk() (*int32, bool)`

GetNumRulesOk returns a tuple with the NumRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRules

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetNumRules(v int32)`

SetNumRules sets NumRules field to given value.


### GetNumTypes

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumTypes() int32`

GetNumTypes returns the NumTypes field if non-nil, zero value otherwise.

### GetNumTypesOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumTypesOk() (*int32, bool)`

GetNumTypesOk returns a tuple with the NumTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTypes

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetNumTypes(v int32)`

SetNumTypes sets NumTypes field to given value.


### GetNumWeightSets

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumWeightSets() int32`

GetNumWeightSets returns the NumWeightSets field if non-nil, zero value otherwise.

### GetNumWeightSetsOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumWeightSetsOk() (*int32, bool)`

GetNumWeightSetsOk returns a tuple with the NumWeightSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWeightSets

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetNumWeightSets(v int32)`

SetNumWeightSets sets NumWeightSets field to given value.


### GetTunables

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetTunables() ApiTelemetryReportGet200ResponseReportCrushTunables`

GetTunables returns the Tunables field if non-nil, zero value otherwise.

### GetTunablesOk

`func (o *ApiTelemetryReportGet200ResponseReportCrush) GetTunablesOk() (*ApiTelemetryReportGet200ResponseReportCrushTunables, bool)`

GetTunablesOk returns a tuple with the Tunables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunables

`func (o *ApiTelemetryReportGet200ResponseReportCrush) SetTunables(v ApiTelemetryReportGet200ResponseReportCrushTunables)`

SetTunables sets Tunables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


