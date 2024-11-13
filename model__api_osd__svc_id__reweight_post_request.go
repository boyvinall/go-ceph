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

// checks if the ApiOsdSvcIdReweightPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiOsdSvcIdReweightPostRequest{}

// ApiOsdSvcIdReweightPostRequest struct for ApiOsdSvcIdReweightPostRequest
type ApiOsdSvcIdReweightPostRequest struct {
	Weight string `json:"weight"`
}

type _ApiOsdSvcIdReweightPostRequest ApiOsdSvcIdReweightPostRequest

// NewApiOsdSvcIdReweightPostRequest instantiates a new ApiOsdSvcIdReweightPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiOsdSvcIdReweightPostRequest(weight string) *ApiOsdSvcIdReweightPostRequest {
	this := ApiOsdSvcIdReweightPostRequest{}
	this.Weight = weight
	return &this
}

// NewApiOsdSvcIdReweightPostRequestWithDefaults instantiates a new ApiOsdSvcIdReweightPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiOsdSvcIdReweightPostRequestWithDefaults() *ApiOsdSvcIdReweightPostRequest {
	this := ApiOsdSvcIdReweightPostRequest{}
	return &this
}

// GetWeight returns the Weight field value
func (o *ApiOsdSvcIdReweightPostRequest) GetWeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ApiOsdSvcIdReweightPostRequest) GetWeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *ApiOsdSvcIdReweightPostRequest) SetWeight(v string) {
	o.Weight = v
}

func (o ApiOsdSvcIdReweightPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiOsdSvcIdReweightPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

func (o *ApiOsdSvcIdReweightPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"weight",
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

	varApiOsdSvcIdReweightPostRequest := _ApiOsdSvcIdReweightPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiOsdSvcIdReweightPostRequest)

	if err != nil {
		return err
	}

	*o = ApiOsdSvcIdReweightPostRequest(varApiOsdSvcIdReweightPostRequest)

	return err
}

type NullableApiOsdSvcIdReweightPostRequest struct {
	value *ApiOsdSvcIdReweightPostRequest
	isSet bool
}

func (v NullableApiOsdSvcIdReweightPostRequest) Get() *ApiOsdSvcIdReweightPostRequest {
	return v.value
}

func (v *NullableApiOsdSvcIdReweightPostRequest) Set(val *ApiOsdSvcIdReweightPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiOsdSvcIdReweightPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiOsdSvcIdReweightPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiOsdSvcIdReweightPostRequest(val *ApiOsdSvcIdReweightPostRequest) *NullableApiOsdSvcIdReweightPostRequest {
	return &NullableApiOsdSvcIdReweightPostRequest{value: val, isSet: true}
}

func (v NullableApiOsdSvcIdReweightPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiOsdSvcIdReweightPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


