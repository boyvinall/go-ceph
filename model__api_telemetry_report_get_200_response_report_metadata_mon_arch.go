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

// checks if the ApiTelemetryReportGet200ResponseReportMetadataMonArch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportMetadataMonArch{}

// ApiTelemetryReportGet200ResponseReportMetadataMonArch 
type ApiTelemetryReportGet200ResponseReportMetadataMonArch struct {
	// 
	X8664 int32 `json:"x86_64"`
}

type _ApiTelemetryReportGet200ResponseReportMetadataMonArch ApiTelemetryReportGet200ResponseReportMetadataMonArch

// NewApiTelemetryReportGet200ResponseReportMetadataMonArch instantiates a new ApiTelemetryReportGet200ResponseReportMetadataMonArch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportMetadataMonArch(x8664 int32) *ApiTelemetryReportGet200ResponseReportMetadataMonArch {
	this := ApiTelemetryReportGet200ResponseReportMetadataMonArch{}
	this.X8664 = x8664
	return &this
}

// NewApiTelemetryReportGet200ResponseReportMetadataMonArchWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportMetadataMonArch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportMetadataMonArchWithDefaults() *ApiTelemetryReportGet200ResponseReportMetadataMonArch {
	this := ApiTelemetryReportGet200ResponseReportMetadataMonArch{}
	return &this
}

// GetX8664 returns the X8664 field value
func (o *ApiTelemetryReportGet200ResponseReportMetadataMonArch) GetX8664() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.X8664
}

// GetX8664Ok returns a tuple with the X8664 field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportMetadataMonArch) GetX8664Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X8664, true
}

// SetX8664 sets field value
func (o *ApiTelemetryReportGet200ResponseReportMetadataMonArch) SetX8664(v int32) {
	o.X8664 = v
}

func (o ApiTelemetryReportGet200ResponseReportMetadataMonArch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportMetadataMonArch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x86_64"] = o.X8664
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportMetadataMonArch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"x86_64",
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

	varApiTelemetryReportGet200ResponseReportMetadataMonArch := _ApiTelemetryReportGet200ResponseReportMetadataMonArch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportMetadataMonArch)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportMetadataMonArch(varApiTelemetryReportGet200ResponseReportMetadataMonArch)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportMetadataMonArch struct {
	value *ApiTelemetryReportGet200ResponseReportMetadataMonArch
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportMetadataMonArch) Get() *ApiTelemetryReportGet200ResponseReportMetadataMonArch {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportMetadataMonArch) Set(val *ApiTelemetryReportGet200ResponseReportMetadataMonArch) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportMetadataMonArch) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportMetadataMonArch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportMetadataMonArch(val *ApiTelemetryReportGet200ResponseReportMetadataMonArch) *NullableApiTelemetryReportGet200ResponseReportMetadataMonArch {
	return &NullableApiTelemetryReportGet200ResponseReportMetadataMonArch{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportMetadataMonArch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportMetadataMonArch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


