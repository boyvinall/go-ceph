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

// checks if the ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi{}

// ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi 
type ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi struct {
	// 
	HumanReadableSize string `json:"human_readable_size"`
	// 
	Locked int32 `json:"locked"`
	// 
	Model string `json:"model"`
	// 
	NrRequests string `json:"nr_requests"`
	Partitions ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions `json:"partitions"`
	// 
	Path string `json:"path"`
	// 
	Removable string `json:"removable"`
	// 
	Rev string `json:"rev"`
	// 
	Ro string `json:"ro"`
	// 
	Rotational string `json:"rotational"`
	// 
	SasAddress string `json:"sas_address"`
	// 
	SasDeviceHandle string `json:"sas_device_handle"`
	// 
	SchedulerMode string `json:"scheduler_mode"`
	// 
	Sectors int32 `json:"sectors"`
	// 
	Sectorsize string `json:"sectorsize"`
	// 
	Size int32 `json:"size"`
	// 
	SupportDiscard string `json:"support_discard"`
	// 
	Vendor string `json:"vendor"`
}

type _ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi

// NewApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi instantiates a new ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi(humanReadableSize string, locked int32, model string, nrRequests string, partitions ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions, path string, removable string, rev string, ro string, rotational string, sasAddress string, sasDeviceHandle string, schedulerMode string, sectors int32, sectorsize string, size int32, supportDiscard string, vendor string) *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi {
	this := ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi{}
	this.HumanReadableSize = humanReadableSize
	this.Locked = locked
	this.Model = model
	this.NrRequests = nrRequests
	this.Partitions = partitions
	this.Path = path
	this.Removable = removable
	this.Rev = rev
	this.Ro = ro
	this.Rotational = rotational
	this.SasAddress = sasAddress
	this.SasDeviceHandle = sasDeviceHandle
	this.SchedulerMode = schedulerMode
	this.Sectors = sectors
	this.Sectorsize = sectorsize
	this.Size = size
	this.SupportDiscard = supportDiscard
	this.Vendor = vendor
	return &this
}

// NewApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiWithDefaults instantiates a new ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiWithDefaults() *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi {
	this := ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi{}
	return &this
}

// GetHumanReadableSize returns the HumanReadableSize field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetHumanReadableSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HumanReadableSize
}

// GetHumanReadableSizeOk returns a tuple with the HumanReadableSize field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetHumanReadableSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HumanReadableSize, true
}

// SetHumanReadableSize sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetHumanReadableSize(v string) {
	o.HumanReadableSize = v
}

// GetLocked returns the Locked field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetLocked() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetLockedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetLocked(v int32) {
	o.Locked = v
}

// GetModel returns the Model field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetModel(v string) {
	o.Model = v
}

// GetNrRequests returns the NrRequests field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetNrRequests() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NrRequests
}

// GetNrRequestsOk returns a tuple with the NrRequests field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetNrRequestsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NrRequests, true
}

// SetNrRequests sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetNrRequests(v string) {
	o.NrRequests = v
}

// GetPartitions returns the Partitions field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetPartitions() ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions {
	if o == nil {
		var ret ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions
		return ret
	}

	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetPartitionsOk() (*ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partitions, true
}

// SetPartitions sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetPartitions(v ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApiPartitions) {
	o.Partitions = v
}

// GetPath returns the Path field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetPath(v string) {
	o.Path = v
}

// GetRemovable returns the Removable field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetRemovable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Removable
}

// GetRemovableOk returns a tuple with the Removable field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetRemovableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Removable, true
}

// SetRemovable sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetRemovable(v string) {
	o.Removable = v
}

// GetRev returns the Rev field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetRev() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rev
}

// GetRevOk returns a tuple with the Rev field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetRevOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rev, true
}

// SetRev sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetRev(v string) {
	o.Rev = v
}

// GetRo returns the Ro field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetRo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ro
}

// GetRoOk returns a tuple with the Ro field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetRoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ro, true
}

