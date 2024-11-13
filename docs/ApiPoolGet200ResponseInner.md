# ApiPoolGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationMetadata** | Pointer to **[]string** |  | [optional] 
**Auid** | Pointer to **int32** |  | [optional] 
**CacheMinEvictAge** | Pointer to **int32** |  | [optional] 
**CacheMinFlushAge** | Pointer to **int32** |  | [optional] 
**CacheMode** | Pointer to **string** |  | [optional] 
**CacheTargetDirtyHighRatioMicro** | Pointer to **int32** |  | [optional] 
**CacheTargetDirtyRatioMicro** | Pointer to **int32** |  | [optional] 
**CacheTargetFullRatioMicro** | Pointer to **int32** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**CrushRule** | Pointer to **string** |  | [optional] 
**ErasureCodeProfile** | Pointer to **string** |  | [optional] 
**ExpectedNumObjects** | Pointer to **int32** |  | [optional] 
**FastRead** | Pointer to **bool** |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 
**FlagsNames** | Pointer to **string** | flags name | [optional] 
**GradeTable** | Pointer to **[]string** |  | [optional] 
**HitSetCount** | Pointer to **int32** |  | [optional] 
**HitSetGradeDecayRate** | Pointer to **int32** |  | [optional] 
**HitSetParams** | Pointer to [**ApiPoolGet200ResponseInnerHitSetParams**](ApiPoolGet200ResponseInnerHitSetParams.md) |  | [optional] 
**HitSetPeriod** | Pointer to **int32** |  | [optional] 
**HitSetSearchLastN** | Pointer to **int32** |  | [optional] 
**LastChange** | Pointer to **string** |  | [optional] 
**LastForceOpResend** | Pointer to **string** |  | [optional] 
**LastForceOpResendPreluminous** | Pointer to **string** |  | [optional] 
**LastForceOpResendPrenautilus** | Pointer to **string** |  | [optional] 
**LastPgMergeMeta** | Pointer to [**ApiPoolGet200ResponseInnerLastPgMergeMeta**](ApiPoolGet200ResponseInnerLastPgMergeMeta.md) |  | [optional] 
**MinReadRecencyForPromote** | Pointer to **int32** |  | [optional] 
**MinSize** | Pointer to **int32** |  | [optional] 
**MinWriteRecencyForPromote** | Pointer to **int32** |  | [optional] 
**ObjectHash** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**ApiPoolGet200ResponseInnerOptions**](ApiPoolGet200ResponseInnerOptions.md) |  | [optional] 
**PgAutoscaleMode** | Pointer to **string** |  | [optional] 
**PgNum** | Pointer to **int32** |  | [optional] 
**PgNumPending** | Pointer to **int32** |  | [optional] 
**PgNumTarget** | Pointer to **int32** |  | [optional] 
**PgPlacementNum** | Pointer to **int32** |  | [optional] 
**PgPlacementNumTarget** | Pointer to **int32** |  | [optional] 
**Pool** | Pointer to **int32** | pool id | [optional] 
**PoolName** | Pointer to **string** | pool name | [optional] 
**PoolSnaps** | Pointer to **[]string** |  | [optional] 
**QuotaMaxBytes** | Pointer to **int32** |  | [optional] 
**QuotaMaxObjects** | Pointer to **int32** |  | [optional] 
**ReadTier** | Pointer to **int32** |  | [optional] 
**RemovedSnaps** | Pointer to **[]string** |  | [optional] 
**Size** | Pointer to **int32** | pool size | [optional] 
**SnapEpoch** | Pointer to **int32** |  | [optional] 
**SnapMode** | Pointer to **string** |  | [optional] 
**SnapSeq** | Pointer to **int32** |  | [optional] 
**StripeWidth** | Pointer to **int32** |  | [optional] 
**TargetMaxBytes** | Pointer to **int32** |  | [optional] 
**TargetMaxObjects** | Pointer to **int32** |  | [optional] 
**TierOf** | Pointer to **int32** |  | [optional] 
**Tiers** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | type of pool | [optional] 
**UseGmtHitset** | Pointer to **bool** |  | [optional] 
**WriteTier** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiPoolGet200ResponseInner

`func NewApiPoolGet200ResponseInner() *ApiPoolGet200ResponseInner`

NewApiPoolGet200ResponseInner instantiates a new ApiPoolGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPoolGet200ResponseInnerWithDefaults

`func NewApiPoolGet200ResponseInnerWithDefaults() *ApiPoolGet200ResponseInner`

NewApiPoolGet200ResponseInnerWithDefaults instantiates a new ApiPoolGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationMetadata

`func (o *ApiPoolGet200ResponseInner) GetApplicationMetadata() []string`

GetApplicationMetadata returns the ApplicationMetadata field if non-nil, zero value otherwise.

### GetApplicationMetadataOk

`func (o *ApiPoolGet200ResponseInner) GetApplicationMetadataOk() (*[]string, bool)`

GetApplicationMetadataOk returns a tuple with the ApplicationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationMetadata

`func (o *ApiPoolGet200ResponseInner) SetApplicationMetadata(v []string)`

SetApplicationMetadata sets ApplicationMetadata field to given value.

### HasApplicationMetadata

`func (o *ApiPoolGet200ResponseInner) HasApplicationMetadata() bool`

HasApplicationMetadata returns a boolean if a field has been set.

### GetAuid

`func (o *ApiPoolGet200ResponseInner) GetAuid() int32`

