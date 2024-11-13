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

// checks if the ApiCephfsSubvolumeSnapshotClonePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCephfsSubvolumeSnapshotClonePostRequest{}

// ApiCephfsSubvolumeSnapshotClonePostRequest struct for ApiCephfsSubvolumeSnapshotClonePostRequest
type ApiCephfsSubvolumeSnapshotClonePostRequest struct {
	CloneName string `json:"clone_name"`
	GroupName *string `json:"group_name,omitempty"`
	SnapName string `json:"snap_name"`
	SubvolName string `json:"subvol_name"`
	TargetGroupName *string `json:"target_group_name,omitempty"`
	VolName string `json:"vol_name"`
}

type _ApiCephfsSubvolumeSnapshotClonePostRequest ApiCephfsSubvolumeSnapshotClonePostRequest

// NewApiCephfsSubvolumeSnapshotClonePostRequest instantiates a new ApiCephfsSubvolumeSnapshotClonePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCephfsSubvolumeSnapshotClonePostRequest(cloneName string, snapName string, subvolName string, volName string) *ApiCephfsSubvolumeSnapshotClonePostRequest {
	this := ApiCephfsSubvolumeSnapshotClonePostRequest{}
	this.CloneName = cloneName
	var groupName string = ""
	this.GroupName = &groupName
	this.SnapName = snapName
	this.SubvolName = subvolName
	var targetGroupName string = ""
	this.TargetGroupName = &targetGroupName
	this.VolName = volName
	return &this
}

// NewApiCephfsSubvolumeSnapshotClonePostRequestWithDefaults instantiates a new ApiCephfsSubvolumeSnapshotClonePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCephfsSubvolumeSnapshotClonePostRequestWithDefaults() *ApiCephfsSubvolumeSnapshotClonePostRequest {
	this := ApiCephfsSubvolumeSnapshotClonePostRequest{}
	var groupName string = ""
	this.GroupName = &groupName
	var targetGroupName string = ""
	this.TargetGroupName = &targetGroupName
	return &this
}

// GetCloneName returns the CloneName field value
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetCloneName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloneName
}

// GetCloneNameOk returns a tuple with the CloneName field value
// and a boolean to check if the value has been set.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetCloneNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloneName, true
}

// SetCloneName sets field value
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetCloneName(v string) {
	o.CloneName = v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetGroupName(v string) {
	o.GroupName = &v
}

// GetSnapName returns the SnapName field value
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetSnapName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapName
}

// GetSnapNameOk returns a tuple with the SnapName field value
// and a boolean to check if the value has been set.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetSnapNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapName, true
}

// SetSnapName sets field value
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetSnapName(v string) {
	o.SnapName = v
}

// GetSubvolName returns the SubvolName field value
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetSubvolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubvolName
}

// GetSubvolNameOk returns a tuple with the SubvolName field value
// and a boolean to check if the value has been set.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetSubvolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubvolName, true
}

// SetSubvolName sets field value
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetSubvolName(v string) {
	o.SubvolName = v
}

// GetTargetGroupName returns the TargetGroupName field value if set, zero value otherwise.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetTargetGroupName() string {
	if o == nil || IsNil(o.TargetGroupName) {
		var ret string
		return ret
	}
	return *o.TargetGroupName
}

// GetTargetGroupNameOk returns a tuple with the TargetGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetTargetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetGroupName) {
		return nil, false
	}
	return o.TargetGroupName, true
}

// HasTargetGroupName returns a boolean if a field has been set.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) HasTargetGroupName() bool {
	if o != nil && !IsNil(o.TargetGroupName) {
		return true
	}

	return false
}

// SetTargetGroupName gets a reference to the given string and assigns it to the TargetGroupName field.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetTargetGroupName(v string) {
	o.TargetGroupName = &v
}

// GetVolName returns the VolName field value
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetVolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolName
}

// GetVolNameOk returns a tuple with the VolName field value
// and a boolean to check if the value has been set.
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetVolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolName, true
}

// SetVolName sets field value
func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetVolName(v string) {
	o.VolName = v
}

func (o ApiCephfsSubvolumeSnapshotClonePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCephfsSubvolumeSnapshotClonePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clone_name"] = o.CloneName
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	toSerialize["snap_name"] = o.SnapName
	toSerialize["subvol_name"] = o.SubvolName
	if !IsNil(o.TargetGroupName) {
		toSerialize["target_group_name"] = o.TargetGroupName
	}
	toSerialize["vol_name"] = o.VolName
	return toSerialize, nil
}

func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clone_name",
		"snap_name",
		"subvol_name",
		"vol_name",
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

	varApiCephfsSubvolumeSnapshotClonePostRequest := _ApiCephfsSubvolumeSnapshotClonePostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiCephfsSubvolumeSnapshotClonePostRequest)

	if err != nil {
		return err
	}

	*o = ApiCephfsSubvolumeSnapshotClonePostRequest(varApiCephfsSubvolumeSnapshotClonePostRequest)

	return err
}

type NullableApiCephfsSubvolumeSnapshotClonePostRequest struct {
	value *ApiCephfsSubvolumeSnapshotClonePostRequest
	isSet bool
}

func (v NullableApiCephfsSubvolumeSnapshotClonePostRequest) Get() *ApiCephfsSubvolumeSnapshotClonePostRequest {
	return v.value
}

func (v *NullableApiCephfsSubvolumeSnapshotClonePostRequest) Set(val *ApiCephfsSubvolumeSnapshotClonePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCephfsSubvolumeSnapshotClonePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCephfsSubvolumeSnapshotClonePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCephfsSubvolumeSnapshotClonePostRequest(val *ApiCephfsSubvolumeSnapshotClonePostRequest) *NullableApiCephfsSubvolumeSnapshotClonePostRequest {
	return &NullableApiCephfsSubvolumeSnapshotClonePostRequest{value: val, isSet: true}
}

func (v NullableApiCephfsSubvolumeSnapshotClonePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCephfsSubvolumeSnapshotClonePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

