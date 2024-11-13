# ApiRgwBucketBucketPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketId** | **string** |  | 
**BucketPolicy** | Pointer to **string** |  | [optional] 
**CannedAcl** | Pointer to **string** |  | [optional] 
**DaemonName** | Pointer to **string** |  | [optional] 
**EncryptionState** | Pointer to **string** |  | [optional] [default to "false"]
**EncryptionType** | Pointer to **string** |  | [optional] 
**KeyId** | Pointer to **string** |  | [optional] 
**Lifecycle** | Pointer to **string** |  | [optional] 
**LockMode** | Pointer to **string** |  | [optional] 
**LockRetentionPeriodDays** | Pointer to **string** |  | [optional] 
**LockRetentionPeriodYears** | Pointer to **string** |  | [optional] 
**MfaDelete** | Pointer to **string** |  | [optional] 
**MfaTokenPin** | Pointer to **string** |  | [optional] 
**MfaTokenSerial** | Pointer to **string** |  | [optional] 
**Replication** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**VersioningState** | Pointer to **string** |  | [optional] 

## Methods

### NewApiRgwBucketBucketPutRequest

`func NewApiRgwBucketBucketPutRequest(bucketId string, ) *ApiRgwBucketBucketPutRequest`

NewApiRgwBucketBucketPutRequest instantiates a new ApiRgwBucketBucketPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwBucketBucketPutRequestWithDefaults

`func NewApiRgwBucketBucketPutRequestWithDefaults() *ApiRgwBucketBucketPutRequest`

NewApiRgwBucketBucketPutRequestWithDefaults instantiates a new ApiRgwBucketBucketPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketId

`func (o *ApiRgwBucketBucketPutRequest) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *ApiRgwBucketBucketPutRequest) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *ApiRgwBucketBucketPutRequest) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.


### GetBucketPolicy

`func (o *ApiRgwBucketBucketPutRequest) GetBucketPolicy() string`

GetBucketPolicy returns the BucketPolicy field if non-nil, zero value otherwise.

### GetBucketPolicyOk

`func (o *ApiRgwBucketBucketPutRequest) GetBucketPolicyOk() (*string, bool)`

GetBucketPolicyOk returns a tuple with the BucketPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicy

`func (o *ApiRgwBucketBucketPutRequest) SetBucketPolicy(v string)`

SetBucketPolicy sets BucketPolicy field to given value.

### HasBucketPolicy

`func (o *ApiRgwBucketBucketPutRequest) HasBucketPolicy() bool`

HasBucketPolicy returns a boolean if a field has been set.

### GetCannedAcl

`func (o *ApiRgwBucketBucketPutRequest) GetCannedAcl() string`

GetCannedAcl returns the CannedAcl field if non-nil, zero value otherwise.

### GetCannedAclOk

`func (o *ApiRgwBucketBucketPutRequest) GetCannedAclOk() (*string, bool)`

GetCannedAclOk returns a tuple with the CannedAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCannedAcl

`func (o *ApiRgwBucketBucketPutRequest) SetCannedAcl(v string)`

SetCannedAcl sets CannedAcl field to given value.

### HasCannedAcl

`func (o *ApiRgwBucketBucketPutRequest) HasCannedAcl() bool`

HasCannedAcl returns a boolean if a field has been set.

### GetDaemonName

`func (o *ApiRgwBucketBucketPutRequest) GetDaemonName() string`

GetDaemonName returns the DaemonName field if non-nil, zero value otherwise.

### GetDaemonNameOk

`func (o *ApiRgwBucketBucketPutRequest) GetDaemonNameOk() (*string, bool)`

GetDaemonNameOk returns a tuple with the DaemonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonName

`func (o *ApiRgwBucketBucketPutRequest) SetDaemonName(v string)`

SetDaemonName sets DaemonName field to given value.

### HasDaemonName

`func (o *ApiRgwBucketBucketPutRequest) HasDaemonName() bool`

HasDaemonName returns a boolean if a field has been set.

### GetEncryptionState

`func (o *ApiRgwBucketBucketPutRequest) GetEncryptionState() string`

GetEncryptionState returns the EncryptionState field if non-nil, zero value otherwise.

### GetEncryptionStateOk

`func (o *ApiRgwBucketBucketPutRequest) GetEncryptionStateOk() (*string, bool)`

GetEncryptionStateOk returns a tuple with the EncryptionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionState

`func (o *ApiRgwBucketBucketPutRequest) SetEncryptionState(v string)`

SetEncryptionState sets EncryptionState field to given value.

### HasEncryptionState

`func (o *ApiRgwBucketBucketPutRequest) HasEncryptionState() bool`

HasEncryptionState returns a boolean if a field has been set.

### GetEncryptionType

`func (o *ApiRgwBucketBucketPutRequest) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *ApiRgwBucketBucketPutRequest) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *ApiRgwBucketBucketPutRequest) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *ApiRgwBucketBucketPutRequest) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.