GetAuid returns the Auid field if non-nil, zero value otherwise.

### GetAuidOk

`func (o *ApiPoolGet200ResponseInner) GetAuidOk() (*int32, bool)`

GetAuidOk returns a tuple with the Auid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuid

`func (o *ApiPoolGet200ResponseInner) SetAuid(v int32)`

SetAuid sets Auid field to given value.

### HasAuid

`func (o *ApiPoolGet200ResponseInner) HasAuid() bool`

HasAuid returns a boolean if a field has been set.

### GetCacheMinEvictAge

`func (o *ApiPoolGet200ResponseInner) GetCacheMinEvictAge() int32`

GetCacheMinEvictAge returns the CacheMinEvictAge field if non-nil, zero value otherwise.

### GetCacheMinEvictAgeOk

`func (o *ApiPoolGet200ResponseInner) GetCacheMinEvictAgeOk() (*int32, bool)`

GetCacheMinEvictAgeOk returns a tuple with the CacheMinEvictAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMinEvictAge

`func (o *ApiPoolGet200ResponseInner) SetCacheMinEvictAge(v int32)`

SetCacheMinEvictAge sets CacheMinEvictAge field to given value.

### HasCacheMinEvictAge

`func (o *ApiPoolGet200ResponseInner) HasCacheMinEvictAge() bool`

HasCacheMinEvictAge returns a boolean if a field has been set.

### GetCacheMinFlushAge

`func (o *ApiPoolGet200ResponseInner) GetCacheMinFlushAge() int32`

GetCacheMinFlushAge returns the CacheMinFlushAge field if non-nil, zero value otherwise.

### GetCacheMinFlushAgeOk

`func (o *ApiPoolGet200ResponseInner) GetCacheMinFlushAgeOk() (*int32, bool)`

GetCacheMinFlushAgeOk returns a tuple with the CacheMinFlushAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMinFlushAge

`func (o *ApiPoolGet200ResponseInner) SetCacheMinFlushAge(v int32)`

SetCacheMinFlushAge sets CacheMinFlushAge field to given value.

### HasCacheMinFlushAge

`func (o *ApiPoolGet200ResponseInner) HasCacheMinFlushAge() bool`

HasCacheMinFlushAge returns a boolean if a field has been set.

### GetCacheMode

`func (o *ApiPoolGet200ResponseInner) GetCacheMode() string`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *ApiPoolGet200ResponseInner) GetCacheModeOk() (*string, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *ApiPoolGet200ResponseInner) SetCacheMode(v string)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *ApiPoolGet200ResponseInner) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.

### GetCacheTargetDirtyHighRatioMicro

`func (o *ApiPoolGet200ResponseInner) GetCacheTargetDirtyHighRatioMicro() int32`

GetCacheTargetDirtyHighRatioMicro returns the CacheTargetDirtyHighRatioMicro field if non-nil, zero value otherwise.

### GetCacheTargetDirtyHighRatioMicroOk

`func (o *ApiPoolGet200ResponseInner) GetCacheTargetDirtyHighRatioMicroOk() (*int32, bool)`

GetCacheTargetDirtyHighRatioMicroOk returns a tuple with the CacheTargetDirtyHighRatioMicro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTargetDirtyHighRatioMicro

`func (o *ApiPoolGet200ResponseInner) SetCacheTargetDirtyHighRatioMicro(v int32)`

SetCacheTargetDirtyHighRatioMicro sets CacheTargetDirtyHighRatioMicro field to given value.

### HasCacheTargetDirtyHighRatioMicro

`func (o *ApiPoolGet200ResponseInner) HasCacheTargetDirtyHighRatioMicro() bool`

HasCacheTargetDirtyHighRatioMicro returns a boolean if a field has been set.

### GetCacheTargetDirtyRatioMicro

`func (o *ApiPoolGet200ResponseInner) GetCacheTargetDirtyRatioMicro() int32`

GetCacheTargetDirtyRatioMicro returns the CacheTargetDirtyRatioMicro field if non-nil, zero value otherwise.

### GetCacheTargetDirtyRatioMicroOk

`func (o *ApiPoolGet200ResponseInner) GetCacheTargetDirtyRatioMicroOk() (*int32, bool)`

GetCacheTargetDirtyRatioMicroOk returns a tuple with the CacheTargetDirtyRatioMicro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTargetDirtyRatioMicro

`func (o *ApiPoolGet200ResponseInner) SetCacheTargetDirtyRatioMicro(v int32)`

SetCacheTargetDirtyRatioMicro sets CacheTargetDirtyRatioMicro field to given value.

### HasCacheTargetDirtyRatioMicro

`func (o *ApiPoolGet200ResponseInner) HasCacheTargetDirtyRatioMicro() bool`

HasCacheTargetDirtyRatioMicro returns a boolean if a field has been set.

### GetCacheTargetFullRatioMicro

`func (o *ApiPoolGet200ResponseInner) GetCacheTargetFullRatioMicro() int32`

GetCacheTargetFullRatioMicro returns the CacheTargetFullRatioMicro field if non-nil, zero value otherwise.

### GetCacheTargetFullRatioMicroOk

`func (o *ApiPoolGet200ResponseInner) GetCacheTargetFullRatioMicroOk() (*int32, bool)`

GetCacheTargetFullRatioMicroOk returns a tuple with the CacheTargetFullRatioMicro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTargetFullRatioMicro

