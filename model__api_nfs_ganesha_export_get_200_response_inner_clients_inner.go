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

// checks if the ApiNfsGaneshaExportGet200ResponseInnerClientsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiNfsGaneshaExportGet200ResponseInnerClientsInner{}

// ApiNfsGaneshaExportGet200ResponseInnerClientsInner struct for ApiNfsGaneshaExportGet200ResponseInnerClientsInner
type ApiNfsGaneshaExportGet200ResponseInnerClientsInner struct {
	// Client access type
	AccessType string `json:"access_type"`
	// list of IP addresses
	Addresses []string `json:"addresses"`
	// Client squash policy
	Squash string `json:"squash"`
}

type _ApiNfsGaneshaExportGet200ResponseInnerClientsInner ApiNfsGaneshaExportGet200ResponseInnerClientsInner

// NewApiNfsGaneshaExportGet200ResponseInnerClientsInner instantiates a new ApiNfsGaneshaExportGet200ResponseInnerClientsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNfsGaneshaExportGet200ResponseInnerClientsInner(accessType string, addresses []string, squash string) *ApiNfsGaneshaExportGet200ResponseInnerClientsInner {
	this := ApiNfsGaneshaExportGet200ResponseInnerClientsInner{}
	this.AccessType = accessType
	this.Addresses = addresses
	this.Squash = squash
	return &this
}

// NewApiNfsGaneshaExportGet200ResponseInnerClientsInnerWithDefaults instantiates a new ApiNfsGaneshaExportGet200ResponseInnerClientsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNfsGaneshaExportGet200ResponseInnerClientsInnerWithDefaults() *ApiNfsGaneshaExportGet200ResponseInnerClientsInner {
	this := ApiNfsGaneshaExportGet200ResponseInnerClientsInner{}
	return &this
}

// GetAccessType returns the AccessType field value
func (o *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) GetAccessType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) GetAccessTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) SetAccessType(v string) {
	o.AccessType = v
}

// GetAddresses returns the Addresses field value
func (o *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) SetAddresses(v []string) {
	o.Addresses = v
}

// GetSquash returns the Squash field value
func (o *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) GetSquash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Squash
}

// GetSquashOk returns a tuple with the Squash field value
// and a boolean to check if the value has been set.
func (o *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) GetSquashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Squash, true
}

// SetSquash sets field value
func (o *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) SetSquash(v string) {
	o.Squash = v
}

func (o ApiNfsGaneshaExportGet200ResponseInnerClientsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNfsGaneshaExportGet200ResponseInnerClientsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_type"] = o.AccessType
	toSerialize["addresses"] = o.Addresses
	toSerialize["squash"] = o.Squash
	return toSerialize, nil
}

func (o *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_type",
		"addresses",
		"squash",
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

	varApiNfsGaneshaExportGet200ResponseInnerClientsInner := _ApiNfsGaneshaExportGet200ResponseInnerClientsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiNfsGaneshaExportGet200ResponseInnerClientsInner)

	if err != nil {
		return err
	}

	*o = ApiNfsGaneshaExportGet200ResponseInnerClientsInner(varApiNfsGaneshaExportGet200ResponseInnerClientsInner)

	return err
}

type NullableApiNfsGaneshaExportGet200ResponseInnerClientsInner struct {
	value *ApiNfsGaneshaExportGet200ResponseInnerClientsInner
	isSet bool
}

func (v NullableApiNfsGaneshaExportGet200ResponseInnerClientsInner) Get() *ApiNfsGaneshaExportGet200ResponseInnerClientsInner {
	return v.value
}

func (v *NullableApiNfsGaneshaExportGet200ResponseInnerClientsInner) Set(val *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNfsGaneshaExportGet200ResponseInnerClientsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNfsGaneshaExportGet200ResponseInnerClientsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNfsGaneshaExportGet200ResponseInnerClientsInner(val *ApiNfsGaneshaExportGet200ResponseInnerClientsInner) *NullableApiNfsGaneshaExportGet200ResponseInnerClientsInner {
	return &NullableApiNfsGaneshaExportGet200ResponseInnerClientsInner{value: val, isSet: true}
}

func (v NullableApiNfsGaneshaExportGet200ResponseInnerClientsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNfsGaneshaExportGet200ResponseInnerClientsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