### GetKeyId

`func (o *ApiRgwBucketBucketPutRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ApiRgwBucketBucketPutRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ApiRgwBucketBucketPutRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ApiRgwBucketBucketPutRequest) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetLifecycle

`func (o *ApiRgwBucketBucketPutRequest) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *ApiRgwBucketBucketPutRequest) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *ApiRgwBucketBucketPutRequest) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *ApiRgwBucketBucketPutRequest) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLockMode

`func (o *ApiRgwBucketBucketPutRequest) GetLockMode() string`

GetLockMode returns the LockMode field if non-nil, zero value otherwise.

### GetLockModeOk

`func (o *ApiRgwBucketBucketPutRequest) GetLockModeOk() (*string, bool)`

GetLockModeOk returns a tuple with the LockMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockMode

`func (o *ApiRgwBucketBucketPutRequest) SetLockMode(v string)`

SetLockMode sets LockMode field to given value.

### HasLockMode

`func (o *ApiRgwBucketBucketPutRequest) HasLockMode() bool`

HasLockMode returns a boolean if a field has been set.

### GetLockRetentionPeriodDays

`func (o *ApiRgwBucketBucketPutRequest) GetLockRetentionPeriodDays() string`

GetLockRetentionPeriodDays returns the LockRetentionPeriodDays field if non-nil, zero value otherwise.

### GetLockRetentionPeriodDaysOk

`func (o *ApiRgwBucketBucketPutRequest) GetLockRetentionPeriodDaysOk() (*string, bool)`

GetLockRetentionPeriodDaysOk returns a tuple with the LockRetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockRetentionPeriodDays

`func (o *ApiRgwBucketBucketPutRequest) SetLockRetentionPeriodDays(v string)`

SetLockRetentionPeriodDays sets LockRetentionPeriodDays field to given value.

### HasLockRetentionPeriodDays

`func (o *ApiRgwBucketBucketPutRequest) HasLockRetentionPeriodDays() bool`

HasLockRetentionPeriodDays returns a boolean if a field has been set.

### GetLockRetentionPeriodYears

`func (o *ApiRgwBucketBucketPutRequest) GetLockRetentionPeriodYears() string`

GetLockRetentionPeriodYears returns the LockRetentionPeriodYears field if non-nil, zero value otherwise.

### GetLockRetentionPeriodYearsOk

`func (o *ApiRgwBucketBucketPutRequest) GetLockRetentionPeriodYearsOk() (*string, bool)`

GetLockRetentionPeriodYearsOk returns a tuple with the LockRetentionPeriodYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockRetentionPeriodYears

`func (o *ApiRgwBucketBucketPutRequest) SetLockRetentionPeriodYears(v string)`

SetLockRetentionPeriodYears sets LockRetentionPeriodYears field to given value.

### HasLockRetentionPeriodYears

