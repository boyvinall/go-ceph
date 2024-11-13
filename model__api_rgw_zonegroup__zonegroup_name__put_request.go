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

// checks if the ApiRgwZonegroupZonegroupNamePutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRgwZonegroupZonegroupNamePutRequest{}

// ApiRgwZonegroupZonegroupNamePutRequest struct for ApiRgwZonegroupZonegroupNamePutRequest
type ApiRgwZonegroupZonegroupNamePutRequest struct {
	AddZones *string `json:"add_zones,omitempty"`
	Default *string `json:"default,omitempty"`
	Master *string `json:"master,omitempty"`
	NewZonegroupName string `json:"new_zonegroup_name"`
	PlacementTargets *string `json:"placement_targets,omitempty"`
	RealmName string `json:"realm_name"`
	RemoveZones *string `json:"remove_zones,omitempty"`
	ZonegroupEndpoints *string `json:"zonegroup_endpoints,omitempty"`
}

type _ApiRgwZonegroupZonegroupNamePutRequest ApiRgwZonegroupZonegroupNamePutRequest

// NewApiRgwZonegroupZonegroupNamePutRequest instantiates a new ApiRgwZonegroupZonegroupNamePutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRgwZonegroupZonegroupNamePutRequest(newZonegroupName string, realmName string) *ApiRgwZonegroupZonegroupNamePutRequest {
	this := ApiRgwZonegroupZonegroupNamePutRequest{}
	var default_ string = ""
	this.Default = &default_
	var master string = ""
	this.Master = &master
	this.NewZonegroupName = newZonegroupName
	this.RealmName = realmName
	var zonegroupEndpoints string = ""
	this.ZonegroupEndpoints = &zonegroupEndpoints
	return &this
}

// NewApiRgwZonegroupZonegroupNamePutRequestWithDefaults instantiates a new ApiRgwZonegroupZonegroupNamePutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRgwZonegroupZonegroupNamePutRequestWithDefaults() *ApiRgwZonegroupZonegroupNamePutRequest {
	this := ApiRgwZonegroupZonegroupNamePutRequest{}
	var default_ string = ""
	this.Default = &default_
	var master string = ""
	this.Master = &master
	var zonegroupEndpoints string = ""
	this.ZonegroupEndpoints = &zonegroupEndpoints
	return &this
}

// GetAddZones returns the AddZones field value if set, zero value otherwise.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetAddZones() string {
	if o == nil || IsNil(o.AddZones) {
		var ret string
		return ret
	}
	return *o.AddZones
}

// GetAddZonesOk returns a tuple with the AddZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetAddZonesOk() (*string, bool) {
	if o == nil || IsNil(o.AddZones) {
		return nil, false
	}
	return o.AddZones, true
}

// HasAddZones returns a boolean if a field has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasAddZones() bool {
	if o != nil && !IsNil(o.AddZones) {
		return true
	}

	return false
}

// SetAddZones gets a reference to the given string and assigns it to the AddZones field.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetAddZones(v string) {
	o.AddZones = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetDefault() string {
	if o == nil || IsNil(o.Default) {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetDefault(v string) {
	o.Default = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetMaster() string {
	if o == nil || IsNil(o.Master) {
		var ret string
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetMasterOk() (*string, bool) {
	if o == nil || IsNil(o.Master) {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasMaster() bool {
	if o != nil && !IsNil(o.Master) {
		return true
	}

	return false
}

// SetMaster gets a reference to the given string and assigns it to the Master field.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetMaster(v string) {
	o.Master = &v
}

// GetNewZonegroupName returns the NewZonegroupName field value
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetNewZonegroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewZonegroupName
}

// GetNewZonegroupNameOk returns a tuple with the NewZonegroupName field value
// and a boolean to check if the value has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetNewZonegroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewZonegroupName, true
}

// SetNewZonegroupName sets field value
func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetNewZonegroupName(v string) {
	o.NewZonegroupName = v
}

// GetPlacementTargets returns the PlacementTargets field value if set, zero value otherwise.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetPlacementTargets() string {
	if o == nil || IsNil(o.PlacementTargets) {
		var ret string
		return ret
	}
	return *o.PlacementTargets
}

// GetPlacementTargetsOk returns a tuple with the PlacementTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetPlacementTargetsOk() (*string, bool) {
	if o == nil || IsNil(o.PlacementTargets) {
		return nil, false
	}
	return o.PlacementTargets, true
}

// HasPlacementTargets returns a boolean if a field has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasPlacementTargets() bool {
	if o != nil && !IsNil(o.PlacementTargets) {
		return true
	}

	return false
}

// SetPlacementTargets gets a reference to the given string and assigns it to the PlacementTargets field.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetPlacementTargets(v string) {
	o.PlacementTargets = &v
}