`func (o *ApiPoolGet200ResponseInner) SetCacheTargetFullRatioMicro(v int32)`

SetCacheTargetFullRatioMicro sets CacheTargetFullRatioMicro field to given value.

### HasCacheTargetFullRatioMicro

`func (o *ApiPoolGet200ResponseInner) HasCacheTargetFullRatioMicro() bool`

HasCacheTargetFullRatioMicro returns a boolean if a field has been set.

### GetCreateTime

`func (o *ApiPoolGet200ResponseInner) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ApiPoolGet200ResponseInner) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ApiPoolGet200ResponseInner) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ApiPoolGet200ResponseInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCrushRule

`func (o *ApiPoolGet200ResponseInner) GetCrushRule() string`

GetCrushRule returns the CrushRule field if non-nil, zero value otherwise.

### GetCrushRuleOk

`func (o *ApiPoolGet200ResponseInner) GetCrushRuleOk() (*string, bool)`

GetCrushRuleOk returns a tuple with the CrushRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrushRule

`func (o *ApiPoolGet200ResponseInner) SetCrushRule(v string)`

SetCrushRule sets CrushRule field to given value.

### HasCrushRule

`func (o *ApiPoolGet200ResponseInner) HasCrushRule() bool`

HasCrushRule returns a boolean if a field has been set.

### GetErasureCodeProfile

`func (o *ApiPoolGet200ResponseInner) GetErasureCodeProfile() string`

GetErasureCodeProfile returns the ErasureCodeProfile field if non-nil, zero value otherwise.

### GetErasureCodeProfileOk

`func (o *ApiPoolGet200ResponseInner) GetErasureCodeProfileOk() (*string, bool)`

GetErasureCodeProfileOk returns a tuple with the ErasureCodeProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErasureCodeProfile

`func (o *ApiPoolGet200ResponseInner) SetErasureCodeProfile(v string)`

SetErasureCodeProfile sets ErasureCodeProfile field to given value.

### HasErasureCodeProfile

`func (o *ApiPoolGet200ResponseInner) HasErasureCodeProfile() bool`

HasErasureCodeProfile returns a boolean if a field has been set.

### GetExpectedNumObjects

`func (o *ApiPoolGet200ResponseInner) GetExpectedNumObjects() int32`

GetExpectedNumObjects returns the ExpectedNumObjects field if non-nil, zero value otherwise.

### GetExpectedNumObjectsOk

`func (o *ApiPoolGet200ResponseInner) GetExpectedNumObjectsOk() (*int32, bool)`

GetExpectedNumObjectsOk returns a tuple with the ExpectedNumObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNumObjects

`func (o *ApiPoolGet200ResponseInner) SetExpectedNumObjects(v int32)`

SetExpectedNumObjects sets ExpectedNumObjects field to given value.

### HasExpectedNumObjects

`func (o *ApiPoolGet200ResponseInner) HasExpectedNumObjects() bool`

HasExpectedNumObjects returns a boolean if a field has been set.

### GetFastRead

`func (o *ApiPoolGet200ResponseInner) GetFastRead() bool`

GetFastRead returns the FastRead field if non-nil, zero value otherwise.

### GetFastReadOk

`func (o *ApiPoolGet200ResponseInner) GetFastReadOk() (*bool, bool)`

GetFastReadOk returns a tuple with the FastRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastRead

`func (o *ApiPoolGet200ResponseInner) SetFastRead(v bool)`

SetFastRead sets FastRead field to given value.

### HasFastRead

`func (o *ApiPoolGet200ResponseInner) HasFastRead() bool`

HasFastRead returns a boolean if a field has been set.

### GetFlags

`func (o *ApiPoolGet200ResponseInner) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ApiPoolGet200ResponseInner) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ApiPoolGet200ResponseInner) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ApiPoolGet200ResponseInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetFlagsNames

`func (o *ApiPoolGet200ResponseInner) GetFlagsNames() string`

GetFlagsNames returns the FlagsNames field if non-nil, zero value otherwise.

### GetFlagsNamesOk

`func (o *ApiPoolGet200ResponseInner) GetFlagsNamesOk() (*string, bool)`

GetFlagsNamesOk returns a tuple with the FlagsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsNames

`func (o *ApiPoolGet200ResponseInner) SetFlagsNames(v string)`

SetFlagsNames sets FlagsNames field to given value.

### HasFlagsNames

`func (o *ApiPoolGet200ResponseInner) HasFlagsNames() bool`

HasFlagsNames returns a boolean if a field has been set.

### GetGradeTable

`func (o *ApiPoolGet200ResponseInner) GetGradeTable() []string`

GetGradeTable returns the GradeTable field if non-nil, zero value otherwise.

### GetGradeTableOk

`func (o *ApiPoolGet200ResponseInner) GetGradeTableOk() (*[]string, bool)`

GetGradeTableOk returns a tuple with the GradeTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeTable

`func (o *ApiPoolGet200ResponseInner) SetGradeTable(v []string)`

SetGradeTable sets GradeTable field to given value.

### HasGradeTable

`func (o *ApiPoolGet200ResponseInner) HasGradeTable() bool`

HasGradeTable returns a boolean if a field has been set.

### GetHitSetCount

`func (o *ApiPoolGet200ResponseInner) GetHitSetCount() int32`

GetHitSetCount returns the HitSetCount field if non-nil, zero value otherwise.

### GetHitSetCountOk

`func (o *ApiPoolGet200ResponseInner) GetHitSetCountOk() (*int32, bool)`

