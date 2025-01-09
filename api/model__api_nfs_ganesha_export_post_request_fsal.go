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

// checks if the ApiNfsGaneshaExportPostRequestFsal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiNfsGaneshaExportPostRequestFsal{}

// ApiNfsGaneshaExportPostRequestFsal FSAL configuration
type ApiNfsGaneshaExportPostRequestFsal struct {
	// CephFS filesystem name
	FsName *string `json:"fs_name,omitempty"`
	// name of FSAL
	Name string `json:"name"`
	// Name of xattr for security label
	SecLabelXattr *string `json:"sec_label_xattr,omitempty"`
}

type _ApiNfsGaneshaExportPostRequestFsal ApiNfsGaneshaExportPostRequestFsal

// NewApiNfsGaneshaExportPostRequestFsal instantiates a new ApiNfsGaneshaExportPostRequestFsal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNfsGaneshaExportPostRequestFsal(name string) *ApiNfsGaneshaExportPostRequestFsal {
	this := ApiNfsGaneshaExportPostRequestFsal{}
	this.Name = name
	return &this
}

// NewApiNfsGaneshaExportPostRequestFsalWithDefaults instantiates a new ApiNfsGaneshaExportPostRequestFsal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNfsGaneshaExportPostRequestFsalWithDefaults() *ApiNfsGaneshaExportPostRequestFsal {
	this := ApiNfsGaneshaExportPostRequestFsal{}
	return &this
}

// GetFsName returns the FsName field value if set, zero value otherwise.
func (o *ApiNfsGaneshaExportPostRequestFsal) GetFsName() string {
	if o == nil || IsNil(o.FsName) {
		var ret string
		return ret
	}
	return *o.FsName
}

// GetFsNameOk returns a tuple with the FsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNfsGaneshaExportPostRequestFsal) GetFsNameOk() (*string, bool) {
	if o == nil || IsNil(o.FsName) {
		return nil, false
	}
	return o.FsName, true
}

// HasFsName returns a boolean if a field has been set.
func (o *ApiNfsGaneshaExportPostRequestFsal) HasFsName() bool {
	if o != nil && !IsNil(o.FsName) {
		return true
	}

	return false
}

// SetFsName gets a reference to the given string and assigns it to the FsName field.
func (o *ApiNfsGaneshaExportPostRequestFsal) SetFsName(v string) {
	o.FsName = &v
}

// GetName returns the Name field value
func (o *ApiNfsGaneshaExportPostRequestFsal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiNfsGaneshaExportPostRequestFsal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiNfsGaneshaExportPostRequestFsal) SetName(v string) {
	o.Name = v
}

// GetSecLabelXattr returns the SecLabelXattr field value if set, zero value otherwise.
func (o *ApiNfsGaneshaExportPostRequestFsal) GetSecLabelXattr() string {
	if o == nil || IsNil(o.SecLabelXattr) {
		var ret string
		return ret
	}
	return *o.SecLabelXattr
}

// GetSecLabelXattrOk returns a tuple with the SecLabelXattr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNfsGaneshaExportPostRequestFsal) GetSecLabelXattrOk() (*string, bool) {
	if o == nil || IsNil(o.SecLabelXattr) {
		return nil, false
	}
	return o.SecLabelXattr, true
}

// HasSecLabelXattr returns a boolean if a field has been set.
func (o *ApiNfsGaneshaExportPostRequestFsal) HasSecLabelXattr() bool {
	if o != nil && !IsNil(o.SecLabelXattr) {
		return true
	}

	return false
}

// SetSecLabelXattr gets a reference to the given string and assigns it to the SecLabelXattr field.
func (o *ApiNfsGaneshaExportPostRequestFsal) SetSecLabelXattr(v string) {
	o.SecLabelXattr = &v
}

func (o ApiNfsGaneshaExportPostRequestFsal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNfsGaneshaExportPostRequestFsal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsName) {
		toSerialize["fs_name"] = o.FsName
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.SecLabelXattr) {
		toSerialize["sec_label_xattr"] = o.SecLabelXattr
	}
	return toSerialize, nil
}

func (o *ApiNfsGaneshaExportPostRequestFsal) UnmarshalJSON(data []byte) (err error) {
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

	varApiNfsGaneshaExportPostRequestFsal := _ApiNfsGaneshaExportPostRequestFsal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiNfsGaneshaExportPostRequestFsal)

	if err != nil {
		return err
	}

	*o = ApiNfsGaneshaExportPostRequestFsal(varApiNfsGaneshaExportPostRequestFsal)

	return err
}

type NullableApiNfsGaneshaExportPostRequestFsal struct {
	value *ApiNfsGaneshaExportPostRequestFsal
	isSet bool
}

func (v NullableApiNfsGaneshaExportPostRequestFsal) Get() *ApiNfsGaneshaExportPostRequestFsal {
	return v.value
}

func (v *NullableApiNfsGaneshaExportPostRequestFsal) Set(val *ApiNfsGaneshaExportPostRequestFsal) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNfsGaneshaExportPostRequestFsal) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNfsGaneshaExportPostRequestFsal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNfsGaneshaExportPostRequestFsal(val *ApiNfsGaneshaExportPostRequestFsal) *NullableApiNfsGaneshaExportPostRequestFsal {
	return &NullableApiNfsGaneshaExportPostRequestFsal{value: val, isSet: true}
}

func (v NullableApiNfsGaneshaExportPostRequestFsal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNfsGaneshaExportPostRequestFsal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

