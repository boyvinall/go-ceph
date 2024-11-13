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

// checks if the ApiRgwRealmImportRealmTokenPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRgwRealmImportRealmTokenPostRequest{}

// ApiRgwRealmImportRealmTokenPostRequest struct for ApiRgwRealmImportRealmTokenPostRequest
type ApiRgwRealmImportRealmTokenPostRequest struct {
	PlacementSpec *string `json:"placement_spec,omitempty"`
	Port string `json:"port"`
	RealmToken string `json:"realm_token"`
	ZoneName string `json:"zone_name"`
}

type _ApiRgwRealmImportRealmTokenPostRequest ApiRgwRealmImportRealmTokenPostRequest

// NewApiRgwRealmImportRealmTokenPostRequest instantiates a new ApiRgwRealmImportRealmTokenPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRgwRealmImportRealmTokenPostRequest(port string, realmToken string, zoneName string) *ApiRgwRealmImportRealmTokenPostRequest {
	this := ApiRgwRealmImportRealmTokenPostRequest{}
	this.Port = port
	this.RealmToken = realmToken
	this.ZoneName = zoneName
	return &this
}

// NewApiRgwRealmImportRealmTokenPostRequestWithDefaults instantiates a new ApiRgwRealmImportRealmTokenPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRgwRealmImportRealmTokenPostRequestWithDefaults() *ApiRgwRealmImportRealmTokenPostRequest {
	this := ApiRgwRealmImportRealmTokenPostRequest{}
	return &this
}

// GetPlacementSpec returns the PlacementSpec field value if set, zero value otherwise.
func (o *ApiRgwRealmImportRealmTokenPostRequest) GetPlacementSpec() string {
	if o == nil || IsNil(o.PlacementSpec) {
		var ret string
		return ret
	}
	return *o.PlacementSpec
}

// GetPlacementSpecOk returns a tuple with the PlacementSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwRealmImportRealmTokenPostRequest) GetPlacementSpecOk() (*string, bool) {
	if o == nil || IsNil(o.PlacementSpec) {
		return nil, false
	}
	return o.PlacementSpec, true
}

// HasPlacementSpec returns a boolean if a field has been set.
func (o *ApiRgwRealmImportRealmTokenPostRequest) HasPlacementSpec() bool {
	if o != nil && !IsNil(o.PlacementSpec) {
		return true
	}

	return false
}

// SetPlacementSpec gets a reference to the given string and assigns it to the PlacementSpec field.
func (o *ApiRgwRealmImportRealmTokenPostRequest) SetPlacementSpec(v string) {
	o.PlacementSpec = &v
}

// GetPort returns the Port field value
func (o *ApiRgwRealmImportRealmTokenPostRequest) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ApiRgwRealmImportRealmTokenPostRequest) GetPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ApiRgwRealmImportRealmTokenPostRequest) SetPort(v string) {
	o.Port = v
}

// GetRealmToken returns the RealmToken field value
func (o *ApiRgwRealmImportRealmTokenPostRequest) GetRealmToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RealmToken
}

// GetRealmTokenOk returns a tuple with the RealmToken field value
// and a boolean to check if the value has been set.
func (o *ApiRgwRealmImportRealmTokenPostRequest) GetRealmTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RealmToken, true
}

// SetRealmToken sets field value
func (o *ApiRgwRealmImportRealmTokenPostRequest) SetRealmToken(v string) {
	o.RealmToken = v
}

// GetZoneName returns the ZoneName field value
func (o *ApiRgwRealmImportRealmTokenPostRequest) GetZoneName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value
// and a boolean to check if the value has been set.
func (o *ApiRgwRealmImportRealmTokenPostRequest) GetZoneNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneName, true
}

// SetZoneName sets field value
func (o *ApiRgwRealmImportRealmTokenPostRequest) SetZoneName(v string) {
	o.ZoneName = v
}

func (o ApiRgwRealmImportRealmTokenPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRgwRealmImportRealmTokenPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlacementSpec) {
		toSerialize["placement_spec"] = o.PlacementSpec
	}
	toSerialize["port"] = o.Port
	toSerialize["realm_token"] = o.RealmToken
	toSerialize["zone_name"] = o.ZoneName
	return toSerialize, nil
}

func (o *ApiRgwRealmImportRealmTokenPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"port",
		"realm_token",
		"zone_name",
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

	varApiRgwRealmImportRealmTokenPostRequest := _ApiRgwRealmImportRealmTokenPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiRgwRealmImportRealmTokenPostRequest)

	if err != nil {
		return err
	}

	*o = ApiRgwRealmImportRealmTokenPostRequest(varApiRgwRealmImportRealmTokenPostRequest)

	return err
}

type NullableApiRgwRealmImportRealmTokenPostRequest struct {
	value *ApiRgwRealmImportRealmTokenPostRequest
	isSet bool
}

func (v NullableApiRgwRealmImportRealmTokenPostRequest) Get() *ApiRgwRealmImportRealmTokenPostRequest {
	return v.value
}

func (v *NullableApiRgwRealmImportRealmTokenPostRequest) Set(val *ApiRgwRealmImportRealmTokenPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRgwRealmImportRealmTokenPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRgwRealmImportRealmTokenPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRgwRealmImportRealmTokenPostRequest(val *ApiRgwRealmImportRealmTokenPostRequest) *NullableApiRgwRealmImportRealmTokenPostRequest {
	return &NullableApiRgwRealmImportRealmTokenPostRequest{value: val, isSet: true}
}

func (v NullableApiRgwRealmImportRealmTokenPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRgwRealmImportRealmTokenPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


