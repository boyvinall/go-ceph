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

// checks if the ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription{}

// ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription 
type ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription struct {
	// 
	Var1SMPWedJul1195301UTC2020 int32 `json:"#1 SMP Wed Jul 1 19:53:01 UTC 2020"`
}

type _ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription

// NewApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription instantiates a new ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription(var1SMPWedJul1195301UTC2020 int32) *ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription {
	this := ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription{}
	this.Var1SMPWedJul1195301UTC2020 = var1SMPWedJul1195301UTC2020
	return &this
}

// NewApiTelemetryReportGet200ResponseReportMetadataMonKernelDescriptionWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportMetadataMonKernelDescriptionWithDefaults() *ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription {
	this := ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription{}
	return &this
}

// GetVar1SMPWedJul1195301UTC2020 returns the Var1SMPWedJul1195301UTC2020 field value
func (o *ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) GetVar1SMPWedJul1195301UTC2020() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var1SMPWedJul1195301UTC2020
}

// GetVar1SMPWedJul1195301UTC2020Ok returns a tuple with the Var1SMPWedJul1195301UTC2020 field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) GetVar1SMPWedJul1195301UTC2020Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var1SMPWedJul1195301UTC2020, true
}

// SetVar1SMPWedJul1195301UTC2020 sets field value
func (o *ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) SetVar1SMPWedJul1195301UTC2020(v int32) {
	o.Var1SMPWedJul1195301UTC2020 = v
}

func (o ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["#1 SMP Wed Jul 1 19:53:01 UTC 2020"] = o.Var1SMPWedJul1195301UTC2020
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"#1 SMP Wed Jul 1 19:53:01 UTC 2020",
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

	varApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription := _ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription(varApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription struct {
	value *ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) Get() *ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) Set(val *ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription(val *ApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) *NullableApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription {
	return &NullableApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportMetadataMonKernelDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


