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

// checks if the ApiRgwUserUidKeyPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRgwUserUidKeyPostRequest{}

// ApiRgwUserUidKeyPostRequest struct for ApiRgwUserUidKeyPostRequest
type ApiRgwUserUidKeyPostRequest struct {
	AccessKey *string `json:"access_key,omitempty"`
	DaemonName *string `json:"daemon_name,omitempty"`
	GenerateKey *string `json:"generate_key,omitempty"`
	KeyType *string `json:"key_type,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
	Subuser *string `json:"subuser,omitempty"`
}

// NewApiRgwUserUidKeyPostRequest instantiates a new ApiRgwUserUidKeyPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRgwUserUidKeyPostRequest() *ApiRgwUserUidKeyPostRequest {
	this := ApiRgwUserUidKeyPostRequest{}
	var generateKey string = "true"
	this.GenerateKey = &generateKey
	var keyType string = "s3"
	this.KeyType = &keyType
	return &this
}

// NewApiRgwUserUidKeyPostRequestWithDefaults instantiates a new ApiRgwUserUidKeyPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRgwUserUidKeyPostRequestWithDefaults() *ApiRgwUserUidKeyPostRequest {
	this := ApiRgwUserUidKeyPostRequest{}
	var generateKey string = "true"
	this.GenerateKey = &generateKey
	var keyType string = "s3"
	this.KeyType = &keyType
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *ApiRgwUserUidKeyPostRequest) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidKeyPostRequest) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *ApiRgwUserUidKeyPostRequest) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *ApiRgwUserUidKeyPostRequest) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetDaemonName returns the DaemonName field value if set, zero value otherwise.
func (o *ApiRgwUserUidKeyPostRequest) GetDaemonName() string {
	if o == nil || IsNil(o.DaemonName) {
		var ret string
		return ret
	}
	return *o.DaemonName
}

// GetDaemonNameOk returns a tuple with the DaemonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidKeyPostRequest) GetDaemonNameOk() (*string, bool) {
	if o == nil || IsNil(o.DaemonName) {
		return nil, false
	}
	return o.DaemonName, true
}

// HasDaemonName returns a boolean if a field has been set.
func (o *ApiRgwUserUidKeyPostRequest) HasDaemonName() bool {
	if o != nil && !IsNil(o.DaemonName) {
		return true
	}

	return false
}

// SetDaemonName gets a reference to the given string and assigns it to the DaemonName field.
func (o *ApiRgwUserUidKeyPostRequest) SetDaemonName(v string) {
	o.DaemonName = &v
}

// GetGenerateKey returns the GenerateKey field value if set, zero value otherwise.
func (o *ApiRgwUserUidKeyPostRequest) GetGenerateKey() string {
	if o == nil || IsNil(o.GenerateKey) {
		var ret string
		return ret
	}
	return *o.GenerateKey
}

// GetGenerateKeyOk returns a tuple with the GenerateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidKeyPostRequest) GetGenerateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GenerateKey) {
		return nil, false
	}
	return o.GenerateKey, true
}

// HasGenerateKey returns a boolean if a field has been set.
func (o *ApiRgwUserUidKeyPostRequest) HasGenerateKey() bool {
	if o != nil && !IsNil(o.GenerateKey) {
		return true
	}

	return false
}

// SetGenerateKey gets a reference to the given string and assigns it to the GenerateKey field.
func (o *ApiRgwUserUidKeyPostRequest) SetGenerateKey(v string) {
	o.GenerateKey = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *ApiRgwUserUidKeyPostRequest) GetKeyType() string {
	if o == nil || IsNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidKeyPostRequest) GetKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *ApiRgwUserUidKeyPostRequest) HasKeyType() bool {
	if o != nil && !IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *ApiRgwUserUidKeyPostRequest) SetKeyType(v string) {
	o.KeyType = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *ApiRgwUserUidKeyPostRequest) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidKeyPostRequest) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *ApiRgwUserUidKeyPostRequest) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *ApiRgwUserUidKeyPostRequest) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetSubuser returns the Subuser field value if set, zero value otherwise.
func (o *ApiRgwUserUidKeyPostRequest) GetSubuser() string {
	if o == nil || IsNil(o.Subuser) {
		var ret string
		return ret
	}
	return *o.Subuser
}

// GetSubuserOk returns a tuple with the Subuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidKeyPostRequest) GetSubuserOk() (*string, bool) {
	if o == nil || IsNil(o.Subuser) {
		return nil, false
	}
	return o.Subuser, true
}

// HasSubuser returns a boolean if a field has been set.
func (o *ApiRgwUserUidKeyPostRequest) HasSubuser() bool {
	if o != nil && !IsNil(o.Subuser) {
		return true
	}

	return false
}

// SetSubuser gets a reference to the given string and assigns it to the Subuser field.
func (o *ApiRgwUserUidKeyPostRequest) SetSubuser(v string) {
	o.Subuser = &v
}

func (o ApiRgwUserUidKeyPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRgwUserUidKeyPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !IsNil(o.DaemonName) {
		toSerialize["daemon_name"] = o.DaemonName
	}
	if !IsNil(o.GenerateKey) {
		toSerialize["generate_key"] = o.GenerateKey
	}
	if !IsNil(o.KeyType) {
		toSerialize["key_type"] = o.KeyType
	}
	if !IsNil(o.SecretKey) {
		toSerialize["secret_key"] = o.SecretKey
	}
	if !IsNil(o.Subuser) {
		toSerialize["subuser"] = o.Subuser
	}
	return toSerialize, nil
}

type NullableApiRgwUserUidKeyPostRequest struct {
	value *ApiRgwUserUidKeyPostRequest
	isSet bool
}

func (v NullableApiRgwUserUidKeyPostRequest) Get() *ApiRgwUserUidKeyPostRequest {
	return v.value
}

func (v *NullableApiRgwUserUidKeyPostRequest) Set(val *ApiRgwUserUidKeyPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRgwUserUidKeyPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRgwUserUidKeyPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRgwUserUidKeyPostRequest(val *ApiRgwUserUidKeyPostRequest) *NullableApiRgwUserUidKeyPostRequest {
	return &NullableApiRgwUserUidKeyPostRequest{value: val, isSet: true}
}

func (v NullableApiRgwUserUidKeyPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRgwUserUidKeyPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