GetHitSetCountOk returns a tuple with the HitSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSetCount

`func (o *ApiPoolGet200ResponseInner) SetHitSetCount(v int32)`

SetHitSetCount sets HitSetCount field to given value.

### HasHitSetCount

`func (o *ApiPoolGet200ResponseInner) HasHitSetCount() bool`

HasHitSetCount returns a boolean if a field has been set.

### GetHitSetGradeDecayRate

`func (o *ApiPoolGet200ResponseInner) GetHitSetGradeDecayRate() int32`

GetHitSetGradeDecayRate returns the HitSetGradeDecayRate field if non-nil, zero value otherwise.

### GetHitSetGradeDecayRateOk

`func (o *ApiPoolGet200ResponseInner) GetHitSetGradeDecayRateOk() (*int32, bool)`

GetHitSetGradeDecayRateOk returns a tuple with the HitSetGradeDecayRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSetGradeDecayRate

`func (o *ApiPoolGet200ResponseInner) SetHitSetGradeDecayRate(v int32)`

SetHitSetGradeDecayRate sets HitSetGradeDecayRate field to given value.

### HasHitSetGradeDecayRate

`func (o *ApiPoolGet200ResponseInner) HasHitSetGradeDecayRate() bool`

HasHitSetGradeDecayRate returns a boolean if a field has been set.

### GetHitSetParams

`func (o *ApiPoolGet200ResponseInner) GetHitSetParams() ApiPoolGet200ResponseInnerHitSetParams`

GetHitSetParams returns the HitSetParams field if non-nil, zero value otherwise.

### GetHitSetParamsOk

`func (o *ApiPoolGet200ResponseInner) GetHitSetParamsOk() (*ApiPoolGet200ResponseInnerHitSetParams, bool)`

GetHitSetParamsOk returns a tuple with the HitSetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSetParams

`func (o *ApiPoolGet200ResponseInner) SetHitSetParams(v ApiPoolGet200ResponseInnerHitSetParams)`

SetHitSetParams sets HitSetParams field to given value.

### HasHitSetParams

`func (o *ApiPoolGet200ResponseInner) HasHitSetParams() bool`

HasHitSetParams returns a boolean if a field has been set.

### GetHitSetPeriod

`func (o *ApiPoolGet200ResponseInner) GetHitSetPeriod() int32`

GetHitSetPeriod returns the HitSetPeriod field if non-nil, zero value otherwise.

### GetHitSetPeriodOk

`func (o *ApiPoolGet200ResponseInner) GetHitSetPeriodOk() (*int32, bool)`

GetHitSetPeriodOk returns a tuple with the HitSetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSetPeriod

`func (o *ApiPoolGet200ResponseInner) SetHitSetPeriod(v int32)`

SetHitSetPeriod sets HitSetPeriod field to given value.

### HasHitSetPeriod

`func (o *ApiPoolGet200ResponseInner) HasHitSetPeriod() bool`

HasHitSetPeriod returns a boolean if a field has been set.

### GetHitSetSearchLastN

`func (o *ApiPoolGet200ResponseInner) GetHitSetSearchLastN() int32`

GetHitSetSearchLastN returns the HitSetSearchLastN field if non-nil, zero value otherwise.

### GetHitSetSearchLastNOk

`func (o *ApiPoolGet200ResponseInner) GetHitSetSearchLastNOk() (*int32, bool)`

GetHitSetSearchLastNOk returns a tuple with the HitSetSearchLastN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitSetSearchLastN

`func (o *ApiPoolGet200ResponseInner) SetHitSetSearchLastN(v int32)`

SetHitSetSearchLastN sets HitSetSearchLastN field to given value.

### HasHitSetSearchLastN

`func (o *ApiPoolGet200ResponseInner) HasHitSetSearchLastN() bool`

HasHitSetSearchLastN returns a boolean if a field has been set.

### GetLastChange

