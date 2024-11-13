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

// checks if the ApiHealthMinimalGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHealthMinimalGet200Response{}

// ApiHealthMinimalGet200Response struct for ApiHealthMinimalGet200Response
type ApiHealthMinimalGet200Response struct {
	ClientPerf ApiHealthMinimalGet200ResponseClientPerf `json:"client_perf"`
	Df ApiHealthMinimalGet200ResponseDf `json:"df"`
	FsMap ApiHealthMinimalGet200ResponseFsMap `json:"fs_map"`
	Health ApiHealthMinimalGet200ResponseHealth `json:"health"`
	// 
	Hosts int32 `json:"hosts"`
	IscsiDaemons ApiHealthMinimalGet200ResponseIscsiDaemons `json:"iscsi_daemons"`
	MgrMap ApiHealthMinimalGet200ResponseMgrMap `json:"mgr_map"`
	MonStatus ApiHealthMinimalGet200ResponseMonStatus `json:"mon_status"`
	OsdMap ApiHealthMinimalGet200ResponseOsdMap `json:"osd_map"`
	PgInfo ApiHealthMinimalGet200ResponsePgInfo `json:"pg_info"`
	// 
	Pools string `json:"pools"`
	// 
	Rgw int32 `json:"rgw"`
	// 
	ScrubStatus string `json:"scrub_status"`
}

type _ApiHealthMinimalGet200Response ApiHealthMinimalGet200Response

// NewApiHealthMinimalGet200Response instantiates a new ApiHealthMinimalGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHealthMinimalGet200Response(clientPerf ApiHealthMinimalGet200ResponseClientPerf, df ApiHealthMinimalGet200ResponseDf, fsMap ApiHealthMinimalGet200ResponseFsMap, health ApiHealthMinimalGet200ResponseHealth, hosts int32, iscsiDaemons ApiHealthMinimalGet200ResponseIscsiDaemons, mgrMap ApiHealthMinimalGet200ResponseMgrMap, monStatus ApiHealthMinimalGet200ResponseMonStatus, osdMap ApiHealthMinimalGet200ResponseOsdMap, pgInfo ApiHealthMinimalGet200ResponsePgInfo, pools string, rgw int32, scrubStatus string) *ApiHealthMinimalGet200Response {
	this := ApiHealthMinimalGet200Response{}
	this.ClientPerf = clientPerf
	this.Df = df
	this.FsMap = fsMap
	this.Health = health
	this.Hosts = hosts
	this.IscsiDaemons = iscsiDaemons
	this.MgrMap = mgrMap
	this.MonStatus = monStatus
	this.OsdMap = osdMap
	this.PgInfo = pgInfo
	this.Pools = pools
	this.Rgw = rgw
	this.ScrubStatus = scrubStatus
	return &this
}

// NewApiHealthMinimalGet200ResponseWithDefaults instantiates a new ApiHealthMinimalGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHealthMinimalGet200ResponseWithDefaults() *ApiHealthMinimalGet200Response {
	this := ApiHealthMinimalGet200Response{}
	return &this
}

// GetClientPerf returns the ClientPerf field value
func (o *ApiHealthMinimalGet200Response) GetClientPerf() ApiHealthMinimalGet200ResponseClientPerf {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponseClientPerf
		return ret
	}

	return o.ClientPerf
}

// GetClientPerfOk returns a tuple with the ClientPerf field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetClientPerfOk() (*ApiHealthMinimalGet200ResponseClientPerf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientPerf, true
}

// SetClientPerf sets field value
func (o *ApiHealthMinimalGet200Response) SetClientPerf(v ApiHealthMinimalGet200ResponseClientPerf) {
	o.ClientPerf = v
}

// GetDf returns the Df field value
func (o *ApiHealthMinimalGet200Response) GetDf() ApiHealthMinimalGet200ResponseDf {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponseDf
		return ret
	}

	return o.Df
}

// GetDfOk returns a tuple with the Df field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetDfOk() (*ApiHealthMinimalGet200ResponseDf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Df, true
}

