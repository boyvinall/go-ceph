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

// checks if the ApiBlockMirroringSiteNamePutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBlockMirroringSiteNamePutRequest{}

// ApiBlockMirroringSiteNamePutRequest struct for ApiBlockMirroringSiteNamePutRequest
type ApiBlockMirroringSiteNamePutRequest struct {
	SiteName string `json:"site_name"`
}

type _ApiBlockMirroringSiteNamePutRequest ApiBlockMirroringSiteNamePutRequest

// NewApiBlockMirroringSiteNamePutRequest instantiates a new ApiBlockMirroringSiteNamePutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBlockMirroringSiteNamePutRequest(siteName string) *ApiBlockMirroringSiteNamePutRequest {
	this := ApiBlockMirroringSiteNamePutRequest{}
	this.SiteName = siteName
	return &this
}

// NewApiBlockMirroringSiteNamePutRequestWithDefaults instantiates a new ApiBlockMirroringSiteNamePutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBlockMirroringSiteNamePutRequestWithDefaults() *ApiBlockMirroringSiteNamePutRequest {
	this := ApiBlockMirroringSiteNamePutRequest{}
	return &this
}

// GetSiteName returns the SiteName field value
func (o *ApiBlockMirroringSiteNamePutRequest) GetSiteName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value
// and a boolean to check if the value has been set.
func (o *ApiBlockMirroringSiteNamePutRequest) GetSiteNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteName, true
}

// SetSiteName sets field value
func (o *ApiBlockMirroringSiteNamePutRequest) SetSiteName(v string) {
	o.SiteName = v
}

func (o ApiBlockMirroringSiteNamePutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBlockMirroringSiteNamePutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["site_name"] = o.SiteName
	return toSerialize, nil
}

func (o *ApiBlockMirroringSiteNamePutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"site_name",
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

	varApiBlockMirroringSiteNamePutRequest := _ApiBlockMirroringSiteNamePutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiBlockMirroringSiteNamePutRequest)

	if err != nil {
		return err
	}

	*o = ApiBlockMirroringSiteNamePutRequest(varApiBlockMirroringSiteNamePutRequest)

	return err
}

type NullableApiBlockMirroringSiteNamePutRequest struct {
	value *ApiBlockMirroringSiteNamePutRequest
	isSet bool
}

func (v NullableApiBlockMirroringSiteNamePutRequest) Get() *ApiBlockMirroringSiteNamePutRequest {
	return v.value
}

func (v *NullableApiBlockMirroringSiteNamePutRequest) Set(val *ApiBlockMirroringSiteNamePutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBlockMirroringSiteNamePutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBlockMirroringSiteNamePutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBlockMirroringSiteNamePutRequest(val *ApiBlockMirroringSiteNamePutRequest) *NullableApiBlockMirroringSiteNamePutRequest {
	return &NullableApiBlockMirroringSiteNamePutRequest{value: val, isSet: true}
}

func (v NullableApiBlockMirroringSiteNamePutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBlockMirroringSiteNamePutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