`func (o *ApiPoolGet200ResponseInner) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *ApiPoolGet200ResponseInner) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *ApiPoolGet200ResponseInner) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *ApiPoolGet200ResponseInner) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetLastForceOpResend

`func (o *ApiPoolGet200ResponseInner) GetLastForceOpResend() string`

GetLastForceOpResend returns the LastForceOpResend field if non-nil, zero value otherwise.

### GetLastForceOpResendOk

`func (o *ApiPoolGet200ResponseInner) GetLastForceOpResendOk() (*string, bool)`

GetLastForceOpResendOk returns a tuple with the LastForceOpResend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastForceOpResend

`func (o *ApiPoolGet200ResponseInner) SetLastForceOpResend(v string)`

SetLastForceOpResend sets LastForceOpResend field to given value.

### HasLastForceOpResend

`func (o *ApiPoolGet200ResponseInner) HasLastForceOpResend() bool`

HasLastForceOpResend returns a boolean if a field has been set.

### GetLastForceOpResendPreluminous

`func (o *ApiPoolGet200ResponseInner) GetLastForceOpResendPreluminous() string`

GetLastForceOpResendPreluminous returns the LastForceOpResendPreluminous field if non-nil, zero value otherwise.

### GetLastForceOpResendPreluminousOk

`func (o *ApiPoolGet200ResponseInner) GetLastForceOpResendPreluminousOk() (*string, bool)`

GetLastForceOpResendPreluminousOk returns a tuple with the LastForceOpResendPreluminous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastForceOpResendPreluminous

`func (o *ApiPoolGet200ResponseInner) SetLastForceOpResendPreluminous(v string)`

SetLastForceOpResendPreluminous sets LastForceOpResendPreluminous field to given value.

### HasLastForceOpResendPreluminous

`func (o *ApiPoolGet200ResponseInner) HasLastForceOpResendPreluminous() bool`

HasLastForceOpResendPreluminous returns a boolean if a field has been set.

### GetLastForceOpResendPrenautilus

`func (o *ApiPoolGet200ResponseInner) GetLastForceOpResendPrenautilus() string`

GetLastForceOpResendPrenautilus returns the LastForceOpResendPrenautilus field if non-nil, zero value otherwise.

### GetLastForceOpResendPrenautilusOk

`func (o *ApiPoolGet200ResponseInner) GetLastForceOpResendPrenautilusOk() (*string, bool)`

GetLastForceOpResendPrenautilusOk returns a tuple with the LastForceOpResendPrenautilus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastForceOpResendPrenautilus

`func (o *ApiPoolGet200ResponseInner) SetLastForceOpResendPrenautilus(v string)`

SetLastForceOpResendPrenautilus sets LastForceOpResendPrenautilus field to given value.

### HasLastForceOpResendPrenautilus

`func (o *ApiPoolGet200ResponseInner) HasLastForceOpResendPrenautilus() bool`

HasLastForceOpResendPrenautilus returns a boolean if a field has been set.

### GetLastPgMergeMeta

`func (o *ApiPoolGet200ResponseInner) GetLastPgMergeMeta() ApiPoolGet200ResponseInnerLastPgMergeMeta`

GetLastPgMergeMeta returns the LastPgMergeMeta field if non-nil, zero value otherwise.

### GetLastPgMergeMetaOk

`func (o *ApiPoolGet200ResponseInner) GetLastPgMergeMetaOk() (*ApiPoolGet200ResponseInnerLastPgMergeMeta, bool)`

GetLastPgMergeMetaOk returns a tuple with the LastPgMergeMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPgMergeMeta

`func (o *ApiPoolGet200ResponseInner) SetLastPgMergeMeta(v ApiPoolGet200ResponseInnerLastPgMergeMeta)`

SetLastPgMergeMeta sets LastPgMergeMeta field to given value.

### HasLastPgMergeMeta

`func (o *ApiPoolGet200ResponseInner) HasLastPgMergeMeta() bool`

HasLastPgMergeMeta returns a boolean if a field has been set.

### GetMinReadRecencyForPromote

`func (o *ApiPoolGet200ResponseInner) GetMinReadRecencyForPromote() int32`

GetMinReadRecencyForPromote returns the MinReadRecencyForPromote field if non-nil, zero value otherwise.

### GetMinReadRecencyForPromoteOk

`func (o *ApiPoolGet200ResponseInner) GetMinReadRecencyForPromoteOk() (*int32, bool)`

GetMinReadRecencyForPromoteOk returns a tuple with the MinReadRecencyForPromote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadRecencyForPromote

`func (o *ApiPoolGet200ResponseInner) SetMinReadRecencyForPromote(v int32)`

SetMinReadRecencyForPromote sets MinReadRecencyForPromote field to given value.

### HasMinReadRecencyForPromote

`func (o *ApiPoolGet200ResponseInner) HasMinReadRecencyForPromote() bool`

HasMinReadRecencyForPromote returns a boolean if a field has been set.

### GetMinSize

`func (o *ApiPoolGet200ResponseInner) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *ApiPoolGet200ResponseInner) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *ApiPoolGet200ResponseInner) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *ApiPoolGet200ResponseInner) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### GetMinWriteRecencyForPromote

`func (o *ApiPoolGet200ResponseInner) GetMinWriteRecencyForPromote() int32`

GetMinWriteRecencyForPromote returns the MinWriteRecencyForPromote field if non-nil, zero value otherwise.

### GetMinWriteRecencyForPromoteOk

`func (o *ApiPoolGet200ResponseInner) GetMinWriteRecencyForPromoteOk() (*int32, bool)`

GetMinWriteRecencyForPromoteOk returns a tuple with the MinWriteRecencyForPromote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWriteRecencyForPromote

`func (o *ApiPoolGet200ResponseInner) SetMinWriteRecencyForPromote(v int32)`

SetMinWriteRecencyForPromote sets MinWriteRecencyForPromote field to given value.

### HasMinWriteRecencyForPromote

`func (o *ApiPoolGet200ResponseInner) HasMinWriteRecencyForPromote() bool`

HasMinWriteRecencyForPromote returns a boolean if a field has been set.

### GetObjectHash

`func (o *ApiPoolGet200ResponseInner) GetObjectHash() int32`

GetObjectHash returns the ObjectHash field if non-nil, zero value otherwise.

### GetObjectHashOk

`func (o *ApiPoolGet200ResponseInner) GetObjectHashOk() (*int32, bool)`

GetObjectHashOk returns a tuple with the ObjectHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHash

`func (o *ApiPoolGet200ResponseInner) SetObjectHash(v int32)`

SetObjectHash sets ObjectHash field to given value.

### HasObjectHash

`func (o *ApiPoolGet200ResponseInner) HasObjectHash() bool`

HasObjectHash returns a boolean if a field has been set.

### GetOptions

