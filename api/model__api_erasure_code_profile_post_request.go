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

// checks if the ApiErasureCodeProfilePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiErasureCodeProfilePostRequest{}

// ApiErasureCodeProfilePostRequest struct for ApiErasureCodeProfilePostRequest
type ApiErasureCodeProfilePostRequest struct {
	Name string `json:"name"`
}

type _ApiErasureCodeProfilePostRequest ApiErasureCodeProfilePostRequest

// NewApiErasureCodeProfilePostRequest instantiates a new ApiErasureCodeProfilePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErasureCodeProfilePostRequest(name string) *ApiErasureCodeProfilePostRequest {
	this := ApiErasureCodeProfilePostRequest{}
	this.Name = name
	return &this
}

// NewApiErasureCodeProfilePostRequestWithDefaults instantiates a new ApiErasureCodeProfilePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErasureCodeProfilePostRequestWithDefaults() *ApiErasureCodeProfilePostRequest {
	this := ApiErasureCodeProfilePostRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApiErasureCodeProfilePostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiErasureCodeProfilePostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiErasureCodeProfilePostRequest) SetName(v string) {
	o.Name = v
}

func (o ApiErasureCodeProfilePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiErasureCodeProfilePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ApiErasureCodeProfilePostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varApiErasureCodeProfilePostRequest := _ApiErasureCodeProfilePostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiErasureCodeProfilePostRequest)

	if err != nil {
		return err
	}

	*o = ApiErasureCodeProfilePostRequest(varApiErasureCodeProfilePostRequest)

	return err
}

type NullableApiErasureCodeProfilePostRequest struct {
	value *ApiErasureCodeProfilePostRequest
	isSet bool
}

func (v NullableApiErasureCodeProfilePostRequest) Get() *ApiErasureCodeProfilePostRequest {
	return v.value
}

func (v *NullableApiErasureCodeProfilePostRequest) Set(val *ApiErasureCodeProfilePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErasureCodeProfilePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErasureCodeProfilePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErasureCodeProfilePostRequest(val *ApiErasureCodeProfilePostRequest) *NullableApiErasureCodeProfilePostRequest {
	return &NullableApiErasureCodeProfilePostRequest{value: val, isSet: true}
}

func (v NullableApiErasureCodeProfilePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErasureCodeProfilePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


