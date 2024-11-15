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

// checks if the ApiAuthCheckPost201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAuthCheckPost201Response{}

// ApiAuthCheckPost201Response struct for ApiAuthCheckPost201Response
type ApiAuthCheckPost201Response struct {
	Permissions ApiAuthCheckPost201ResponsePermissions `json:"permissions"`
	// Is password update required?
	PwdUpdateRequired bool `json:"pwdUpdateRequired"`
	// Uses single sign on?
	Sso bool `json:"sso"`
	// Username
	Username string `json:"username"`
}

type _ApiAuthCheckPost201Response ApiAuthCheckPost201Response

// NewApiAuthCheckPost201Response instantiates a new ApiAuthCheckPost201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAuthCheckPost201Response(permissions ApiAuthCheckPost201ResponsePermissions, pwdUpdateRequired bool, sso bool, username string) *ApiAuthCheckPost201Response {
	this := ApiAuthCheckPost201Response{}
	this.Permissions = permissions
	this.PwdUpdateRequired = pwdUpdateRequired
	this.Sso = sso
	this.Username = username
	return &this
}

// NewApiAuthCheckPost201ResponseWithDefaults instantiates a new ApiAuthCheckPost201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAuthCheckPost201ResponseWithDefaults() *ApiAuthCheckPost201Response {
	this := ApiAuthCheckPost201Response{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *ApiAuthCheckPost201Response) GetPermissions() ApiAuthCheckPost201ResponsePermissions {
	if o == nil {
		var ret ApiAuthCheckPost201ResponsePermissions
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ApiAuthCheckPost201Response) GetPermissionsOk() (*ApiAuthCheckPost201ResponsePermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *ApiAuthCheckPost201Response) SetPermissions(v ApiAuthCheckPost201ResponsePermissions) {
	o.Permissions = v
}

// GetPwdUpdateRequired returns the PwdUpdateRequired field value
func (o *ApiAuthCheckPost201Response) GetPwdUpdateRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PwdUpdateRequired
}

// GetPwdUpdateRequiredOk returns a tuple with the PwdUpdateRequired field value
// and a boolean to check if the value has been set.
func (o *ApiAuthCheckPost201Response) GetPwdUpdateRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PwdUpdateRequired, true
}

// SetPwdUpdateRequired sets field value
func (o *ApiAuthCheckPost201Response) SetPwdUpdateRequired(v bool) {
	o.PwdUpdateRequired = v
}

// GetSso returns the Sso field value
func (o *ApiAuthCheckPost201Response) GetSso() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sso
}

// GetSsoOk returns a tuple with the Sso field value
// and a boolean to check if the value has been set.
func (o *ApiAuthCheckPost201Response) GetSsoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sso, true
}

// SetSso sets field value
func (o *ApiAuthCheckPost201Response) SetSso(v bool) {
	o.Sso = v
}

// GetUsername returns the Username field value
func (o *ApiAuthCheckPost201Response) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ApiAuthCheckPost201Response) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ApiAuthCheckPost201Response) SetUsername(v string) {
	o.Username = v
}

func (o ApiAuthCheckPost201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAuthCheckPost201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	toSerialize["pwdUpdateRequired"] = o.PwdUpdateRequired
	toSerialize["sso"] = o.Sso
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *ApiAuthCheckPost201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"permissions",
		"pwdUpdateRequired",
		"sso",
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

	varApiAuthCheckPost201Response := _ApiAuthCheckPost201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiAuthCheckPost201Response)

	if err != nil {
		return err
	}

	*o = ApiAuthCheckPost201Response(varApiAuthCheckPost201Response)

	return err
}

type NullableApiAuthCheckPost201Response struct {
	value *ApiAuthCheckPost201Response
	isSet bool
}

func (v NullableApiAuthCheckPost201Response) Get() *ApiAuthCheckPost201Response {
	return v.value
}

func (v *NullableApiAuthCheckPost201Response) Set(val *ApiAuthCheckPost201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAuthCheckPost201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAuthCheckPost201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAuthCheckPost201Response(val *ApiAuthCheckPost201Response) *NullableApiAuthCheckPost201Response {
	return &NullableApiAuthCheckPost201Response{value: val, isSet: true}
}

func (v NullableApiAuthCheckPost201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAuthCheckPost201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