`func (o *ApiPoolGet200ResponseInner) GetOptions() ApiPoolGet200ResponseInnerOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApiPoolGet200ResponseInner) GetOptionsOk() (*ApiPoolGet200ResponseInnerOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApiPoolGet200ResponseInner) SetOptions(v ApiPoolGet200ResponseInnerOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApiPoolGet200ResponseInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPgAutoscaleMode

`func (o *ApiPoolGet200ResponseInner) GetPgAutoscaleMode() string`

GetPgAutoscaleMode returns the PgAutoscaleMode field if non-nil, zero value otherwise.

### GetPgAutoscaleModeOk

`func (o *ApiPoolGet200ResponseInner) GetPgAutoscaleModeOk() (*string, bool)`

GetPgAutoscaleModeOk returns a tuple with the PgAutoscaleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgAutoscaleMode

`func (o *ApiPoolGet200ResponseInner) SetPgAutoscaleMode(v string)`

SetPgAutoscaleMode sets PgAutoscaleMode field to given value.

### HasPgAutoscaleMode

`func (o *ApiPoolGet200ResponseInner) HasPgAutoscaleMode() bool`

HasPgAutoscaleMode returns a boolean if a field has been set.

### GetPgNum

`func (o *ApiPoolGet200ResponseInner) GetPgNum() int32`

GetPgNum returns the PgNum field if non-nil, zero value otherwise.

### GetPgNumOk

`func (o *ApiPoolGet200ResponseInner) GetPgNumOk() (*int32, bool)`

GetPgNumOk returns a tuple with the PgNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgNum

`func (o *ApiPoolGet200ResponseInner) SetPgNum(v int32)`

SetPgNum sets PgNum field to given value.

### HasPgNum

`func (o *ApiPoolGet200ResponseInner) HasPgNum() bool`

HasPgNum returns a boolean if a field has been set.

### GetPgNumPending

`func (o *ApiPoolGet200ResponseInner) GetPgNumPending() int32`

GetPgNumPending returns the PgNumPending field if non-nil, zero value otherwise.

### GetPgNumPendingOk

`func (o *ApiPoolGet200ResponseInner) GetPgNumPendingOk() (*int32, bool)`

GetPgNumPendingOk returns a tuple with the PgNumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgNumPending

`func (o *ApiPoolGet200ResponseInner) SetPgNumPending(v int32)`

SetPgNumPending sets PgNumPending field to given value.

### HasPgNumPending

`func (o *ApiPoolGet200ResponseInner) HasPgNumPending() bool`

HasPgNumPending returns a boolean if a field has been set.

### GetPgNumTarget

`func (o *ApiPoolGet200ResponseInner) GetPgNumTarget() int32`

GetPgNumTarget returns the PgNumTarget field if non-nil, zero value otherwise.

### GetPgNumTargetOk

`func (o *ApiPoolGet200ResponseInner) GetPgNumTargetOk() (*int32, bool)`

GetPgNumTargetOk returns a tuple with the PgNumTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgNumTarget

`func (o *ApiPoolGet200ResponseInner) SetPgNumTarget(v int32)`

SetPgNumTarget sets PgNumTarget field to given value.

### HasPgNumTarget

`func (o *ApiPoolGet200ResponseInner) HasPgNumTarget() bool`

HasPgNumTarget returns a boolean if a field has been set.

### GetPgPlacementNum

`func (o *ApiPoolGet200ResponseInner) GetPgPlacementNum() int32`

GetPgPlacementNum returns the PgPlacementNum field if non-nil, zero value otherwise.

### GetPgPlacementNumOk

`func (o *ApiPoolGet200ResponseInner) GetPgPlacementNumOk() (*int32, bool)`

GetPgPlacementNumOk returns a tuple with the PgPlacementNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgPlacementNum

`func (o *ApiPoolGet200ResponseInner) SetPgPlacementNum(v int32)`

SetPgPlacementNum sets PgPlacementNum field to given value.

### HasPgPlacementNum

`func (o *ApiPoolGet200ResponseInner) HasPgPlacementNum() bool`

HasPgPlacementNum returns a boolean if a field has been set.

### GetPgPlacementNumTarget

`func (o *ApiPoolGet200ResponseInner) GetPgPlacementNumTarget() int32`

GetPgPlacementNumTarget returns the PgPlacementNumTarget field if non-nil, zero value otherwise.

### GetPgPlacementNumTargetOk

`func (o *ApiPoolGet200ResponseInner) GetPgPlacementNumTargetOk() (*int32, bool)`

GetPgPlacementNumTargetOk returns a tuple with the PgPlacementNumTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgPlacementNumTarget

`func (o *ApiPoolGet200ResponseInner) SetPgPlacementNumTarget(v int32)`

SetPgPlacementNumTarget sets PgPlacementNumTarget field to given value.

### HasPgPlacementNumTarget

`func (o *ApiPoolGet200ResponseInner) HasPgPlacementNumTarget() bool`

HasPgPlacementNumTarget returns a boolean if a field has been set.

### GetPool

`func (o *ApiPoolGet200ResponseInner) GetPool() int32`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ApiPoolGet200ResponseInner) GetPoolOk() (*int32, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ApiPoolGet200ResponseInner) SetPool(v int32)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ApiPoolGet200ResponseInner) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolName

`func (o *ApiPoolGet200ResponseInner) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *ApiPoolGet200ResponseInner) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *ApiPoolGet200ResponseInner) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *ApiPoolGet200ResponseInner) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolSnaps

`func (o *ApiPoolGet200ResponseInner) GetPoolSnaps() []string`

