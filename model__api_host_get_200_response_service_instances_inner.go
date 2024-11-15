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

// checks if the ApiHostGet200ResponseServiceInstancesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHostGet200ResponseServiceInstancesInner{}

// ApiHostGet200ResponseServiceInstancesInner struct for ApiHostGet200ResponseServiceInstancesInner
type ApiHostGet200ResponseServiceInstancesInner struct {
	// Number of instances of the service
	Count int32 `json:"count"`
	// type of service
	Type string `json:"type"`
}

type _ApiHostGet200ResponseServiceInstancesInner ApiHostGet200ResponseServiceInstancesInner

// NewApiHostGet200ResponseServiceInstancesInner instantiates a new ApiHostGet200ResponseServiceInstancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHostGet200ResponseServiceInstancesInner(count int32, type_ string) *ApiHostGet200ResponseServiceInstancesInner {
	this := ApiHostGet200ResponseServiceInstancesInner{}
	this.Count = count
	this.Type = type_
	return &this
}

// NewApiHostGet200ResponseServiceInstancesInnerWithDefaults instantiates a new ApiHostGet200ResponseServiceInstancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHostGet200ResponseServiceInstancesInnerWithDefaults() *ApiHostGet200ResponseServiceInstancesInner {
	this := ApiHostGet200ResponseServiceInstancesInner{}
	return &this
}

// GetCount returns the Count field value
func (o *ApiHostGet200ResponseServiceInstancesInner) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200ResponseServiceInstancesInner) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ApiHostGet200ResponseServiceInstancesInner) SetCount(v int32) {
	o.Count = v
}

// GetType returns the Type field value
func (o *ApiHostGet200ResponseServiceInstancesInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200ResponseServiceInstancesInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiHostGet200ResponseServiceInstancesInner) SetType(v string) {
	o.Type = v
}

func (o ApiHostGet200ResponseServiceInstancesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHostGet200ResponseServiceInstancesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ApiHostGet200ResponseServiceInstancesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"type",
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

	varApiHostGet200ResponseServiceInstancesInner := _ApiHostGet200ResponseServiceInstancesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiHostGet200ResponseServiceInstancesInner)

	if err != nil {
		return err
	}

	*o = ApiHostGet200ResponseServiceInstancesInner(varApiHostGet200ResponseServiceInstancesInner)

	return err
}

type NullableApiHostGet200ResponseServiceInstancesInner struct {
	value *ApiHostGet200ResponseServiceInstancesInner
	isSet bool
}

func (v NullableApiHostGet200ResponseServiceInstancesInner) Get() *ApiHostGet200ResponseServiceInstancesInner {
	return v.value
}

func (v *NullableApiHostGet200ResponseServiceInstancesInner) Set(val *ApiHostGet200ResponseServiceInstancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHostGet200ResponseServiceInstancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHostGet200ResponseServiceInstancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHostGet200ResponseServiceInstancesInner(val *ApiHostGet200ResponseServiceInstancesInner) *NullableApiHostGet200ResponseServiceInstancesInner {
	return &NullableApiHostGet200ResponseServiceInstancesInner{value: val, isSet: true}
}

func (v NullableApiHostGet200ResponseServiceInstancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHostGet200ResponseServiceInstancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


