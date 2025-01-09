# ApiTelemetryReportGet200ResponseReportFs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**FeatureFlags** | [**ApiTelemetryReportGet200ResponseReportFsFeatureFlags**](ApiTelemetryReportGet200ResponseReportFsFeatureFlags.md) |  | 
**Filesystems** | **[]int32** |  | 
**NumStandbyMds** | **int32** |  | 
**TotalNumMds** | **int32** |  | 

## Methods

### NewApiTelemetryReportGet200ResponseReportFs

`func NewApiTelemetryReportGet200ResponseReportFs(count int32, featureFlags ApiTelemetryReportGet200ResponseReportFsFeatureFlags, filesystems []int32, numStandbyMds int32, totalNumMds int32, ) *ApiTelemetryReportGet200ResponseReportFs`

NewApiTelemetryReportGet200ResponseReportFs instantiates a new ApiTelemetryReportGet200ResponseReportFs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTelemetryReportGet200ResponseReportFsWithDefaults

`func NewApiTelemetryReportGet200ResponseReportFsWithDefaults() *ApiTelemetryReportGet200ResponseReportFs`

NewApiTelemetryReportGet200ResponseReportFsWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportFs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ApiTelemetryReportGet200ResponseReportFs) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApiTelemetryReportGet200ResponseReportFs) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApiTelemetryReportGet200ResponseReportFs) SetCount(v int32)`

SetCount sets Count field to given value.


### GetFeatureFlags

`func (o *ApiTelemetryReportGet200ResponseReportFs) GetFeatureFlags() ApiTelemetryReportGet200ResponseReportFsFeatureFlags`

GetFeatureFlags returns the FeatureFlags field if non-nil, zero value otherwise.

### GetFeatureFlagsOk

`func (o *ApiTelemetryReportGet200ResponseReportFs) GetFeatureFlagsOk() (*ApiTelemetryReportGet200ResponseReportFsFeatureFlags, bool)`

GetFeatureFlagsOk returns a tuple with the FeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlags

`func (o *ApiTelemetryReportGet200ResponseReportFs) SetFeatureFlags(v ApiTelemetryReportGet200ResponseReportFsFeatureFlags)`

SetFeatureFlags sets FeatureFlags field to given value.


### GetFilesystems

`func (o *ApiTelemetryReportGet200ResponseReportFs) GetFilesystems() []int32`

GetFilesystems returns the Filesystems field if non-nil, zero value otherwise.

### GetFilesystemsOk

`func (o *ApiTelemetryReportGet200ResponseReportFs) GetFilesystemsOk() (*[]int32, bool)`

GetFilesystemsOk returns a tuple with the Filesystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystems

`func (o *ApiTelemetryReportGet200ResponseReportFs) SetFilesystems(v []int32)`

SetFilesystems sets Filesystems field to given value.


### GetNumStandbyMds

`func (o *ApiTelemetryReportGet200ResponseReportFs) GetNumStandbyMds() int32`

GetNumStandbyMds returns the NumStandbyMds field if non-nil, zero value otherwise.

### GetNumStandbyMdsOk

`func (o *ApiTelemetryReportGet200ResponseReportFs) GetNumStandbyMdsOk() (*int32, bool)`

GetNumStandbyMdsOk returns a tuple with the NumStandbyMds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStandbyMds

`func (o *ApiTelemetryReportGet200ResponseReportFs) SetNumStandbyMds(v int32)`

SetNumStandbyMds sets NumStandbyMds field to given value.


### GetTotalNumMds

`func (o *ApiTelemetryReportGet200ResponseReportFs) GetTotalNumMds() int32`

GetTotalNumMds returns the TotalNumMds field if non-nil, zero value otherwise.

### GetTotalNumMdsOk

`func (o *ApiTelemetryReportGet200ResponseReportFs) GetTotalNumMdsOk() (*int32, bool)`

GetTotalNumMdsOk returns a tuple with the TotalNumMds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumMds

`func (o *ApiTelemetryReportGet200ResponseReportFs) SetTotalNumMds(v int32)`

SetTotalNumMds sets TotalNumMds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


