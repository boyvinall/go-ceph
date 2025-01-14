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

// checks if the ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner{}

// ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner struct for ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner
type ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner struct {
	// IP Address
	Addr string `json:"addr"`
	// 
	Nonce int32 `json:"nonce"`
	// 
	Type string `json:"type"`
}

type _ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner

// NewApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner instantiates a new ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner(addr string, nonce int32, type_ string) *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner {
	this := ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner{}
	this.Addr = addr
	this.Nonce = nonce
	this.Type = type_
	return &this
}

// NewApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInnerWithDefaults instantiates a new ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInnerWithDefaults() *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner {
	this := ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner{}
	return &this
}

// GetAddr returns the Addr field value
func (o *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) GetAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Addr
}

// GetAddrOk returns a tuple with the Addr field value
// and a boolean to check if the value has been set.
func (o *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) GetAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addr, true
}

// SetAddr sets field value
func (o *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) SetAddr(v string) {
	o.Addr = v
}

// GetNonce returns the Nonce field value
func (o *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) SetNonce(v int32) {
	o.Nonce = v
}

// GetType returns the Type field value
func (o *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) SetType(v string) {
	o.Type = v
}

func (o ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addr"] = o.Addr
	toSerialize["nonce"] = o.Nonce
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addr",
		"nonce",
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

	varApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner := _ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner)

	if err != nil {
		return err
	}

	*o = ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner(varApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner)

	return err
}

type NullableApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner struct {
	value *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner
	isSet bool
}

func (v NullableApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) Get() *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner {
	return v.value
}

func (v *NullableApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) Set(val *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner(val *ApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) *NullableApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner {
	return &NullableApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner{value: val, isSet: true}
}

func (v NullableApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLogsAllGet200ResponseAuditLogInnerAddrsAddrvecInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


