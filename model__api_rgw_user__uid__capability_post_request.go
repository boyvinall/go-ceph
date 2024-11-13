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

// checks if the ApiRgwUserUidCapabilityPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRgwUserUidCapabilityPostRequest{}

// ApiRgwUserUidCapabilityPostRequest struct for ApiRgwUserUidCapabilityPostRequest
type ApiRgwUserUidCapabilityPostRequest struct {
	DaemonName *string `json:"daemon_name,omitempty"`
	Perm string `json:"perm"`
	Type string `json:"type"`
}

type _ApiRgwUserUidCapabilityPostRequest ApiRgwUserUidCapabilityPostRequest

// NewApiRgwUserUidCapabilityPostRequest instantiates a new ApiRgwUserUidCapabilityPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRgwUserUidCapabilityPostRequest(perm string, type_ string) *ApiRgwUserUidCapabilityPostRequest {
	this := ApiRgwUserUidCapabilityPostRequest{}
	this.Perm = perm
	this.Type = type_
	return &this
}

// NewApiRgwUserUidCapabilityPostRequestWithDefaults instantiates a new ApiRgwUserUidCapabilityPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRgwUserUidCapabilityPostRequestWithDefaults() *ApiRgwUserUidCapabilityPostRequest {
	this := ApiRgwUserUidCapabilityPostRequest{}
	return &this
}

// GetDaemonName returns the DaemonName field value if set, zero value otherwise.
func (o *ApiRgwUserUidCapabilityPostRequest) GetDaemonName() string {
	if o == nil || IsNil(o.DaemonName) {
		var ret string
		return ret
	}
	return *o.DaemonName
}

// GetDaemonNameOk returns a tuple with the DaemonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidCapabilityPostRequest) GetDaemonNameOk() (*string, bool) {
	if o == nil || IsNil(o.DaemonName) {
		return nil, false
	}
	return o.DaemonName, true
}

// HasDaemonName returns a boolean if a field has been set.
func (o *ApiRgwUserUidCapabilityPostRequest) HasDaemonName() bool {
	if o != nil && !IsNil(o.DaemonName) {
		return true
	}

	return false
}

// SetDaemonName gets a reference to the given string and assigns it to the DaemonName field.
func (o *ApiRgwUserUidCapabilityPostRequest) SetDaemonName(v string) {
	o.DaemonName = &v
}

// GetPerm returns the Perm field value
func (o *ApiRgwUserUidCapabilityPostRequest) GetPerm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Perm
}

// GetPermOk returns a tuple with the Perm field value
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidCapabilityPostRequest) GetPermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Perm, true
}

// SetPerm sets field value
func (o *ApiRgwUserUidCapabilityPostRequest) SetPerm(v string) {
	o.Perm = v
}

// GetType returns the Type field value
func (o *ApiRgwUserUidCapabilityPostRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidCapabilityPostRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiRgwUserUidCapabilityPostRequest) SetType(v string) {
	o.Type = v
}

func (o ApiRgwUserUidCapabilityPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRgwUserUidCapabilityPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DaemonName) {
		toSerialize["daemon_name"] = o.DaemonName
	}
	toSerialize["perm"] = o.Perm
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ApiRgwUserUidCapabilityPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"perm",
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

	varApiRgwUserUidCapabilityPostRequest := _ApiRgwUserUidCapabilityPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiRgwUserUidCapabilityPostRequest)

	if err != nil {
		return err
	}

	*o = ApiRgwUserUidCapabilityPostRequest(varApiRgwUserUidCapabilityPostRequest)

	return err
}

type NullableApiRgwUserUidCapabilityPostRequest struct {
	value *ApiRgwUserUidCapabilityPostRequest
	isSet bool
}

func (v NullableApiRgwUserUidCapabilityPostRequest) Get() *ApiRgwUserUidCapabilityPostRequest {
	return v.value
}

func (v *NullableApiRgwUserUidCapabilityPostRequest) Set(val *ApiRgwUserUidCapabilityPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRgwUserUidCapabilityPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRgwUserUidCapabilityPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRgwUserUidCapabilityPostRequest(val *ApiRgwUserUidCapabilityPostRequest) *NullableApiRgwUserUidCapabilityPostRequest {
	return &NullableApiRgwUserUidCapabilityPostRequest{value: val, isSet: true}
}

func (v NullableApiRgwUserUidCapabilityPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRgwUserUidCapabilityPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


