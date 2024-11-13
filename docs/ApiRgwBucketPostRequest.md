# ApiRgwBucketPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**BucketPolicy** | Pointer to **string** |  | [optional] 
**CannedAcl** | Pointer to **string** |  | [optional] 
**DaemonName** | Pointer to **string** |  | [optional] 
**EncryptionState** | Pointer to **string** |  | [optional] [default to "false"]
**EncryptionType** | Pointer to **string** |  | [optional] 
**KeyId** | Pointer to **string** |  | [optional] 
**LockEnabled** | Pointer to **string** |  | [optional] [default to "false"]
**LockMode** | Pointer to **string** |  | [optional] 
**LockRetentionPeriodDays** | Pointer to **string** |  | [optional] 
**LockRetentionPeriodYears** | Pointer to **string** |  | [optional] 
**PlacementTarget** | Pointer to **string** |  | [optional] 
**Replication** | Pointer to **string** |  | [optional] [default to "false"]
**Tags** | Pointer to **string** |  | [optional] 
**Uid** | **string** |  | 
**Zonegroup** | Pointer to **string** |  | [optional] 

## Methods

### NewApiRgwBucketPostRequest

`func NewApiRgwBucketPostRequest(bucket string, uid string, ) *ApiRgwBucketPostRequest`

NewApiRgwBucketPostRequest instantiates a new ApiRgwBucketPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwBucketPostRequestWithDefaults

`func NewApiRgwBucketPostRequestWithDefaults() *ApiRgwBucketPostRequest`

NewApiRgwBucketPostRequestWithDefaults instantiates a new ApiRgwBucketPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *ApiRgwBucketPostRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ApiRgwBucketPostRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ApiRgwBucketPostRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetBucketPolicy

`func (o *ApiRgwBucketPostRequest) GetBucketPolicy() string`

GetBucketPolicy returns the BucketPolicy field if non-nil, zero value otherwise.

### GetBucketPolicyOk

`func (o *ApiRgwBucketPostRequest) GetBucketPolicyOk() (*string, bool)`

GetBucketPolicyOk returns a tuple with the BucketPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicy

`func (o *ApiRgwBucketPostRequest) SetBucketPolicy(v string)`

SetBucketPolicy sets BucketPolicy field to given value.

### HasBucketPolicy

`func (o *ApiRgwBucketPostRequest) HasBucketPolicy() bool`

HasBucketPolicy returns a boolean if a field has been set.

### GetCannedAcl

`func (o *ApiRgwBucketPostRequest) GetCannedAcl() string`

GetCannedAcl returns the CannedAcl field if non-nil, zero value otherwise.

### GetCannedAclOk

`func (o *ApiRgwBucketPostRequest) GetCannedAclOk() (*string, bool)`

GetCannedAclOk returns a tuple with the CannedAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCannedAcl

`func (o *ApiRgwBucketPostRequest) SetCannedAcl(v string)`

SetCannedAcl sets CannedAcl field to given value.

### HasCannedAcl

`func (o *ApiRgwBucketPostRequest) HasCannedAcl() bool`

HasCannedAcl returns a boolean if a field has been set.

### GetDaemonName

`func (o *ApiRgwBucketPostRequest) GetDaemonName() string`

GetDaemonName returns the DaemonName field if non-nil, zero value otherwise.

### GetDaemonNameOk

`func (o *ApiRgwBucketPostRequest) GetDaemonNameOk() (*string, bool)`

GetDaemonNameOk returns a tuple with the DaemonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonName

`func (o *ApiRgwBucketPostRequest) SetDaemonName(v string)`

SetDaemonName sets DaemonName field to given value.

### HasDaemonName

`func (o *ApiRgwBucketPostRequest) HasDaemonName() bool`

HasDaemonName returns a boolean if a field has been set.

### GetEncryptionState

`func (o *ApiRgwBucketPostRequest) GetEncryptionState() string`

GetEncryptionState returns the EncryptionState field if non-nil, zero value otherwise.

### GetEncryptionStateOk

`func (o *ApiRgwBucketPostRequest) GetEncryptionStateOk() (*string, bool)`

GetEncryptionStateOk returns a tuple with the EncryptionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionState

`func (o *ApiRgwBucketPostRequest) SetEncryptionState(v string)`

SetEncryptionState sets EncryptionState field to given value.

### HasEncryptionState

`func (o *ApiRgwBucketPostRequest) HasEncryptionState() bool`

HasEncryptionState returns a boolean if a field has been set.

### GetEncryptionType

`func (o *ApiRgwBucketPostRequest) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *ApiRgwBucketPostRequest) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *ApiRgwBucketPostRequest) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *ApiRgwBucketPostRequest) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.

### GetKeyId

`func (o *ApiRgwBucketPostRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ApiRgwBucketPostRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ApiRgwBucketPostRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ApiRgwBucketPostRequest) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetLockEnabled

`func (o *ApiRgwBucketPostRequest) GetLockEnabled() string`

GetLockEnabled returns the LockEnabled field if non-nil, zero value otherwise.