// SetRo sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetRo(v string) {
	o.Ro = v
}

// GetRotational returns the Rotational field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetRotational() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rotational
}

// GetRotationalOk returns a tuple with the Rotational field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetRotationalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rotational, true
}

// SetRotational sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetRotational(v string) {
	o.Rotational = v
}

// GetSasAddress returns the SasAddress field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSasAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SasAddress
}

// GetSasAddressOk returns a tuple with the SasAddress field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSasAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SasAddress, true
}

// SetSasAddress sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetSasAddress(v string) {
	o.SasAddress = v
}

// GetSasDeviceHandle returns the SasDeviceHandle field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSasDeviceHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SasDeviceHandle
}

// GetSasDeviceHandleOk returns a tuple with the SasDeviceHandle field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSasDeviceHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SasDeviceHandle, true
}

// SetSasDeviceHandle sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetSasDeviceHandle(v string) {
	o.SasDeviceHandle = v
}

// GetSchedulerMode returns the SchedulerMode field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSchedulerMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchedulerMode
}

// GetSchedulerModeOk returns a tuple with the SchedulerMode field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSchedulerModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchedulerMode, true
}

// SetSchedulerMode sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetSchedulerMode(v string) {
	o.SchedulerMode = v
}

// GetSectors returns the Sectors field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSectors() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sectors
}

// GetSectorsOk returns a tuple with the Sectors field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSectorsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sectors, true
}

// SetSectors sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetSectors(v int32) {
	o.Sectors = v
}

// GetSectorsize returns the Sectorsize field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSectorsize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sectorsize
}

// GetSectorsizeOk returns a tuple with the Sectorsize field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSectorsizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sectorsize, true
}

// SetSectorsize sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetSectorsize(v string) {
	o.Sectorsize = v
}

// GetSize returns the Size field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetSize(v int32) {
	o.Size = v
}

// GetSupportDiscard returns the SupportDiscard field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSupportDiscard() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportDiscard
}

// GetSupportDiscardOk returns a tuple with the SupportDiscard field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetSupportDiscardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportDiscard, true
}

// SetSupportDiscard sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetSupportDiscard(v string) {
	o.SupportDiscard = v
}

// GetVendor returns the Vendor field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) SetVendor(v string) {
	o.Vendor = v
}

func (o ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["human_readable_size"] = o.HumanReadableSize
	toSerialize["locked"] = o.Locked
	toSerialize["model"] = o.Model
	toSerialize["nr_requests"] = o.NrRequests
	toSerialize["partitions"] = o.Partitions
	toSerialize["path"] = o.Path
	toSerialize["removable"] = o.Removable
	toSerialize["rev"] = o.Rev
	toSerialize["ro"] = o.Ro
	toSerialize["rotational"] = o.Rotational
	toSerialize["sas_address"] = o.SasAddress
	toSerialize["sas_device_handle"] = o.SasDeviceHandle
	toSerialize["scheduler_mode"] = o.SchedulerMode
	toSerialize["sectors"] = o.Sectors
	toSerialize["sectorsize"] = o.Sectorsize
	toSerialize["size"] = o.Size
	toSerialize["support_discard"] = o.SupportDiscard
	toSerialize["vendor"] = o.Vendor
	return toSerialize, nil
}

func (o *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"human_readable_size",
		"locked",
		"model",
		"nr_requests",
		"partitions",
		"path",
		"removable",
		"rev",
		"ro",
		"rotational",
		"sas_address",
		"sas_device_handle",
		"scheduler_mode",
		"sectors",
		"sectorsize",
		"size",
		"support_discard",
		"vendor",
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

	varApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi := _ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi)

	if err != nil {
		return err
	}

	*o = ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi(varApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi)

	return err
}

type NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi struct {
	value *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi
	isSet bool
}

func (v NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) Get() *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi {
	return v.value
}

func (v *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) Set(val *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi(val *ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi {
	return &NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi{value: val, isSet: true}
}

func (v NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

