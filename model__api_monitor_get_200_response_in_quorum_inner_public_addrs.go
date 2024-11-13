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

// checks if the ApiMonitorGet200ResponseInQuorumInnerPublicAddrs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMonitorGet200ResponseInQuorumInnerPublicAddrs{}

// ApiMonitorGet200ResponseInQuorumInnerPublicAddrs 
type ApiMonitorGet200ResponseInQuorumInnerPublicAddrs struct {
	// 
	Addrvec []ApiMonitorGet200ResponseInQuorumInnerPublicAddrsAddrvecInner `json:"addrvec"`
}

type _ApiMonitorGet200ResponseInQuorumInnerPublicAddrs ApiMonitorGet200ResponseInQuorumInnerPublicAddrs

// NewApiMonitorGet200ResponseInQuorumInnerPublicAddrs instantiates a new ApiMonitorGet200ResponseInQuorumInnerPublicAddrs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMonitorGet200ResponseInQuorumInnerPublicAddrs(addrvec []ApiMonitorGet200ResponseInQuorumInnerPublicAddrsAddrvecInner) *ApiMonitorGet200ResponseInQuorumInnerPublicAddrs {
	this := ApiMonitorGet200ResponseInQuorumInnerPublicAddrs{}
	this.Addrvec = addrvec
	return &this
}

// NewApiMonitorGet200ResponseInQuorumInnerPublicAddrsWithDefaults instantiates a new ApiMonitorGet200ResponseInQuorumInnerPublicAddrs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMonitorGet200ResponseInQuorumInnerPublicAddrsWithDefaults() *ApiMonitorGet200ResponseInQuorumInnerPublicAddrs {
	this := ApiMonitorGet200ResponseInQuorumInnerPublicAddrs{}
	return &this
}

// GetAddrvec returns the Addrvec field value
func (o *ApiMonitorGet200ResponseInQuorumInnerPublicAddrs) GetAddrvec() []ApiMonitorGet200ResponseInQuorumInnerPublicAddrsAddrvecInner {
	if o == nil {
		var ret []ApiMonitorGet200ResponseInQuorumInnerPublicAddrsAddrvecInner
		return ret
	}

	return o.Addrvec
}

// GetAddrvecOk returns a tuple with the Addrvec field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseInQuorumInnerPublicAddrs) GetAddrvecOk() ([]ApiMonitorGet200ResponseInQuorumInnerPublicAddrsAddrvecInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addrvec, true
}

// SetAddrvec sets field value
func (o *ApiMonitorGet200ResponseInQuorumInnerPublicAddrs) SetAddrvec(v []ApiMonitorGet200ResponseInQuorumInnerPublicAddrsAddrvecInner) {
	o.Addrvec = v
}

func (o ApiMonitorGet200ResponseInQuorumInnerPublicAddrs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMonitorGet200ResponseInQuorumInnerPublicAddrs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addrvec"] = o.Addrvec
	return toSerialize, nil
}

func (o *ApiMonitorGet200ResponseInQuorumInnerPublicAddrs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addrvec",
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

	varApiMonitorGet200ResponseInQuorumInnerPublicAddrs := _ApiMonitorGet200ResponseInQuorumInnerPublicAddrs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiMonitorGet200ResponseInQuorumInnerPublicAddrs)

	if err != nil {
		return err
	}

	*o = ApiMonitorGet200ResponseInQuorumInnerPublicAddrs(varApiMonitorGet200ResponseInQuorumInnerPublicAddrs)

	return err
}

type NullableApiMonitorGet200ResponseInQuorumInnerPublicAddrs struct {
	value *ApiMonitorGet200ResponseInQuorumInnerPublicAddrs
	isSet bool
}

func (v NullableApiMonitorGet200ResponseInQuorumInnerPublicAddrs) Get() *ApiMonitorGet200ResponseInQuorumInnerPublicAddrs {
	return v.value
}

func (v *NullableApiMonitorGet200ResponseInQuorumInnerPublicAddrs) Set(val *ApiMonitorGet200ResponseInQuorumInnerPublicAddrs) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMonitorGet200ResponseInQuorumInnerPublicAddrs) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMonitorGet200ResponseInQuorumInnerPublicAddrs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMonitorGet200ResponseInQuorumInnerPublicAddrs(val *ApiMonitorGet200ResponseInQuorumInnerPublicAddrs) *NullableApiMonitorGet200ResponseInQuorumInnerPublicAddrs {
	return &NullableApiMonitorGet200ResponseInQuorumInnerPublicAddrs{value: val, isSet: true}
}

func (v NullableApiMonitorGet200ResponseInQuorumInnerPublicAddrs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMonitorGet200ResponseInQuorumInnerPublicAddrs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


