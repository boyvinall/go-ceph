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

// checks if the ApiRgwUserUidSubuserPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRgwUserUidSubuserPostRequest{}

// ApiRgwUserUidSubuserPostRequest struct for ApiRgwUserUidSubuserPostRequest
type ApiRgwUserUidSubuserPostRequest struct {
	Access string `json:"access"`
	AccessKey *string `json:"access_key,omitempty"`
	DaemonName *string `json:"daemon_name,omitempty"`
	GenerateSecret *string `json:"generate_secret,omitempty"`
	KeyType *string `json:"key_type,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
	Subuser string `json:"subuser"`
}

type _ApiRgwUserUidSubuserPostRequest ApiRgwUserUidSubuserPostRequest

// NewApiRgwUserUidSubuserPostRequest instantiates a new ApiRgwUserUidSubuserPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRgwUserUidSubuserPostRequest(access string, subuser string) *ApiRgwUserUidSubuserPostRequest {
	this := ApiRgwUserUidSubuserPostRequest{}
	this.Access = access
	var generateSecret string = "true"
	this.GenerateSecret = &generateSecret
	var keyType string = "s3"
	this.KeyType = &keyType
	this.Subuser = subuser
	return &this
}

// NewApiRgwUserUidSubuserPostRequestWithDefaults instantiates a new ApiRgwUserUidSubuserPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRgwUserUidSubuserPostRequestWithDefaults() *ApiRgwUserUidSubuserPostRequest {
	this := ApiRgwUserUidSubuserPostRequest{}
	var generateSecret string = "true"
	this.GenerateSecret = &generateSecret
	var keyType string = "s3"
	this.KeyType = &keyType
	return &this
}

// GetAccess returns the Access field value
func (o *ApiRgwUserUidSubuserPostRequest) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidSubuserPostRequest) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *ApiRgwUserUidSubuserPostRequest) SetAccess(v string) {
	o.Access = v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *ApiRgwUserUidSubuserPostRequest) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidSubuserPostRequest) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *ApiRgwUserUidSubuserPostRequest) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *ApiRgwUserUidSubuserPostRequest) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetDaemonName returns the DaemonName field value if set, zero value otherwise.
func (o *ApiRgwUserUidSubuserPostRequest) GetDaemonName() string {
	if o == nil || IsNil(o.DaemonName) {
		var ret string
		return ret
	}
	return *o.DaemonName
}

// GetDaemonNameOk returns a tuple with the DaemonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidSubuserPostRequest) GetDaemonNameOk() (*string, bool) {
	if o == nil || IsNil(o.DaemonName) {
		return nil, false
	}
	return o.DaemonName, true
}

// HasDaemonName returns a boolean if a field has been set.
func (o *ApiRgwUserUidSubuserPostRequest) HasDaemonName() bool {
	if o != nil && !IsNil(o.DaemonName) {
		return true
	}

	return false
}

// SetDaemonName gets a reference to the given string and assigns it to the DaemonName field.
func (o *ApiRgwUserUidSubuserPostRequest) SetDaemonName(v string) {
	o.DaemonName = &v
}

// GetGenerateSecret returns the GenerateSecret field value if set, zero value otherwise.
func (o *ApiRgwUserUidSubuserPostRequest) GetGenerateSecret() string {
	if o == nil || IsNil(o.GenerateSecret) {
		var ret string
		return ret
	}
	return *o.GenerateSecret
}

// GetGenerateSecretOk returns a tuple with the GenerateSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidSubuserPostRequest) GetGenerateSecretOk() (*string, bool) {
	if o == nil || IsNil(o.GenerateSecret) {
		return nil, false
	}
	return o.GenerateSecret, true
}

// HasGenerateSecret returns a boolean if a field has been set.
func (o *ApiRgwUserUidSubuserPostRequest) HasGenerateSecret() bool {
	if o != nil && !IsNil(o.GenerateSecret) {
		return true
	}

	return false
}

// SetGenerateSecret gets a reference to the given string and assigns it to the GenerateSecret field.
func (o *ApiRgwUserUidSubuserPostRequest) SetGenerateSecret(v string) {
	o.GenerateSecret = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *ApiRgwUserUidSubuserPostRequest) GetKeyType() string {
	if o == nil || IsNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidSubuserPostRequest) GetKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *ApiRgwUserUidSubuserPostRequest) HasKeyType() bool {
	if o != nil && !IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *ApiRgwUserUidSubuserPostRequest) SetKeyType(v string) {
	o.KeyType = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *ApiRgwUserUidSubuserPostRequest) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidSubuserPostRequest) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *ApiRgwUserUidSubuserPostRequest) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *ApiRgwUserUidSubuserPostRequest) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetSubuser returns the Subuser field value
func (o *ApiRgwUserUidSubuserPostRequest) GetSubuser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subuser
}

// GetSubuserOk returns a tuple with the Subuser field value
// and a boolean to check if the value has been set.
func (o *ApiRgwUserUidSubuserPostRequest) GetSubuserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subuser, true
}

// SetSubuser sets field value
func (o *ApiRgwUserUidSubuserPostRequest) SetSubuser(v string) {
	o.Subuser = v
}

func (o ApiRgwUserUidSubuserPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRgwUserUidSubuserPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access"] = o.Access
	if !IsNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !IsNil(o.DaemonName) {
		toSerialize["daemon_name"] = o.DaemonName
	}
	if !IsNil(o.GenerateSecret) {
		toSerialize["generate_secret"] = o.GenerateSecret
	}
	if !IsNil(o.KeyType) {
		toSerialize["key_type"] = o.KeyType
	}
	if !IsNil(o.SecretKey) {
		toSerialize["secret_key"] = o.SecretKey
	}
	toSerialize["subuser"] = o.Subuser
	return toSerialize, nil
}

func (o *ApiRgwUserUidSubuserPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access",
		"subuser",
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

	varApiRgwUserUidSubuserPostRequest := _ApiRgwUserUidSubuserPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiRgwUserUidSubuserPostRequest)

	if err != nil {
		return err
	}

	*o = ApiRgwUserUidSubuserPostRequest(varApiRgwUserUidSubuserPostRequest)

	return err
}

type NullableApiRgwUserUidSubuserPostRequest struct {
	value *ApiRgwUserUidSubuserPostRequest
	isSet bool
}

func (v NullableApiRgwUserUidSubuserPostRequest) Get() *ApiRgwUserUidSubuserPostRequest {
	return v.value
}

func (v *NullableApiRgwUserUidSubuserPostRequest) Set(val *ApiRgwUserUidSubuserPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRgwUserUidSubuserPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRgwUserUidSubuserPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRgwUserUidSubuserPostRequest(val *ApiRgwUserUidSubuserPostRequest) *NullableApiRgwUserUidSubuserPostRequest {
	return &NullableApiRgwUserUidSubuserPostRequest{value: val, isSet: true}
}

func (v NullableApiRgwUserUidSubuserPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRgwUserUidSubuserPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


