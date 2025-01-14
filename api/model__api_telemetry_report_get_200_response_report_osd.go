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

// checks if the ApiTelemetryReportGet200ResponseReportOsd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportOsd{}

// ApiTelemetryReportGet200ResponseReportOsd 
type ApiTelemetryReportGet200ResponseReportOsd struct {
	// 
	ClusterNetwork bool `json:"cluster_network"`
	// 
	Count int32 `json:"count"`
	// 
	RequireMinCompatClient string `json:"require_min_compat_client"`
	// 
	RequireOsdRelease string `json:"require_osd_release"`
}

type _ApiTelemetryReportGet200ResponseReportOsd ApiTelemetryReportGet200ResponseReportOsd

// NewApiTelemetryReportGet200ResponseReportOsd instantiates a new ApiTelemetryReportGet200ResponseReportOsd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportOsd(clusterNetwork bool, count int32, requireMinCompatClient string, requireOsdRelease string) *ApiTelemetryReportGet200ResponseReportOsd {
	this := ApiTelemetryReportGet200ResponseReportOsd{}
	this.ClusterNetwork = clusterNetwork
	this.Count = count
	this.RequireMinCompatClient = requireMinCompatClient
	this.RequireOsdRelease = requireOsdRelease
	return &this
}

// NewApiTelemetryReportGet200ResponseReportOsdWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportOsd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportOsdWithDefaults() *ApiTelemetryReportGet200ResponseReportOsd {
	this := ApiTelemetryReportGet200ResponseReportOsd{}
	return &this
}

// GetClusterNetwork returns the ClusterNetwork field value
func (o *ApiTelemetryReportGet200ResponseReportOsd) GetClusterNetwork() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ClusterNetwork
}

// GetClusterNetworkOk returns a tuple with the ClusterNetwork field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportOsd) GetClusterNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterNetwork, true
}

// SetClusterNetwork sets field value
func (o *ApiTelemetryReportGet200ResponseReportOsd) SetClusterNetwork(v bool) {
	o.ClusterNetwork = v
}

// GetCount returns the Count field value
func (o *ApiTelemetryReportGet200ResponseReportOsd) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportOsd) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ApiTelemetryReportGet200ResponseReportOsd) SetCount(v int32) {
	o.Count = v
}

// GetRequireMinCompatClient returns the RequireMinCompatClient field value
func (o *ApiTelemetryReportGet200ResponseReportOsd) GetRequireMinCompatClient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequireMinCompatClient
}

// GetRequireMinCompatClientOk returns a tuple with the RequireMinCompatClient field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportOsd) GetRequireMinCompatClientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireMinCompatClient, true
}

// SetRequireMinCompatClient sets field value
func (o *ApiTelemetryReportGet200ResponseReportOsd) SetRequireMinCompatClient(v string) {
	o.RequireMinCompatClient = v
}

// GetRequireOsdRelease returns the RequireOsdRelease field value
func (o *ApiTelemetryReportGet200ResponseReportOsd) GetRequireOsdRelease() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequireOsdRelease
}

// GetRequireOsdReleaseOk returns a tuple with the RequireOsdRelease field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportOsd) GetRequireOsdReleaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireOsdRelease, true
}

// SetRequireOsdRelease sets field value
func (o *ApiTelemetryReportGet200ResponseReportOsd) SetRequireOsdRelease(v string) {
	o.RequireOsdRelease = v
}

func (o ApiTelemetryReportGet200ResponseReportOsd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportOsd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster_network"] = o.ClusterNetwork
	toSerialize["count"] = o.Count
	toSerialize["require_min_compat_client"] = o.RequireMinCompatClient
	toSerialize["require_osd_release"] = o.RequireOsdRelease
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportOsd) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cluster_network",
		"count",
		"require_min_compat_client",
		"require_osd_release",
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

	varApiTelemetryReportGet200ResponseReportOsd := _ApiTelemetryReportGet200ResponseReportOsd{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportOsd)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportOsd(varApiTelemetryReportGet200ResponseReportOsd)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportOsd struct {
	value *ApiTelemetryReportGet200ResponseReportOsd
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportOsd) Get() *ApiTelemetryReportGet200ResponseReportOsd {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportOsd) Set(val *ApiTelemetryReportGet200ResponseReportOsd) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportOsd) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportOsd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportOsd(val *ApiTelemetryReportGet200ResponseReportOsd) *NullableApiTelemetryReportGet200ResponseReportOsd {
	return &NullableApiTelemetryReportGet200ResponseReportOsd{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportOsd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportOsd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


