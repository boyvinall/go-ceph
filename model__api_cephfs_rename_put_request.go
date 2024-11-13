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

// checks if the ApiCephfsRenamePutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCephfsRenamePutRequest{}

// ApiCephfsRenamePutRequest struct for ApiCephfsRenamePutRequest
type ApiCephfsRenamePutRequest struct {
	// Existing FS Name
	Name string `json:"name"`
	// New FS Name
	NewName string `json:"new_name"`
}

type _ApiCephfsRenamePutRequest ApiCephfsRenamePutRequest

// NewApiCephfsRenamePutRequest instantiates a new ApiCephfsRenamePutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCephfsRenamePutRequest(name string, newName string) *ApiCephfsRenamePutRequest {
	this := ApiCephfsRenamePutRequest{}
	this.Name = name
	this.NewName = newName
	return &this
}

// NewApiCephfsRenamePutRequestWithDefaults instantiates a new ApiCephfsRenamePutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCephfsRenamePutRequestWithDefaults() *ApiCephfsRenamePutRequest {
	this := ApiCephfsRenamePutRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApiCephfsRenamePutRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiCephfsRenamePutRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiCephfsRenamePutRequest) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value
func (o *ApiCephfsRenamePutRequest) GetNewName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value
// and a boolean to check if the value has been set.
func (o *ApiCephfsRenamePutRequest) GetNewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewName, true
}

// SetNewName sets field value
func (o *ApiCephfsRenamePutRequest) SetNewName(v string) {
	o.NewName = v
}

func (o ApiCephfsRenamePutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCephfsRenamePutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["new_name"] = o.NewName
	return toSerialize, nil
}

func (o *ApiCephfsRenamePutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"new_name",
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

	varApiCephfsRenamePutRequest := _ApiCephfsRenamePutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiCephfsRenamePutRequest)

	if err != nil {
		return err
	}

	*o = ApiCephfsRenamePutRequest(varApiCephfsRenamePutRequest)

	return err
}

type NullableApiCephfsRenamePutRequest struct {
	value *ApiCephfsRenamePutRequest
	isSet bool
}

func (v NullableApiCephfsRenamePutRequest) Get() *ApiCephfsRenamePutRequest {
	return v.value
}

func (v *NullableApiCephfsRenamePutRequest) Set(val *ApiCephfsRenamePutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCephfsRenamePutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCephfsRenamePutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCephfsRenamePutRequest(val *ApiCephfsRenamePutRequest) *NullableApiCephfsRenamePutRequest {
	return &NullableApiCephfsRenamePutRequest{value: val, isSet: true}
}

func (v NullableApiCephfsRenamePutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCephfsRenamePutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


