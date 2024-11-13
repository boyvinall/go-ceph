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

// checks if the ApiMonitorGet200ResponseInQuorumInnerStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMonitorGet200ResponseInQuorumInnerStats{}

// ApiMonitorGet200ResponseInQuorumInnerStats 
type ApiMonitorGet200ResponseInQuorumInnerStats struct {
	// 
	NumSessions []int32 `json:"num_sessions"`
}

type _ApiMonitorGet200ResponseInQuorumInnerStats ApiMonitorGet200ResponseInQuorumInnerStats

// NewApiMonitorGet200ResponseInQuorumInnerStats instantiates a new ApiMonitorGet200ResponseInQuorumInnerStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMonitorGet200ResponseInQuorumInnerStats(numSessions []int32) *ApiMonitorGet200ResponseInQuorumInnerStats {
	this := ApiMonitorGet200ResponseInQuorumInnerStats{}
	this.NumSessions = numSessions
	return &this
}

// NewApiMonitorGet200ResponseInQuorumInnerStatsWithDefaults instantiates a new ApiMonitorGet200ResponseInQuorumInnerStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMonitorGet200ResponseInQuorumInnerStatsWithDefaults() *ApiMonitorGet200ResponseInQuorumInnerStats {
	this := ApiMonitorGet200ResponseInQuorumInnerStats{}
	return &this
}

// GetNumSessions returns the NumSessions field value
func (o *ApiMonitorGet200ResponseInQuorumInnerStats) GetNumSessions() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.NumSessions
}

// GetNumSessionsOk returns a tuple with the NumSessions field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseInQuorumInnerStats) GetNumSessionsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumSessions, true
}

// SetNumSessions sets field value
func (o *ApiMonitorGet200ResponseInQuorumInnerStats) SetNumSessions(v []int32) {
	o.NumSessions = v
}

func (o ApiMonitorGet200ResponseInQuorumInnerStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMonitorGet200ResponseInQuorumInnerStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["num_sessions"] = o.NumSessions
	return toSerialize, nil
}

func (o *ApiMonitorGet200ResponseInQuorumInnerStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"num_sessions",
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

	varApiMonitorGet200ResponseInQuorumInnerStats := _ApiMonitorGet200ResponseInQuorumInnerStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiMonitorGet200ResponseInQuorumInnerStats)

	if err != nil {
		return err
	}

	*o = ApiMonitorGet200ResponseInQuorumInnerStats(varApiMonitorGet200ResponseInQuorumInnerStats)

	return err
}

type NullableApiMonitorGet200ResponseInQuorumInnerStats struct {
	value *ApiMonitorGet200ResponseInQuorumInnerStats
	isSet bool
}

func (v NullableApiMonitorGet200ResponseInQuorumInnerStats) Get() *ApiMonitorGet200ResponseInQuorumInnerStats {
	return v.value
}

func (v *NullableApiMonitorGet200ResponseInQuorumInnerStats) Set(val *ApiMonitorGet200ResponseInQuorumInnerStats) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMonitorGet200ResponseInQuorumInnerStats) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMonitorGet200ResponseInQuorumInnerStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMonitorGet200ResponseInQuorumInnerStats(val *ApiMonitorGet200ResponseInQuorumInnerStats) *NullableApiMonitorGet200ResponseInQuorumInnerStats {
	return &NullableApiMonitorGet200ResponseInQuorumInnerStats{value: val, isSet: true}
}

func (v NullableApiMonitorGet200ResponseInQuorumInnerStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMonitorGet200ResponseInQuorumInnerStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


