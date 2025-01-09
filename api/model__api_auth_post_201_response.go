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

// checks if the ApiAuthPost201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAuthPost201Response{}

// ApiAuthPost201Response struct for ApiAuthPost201Response
type ApiAuthPost201Response struct {
	Permissions ApiAuthPost201ResponsePermissions `json:"permissions"`
	// Password expiration date
	PwdExpirationDate string `json:"pwdExpirationDate"`
	// Is password update required?
	PwdUpdateRequired bool `json:"pwdUpdateRequired"`
	// Uses single sign on?
	Sso bool `json:"sso"`
	// Authentication Token
	Token string `json:"token"`
	// Username
	Username string `json:"username"`
}

type _ApiAuthPost201Response ApiAuthPost201Response

// NewApiAuthPost201Response instantiates a new ApiAuthPost201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAuthPost201Response(permissions ApiAuthPost201ResponsePermissions, pwdExpirationDate string, pwdUpdateRequired bool, sso bool, token string, username string) *ApiAuthPost201Response {
	this := ApiAuthPost201Response{}
	this.Permissions = permissions
	this.PwdExpirationDate = pwdExpirationDate
	this.PwdUpdateRequired = pwdUpdateRequired
	this.Sso = sso
	this.Token = token
	this.Username = username
	return &this
}

// NewApiAuthPost201ResponseWithDefaults instantiates a new ApiAuthPost201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAuthPost201ResponseWithDefaults() *ApiAuthPost201Response {
	this := ApiAuthPost201Response{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *ApiAuthPost201Response) GetPermissions() ApiAuthPost201ResponsePermissions {
	if o == nil {
		var ret ApiAuthPost201ResponsePermissions
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ApiAuthPost201Response) GetPermissionsOk() (*ApiAuthPost201ResponsePermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *ApiAuthPost201Response) SetPermissions(v ApiAuthPost201ResponsePermissions) {
	o.Permissions = v
}

// GetPwdExpirationDate returns the PwdExpirationDate field value
func (o *ApiAuthPost201Response) GetPwdExpirationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PwdExpirationDate
}

// GetPwdExpirationDateOk returns a tuple with the PwdExpirationDate field value
// and a boolean to check if the value has been set.
func (o *ApiAuthPost201Response) GetPwdExpirationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PwdExpirationDate, true
}

// SetPwdExpirationDate sets field value
func (o *ApiAuthPost201Response) SetPwdExpirationDate(v string) {
	o.PwdExpirationDate = v
}

// GetPwdUpdateRequired returns the PwdUpdateRequired field value
func (o *ApiAuthPost201Response) GetPwdUpdateRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PwdUpdateRequired
}

// GetPwdUpdateRequiredOk returns a tuple with the PwdUpdateRequired field value
// and a boolean to check if the value has been set.
func (o *ApiAuthPost201Response) GetPwdUpdateRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PwdUpdateRequired, true
}

// SetPwdUpdateRequired sets field value
func (o *ApiAuthPost201Response) SetPwdUpdateRequired(v bool) {
	o.PwdUpdateRequired = v
}

// GetSso returns the Sso field value
func (o *ApiAuthPost201Response) GetSso() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sso
}

// GetSsoOk returns a tuple with the Sso field value
// and a boolean to check if the value has been set.
func (o *ApiAuthPost201Response) GetSsoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sso, true
}

// SetSso sets field value
func (o *ApiAuthPost201Response) SetSso(v bool) {
	o.Sso = v
}

// GetToken returns the Token field value
func (o *ApiAuthPost201Response) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ApiAuthPost201Response) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ApiAuthPost201Response) SetToken(v string) {
	o.Token = v
}

// GetUsername returns the Username field value
func (o *ApiAuthPost201Response) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ApiAuthPost201Response) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ApiAuthPost201Response) SetUsername(v string) {
	o.Username = v
}

func (o ApiAuthPost201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAuthPost201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	toSerialize["pwdExpirationDate"] = o.PwdExpirationDate
	toSerialize["pwdUpdateRequired"] = o.PwdUpdateRequired
	toSerialize["sso"] = o.Sso
	toSerialize["token"] = o.Token
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *ApiAuthPost201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"permissions",
		"pwdExpirationDate",
		"pwdUpdateRequired",
		"sso",
		"token",
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

	varApiAuthPost201Response := _ApiAuthPost201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiAuthPost201Response)

	if err != nil {
		return err
	}

	*o = ApiAuthPost201Response(varApiAuthPost201Response)

	return err
}

type NullableApiAuthPost201Response struct {
	value *ApiAuthPost201Response
	isSet bool
}

func (v NullableApiAuthPost201Response) Get() *ApiAuthPost201Response {
	return v.value
}

func (v *NullableApiAuthPost201Response) Set(val *ApiAuthPost201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAuthPost201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAuthPost201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAuthPost201Response(val *ApiAuthPost201Response) *NullableApiAuthPost201Response {
	return &NullableApiAuthPost201Response{value: val, isSet: true}
}

func (v NullableApiAuthPost201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAuthPost201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

