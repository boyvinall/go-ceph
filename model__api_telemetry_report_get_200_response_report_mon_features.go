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

// checks if the ApiTelemetryReportGet200ResponseReportMonFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportMonFeatures{}

// ApiTelemetryReportGet200ResponseReportMonFeatures 
type ApiTelemetryReportGet200ResponseReportMonFeatures struct {
	// 
	Optional []int32 `json:"optional"`
	// 
	Persistent []string `json:"persistent"`
}

type _ApiTelemetryReportGet200ResponseReportMonFeatures ApiTelemetryReportGet200ResponseReportMonFeatures

// NewApiTelemetryReportGet200ResponseReportMonFeatures instantiates a new ApiTelemetryReportGet200ResponseReportMonFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportMonFeatures(optional []int32, persistent []string) *ApiTelemetryReportGet200ResponseReportMonFeatures {
	this := ApiTelemetryReportGet200ResponseReportMonFeatures{}
	this.Optional = optional
	this.Persistent = persistent
	return &this
}

// NewApiTelemetryReportGet200ResponseReportMonFeaturesWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportMonFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportMonFeaturesWithDefaults() *ApiTelemetryReportGet200ResponseReportMonFeatures {
	this := ApiTelemetryReportGet200ResponseReportMonFeatures{}
	return &this
}

// GetOptional returns the Optional field value
func (o *ApiTelemetryReportGet200ResponseReportMonFeatures) GetOptional() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportMonFeatures) GetOptionalOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Optional, true
}

// SetOptional sets field value
func (o *ApiTelemetryReportGet200ResponseReportMonFeatures) SetOptional(v []int32) {
	o.Optional = v
}

// GetPersistent returns the Persistent field value
func (o *ApiTelemetryReportGet200ResponseReportMonFeatures) GetPersistent() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Persistent
}

// GetPersistentOk returns a tuple with the Persistent field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportMonFeatures) GetPersistentOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Persistent, true
}

// SetPersistent sets field value
func (o *ApiTelemetryReportGet200ResponseReportMonFeatures) SetPersistent(v []string) {
	o.Persistent = v
}

func (o ApiTelemetryReportGet200ResponseReportMonFeatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportMonFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optional"] = o.Optional
	toSerialize["persistent"] = o.Persistent
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportMonFeatures) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"optional",
		"persistent",
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

	varApiTelemetryReportGet200ResponseReportMonFeatures := _ApiTelemetryReportGet200ResponseReportMonFeatures{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportMonFeatures)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportMonFeatures(varApiTelemetryReportGet200ResponseReportMonFeatures)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportMonFeatures struct {
	value *ApiTelemetryReportGet200ResponseReportMonFeatures
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportMonFeatures) Get() *ApiTelemetryReportGet200ResponseReportMonFeatures {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportMonFeatures) Set(val *ApiTelemetryReportGet200ResponseReportMonFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportMonFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportMonFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportMonFeatures(val *ApiTelemetryReportGet200ResponseReportMonFeatures) *NullableApiTelemetryReportGet200ResponseReportMonFeatures {
	return &NullableApiTelemetryReportGet200ResponseReportMonFeatures{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportMonFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportMonFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


