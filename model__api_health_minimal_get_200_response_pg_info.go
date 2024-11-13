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

// checks if the ApiHealthMinimalGet200ResponsePgInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHealthMinimalGet200ResponsePgInfo{}

// ApiHealthMinimalGet200ResponsePgInfo 
type ApiHealthMinimalGet200ResponsePgInfo struct {
	ObjectStats ApiHealthMinimalGet200ResponsePgInfoObjectStats `json:"object_stats"`
	// 
	PgsPerOsd int32 `json:"pgs_per_osd"`
	// 
	Statuses string `json:"statuses"`
}

type _ApiHealthMinimalGet200ResponsePgInfo ApiHealthMinimalGet200ResponsePgInfo

// NewApiHealthMinimalGet200ResponsePgInfo instantiates a new ApiHealthMinimalGet200ResponsePgInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHealthMinimalGet200ResponsePgInfo(objectStats ApiHealthMinimalGet200ResponsePgInfoObjectStats, pgsPerOsd int32, statuses string) *ApiHealthMinimalGet200ResponsePgInfo {
	this := ApiHealthMinimalGet200ResponsePgInfo{}
	this.ObjectStats = objectStats
	this.PgsPerOsd = pgsPerOsd
	this.Statuses = statuses
	return &this
}

// NewApiHealthMinimalGet200ResponsePgInfoWithDefaults instantiates a new ApiHealthMinimalGet200ResponsePgInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHealthMinimalGet200ResponsePgInfoWithDefaults() *ApiHealthMinimalGet200ResponsePgInfo {
	this := ApiHealthMinimalGet200ResponsePgInfo{}
	return &this
}

// GetObjectStats returns the ObjectStats field value
func (o *ApiHealthMinimalGet200ResponsePgInfo) GetObjectStats() ApiHealthMinimalGet200ResponsePgInfoObjectStats {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponsePgInfoObjectStats
		return ret
	}

	return o.ObjectStats
}

// GetObjectStatsOk returns a tuple with the ObjectStats field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200ResponsePgInfo) GetObjectStatsOk() (*ApiHealthMinimalGet200ResponsePgInfoObjectStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectStats, true
}

// SetObjectStats sets field value
func (o *ApiHealthMinimalGet200ResponsePgInfo) SetObjectStats(v ApiHealthMinimalGet200ResponsePgInfoObjectStats) {
	o.ObjectStats = v
}

// GetPgsPerOsd returns the PgsPerOsd field value
func (o *ApiHealthMinimalGet200ResponsePgInfo) GetPgsPerOsd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PgsPerOsd
}

// GetPgsPerOsdOk returns a tuple with the PgsPerOsd field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200ResponsePgInfo) GetPgsPerOsdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PgsPerOsd, true
}

// SetPgsPerOsd sets field value
func (o *ApiHealthMinimalGet200ResponsePgInfo) SetPgsPerOsd(v int32) {
	o.PgsPerOsd = v
}

// GetStatuses returns the Statuses field value
func (o *ApiHealthMinimalGet200ResponsePgInfo) GetStatuses() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200ResponsePgInfo) GetStatusesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statuses, true
}

// SetStatuses sets field value
func (o *ApiHealthMinimalGet200ResponsePgInfo) SetStatuses(v string) {
	o.Statuses = v
}

func (o ApiHealthMinimalGet200ResponsePgInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHealthMinimalGet200ResponsePgInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_stats"] = o.ObjectStats
	toSerialize["pgs_per_osd"] = o.PgsPerOsd
	toSerialize["statuses"] = o.Statuses
	return toSerialize, nil
}

func (o *ApiHealthMinimalGet200ResponsePgInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_stats",
		"pgs_per_osd",
		"statuses",
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

	varApiHealthMinimalGet200ResponsePgInfo := _ApiHealthMinimalGet200ResponsePgInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiHealthMinimalGet200ResponsePgInfo)

	if err != nil {
		return err
	}

	*o = ApiHealthMinimalGet200ResponsePgInfo(varApiHealthMinimalGet200ResponsePgInfo)

	return err
}

type NullableApiHealthMinimalGet200ResponsePgInfo struct {
	value *ApiHealthMinimalGet200ResponsePgInfo
	isSet bool
}

func (v NullableApiHealthMinimalGet200ResponsePgInfo) Get() *ApiHealthMinimalGet200ResponsePgInfo {
	return v.value
}

func (v *NullableApiHealthMinimalGet200ResponsePgInfo) Set(val *ApiHealthMinimalGet200ResponsePgInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHealthMinimalGet200ResponsePgInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHealthMinimalGet200ResponsePgInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHealthMinimalGet200ResponsePgInfo(val *ApiHealthMinimalGet200ResponsePgInfo) *NullableApiHealthMinimalGet200ResponsePgInfo {
	return &NullableApiHealthMinimalGet200ResponsePgInfo{value: val, isSet: true}
}

func (v NullableApiHealthMinimalGet200ResponsePgInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHealthMinimalGet200ResponsePgInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

