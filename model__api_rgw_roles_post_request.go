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

// checks if the ApiRgwRolesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRgwRolesPostRequest{}

// ApiRgwRolesPostRequest struct for ApiRgwRolesPostRequest
type ApiRgwRolesPostRequest struct {
	RoleAssumePolicyDoc *string `json:"role_assume_policy_doc,omitempty"`
	RoleName *string `json:"role_name,omitempty"`
	RolePath *string `json:"role_path,omitempty"`
}

// NewApiRgwRolesPostRequest instantiates a new ApiRgwRolesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRgwRolesPostRequest() *ApiRgwRolesPostRequest {
	this := ApiRgwRolesPostRequest{}
	var roleAssumePolicyDoc string = ""
	this.RoleAssumePolicyDoc = &roleAssumePolicyDoc
	var roleName string = ""
	this.RoleName = &roleName
	var rolePath string = ""
	this.RolePath = &rolePath
	return &this
}

// NewApiRgwRolesPostRequestWithDefaults instantiates a new ApiRgwRolesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRgwRolesPostRequestWithDefaults() *ApiRgwRolesPostRequest {
	this := ApiRgwRolesPostRequest{}
	var roleAssumePolicyDoc string = ""
	this.RoleAssumePolicyDoc = &roleAssumePolicyDoc
	var roleName string = ""
	this.RoleName = &roleName
	var rolePath string = ""
	this.RolePath = &rolePath
	return &this
}

// GetRoleAssumePolicyDoc returns the RoleAssumePolicyDoc field value if set, zero value otherwise.
func (o *ApiRgwRolesPostRequest) GetRoleAssumePolicyDoc() string {
	if o == nil || IsNil(o.RoleAssumePolicyDoc) {
		var ret string
		return ret
	}
	return *o.RoleAssumePolicyDoc
}

// GetRoleAssumePolicyDocOk returns a tuple with the RoleAssumePolicyDoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwRolesPostRequest) GetRoleAssumePolicyDocOk() (*string, bool) {
	if o == nil || IsNil(o.RoleAssumePolicyDoc) {
		return nil, false
	}
	return o.RoleAssumePolicyDoc, true
}

// HasRoleAssumePolicyDoc returns a boolean if a field has been set.
func (o *ApiRgwRolesPostRequest) HasRoleAssumePolicyDoc() bool {
	if o != nil && !IsNil(o.RoleAssumePolicyDoc) {
		return true
	}

	return false
}

// SetRoleAssumePolicyDoc gets a reference to the given string and assigns it to the RoleAssumePolicyDoc field.
func (o *ApiRgwRolesPostRequest) SetRoleAssumePolicyDoc(v string) {
	o.RoleAssumePolicyDoc = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *ApiRgwRolesPostRequest) GetRoleName() string {
	if o == nil || IsNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwRolesPostRequest) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *ApiRgwRolesPostRequest) HasRoleName() bool {
	if o != nil && !IsNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *ApiRgwRolesPostRequest) SetRoleName(v string) {
	o.RoleName = &v
}

// GetRolePath returns the RolePath field value if set, zero value otherwise.
func (o *ApiRgwRolesPostRequest) GetRolePath() string {
	if o == nil || IsNil(o.RolePath) {
		var ret string
		return ret
	}
	return *o.RolePath
}

// GetRolePathOk returns a tuple with the RolePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwRolesPostRequest) GetRolePathOk() (*string, bool) {
	if o == nil || IsNil(o.RolePath) {
		return nil, false
	}
	return o.RolePath, true
}

// HasRolePath returns a boolean if a field has been set.
func (o *ApiRgwRolesPostRequest) HasRolePath() bool {
	if o != nil && !IsNil(o.RolePath) {
		return true
	}

	return false
}

// SetRolePath gets a reference to the given string and assigns it to the RolePath field.
func (o *ApiRgwRolesPostRequest) SetRolePath(v string) {
	o.RolePath = &v
}

func (o ApiRgwRolesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRgwRolesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoleAssumePolicyDoc) {
		toSerialize["role_assume_policy_doc"] = o.RoleAssumePolicyDoc
	}
	if !IsNil(o.RoleName) {
		toSerialize["role_name"] = o.RoleName
	}
	if !IsNil(o.RolePath) {
		toSerialize["role_path"] = o.RolePath
	}
	return toSerialize, nil
}

type NullableApiRgwRolesPostRequest struct {
	value *ApiRgwRolesPostRequest
	isSet bool
}

func (v NullableApiRgwRolesPostRequest) Get() *ApiRgwRolesPostRequest {
	return v.value
}

func (v *NullableApiRgwRolesPostRequest) Set(val *ApiRgwRolesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRgwRolesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRgwRolesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRgwRolesPostRequest(val *ApiRgwRolesPostRequest) *NullableApiRgwRolesPostRequest {
	return &NullableApiRgwRolesPostRequest{value: val, isSet: true}
}

func (v NullableApiRgwRolesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRgwRolesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