`func (o *ApiRgwBucketBucketPutRequest) HasLockRetentionPeriodYears() bool`

HasLockRetentionPeriodYears returns a boolean if a field has been set.

### GetMfaDelete

`func (o *ApiRgwBucketBucketPutRequest) GetMfaDelete() string`

GetMfaDelete returns the MfaDelete field if non-nil, zero value otherwise.

### GetMfaDeleteOk

`func (o *ApiRgwBucketBucketPutRequest) GetMfaDeleteOk() (*string, bool)`

GetMfaDeleteOk returns a tuple with the MfaDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaDelete

`func (o *ApiRgwBucketBucketPutRequest) SetMfaDelete(v string)`

SetMfaDelete sets MfaDelete field to given value.

### HasMfaDelete

`func (o *ApiRgwBucketBucketPutRequest) HasMfaDelete() bool`

HasMfaDelete returns a boolean if a field has been set.

### GetMfaTokenPin

`func (o *ApiRgwBucketBucketPutRequest) GetMfaTokenPin() string`

GetMfaTokenPin returns the MfaTokenPin field if non-nil, zero value otherwise.

### GetMfaTokenPinOk

`func (o *ApiRgwBucketBucketPutRequest) GetMfaTokenPinOk() (*string, bool)`

GetMfaTokenPinOk returns a tuple with the MfaTokenPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaTokenPin

`func (o *ApiRgwBucketBucketPutRequest) SetMfaTokenPin(v string)`

SetMfaTokenPin sets MfaTokenPin field to given value.

### HasMfaTokenPin

`func (o *ApiRgwBucketBucketPutRequest) HasMfaTokenPin() bool`

HasMfaTokenPin returns a boolean if a field has been set.

### GetMfaTokenSerial

`func (o *ApiRgwBucketBucketPutRequest) GetMfaTokenSerial() string`

GetMfaTokenSerial returns the MfaTokenSerial field if non-nil, zero value otherwise.

### GetMfaTokenSerialOk

`func (o *ApiRgwBucketBucketPutRequest) GetMfaTokenSerialOk() (*string, bool)`

GetMfaTokenSerialOk returns a tuple with the MfaTokenSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaTokenSerial

`func (o *ApiRgwBucketBucketPutRequest) SetMfaTokenSerial(v string)`

SetMfaTokenSerial sets MfaTokenSerial field to given value.

### HasMfaTokenSerial

`func (o *ApiRgwBucketBucketPutRequest) HasMfaTokenSerial() bool`

HasMfaTokenSerial returns a boolean if a field has been set.

### GetReplication

`func (o *ApiRgwBucketBucketPutRequest) GetReplication() string`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *ApiRgwBucketBucketPutRequest) GetReplicationOk() (*string, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *ApiRgwBucketBucketPutRequest) SetReplication(v string)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *ApiRgwBucketBucketPutRequest) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetTags

`func (o *ApiRgwBucketBucketPutRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApiRgwBucketBucketPutRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApiRgwBucketBucketPutRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApiRgwBucketBucketPutRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUid

`func (o *ApiRgwBucketBucketPutRequest) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApiRgwBucketBucketPutRequest) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApiRgwBucketBucketPutRequest) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ApiRgwBucketBucketPutRequest) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetVersioningState

`func (o *ApiRgwBucketBucketPutRequest) GetVersioningState() string`

GetVersioningState returns the VersioningState field if non-nil, zero value otherwise.

### GetVersioningStateOk

`func (o *ApiRgwBucketBucketPutRequest) GetVersioningStateOk() (*string, bool)`

GetVersioningStateOk returns a tuple with the VersioningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioningState

`func (o *ApiRgwBucketBucketPutRequest) SetVersioningState(v string)`

SetVersioningState sets VersioningState field to given value.

### HasVersioningState

`func (o *ApiRgwBucketBucketPutRequest) HasVersioningState() bool`

HasVersioningState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


