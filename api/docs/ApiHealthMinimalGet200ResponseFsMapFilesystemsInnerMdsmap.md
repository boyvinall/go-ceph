# ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balancer** | **string** |  | 
**Btime** | **string** |  | 
**Compat** | [**ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmapCompat**](ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmapCompat.md) |  | 
**Created** | **string** |  | 
**Damaged** | **[]int32** |  | 
**DataPools** | **[]int32** |  | 
**Enabled** | **bool** |  | 
**Epoch** | **int32** |  | 
**EverAllowedFeatures** | **int32** |  | 
**ExplicitlyAllowedFeatures** | **int32** |  | 
**Failed** | **[]int32** |  | 
**Flags** | **int32** |  | 
**FsName** | **string** |  | 
**In** | **[]int32** |  | 
**Info** | **string** |  | 
**LastFailure** | **int32** |  | 
**LastFailureOsdEpoch** | **int32** |  | 
**MaxFileSize** | **int32** |  | 
**MaxMds** | **int32** |  | 
**MetadataPool** | **int32** |  | 
**Modified** | **string** |  | 
**RequiredClientFeatures** | **string** |  | 
**Root** | **int32** |  | 
**SessionAutoclose** | **int32** |  | 
**SessionTimeout** | **int32** |  | 
**StandbyCountWanted** | **int32** |  | 
**Stopped** | **[]int32** |  | 
**Tableserver** | **int32** |  | 
**Up** | **string** |  | 

## Methods

### NewApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap

`func NewApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap(balancer string, btime string, compat ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmapCompat, created string, damaged []int32, dataPools []int32, enabled bool, epoch int32, everAllowedFeatures int32, explicitlyAllowedFeatures int32, failed []int32, flags int32, fsName string, in []int32, info string, lastFailure int32, lastFailureOsdEpoch int32, maxFileSize int32, maxMds int32, metadataPool int32, modified string, requiredClientFeatures string, root int32, sessionAutoclose int32, sessionTimeout int32, standbyCountWanted int32, stopped []int32, tableserver int32, up string, ) *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap`

NewApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap instantiates a new ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmapWithDefaults

`func NewApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmapWithDefaults() *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap`

NewApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmapWithDefaults instantiates a new ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalancer

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetBalancer() string`

GetBalancer returns the Balancer field if non-nil, zero value otherwise.

### GetBalancerOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetBalancerOk() (*string, bool)`

GetBalancerOk returns a tuple with the Balancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancer

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetBalancer(v string)`

SetBalancer sets Balancer field to given value.


### GetBtime

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetBtime() string`

GetBtime returns the Btime field if non-nil, zero value otherwise.

### GetBtimeOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetBtimeOk() (*string, bool)`

GetBtimeOk returns a tuple with the Btime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtime

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetBtime(v string)`

SetBtime sets Btime field to given value.


### GetCompat

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetCompat() ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmapCompat`

GetCompat returns the Compat field if non-nil, zero value otherwise.

### GetCompatOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetCompatOk() (*ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmapCompat, bool)`

GetCompatOk returns a tuple with the Compat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompat

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetCompat(v ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmapCompat)`

SetCompat sets Compat field to given value.


### GetCreated

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetDamaged

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetDamaged() []int32`

GetDamaged returns the Damaged field if non-nil, zero value otherwise.

### GetDamagedOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetDamagedOk() (*[]int32, bool)`

GetDamagedOk returns a tuple with the Damaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamaged

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetDamaged(v []int32)`

SetDamaged sets Damaged field to given value.


### GetDataPools

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetDataPools() []int32`

GetDataPools returns the DataPools field if non-nil, zero value otherwise.

### GetDataPoolsOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetDataPoolsOk() (*[]int32, bool)`