GetPoolSnaps returns the PoolSnaps field if non-nil, zero value otherwise.

### GetPoolSnapsOk

`func (o *ApiPoolGet200ResponseInner) GetPoolSnapsOk() (*[]string, bool)`

GetPoolSnapsOk returns a tuple with the PoolSnaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSnaps

`func (o *ApiPoolGet200ResponseInner) SetPoolSnaps(v []string)`

SetPoolSnaps sets PoolSnaps field to given value.

### HasPoolSnaps

`func (o *ApiPoolGet200ResponseInner) HasPoolSnaps() bool`

HasPoolSnaps returns a boolean if a field has been set.

### GetQuotaMaxBytes

`func (o *ApiPoolGet200ResponseInner) GetQuotaMaxBytes() int32`

GetQuotaMaxBytes returns the QuotaMaxBytes field if non-nil, zero value otherwise.

### GetQuotaMaxBytesOk

`func (o *ApiPoolGet200ResponseInner) GetQuotaMaxBytesOk() (*int32, bool)`

GetQuotaMaxBytesOk returns a tuple with the QuotaMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMaxBytes

`func (o *ApiPoolGet200ResponseInner) SetQuotaMaxBytes(v int32)`

SetQuotaMaxBytes sets QuotaMaxBytes field to given value.

### HasQuotaMaxBytes

`func (o *ApiPoolGet200ResponseInner) HasQuotaMaxBytes() bool`

HasQuotaMaxBytes returns a boolean if a field has been set.

### GetQuotaMaxObjects

`func (o *ApiPoolGet200ResponseInner) GetQuotaMaxObjects() int32`

GetQuotaMaxObjects returns the QuotaMaxObjects field if non-nil, zero value otherwise.

### GetQuotaMaxObjectsOk

`func (o *ApiPoolGet200ResponseInner) GetQuotaMaxObjectsOk() (*int32, bool)`

GetQuotaMaxObjectsOk returns a tuple with the QuotaMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaMaxObjects

`func (o *ApiPoolGet200ResponseInner) SetQuotaMaxObjects(v int32)`

SetQuotaMaxObjects sets QuotaMaxObjects field to given value.

### HasQuotaMaxObjects

`func (o *ApiPoolGet200ResponseInner) HasQuotaMaxObjects() bool`

HasQuotaMaxObjects returns a boolean if a field has been set.

### GetReadTier

`func (o *ApiPoolGet200ResponseInner) GetReadTier() int32`

GetReadTier returns the ReadTier field if non-nil, zero value otherwise.

### GetReadTierOk

`func (o *ApiPoolGet200ResponseInner) GetReadTierOk() (*int32, bool)`

GetReadTierOk returns a tuple with the ReadTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTier

`func (o *ApiPoolGet200ResponseInner) SetReadTier(v int32)`

SetReadTier sets ReadTier field to given value.

### HasReadTier

`func (o *ApiPoolGet200ResponseInner) HasReadTier() bool`

HasReadTier returns a boolean if a field has been set.

### GetRemovedSnaps

`func (o *ApiPoolGet200ResponseInner) GetRemovedSnaps() []string`

GetRemovedSnaps returns the RemovedSnaps field if non-nil, zero value otherwise.

### GetRemovedSnapsOk

`func (o *ApiPoolGet200ResponseInner) GetRemovedSnapsOk() (*[]string, bool)`

GetRemovedSnapsOk returns a tuple with the RemovedSnaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedSnaps

`func (o *ApiPoolGet200ResponseInner) SetRemovedSnaps(v []string)`

SetRemovedSnaps sets RemovedSnaps field to given value.

### HasRemovedSnaps

`func (o *ApiPoolGet200ResponseInner) HasRemovedSnaps() bool`

HasRemovedSnaps returns a boolean if a field has been set.

### GetSize