// GetRealmName returns the RealmName field value
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetRealmName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RealmName
}

// GetRealmNameOk returns a tuple with the RealmName field value
// and a boolean to check if the value has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetRealmNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RealmName, true
}

// SetRealmName sets field value
func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetRealmName(v string) {
	o.RealmName = v
}

// GetRemoveZones returns the RemoveZones field value if set, zero value otherwise.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetRemoveZones() string {
	if o == nil || IsNil(o.RemoveZones) {
		var ret string
		return ret
	}
	return *o.RemoveZones
}

// GetRemoveZonesOk returns a tuple with the RemoveZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetRemoveZonesOk() (*string, bool) {
	if o == nil || IsNil(o.RemoveZones) {
		return nil, false
	}
	return o.RemoveZones, true
}

// HasRemoveZones returns a boolean if a field has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasRemoveZones() bool {
	if o != nil && !IsNil(o.RemoveZones) {
		return true
	}

	return false
}

// SetRemoveZones gets a reference to the given string and assigns it to the RemoveZones field.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetRemoveZones(v string) {
	o.RemoveZones = &v
}

// GetZonegroupEndpoints returns the ZonegroupEndpoints field value if set, zero value otherwise.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetZonegroupEndpoints() string {
	if o == nil || IsNil(o.ZonegroupEndpoints) {
		var ret string
		return ret
	}
	return *o.ZonegroupEndpoints
}

// GetZonegroupEndpointsOk returns a tuple with the ZonegroupEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetZonegroupEndpointsOk() (*string, bool) {
	if o == nil || IsNil(o.ZonegroupEndpoints) {
		return nil, false
	}
	return o.ZonegroupEndpoints, true
}

// HasZonegroupEndpoints returns a boolean if a field has been set.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasZonegroupEndpoints() bool {
	if o != nil && !IsNil(o.ZonegroupEndpoints) {
		return true
	}

	return false
}

// SetZonegroupEndpoints gets a reference to the given string and assigns it to the ZonegroupEndpoints field.
func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetZonegroupEndpoints(v string) {
	o.ZonegroupEndpoints = &v
}

func (o ApiRgwZonegroupZonegroupNamePutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRgwZonegroupZonegroupNamePutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddZones) {
		toSerialize["add_zones"] = o.AddZones
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Master) {
		toSerialize["master"] = o.Master
	}
	toSerialize["new_zonegroup_name"] = o.NewZonegroupName
	if !IsNil(o.PlacementTargets) {
		toSerialize["placement_targets"] = o.PlacementTargets
	}
	toSerialize["realm_name"] = o.RealmName
	if !IsNil(o.RemoveZones) {
		toSerialize["remove_zones"] = o.RemoveZones
	}
	if !IsNil(o.ZonegroupEndpoints) {
		toSerialize["zonegroup_endpoints"] = o.ZonegroupEndpoints
	}
	return toSerialize, nil
}

func (o *ApiRgwZonegroupZonegroupNamePutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"new_zonegroup_name",
		"realm_name",
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

	varApiRgwZonegroupZonegroupNamePutRequest := _ApiRgwZonegroupZonegroupNamePutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiRgwZonegroupZonegroupNamePutRequest)

	if err != nil {
		return err
	}

	*o = ApiRgwZonegroupZonegroupNamePutRequest(varApiRgwZonegroupZonegroupNamePutRequest)

	return err
}

type NullableApiRgwZonegroupZonegroupNamePutRequest struct {
	value *ApiRgwZonegroupZonegroupNamePutRequest
	isSet bool
}

func (v NullableApiRgwZonegroupZonegroupNamePutRequest) Get() *ApiRgwZonegroupZonegroupNamePutRequest {
	return v.value
}

func (v *NullableApiRgwZonegroupZonegroupNamePutRequest) Set(val *ApiRgwZonegroupZonegroupNamePutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRgwZonegroupZonegroupNamePutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRgwZonegroupZonegroupNamePutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRgwZonegroupZonegroupNamePutRequest(val *ApiRgwZonegroupZonegroupNamePutRequest) *NullableApiRgwZonegroupZonegroupNamePutRequest {
	return &NullableApiRgwZonegroupZonegroupNamePutRequest{value: val, isSet: true}
}

func (v NullableApiRgwZonegroupZonegroupNamePutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRgwZonegroupZonegroupNamePutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

