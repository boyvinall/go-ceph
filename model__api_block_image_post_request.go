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

// checks if the ApiBlockImagePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBlockImagePostRequest{}

// ApiBlockImagePostRequest struct for ApiBlockImagePostRequest
type ApiBlockImagePostRequest struct {
	Configuration *string `json:"configuration,omitempty"`
	DataPool *string `json:"data_pool,omitempty"`
	Features *string `json:"features,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
	MirrorMode *string `json:"mirror_mode,omitempty"`
	Name string `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
	ObjSize *int32 `json:"obj_size,omitempty"`
	PoolName string `json:"pool_name"`
	ScheduleInterval *string `json:"schedule_interval,omitempty"`
	Size int32 `json:"size"`
	StripeCount *int32 `json:"stripe_count,omitempty"`
	StripeUnit *string `json:"stripe_unit,omitempty"`
}

type _ApiBlockImagePostRequest ApiBlockImagePostRequest

// NewApiBlockImagePostRequest instantiates a new ApiBlockImagePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBlockImagePostRequest(name string, poolName string, size int32) *ApiBlockImagePostRequest {
	this := ApiBlockImagePostRequest{}
	this.Name = name
	this.PoolName = poolName
	var scheduleInterval string = ""
	this.ScheduleInterval = &scheduleInterval
	this.Size = size
	return &this
}

// NewApiBlockImagePostRequestWithDefaults instantiates a new ApiBlockImagePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBlockImagePostRequestWithDefaults() *ApiBlockImagePostRequest {
	this := ApiBlockImagePostRequest{}
	var scheduleInterval string = ""
	this.ScheduleInterval = &scheduleInterval
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ApiBlockImagePostRequest) GetConfiguration() string {
	if o == nil || IsNil(o.Configuration) {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetConfigurationOk() (*string, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ApiBlockImagePostRequest) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *ApiBlockImagePostRequest) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDataPool returns the DataPool field value if set, zero value otherwise.
func (o *ApiBlockImagePostRequest) GetDataPool() string {
	if o == nil || IsNil(o.DataPool) {
		var ret string
		return ret
	}
	return *o.DataPool
}

// GetDataPoolOk returns a tuple with the DataPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetDataPoolOk() (*string, bool) {
	if o == nil || IsNil(o.DataPool) {
		return nil, false
	}
	return o.DataPool, true
}

// HasDataPool returns a boolean if a field has been set.
func (o *ApiBlockImagePostRequest) HasDataPool() bool {
	if o != nil && !IsNil(o.DataPool) {
		return true
	}

	return false
}

// SetDataPool gets a reference to the given string and assigns it to the DataPool field.
func (o *ApiBlockImagePostRequest) SetDataPool(v string) {
	o.DataPool = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ApiBlockImagePostRequest) GetFeatures() string {
	if o == nil || IsNil(o.Features) {
		var ret string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ApiBlockImagePostRequest) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given string and assigns it to the Features field.
func (o *ApiBlockImagePostRequest) SetFeatures(v string) {
	o.Features = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApiBlockImagePostRequest) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApiBlockImagePostRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *ApiBlockImagePostRequest) SetMetadata(v string) {
	o.Metadata = &v
}

// GetMirrorMode returns the MirrorMode field value if set, zero value otherwise.
func (o *ApiBlockImagePostRequest) GetMirrorMode() string {
	if o == nil || IsNil(o.MirrorMode) {
		var ret string
		return ret
	}
	return *o.MirrorMode
}

// GetMirrorModeOk returns a tuple with the MirrorMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetMirrorModeOk() (*string, bool) {
	if o == nil || IsNil(o.MirrorMode) {
		return nil, false
	}
	return o.MirrorMode, true
}

// HasMirrorMode returns a boolean if a field has been set.
func (o *ApiBlockImagePostRequest) HasMirrorMode() bool {
	if o != nil && !IsNil(o.MirrorMode) {
		return true
	}

	return false
}

// SetMirrorMode gets a reference to the given string and assigns it to the MirrorMode field.
func (o *ApiBlockImagePostRequest) SetMirrorMode(v string) {
	o.MirrorMode = &v
}

// GetName returns the Name field value
func (o *ApiBlockImagePostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiBlockImagePostRequest) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ApiBlockImagePostRequest) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ApiBlockImagePostRequest) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ApiBlockImagePostRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetObjSize returns the ObjSize field value if set, zero value otherwise.
func (o *ApiBlockImagePostRequest) GetObjSize() int32 {
	if o == nil || IsNil(o.ObjSize) {
		var ret int32
		return ret
	}
	return *o.ObjSize
}

// GetObjSizeOk returns a tuple with the ObjSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetObjSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ObjSize) {
		return nil, false
	}
	return o.ObjSize, true
}

// HasObjSize returns a boolean if a field has been set.
func (o *ApiBlockImagePostRequest) HasObjSize() bool {
	if o != nil && !IsNil(o.ObjSize) {
		return true
	}

	return false
}

