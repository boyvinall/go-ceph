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

// checks if the ApiMonitorGet200ResponseMonStatusFeatureMapClientInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMonitorGet200ResponseMonStatusFeatureMapClientInner{}

// ApiMonitorGet200ResponseMonStatusFeatureMapClientInner struct for ApiMonitorGet200ResponseMonStatusFeatureMapClientInner
type ApiMonitorGet200ResponseMonStatusFeatureMapClientInner struct {
	// 
	Features string `json:"features"`
	// 
	Num int32 `json:"num"`
	// 
	Release string `json:"release"`
}

type _ApiMonitorGet200ResponseMonStatusFeatureMapClientInner ApiMonitorGet200ResponseMonStatusFeatureMapClientInner

// NewApiMonitorGet200ResponseMonStatusFeatureMapClientInner instantiates a new ApiMonitorGet200ResponseMonStatusFeatureMapClientInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMonitorGet200ResponseMonStatusFeatureMapClientInner(features string, num int32, release string) *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner {
	this := ApiMonitorGet200ResponseMonStatusFeatureMapClientInner{}
	this.Features = features
	this.Num = num
	this.Release = release
	return &this
}

// NewApiMonitorGet200ResponseMonStatusFeatureMapClientInnerWithDefaults instantiates a new ApiMonitorGet200ResponseMonStatusFeatureMapClientInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMonitorGet200ResponseMonStatusFeatureMapClientInnerWithDefaults() *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner {
	this := ApiMonitorGet200ResponseMonStatusFeatureMapClientInner{}
	return &this
}

// GetFeatures returns the Features field value
func (o *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) GetFeatures() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) GetFeaturesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Features, true
}

// SetFeatures sets field value
func (o *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) SetFeatures(v string) {
	o.Features = v
}

// GetNum returns the Num field value
func (o *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) GetNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Num
}

// GetNumOk returns a tuple with the Num field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) GetNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Num, true
}

// SetNum sets field value
func (o *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) SetNum(v int32) {
	o.Num = v
}

// GetRelease returns the Release field value
func (o *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) GetRelease() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Release
}

// GetReleaseOk returns a tuple with the Release field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) GetReleaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Release, true
}

// SetRelease sets field value
func (o *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) SetRelease(v string) {
	o.Release = v
}

func (o ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["features"] = o.Features
	toSerialize["num"] = o.Num
	toSerialize["release"] = o.Release
	return toSerialize, nil
}

func (o *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"features",
		"num",
		"release",
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

	varApiMonitorGet200ResponseMonStatusFeatureMapClientInner := _ApiMonitorGet200ResponseMonStatusFeatureMapClientInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiMonitorGet200ResponseMonStatusFeatureMapClientInner)

	if err != nil {
		return err
	}

	*o = ApiMonitorGet200ResponseMonStatusFeatureMapClientInner(varApiMonitorGet200ResponseMonStatusFeatureMapClientInner)

	return err
}

type NullableApiMonitorGet200ResponseMonStatusFeatureMapClientInner struct {
	value *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner
	isSet bool
}

func (v NullableApiMonitorGet200ResponseMonStatusFeatureMapClientInner) Get() *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner {
	return v.value
}

func (v *NullableApiMonitorGet200ResponseMonStatusFeatureMapClientInner) Set(val *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMonitorGet200ResponseMonStatusFeatureMapClientInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMonitorGet200ResponseMonStatusFeatureMapClientInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMonitorGet200ResponseMonStatusFeatureMapClientInner(val *ApiMonitorGet200ResponseMonStatusFeatureMapClientInner) *NullableApiMonitorGet200ResponseMonStatusFeatureMapClientInner {
	return &NullableApiMonitorGet200ResponseMonStatusFeatureMapClientInner{value: val, isSet: true}
}

func (v NullableApiMonitorGet200ResponseMonStatusFeatureMapClientInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMonitorGet200ResponseMonStatusFeatureMapClientInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


