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

// checks if the ApiCephfsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCephfsPostRequest{}

// ApiCephfsPostRequest struct for ApiCephfsPostRequest
type ApiCephfsPostRequest struct {
	Name string `json:"name"`
	ServiceSpec string `json:"service_spec"`
}

type _ApiCephfsPostRequest ApiCephfsPostRequest

// NewApiCephfsPostRequest instantiates a new ApiCephfsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCephfsPostRequest(name string, serviceSpec string) *ApiCephfsPostRequest {
	this := ApiCephfsPostRequest{}
	this.Name = name
	this.ServiceSpec = serviceSpec
	return &this
}

// NewApiCephfsPostRequestWithDefaults instantiates a new ApiCephfsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCephfsPostRequestWithDefaults() *ApiCephfsPostRequest {
	this := ApiCephfsPostRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApiCephfsPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiCephfsPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiCephfsPostRequest) SetName(v string) {
	o.Name = v
}

// GetServiceSpec returns the ServiceSpec field value
func (o *ApiCephfsPostRequest) GetServiceSpec() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceSpec
}

// GetServiceSpecOk returns a tuple with the ServiceSpec field value
// and a boolean to check if the value has been set.
func (o *ApiCephfsPostRequest) GetServiceSpecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceSpec, true
}

// SetServiceSpec sets field value
func (o *ApiCephfsPostRequest) SetServiceSpec(v string) {
	o.ServiceSpec = v
}

func (o ApiCephfsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCephfsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["service_spec"] = o.ServiceSpec
	return toSerialize, nil
}

func (o *ApiCephfsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"service_spec",
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

	varApiCephfsPostRequest := _ApiCephfsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiCephfsPostRequest)

	if err != nil {
		return err
	}

	*o = ApiCephfsPostRequest(varApiCephfsPostRequest)

	return err
}

type NullableApiCephfsPostRequest struct {
	value *ApiCephfsPostRequest
	isSet bool
}

func (v NullableApiCephfsPostRequest) Get() *ApiCephfsPostRequest {
	return v.value
}

func (v *NullableApiCephfsPostRequest) Set(val *ApiCephfsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCephfsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCephfsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCephfsPostRequest(val *ApiCephfsPostRequest) *NullableApiCephfsPostRequest {
	return &NullableApiCephfsPostRequest{value: val, isSet: true}
}

func (v NullableApiCephfsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCephfsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

