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

// checks if the ApiTelemetryReportGet200ResponseReportCrushBucketAlgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportCrushBucketAlgs{}

// ApiTelemetryReportGet200ResponseReportCrushBucketAlgs 
type ApiTelemetryReportGet200ResponseReportCrushBucketAlgs struct {
	// 
	Straw2 int32 `json:"straw2"`
}

type _ApiTelemetryReportGet200ResponseReportCrushBucketAlgs ApiTelemetryReportGet200ResponseReportCrushBucketAlgs

// NewApiTelemetryReportGet200ResponseReportCrushBucketAlgs instantiates a new ApiTelemetryReportGet200ResponseReportCrushBucketAlgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportCrushBucketAlgs(straw2 int32) *ApiTelemetryReportGet200ResponseReportCrushBucketAlgs {
	this := ApiTelemetryReportGet200ResponseReportCrushBucketAlgs{}
	this.Straw2 = straw2
	return &this
}

// NewApiTelemetryReportGet200ResponseReportCrushBucketAlgsWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportCrushBucketAlgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportCrushBucketAlgsWithDefaults() *ApiTelemetryReportGet200ResponseReportCrushBucketAlgs {
	this := ApiTelemetryReportGet200ResponseReportCrushBucketAlgs{}
	return &this
}

// GetStraw2 returns the Straw2 field value
func (o *ApiTelemetryReportGet200ResponseReportCrushBucketAlgs) GetStraw2() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Straw2
}

// GetStraw2Ok returns a tuple with the Straw2 field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushBucketAlgs) GetStraw2Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Straw2, true
}

// SetStraw2 sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushBucketAlgs) SetStraw2(v int32) {
	o.Straw2 = v
}

func (o ApiTelemetryReportGet200ResponseReportCrushBucketAlgs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportCrushBucketAlgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["straw2"] = o.Straw2
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportCrushBucketAlgs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"straw2",
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

	varApiTelemetryReportGet200ResponseReportCrushBucketAlgs := _ApiTelemetryReportGet200ResponseReportCrushBucketAlgs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportCrushBucketAlgs)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportCrushBucketAlgs(varApiTelemetryReportGet200ResponseReportCrushBucketAlgs)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportCrushBucketAlgs struct {
	value *ApiTelemetryReportGet200ResponseReportCrushBucketAlgs
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportCrushBucketAlgs) Get() *ApiTelemetryReportGet200ResponseReportCrushBucketAlgs {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrushBucketAlgs) Set(val *ApiTelemetryReportGet200ResponseReportCrushBucketAlgs) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportCrushBucketAlgs) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrushBucketAlgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportCrushBucketAlgs(val *ApiTelemetryReportGet200ResponseReportCrushBucketAlgs) *NullableApiTelemetryReportGet200ResponseReportCrushBucketAlgs {
	return &NullableApiTelemetryReportGet200ResponseReportCrushBucketAlgs{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportCrushBucketAlgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrushBucketAlgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