`func (o *ApiPoolGet200ResponseInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApiPoolGet200ResponseInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApiPoolGet200ResponseInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ApiPoolGet200ResponseInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSnapEpoch

`func (o *ApiPoolGet200ResponseInner) GetSnapEpoch() int32`

GetSnapEpoch returns the SnapEpoch field if non-nil, zero value otherwise.

### GetSnapEpochOk

`func (o *ApiPoolGet200ResponseInner) GetSnapEpochOk() (*int32, bool)`

GetSnapEpochOk returns a tuple with the SnapEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapEpoch

`func (o *ApiPoolGet200ResponseInner) SetSnapEpoch(v int32)`

SetSnapEpoch sets SnapEpoch field to given value.

### HasSnapEpoch

`func (o *ApiPoolGet200ResponseInner) HasSnapEpoch() bool`

HasSnapEpoch returns a boolean if a field has been set.

### GetSnapMode

`func (o *ApiPoolGet200ResponseInner) GetSnapMode() string`

GetSnapMode returns the SnapMode field if non-nil, zero value otherwise.

### GetSnapModeOk

`func (o *ApiPoolGet200ResponseInner) GetSnapModeOk() (*string, bool)`

GetSnapModeOk returns a tuple with the SnapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapMode

`func (o *ApiPoolGet200ResponseInner) SetSnapMode(v string)`

SetSnapMode sets SnapMode field to given value.

### HasSnapMode

`func (o *ApiPoolGet200ResponseInner) HasSnapMode() bool`

HasSnapMode returns a boolean if a field has been set.

### GetSnapSeq

`func (o *ApiPoolGet200ResponseInner) GetSnapSeq() int32`

GetSnapSeq returns the SnapSeq field if non-nil, zero value otherwise.

### GetSnapSeqOk

`func (o *ApiPoolGet200ResponseInner) GetSnapSeqOk() (*int32, bool)`

GetSnapSeqOk returns a tuple with the SnapSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapSeq

`func (o *ApiPoolGet200ResponseInner) SetSnapSeq(v int32)`

SetSnapSeq sets SnapSeq field to given value.

### HasSnapSeq

`func (o *ApiPoolGet200ResponseInner) HasSnapSeq() bool`

HasSnapSeq returns a boolean if a field has been set.

### GetStripeWidth

`func (o *ApiPoolGet200ResponseInner) GetStripeWidth() int32`

GetStripeWidth returns the StripeWidth field if non-nil, zero value otherwise.

### GetStripeWidthOk

`func (o *ApiPoolGet200ResponseInner) GetStripeWidthOk() (*int32, bool)`

GetStripeWidthOk returns a tuple with the StripeWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeWidth

`func (o *ApiPoolGet200ResponseInner) SetStripeWidth(v int32)`

SetStripeWidth sets StripeWidth field to given value.

### HasStripeWidth

`func (o *ApiPoolGet200ResponseInner) HasStripeWidth() bool`

HasStripeWidth returns a boolean if a field has been set.

### GetTargetMaxBytes

`func (o *ApiPoolGet200ResponseInner) GetTargetMaxBytes() int32`

GetTargetMaxBytes returns the TargetMaxBytes field if non-nil, zero value otherwise.

### GetTargetMaxBytesOk

`func (o *ApiPoolGet200ResponseInner) GetTargetMaxBytesOk() (*int32, bool)`

GetTargetMaxBytesOk returns a tuple with the TargetMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMaxBytes

`func (o *ApiPoolGet200ResponseInner) SetTargetMaxBytes(v int32)`

SetTargetMaxBytes sets TargetMaxBytes field to given value.

### HasTargetMaxBytes

`func (o *ApiPoolGet200ResponseInner) HasTargetMaxBytes() bool`

HasTargetMaxBytes returns a boolean if a field has been set.

### GetTargetMaxObjects

`func (o *ApiPoolGet200ResponseInner) GetTargetMaxObjects() int32`

GetTargetMaxObjects returns the TargetMaxObjects field if non-nil, zero value otherwise.

### GetTargetMaxObjectsOk

`func (o *ApiPoolGet200ResponseInner) GetTargetMaxObjectsOk() (*int32, bool)`

GetTargetMaxObjectsOk returns a tuple with the TargetMaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMaxObjects

`func (o *ApiPoolGet200ResponseInner) SetTargetMaxObjects(v int32)`

SetTargetMaxObjects sets TargetMaxObjects field to given value.

### HasTargetMaxObjects

`func (o *ApiPoolGet200ResponseInner) HasTargetMaxObjects() bool`

HasTargetMaxObjects returns a boolean if a field has been set.

### GetTierOf

`func (o *ApiPoolGet200ResponseInner) GetTierOf() int32`

GetTierOf returns the TierOf field if non-nil, zero value otherwise.

### GetTierOfOk

`func (o *ApiPoolGet200ResponseInner) GetTierOfOk() (*int32, bool)`

GetTierOfOk returns a tuple with the TierOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierOf

`func (o *ApiPoolGet200ResponseInner) SetTierOf(v int32)`

SetTierOf sets TierOf field to given value.

### HasTierOf

`func (o *ApiPoolGet200ResponseInner) HasTierOf() bool`

HasTierOf returns a boolean if a field has been set.

### GetTiers

`func (o *ApiPoolGet200ResponseInner) GetTiers() []string`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *ApiPoolGet200ResponseInner) GetTiersOk() (*[]string, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *ApiPoolGet200ResponseInner) SetTiers(v []string)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *ApiPoolGet200ResponseInner) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetType

`func (o *ApiPoolGet200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPoolGet200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPoolGet200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiPoolGet200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUseGmtHitset

`func (o *ApiPoolGet200ResponseInner) GetUseGmtHitset() bool`

GetUseGmtHitset returns the UseGmtHitset field if non-nil, zero value otherwise.

### GetUseGmtHitsetOk

`func (o *ApiPoolGet200ResponseInner) GetUseGmtHitsetOk() (*bool, bool)`

GetUseGmtHitsetOk returns a tuple with the UseGmtHitset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGmtHitset

`func (o *ApiPoolGet200ResponseInner) SetUseGmtHitset(v bool)`

SetUseGmtHitset sets UseGmtHitset field to given value.

### HasUseGmtHitset

`func (o *ApiPoolGet200ResponseInner) HasUseGmtHitset() bool`

HasUseGmtHitset returns a boolean if a field has been set.

### GetWriteTier

`func (o *ApiPoolGet200ResponseInner) GetWriteTier() int32`

GetWriteTier returns the WriteTier field if non-nil, zero value otherwise.

### GetWriteTierOk

`func (o *ApiPoolGet200ResponseInner) GetWriteTierOk() (*int32, bool)`

GetWriteTierOk returns a tuple with the WriteTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteTier

`func (o *ApiPoolGet200ResponseInner) SetWriteTier(v int32)`

SetWriteTier sets WriteTier field to given value.

### HasWriteTier

`func (o *ApiPoolGet200ResponseInner) HasWriteTier() bool`

HasWriteTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


