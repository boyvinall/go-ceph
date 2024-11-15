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

// checks if the ApiOsdSvcIdMarkPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiOsdSvcIdMarkPutRequest{}

// ApiOsdSvcIdMarkPutRequest struct for ApiOsdSvcIdMarkPutRequest
type ApiOsdSvcIdMarkPutRequest struct {
	Action string `json:"action"`
}

type _ApiOsdSvcIdMarkPutRequest ApiOsdSvcIdMarkPutRequest

// NewApiOsdSvcIdMarkPutRequest instantiates a new ApiOsdSvcIdMarkPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiOsdSvcIdMarkPutRequest(action string) *ApiOsdSvcIdMarkPutRequest {
	this := ApiOsdSvcIdMarkPutRequest{}
	this.Action = action
	return &this
}

// NewApiOsdSvcIdMarkPutRequestWithDefaults instantiates a new ApiOsdSvcIdMarkPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiOsdSvcIdMarkPutRequestWithDefaults() *ApiOsdSvcIdMarkPutRequest {
	this := ApiOsdSvcIdMarkPutRequest{}
	return &this
}

// GetAction returns the Action field value
func (o *ApiOsdSvcIdMarkPutRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ApiOsdSvcIdMarkPutRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ApiOsdSvcIdMarkPutRequest) SetAction(v string) {
	o.Action = v
}

func (o ApiOsdSvcIdMarkPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiOsdSvcIdMarkPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

func (o *ApiOsdSvcIdMarkPutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varApiOsdSvcIdMarkPutRequest := _ApiOsdSvcIdMarkPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiOsdSvcIdMarkPutRequest)

	if err != nil {
		return err
	}

	*o = ApiOsdSvcIdMarkPutRequest(varApiOsdSvcIdMarkPutRequest)

	return err
}

type NullableApiOsdSvcIdMarkPutRequest struct {
	value *ApiOsdSvcIdMarkPutRequest
	isSet bool
}

func (v NullableApiOsdSvcIdMarkPutRequest) Get() *ApiOsdSvcIdMarkPutRequest {
	return v.value
}

func (v *NullableApiOsdSvcIdMarkPutRequest) Set(val *ApiOsdSvcIdMarkPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiOsdSvcIdMarkPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiOsdSvcIdMarkPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiOsdSvcIdMarkPutRequest(val *ApiOsdSvcIdMarkPutRequest) *NullableApiOsdSvcIdMarkPutRequest {
	return &NullableApiOsdSvcIdMarkPutRequest{value: val, isSet: true}
}

func (v NullableApiOsdSvcIdMarkPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiOsdSvcIdMarkPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


