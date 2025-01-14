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

// checks if the ApiTelemetryReportGet200ResponseReportFsFeatureFlags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportFsFeatureFlags{}

// ApiTelemetryReportGet200ResponseReportFsFeatureFlags 
type ApiTelemetryReportGet200ResponseReportFsFeatureFlags struct {
	// 
	EnableMultiple bool `json:"enable_multiple"`
	// 
	EverEnabledMultiple bool `json:"ever_enabled_multiple"`
}

type _ApiTelemetryReportGet200ResponseReportFsFeatureFlags ApiTelemetryReportGet200ResponseReportFsFeatureFlags

// NewApiTelemetryReportGet200ResponseReportFsFeatureFlags instantiates a new ApiTelemetryReportGet200ResponseReportFsFeatureFlags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportFsFeatureFlags(enableMultiple bool, everEnabledMultiple bool) *ApiTelemetryReportGet200ResponseReportFsFeatureFlags {
	this := ApiTelemetryReportGet200ResponseReportFsFeatureFlags{}
	this.EnableMultiple = enableMultiple
	this.EverEnabledMultiple = everEnabledMultiple
	return &this
}

// NewApiTelemetryReportGet200ResponseReportFsFeatureFlagsWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportFsFeatureFlags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportFsFeatureFlagsWithDefaults() *ApiTelemetryReportGet200ResponseReportFsFeatureFlags {
	this := ApiTelemetryReportGet200ResponseReportFsFeatureFlags{}
	return &this
}

// GetEnableMultiple returns the EnableMultiple field value
func (o *ApiTelemetryReportGet200ResponseReportFsFeatureFlags) GetEnableMultiple() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableMultiple
}

// GetEnableMultipleOk returns a tuple with the EnableMultiple field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportFsFeatureFlags) GetEnableMultipleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableMultiple, true
}

// SetEnableMultiple sets field value
func (o *ApiTelemetryReportGet200ResponseReportFsFeatureFlags) SetEnableMultiple(v bool) {
	o.EnableMultiple = v
}

// GetEverEnabledMultiple returns the EverEnabledMultiple field value
func (o *ApiTelemetryReportGet200ResponseReportFsFeatureFlags) GetEverEnabledMultiple() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EverEnabledMultiple
}

// GetEverEnabledMultipleOk returns a tuple with the EverEnabledMultiple field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportFsFeatureFlags) GetEverEnabledMultipleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EverEnabledMultiple, true
}

// SetEverEnabledMultiple sets field value
func (o *ApiTelemetryReportGet200ResponseReportFsFeatureFlags) SetEverEnabledMultiple(v bool) {
	o.EverEnabledMultiple = v
}

func (o ApiTelemetryReportGet200ResponseReportFsFeatureFlags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportFsFeatureFlags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable_multiple"] = o.EnableMultiple
	toSerialize["ever_enabled_multiple"] = o.EverEnabledMultiple
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportFsFeatureFlags) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enable_multiple",
		"ever_enabled_multiple",
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

	varApiTelemetryReportGet200ResponseReportFsFeatureFlags := _ApiTelemetryReportGet200ResponseReportFsFeatureFlags{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportFsFeatureFlags)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportFsFeatureFlags(varApiTelemetryReportGet200ResponseReportFsFeatureFlags)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportFsFeatureFlags struct {
	value *ApiTelemetryReportGet200ResponseReportFsFeatureFlags
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportFsFeatureFlags) Get() *ApiTelemetryReportGet200ResponseReportFsFeatureFlags {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportFsFeatureFlags) Set(val *ApiTelemetryReportGet200ResponseReportFsFeatureFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportFsFeatureFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportFsFeatureFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportFsFeatureFlags(val *ApiTelemetryReportGet200ResponseReportFsFeatureFlags) *NullableApiTelemetryReportGet200ResponseReportFsFeatureFlags {
	return &NullableApiTelemetryReportGet200ResponseReportFsFeatureFlags{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportFsFeatureFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportFsFeatureFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


