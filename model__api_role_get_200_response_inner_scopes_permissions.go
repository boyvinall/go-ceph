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

// checks if the ApiRoleGet200ResponseInnerScopesPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRoleGet200ResponseInnerScopesPermissions{}

// ApiRoleGet200ResponseInnerScopesPermissions 
type ApiRoleGet200ResponseInnerScopesPermissions struct {
	// 
	Cephfs []string `json:"cephfs"`
}

type _ApiRoleGet200ResponseInnerScopesPermissions ApiRoleGet200ResponseInnerScopesPermissions

// NewApiRoleGet200ResponseInnerScopesPermissions instantiates a new ApiRoleGet200ResponseInnerScopesPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRoleGet200ResponseInnerScopesPermissions(cephfs []string) *ApiRoleGet200ResponseInnerScopesPermissions {
	this := ApiRoleGet200ResponseInnerScopesPermissions{}
	this.Cephfs = cephfs
	return &this
}

// NewApiRoleGet200ResponseInnerScopesPermissionsWithDefaults instantiates a new ApiRoleGet200ResponseInnerScopesPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRoleGet200ResponseInnerScopesPermissionsWithDefaults() *ApiRoleGet200ResponseInnerScopesPermissions {
	this := ApiRoleGet200ResponseInnerScopesPermissions{}
	return &this
}

// GetCephfs returns the Cephfs field value
func (o *ApiRoleGet200ResponseInnerScopesPermissions) GetCephfs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Cephfs
}

// GetCephfsOk returns a tuple with the Cephfs field value
// and a boolean to check if the value has been set.
func (o *ApiRoleGet200ResponseInnerScopesPermissions) GetCephfsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cephfs, true
}

// SetCephfs sets field value
func (o *ApiRoleGet200ResponseInnerScopesPermissions) SetCephfs(v []string) {
	o.Cephfs = v
}

func (o ApiRoleGet200ResponseInnerScopesPermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRoleGet200ResponseInnerScopesPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cephfs"] = o.Cephfs
	return toSerialize, nil
}

func (o *ApiRoleGet200ResponseInnerScopesPermissions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cephfs",
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

	varApiRoleGet200ResponseInnerScopesPermissions := _ApiRoleGet200ResponseInnerScopesPermissions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiRoleGet200ResponseInnerScopesPermissions)

	if err != nil {
		return err
	}

	*o = ApiRoleGet200ResponseInnerScopesPermissions(varApiRoleGet200ResponseInnerScopesPermissions)

	return err
}

type NullableApiRoleGet200ResponseInnerScopesPermissions struct {
	value *ApiRoleGet200ResponseInnerScopesPermissions
	isSet bool
}

func (v NullableApiRoleGet200ResponseInnerScopesPermissions) Get() *ApiRoleGet200ResponseInnerScopesPermissions {
	return v.value
}

func (v *NullableApiRoleGet200ResponseInnerScopesPermissions) Set(val *ApiRoleGet200ResponseInnerScopesPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRoleGet200ResponseInnerScopesPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRoleGet200ResponseInnerScopesPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRoleGet200ResponseInnerScopesPermissions(val *ApiRoleGet200ResponseInnerScopesPermissions) *NullableApiRoleGet200ResponseInnerScopesPermissions {
	return &NullableApiRoleGet200ResponseInnerScopesPermissions{value: val, isSet: true}
}

func (v NullableApiRoleGet200ResponseInnerScopesPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRoleGet200ResponseInnerScopesPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