// SetObjSize gets a reference to the given int32 and assigns it to the ObjSize field.
func (o *ApiBlockImagePostRequest) SetObjSize(v int32) {
	o.ObjSize = &v
}

// GetPoolName returns the PoolName field value
func (o *ApiBlockImagePostRequest) GetPoolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetPoolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolName, true
}

// SetPoolName sets field value
func (o *ApiBlockImagePostRequest) SetPoolName(v string) {
	o.PoolName = v
}

// GetScheduleInterval returns the ScheduleInterval field value if set, zero value otherwise.
func (o *ApiBlockImagePostRequest) GetScheduleInterval() string {
	if o == nil || IsNil(o.ScheduleInterval) {
		var ret string
		return ret
	}
	return *o.ScheduleInterval
}

// GetScheduleIntervalOk returns a tuple with the ScheduleInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetScheduleIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduleInterval) {
		return nil, false
	}
	return o.ScheduleInterval, true
}

// HasScheduleInterval returns a boolean if a field has been set.
func (o *ApiBlockImagePostRequest) HasScheduleInterval() bool {
	if o != nil && !IsNil(o.ScheduleInterval) {
		return true
	}

	return false
}

// SetScheduleInterval gets a reference to the given string and assigns it to the ScheduleInterval field.
func (o *ApiBlockImagePostRequest) SetScheduleInterval(v string) {
	o.ScheduleInterval = &v
}

// GetSize returns the Size field value
func (o *ApiBlockImagePostRequest) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ApiBlockImagePostRequest) SetSize(v int32) {
	o.Size = v
}

// GetStripeCount returns the StripeCount field value if set, zero value otherwise.
func (o *ApiBlockImagePostRequest) GetStripeCount() int32 {
	if o == nil || IsNil(o.StripeCount) {
		var ret int32
		return ret
	}
	return *o.StripeCount
}

// GetStripeCountOk returns a tuple with the StripeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetStripeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.StripeCount) {
		return nil, false
	}
	return o.StripeCount, true
}

// HasStripeCount returns a boolean if a field has been set.
func (o *ApiBlockImagePostRequest) HasStripeCount() bool {
	if o != nil && !IsNil(o.StripeCount) {
		return true
	}

	return false
}

// SetStripeCount gets a reference to the given int32 and assigns it to the StripeCount field.
func (o *ApiBlockImagePostRequest) SetStripeCount(v int32) {
	o.StripeCount = &v
}

// GetStripeUnit returns the StripeUnit field value if set, zero value otherwise.
func (o *ApiBlockImagePostRequest) GetStripeUnit() string {
	if o == nil || IsNil(o.StripeUnit) {
		var ret string
		return ret
	}
	return *o.StripeUnit
}

// GetStripeUnitOk returns a tuple with the StripeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImagePostRequest) GetStripeUnitOk() (*string, bool) {
	if o == nil || IsNil(o.StripeUnit) {
		return nil, false
	}
	return o.StripeUnit, true
}

// HasStripeUnit returns a boolean if a field has been set.
func (o *ApiBlockImagePostRequest) HasStripeUnit() bool {
	if o != nil && !IsNil(o.StripeUnit) {
		return true
	}

	return false
}

// SetStripeUnit gets a reference to the given string and assigns it to the StripeUnit field.
func (o *ApiBlockImagePostRequest) SetStripeUnit(v string) {
	o.StripeUnit = &v
}

func (o ApiBlockImagePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBlockImagePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.DataPool) {
		toSerialize["data_pool"] = o.DataPool
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MirrorMode) {
		toSerialize["mirror_mode"] = o.MirrorMode
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.ObjSize) {
		toSerialize["obj_size"] = o.ObjSize
	}
	toSerialize["pool_name"] = o.PoolName
	if !IsNil(o.ScheduleInterval) {
		toSerialize["schedule_interval"] = o.ScheduleInterval
	}
	toSerialize["size"] = o.Size
	if !IsNil(o.StripeCount) {
		toSerialize["stripe_count"] = o.StripeCount
	}
	if !IsNil(o.StripeUnit) {
		toSerialize["stripe_unit"] = o.StripeUnit
	}
	return toSerialize, nil
}

func (o *ApiBlockImagePostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"pool_name",
		"size",
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

	varApiBlockImagePostRequest := _ApiBlockImagePostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiBlockImagePostRequest)

	if err != nil {
		return err
	}

	*o = ApiBlockImagePostRequest(varApiBlockImagePostRequest)

	return err
}

type NullableApiBlockImagePostRequest struct {
	value *ApiBlockImagePostRequest
	isSet bool
}

func (v NullableApiBlockImagePostRequest) Get() *ApiBlockImagePostRequest {
	return v.value
}

func (v *NullableApiBlockImagePostRequest) Set(val *ApiBlockImagePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBlockImagePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBlockImagePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBlockImagePostRequest(val *ApiBlockImagePostRequest) *NullableApiBlockImagePostRequest {
	return &NullableApiBlockImagePostRequest{value: val, isSet: true}
}

func (v NullableApiBlockImagePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBlockImagePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


