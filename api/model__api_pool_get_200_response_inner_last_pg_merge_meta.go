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

// checks if the ApiPoolGet200ResponseInnerLastPgMergeMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPoolGet200ResponseInnerLastPgMergeMeta{}

// ApiPoolGet200ResponseInnerLastPgMergeMeta 
type ApiPoolGet200ResponseInnerLastPgMergeMeta struct {
	// 
	LastEpochClean int32 `json:"last_epoch_clean"`
	// 
	LastEpochStarted int32 `json:"last_epoch_started"`
	// 
	ReadyEpoch int32 `json:"ready_epoch"`
	// 
	SourcePgid string `json:"source_pgid"`
	// 
	SourceVersion string `json:"source_version"`
	// 
	TargetVersion string `json:"target_version"`
}

type _ApiPoolGet200ResponseInnerLastPgMergeMeta ApiPoolGet200ResponseInnerLastPgMergeMeta

// NewApiPoolGet200ResponseInnerLastPgMergeMeta instantiates a new ApiPoolGet200ResponseInnerLastPgMergeMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPoolGet200ResponseInnerLastPgMergeMeta(lastEpochClean int32, lastEpochStarted int32, readyEpoch int32, sourcePgid string, sourceVersion string, targetVersion string) *ApiPoolGet200ResponseInnerLastPgMergeMeta {
	this := ApiPoolGet200ResponseInnerLastPgMergeMeta{}
	this.LastEpochClean = lastEpochClean
	this.LastEpochStarted = lastEpochStarted
	this.ReadyEpoch = readyEpoch
	this.SourcePgid = sourcePgid
	this.SourceVersion = sourceVersion
	this.TargetVersion = targetVersion
	return &this
}

// NewApiPoolGet200ResponseInnerLastPgMergeMetaWithDefaults instantiates a new ApiPoolGet200ResponseInnerLastPgMergeMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPoolGet200ResponseInnerLastPgMergeMetaWithDefaults() *ApiPoolGet200ResponseInnerLastPgMergeMeta {
	this := ApiPoolGet200ResponseInnerLastPgMergeMeta{}
	return &this
}

// GetLastEpochClean returns the LastEpochClean field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetLastEpochClean() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LastEpochClean
}

// GetLastEpochCleanOk returns a tuple with the LastEpochClean field value
// and a boolean to check if the value has been set.
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetLastEpochCleanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEpochClean, true
}

// SetLastEpochClean sets field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) SetLastEpochClean(v int32) {
	o.LastEpochClean = v
}

// GetLastEpochStarted returns the LastEpochStarted field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetLastEpochStarted() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LastEpochStarted
}

// GetLastEpochStartedOk returns a tuple with the LastEpochStarted field value
// and a boolean to check if the value has been set.
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetLastEpochStartedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEpochStarted, true
}

// SetLastEpochStarted sets field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) SetLastEpochStarted(v int32) {
	o.LastEpochStarted = v
}

// GetReadyEpoch returns the ReadyEpoch field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetReadyEpoch() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReadyEpoch
}

// GetReadyEpochOk returns a tuple with the ReadyEpoch field value
// and a boolean to check if the value has been set.
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetReadyEpochOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadyEpoch, true
}

// SetReadyEpoch sets field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) SetReadyEpoch(v int32) {
	o.ReadyEpoch = v
}

// GetSourcePgid returns the SourcePgid field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetSourcePgid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourcePgid
}

// GetSourcePgidOk returns a tuple with the SourcePgid field value
// and a boolean to check if the value has been set.
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetSourcePgidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcePgid, true
}

// SetSourcePgid sets field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) SetSourcePgid(v string) {
	o.SourcePgid = v
}

// GetSourceVersion returns the SourceVersion field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetSourceVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceVersion
}

// GetSourceVersionOk returns a tuple with the SourceVersion field value
// and a boolean to check if the value has been set.
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetSourceVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceVersion, true
}

// SetSourceVersion sets field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) SetSourceVersion(v string) {
	o.SourceVersion = v
}

// GetTargetVersion returns the TargetVersion field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetTargetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value
// and a boolean to check if the value has been set.
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) GetTargetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetVersion, true
}

// SetTargetVersion sets field value
func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) SetTargetVersion(v string) {
	o.TargetVersion = v
}

func (o ApiPoolGet200ResponseInnerLastPgMergeMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPoolGet200ResponseInnerLastPgMergeMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["last_epoch_clean"] = o.LastEpochClean
	toSerialize["last_epoch_started"] = o.LastEpochStarted
	toSerialize["ready_epoch"] = o.ReadyEpoch
	toSerialize["source_pgid"] = o.SourcePgid
	toSerialize["source_version"] = o.SourceVersion
	toSerialize["target_version"] = o.TargetVersion
	return toSerialize, nil
}

func (o *ApiPoolGet200ResponseInnerLastPgMergeMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"last_epoch_clean",
		"last_epoch_started",
		"ready_epoch",
		"source_pgid",
		"source_version",
		"target_version",
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

	varApiPoolGet200ResponseInnerLastPgMergeMeta := _ApiPoolGet200ResponseInnerLastPgMergeMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiPoolGet200ResponseInnerLastPgMergeMeta)

	if err != nil {
		return err
	}

	*o = ApiPoolGet200ResponseInnerLastPgMergeMeta(varApiPoolGet200ResponseInnerLastPgMergeMeta)

	return err
}

type NullableApiPoolGet200ResponseInnerLastPgMergeMeta struct {
	value *ApiPoolGet200ResponseInnerLastPgMergeMeta
	isSet bool
}

func (v NullableApiPoolGet200ResponseInnerLastPgMergeMeta) Get() *ApiPoolGet200ResponseInnerLastPgMergeMeta {
	return v.value
}

func (v *NullableApiPoolGet200ResponseInnerLastPgMergeMeta) Set(val *ApiPoolGet200ResponseInnerLastPgMergeMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPoolGet200ResponseInnerLastPgMergeMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPoolGet200ResponseInnerLastPgMergeMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPoolGet200ResponseInnerLastPgMergeMeta(val *ApiPoolGet200ResponseInnerLastPgMergeMeta) *NullableApiPoolGet200ResponseInnerLastPgMergeMeta {
	return &NullableApiPoolGet200ResponseInnerLastPgMergeMeta{value: val, isSet: true}
}

func (v NullableApiPoolGet200ResponseInnerLastPgMergeMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPoolGet200ResponseInnerLastPgMergeMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


