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

// checks if the ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest{}

// ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest struct for ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest
type ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest struct {
	ClientId *string `json:"client_id,omitempty"`
	ClusterName *string `json:"cluster_name,omitempty"`
	Key *string `json:"key,omitempty"`
	MonHost *string `json:"mon_host,omitempty"`
}

// NewApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest instantiates a new ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest() *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest {
	this := ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest{}
	return &this
}

// NewApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequestWithDefaults instantiates a new ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequestWithDefaults() *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest {
	this := ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) SetKey(v string) {
	o.Key = &v
}

// GetMonHost returns the MonHost field value if set, zero value otherwise.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) GetMonHost() string {
	if o == nil || IsNil(o.MonHost) {
		var ret string
		return ret
	}
	return *o.MonHost
}

// GetMonHostOk returns a tuple with the MonHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) GetMonHostOk() (*string, bool) {
	if o == nil || IsNil(o.MonHost) {
		return nil, false
	}
	return o.MonHost, true
}

// HasMonHost returns a boolean if a field has been set.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) HasMonHost() bool {
	if o != nil && !IsNil(o.MonHost) {
		return true
	}

	return false
}

// SetMonHost gets a reference to the given string and assigns it to the MonHost field.
func (o *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) SetMonHost(v string) {
	o.MonHost = &v
}

func (o ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClusterName) {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.MonHost) {
		toSerialize["mon_host"] = o.MonHost
	}
	return toSerialize, nil
}

type NullableApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest struct {
	value *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest
	isSet bool
}

func (v NullableApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) Get() *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest {
	return v.value
}

func (v *NullableApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) Set(val *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest(val *ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) *NullableApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest {
	return &NullableApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest{value: val, isSet: true}
}

func (v NullableApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


