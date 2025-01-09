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

// checks if the ApiCephfsFsIdWriteToFilePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCephfsFsIdWriteToFilePostRequest{}

// ApiCephfsFsIdWriteToFilePostRequest struct for ApiCephfsFsIdWriteToFilePostRequest
type ApiCephfsFsIdWriteToFilePostRequest struct {
	Buf string `json:"buf"`
	Path string `json:"path"`
}

type _ApiCephfsFsIdWriteToFilePostRequest ApiCephfsFsIdWriteToFilePostRequest

// NewApiCephfsFsIdWriteToFilePostRequest instantiates a new ApiCephfsFsIdWriteToFilePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCephfsFsIdWriteToFilePostRequest(buf string, path string) *ApiCephfsFsIdWriteToFilePostRequest {
	this := ApiCephfsFsIdWriteToFilePostRequest{}
	this.Buf = buf
	this.Path = path
	return &this
}

// NewApiCephfsFsIdWriteToFilePostRequestWithDefaults instantiates a new ApiCephfsFsIdWriteToFilePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCephfsFsIdWriteToFilePostRequestWithDefaults() *ApiCephfsFsIdWriteToFilePostRequest {
	this := ApiCephfsFsIdWriteToFilePostRequest{}
	return &this
}

// GetBuf returns the Buf field value
func (o *ApiCephfsFsIdWriteToFilePostRequest) GetBuf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Buf
}

// GetBufOk returns a tuple with the Buf field value
// and a boolean to check if the value has been set.
func (o *ApiCephfsFsIdWriteToFilePostRequest) GetBufOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Buf, true
}

// SetBuf sets field value
func (o *ApiCephfsFsIdWriteToFilePostRequest) SetBuf(v string) {
	o.Buf = v
}

// GetPath returns the Path field value
func (o *ApiCephfsFsIdWriteToFilePostRequest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ApiCephfsFsIdWriteToFilePostRequest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ApiCephfsFsIdWriteToFilePostRequest) SetPath(v string) {
	o.Path = v
}

func (o ApiCephfsFsIdWriteToFilePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCephfsFsIdWriteToFilePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buf"] = o.Buf
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

func (o *ApiCephfsFsIdWriteToFilePostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"buf",
		"path",
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

	varApiCephfsFsIdWriteToFilePostRequest := _ApiCephfsFsIdWriteToFilePostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiCephfsFsIdWriteToFilePostRequest)

	if err != nil {
		return err
	}

	*o = ApiCephfsFsIdWriteToFilePostRequest(varApiCephfsFsIdWriteToFilePostRequest)

	return err
}

type NullableApiCephfsFsIdWriteToFilePostRequest struct {
	value *ApiCephfsFsIdWriteToFilePostRequest
	isSet bool
}

func (v NullableApiCephfsFsIdWriteToFilePostRequest) Get() *ApiCephfsFsIdWriteToFilePostRequest {
	return v.value
}

func (v *NullableApiCephfsFsIdWriteToFilePostRequest) Set(val *ApiCephfsFsIdWriteToFilePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCephfsFsIdWriteToFilePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCephfsFsIdWriteToFilePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCephfsFsIdWriteToFilePostRequest(val *ApiCephfsFsIdWriteToFilePostRequest) *NullableApiCephfsFsIdWriteToFilePostRequest {
	return &NullableApiCephfsFsIdWriteToFilePostRequest{value: val, isSet: true}
}

func (v NullableApiCephfsFsIdWriteToFilePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCephfsFsIdWriteToFilePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

