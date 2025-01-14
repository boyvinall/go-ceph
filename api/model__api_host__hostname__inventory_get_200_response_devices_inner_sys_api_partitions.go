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

// checks if the ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions{}

// ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions 
type ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions struct {
	PartitionName ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitionsPartitionName `json:"partition_name"`
}

type _ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions

// NewApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions instantiates a new ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions(partitionName ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitionsPartitionName) *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions {
	this := ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions{}
	this.PartitionName = partitionName
	return &this
}

// NewApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitionsWithDefaults instantiates a new ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitionsWithDefaults() *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions {
	this := ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions{}
	return &this
}

// GetPartitionName returns the PartitionName field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) GetPartitionName() ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitionsPartitionName {
	if o == nil {
		var ret ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitionsPartitionName
		return ret
	}

	return o.PartitionName
}

// GetPartitionNameOk returns a tuple with the PartitionName field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) GetPartitionNameOk() (*ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitionsPartitionName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionName, true
}

// SetPartitionName sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) SetPartitionName(v ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitionsPartitionName) {
	o.PartitionName = v
}

func (o ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["partition_name"] = o.PartitionName
	return toSerialize, nil
}

func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"partition_name",
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

	varApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions := _ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions)

	if err != nil {
		return err
	}

	*o = ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions(varApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions)

	return err
}

type NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions struct {
	value *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions
	isSet bool
}

func (v NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) Get() *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions {
	return v.value
}

func (v *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) Set(val *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions(val *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions {
	return &NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions{value: val, isSet: true}
}

func (v NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


