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

// checks if the ApiHealthMinimalGet200ResponseFsMapFilesystemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHealthMinimalGet200ResponseFsMapFilesystemsInner{}

// ApiHealthMinimalGet200ResponseFsMapFilesystemsInner struct for ApiHealthMinimalGet200ResponseFsMapFilesystemsInner
type ApiHealthMinimalGet200ResponseFsMapFilesystemsInner struct {
	Mdsmap ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap `json:"mdsmap"`
	// 
	Standbys string `json:"standbys"`
}

type _ApiHealthMinimalGet200ResponseFsMapFilesystemsInner ApiHealthMinimalGet200ResponseFsMapFilesystemsInner

// NewApiHealthMinimalGet200ResponseFsMapFilesystemsInner instantiates a new ApiHealthMinimalGet200ResponseFsMapFilesystemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHealthMinimalGet200ResponseFsMapFilesystemsInner(mdsmap ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap, standbys string) *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner {
	this := ApiHealthMinimalGet200ResponseFsMapFilesystemsInner{}
	this.Mdsmap = mdsmap
	this.Standbys = standbys
	return &this
}

// NewApiHealthMinimalGet200ResponseFsMapFilesystemsInnerWithDefaults instantiates a new ApiHealthMinimalGet200ResponseFsMapFilesystemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHealthMinimalGet200ResponseFsMapFilesystemsInnerWithDefaults() *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner {
	this := ApiHealthMinimalGet200ResponseFsMapFilesystemsInner{}
	return &this
}

// GetMdsmap returns the Mdsmap field value
func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) GetMdsmap() ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap
		return ret
	}

	return o.Mdsmap
}

// GetMdsmapOk returns a tuple with the Mdsmap field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) GetMdsmapOk() (*ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mdsmap, true
}

// SetMdsmap sets field value
func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) SetMdsmap(v ApiHealthMinimalGet200ResponseFsMapFilesystemsInnerMdsmap) {
	o.Mdsmap = v
}

// GetStandbys returns the Standbys field value
func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) GetStandbys() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Standbys
}

// GetStandbysOk returns a tuple with the Standbys field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) GetStandbysOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Standbys, true
}

// SetStandbys sets field value
func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) SetStandbys(v string) {
	o.Standbys = v
}

func (o ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mdsmap"] = o.Mdsmap
	toSerialize["standbys"] = o.Standbys
	return toSerialize, nil
}

func (o *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mdsmap",
		"standbys",
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

	varApiHealthMinimalGet200ResponseFsMapFilesystemsInner := _ApiHealthMinimalGet200ResponseFsMapFilesystemsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiHealthMinimalGet200ResponseFsMapFilesystemsInner)

	if err != nil {
		return err
	}

	*o = ApiHealthMinimalGet200ResponseFsMapFilesystemsInner(varApiHealthMinimalGet200ResponseFsMapFilesystemsInner)

	return err
}

type NullableApiHealthMinimalGet200ResponseFsMapFilesystemsInner struct {
	value *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner
	isSet bool
}

func (v NullableApiHealthMinimalGet200ResponseFsMapFilesystemsInner) Get() *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner {
	return v.value
}

func (v *NullableApiHealthMinimalGet200ResponseFsMapFilesystemsInner) Set(val *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHealthMinimalGet200ResponseFsMapFilesystemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHealthMinimalGet200ResponseFsMapFilesystemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHealthMinimalGet200ResponseFsMapFilesystemsInner(val *ApiHealthMinimalGet200ResponseFsMapFilesystemsInner) *NullableApiHealthMinimalGet200ResponseFsMapFilesystemsInner {
	return &NullableApiHealthMinimalGet200ResponseFsMapFilesystemsInner{value: val, isSet: true}
}

func (v NullableApiHealthMinimalGet200ResponseFsMapFilesystemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHealthMinimalGet200ResponseFsMapFilesystemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


