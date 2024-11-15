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

// checks if the ApiPoolPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPoolPostRequest{}

// ApiPoolPostRequest struct for ApiPoolPostRequest
type ApiPoolPostRequest struct {
	Pool *string `json:"pool,omitempty"`
}

// NewApiPoolPostRequest instantiates a new ApiPoolPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPoolPostRequest() *ApiPoolPostRequest {
	this := ApiPoolPostRequest{}
	var pool string = "rbd-mirror"
	this.Pool = &pool
	return &this
}

// NewApiPoolPostRequestWithDefaults instantiates a new ApiPoolPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPoolPostRequestWithDefaults() *ApiPoolPostRequest {
	this := ApiPoolPostRequest{}
	var pool string = "rbd-mirror"
	this.Pool = &pool
	return &this
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *ApiPoolPostRequest) GetPool() string {
	if o == nil || IsNil(o.Pool) {
		var ret string
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPoolPostRequest) GetPoolOk() (*string, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *ApiPoolPostRequest) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given string and assigns it to the Pool field.
func (o *ApiPoolPostRequest) SetPool(v string) {
	o.Pool = &v
}

func (o ApiPoolPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPoolPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	return toSerialize, nil
}

type NullableApiPoolPostRequest struct {
	value *ApiPoolPostRequest
	isSet bool
}

func (v NullableApiPoolPostRequest) Get() *ApiPoolPostRequest {
	return v.value
}

func (v *NullableApiPoolPostRequest) Set(val *ApiPoolPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPoolPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPoolPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPoolPostRequest(val *ApiPoolPostRequest) *NullableApiPoolPostRequest {
	return &NullableApiPoolPostRequest{value: val, isSet: true}
}

func (v NullableApiPoolPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPoolPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


