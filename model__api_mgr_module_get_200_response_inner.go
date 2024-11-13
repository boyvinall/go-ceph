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

// checks if the ApiMgrModuleGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMgrModuleGet200ResponseInner{}

// ApiMgrModuleGet200ResponseInner struct for ApiMgrModuleGet200ResponseInner
type ApiMgrModuleGet200ResponseInner struct {
	// Is it an always on module?
	AlwaysOn *bool `json:"always_on,omitempty"`
	// Is Module Enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Module Name
	Name *string `json:"name,omitempty"`
	Options *ApiMgrModuleGet200ResponseInnerOptions `json:"options,omitempty"`
}

// NewApiMgrModuleGet200ResponseInner instantiates a new ApiMgrModuleGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMgrModuleGet200ResponseInner() *ApiMgrModuleGet200ResponseInner {
	this := ApiMgrModuleGet200ResponseInner{}
	return &this
}

// NewApiMgrModuleGet200ResponseInnerWithDefaults instantiates a new ApiMgrModuleGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMgrModuleGet200ResponseInnerWithDefaults() *ApiMgrModuleGet200ResponseInner {
	this := ApiMgrModuleGet200ResponseInner{}
	return &this
}

// GetAlwaysOn returns the AlwaysOn field value if set, zero value otherwise.
func (o *ApiMgrModuleGet200ResponseInner) GetAlwaysOn() bool {
	if o == nil || IsNil(o.AlwaysOn) {
		var ret bool
		return ret
	}
	return *o.AlwaysOn
}

// GetAlwaysOnOk returns a tuple with the AlwaysOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInner) GetAlwaysOnOk() (*bool, bool) {
	if o == nil || IsNil(o.AlwaysOn) {
		return nil, false
	}
	return o.AlwaysOn, true
}

// HasAlwaysOn returns a boolean if a field has been set.
func (o *ApiMgrModuleGet200ResponseInner) HasAlwaysOn() bool {
	if o != nil && !IsNil(o.AlwaysOn) {
		return true
	}

	return false
}

// SetAlwaysOn gets a reference to the given bool and assigns it to the AlwaysOn field.
func (o *ApiMgrModuleGet200ResponseInner) SetAlwaysOn(v bool) {
	o.AlwaysOn = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiMgrModuleGet200ResponseInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiMgrModuleGet200ResponseInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiMgrModuleGet200ResponseInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiMgrModuleGet200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiMgrModuleGet200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiMgrModuleGet200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ApiMgrModuleGet200ResponseInner) GetOptions() ApiMgrModuleGet200ResponseInnerOptions {
	if o == nil || IsNil(o.Options) {
		var ret ApiMgrModuleGet200ResponseInnerOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInner) GetOptionsOk() (*ApiMgrModuleGet200ResponseInnerOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApiMgrModuleGet200ResponseInner) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ApiMgrModuleGet200ResponseInnerOptions and assigns it to the Options field.
func (o *ApiMgrModuleGet200ResponseInner) SetOptions(v ApiMgrModuleGet200ResponseInnerOptions) {
	o.Options = &v
}

func (o ApiMgrModuleGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMgrModuleGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlwaysOn) {
		toSerialize["always_on"] = o.AlwaysOn
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableApiMgrModuleGet200ResponseInner struct {
	value *ApiMgrModuleGet200ResponseInner
	isSet bool
}

func (v NullableApiMgrModuleGet200ResponseInner) Get() *ApiMgrModuleGet200ResponseInner {
	return v.value
}

func (v *NullableApiMgrModuleGet200ResponseInner) Set(val *ApiMgrModuleGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMgrModuleGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMgrModuleGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMgrModuleGet200ResponseInner(val *ApiMgrModuleGet200ResponseInner) *NullableApiMgrModuleGet200ResponseInner {
	return &NullableApiMgrModuleGet200ResponseInner{value: val, isSet: true}
}

func (v NullableApiMgrModuleGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMgrModuleGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

