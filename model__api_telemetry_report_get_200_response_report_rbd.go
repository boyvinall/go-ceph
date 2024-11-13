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

// checks if the ApiTelemetryReportGet200ResponseReportRbd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportRbd{}

// ApiTelemetryReportGet200ResponseReportRbd 
type ApiTelemetryReportGet200ResponseReportRbd struct {
	// 
	MirroringByPool []bool `json:"mirroring_by_pool"`
	// 
	NumImagesByPool []int32 `json:"num_images_by_pool"`
	// 
	NumPools int32 `json:"num_pools"`
}

type _ApiTelemetryReportGet200ResponseReportRbd ApiTelemetryReportGet200ResponseReportRbd

// NewApiTelemetryReportGet200ResponseReportRbd instantiates a new ApiTelemetryReportGet200ResponseReportRbd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportRbd(mirroringByPool []bool, numImagesByPool []int32, numPools int32) *ApiTelemetryReportGet200ResponseReportRbd {
	this := ApiTelemetryReportGet200ResponseReportRbd{}
	this.MirroringByPool = mirroringByPool
	this.NumImagesByPool = numImagesByPool
	this.NumPools = numPools
	return &this
}

// NewApiTelemetryReportGet200ResponseReportRbdWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportRbd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportRbdWithDefaults() *ApiTelemetryReportGet200ResponseReportRbd {
	this := ApiTelemetryReportGet200ResponseReportRbd{}
	return &this
}

// GetMirroringByPool returns the MirroringByPool field value
func (o *ApiTelemetryReportGet200ResponseReportRbd) GetMirroringByPool() []bool {
	if o == nil {
		var ret []bool
		return ret
	}

	return o.MirroringByPool
}

// GetMirroringByPoolOk returns a tuple with the MirroringByPool field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportRbd) GetMirroringByPoolOk() ([]bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.MirroringByPool, true
}

// SetMirroringByPool sets field value
func (o *ApiTelemetryReportGet200ResponseReportRbd) SetMirroringByPool(v []bool) {
	o.MirroringByPool = v
}

// GetNumImagesByPool returns the NumImagesByPool field value
func (o *ApiTelemetryReportGet200ResponseReportRbd) GetNumImagesByPool() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.NumImagesByPool
}

// GetNumImagesByPoolOk returns a tuple with the NumImagesByPool field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportRbd) GetNumImagesByPoolOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumImagesByPool, true
}

// SetNumImagesByPool sets field value
func (o *ApiTelemetryReportGet200ResponseReportRbd) SetNumImagesByPool(v []int32) {
	o.NumImagesByPool = v
}

// GetNumPools returns the NumPools field value
func (o *ApiTelemetryReportGet200ResponseReportRbd) GetNumPools() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPools
}

// GetNumPoolsOk returns a tuple with the NumPools field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportRbd) GetNumPoolsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPools, true
}

// SetNumPools sets field value
func (o *ApiTelemetryReportGet200ResponseReportRbd) SetNumPools(v int32) {
	o.NumPools = v
}

func (o ApiTelemetryReportGet200ResponseReportRbd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportRbd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mirroring_by_pool"] = o.MirroringByPool
	toSerialize["num_images_by_pool"] = o.NumImagesByPool
	toSerialize["num_pools"] = o.NumPools
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportRbd) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mirroring_by_pool",
		"num_images_by_pool",
		"num_pools",
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

	varApiTelemetryReportGet200ResponseReportRbd := _ApiTelemetryReportGet200ResponseReportRbd{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportRbd)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportRbd(varApiTelemetryReportGet200ResponseReportRbd)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportRbd struct {
	value *ApiTelemetryReportGet200ResponseReportRbd
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportRbd) Get() *ApiTelemetryReportGet200ResponseReportRbd {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportRbd) Set(val *ApiTelemetryReportGet200ResponseReportRbd) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportRbd) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportRbd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportRbd(val *ApiTelemetryReportGet200ResponseReportRbd) *NullableApiTelemetryReportGet200ResponseReportRbd {
	return &NullableApiTelemetryReportGet200ResponseReportRbd{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportRbd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportRbd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


