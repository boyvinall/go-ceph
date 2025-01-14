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

// checks if the ApiPerfCountersGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPerfCountersGet200Response{}

// ApiPerfCountersGet200Response struct for ApiPerfCountersGet200Response
type ApiPerfCountersGet200Response struct {
	MonA ApiPerfCountersGet200ResponseMonA `json:"mon.a"`
}

type _ApiPerfCountersGet200Response ApiPerfCountersGet200Response

// NewApiPerfCountersGet200Response instantiates a new ApiPerfCountersGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPerfCountersGet200Response(monA ApiPerfCountersGet200ResponseMonA) *ApiPerfCountersGet200Response {
	this := ApiPerfCountersGet200Response{}
	this.MonA = monA
	return &this
}

// NewApiPerfCountersGet200ResponseWithDefaults instantiates a new ApiPerfCountersGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPerfCountersGet200ResponseWithDefaults() *ApiPerfCountersGet200Response {
	this := ApiPerfCountersGet200Response{}
	return &this
}

// GetMonA returns the MonA field value
func (o *ApiPerfCountersGet200Response) GetMonA() ApiPerfCountersGet200ResponseMonA {
	if o == nil {
		var ret ApiPerfCountersGet200ResponseMonA
		return ret
	}

	return o.MonA
}

// GetMonAOk returns a tuple with the MonA field value
// and a boolean to check if the value has been set.
func (o *ApiPerfCountersGet200Response) GetMonAOk() (*ApiPerfCountersGet200ResponseMonA, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonA, true
}

// SetMonA sets field value
func (o *ApiPerfCountersGet200Response) SetMonA(v ApiPerfCountersGet200ResponseMonA) {
	o.MonA = v
}

func (o ApiPerfCountersGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPerfCountersGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mon.a"] = o.MonA
	return toSerialize, nil
}

func (o *ApiPerfCountersGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mon.a",
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

	varApiPerfCountersGet200Response := _ApiPerfCountersGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiPerfCountersGet200Response)

	if err != nil {
		return err
	}

	*o = ApiPerfCountersGet200Response(varApiPerfCountersGet200Response)

	return err
}

type NullableApiPerfCountersGet200Response struct {
	value *ApiPerfCountersGet200Response
	isSet bool
}

func (v NullableApiPerfCountersGet200Response) Get() *ApiPerfCountersGet200Response {
	return v.value
}

func (v *NullableApiPerfCountersGet200Response) Set(val *ApiPerfCountersGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPerfCountersGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPerfCountersGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPerfCountersGet200Response(val *ApiPerfCountersGet200Response) *NullableApiPerfCountersGet200Response {
	return &NullableApiPerfCountersGet200Response{value: val, isSet: true}
}

func (v NullableApiPerfCountersGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPerfCountersGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


