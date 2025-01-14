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

// checks if the ApiRgwMultisiteSyncPolicyGroupPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRgwMultisiteSyncPolicyGroupPutRequest{}

// ApiRgwMultisiteSyncPolicyGroupPutRequest struct for ApiRgwMultisiteSyncPolicyGroupPutRequest
type ApiRgwMultisiteSyncPolicyGroupPutRequest struct {
	BucketName *string `json:"bucket_name,omitempty"`
	GroupId string `json:"group_id"`
	Status string `json:"status"`
}

type _ApiRgwMultisiteSyncPolicyGroupPutRequest ApiRgwMultisiteSyncPolicyGroupPutRequest

// NewApiRgwMultisiteSyncPolicyGroupPutRequest instantiates a new ApiRgwMultisiteSyncPolicyGroupPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRgwMultisiteSyncPolicyGroupPutRequest(groupId string, status string) *ApiRgwMultisiteSyncPolicyGroupPutRequest {
	this := ApiRgwMultisiteSyncPolicyGroupPutRequest{}
	var bucketName string = ""
	this.BucketName = &bucketName
	this.GroupId = groupId
	this.Status = status
	return &this
}

// NewApiRgwMultisiteSyncPolicyGroupPutRequestWithDefaults instantiates a new ApiRgwMultisiteSyncPolicyGroupPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRgwMultisiteSyncPolicyGroupPutRequestWithDefaults() *ApiRgwMultisiteSyncPolicyGroupPutRequest {
	this := ApiRgwMultisiteSyncPolicyGroupPutRequest{}
	var bucketName string = ""
	this.BucketName = &bucketName
	return &this
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) SetBucketName(v string) {
	o.BucketName = &v
}

// GetGroupId returns the GroupId field value
func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) SetGroupId(v string) {
	o.GroupId = v
}

// GetStatus returns the Status field value
func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) SetStatus(v string) {
	o.Status = v
}

func (o ApiRgwMultisiteSyncPolicyGroupPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRgwMultisiteSyncPolicyGroupPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BucketName) {
		toSerialize["bucket_name"] = o.BucketName
	}
	toSerialize["group_id"] = o.GroupId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group_id",
		"status",
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

	varApiRgwMultisiteSyncPolicyGroupPutRequest := _ApiRgwMultisiteSyncPolicyGroupPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiRgwMultisiteSyncPolicyGroupPutRequest)

	if err != nil {
		return err
	}

	*o = ApiRgwMultisiteSyncPolicyGroupPutRequest(varApiRgwMultisiteSyncPolicyGroupPutRequest)

	return err
}

type NullableApiRgwMultisiteSyncPolicyGroupPutRequest struct {
	value *ApiRgwMultisiteSyncPolicyGroupPutRequest
	isSet bool
}

func (v NullableApiRgwMultisiteSyncPolicyGroupPutRequest) Get() *ApiRgwMultisiteSyncPolicyGroupPutRequest {
	return v.value
}

func (v *NullableApiRgwMultisiteSyncPolicyGroupPutRequest) Set(val *ApiRgwMultisiteSyncPolicyGroupPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRgwMultisiteSyncPolicyGroupPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRgwMultisiteSyncPolicyGroupPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRgwMultisiteSyncPolicyGroupPutRequest(val *ApiRgwMultisiteSyncPolicyGroupPutRequest) *NullableApiRgwMultisiteSyncPolicyGroupPutRequest {
	return &NullableApiRgwMultisiteSyncPolicyGroupPutRequest{value: val, isSet: true}
}

func (v NullableApiRgwMultisiteSyncPolicyGroupPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRgwMultisiteSyncPolicyGroupPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


