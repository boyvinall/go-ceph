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

// checks if the ApiClusterUserPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiClusterUserPutRequest{}

// ApiClusterUserPutRequest struct for ApiClusterUserPutRequest
type ApiClusterUserPutRequest struct {
	// List of updated capabilities to user_entity
	Capabilities []ApiClusterUserPutRequestCapabilitiesInner `json:"capabilities,omitempty"`
	// Entity to edit
	UserEntity *string `json:"user_entity,omitempty"`
}

// NewApiClusterUserPutRequest instantiates a new ApiClusterUserPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiClusterUserPutRequest() *ApiClusterUserPutRequest {
	this := ApiClusterUserPutRequest{}
	var userEntity string = ""
	this.UserEntity = &userEntity
	return &this
}

// NewApiClusterUserPutRequestWithDefaults instantiates a new ApiClusterUserPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiClusterUserPutRequestWithDefaults() *ApiClusterUserPutRequest {
	this := ApiClusterUserPutRequest{}
	var userEntity string = ""
	this.UserEntity = &userEntity
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ApiClusterUserPutRequest) GetCapabilities() []ApiClusterUserPutRequestCapabilitiesInner {
	if o == nil || IsNil(o.Capabilities) {
		var ret []ApiClusterUserPutRequestCapabilitiesInner
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterUserPutRequest) GetCapabilitiesOk() ([]ApiClusterUserPutRequestCapabilitiesInner, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ApiClusterUserPutRequest) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []ApiClusterUserPutRequestCapabilitiesInner and assigns it to the Capabilities field.
func (o *ApiClusterUserPutRequest) SetCapabilities(v []ApiClusterUserPutRequestCapabilitiesInner) {
	o.Capabilities = v
}

// GetUserEntity returns the UserEntity field value if set, zero value otherwise.
func (o *ApiClusterUserPutRequest) GetUserEntity() string {
	if o == nil || IsNil(o.UserEntity) {
		var ret string
		return ret
	}
	return *o.UserEntity
}

// GetUserEntityOk returns a tuple with the UserEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterUserPutRequest) GetUserEntityOk() (*string, bool) {
	if o == nil || IsNil(o.UserEntity) {
		return nil, false
	}
	return o.UserEntity, true
}

// HasUserEntity returns a boolean if a field has been set.
func (o *ApiClusterUserPutRequest) HasUserEntity() bool {
	if o != nil && !IsNil(o.UserEntity) {
		return true
	}

	return false
}

// SetUserEntity gets a reference to the given string and assigns it to the UserEntity field.
func (o *ApiClusterUserPutRequest) SetUserEntity(v string) {
	o.UserEntity = &v
}

func (o ApiClusterUserPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiClusterUserPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.UserEntity) {
		toSerialize["user_entity"] = o.UserEntity
	}
	return toSerialize, nil
}

type NullableApiClusterUserPutRequest struct {
	value *ApiClusterUserPutRequest
	isSet bool
}

func (v NullableApiClusterUserPutRequest) Get() *ApiClusterUserPutRequest {
	return v.value
}

func (v *NullableApiClusterUserPutRequest) Set(val *ApiClusterUserPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiClusterUserPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiClusterUserPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiClusterUserPutRequest(val *ApiClusterUserPutRequest) *NullableApiClusterUserPutRequest {
	return &NullableApiClusterUserPutRequest{value: val, isSet: true}
}

func (v NullableApiClusterUserPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiClusterUserPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

