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

// checks if the ApiMonitorGet200ResponseInQuorumInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMonitorGet200ResponseInQuorumInner{}

// ApiMonitorGet200ResponseInQuorumInner struct for ApiMonitorGet200ResponseInQuorumInner
type ApiMonitorGet200ResponseInQuorumInner struct {
	// 
	Addr string `json:"addr"`
	// 
	Name string `json:"name"`
	// 
	Priority int32 `json:"priority"`
	// 
	PublicAddr string `json:"public_addr"`
	PublicAddrs ApiMonitorGet200ResponseInQuorumInnerPublicAddrs `json:"public_addrs"`
	// 
	Rank int32 `json:"rank"`
	Stats ApiMonitorGet200ResponseInQuorumInnerStats `json:"stats"`
	// 
	Weight int32 `json:"weight"`
}

type _ApiMonitorGet200ResponseInQuorumInner ApiMonitorGet200ResponseInQuorumInner

// NewApiMonitorGet200ResponseInQuorumInner instantiates a new ApiMonitorGet200ResponseInQuorumInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMonitorGet200ResponseInQuorumInner(addr string, name string, priority int32, publicAddr string, publicAddrs ApiMonitorGet200ResponseInQuorumInnerPublicAddrs, rank int32, stats ApiMonitorGet200ResponseInQuorumInnerStats, weight int32) *ApiMonitorGet200ResponseInQuorumInner {
	this := ApiMonitorGet200ResponseInQuorumInner{}
	this.Addr = addr
	this.Name = name
	this.Priority = priority
	this.PublicAddr = publicAddr
	this.PublicAddrs = publicAddrs
	this.Rank = rank
	this.Stats = stats
	this.Weight = weight
	return &this
}

// NewApiMonitorGet200ResponseInQuorumInnerWithDefaults instantiates a new ApiMonitorGet200ResponseInQuorumInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMonitorGet200ResponseInQuorumInnerWithDefaults() *ApiMonitorGet200ResponseInQuorumInner {
	this := ApiMonitorGet200ResponseInQuorumInner{}
	return &this
}

// GetAddr returns the Addr field value
func (o *ApiMonitorGet200ResponseInQuorumInner) GetAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Addr
}

// GetAddrOk returns a tuple with the Addr field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseInQuorumInner) GetAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addr, true
}

// SetAddr sets field value
func (o *ApiMonitorGet200ResponseInQuorumInner) SetAddr(v string) {
	o.Addr = v
}

// GetName returns the Name field value
func (o *ApiMonitorGet200ResponseInQuorumInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseInQuorumInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiMonitorGet200ResponseInQuorumInner) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value
func (o *ApiMonitorGet200ResponseInQuorumInner) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseInQuorumInner) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *ApiMonitorGet200ResponseInQuorumInner) SetPriority(v int32) {
	o.Priority = v
}

// GetPublicAddr returns the PublicAddr field value
func (o *ApiMonitorGet200ResponseInQuorumInner) GetPublicAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicAddr
}

// GetPublicAddrOk returns a tuple with the PublicAddr field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseInQuorumInner) GetPublicAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicAddr, true
}

// SetPublicAddr sets field value
func (o *ApiMonitorGet200ResponseInQuorumInner) SetPublicAddr(v string) {
	o.PublicAddr = v
}

// GetPublicAddrs returns the PublicAddrs field value
func (o *ApiMonitorGet200ResponseInQuorumInner) GetPublicAddrs() ApiMonitorGet200ResponseInQuorumInnerPublicAddrs {
	if o == nil {
		var ret ApiMonitorGet200ResponseInQuorumInnerPublicAddrs
		return ret
	}

	return o.PublicAddrs
}

// GetPublicAddrsOk returns a tuple with the PublicAddrs field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseInQuorumInner) GetPublicAddrsOk() (*ApiMonitorGet200ResponseInQuorumInnerPublicAddrs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicAddrs, true
}

// SetPublicAddrs sets field value
func (o *ApiMonitorGet200ResponseInQuorumInner) SetPublicAddrs(v ApiMonitorGet200ResponseInQuorumInnerPublicAddrs) {
	o.PublicAddrs = v
}

// GetRank returns the Rank field value
func (o *ApiMonitorGet200ResponseInQuorumInner) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseInQuorumInner) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *ApiMonitorGet200ResponseInQuorumInner) SetRank(v int32) {
	o.Rank = v
}

// GetStats returns the Stats field value
func (o *ApiMonitorGet200ResponseInQuorumInner) GetStats() ApiMonitorGet200ResponseInQuorumInnerStats {
	if o == nil {
		var ret ApiMonitorGet200ResponseInQuorumInnerStats
		return ret
	}

	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseInQuorumInner) GetStatsOk() (*ApiMonitorGet200ResponseInQuorumInnerStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value
func (o *ApiMonitorGet200ResponseInQuorumInner) SetStats(v ApiMonitorGet200ResponseInQuorumInnerStats) {
	o.Stats = v
}

// GetWeight returns the Weight field value
func (o *ApiMonitorGet200ResponseInQuorumInner) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ApiMonitorGet200ResponseInQuorumInner) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *ApiMonitorGet200ResponseInQuorumInner) SetWeight(v int32) {
	o.Weight = v
}

func (o ApiMonitorGet200ResponseInQuorumInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMonitorGet200ResponseInQuorumInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addr"] = o.Addr
	toSerialize["name"] = o.Name
	toSerialize["priority"] = o.Priority
	toSerialize["public_addr"] = o.PublicAddr
	toSerialize["public_addrs"] = o.PublicAddrs
	toSerialize["rank"] = o.Rank
	toSerialize["stats"] = o.Stats
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

func (o *ApiMonitorGet200ResponseInQuorumInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addr",
		"name",
		"priority",
		"public_addr",
		"public_addrs",
		"rank",
		"stats",
		"weight",
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

	varApiMonitorGet200ResponseInQuorumInner := _ApiMonitorGet200ResponseInQuorumInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiMonitorGet200ResponseInQuorumInner)

	if err != nil {
		return err
	}

	*o = ApiMonitorGet200ResponseInQuorumInner(varApiMonitorGet200ResponseInQuorumInner)

	return err
}

type NullableApiMonitorGet200ResponseInQuorumInner struct {
	value *ApiMonitorGet200ResponseInQuorumInner
	isSet bool
}

func (v NullableApiMonitorGet200ResponseInQuorumInner) Get() *ApiMonitorGet200ResponseInQuorumInner {
	return v.value
}

func (v *NullableApiMonitorGet200ResponseInQuorumInner) Set(val *ApiMonitorGet200ResponseInQuorumInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMonitorGet200ResponseInQuorumInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMonitorGet200ResponseInQuorumInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMonitorGet200ResponseInQuorumInner(val *ApiMonitorGet200ResponseInQuorumInner) *NullableApiMonitorGet200ResponseInQuorumInner {
	return &NullableApiMonitorGet200ResponseInQuorumInner{value: val, isSet: true}
}

func (v NullableApiMonitorGet200ResponseInQuorumInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMonitorGet200ResponseInQuorumInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


