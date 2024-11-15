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

// checks if the ApiMonitorGet200ResponseMonStatusFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMonitorGet200ResponseMonStatusFeatures{}

// ApiMonitorGet200ResponseMonStatusFeatures 
type ApiMonitorGet200ResponseMonStatusFeatures struct {
	// 
	QuorumCon string `json:"quorum_con"`
	// 
	QuorumMon []string `json:"quorum_mon"`
	// 
	RequiredCon string `json:"required_con"`
	// 
	RequiredMon []int32 `json:"required_mon"`
}

type _ApiMonitorGet200ResponseMonStatusFeatures ApiMonitorGet200ResponseMonStatusFeatures

// NewApiMonitorGet200ResponseMonStatusFeatures instantiates a new ApiMonitorGet200ResponseMonStatusFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMonitorGet200ResponseMonStatusFeatures(quorumCon string, quorumMon []string, requiredCon string, requiredMon []int32) *ApiMonitorGet200ResponseMonStatusFeatures {
	this := ApiMonitorGet200ResponseMonStatusFeatures{}
	this.QuorumCon = quorumCon
	this.QuorumMon = quorumMon
	this.RequiredCon = requiredCon
	this.RequiredMon = requiredMon
	return &this
}

// NewApiMonitorGet200ResponseMonStatusFeaturesWithDefaults instantiates a new ApiMonitorGet200ResponseMonStatusFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMonitorGet200ResponseMonStatusFeaturesWithDefaults() *ApiMonitorGet200ResponseMonStatusFeatures {
	this := ApiMonitorGet200ResponseMonStatusFeatures{}
	return &this
}

// GetQuorumCon returns the QuorumCon field value
func (o *ApiMonitorGet200ResponseMonStatusFeatures) GetQuorumCon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuorumCon
}

// GetQuorumConOk returns a tuple with the QuorumCon field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseMonStatusFeatures) GetQuorumConOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuorumCon, true
}

// SetQuorumCon sets field value
func (o *ApiMonitorGet200ResponseMonStatusFeatures) SetQuorumCon(v string) {
	o.QuorumCon = v
}

// GetQuorumMon returns the QuorumMon field value
func (o *ApiMonitorGet200ResponseMonStatusFeatures) GetQuorumMon() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.QuorumMon
}

// GetQuorumMonOk returns a tuple with the QuorumMon field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseMonStatusFeatures) GetQuorumMonOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuorumMon, true
}

// SetQuorumMon sets field value
func (o *ApiMonitorGet200ResponseMonStatusFeatures) SetQuorumMon(v []string) {
	o.QuorumMon = v
}

// GetRequiredCon returns the RequiredCon field value
func (o *ApiMonitorGet200ResponseMonStatusFeatures) GetRequiredCon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequiredCon
}

// GetRequiredConOk returns a tuple with the RequiredCon field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseMonStatusFeatures) GetRequiredConOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiredCon, true
}

// SetRequiredCon sets field value
func (o *ApiMonitorGet200ResponseMonStatusFeatures) SetRequiredCon(v string) {
	o.RequiredCon = v
}

// GetRequiredMon returns the RequiredMon field value
func (o *ApiMonitorGet200ResponseMonStatusFeatures) GetRequiredMon() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.RequiredMon
}

// GetRequiredMonOk returns a tuple with the RequiredMon field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseMonStatusFeatures) GetRequiredMonOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredMon, true
}

// SetRequiredMon sets field value
func (o *ApiMonitorGet200ResponseMonStatusFeatures) SetRequiredMon(v []int32) {
	o.RequiredMon = v
}

func (o ApiMonitorGet200ResponseMonStatusFeatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMonitorGet200ResponseMonStatusFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quorum_con"] = o.QuorumCon
	toSerialize["quorum_mon"] = o.QuorumMon
	toSerialize["required_con"] = o.RequiredCon
	toSerialize["required_mon"] = o.RequiredMon
	return toSerialize, nil
}

func (o *ApiMonitorGet200ResponseMonStatusFeatures) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"quorum_con",
		"quorum_mon",
		"required_con",
		"required_mon",
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

	varApiMonitorGet200ResponseMonStatusFeatures := _ApiMonitorGet200ResponseMonStatusFeatures{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiMonitorGet200ResponseMonStatusFeatures)

	if err != nil {
		return err
	}

	*o = ApiMonitorGet200ResponseMonStatusFeatures(varApiMonitorGet200ResponseMonStatusFeatures)

	return err
}

type NullableApiMonitorGet200ResponseMonStatusFeatures struct {
	value *ApiMonitorGet200ResponseMonStatusFeatures
	isSet bool
}

func (v NullableApiMonitorGet200ResponseMonStatusFeatures) Get() *ApiMonitorGet200ResponseMonStatusFeatures {
	return v.value
}

func (v *NullableApiMonitorGet200ResponseMonStatusFeatures) Set(val *ApiMonitorGet200ResponseMonStatusFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMonitorGet200ResponseMonStatusFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMonitorGet200ResponseMonStatusFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMonitorGet200ResponseMonStatusFeatures(val *ApiMonitorGet200ResponseMonStatusFeatures) *NullableApiMonitorGet200ResponseMonStatusFeatures {
	return &NullableApiMonitorGet200ResponseMonStatusFeatures{value: val, isSet: true}
}

func (v NullableApiMonitorGet200ResponseMonStatusFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMonitorGet200ResponseMonStatusFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


