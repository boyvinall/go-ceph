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

// checks if the ApiTelemetryReportGet200ResponseReportCrush type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportCrush{}

// ApiTelemetryReportGet200ResponseReportCrush 
type ApiTelemetryReportGet200ResponseReportCrush struct {
	BucketAlgs ApiTelemetryReportGet200ResponseReportCrushBucketAlgs `json:"bucket_algs"`
	BucketSizes ApiTelemetryReportGet200ResponseReportCrushBucketSizes `json:"bucket_sizes"`
	BucketTypes ApiTelemetryReportGet200ResponseReportCrushBucketTypes `json:"bucket_types"`
	// 
	CompatWeightSet bool `json:"compat_weight_set"`
	// 
	DeviceClasses []int32 `json:"device_classes"`
	// 
	NumBuckets int32 `json:"num_buckets"`
	// 
	NumDevices int32 `json:"num_devices"`
	// 
	NumRules int32 `json:"num_rules"`
	// 
	NumTypes int32 `json:"num_types"`
	// 
	NumWeightSets int32 `json:"num_weight_sets"`
	Tunables ApiTelemetryReportGet200ResponseReportCrushTunables `json:"tunables"`
}

type _ApiTelemetryReportGet200ResponseReportCrush ApiTelemetryReportGet200ResponseReportCrush

// NewApiTelemetryReportGet200ResponseReportCrush instantiates a new ApiTelemetryReportGet200ResponseReportCrush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportCrush(bucketAlgs ApiTelemetryReportGet200ResponseReportCrushBucketAlgs, bucketSizes ApiTelemetryReportGet200ResponseReportCrushBucketSizes, bucketTypes ApiTelemetryReportGet200ResponseReportCrushBucketTypes, compatWeightSet bool, deviceClasses []int32, numBuckets int32, numDevices int32, numRules int32, numTypes int32, numWeightSets int32, tunables ApiTelemetryReportGet200ResponseReportCrushTunables) *ApiTelemetryReportGet200ResponseReportCrush {
	this := ApiTelemetryReportGet200ResponseReportCrush{}
	this.BucketAlgs = bucketAlgs
	this.BucketSizes = bucketSizes
	this.BucketTypes = bucketTypes
	this.CompatWeightSet = compatWeightSet
	this.DeviceClasses = deviceClasses
	this.NumBuckets = numBuckets
	this.NumDevices = numDevices
	this.NumRules = numRules
	this.NumTypes = numTypes
	this.NumWeightSets = numWeightSets
	this.Tunables = tunables
	return &this
}

// NewApiTelemetryReportGet200ResponseReportCrushWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportCrush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportCrushWithDefaults() *ApiTelemetryReportGet200ResponseReportCrush {
	this := ApiTelemetryReportGet200ResponseReportCrush{}
	return &this
}

// GetBucketAlgs returns the BucketAlgs field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketAlgs() ApiTelemetryReportGet200ResponseReportCrushBucketAlgs {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportCrushBucketAlgs
		return ret
	}

	return o.BucketAlgs
}

// GetBucketAlgsOk returns a tuple with the BucketAlgs field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketAlgsOk() (*ApiTelemetryReportGet200ResponseReportCrushBucketAlgs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketAlgs, true
}

// SetBucketAlgs sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetBucketAlgs(v ApiTelemetryReportGet200ResponseReportCrushBucketAlgs) {
	o.BucketAlgs = v
}

// GetBucketSizes returns the BucketSizes field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketSizes() ApiTelemetryReportGet200ResponseReportCrushBucketSizes {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportCrushBucketSizes
		return ret
	}

	return o.BucketSizes
}

// GetBucketSizesOk returns a tuple with the BucketSizes field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketSizesOk() (*ApiTelemetryReportGet200ResponseReportCrushBucketSizes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketSizes, true
}

// SetBucketSizes sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetBucketSizes(v ApiTelemetryReportGet200ResponseReportCrushBucketSizes) {
	o.BucketSizes = v
}

// GetBucketTypes returns the BucketTypes field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketTypes() ApiTelemetryReportGet200ResponseReportCrushBucketTypes {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportCrushBucketTypes
		return ret
	}

	return o.BucketTypes
}

// GetBucketTypesOk returns a tuple with the BucketTypes field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetBucketTypesOk() (*ApiTelemetryReportGet200ResponseReportCrushBucketTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketTypes, true
}

