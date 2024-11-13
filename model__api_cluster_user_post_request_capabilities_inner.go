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

// checks if the ApiClusterUserPostRequestCapabilitiesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiClusterUserPostRequestCapabilitiesInner{}

// ApiClusterUserPostRequestCapabilitiesInner struct for ApiClusterUserPostRequestCapabilitiesInner
type ApiClusterUserPostRequestCapabilitiesInner struct {
	// Capability to add; eg. allow *
	Cap string `json:"cap"`
	// Entity to add
	Entity string `json:"entity"`
}

type _ApiClusterUserPostRequestCapabilitiesInner ApiClusterUserPostRequestCapabilitiesInner

// NewApiClusterUserPostRequestCapabilitiesInner instantiates a new ApiClusterUserPostRequestCapabilitiesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiClusterUserPostRequestCapabilitiesInner(cap string, entity string) *ApiClusterUserPostRequestCapabilitiesInner {
	this := ApiClusterUserPostRequestCapabilitiesInner{}
	this.Cap = cap
	this.Entity = entity
	return &this
}

// NewApiClusterUserPostRequestCapabilitiesInnerWithDefaults instantiates a new ApiClusterUserPostRequestCapabilitiesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiClusterUserPostRequestCapabilitiesInnerWithDefaults() *ApiClusterUserPostRequestCapabilitiesInner {
	this := ApiClusterUserPostRequestCapabilitiesInner{}
	return &this
}

// GetCap returns the Cap field value
func (o *ApiClusterUserPostRequestCapabilitiesInner) GetCap() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cap
}

// GetCapOk returns a tuple with the Cap field value
// and a boolean to check if the value has been set.
func (o *ApiClusterUserPostRequestCapabilitiesInner) GetCapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cap, true
}

// SetCap sets field value
func (o *ApiClusterUserPostRequestCapabilitiesInner) SetCap(v string) {
	o.Cap = v
}

// GetEntity returns the Entity field value
func (o *ApiClusterUserPostRequestCapabilitiesInner) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *ApiClusterUserPostRequestCapabilitiesInner) GetEntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *ApiClusterUserPostRequestCapabilitiesInner) SetEntity(v string) {
	o.Entity = v
}

func (o ApiClusterUserPostRequestCapabilitiesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiClusterUserPostRequestCapabilitiesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cap"] = o.Cap
	toSerialize["entity"] = o.Entity
	return toSerialize, nil
}

func (o *ApiClusterUserPostRequestCapabilitiesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cap",
		"entity",
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

	varApiClusterUserPostRequestCapabilitiesInner := _ApiClusterUserPostRequestCapabilitiesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiClusterUserPostRequestCapabilitiesInner)

	if err != nil {
		return err
	}

	*o = ApiClusterUserPostRequestCapabilitiesInner(varApiClusterUserPostRequestCapabilitiesInner)

	return err
}

type NullableApiClusterUserPostRequestCapabilitiesInner struct {
	value *ApiClusterUserPostRequestCapabilitiesInner
	isSet bool
}

func (v NullableApiClusterUserPostRequestCapabilitiesInner) Get() *ApiClusterUserPostRequestCapabilitiesInner {
	return v.value
}

func (v *NullableApiClusterUserPostRequestCapabilitiesInner) Set(val *ApiClusterUserPostRequestCapabilitiesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiClusterUserPostRequestCapabilitiesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiClusterUserPostRequestCapabilitiesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiClusterUserPostRequestCapabilitiesInner(val *ApiClusterUserPostRequestCapabilitiesInner) *NullableApiClusterUserPostRequestCapabilitiesInner {
	return &NullableApiClusterUserPostRequestCapabilitiesInner{value: val, isSet: true}
}

func (v NullableApiClusterUserPostRequestCapabilitiesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiClusterUserPostRequestCapabilitiesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