GetDataPoolsOk returns a tuple with the DataPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPools

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetDataPools(v []int32)`

SetDataPools sets DataPools field to given value.


### GetEnabled

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEpoch

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetEpoch() int32`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetEpochOk() (*int32, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetEpoch(v int32)`

SetEpoch sets Epoch field to given value.


### GetEverAllowedFeatures

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetEverAllowedFeatures() int32`

GetEverAllowedFeatures returns the EverAllowedFeatures field if non-nil, zero value otherwise.

### GetEverAllowedFeaturesOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetEverAllowedFeaturesOk() (*int32, bool)`

GetEverAllowedFeaturesOk returns a tuple with the EverAllowedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEverAllowedFeatures

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetEverAllowedFeatures(v int32)`

SetEverAllowedFeatures sets EverAllowedFeatures field to given value.


### GetExplicitlyAllowedFeatures

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetExplicitlyAllowedFeatures() int32`

GetExplicitlyAllowedFeatures returns the ExplicitlyAllowedFeatures field if non-nil, zero value otherwise.

### GetExplicitlyAllowedFeaturesOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetExplicitlyAllowedFeaturesOk() (*int32, bool)`

GetExplicitlyAllowedFeaturesOk returns a tuple with the ExplicitlyAllowedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitlyAllowedFeatures

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetExplicitlyAllowedFeatures(v int32)`

SetExplicitlyAllowedFeatures sets ExplicitlyAllowedFeatures field to given value.


### GetFailed

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetFailed() []int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetFailedOk() (*[]int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetFailed(v []int32)`

SetFailed sets Failed field to given value.


### GetFlags

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetFsName

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetFsName() string`

GetFsName returns the FsName field if non-nil, zero value otherwise.

### GetFsNameOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetFsNameOk() (*string, bool)`

GetFsNameOk returns a tuple with the FsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsName

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetFsName(v string)`

SetFsName sets FsName field to given value.


### GetIn

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetIn() []int32`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetInOk() (*[]int32, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetIn(v []int32)`

SetIn sets In field to given value.


### GetInfo

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetInfo(v string)`

SetInfo sets Info field to given value.


### GetLastFailure

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetLastFailure() int32`

GetLastFailure returns the LastFailure field if non-nil, zero value otherwise.

### GetLastFailureOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetLastFailureOk() (*int32, bool)`

GetLastFailureOk returns a tuple with the LastFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailure

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetLastFailure(v int32)`

SetLastFailure sets LastFailure field to given value.


### GetLastFailureOsdEpoch

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetLastFailureOsdEpoch() int32`

GetLastFailureOsdEpoch returns the LastFailureOsdEpoch field if non-nil, zero value otherwise.

### GetLastFailureOsdEpochOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetLastFailureOsdEpochOk() (*int32, bool)`

GetLastFailureOsdEpochOk returns a tuple with the LastFailureOsdEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureOsdEpoch

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetLastFailureOsdEpoch(v int32)`

SetLastFailureOsdEpoch sets LastFailureOsdEpoch field to given value.


### GetMaxFileSize

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetMaxFileSize() int32`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetMaxFileSizeOk() (*int32, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetMaxFileSize(v int32)`

SetMaxFileSize sets MaxFileSize field to given value.


### GetMaxMds

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetMaxMds() int32`

GetMaxMds returns the MaxMds field if non-nil, zero value otherwise.

### GetMaxMdsOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetMaxMdsOk() (*int32, bool)`

GetMaxMdsOk returns a tuple with the MaxMds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMds

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetMaxMds(v int32)`

SetMaxMds sets MaxMds field to given value.


### GetMetadataPool

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetMetadataPool() int32`

GetMetadataPool returns the MetadataPool field if non-nil, zero value otherwise.

### GetMetadataPoolOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetMetadataPoolOk() (*int32, bool)`

GetMetadataPoolOk returns a tuple with the MetadataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataPool

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetMetadataPool(v int32)`

SetMetadataPool sets MetadataPool field to given value.


### GetModified

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetModified(v string)`

SetModified sets Modified field to given value.


### GetRequiredClientFeatures

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetRequiredClientFeatures() string`

GetRequiredClientFeatures returns the RequiredClientFeatures field if non-nil, zero value otherwise.

### GetRequiredClientFeaturesOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetRequiredClientFeaturesOk() (*string, bool)`

GetRequiredClientFeaturesOk returns a tuple with the RequiredClientFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredClientFeatures

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetRequiredClientFeatures(v string)`

SetRequiredClientFeatures sets RequiredClientFeatures field to given value.


### GetRoot

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetRoot() int32`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetRootOk() (*int32, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetRoot(v int32)`

SetRoot sets Root field to given value.


### GetSessionAutoclose

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetSessionAutoclose() int32`

GetSessionAutoclose returns the SessionAutoclose field if non-nil, zero value otherwise.

### GetSessionAutocloseOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetSessionAutocloseOk() (*int32, bool)`

GetSessionAutocloseOk returns a tuple with the SessionAutoclose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAutoclose

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetSessionAutoclose(v int32)`

SetSessionAutoclose sets SessionAutoclose field to given value.


### GetSessionTimeout

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.


### GetStandbyCountWanted

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetStandbyCountWanted() int32`

GetStandbyCountWanted returns the StandbyCountWanted field if non-nil, zero value otherwise.

### GetStandbyCountWantedOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetStandbyCountWantedOk() (*int32, bool)`

GetStandbyCountWantedOk returns a tuple with the StandbyCountWanted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyCountWanted

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetStandbyCountWanted(v int32)`

SetStandbyCountWanted sets StandbyCountWanted field to given value.


### GetStopped

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetStopped() []int32`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetStoppedOk() (*[]int32, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetStopped(v []int32)`

SetStopped sets Stopped field to given value.


### GetTableserver

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetTableserver() int32`

GetTableserver returns the Tableserver field if non-nil, zero value otherwise.

### GetTableserverOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetTableserverOk() (*int32, bool)`

GetTableserverOk returns a tuple with the Tableserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableserver

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetTableserver(v int32)`

SetTableserver sets Tableserver field to given value.


### GetUp

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetUp() string`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) GetUpOk() (*string, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) SetUp(v string)`

SetUp sets Up field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


