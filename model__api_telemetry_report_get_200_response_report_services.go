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

// checks if the ApiTelemetryReportGet200ResponseReportServices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportServices{}

// ApiTelemetryReportGet200ResponseReportServices 
type ApiTelemetryReportGet200ResponseReportServices struct {
	// 
	Rgw int32 `json:"rgw"`
}

type _ApiTelemetryReportGet200ResponseReportServices ApiTelemetryReportGet200ResponseReportServices

// NewApiTelemetryReportGet200ResponseReportServices instantiates a new ApiTelemetryReportGet200ResponseReportServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportServices(rgw int32) *ApiTelemetryReportGet200ResponseReportServices {
	this := ApiTelemetryReportGet200ResponseReportServices{}
	this.Rgw = rgw
	return &this
}

// NewApiTelemetryReportGet200ResponseReportServicesWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportServicesWithDefaults() *ApiTelemetryReportGet200ResponseReportServices {
	this := ApiTelemetryReportGet200ResponseReportServices{}
	return &this
}

// GetRgw returns the Rgw field value
func (o *ApiTelemetryReportGet200ResponseReportServices) GetRgw() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rgw
}

// GetRgwOk returns a tuple with the Rgw field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportServices) GetRgwOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rgw, true
}

// SetRgw sets field value
func (o *ApiTelemetryReportGet200ResponseReportServices) SetRgw(v int32) {
	o.Rgw = v
}

func (o ApiTelemetryReportGet200ResponseReportServices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportServices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rgw"] = o.Rgw
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportServices) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rgw",
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

	varApiTelemetryReportGet200ResponseReportServices := _ApiTelemetryReportGet200ResponseReportServices{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportServices)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportServices(varApiTelemetryReportGet200ResponseReportServices)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportServices struct {
	value *ApiTelemetryReportGet200ResponseReportServices
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportServices) Get() *ApiTelemetryReportGet200ResponseReportServices {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportServices) Set(val *ApiTelemetryReportGet200ResponseReportServices) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportServices) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportServices(val *ApiTelemetryReportGet200ResponseReportServices) *NullableApiTelemetryReportGet200ResponseReportServices {
	return &NullableApiTelemetryReportGet200ResponseReportServices{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


