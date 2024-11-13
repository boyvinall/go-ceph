/*
Ceph REST API

This is the official Ceph REST API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ceph

import (
	"encoding/json"
)

// checks if the ApiBlockImageImageSpecPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBlockImageImageSpecPutRequest{}

// ApiBlockImageImageSpecPutRequest struct for ApiBlockImageImageSpecPutRequest
type ApiBlockImageImageSpecPutRequest struct {
	Configuration *string `json:"configuration,omitempty"`
	EnableMirror *string `json:"enable_mirror,omitempty"`
	Features *string `json:"features,omitempty"`
	Force *bool `json:"force,omitempty"`
	ImageMirrorMode *string `json:"image_mirror_mode,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
	MirrorMode *string `json:"mirror_mode,omitempty"`
	Name *string `json:"name,omitempty"`
	Primary *string `json:"primary,omitempty"`
	RemoveScheduling *bool `json:"remove_scheduling,omitempty"`
	Resync *bool `json:"resync,omitempty"`
	ScheduleInterval *string `json:"schedule_interval,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewApiBlockImageImageSpecPutRequest instantiates a new ApiBlockImageImageSpecPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBlockImageImageSpecPutRequest() *ApiBlockImageImageSpecPutRequest {
	this := ApiBlockImageImageSpecPutRequest{}
	var force bool = false
	this.Force = &force
	var removeScheduling bool = false
	this.RemoveScheduling = &removeScheduling
	var resync bool = false
	this.Resync = &resync
	var scheduleInterval string = ""
	this.ScheduleInterval = &scheduleInterval
	return &this
}

// NewApiBlockImageImageSpecPutRequestWithDefaults instantiates a new ApiBlockImageImageSpecPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBlockImageImageSpecPutRequestWithDefaults() *ApiBlockImageImageSpecPutRequest {
	this := ApiBlockImageImageSpecPutRequest{}
	var force bool = false
	this.Force = &force
	var removeScheduling bool = false
	this.RemoveScheduling = &removeScheduling
	var resync bool = false
	this.Resync = &resync
	var scheduleInterval string = ""
	this.ScheduleInterval = &scheduleInterval
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetConfiguration() string {
	if o == nil || IsNil(o.Configuration) {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetConfigurationOk() (*string, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *ApiBlockImageImageSpecPutRequest) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetEnableMirror returns the EnableMirror field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetEnableMirror() string {
	if o == nil || IsNil(o.EnableMirror) {
		var ret string
		return ret
	}
	return *o.EnableMirror
}

// GetEnableMirrorOk returns a tuple with the EnableMirror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetEnableMirrorOk() (*string, bool) {
	if o == nil || IsNil(o.EnableMirror) {
		return nil, false
	}
	return o.EnableMirror, true
}

// HasEnableMirror returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasEnableMirror() bool {
	if o != nil && !IsNil(o.EnableMirror) {
		return true
	}

	return false
}

// SetEnableMirror gets a reference to the given string and assigns it to the EnableMirror field.
func (o *ApiBlockImageImageSpecPutRequest) SetEnableMirror(v string) {
	o.EnableMirror = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetFeatures() string {
	if o == nil || IsNil(o.Features) {
		var ret string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given string and assigns it to the Features field.
func (o *ApiBlockImageImageSpecPutRequest) SetFeatures(v string) {
	o.Features = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *ApiBlockImageImageSpecPutRequest) SetForce(v bool) {
	o.Force = &v
}

// GetImageMirrorMode returns the ImageMirrorMode field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetImageMirrorMode() string {
	if o == nil || IsNil(o.ImageMirrorMode) {
		var ret string
		return ret
	}
	return *o.ImageMirrorMode
}

// GetImageMirrorModeOk returns a tuple with the ImageMirrorMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetImageMirrorModeOk() (*string, bool) {
	if o == nil || IsNil(o.ImageMirrorMode) {
		return nil, false
	}
	return o.ImageMirrorMode, true
}

// HasImageMirrorMode returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasImageMirrorMode() bool {
	if o != nil && !IsNil(o.ImageMirrorMode) {
		return true
	}

	return false
}

// SetImageMirrorMode gets a reference to the given string and assigns it to the ImageMirrorMode field.
func (o *ApiBlockImageImageSpecPutRequest) SetImageMirrorMode(v string) {
	o.ImageMirrorMode = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *ApiBlockImageImageSpecPutRequest) SetMetadata(v string) {
	o.Metadata = &v
}

// GetMirrorMode returns the MirrorMode field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetMirrorMode() string {
	if o == nil || IsNil(o.MirrorMode) {
		var ret string
		return ret
	}
	return *o.MirrorMode
}

// GetMirrorModeOk returns a tuple with the MirrorMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetMirrorModeOk() (*string, bool) {
	if o == nil || IsNil(o.MirrorMode) {
		return nil, false
	}
	return o.MirrorMode, true
}

// HasMirrorMode returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasMirrorMode() bool {
	if o != nil && !IsNil(o.MirrorMode) {
		return true
	}

	return false
}

// SetMirrorMode gets a reference to the given string and assigns it to the MirrorMode field.
func (o *ApiBlockImageImageSpecPutRequest) SetMirrorMode(v string) {
	o.MirrorMode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiBlockImageImageSpecPutRequest) SetName(v string) {
	o.Name = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetPrimary() string {
	if o == nil || IsNil(o.Primary) {
		var ret string
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetPrimaryOk() (*string, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given string and assigns it to the Primary field.
func (o *ApiBlockImageImageSpecPutRequest) SetPrimary(v string) {
	o.Primary = &v
}

// GetRemoveScheduling returns the RemoveScheduling field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetRemoveScheduling() bool {
	if o == nil || IsNil(o.RemoveScheduling) {
		var ret bool
		return ret
	}
	return *o.RemoveScheduling
}

// GetRemoveSchedulingOk returns a tuple with the RemoveScheduling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetRemoveSchedulingOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveScheduling) {
		return nil, false
	}
	return o.RemoveScheduling, true
}

// HasRemoveScheduling returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasRemoveScheduling() bool {
	if o != nil && !IsNil(o.RemoveScheduling) {
		return true
	}

	return false
}

// SetRemoveScheduling gets a reference to the given bool and assigns it to the RemoveScheduling field.
func (o *ApiBlockImageImageSpecPutRequest) SetRemoveScheduling(v bool) {
	o.RemoveScheduling = &v
}

// GetResync returns the Resync field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetResync() bool {
	if o == nil || IsNil(o.Resync) {
		var ret bool
		return ret
	}
	return *o.Resync
}

// GetResyncOk returns a tuple with the Resync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetResyncOk() (*bool, bool) {
	if o == nil || IsNil(o.Resync) {
		return nil, false
	}
	return o.Resync, true
}

// HasResync returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasResync() bool {
	if o != nil && !IsNil(o.Resync) {
		return true
	}

	return false
}

// SetResync gets a reference to the given bool and assigns it to the Resync field.
func (o *ApiBlockImageImageSpecPutRequest) SetResync(v bool) {
	o.Resync = &v
}

// GetScheduleInterval returns the ScheduleInterval field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetScheduleInterval() string {
	if o == nil || IsNil(o.ScheduleInterval) {
		var ret string
		return ret
	}
	return *o.ScheduleInterval
}

// GetScheduleIntervalOk returns a tuple with the ScheduleInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetScheduleIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduleInterval) {
		return nil, false
	}
	return o.ScheduleInterval, true
}

// HasScheduleInterval returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasScheduleInterval() bool {
	if o != nil && !IsNil(o.ScheduleInterval) {
		return true
	}

	return false
}

// SetScheduleInterval gets a reference to the given string and assigns it to the ScheduleInterval field.
func (o *ApiBlockImageImageSpecPutRequest) SetScheduleInterval(v string) {
	o.ScheduleInterval = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecPutRequest) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecPutRequest) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecPutRequest) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ApiBlockImageImageSpecPutRequest) SetSize(v int32) {
	o.Size = &v
}

func (o ApiBlockImageImageSpecPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBlockImageImageSpecPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.EnableMirror) {
		toSerialize["enable_mirror"] = o.EnableMirror
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.ImageMirrorMode) {
		toSerialize["image_mirror_mode"] = o.ImageMirrorMode
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MirrorMode) {
		toSerialize["mirror_mode"] = o.MirrorMode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.RemoveScheduling) {
		toSerialize["remove_scheduling"] = o.RemoveScheduling
	}
	if !IsNil(o.Resync) {
		toSerialize["resync"] = o.Resync
	}
	if !IsNil(o.ScheduleInterval) {
		toSerialize["schedule_interval"] = o.ScheduleInterval
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableApiBlockImageImageSpecPutRequest struct {
	value *ApiBlockImageImageSpecPutRequest
	isSet bool
}

func (v NullableApiBlockImageImageSpecPutRequest) Get() *ApiBlockImageImageSpecPutRequest {
	return v.value
}

func (v *NullableApiBlockImageImageSpecPutRequest) Set(val *ApiBlockImageImageSpecPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBlockImageImageSpecPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBlockImageImageSpecPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBlockImageImageSpecPutRequest(val *ApiBlockImageImageSpecPutRequest) *NullableApiBlockImageImageSpecPutRequest {
	return &NullableApiBlockImageImageSpecPutRequest{value: val, isSet: true}
}

func (v NullableApiBlockImageImageSpecPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBlockImageImageSpecPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


