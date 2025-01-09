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

// checks if the ApiClusterUserPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiClusterUserPostRequest{}

// ApiClusterUserPostRequest struct for ApiClusterUserPostRequest
type ApiClusterUserPostRequest struct {
	// List of capabilities to add to user_entity
	Capabilities []ApiClusterUserPostRequestCapabilitiesInner `json:"capabilities,omitempty"`
	ImportData *string `json:"import_data,omitempty"`
	// Entity to add
	UserEntity *string `json:"user_entity,omitempty"`
}

// NewApiClusterUserPostRequest instantiates a new ApiClusterUserPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiClusterUserPostRequest() *ApiClusterUserPostRequest {
	this := ApiClusterUserPostRequest{}
	var importData string = ""
	this.ImportData = &importData
	var userEntity string = ""
	this.UserEntity = &userEntity
	return &this
}

// NewApiClusterUserPostRequestWithDefaults instantiates a new ApiClusterUserPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiClusterUserPostRequestWithDefaults() *ApiClusterUserPostRequest {
	this := ApiClusterUserPostRequest{}
	var importData string = ""
	this.ImportData = &importData
	var userEntity string = ""
	this.UserEntity = &userEntity
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ApiClusterUserPostRequest) GetCapabilities() []ApiClusterUserPostRequestCapabilitiesInner {
	if o == nil || IsNil(o.Capabilities) {
		var ret []ApiClusterUserPostRequestCapabilitiesInner
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterUserPostRequest) GetCapabilitiesOk() ([]ApiClusterUserPostRequestCapabilitiesInner, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ApiClusterUserPostRequest) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []ApiClusterUserPostRequestCapabilitiesInner and assigns it to the Capabilities field.
func (o *ApiClusterUserPostRequest) SetCapabilities(v []ApiClusterUserPostRequestCapabilitiesInner) {
	o.Capabilities = v
}

// GetImportData returns the ImportData field value if set, zero value otherwise.
func (o *ApiClusterUserPostRequest) GetImportData() string {
	if o == nil || IsNil(o.ImportData) {
		var ret string
		return ret
	}
	return *o.ImportData
}

// GetImportDataOk returns a tuple with the ImportData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterUserPostRequest) GetImportDataOk() (*string, bool) {
	if o == nil || IsNil(o.ImportData) {
		return nil, false
	}
	return o.ImportData, true
}

// HasImportData returns a boolean if a field has been set.
func (o *ApiClusterUserPostRequest) HasImportData() bool {
	if o != nil && !IsNil(o.ImportData) {
		return true
	}

	return false
}

// SetImportData gets a reference to the given string and assigns it to the ImportData field.
func (o *ApiClusterUserPostRequest) SetImportData(v string) {
	o.ImportData = &v
}

// GetUserEntity returns the UserEntity field value if set, zero value otherwise.
func (o *ApiClusterUserPostRequest) GetUserEntity() string {
	if o == nil || IsNil(o.UserEntity) {
		var ret string
		return ret
	}
	return *o.UserEntity
}

// GetUserEntityOk returns a tuple with the UserEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterUserPostRequest) GetUserEntityOk() (*string, bool) {
	if o == nil || IsNil(o.UserEntity) {
		return nil, false
	}
	return o.UserEntity, true
}

// HasUserEntity returns a boolean if a field has been set.
func (o *ApiClusterUserPostRequest) HasUserEntity() bool {
	if o != nil && !IsNil(o.UserEntity) {
		return true
	}

	return false
}

// SetUserEntity gets a reference to the given string and assigns it to the UserEntity field.
func (o *ApiClusterUserPostRequest) SetUserEntity(v string) {
	o.UserEntity = &v
}

func (o ApiClusterUserPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiClusterUserPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.ImportData) {
		toSerialize["import_data"] = o.ImportData
	}
	if !IsNil(o.UserEntity) {
		toSerialize["user_entity"] = o.UserEntity
	}
	return toSerialize, nil
}

type NullableApiClusterUserPostRequest struct {
	value *ApiClusterUserPostRequest
	isSet bool
}

func (v NullableApiClusterUserPostRequest) Get() *ApiClusterUserPostRequest {
	return v.value
}

func (v *NullableApiClusterUserPostRequest) Set(val *ApiClusterUserPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiClusterUserPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiClusterUserPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiClusterUserPostRequest(val *ApiClusterUserPostRequest) *NullableApiClusterUserPostRequest {
	return &NullableApiClusterUserPostRequest{value: val, isSet: true}
}

func (v NullableApiClusterUserPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiClusterUserPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

