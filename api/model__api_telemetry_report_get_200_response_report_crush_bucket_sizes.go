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

// checks if the ApiTelemetryReportGet200ResponseReportCrushBucketSizes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportCrushBucketSizes{}

// ApiTelemetryReportGet200ResponseReportCrushBucketSizes 
type ApiTelemetryReportGet200ResponseReportCrushBucketSizes struct {
	// 
	Var1 int32 `json:"1"`
	// 
	Var3 int32 `json:"3"`
}

type _ApiTelemetryReportGet200ResponseReportCrushBucketSizes ApiTelemetryReportGet200ResponseReportCrushBucketSizes

// NewApiTelemetryReportGet200ResponseReportCrushBucketSizes instantiates a new ApiTelemetryReportGet200ResponseReportCrushBucketSizes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportCrushBucketSizes(var1 int32, var3 int32) *ApiTelemetryReportGet200ResponseReportCrushBucketSizes {
	this := ApiTelemetryReportGet200ResponseReportCrushBucketSizes{}
	this.Var1 = var1
	this.Var3 = var3
	return &this
}

// NewApiTelemetryReportGet200ResponseReportCrushBucketSizesWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportCrushBucketSizes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportCrushBucketSizesWithDefaults() *ApiTelemetryReportGet200ResponseReportCrushBucketSizes {
	this := ApiTelemetryReportGet200ResponseReportCrushBucketSizes{}
	return &this
}

// GetVar1 returns the Var1 field value
func (o *ApiTelemetryReportGet200ResponseReportCrushBucketSizes) GetVar1() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushBucketSizes) GetVar1Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var1, true
}

// SetVar1 sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushBucketSizes) SetVar1(v int32) {
	o.Var1 = v
}

// GetVar3 returns the Var3 field value
func (o *ApiTelemetryReportGet200ResponseReportCrushBucketSizes) GetVar3() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushBucketSizes) GetVar3Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var3, true
}

// SetVar3 sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushBucketSizes) SetVar3(v int32) {
	o.Var3 = v
}

func (o ApiTelemetryReportGet200ResponseReportCrushBucketSizes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportCrushBucketSizes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["1"] = o.Var1
	toSerialize["3"] = o.Var3
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportCrushBucketSizes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"1",
		"3",
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

	varApiTelemetryReportGet200ResponseReportCrushBucketSizes := _ApiTelemetryReportGet200ResponseReportCrushBucketSizes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportCrushBucketSizes)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportCrushBucketSizes(varApiTelemetryReportGet200ResponseReportCrushBucketSizes)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportCrushBucketSizes struct {
	value *ApiTelemetryReportGet200ResponseReportCrushBucketSizes
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportCrushBucketSizes) Get() *ApiTelemetryReportGet200ResponseReportCrushBucketSizes {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrushBucketSizes) Set(val *ApiTelemetryReportGet200ResponseReportCrushBucketSizes) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportCrushBucketSizes) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrushBucketSizes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportCrushBucketSizes(val *ApiTelemetryReportGet200ResponseReportCrushBucketSizes) *NullableApiTelemetryReportGet200ResponseReportCrushBucketSizes {
	return &NullableApiTelemetryReportGet200ResponseReportCrushBucketSizes{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportCrushBucketSizes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrushBucketSizes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

