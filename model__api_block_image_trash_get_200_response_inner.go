/*
Ceph REST API

This is the official Ceph REST API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ceph

import (
	"encoding/json"
)

// checks if the ApiBlockImageTrashGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBlockImageTrashGet200ResponseInner{}

// ApiBlockImageTrashGet200ResponseInner struct for ApiBlockImageTrashGet200ResponseInner
type ApiBlockImageTrashGet200ResponseInner struct {
	// pool name
	PoolName *string `json:"pool_name,omitempty"`
	// 
	Status *int32 `json:"status,omitempty"`
	// 
	Value []string `json:"value,omitempty"`
}

// NewApiBlockImageTrashGet200ResponseInner instantiates a new ApiBlockImageTrashGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBlockImageTrashGet200ResponseInner() *ApiBlockImageTrashGet200ResponseInner {
	this := ApiBlockImageTrashGet200ResponseInner{}
	return &this
}

// NewApiBlockImageTrashGet200ResponseInnerWithDefaults instantiates a new ApiBlockImageTrashGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBlockImageTrashGet200ResponseInnerWithDefaults() *ApiBlockImageTrashGet200ResponseInner {
	this := ApiBlockImageTrashGet200ResponseInner{}
	return &this
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *ApiBlockImageTrashGet200ResponseInner) GetPoolName() string {
	if o == nil || IsNil(o.PoolName) {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageTrashGet200ResponseInner) GetPoolNameOk() (*string, bool) {
	if o == nil || IsNil(o.PoolName) {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *ApiBlockImageTrashGet200ResponseInner) HasPoolName() bool {
	if o != nil && !IsNil(o.PoolName) {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *ApiBlockImageTrashGet200ResponseInner) SetPoolName(v string) {
	o.PoolName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiBlockImageTrashGet200ResponseInner) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageTrashGet200ResponseInner) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiBlockImageTrashGet200ResponseInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ApiBlockImageTrashGet200ResponseInner) SetStatus(v int32) {
	o.Status = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiBlockImageTrashGet200ResponseInner) GetValue() []string {
	if o == nil || IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageTrashGet200ResponseInner) GetValueOk() ([]string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiBlockImageTrashGet200ResponseInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *ApiBlockImageTrashGet200ResponseInner) SetValue(v []string) {
	o.Value = v
}

func (o ApiBlockImageTrashGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBlockImageTrashGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PoolName) {
		toSerialize["pool_name"] = o.PoolName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApiBlockImageTrashGet200ResponseInner struct {
	value *ApiBlockImageTrashGet200ResponseInner
	isSet bool
}

func (v NullableApiBlockImageTrashGet200ResponseInner) Get() *ApiBlockImageTrashGet200ResponseInner {
	return v.value
}

func (v *NullableApiBlockImageTrashGet200ResponseInner) Set(val *ApiBlockImageTrashGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBlockImageTrashGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBlockImageTrashGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBlockImageTrashGet200ResponseInner(val *ApiBlockImageTrashGet200ResponseInner) *NullableApiBlockImageTrashGet200ResponseInner {
	return &NullableApiBlockImageTrashGet200ResponseInner{value: val, isSet: true}
}

func (v NullableApiBlockImageTrashGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBlockImageTrashGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