// SetDf sets field value
func (o *ApiHealthMinimalGet200Response) SetDf(v ApiHealthMinimalGet200ResponseDf) {
	o.Df = v
}

// GetFsMap returns the FsMap field value
func (o *ApiHealthMinimalGet200Response) GetFsMap() ApiHealthMinimalGet200ResponseFsMap {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponseFsMap
		return ret
	}

	return o.FsMap
}

// GetFsMapOk returns a tuple with the FsMap field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetFsMapOk() (*ApiHealthMinimalGet200ResponseFsMap, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FsMap, true
}

// SetFsMap sets field value
func (o *ApiHealthMinimalGet200Response) SetFsMap(v ApiHealthMinimalGet200ResponseFsMap) {
	o.FsMap = v
}

// GetHealth returns the Health field value
func (o *ApiHealthMinimalGet200Response) GetHealth() ApiHealthMinimalGet200ResponseHealth {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponseHealth
		return ret
	}

	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetHealthOk() (*ApiHealthMinimalGet200ResponseHealth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value
func (o *ApiHealthMinimalGet200Response) SetHealth(v ApiHealthMinimalGet200ResponseHealth) {
	o.Health = v
}

// GetHosts returns the Hosts field value
func (o *ApiHealthMinimalGet200Response) GetHosts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetHostsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// SetHosts sets field value
func (o *ApiHealthMinimalGet200Response) SetHosts(v int32) {
	o.Hosts = v
}

// GetIscsiDaemons returns the IscsiDaemons field value
func (o *ApiHealthMinimalGet200Response) GetIscsiDaemons() ApiHealthMinimalGet200ResponseIscsiDaemons {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponseIscsiDaemons
		return ret
	}

	return o.IscsiDaemons
}

// GetIscsiDaemonsOk returns a tuple with the IscsiDaemons field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetIscsiDaemonsOk() (*ApiHealthMinimalGet200ResponseIscsiDaemons, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IscsiDaemons, true
}

// SetIscsiDaemons sets field value
func (o *ApiHealthMinimalGet200Response) SetIscsiDaemons(v ApiHealthMinimalGet200ResponseIscsiDaemons) {
	o.IscsiDaemons = v
}

// GetMgrMap returns the MgrMap field value
func (o *ApiHealthMinimalGet200Response) GetMgrMap() ApiHealthMinimalGet200ResponseMgrMap {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponseMgrMap
		return ret
	}

	return o.MgrMap
}

// GetMgrMapOk returns a tuple with the MgrMap field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetMgrMapOk() (*ApiHealthMinimalGet200ResponseMgrMap, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MgrMap, true
}

// SetMgrMap sets field value
func (o *ApiHealthMinimalGet200Response) SetMgrMap(v ApiHealthMinimalGet200ResponseMgrMap) {
	o.MgrMap = v
}

// GetMonStatus returns the MonStatus field value
func (o *ApiHealthMinimalGet200Response) GetMonStatus() ApiHealthMinimalGet200ResponseMonStatus {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponseMonStatus
		return ret
	}

	return o.MonStatus
}

// GetMonStatusOk returns a tuple with the MonStatus field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetMonStatusOk() (*ApiHealthMinimalGet200ResponseMonStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonStatus, true
}

// SetMonStatus sets field value
func (o *ApiHealthMinimalGet200Response) SetMonStatus(v ApiHealthMinimalGet200ResponseMonStatus) {
	o.MonStatus = v
}

// GetOsdMap returns the OsdMap field value
func (o *ApiHealthMinimalGet200Response) GetOsdMap() ApiHealthMinimalGet200ResponseOsdMap {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponseOsdMap
		return ret
	}

	return o.OsdMap
}

// GetOsdMapOk returns a tuple with the OsdMap field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetOsdMapOk() (*ApiHealthMinimalGet200ResponseOsdMap, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsdMap, true
}