### GetLockEnabledOk

`func (o *ApiRgwBucketPostRequest) GetLockEnabledOk() (*string, bool)`

GetLockEnabledOk returns a tuple with the LockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockEnabled

`func (o *ApiRgwBucketPostRequest) SetLockEnabled(v string)`

SetLockEnabled sets LockEnabled field to given value.

### HasLockEnabled

`func (o *ApiRgwBucketPostRequest) HasLockEnabled() bool`

HasLockEnabled returns a boolean if a field has been set.

### GetLockMode

`func (o *ApiRgwBucketPostRequest) GetLockMode() string`

GetLockMode returns the LockMode field if non-nil, zero value otherwise.

### GetLockModeOk

`func (o *ApiRgwBucketPostRequest) GetLockModeOk() (*string, bool)`

GetLockModeOk returns a tuple with the LockMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockMode

`func (o *ApiRgwBucketPostRequest) SetLockMode(v string)`

SetLockMode sets LockMode field to given value.

### HasLockMode

`func (o *ApiRgwBucketPostRequest) HasLockMode() bool`

HasLockMode returns a boolean if a field has been set.

### GetLockRetentionPeriodDays

`func (o *ApiRgwBucketPostRequest) GetLockRetentionPeriodDays() string`

GetLockRetentionPeriodDays returns the LockRetentionPeriodDays field if non-nil, zero value otherwise.

### GetLockRetentionPeriodDaysOk

`func (o *ApiRgwBucketPostRequest) GetLockRetentionPeriodDaysOk() (*string, bool)`

GetLockRetentionPeriodDaysOk returns a tuple with the LockRetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockRetentionPeriodDays

`func (o *ApiRgwBucketPostRequest) SetLockRetentionPeriodDays(v string)`

SetLockRetentionPeriodDays sets LockRetentionPeriodDays field to given value.

### HasLockRetentionPeriodDays

`func (o *ApiRgwBucketPostRequest) HasLockRetentionPeriodDays() bool`

HasLockRetentionPeriodDays returns a boolean if a field has been set.

### GetLockRetentionPeriodYears

`func (o *ApiRgwBucketPostRequest) GetLockRetentionPeriodYears() string`

GetLockRetentionPeriodYears returns the LockRetentionPeriodYears field if non-nil, zero value otherwise.

### GetLockRetentionPeriodYearsOk

`func (o *ApiRgwBucketPostRequest) GetLockRetentionPeriodYearsOk() (*string, bool)`

GetLockRetentionPeriodYearsOk returns a tuple with the LockRetentionPeriodYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockRetentionPeriodYears

`func (o *ApiRgwBucketPostRequest) SetLockRetentionPeriodYears(v string)`

SetLockRetentionPeriodYears sets LockRetentionPeriodYears field to given value.

### HasLockRetentionPeriodYears

`func (o *ApiRgwBucketPostRequest) HasLockRetentionPeriodYears() bool`

HasLockRetentionPeriodYears returns a boolean if a field has been set.

### GetPlacementTarget

`func (o *ApiRgwBucketPostRequest) GetPlacementTarget() string`

GetPlacementTarget returns the PlacementTarget field if non-nil, zero value otherwise.

### GetPlacementTargetOk

`func (o *ApiRgwBucketPostRequest) GetPlacementTargetOk() (*string, bool)`

GetPlacementTargetOk returns a tuple with the PlacementTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementTarget

`func (o *ApiRgwBucketPostRequest) SetPlacementTarget(v string)`

SetPlacementTarget sets PlacementTarget field to given value.

### HasPlacementTarget

`func (o *ApiRgwBucketPostRequest) HasPlacementTarget() bool`

HasPlacementTarget returns a boolean if a field has been set.

### GetReplication

`func (o *ApiRgwBucketPostRequest) GetReplication() string`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *ApiRgwBucketPostRequest) GetReplicationOk() (*string, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *ApiRgwBucketPostRequest) SetReplication(v string)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *ApiRgwBucketPostRequest) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetTags

`func (o *ApiRgwBucketPostRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApiRgwBucketPostRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApiRgwBucketPostRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApiRgwBucketPostRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUid

`func (o *ApiRgwBucketPostRequest) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApiRgwBucketPostRequest) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApiRgwBucketPostRequest) SetUid(v string)`

SetUid sets Uid field to given value.


### GetZonegroup

`func (o *ApiRgwBucketPostRequest) GetZonegroup() string`

GetZonegroup returns the Zonegroup field if non-nil, zero value otherwise.

### GetZonegroupOk

`func (o *ApiRgwBucketPostRequest) GetZonegroupOk() (*string, bool)`

GetZonegroupOk returns a tuple with the Zonegroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonegroup

`func (o *ApiRgwBucketPostRequest) SetZonegroup(v string)`

SetZonegroup sets Zonegroup field to given value.

### HasZonegroup

`func (o *ApiRgwBucketPostRequest) HasZonegroup() bool`

HasZonegroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


