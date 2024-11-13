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

// checks if the ApiTaskGet200ResponseFinishedTasksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTaskGet200ResponseFinishedTasksInner{}

// ApiTaskGet200ResponseFinishedTasksInner struct for ApiTaskGet200ResponseFinishedTasksInner
type ApiTaskGet200ResponseFinishedTasksInner struct {
	// Task begin time
	BeginTime string `json:"begin_time"`
	// 
	Duration int32 `json:"duration"`
	// Task end time
	EndTime string `json:"end_time"`
	// 
	Exception bool `json:"exception"`
	Metadata ApiSummaryGet200ResponseFinishedTasksInnerMetadata `json:"metadata"`
	// finished tasks name
	Name string `json:"name"`
	// Progress of tasks
	Progress int32 `json:"progress"`
	// 
	RetValue bool `json:"ret_value"`
	// 
	Success bool `json:"success"`
}

type _ApiTaskGet200ResponseFinishedTasksInner ApiTaskGet200ResponseFinishedTasksInner

// NewApiTaskGet200ResponseFinishedTasksInner instantiates a new ApiTaskGet200ResponseFinishedTasksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTaskGet200ResponseFinishedTasksInner(beginTime string, duration int32, endTime string, exception bool, metadata ApiSummaryGet200ResponseFinishedTasksInnerMetadata, name string, progress int32, retValue bool, success bool) *ApiTaskGet200ResponseFinishedTasksInner {
	this := ApiTaskGet200ResponseFinishedTasksInner{}
	this.BeginTime = beginTime
	this.Duration = duration
	this.EndTime = endTime
	this.Exception = exception
	this.Metadata = metadata
	this.Name = name
	this.Progress = progress
	this.RetValue = retValue
	this.Success = success
	return &this
}

// NewApiTaskGet200ResponseFinishedTasksInnerWithDefaults instantiates a new ApiTaskGet200ResponseFinishedTasksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTaskGet200ResponseFinishedTasksInnerWithDefaults() *ApiTaskGet200ResponseFinishedTasksInner {
	this := ApiTaskGet200ResponseFinishedTasksInner{}
	return &this
}

// GetBeginTime returns the BeginTime field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetBeginTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeginTime
}

// GetBeginTimeOk returns a tuple with the BeginTime field value
// and a boolean to check if the value has been set.
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetBeginTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeginTime, true
}

// SetBeginTime sets field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) SetBeginTime(v string) {
	o.BeginTime = v
}

// GetDuration returns the Duration field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) SetDuration(v int32) {
	o.Duration = v
}

// GetEndTime returns the EndTime field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) SetEndTime(v string) {
	o.EndTime = v
}

// GetException returns the Exception field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetException() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value
// and a boolean to check if the value has been set.
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetExceptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exception, true
}

// SetException sets field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) SetException(v bool) {
	o.Exception = v
}

// GetMetadata returns the Metadata field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetMetadata() ApiSummaryGet200ResponseFinishedTasksInnerMetadata {
	if o == nil {
		var ret ApiSummaryGet200ResponseFinishedTasksInnerMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetMetadataOk() (*ApiSummaryGet200ResponseFinishedTasksInnerMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) SetMetadata(v ApiSummaryGet200ResponseFinishedTasksInnerMetadata) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) SetName(v string) {
	o.Name = v
}

// GetProgress returns the Progress field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetProgress() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetProgressOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) SetProgress(v int32) {
	o.Progress = v
}

// GetRetValue returns the RetValue field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetRetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RetValue
}

// GetRetValueOk returns a tuple with the RetValue field value
// and a boolean to check if the value has been set.
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetRetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetValue, true
}

// SetRetValue sets field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) SetRetValue(v bool) {
	o.RetValue = v
}

// GetSuccess returns the Success field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ApiTaskGet200ResponseFinishedTasksInner) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ApiTaskGet200ResponseFinishedTasksInner) SetSuccess(v bool) {
	o.Success = v
}

func (o ApiTaskGet200ResponseFinishedTasksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTaskGet200ResponseFinishedTasksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["begin_time"] = o.BeginTime
	toSerialize["duration"] = o.Duration
	toSerialize["end_time"] = o.EndTime
	toSerialize["exception"] = o.Exception
	toSerialize["metadata"] = o.Metadata
	toSerialize["name"] = o.Name
	toSerialize["progress"] = o.Progress
	toSerialize["ret_value"] = o.RetValue
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *ApiTaskGet200ResponseFinishedTasksInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"begin_time",
		"duration",
		"end_time",
		"exception",
		"metadata",
		"name",
		"progress",
		"ret_value",
		"success",
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

	varApiTaskGet200ResponseFinishedTasksInner := _ApiTaskGet200ResponseFinishedTasksInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTaskGet200ResponseFinishedTasksInner)

	if err != nil {
		return err
	}

	*o = ApiTaskGet200ResponseFinishedTasksInner(varApiTaskGet200ResponseFinishedTasksInner)

	return err
}

type NullableApiTaskGet200ResponseFinishedTasksInner struct {
	value *ApiTaskGet200ResponseFinishedTasksInner
	isSet bool
}

func (v NullableApiTaskGet200ResponseFinishedTasksInner) Get() *ApiTaskGet200ResponseFinishedTasksInner {
	return v.value
}

func (v *NullableApiTaskGet200ResponseFinishedTasksInner) Set(val *ApiTaskGet200ResponseFinishedTasksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTaskGet200ResponseFinishedTasksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTaskGet200ResponseFinishedTasksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTaskGet200ResponseFinishedTasksInner(val *ApiTaskGet200ResponseFinishedTasksInner) *NullableApiTaskGet200ResponseFinishedTasksInner {
	return &NullableApiTaskGet200ResponseFinishedTasksInner{value: val, isSet: true}
}

func (v NullableApiTaskGet200ResponseFinishedTasksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTaskGet200ResponseFinishedTasksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

