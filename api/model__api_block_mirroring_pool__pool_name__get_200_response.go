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

// checks if the ApiBlockMirroringPoolPoolNameGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBlockMirroringPoolPoolNameGet200Response{}

// ApiBlockMirroringPoolPoolNameGet200Response struct for ApiBlockMirroringPoolPoolNameGet200Response
type ApiBlockMirroringPoolPoolNameGet200Response struct {
	// Mirror Mode
	MirrorMode string `json:"mirror_mode"`
}

type _ApiBlockMirroringPoolPoolNameGet200Response ApiBlockMirroringPoolPoolNameGet200Response

// NewApiBlockMirroringPoolPoolNameGet200Response instantiates a new ApiBlockMirroringPoolPoolNameGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBlockMirroringPoolPoolNameGet200Response(mirrorMode string) *ApiBlockMirroringPoolPoolNameGet200Response {
	this := ApiBlockMirroringPoolPoolNameGet200Response{}
	this.MirrorMode = mirrorMode
	return &this
}

// NewApiBlockMirroringPoolPoolNameGet200ResponseWithDefaults instantiates a new ApiBlockMirroringPoolPoolNameGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBlockMirroringPoolPoolNameGet200ResponseWithDefaults() *ApiBlockMirroringPoolPoolNameGet200Response {
	this := ApiBlockMirroringPoolPoolNameGet200Response{}
	return &this
}

// GetMirrorMode returns the MirrorMode field value
func (o *ApiBlockMirroringPoolPoolNameGet200Response) GetMirrorMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MirrorMode
}

// GetMirrorModeOk returns a tuple with the MirrorMode field value
// and a boolean to check if the value has been set.
func (o *ApiBlockMirroringPoolPoolNameGet200Response) GetMirrorModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MirrorMode, true
}

// SetMirrorMode sets field value
func (o *ApiBlockMirroringPoolPoolNameGet200Response) SetMirrorMode(v string) {
	o.MirrorMode = v
}

func (o ApiBlockMirroringPoolPoolNameGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBlockMirroringPoolPoolNameGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mirror_mode"] = o.MirrorMode
	return toSerialize, nil
}

func (o *ApiBlockMirroringPoolPoolNameGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mirror_mode",
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

	varApiBlockMirroringPoolPoolNameGet200Response := _ApiBlockMirroringPoolPoolNameGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiBlockMirroringPoolPoolNameGet200Response)

	if err != nil {
		return err
	}

	*o = ApiBlockMirroringPoolPoolNameGet200Response(varApiBlockMirroringPoolPoolNameGet200Response)

	return err
}

type NullableApiBlockMirroringPoolPoolNameGet200Response struct {
	value *ApiBlockMirroringPoolPoolNameGet200Response
	isSet bool
}

func (v NullableApiBlockMirroringPoolPoolNameGet200Response) Get() *ApiBlockMirroringPoolPoolNameGet200Response {
	return v.value
}

func (v *NullableApiBlockMirroringPoolPoolNameGet200Response) Set(val *ApiBlockMirroringPoolPoolNameGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBlockMirroringPoolPoolNameGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBlockMirroringPoolPoolNameGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBlockMirroringPoolPoolNameGet200Response(val *ApiBlockMirroringPoolPoolNameGet200Response) *NullableApiBlockMirroringPoolPoolNameGet200Response {
	return &NullableApiBlockMirroringPoolPoolNameGet200Response{value: val, isSet: true}
}

func (v NullableApiBlockMirroringPoolPoolNameGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBlockMirroringPoolPoolNameGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


