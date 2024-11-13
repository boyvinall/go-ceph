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

// checks if the ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest{}

// ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest struct for ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest
type ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest struct {
	ChildImageName string `json:"child_image_name"`
	ChildNamespace *string `json:"child_namespace,omitempty"`
	ChildPoolName string `json:"child_pool_name"`
	Configuration *string `json:"configuration,omitempty"`
	DataPool *string `json:"data_pool,omitempty"`
	Features *string `json:"features,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
	ObjSize *int32 `json:"obj_size,omitempty"`
	StripeCount *int32 `json:"stripe_count,omitempty"`
	StripeUnit *string `json:"stripe_unit,omitempty"`
}

type _ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest

// NewApiBlockImageImageSpecSnapSnapshotNameClonePostRequest instantiates a new ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBlockImageImageSpecSnapSnapshotNameClonePostRequest(childImageName string, childPoolName string) *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest {
	this := ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest{}
	this.ChildImageName = childImageName
	this.ChildPoolName = childPoolName
	return &this
}

// NewApiBlockImageImageSpecSnapSnapshotNameClonePostRequestWithDefaults instantiates a new ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBlockImageImageSpecSnapSnapshotNameClonePostRequestWithDefaults() *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest {
	this := ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest{}
	return &this
}

// GetChildImageName returns the ChildImageName field value
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetChildImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChildImageName
}

// GetChildImageNameOk returns a tuple with the ChildImageName field value
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetChildImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChildImageName, true
}

// SetChildImageName sets field value
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) SetChildImageName(v string) {
	o.ChildImageName = v
}

// GetChildNamespace returns the ChildNamespace field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetChildNamespace() string {
	if o == nil || IsNil(o.ChildNamespace) {
		var ret string
		return ret
	}
	return *o.ChildNamespace
}

// GetChildNamespaceOk returns a tuple with the ChildNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetChildNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.ChildNamespace) {
		return nil, false
	}
	return o.ChildNamespace, true
}

// HasChildNamespace returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) HasChildNamespace() bool {
	if o != nil && !IsNil(o.ChildNamespace) {
		return true
	}

	return false
}

// SetChildNamespace gets a reference to the given string and assigns it to the ChildNamespace field.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) SetChildNamespace(v string) {
	o.ChildNamespace = &v
}

// GetChildPoolName returns the ChildPoolName field value
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetChildPoolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChildPoolName
}

// GetChildPoolNameOk returns a tuple with the ChildPoolName field value
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetChildPoolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChildPoolName, true
}

// SetChildPoolName sets field value
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) SetChildPoolName(v string) {
	o.ChildPoolName = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetConfiguration() string {
	if o == nil || IsNil(o.Configuration) {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetConfigurationOk() (*string, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDataPool returns the DataPool field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetDataPool() string {
	if o == nil || IsNil(o.DataPool) {
		var ret string
		return ret
	}
	return *o.DataPool
}

// GetDataPoolOk returns a tuple with the DataPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetDataPoolOk() (*string, bool) {
	if o == nil || IsNil(o.DataPool) {
		return nil, false
	}
	return o.DataPool, true
}

// HasDataPool returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) HasDataPool() bool {
	if o != nil && !IsNil(o.DataPool) {
		return true
	}

	return false
}

// SetDataPool gets a reference to the given string and assigns it to the DataPool field.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) SetDataPool(v string) {
	o.DataPool = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetFeatures() string {
	if o == nil || IsNil(o.Features) {
		var ret string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given string and assigns it to the Features field.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) SetFeatures(v string) {
	o.Features = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) SetMetadata(v string) {
	o.Metadata = &v
}

// GetObjSize returns the ObjSize field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetObjSize() int32 {
	if o == nil || IsNil(o.ObjSize) {
		var ret int32
		return ret
	}
	return *o.ObjSize
}

// GetObjSizeOk returns a tuple with the ObjSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetObjSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ObjSize) {
		return nil, false
	}
	return o.ObjSize, true
}

// HasObjSize returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) HasObjSize() bool {
	if o != nil && !IsNil(o.ObjSize) {
		return true
	}

	return false
}

// SetObjSize gets a reference to the given int32 and assigns it to the ObjSize field.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) SetObjSize(v int32) {
	o.ObjSize = &v
}

// GetStripeCount returns the StripeCount field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetStripeCount() int32 {
	if o == nil || IsNil(o.StripeCount) {
		var ret int32
		return ret
	}
	return *o.StripeCount
}

// GetStripeCountOk returns a tuple with the StripeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetStripeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.StripeCount) {
		return nil, false
	}
	return o.StripeCount, true
}

// HasStripeCount returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) HasStripeCount() bool {
	if o != nil && !IsNil(o.StripeCount) {
		return true
	}

	return false
}

// SetStripeCount gets a reference to the given int32 and assigns it to the StripeCount field.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) SetStripeCount(v int32) {
	o.StripeCount = &v
}

// GetStripeUnit returns the StripeUnit field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetStripeUnit() string {
	if o == nil || IsNil(o.StripeUnit) {
		var ret string
		return ret
	}
	return *o.StripeUnit
}

// GetStripeUnitOk returns a tuple with the StripeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) GetStripeUnitOk() (*string, bool) {
	if o == nil || IsNil(o.StripeUnit) {
		return nil, false
	}
	return o.StripeUnit, true
}

// HasStripeUnit returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) HasStripeUnit() bool {
	if o != nil && !IsNil(o.StripeUnit) {
		return true
	}

	return false
}

// SetStripeUnit gets a reference to the given string and assigns it to the StripeUnit field.
func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) SetStripeUnit(v string) {
	o.StripeUnit = &v
}

func (o ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["child_image_name"] = o.ChildImageName
	if !IsNil(o.ChildNamespace) {
		toSerialize["child_namespace"] = o.ChildNamespace
	}
	toSerialize["child_pool_name"] = o.ChildPoolName
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
	if !IsNil(o.ObjSize) {
		toSerialize["obj_size"] = o.ObjSize
	}
	if !IsNil(o.StripeCount) {
		toSerialize["stripe_count"] = o.StripeCount
	}
	if !IsNil(o.StripeUnit) {
		toSerialize["stripe_unit"] = o.StripeUnit
	}
	return toSerialize, nil
}

func (o *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"child_image_name",
		"child_pool_name",
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

	varApiBlockImageImageSpecSnapSnapshotNameClonePostRequest := _ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiBlockImageImageSpecSnapSnapshotNameClonePostRequest)

	if err != nil {
		return err
	}

	*o = ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest(varApiBlockImageImageSpecSnapSnapshotNameClonePostRequest)

	return err
}

type NullableApiBlockImageImageSpecSnapSnapshotNameClonePostRequest struct {
	value *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest
	isSet bool
}

func (v NullableApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) Get() *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest {
	return v.value
}

func (v *NullableApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) Set(val *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBlockImageImageSpecSnapSnapshotNameClonePostRequest(val *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) *NullableApiBlockImageImageSpecSnapSnapshotNameClonePostRequest {
	return &NullableApiBlockImageImageSpecSnapSnapshotNameClonePostRequest{value: val, isSet: true}
}

func (v NullableApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


