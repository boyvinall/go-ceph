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

// checks if the ApiPerfCountersGet200ResponseMonACacheBytes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPerfCountersGet200ResponseMonACacheBytes{}

// ApiPerfCountersGet200ResponseMonACacheBytes 
type ApiPerfCountersGet200ResponseMonACacheBytes struct {
	// 
	Description string `json:"description"`
	// 
	Nick string `json:"nick"`
	// 
	Priority int32 `json:"priority"`
	// 
	Type int32 `json:"type"`
	// 
	Units int32 `json:"units"`
	// 
	Value int32 `json:"value"`
}

type _ApiPerfCountersGet200ResponseMonACacheBytes ApiPerfCountersGet200ResponseMonACacheBytes

// NewApiPerfCountersGet200ResponseMonACacheBytes instantiates a new ApiPerfCountersGet200ResponseMonACacheBytes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPerfCountersGet200ResponseMonACacheBytes(description string, nick string, priority int32, type_ int32, units int32, value int32) *ApiPerfCountersGet200ResponseMonACacheBytes {
	this := ApiPerfCountersGet200ResponseMonACacheBytes{}
	this.Description = description
	this.Nick = nick
	this.Priority = priority
	this.Type = type_
	this.Units = units
	this.Value = value
	return &this
}

// NewApiPerfCountersGet200ResponseMonACacheBytesWithDefaults instantiates a new ApiPerfCountersGet200ResponseMonACacheBytes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPerfCountersGet200ResponseMonACacheBytesWithDefaults() *ApiPerfCountersGet200ResponseMonACacheBytes {
	this := ApiPerfCountersGet200ResponseMonACacheBytes{}
	return &this
}

// GetDescription returns the Description field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) SetDescription(v string) {
	o.Description = v
}

// GetNick returns the Nick field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetNick() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nick
}

// GetNickOk returns a tuple with the Nick field value
// and a boolean to check if the value has been set.
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetNickOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nick, true
}

// SetNick sets field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) SetNick(v string) {
	o.Nick = v
}

// GetPriority returns the Priority field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) SetPriority(v int32) {
	o.Priority = v
}

// GetType returns the Type field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) SetType(v int32) {
	o.Type = v
}

// GetUnits returns the Units field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) SetUnits(v int32) {
	o.Units = v
}

// GetValue returns the Value field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ApiPerfCountersGet200ResponseMonACacheBytes) SetValue(v int32) {
	o.Value = v
}

func (o ApiPerfCountersGet200ResponseMonACacheBytes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPerfCountersGet200ResponseMonACacheBytes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["nick"] = o.Nick
	toSerialize["priority"] = o.Priority
	toSerialize["type"] = o.Type
	toSerialize["units"] = o.Units
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ApiPerfCountersGet200ResponseMonACacheBytes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"nick",
		"priority",
		"type",
		"units",
		"value",
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

	varApiPerfCountersGet200ResponseMonACacheBytes := _ApiPerfCountersGet200ResponseMonACacheBytes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiPerfCountersGet200ResponseMonACacheBytes)

	if err != nil {
		return err
	}

	*o = ApiPerfCountersGet200ResponseMonACacheBytes(varApiPerfCountersGet200ResponseMonACacheBytes)

	return err
}

type NullableApiPerfCountersGet200ResponseMonACacheBytes struct {
	value *ApiPerfCountersGet200ResponseMonACacheBytes
	isSet bool
}

func (v NullableApiPerfCountersGet200ResponseMonACacheBytes) Get() *ApiPerfCountersGet200ResponseMonACacheBytes {
	return v.value
}

func (v *NullableApiPerfCountersGet200ResponseMonACacheBytes) Set(val *ApiPerfCountersGet200ResponseMonACacheBytes) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPerfCountersGet200ResponseMonACacheBytes) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPerfCountersGet200ResponseMonACacheBytes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPerfCountersGet200ResponseMonACacheBytes(val *ApiPerfCountersGet200ResponseMonACacheBytes) *NullableApiPerfCountersGet200ResponseMonACacheBytes {
	return &NullableApiPerfCountersGet200ResponseMonACacheBytes{value: val, isSet: true}
}

func (v NullableApiPerfCountersGet200ResponseMonACacheBytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPerfCountersGet200ResponseMonACacheBytes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


