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

// checks if the ApiAuthPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAuthPostRequest{}

// ApiAuthPostRequest struct for ApiAuthPostRequest
type ApiAuthPostRequest struct {
	// Password
	Password string `json:"password"`
	// Token Time to Live (in hours)
	Ttl *int32 `json:"ttl,omitempty"`
	// Username
	Username string `json:"username"`
}

type _ApiAuthPostRequest ApiAuthPostRequest

// NewApiAuthPostRequest instantiates a new ApiAuthPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAuthPostRequest(password string, username string) *ApiAuthPostRequest {
	this := ApiAuthPostRequest{}
	this.Password = password
	this.Username = username
	return &this
}

// NewApiAuthPostRequestWithDefaults instantiates a new ApiAuthPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAuthPostRequestWithDefaults() *ApiAuthPostRequest {
	this := ApiAuthPostRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *ApiAuthPostRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ApiAuthPostRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ApiAuthPostRequest) SetPassword(v string) {
	o.Password = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ApiAuthPostRequest) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAuthPostRequest) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ApiAuthPostRequest) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *ApiAuthPostRequest) SetTtl(v int32) {
	o.Ttl = &v
}

// GetUsername returns the Username field value
func (o *ApiAuthPostRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ApiAuthPostRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ApiAuthPostRequest) SetUsername(v string) {
	o.Username = v
}

func (o ApiAuthPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAuthPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *ApiAuthPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"username",
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

	varApiAuthPostRequest := _ApiAuthPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiAuthPostRequest)

	if err != nil {
		return err
	}

	*o = ApiAuthPostRequest(varApiAuthPostRequest)

	return err
}

type NullableApiAuthPostRequest struct {
	value *ApiAuthPostRequest
	isSet bool
}

func (v NullableApiAuthPostRequest) Get() *ApiAuthPostRequest {
	return v.value
}

func (v *NullableApiAuthPostRequest) Set(val *ApiAuthPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAuthPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAuthPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAuthPostRequest(val *ApiAuthPostRequest) *NullableApiAuthPostRequest {
	return &NullableApiAuthPostRequest{value: val, isSet: true}
}

func (v NullableApiAuthPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAuthPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

