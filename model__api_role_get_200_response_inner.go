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

// checks if the ApiRoleGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRoleGet200ResponseInner{}

// ApiRoleGet200ResponseInner struct for ApiRoleGet200ResponseInner
type ApiRoleGet200ResponseInner struct {
	// Role Descriptions
	Description *string `json:"description,omitempty"`
	// Role Name
	Name *string `json:"name,omitempty"`
	ScopesPermissions *ApiRoleGet200ResponseInnerScopesPermissions `json:"scopes_permissions,omitempty"`
	// 
	System *bool `json:"system,omitempty"`
}

// NewApiRoleGet200ResponseInner instantiates a new ApiRoleGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRoleGet200ResponseInner() *ApiRoleGet200ResponseInner {
	this := ApiRoleGet200ResponseInner{}
	return &this
}

// NewApiRoleGet200ResponseInnerWithDefaults instantiates a new ApiRoleGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRoleGet200ResponseInnerWithDefaults() *ApiRoleGet200ResponseInner {
	this := ApiRoleGet200ResponseInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiRoleGet200ResponseInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleGet200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiRoleGet200ResponseInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiRoleGet200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiRoleGet200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleGet200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiRoleGet200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiRoleGet200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetScopesPermissions returns the ScopesPermissions field value if set, zero value otherwise.
func (o *ApiRoleGet200ResponseInner) GetScopesPermissions() ApiRoleGet200ResponseInnerScopesPermissions {
	if o == nil || IsNil(o.ScopesPermissions) {
		var ret ApiRoleGet200ResponseInnerScopesPermissions
		return ret
	}
	return *o.ScopesPermissions
}

// GetScopesPermissionsOk returns a tuple with the ScopesPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleGet200ResponseInner) GetScopesPermissionsOk() (*ApiRoleGet200ResponseInnerScopesPermissions, bool) {
	if o == nil || IsNil(o.ScopesPermissions) {
		return nil, false
	}
	return o.ScopesPermissions, true
}

// HasScopesPermissions returns a boolean if a field has been set.
func (o *ApiRoleGet200ResponseInner) HasScopesPermissions() bool {
	if o != nil && !IsNil(o.ScopesPermissions) {
		return true
	}

	return false
}

// SetScopesPermissions gets a reference to the given ApiRoleGet200ResponseInnerScopesPermissions and assigns it to the ScopesPermissions field.
func (o *ApiRoleGet200ResponseInner) SetScopesPermissions(v ApiRoleGet200ResponseInnerScopesPermissions) {
	o.ScopesPermissions = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *ApiRoleGet200ResponseInner) GetSystem() bool {
	if o == nil || IsNil(o.System) {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleGet200ResponseInner) GetSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *ApiRoleGet200ResponseInner) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *ApiRoleGet200ResponseInner) SetSystem(v bool) {
	o.System = &v
}

func (o ApiRoleGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRoleGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ScopesPermissions) {
		toSerialize["scopes_permissions"] = o.ScopesPermissions
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	return toSerialize, nil
}

type NullableApiRoleGet200ResponseInner struct {
	value *ApiRoleGet200ResponseInner
	isSet bool
}

func (v NullableApiRoleGet200ResponseInner) Get() *ApiRoleGet200ResponseInner {
	return v.value
}

func (v *NullableApiRoleGet200ResponseInner) Set(val *ApiRoleGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRoleGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRoleGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRoleGet200ResponseInner(val *ApiRoleGet200ResponseInner) *NullableApiRoleGet200ResponseInner {
	return &NullableApiRoleGet200ResponseInner{value: val, isSet: true}
}

func (v NullableApiRoleGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRoleGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


