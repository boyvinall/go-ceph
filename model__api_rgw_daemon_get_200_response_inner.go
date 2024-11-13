/*
Ceph REST API

This is the official Ceph REST API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ceph

import (
	"encoding/json"
)

// checks if the ApiRgwDaemonGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRgwDaemonGet200ResponseInner{}

// ApiRgwDaemonGet200ResponseInner struct for ApiRgwDaemonGet200ResponseInner
type ApiRgwDaemonGet200ResponseInner struct {
	// Daemon ID
	Id *string `json:"id,omitempty"`
	// Port
	Port *int32 `json:"port,omitempty"`
	// 
	ServerHostname *string `json:"server_hostname,omitempty"`
	// Ceph Version
	Version *string `json:"version,omitempty"`
	// Zone
	ZoneName *string `json:"zone_name,omitempty"`
	// Zone Group
	ZonegroupName *string `json:"zonegroup_name,omitempty"`
}

// NewApiRgwDaemonGet200ResponseInner instantiates a new ApiRgwDaemonGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRgwDaemonGet200ResponseInner() *ApiRgwDaemonGet200ResponseInner {
	this := ApiRgwDaemonGet200ResponseInner{}
	return &this
}

// NewApiRgwDaemonGet200ResponseInnerWithDefaults instantiates a new ApiRgwDaemonGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRgwDaemonGet200ResponseInnerWithDefaults() *ApiRgwDaemonGet200ResponseInner {
	this := ApiRgwDaemonGet200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiRgwDaemonGet200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwDaemonGet200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiRgwDaemonGet200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiRgwDaemonGet200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ApiRgwDaemonGet200ResponseInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwDaemonGet200ResponseInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ApiRgwDaemonGet200ResponseInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ApiRgwDaemonGet200ResponseInner) SetPort(v int32) {
	o.Port = &v
}

// GetServerHostname returns the ServerHostname field value if set, zero value otherwise.
func (o *ApiRgwDaemonGet200ResponseInner) GetServerHostname() string {
	if o == nil || IsNil(o.ServerHostname) {
		var ret string
		return ret
	}
	return *o.ServerHostname
}

// GetServerHostnameOk returns a tuple with the ServerHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwDaemonGet200ResponseInner) GetServerHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerHostname) {
		return nil, false
	}
	return o.ServerHostname, true
}

// HasServerHostname returns a boolean if a field has been set.
func (o *ApiRgwDaemonGet200ResponseInner) HasServerHostname() bool {
	if o != nil && !IsNil(o.ServerHostname) {
		return true
	}

	return false
}

// SetServerHostname gets a reference to the given string and assigns it to the ServerHostname field.
func (o *ApiRgwDaemonGet200ResponseInner) SetServerHostname(v string) {
	o.ServerHostname = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApiRgwDaemonGet200ResponseInner) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwDaemonGet200ResponseInner) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApiRgwDaemonGet200ResponseInner) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApiRgwDaemonGet200ResponseInner) SetVersion(v string) {
	o.Version = &v
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise.
func (o *ApiRgwDaemonGet200ResponseInner) GetZoneName() string {
	if o == nil || IsNil(o.ZoneName) {
		var ret string
		return ret
	}
	return *o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwDaemonGet200ResponseInner) GetZoneNameOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneName) {
		return nil, false
	}
	return o.ZoneName, true
}

// HasZoneName returns a boolean if a field has been set.
func (o *ApiRgwDaemonGet200ResponseInner) HasZoneName() bool {
	if o != nil && !IsNil(o.ZoneName) {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given string and assigns it to the ZoneName field.
func (o *ApiRgwDaemonGet200ResponseInner) SetZoneName(v string) {
	o.ZoneName = &v
}

// GetZonegroupName returns the ZonegroupName field value if set, zero value otherwise.
func (o *ApiRgwDaemonGet200ResponseInner) GetZonegroupName() string {
	if o == nil || IsNil(o.ZonegroupName) {
		var ret string
		return ret
	}
	return *o.ZonegroupName
}

// GetZonegroupNameOk returns a tuple with the ZonegroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwDaemonGet200ResponseInner) GetZonegroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.ZonegroupName) {
		return nil, false
	}
	return o.ZonegroupName, true
}

// HasZonegroupName returns a boolean if a field has been set.
func (o *ApiRgwDaemonGet200ResponseInner) HasZonegroupName() bool {
	if o != nil && !IsNil(o.ZonegroupName) {
		return true
	}

	return false
}

// SetZonegroupName gets a reference to the given string and assigns it to the ZonegroupName field.
func (o *ApiRgwDaemonGet200ResponseInner) SetZonegroupName(v string) {
	o.ZonegroupName = &v
}

func (o ApiRgwDaemonGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRgwDaemonGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.ServerHostname) {
		toSerialize["server_hostname"] = o.ServerHostname
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.ZoneName) {
		toSerialize["zone_name"] = o.ZoneName
	}
	if !IsNil(o.ZonegroupName) {
		toSerialize["zonegroup_name"] = o.ZonegroupName
	}
	return toSerialize, nil
}

type NullableApiRgwDaemonGet200ResponseInner struct {
	value *ApiRgwDaemonGet200ResponseInner
	isSet bool
}

func (v NullableApiRgwDaemonGet200ResponseInner) Get() *ApiRgwDaemonGet200ResponseInner {
	return v.value
}

func (v *NullableApiRgwDaemonGet200ResponseInner) Set(val *ApiRgwDaemonGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRgwDaemonGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRgwDaemonGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRgwDaemonGet200ResponseInner(val *ApiRgwDaemonGet200ResponseInner) *NullableApiRgwDaemonGet200ResponseInner {
	return &NullableApiRgwDaemonGet200ResponseInner{value: val, isSet: true}
}

func (v NullableApiRgwDaemonGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRgwDaemonGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

