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

// checks if the ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner{}

// ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner struct for ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner
type ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner struct {
	// 
	BlockUuid string `json:"block_uuid"`
	// 
	ClusterFsid string `json:"cluster_fsid"`
	// 
	ClusterName string `json:"cluster_name"`
	// 
	Name string `json:"name"`
	// 
	OsdFsid string `json:"osd_fsid"`
	// 
	OsdId string `json:"osd_id"`
	// 
	OsdspecAffinity string `json:"osdspec_affinity"`
	// 
	Type string `json:"type"`
}

type _ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner

// NewApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner instantiates a new ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner(blockUuid string, clusterFsid string, clusterName string, name string, osdFsid string, osdId string, osdspecAffinity string, type_ string) *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner {
	this := ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner{}
	this.BlockUuid = blockUuid
	this.ClusterFsid = clusterFsid
	this.ClusterName = clusterName
	this.Name = name
	this.OsdFsid = osdFsid
	this.OsdId = osdId
	this.OsdspecAffinity = osdspecAffinity
	this.Type = type_
	return &this
}

// NewApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInnerWithDefaults instantiates a new ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInnerWithDefaults() *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner {
	this := ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner{}
	return &this
}

// GetBlockUuid returns the BlockUuid field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetBlockUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockUuid
}

// GetBlockUuidOk returns a tuple with the BlockUuid field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetBlockUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockUuid, true
}

// SetBlockUuid sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) SetBlockUuid(v string) {
	o.BlockUuid = v
}

// GetClusterFsid returns the ClusterFsid field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetClusterFsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterFsid
}

// GetClusterFsidOk returns a tuple with the ClusterFsid field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetClusterFsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterFsid, true
}

// SetClusterFsid sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) SetClusterFsid(v string) {
	o.ClusterFsid = v
}

// GetClusterName returns the ClusterName field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) SetClusterName(v string) {
	o.ClusterName = v
}

// GetName returns the Name field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) SetName(v string) {
	o.Name = v
}

// GetOsdFsid returns the OsdFsid field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetOsdFsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsdFsid
}

// GetOsdFsidOk returns a tuple with the OsdFsid field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetOsdFsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsdFsid, true
}

// SetOsdFsid sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) SetOsdFsid(v string) {
	o.OsdFsid = v
}

// GetOsdId returns the OsdId field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetOsdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsdId
}

// GetOsdIdOk returns a tuple with the OsdId field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetOsdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsdId, true
}

// SetOsdId sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) SetOsdId(v string) {
	o.OsdId = v
}

// GetOsdspecAffinity returns the OsdspecAffinity field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetOsdspecAffinity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsdspecAffinity
}

// GetOsdspecAffinityOk returns a tuple with the OsdspecAffinity field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetOsdspecAffinityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsdspecAffinity, true
}

// SetOsdspecAffinity sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) SetOsdspecAffinity(v string) {
	o.OsdspecAffinity = v
}

// GetType returns the Type field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) SetType(v string) {
	o.Type = v
}

func (o ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block_uuid"] = o.BlockUuid
	toSerialize["cluster_fsid"] = o.ClusterFsid
	toSerialize["cluster_name"] = o.ClusterName
	toSerialize["name"] = o.Name
	toSerialize["osd_fsid"] = o.OsdFsid
	toSerialize["osd_id"] = o.OsdId
	toSerialize["osdspec_affinity"] = o.OsdspecAffinity
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"block_uuid",
		"cluster_fsid",
		"cluster_name",
		"name",
		"osd_fsid",
		"osd_id",
		"osdspec_affinity",
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

	varApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner := _ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner)

	if err != nil {
		return err
	}

	*o = ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner(varApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner)

	return err
}

type NullableApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner struct {
	value *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner
	isSet bool
}

func (v NullableApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) Get() *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner {
	return v.value
}

func (v *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) Set(val *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner(val *ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner {
	return &NullableApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner{value: val, isSet: true}
}

func (v NullableApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


