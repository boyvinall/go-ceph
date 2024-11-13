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

// checks if the ApiIscsiDiscoveryauthPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiIscsiDiscoveryauthPutRequest{}

// ApiIscsiDiscoveryauthPutRequest struct for ApiIscsiDiscoveryauthPutRequest
type ApiIscsiDiscoveryauthPutRequest struct {
	// Mutual Password
	MutualPassword string `json:"mutual_password"`
	// Mutual UserName
	MutualUser string `json:"mutual_user"`
	// Password
	Password string `json:"password"`
	// Username
	User string `json:"user"`
}

type _ApiIscsiDiscoveryauthPutRequest ApiIscsiDiscoveryauthPutRequest

// NewApiIscsiDiscoveryauthPutRequest instantiates a new ApiIscsiDiscoveryauthPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiIscsiDiscoveryauthPutRequest(mutualPassword string, mutualUser string, password string, user string) *ApiIscsiDiscoveryauthPutRequest {
	this := ApiIscsiDiscoveryauthPutRequest{}
	this.MutualPassword = mutualPassword
	this.MutualUser = mutualUser
	this.Password = password
	this.User = user
	return &this
}

// NewApiIscsiDiscoveryauthPutRequestWithDefaults instantiates a new ApiIscsiDiscoveryauthPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiIscsiDiscoveryauthPutRequestWithDefaults() *ApiIscsiDiscoveryauthPutRequest {
	this := ApiIscsiDiscoveryauthPutRequest{}
	return &this
}

// GetMutualPassword returns the MutualPassword field value
func (o *ApiIscsiDiscoveryauthPutRequest) GetMutualPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MutualPassword
}

// GetMutualPasswordOk returns a tuple with the MutualPassword field value
// and a boolean to check if the value has been set.
func (o *ApiIscsiDiscoveryauthPutRequest) GetMutualPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutualPassword, true
}

// SetMutualPassword sets field value
func (o *ApiIscsiDiscoveryauthPutRequest) SetMutualPassword(v string) {
	o.MutualPassword = v
}

// GetMutualUser returns the MutualUser field value
func (o *ApiIscsiDiscoveryauthPutRequest) GetMutualUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MutualUser
}

// GetMutualUserOk returns a tuple with the MutualUser field value
// and a boolean to check if the value has been set.
func (o *ApiIscsiDiscoveryauthPutRequest) GetMutualUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutualUser, true
}

// SetMutualUser sets field value
func (o *ApiIscsiDiscoveryauthPutRequest) SetMutualUser(v string) {
	o.MutualUser = v
}

// GetPassword returns the Password field value
func (o *ApiIscsiDiscoveryauthPutRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ApiIscsiDiscoveryauthPutRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ApiIscsiDiscoveryauthPutRequest) SetPassword(v string) {
	o.Password = v
}

// GetUser returns the User field value
func (o *ApiIscsiDiscoveryauthPutRequest) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ApiIscsiDiscoveryauthPutRequest) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ApiIscsiDiscoveryauthPutRequest) SetUser(v string) {
	o.User = v
}

func (o ApiIscsiDiscoveryauthPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiIscsiDiscoveryauthPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mutual_password"] = o.MutualPassword
	toSerialize["mutual_user"] = o.MutualUser
	toSerialize["password"] = o.Password
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *ApiIscsiDiscoveryauthPutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mutual_password",
		"mutual_user",
		"password",
		"user",
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

	varApiIscsiDiscoveryauthPutRequest := _ApiIscsiDiscoveryauthPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiIscsiDiscoveryauthPutRequest)

	if err != nil {
		return err
	}

	*o = ApiIscsiDiscoveryauthPutRequest(varApiIscsiDiscoveryauthPutRequest)

	return err
}

type NullableApiIscsiDiscoveryauthPutRequest struct {
	value *ApiIscsiDiscoveryauthPutRequest
	isSet bool
}

func (v NullableApiIscsiDiscoveryauthPutRequest) Get() *ApiIscsiDiscoveryauthPutRequest {
	return v.value
}

func (v *NullableApiIscsiDiscoveryauthPutRequest) Set(val *ApiIscsiDiscoveryauthPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiIscsiDiscoveryauthPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiIscsiDiscoveryauthPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiIscsiDiscoveryauthPutRequest(val *ApiIscsiDiscoveryauthPutRequest) *NullableApiIscsiDiscoveryauthPutRequest {
	return &NullableApiIscsiDiscoveryauthPutRequest{value: val, isSet: true}
}

func (v NullableApiIscsiDiscoveryauthPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiIscsiDiscoveryauthPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

