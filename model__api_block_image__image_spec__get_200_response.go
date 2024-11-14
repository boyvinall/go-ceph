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

// checks if the ApiBlockImageImageSpecGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBlockImageImageSpecGet200Response{}

// ApiBlockImageImageSpecGet200Response struct for ApiBlockImageImageSpecGet200Response
type ApiBlockImageImageSpecGet200Response struct {
	// 
	Size *float32 `json:"size,omitempty"`
	// 
	ObjSize *int32 `json:"obj_size,omitempty"`
	// 
	NumObjs *int32 `json:"num_objs,omitempty"`
	// 
	Order *int32 `json:"order,omitempty"`
	// 
	BlockNamePrefix *string `json:"block_name_prefix,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	// 
	UniqueId *string `json:"unique_id,omitempty"`
	// 
	Id *string `json:"id,omitempty"`
	// 
	ImageFormat *int32 `json:"image_format,omitempty"`
	// 
	PoolName *string `json:"pool_name,omitempty"`
	// 
	Namespace *string `json:"namespace,omitempty"`
	// 
	Features *int32 `json:"features,omitempty"`
	// 
	FeaturesName []string `json:"features_name,omitempty"`
	// 
	Timestamp *string `json:"timestamp,omitempty"`
	// 
	StripeCount *int32 `json:"stripe_count,omitempty"`
	// 
	StripeUnit *int32 `json:"stripe_unit,omitempty"`
	// 
	DataPool NullableString `json:"data_pool,omitempty"`
	// 
	Parent NullableString `json:"parent,omitempty"`
	// 
	Snapshots []ApiBlockImageImageSpecGet200ResponseSnapshotsInner `json:"snapshots,omitempty"`
	// 
	TotalDiskUsage *float32 `json:"total_disk_usage,omitempty"`
	// 
	DiskUsage *int32 `json:"disk_usage,omitempty"`
	// 
	Configuration []ApiBlockImageImageSpecGet200ResponseConfigurationInner `json:"configuration,omitempty"`
}

// NewApiBlockImageImageSpecGet200Response instantiates a new ApiBlockImageImageSpecGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBlockImageImageSpecGet200Response() *ApiBlockImageImageSpecGet200Response {
	this := ApiBlockImageImageSpecGet200Response{}
	return &this
}

// NewApiBlockImageImageSpecGet200ResponseWithDefaults instantiates a new ApiBlockImageImageSpecGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBlockImageImageSpecGet200ResponseWithDefaults() *ApiBlockImageImageSpecGet200Response {
	this := ApiBlockImageImageSpecGet200Response{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ApiBlockImageImageSpecGet200Response) SetSize(v float32) {
	o.Size = &v
}

// GetObjSize returns the ObjSize field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetObjSize() int32 {
	if o == nil || IsNil(o.ObjSize) {
		var ret int32
		return ret
	}
	return *o.ObjSize
}

// GetObjSizeOk returns a tuple with the ObjSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetObjSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ObjSize) {
		return nil, false
	}
	return o.ObjSize, true
}

// HasObjSize returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasObjSize() bool {
	if o != nil && !IsNil(o.ObjSize) {
		return true
	}

	return false
}

// SetObjSize gets a reference to the given int32 and assigns it to the ObjSize field.
func (o *ApiBlockImageImageSpecGet200Response) SetObjSize(v int32) {
	o.ObjSize = &v
}

// GetNumObjs returns the NumObjs field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetNumObjs() int32 {
	if o == nil || IsNil(o.NumObjs) {
		var ret int32
		return ret
	}
	return *o.NumObjs
}

// GetNumObjsOk returns a tuple with the NumObjs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetNumObjsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumObjs) {
		return nil, false
	}
	return o.NumObjs, true
}

// HasNumObjs returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasNumObjs() bool {
	if o != nil && !IsNil(o.NumObjs) {
		return true
	}

	return false
}

// SetNumObjs gets a reference to the given int32 and assigns it to the NumObjs field.
func (o *ApiBlockImageImageSpecGet200Response) SetNumObjs(v int32) {
	o.NumObjs = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ApiBlockImageImageSpecGet200Response) SetOrder(v int32) {
	o.Order = &v
}

// GetBlockNamePrefix returns the BlockNamePrefix field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetBlockNamePrefix() string {
	if o == nil || IsNil(o.BlockNamePrefix) {
		var ret string
		return ret
	}
	return *o.BlockNamePrefix
}

// GetBlockNamePrefixOk returns a tuple with the BlockNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetBlockNamePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.BlockNamePrefix) {
		return nil, false
	}
	return o.BlockNamePrefix, true
}

// HasBlockNamePrefix returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasBlockNamePrefix() bool {
	if o != nil && !IsNil(o.BlockNamePrefix) {
		return true
	}

	return false
}

// SetBlockNamePrefix gets a reference to the given string and assigns it to the BlockNamePrefix field.
func (o *ApiBlockImageImageSpecGet200Response) SetBlockNamePrefix(v string) {
	o.BlockNamePrefix = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiBlockImageImageSpecGet200Response) SetName(v string) {
	o.Name = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetUniqueId() string {
	if o == nil || IsNil(o.UniqueId) {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetUniqueIdOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueId) {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasUniqueId() bool {
	if o != nil && !IsNil(o.UniqueId) {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *ApiBlockImageImageSpecGet200Response) SetUniqueId(v string) {
	o.UniqueId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiBlockImageImageSpecGet200Response) SetId(v string) {
	o.Id = &v
}

// GetImageFormat returns the ImageFormat field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetImageFormat() int32 {
	if o == nil || IsNil(o.ImageFormat) {
		var ret int32
		return ret
	}
	return *o.ImageFormat
}

// GetImageFormatOk returns a tuple with the ImageFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetImageFormatOk() (*int32, bool) {
	if o == nil || IsNil(o.ImageFormat) {
		return nil, false
	}
	return o.ImageFormat, true
}

// HasImageFormat returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasImageFormat() bool {
	if o != nil && !IsNil(o.ImageFormat) {
		return true
	}

	return false
}

// SetImageFormat gets a reference to the given int32 and assigns it to the ImageFormat field.
func (o *ApiBlockImageImageSpecGet200Response) SetImageFormat(v int32) {
	o.ImageFormat = &v
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetPoolName() string {
	if o == nil || IsNil(o.PoolName) {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetPoolNameOk() (*string, bool) {
	if o == nil || IsNil(o.PoolName) {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasPoolName() bool {
	if o != nil && !IsNil(o.PoolName) {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *ApiBlockImageImageSpecGet200Response) SetPoolName(v string) {
	o.PoolName = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ApiBlockImageImageSpecGet200Response) SetNamespace(v string) {
	o.Namespace = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetFeatures() int32 {
	if o == nil || IsNil(o.Features) {
		var ret int32
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetFeaturesOk() (*int32, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given int32 and assigns it to the Features field.
func (o *ApiBlockImageImageSpecGet200Response) SetFeatures(v int32) {
	o.Features = &v
}

// GetFeaturesName returns the FeaturesName field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetFeaturesName() []string {
	if o == nil || IsNil(o.FeaturesName) {
		var ret []string
		return ret
	}
	return o.FeaturesName
}

// GetFeaturesNameOk returns a tuple with the FeaturesName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetFeaturesNameOk() ([]string, bool) {
	if o == nil || IsNil(o.FeaturesName) {
		return nil, false
	}
	return o.FeaturesName, true
}

// HasFeaturesName returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasFeaturesName() bool {
	if o != nil && !IsNil(o.FeaturesName) {
		return true
	}

	return false
}

// SetFeaturesName gets a reference to the given []string and assigns it to the FeaturesName field.
func (o *ApiBlockImageImageSpecGet200Response) SetFeaturesName(v []string) {
	o.FeaturesName = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ApiBlockImageImageSpecGet200Response) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetStripeCount returns the StripeCount field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetStripeCount() int32 {
	if o == nil || IsNil(o.StripeCount) {
		var ret int32
		return ret
	}
	return *o.StripeCount
}

// GetStripeCountOk returns a tuple with the StripeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetStripeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.StripeCount) {
		return nil, false
	}
	return o.StripeCount, true
}

// HasStripeCount returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasStripeCount() bool {
	if o != nil && !IsNil(o.StripeCount) {
		return true
	}

	return false
}

// SetStripeCount gets a reference to the given int32 and assigns it to the StripeCount field.
func (o *ApiBlockImageImageSpecGet200Response) SetStripeCount(v int32) {
	o.StripeCount = &v
}

// GetStripeUnit returns the StripeUnit field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetStripeUnit() int32 {
	if o == nil || IsNil(o.StripeUnit) {
		var ret int32
		return ret
	}
	return *o.StripeUnit
}

// GetStripeUnitOk returns a tuple with the StripeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetStripeUnitOk() (*int32, bool) {
	if o == nil || IsNil(o.StripeUnit) {
		return nil, false
	}
	return o.StripeUnit, true
}

// HasStripeUnit returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasStripeUnit() bool {
	if o != nil && !IsNil(o.StripeUnit) {
		return true
	}

	return false
}

// SetStripeUnit gets a reference to the given int32 and assigns it to the StripeUnit field.
func (o *ApiBlockImageImageSpecGet200Response) SetStripeUnit(v int32) {
	o.StripeUnit = &v
}

// GetDataPool returns the DataPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiBlockImageImageSpecGet200Response) GetDataPool() string {
	if o == nil || IsNil(o.DataPool.Get()) {
		var ret string
		return ret
	}
	return *o.DataPool.Get()
}

// GetDataPoolOk returns a tuple with the DataPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiBlockImageImageSpecGet200Response) GetDataPoolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataPool.Get(), o.DataPool.IsSet()
}

// HasDataPool returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasDataPool() bool {
	if o != nil && o.DataPool.IsSet() {
		return true
	}

	return false
}

// SetDataPool gets a reference to the given NullableString and assigns it to the DataPool field.
func (o *ApiBlockImageImageSpecGet200Response) SetDataPool(v string) {
	o.DataPool.Set(&v)
}
// SetDataPoolNil sets the value for DataPool to be an explicit nil
func (o *ApiBlockImageImageSpecGet200Response) SetDataPoolNil() {
	o.DataPool.Set(nil)
}

// UnsetDataPool ensures that no value is present for DataPool, not even an explicit nil
func (o *ApiBlockImageImageSpecGet200Response) UnsetDataPool() {
	o.DataPool.Unset()
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiBlockImageImageSpecGet200Response) GetParent() string {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiBlockImageImageSpecGet200Response) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *ApiBlockImageImageSpecGet200Response) SetParent(v string) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *ApiBlockImageImageSpecGet200Response) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *ApiBlockImageImageSpecGet200Response) UnsetParent() {
	o.Parent.Unset()
}

// GetSnapshots returns the Snapshots field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetSnapshots() []ApiBlockImageImageSpecGet200ResponseSnapshotsInner {
	if o == nil || IsNil(o.Snapshots) {
		var ret []ApiBlockImageImageSpecGet200ResponseSnapshotsInner
		return ret
	}
	return o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetSnapshotsOk() ([]ApiBlockImageImageSpecGet200ResponseSnapshotsInner, bool) {
	if o == nil || IsNil(o.Snapshots) {
		return nil, false
	}
	return o.Snapshots, true
}

// HasSnapshots returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasSnapshots() bool {
	if o != nil && !IsNil(o.Snapshots) {
		return true
	}

	return false
}

// SetSnapshots gets a reference to the given []ApiBlockImageImageSpecGet200ResponseSnapshotsInner and assigns it to the Snapshots field.
func (o *ApiBlockImageImageSpecGet200Response) SetSnapshots(v []ApiBlockImageImageSpecGet200ResponseSnapshotsInner) {
	o.Snapshots = v
}

// GetTotalDiskUsage returns the TotalDiskUsage field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetTotalDiskUsage() float32 {
	if o == nil || IsNil(o.TotalDiskUsage) {
		var ret float32
		return ret
	}
	return *o.TotalDiskUsage
}

// GetTotalDiskUsageOk returns a tuple with the TotalDiskUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetTotalDiskUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalDiskUsage) {
		return nil, false
	}
	return o.TotalDiskUsage, true
}

// HasTotalDiskUsage returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasTotalDiskUsage() bool {
	if o != nil && !IsNil(o.TotalDiskUsage) {
		return true
	}

	return false
}

// SetTotalDiskUsage gets a reference to the given float32 and assigns it to the TotalDiskUsage field.
func (o *ApiBlockImageImageSpecGet200Response) SetTotalDiskUsage(v float32) {
	o.TotalDiskUsage = &v
}

// GetDiskUsage returns the DiskUsage field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetDiskUsage() int32 {
	if o == nil || IsNil(o.DiskUsage) {
		var ret int32
		return ret
	}
	return *o.DiskUsage
}

// GetDiskUsageOk returns a tuple with the DiskUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetDiskUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskUsage) {
		return nil, false
	}
	return o.DiskUsage, true
}

// HasDiskUsage returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasDiskUsage() bool {
	if o != nil && !IsNil(o.DiskUsage) {
		return true
	}

	return false
}

// SetDiskUsage gets a reference to the given int32 and assigns it to the DiskUsage field.
func (o *ApiBlockImageImageSpecGet200Response) SetDiskUsage(v int32) {
	o.DiskUsage = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ApiBlockImageImageSpecGet200Response) GetConfiguration() []ApiBlockImageImageSpecGet200ResponseConfigurationInner {
	if o == nil || IsNil(o.Configuration) {
		var ret []ApiBlockImageImageSpecGet200ResponseConfigurationInner
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBlockImageImageSpecGet200Response) GetConfigurationOk() ([]ApiBlockImageImageSpecGet200ResponseConfigurationInner, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ApiBlockImageImageSpecGet200Response) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given []ApiBlockImageImageSpecGet200ResponseConfigurationInner and assigns it to the Configuration field.
func (o *ApiBlockImageImageSpecGet200Response) SetConfiguration(v []ApiBlockImageImageSpecGet200ResponseConfigurationInner) {
	o.Configuration = v
}

func (o ApiBlockImageImageSpecGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBlockImageImageSpecGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.ObjSize) {
		toSerialize["obj_size"] = o.ObjSize
	}
	if !IsNil(o.NumObjs) {
		toSerialize["num_objs"] = o.NumObjs
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.BlockNamePrefix) {
		toSerialize["block_name_prefix"] = o.BlockNamePrefix
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UniqueId) {
		toSerialize["unique_id"] = o.UniqueId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImageFormat) {
		toSerialize["image_format"] = o.ImageFormat
	}
	if !IsNil(o.PoolName) {
		toSerialize["pool_name"] = o.PoolName
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.FeaturesName) {
		toSerialize["features_name"] = o.FeaturesName
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.StripeCount) {
		toSerialize["stripe_count"] = o.StripeCount
	}
	if !IsNil(o.StripeUnit) {
		toSerialize["stripe_unit"] = o.StripeUnit
	}
	if o.DataPool.IsSet() {
		toSerialize["data_pool"] = o.DataPool.Get()
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Snapshots) {
		toSerialize["snapshots"] = o.Snapshots
	}
	if !IsNil(o.TotalDiskUsage) {
		toSerialize["total_disk_usage"] = o.TotalDiskUsage
	}
	if !IsNil(o.DiskUsage) {
		toSerialize["disk_usage"] = o.DiskUsage
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
}

type NullableApiBlockImageImageSpecGet200Response struct {
	value *ApiBlockImageImageSpecGet200Response
	isSet bool
}

func (v NullableApiBlockImageImageSpecGet200Response) Get() *ApiBlockImageImageSpecGet200Response {
	return v.value
}

func (v *NullableApiBlockImageImageSpecGet200Response) Set(val *ApiBlockImageImageSpecGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBlockImageImageSpecGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBlockImageImageSpecGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBlockImageImageSpecGet200Response(val *ApiBlockImageImageSpecGet200Response) *NullableApiBlockImageImageSpecGet200Response {
	return &NullableApiBlockImageImageSpecGet200Response{value: val, isSet: true}
}

func (v NullableApiBlockImageImageSpecGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBlockImageImageSpecGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


