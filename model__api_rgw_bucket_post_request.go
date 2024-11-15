/*
Ceph REST API

This is the official Ceph REST API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ceph

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiRgwBucketPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRgwBucketPostRequest{}

// ApiRgwBucketPostRequest struct for ApiRgwBucketPostRequest
type ApiRgwBucketPostRequest struct {
	Bucket string `json:"bucket"`
	BucketPolicy *string `json:"bucket_policy,omitempty"`
	CannedAcl *string `json:"canned_acl,omitempty"`
	DaemonName *string `json:"daemon_name,omitempty"`
	EncryptionState *string `json:"encryption_state,omitempty"`
	EncryptionType *string `json:"encryption_type,omitempty"`
	KeyId *string `json:"key_id,omitempty"`
	LockEnabled *string `json:"lock_enabled,omitempty"`
	LockMode *string `json:"lock_mode,omitempty"`
	LockRetentionPeriodDays *string `json:"lock_retention_period_days,omitempty"`
	LockRetentionPeriodYears *string `json:"lock_retention_period_years,omitempty"`
	PlacementTarget *string `json:"placement_target,omitempty"`
	Replication *string `json:"replication,omitempty"`
	Tags *string `json:"tags,omitempty"`
	Uid string `json:"uid"`
	Zonegroup *string `json:"zonegroup,omitempty"`
}

type _ApiRgwBucketPostRequest ApiRgwBucketPostRequest

// NewApiRgwBucketPostRequest instantiates a new ApiRgwBucketPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRgwBucketPostRequest(bucket string, uid string) *ApiRgwBucketPostRequest {
	this := ApiRgwBucketPostRequest{}
	this.Bucket = bucket
	var encryptionState string = "false"
	this.EncryptionState = &encryptionState
	var lockEnabled string = "false"
	this.LockEnabled = &lockEnabled
	var replication string = "false"
	this.Replication = &replication
	this.Uid = uid
	return &this
}

// NewApiRgwBucketPostRequestWithDefaults instantiates a new ApiRgwBucketPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRgwBucketPostRequestWithDefaults() *ApiRgwBucketPostRequest {
	this := ApiRgwBucketPostRequest{}
	var encryptionState string = "false"
	this.EncryptionState = &encryptionState
	var lockEnabled string = "false"
	this.LockEnabled = &lockEnabled
	var replication string = "false"
	this.Replication = &replication
	return &this
}

// GetBucket returns the Bucket field value
func (o *ApiRgwBucketPostRequest) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *ApiRgwBucketPostRequest) SetBucket(v string) {
	o.Bucket = v
}

// GetBucketPolicy returns the BucketPolicy field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetBucketPolicy() string {
	if o == nil || IsNil(o.BucketPolicy) {
		var ret string
		return ret
	}
	return *o.BucketPolicy
}

// GetBucketPolicyOk returns a tuple with the BucketPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetBucketPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.BucketPolicy) {
		return nil, false
	}
	return o.BucketPolicy, true
}

// HasBucketPolicy returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasBucketPolicy() bool {
	if o != nil && !IsNil(o.BucketPolicy) {
		return true
	}

	return false
}

// SetBucketPolicy gets a reference to the given string and assigns it to the BucketPolicy field.
func (o *ApiRgwBucketPostRequest) SetBucketPolicy(v string) {
	o.BucketPolicy = &v
}

// GetCannedAcl returns the CannedAcl field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetCannedAcl() string {
	if o == nil || IsNil(o.CannedAcl) {
		var ret string
		return ret
	}
	return *o.CannedAcl
}

// GetCannedAclOk returns a tuple with the CannedAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetCannedAclOk() (*string, bool) {
	if o == nil || IsNil(o.CannedAcl) {
		return nil, false
	}
	return o.CannedAcl, true
}

// HasCannedAcl returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasCannedAcl() bool {
	if o != nil && !IsNil(o.CannedAcl) {
		return true
	}

	return false
}

// SetCannedAcl gets a reference to the given string and assigns it to the CannedAcl field.
func (o *ApiRgwBucketPostRequest) SetCannedAcl(v string) {
	o.CannedAcl = &v
}

// GetDaemonName returns the DaemonName field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetDaemonName() string {
	if o == nil || IsNil(o.DaemonName) {
		var ret string
		return ret
	}
	return *o.DaemonName
}

// GetDaemonNameOk returns a tuple with the DaemonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetDaemonNameOk() (*string, bool) {
	if o == nil || IsNil(o.DaemonName) {
		return nil, false
	}
	return o.DaemonName, true
}

// HasDaemonName returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasDaemonName() bool {
	if o != nil && !IsNil(o.DaemonName) {
		return true
	}

	return false
}

// SetDaemonName gets a reference to the given string and assigns it to the DaemonName field.
func (o *ApiRgwBucketPostRequest) SetDaemonName(v string) {
	o.DaemonName = &v
}

// GetEncryptionState returns the EncryptionState field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetEncryptionState() string {
	if o == nil || IsNil(o.EncryptionState) {
		var ret string
		return ret
	}
	return *o.EncryptionState
}

// GetEncryptionStateOk returns a tuple with the EncryptionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetEncryptionStateOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionState) {
		return nil, false
	}
	return o.EncryptionState, true
}

// HasEncryptionState returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasEncryptionState() bool {
	if o != nil && !IsNil(o.EncryptionState) {
		return true
	}

	return false
}

// SetEncryptionState gets a reference to the given string and assigns it to the EncryptionState field.
func (o *ApiRgwBucketPostRequest) SetEncryptionState(v string) {
	o.EncryptionState = &v
}

// GetEncryptionType returns the EncryptionType field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetEncryptionType() string {
	if o == nil || IsNil(o.EncryptionType) {
		var ret string
		return ret
	}
	return *o.EncryptionType
}

// GetEncryptionTypeOk returns a tuple with the EncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetEncryptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionType) {
		return nil, false
	}
	return o.EncryptionType, true
}

// HasEncryptionType returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasEncryptionType() bool {
	if o != nil && !IsNil(o.EncryptionType) {
		return true
	}

	return false
}

// SetEncryptionType gets a reference to the given string and assigns it to the EncryptionType field.
func (o *ApiRgwBucketPostRequest) SetEncryptionType(v string) {
	o.EncryptionType = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *ApiRgwBucketPostRequest) SetKeyId(v string) {
	o.KeyId = &v
}

// GetLockEnabled returns the LockEnabled field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetLockEnabled() string {
	if o == nil || IsNil(o.LockEnabled) {
		var ret string
		return ret
	}
	return *o.LockEnabled
}

// GetLockEnabledOk returns a tuple with the LockEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetLockEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.LockEnabled) {
		return nil, false
	}
	return o.LockEnabled, true
}

// HasLockEnabled returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasLockEnabled() bool {
	if o != nil && !IsNil(o.LockEnabled) {
		return true
	}

	return false
}

// SetLockEnabled gets a reference to the given string and assigns it to the LockEnabled field.
func (o *ApiRgwBucketPostRequest) SetLockEnabled(v string) {
	o.LockEnabled = &v
}

// GetLockMode returns the LockMode field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetLockMode() string {
	if o == nil || IsNil(o.LockMode) {
		var ret string
		return ret
	}
	return *o.LockMode
}

// GetLockModeOk returns a tuple with the LockMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetLockModeOk() (*string, bool) {
	if o == nil || IsNil(o.LockMode) {
		return nil, false
	}
	return o.LockMode, true
}

// HasLockMode returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasLockMode() bool {
	if o != nil && !IsNil(o.LockMode) {
		return true
	}

	return false
}

// SetLockMode gets a reference to the given string and assigns it to the LockMode field.
func (o *ApiRgwBucketPostRequest) SetLockMode(v string) {
	o.LockMode = &v
}

// GetLockRetentionPeriodDays returns the LockRetentionPeriodDays field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetLockRetentionPeriodDays() string {
	if o == nil || IsNil(o.LockRetentionPeriodDays) {
		var ret string
		return ret
	}
	return *o.LockRetentionPeriodDays
}

// GetLockRetentionPeriodDaysOk returns a tuple with the LockRetentionPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetLockRetentionPeriodDaysOk() (*string, bool) {
	if o == nil || IsNil(o.LockRetentionPeriodDays) {
		return nil, false
	}
	return o.LockRetentionPeriodDays, true
}

// HasLockRetentionPeriodDays returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasLockRetentionPeriodDays() bool {
	if o != nil && !IsNil(o.LockRetentionPeriodDays) {
		return true
	}

	return false
}

// SetLockRetentionPeriodDays gets a reference to the given string and assigns it to the LockRetentionPeriodDays field.
func (o *ApiRgwBucketPostRequest) SetLockRetentionPeriodDays(v string) {
	o.LockRetentionPeriodDays = &v
}

// GetLockRetentionPeriodYears returns the LockRetentionPeriodYears field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetLockRetentionPeriodYears() string {
	if o == nil || IsNil(o.LockRetentionPeriodYears) {
		var ret string
		return ret
	}
	return *o.LockRetentionPeriodYears
}

// GetLockRetentionPeriodYearsOk returns a tuple with the LockRetentionPeriodYears field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetLockRetentionPeriodYearsOk() (*string, bool) {
	if o == nil || IsNil(o.LockRetentionPeriodYears) {
		return nil, false
	}
	return o.LockRetentionPeriodYears, true
}

// HasLockRetentionPeriodYears returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasLockRetentionPeriodYears() bool {
	if o != nil && !IsNil(o.LockRetentionPeriodYears) {
		return true
	}

	return false
}

// SetLockRetentionPeriodYears gets a reference to the given string and assigns it to the LockRetentionPeriodYears field.
func (o *ApiRgwBucketPostRequest) SetLockRetentionPeriodYears(v string) {
	o.LockRetentionPeriodYears = &v
}

// GetPlacementTarget returns the PlacementTarget field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetPlacementTarget() string {
	if o == nil || IsNil(o.PlacementTarget) {
		var ret string
		return ret
	}
	return *o.PlacementTarget
}

// GetPlacementTargetOk returns a tuple with the PlacementTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetPlacementTargetOk() (*string, bool) {
	if o == nil || IsNil(o.PlacementTarget) {
		return nil, false
	}
	return o.PlacementTarget, true
}

// HasPlacementTarget returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasPlacementTarget() bool {
	if o != nil && !IsNil(o.PlacementTarget) {
		return true
	}

	return false
}

// SetPlacementTarget gets a reference to the given string and assigns it to the PlacementTarget field.
func (o *ApiRgwBucketPostRequest) SetPlacementTarget(v string) {
	o.PlacementTarget = &v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetReplication() string {
	if o == nil || IsNil(o.Replication) {
		var ret string
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetReplicationOk() (*string, bool) {
	if o == nil || IsNil(o.Replication) {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasReplication() bool {
	if o != nil && !IsNil(o.Replication) {
		return true
	}

	return false
}

// SetReplication gets a reference to the given string and assigns it to the Replication field.
func (o *ApiRgwBucketPostRequest) SetReplication(v string) {
	o.Replication = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *ApiRgwBucketPostRequest) SetTags(v string) {
	o.Tags = &v
}

// GetUid returns the Uid field value
func (o *ApiRgwBucketPostRequest) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *ApiRgwBucketPostRequest) SetUid(v string) {
	o.Uid = v
}

// GetZonegroup returns the Zonegroup field value if set, zero value otherwise.
func (o *ApiRgwBucketPostRequest) GetZonegroup() string {
	if o == nil || IsNil(o.Zonegroup) {
		var ret string
		return ret
	}
	return *o.Zonegroup
}

// GetZonegroupOk returns a tuple with the Zonegroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwBucketPostRequest) GetZonegroupOk() (*string, bool) {
	if o == nil || IsNil(o.Zonegroup) {
		return nil, false
	}
	return o.Zonegroup, true
}

// HasZonegroup returns a boolean if a field has been set.
func (o *ApiRgwBucketPostRequest) HasZonegroup() bool {
	if o != nil && !IsNil(o.Zonegroup) {
		return true
	}

	return false
}

// SetZonegroup gets a reference to the given string and assigns it to the Zonegroup field.
func (o *ApiRgwBucketPostRequest) SetZonegroup(v string) {
	o.Zonegroup = &v
}

func (o ApiRgwBucketPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRgwBucketPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket"] = o.Bucket
	if !IsNil(o.BucketPolicy) {
		toSerialize["bucket_policy"] = o.BucketPolicy
	}
	if !IsNil(o.CannedAcl) {
		toSerialize["canned_acl"] = o.CannedAcl
	}
	if !IsNil(o.DaemonName) {
		toSerialize["daemon_name"] = o.DaemonName
	}
	if !IsNil(o.EncryptionState) {
		toSerialize["encryption_state"] = o.EncryptionState
	}
	if !IsNil(o.EncryptionType) {
		toSerialize["encryption_type"] = o.EncryptionType
	}
	if !IsNil(o.KeyId) {
		toSerialize["key_id"] = o.KeyId
	}
	if !IsNil(o.LockEnabled) {
		toSerialize["lock_enabled"] = o.LockEnabled
	}
	if !IsNil(o.LockMode) {
		toSerialize["lock_mode"] = o.LockMode
	}
	if !IsNil(o.LockRetentionPeriodDays) {
		toSerialize["lock_retention_period_days"] = o.LockRetentionPeriodDays
	}
	if !IsNil(o.LockRetentionPeriodYears) {
		toSerialize["lock_retention_period_years"] = o.LockRetentionPeriodYears
	}
	if !IsNil(o.PlacementTarget) {
		toSerialize["placement_target"] = o.PlacementTarget
	}
	if !IsNil(o.Replication) {
		toSerialize["replication"] = o.Replication
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["uid"] = o.Uid
	if !IsNil(o.Zonegroup) {
		toSerialize["zonegroup"] = o.Zonegroup
	}
	return toSerialize, nil
}

func (o *ApiRgwBucketPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bucket",
		"uid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiRgwBucketPostRequest := _ApiRgwBucketPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiRgwBucketPostRequest)

	if err != nil {
		return err
	}

	*o = ApiRgwBucketPostRequest(varApiRgwBucketPostRequest)

	return err
}

type NullableApiRgwBucketPostRequest struct {
	value *ApiRgwBucketPostRequest
	isSet bool
}

func (v NullableApiRgwBucketPostRequest) Get() *ApiRgwBucketPostRequest {
	return v.value
}

func (v *NullableApiRgwBucketPostRequest) Set(val *ApiRgwBucketPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRgwBucketPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRgwBucketPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRgwBucketPostRequest(val *ApiRgwBucketPostRequest) *NullableApiRgwBucketPostRequest {
	return &NullableApiRgwBucketPostRequest{value: val, isSet: true}
}

func (v NullableApiRgwBucketPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRgwBucketPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