// SetBucketTypes sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetBucketTypes(v ApiTelemetryReportGet200ResponseReportCrushBucketTypes) {
	o.BucketTypes = v
}

// GetCompatWeightSet returns the CompatWeightSet field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetCompatWeightSet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CompatWeightSet
}

// GetCompatWeightSetOk returns a tuple with the CompatWeightSet field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetCompatWeightSetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompatWeightSet, true
}

// SetCompatWeightSet sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetCompatWeightSet(v bool) {
	o.CompatWeightSet = v
}

// GetDeviceClasses returns the DeviceClasses field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetDeviceClasses() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.DeviceClasses
}

// GetDeviceClassesOk returns a tuple with the DeviceClasses field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetDeviceClassesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceClasses, true
}

// SetDeviceClasses sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetDeviceClasses(v []int32) {
	o.DeviceClasses = v
}

// GetNumBuckets returns the NumBuckets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumBuckets() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumBuckets
}

// GetNumBucketsOk returns a tuple with the NumBuckets field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumBucketsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumBuckets, true
}

// SetNumBuckets sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetNumBuckets(v int32) {
	o.NumBuckets = v
}

// GetNumDevices returns the NumDevices field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumDevices() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumDevices
}

// GetNumDevicesOk returns a tuple with the NumDevices field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumDevicesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumDevices, true
}

// SetNumDevices sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetNumDevices(v int32) {
	o.NumDevices = v
}

// GetNumRules returns the NumRules field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumRules() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumRules
}

// GetNumRulesOk returns a tuple with the NumRules field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumRulesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumRules, true
}

// SetNumRules sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetNumRules(v int32) {
	o.NumRules = v
}

// GetNumTypes returns the NumTypes field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumTypes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumTypes
}

// GetNumTypesOk returns a tuple with the NumTypes field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumTypesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumTypes, true
}

// SetNumTypes sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetNumTypes(v int32) {
	o.NumTypes = v
}

// GetNumWeightSets returns the NumWeightSets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumWeightSets() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumWeightSets
}

// GetNumWeightSetsOk returns a tuple with the NumWeightSets field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetNumWeightSetsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumWeightSets, true
}

// SetNumWeightSets sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetNumWeightSets(v int32) {
	o.NumWeightSets = v
}

// GetTunables returns the Tunables field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetTunables() ApiTelemetryReportGet200ResponseReportCrushTunables {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportCrushTunables
		return ret
	}

	return o.Tunables
}

// GetTunablesOk returns a tuple with the Tunables field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrush) GetTunablesOk() (*ApiTelemetryReportGet200ResponseReportCrushTunables, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tunables, true
}

// SetTunables sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrush) SetTunables(v ApiTelemetryReportGet200ResponseReportCrushTunables) {
	o.Tunables = v
}

func (o ApiTelemetryReportGet200ResponseReportCrush) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportCrush) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket_algs"] = o.BucketAlgs
	toSerialize["bucket_sizes"] = o.BucketSizes
	toSerialize["bucket_types"] = o.BucketTypes
	toSerialize["compat_weight_set"] = o.CompatWeightSet
	toSerialize["device_classes"] = o.DeviceClasses
	toSerialize["num_buckets"] = o.NumBuckets
	toSerialize["num_devices"] = o.NumDevices
	toSerialize["num_rules"] = o.NumRules
	toSerialize["num_types"] = o.NumTypes
	toSerialize["num_weight_sets"] = o.NumWeightSets
	toSerialize["tunables"] = o.Tunables
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportCrush) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bucket_algs",
		"bucket_sizes",
		"bucket_types",
		"compat_weight_set",
		"device_classes",
		"num_buckets",
		"num_devices",
		"num_rules",
		"num_types",
		"num_weight_sets",
		"tunables",
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

	varApiTelemetryReportGet200ResponseReportCrush := _ApiTelemetryReportGet200ResponseReportCrush{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportCrush)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportCrush(varApiTelemetryReportGet200ResponseReportCrush)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportCrush struct {
	value *ApiTelemetryReportGet200ResponseReportCrush
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportCrush) Get() *ApiTelemetryReportGet200ResponseReportCrush {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrush) Set(val *ApiTelemetryReportGet200ResponseReportCrush) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportCrush) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportCrush(val *ApiTelemetryReportGet200ResponseReportCrush) *NullableApiTelemetryReportGet200ResponseReportCrush {
	return &NullableApiTelemetryReportGet200ResponseReportCrush{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportCrush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