// SetOsdMap sets field value
func (o *ApiHealthMinimalGet200Response) SetOsdMap(v ApiHealthMinimalGet200ResponseOsdMap) {
	o.OsdMap = v
}

// GetPgInfo returns the PgInfo field value
func (o *ApiHealthMinimalGet200Response) GetPgInfo() ApiHealthMinimalGet200ResponsePgInfo {
	if o == nil {
		var ret ApiHealthMinimalGet200ResponsePgInfo
		return ret
	}

	return o.PgInfo
}

// GetPgInfoOk returns a tuple with the PgInfo field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetPgInfoOk() (*ApiHealthMinimalGet200ResponsePgInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PgInfo, true
}

// SetPgInfo sets field value
func (o *ApiHealthMinimalGet200Response) SetPgInfo(v ApiHealthMinimalGet200ResponsePgInfo) {
	o.PgInfo = v
}

// GetPools returns the Pools field value
func (o *ApiHealthMinimalGet200Response) GetPools() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetPoolsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pools, true
}

// SetPools sets field value
func (o *ApiHealthMinimalGet200Response) SetPools(v string) {
	o.Pools = v
}

// GetRgw returns the Rgw field value
func (o *ApiHealthMinimalGet200Response) GetRgw() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rgw
}

// GetRgwOk returns a tuple with the Rgw field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetRgwOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rgw, true
}

// SetRgw sets field value
func (o *ApiHealthMinimalGet200Response) SetRgw(v int32) {
	o.Rgw = v
}

// GetScrubStatus returns the ScrubStatus field value
func (o *ApiHealthMinimalGet200Response) GetScrubStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScrubStatus
}

// GetScrubStatusOk returns a tuple with the ScrubStatus field value
// and a boolean to check if the value has been set.
func (o *ApiHealthMinimalGet200Response) GetScrubStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScrubStatus, true
}

// SetScrubStatus sets field value
func (o *ApiHealthMinimalGet200Response) SetScrubStatus(v string) {
	o.ScrubStatus = v
}

func (o ApiHealthMinimalGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHealthMinimalGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client_perf"] = o.ClientPerf
	toSerialize["df"] = o.Df
	toSerialize["fs_map"] = o.FsMap
	toSerialize["health"] = o.Health
	toSerialize["hosts"] = o.Hosts
	toSerialize["iscsi_daemons"] = o.IscsiDaemons
	toSerialize["mgr_map"] = o.MgrMap
	toSerialize["mon_status"] = o.MonStatus
	toSerialize["osd_map"] = o.OsdMap
	toSerialize["pg_info"] = o.PgInfo
	toSerialize["pools"] = o.Pools
	toSerialize["rgw"] = o.Rgw
	toSerialize["scrub_status"] = o.ScrubStatus
	return toSerialize, nil
}

func (o *ApiHealthMinimalGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_perf",
		"df",
		"fs_map",
		"health",
		"hosts",
		"iscsi_daemons",
		"mgr_map",
		"mon_status",
		"osd_map",
		"pg_info",
		"pools",
		"rgw",
		"scrub_status",
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

	varApiHealthMinimalGet200Response := _ApiHealthMinimalGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiHealthMinimalGet200Response)

	if err != nil {
		return err
	}

	*o = ApiHealthMinimalGet200Response(varApiHealthMinimalGet200Response)

	return err
}

type NullableApiHealthMinimalGet200Response struct {
	value *ApiHealthMinimalGet200Response
	isSet bool
}

func (v NullableApiHealthMinimalGet200Response) Get() *ApiHealthMinimalGet200Response {
	return v.value
}

func (v *NullableApiHealthMinimalGet200Response) Set(val *ApiHealthMinimalGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHealthMinimalGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHealthMinimalGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHealthMinimalGet200Response(val *ApiHealthMinimalGet200Response) *NullableApiHealthMinimalGet200Response {
	return &NullableApiHealthMinimalGet200Response{value: val, isSet: true}
}

func (v NullableApiHealthMinimalGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHealthMinimalGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


