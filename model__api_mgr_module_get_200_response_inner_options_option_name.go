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

// checks if the ApiMgrModuleGet200ResponseInnerOptionsOptionName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMgrModuleGet200ResponseInnerOptionsOptionName{}

// ApiMgrModuleGet200ResponseInnerOptionsOptionName Options
type ApiMgrModuleGet200ResponseInnerOptionsOptionName struct {
	// Default value for the option
	DefaultValue int32 `json:"default_value"`
	// Description of the option
	Desc string `json:"desc"`
	// 
	EnumAllowed []string `json:"enum_allowed"`
	// List of flags associated
	Flags int32 `json:"flags"`
	// Option level
	Level string `json:"level"`
	// Elaborated description
	LongDesc string `json:"long_desc"`
	// Maximum value
	Max string `json:"max"`
	// Minimum value
	Min string `json:"min"`
	// Name of the option
	Name string `json:"name"`
	// Related options
	SeeAlso []string `json:"see_also"`
	// Tags associated with the option
	Tags []string `json:"tags"`
	// Type of the option
	Type string `json:"type"`
}

type _ApiMgrModuleGet200ResponseInnerOptionsOptionName ApiMgrModuleGet200ResponseInnerOptionsOptionName

// NewApiMgrModuleGet200ResponseInnerOptionsOptionName instantiates a new ApiMgrModuleGet200ResponseInnerOptionsOptionName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMgrModuleGet200ResponseInnerOptionsOptionName(defaultValue int32, desc string, enumAllowed []string, flags int32, level string, longDesc string, max string, min string, name string, seeAlso []string, tags []string, type_ string) *ApiMgrModuleGet200ResponseInnerOptionsOptionName {
	this := ApiMgrModuleGet200ResponseInnerOptionsOptionName{}
	this.DefaultValue = defaultValue
	this.Desc = desc
	this.EnumAllowed = enumAllowed
	this.Flags = flags
	this.Level = level
	this.LongDesc = longDesc
	this.Max = max
	this.Min = min
	this.Name = name
	this.SeeAlso = seeAlso
	this.Tags = tags
	this.Type = type_
	return &this
}

// NewApiMgrModuleGet200ResponseInnerOptionsOptionNameWithDefaults instantiates a new ApiMgrModuleGet200ResponseInnerOptionsOptionName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMgrModuleGet200ResponseInnerOptionsOptionNameWithDefaults() *ApiMgrModuleGet200ResponseInnerOptionsOptionName {
	this := ApiMgrModuleGet200ResponseInnerOptionsOptionName{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetDefaultValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetDefaultValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// SetDefaultValue sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetDefaultValue(v int32) {
	o.DefaultValue = v
}

// GetDesc returns the Desc field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Desc, true
}

// SetDesc sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetDesc(v string) {
	o.Desc = v
}

// GetEnumAllowed returns the EnumAllowed field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetEnumAllowed() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EnumAllowed
}

// GetEnumAllowedOk returns a tuple with the EnumAllowed field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetEnumAllowedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnumAllowed, true
}

// SetEnumAllowed sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetEnumAllowed(v []string) {
	o.EnumAllowed = v
}

// GetFlags returns the Flags field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetFlags(v int32) {
	o.Flags = v
}

// GetLevel returns the Level field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetLevel(v string) {
	o.Level = v
}

// GetLongDesc returns the LongDesc field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetLongDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LongDesc
}

// GetLongDescOk returns a tuple with the LongDesc field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetLongDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LongDesc, true
}

// SetLongDesc sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetLongDesc(v string) {
	o.LongDesc = v
}

// GetMax returns the Max field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetMax() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetMaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetMax(v string) {
	o.Max = v
}

// GetMin returns the Min field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetMin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetMinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetMin(v string) {
	o.Min = v
}

// GetName returns the Name field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetName(v string) {
	o.Name = v
}

// GetSeeAlso returns the SeeAlso field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetSeeAlso() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SeeAlso
}

// GetSeeAlsoOk returns a tuple with the SeeAlso field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetSeeAlsoOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeeAlso, true
}

// SetSeeAlso sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetSeeAlso(v []string) {
	o.SeeAlso = v
}

// GetTags returns the Tags field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) SetType(v string) {
	o.Type = v
}

func (o ApiMgrModuleGet200ResponseInnerOptionsOptionName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMgrModuleGet200ResponseInnerOptionsOptionName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["default_value"] = o.DefaultValue
	toSerialize["desc"] = o.Desc
	toSerialize["enum_allowed"] = o.EnumAllowed
	toSerialize["flags"] = o.Flags
	toSerialize["level"] = o.Level
	toSerialize["long_desc"] = o.LongDesc
	toSerialize["max"] = o.Max
	toSerialize["min"] = o.Min
	toSerialize["name"] = o.Name
	toSerialize["see_also"] = o.SeeAlso
	toSerialize["tags"] = o.Tags
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ApiMgrModuleGet200ResponseInnerOptionsOptionName) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"default_value",
		"desc",
		"enum_allowed",
		"flags",
		"level",
		"long_desc",
		"max",
		"min",
		"name",
		"see_also",
		"tags",
		"type",
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

	varApiMgrModuleGet200ResponseInnerOptionsOptionName := _ApiMgrModuleGet200ResponseInnerOptionsOptionName{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiMgrModuleGet200ResponseInnerOptionsOptionName)

	if err != nil {
		return err
	}

	*o = ApiMgrModuleGet200ResponseInnerOptionsOptionName(varApiMgrModuleGet200ResponseInnerOptionsOptionName)

	return err
}

type NullableApiMgrModuleGet200ResponseInnerOptionsOptionName struct {
	value *ApiMgrModuleGet200ResponseInnerOptionsOptionName
	isSet bool
}

func (v NullableApiMgrModuleGet200ResponseInnerOptionsOptionName) Get() *ApiMgrModuleGet200ResponseInnerOptionsOptionName {
	return v.value
}

func (v *NullableApiMgrModuleGet200ResponseInnerOptionsOptionName) Set(val *ApiMgrModuleGet200ResponseInnerOptionsOptionName) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMgrModuleGet200ResponseInnerOptionsOptionName) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMgrModuleGet200ResponseInnerOptionsOptionName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMgrModuleGet200ResponseInnerOptionsOptionName(val *ApiMgrModuleGet200ResponseInnerOptionsOptionName) *NullableApiMgrModuleGet200ResponseInnerOptionsOptionName {
	return &NullableApiMgrModuleGet200ResponseInnerOptionsOptionName{value: val, isSet: true}
}

func (v NullableApiMgrModuleGet200ResponseInnerOptionsOptionName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMgrModuleGet200ResponseInnerOptionsOptionName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


