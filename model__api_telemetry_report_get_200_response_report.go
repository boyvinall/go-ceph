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

// checks if the ApiTelemetryReportGet200ResponseReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReport{}

// ApiTelemetryReportGet200ResponseReport 
type ApiTelemetryReportGet200ResponseReport struct {
	Balancer ApiTelemetryReportGet200ResponseReportBalancer `json:"balancer"`
	// 
	Channels []string `json:"channels"`
	// 
	ChannelsAvailable []string `json:"channels_available"`
	Config ApiTelemetryReportGet200ResponseReportConfig `json:"config"`
	// 
	Crashes []int32 `json:"crashes"`
	// 
	Created string `json:"created"`
	Crush ApiTelemetryReportGet200ResponseReportCrush `json:"crush"`
	Fs ApiTelemetryReportGet200ResponseReportFs `json:"fs"`
	Hosts ApiTelemetryReportGet200ResponseReportHosts `json:"hosts"`
	// 
	Leaderboard bool `json:"leaderboard"`
	// 
	License string `json:"license"`
	Metadata ApiTelemetryReportGet200ResponseReportMetadata `json:"metadata"`
	Mon ApiTelemetryReportGet200ResponseReportMon `json:"mon"`
	Osd ApiTelemetryReportGet200ResponseReportOsd `json:"osd"`
	// 
	Pools []ApiTelemetryReportGet200ResponseReportPoolsInner `json:"pools"`
	Rbd ApiTelemetryReportGet200ResponseReportRbd `json:"rbd"`
	// 
	ReportId string `json:"report_id"`
	// 
	ReportTimestamp string `json:"report_timestamp"`
	// 
	ReportVersion int32 `json:"report_version"`
	Rgw ApiTelemetryReportGet200ResponseReportRgw `json:"rgw"`
	Services ApiTelemetryReportGet200ResponseReportServices `json:"services"`
	Usage ApiTelemetryReportGet200ResponseReportUsage `json:"usage"`
}

type _ApiTelemetryReportGet200ResponseReport ApiTelemetryReportGet200ResponseReport

// NewApiTelemetryReportGet200ResponseReport instantiates a new ApiTelemetryReportGet200ResponseReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReport(balancer ApiTelemetryReportGet200ResponseReportBalancer, channels []string, channelsAvailable []string, config ApiTelemetryReportGet200ResponseReportConfig, crashes []int32, created string, crush ApiTelemetryReportGet200ResponseReportCrush, fs ApiTelemetryReportGet200ResponseReportFs, hosts ApiTelemetryReportGet200ResponseReportHosts, leaderboard bool, license string, metadata ApiTelemetryReportGet200ResponseReportMetadata, mon ApiTelemetryReportGet200ResponseReportMon, osd ApiTelemetryReportGet200ResponseReportOsd, pools []ApiTelemetryReportGet200ResponseReportPoolsInner, rbd ApiTelemetryReportGet200ResponseReportRbd, reportId string, reportTimestamp string, reportVersion int32, rgw ApiTelemetryReportGet200ResponseReportRgw, services ApiTelemetryReportGet200ResponseReportServices, usage ApiTelemetryReportGet200ResponseReportUsage) *ApiTelemetryReportGet200ResponseReport {
	this := ApiTelemetryReportGet200ResponseReport{}
	this.Balancer = balancer
	this.Channels = channels
	this.ChannelsAvailable = channelsAvailable
	this.Config = config
	this.Crashes = crashes
	this.Created = created
	this.Crush = crush
	this.Fs = fs
	this.Hosts = hosts
	this.Leaderboard = leaderboard
	this.License = license
	this.Metadata = metadata
	this.Mon = mon
	this.Osd = osd
	this.Pools = pools
	this.Rbd = rbd
	this.ReportId = reportId
	this.ReportTimestamp = reportTimestamp
	this.ReportVersion = reportVersion
	this.Rgw = rgw
	this.Services = services
	this.Usage = usage
	return &this
}

// NewApiTelemetryReportGet200ResponseReportWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportWithDefaults() *ApiTelemetryReportGet200ResponseReport {
	this := ApiTelemetryReportGet200ResponseReport{}
	return &this
}

// GetBalancer returns the Balancer field value
func (o *ApiTelemetryReportGet200ResponseReport) GetBalancer() ApiTelemetryReportGet200ResponseReportBalancer {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportBalancer
		return ret
	}

	return o.Balancer
}

// GetBalancerOk returns a tuple with the Balancer field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetBalancerOk() (*ApiTelemetryReportGet200ResponseReportBalancer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balancer, true
}

// SetBalancer sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetBalancer(v ApiTelemetryReportGet200ResponseReportBalancer) {
	o.Balancer = v
}

// GetChannels returns the Channels field value
func (o *ApiTelemetryReportGet200ResponseReport) GetChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetChannelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels, true
}

// SetChannels sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetChannels(v []string) {
	o.Channels = v
}

// GetChannelsAvailable returns the ChannelsAvailable field value
func (o *ApiTelemetryReportGet200ResponseReport) GetChannelsAvailable() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ChannelsAvailable
}

// GetChannelsAvailableOk returns a tuple with the ChannelsAvailable field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetChannelsAvailableOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelsAvailable, true
}

// SetChannelsAvailable sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetChannelsAvailable(v []string) {
	o.ChannelsAvailable = v
}

// GetConfig returns the Config field value
func (o *ApiTelemetryReportGet200ResponseReport) GetConfig() ApiTelemetryReportGet200ResponseReportConfig {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetConfigOk() (*ApiTelemetryReportGet200ResponseReportConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetConfig(v ApiTelemetryReportGet200ResponseReportConfig) {
	o.Config = v
}

// GetCrashes returns the Crashes field value
func (o *ApiTelemetryReportGet200ResponseReport) GetCrashes() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Crashes
}

// GetCrashesOk returns a tuple with the Crashes field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetCrashesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Crashes, true
}

// SetCrashes sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetCrashes(v []int32) {
	o.Crashes = v
}

// GetCreated returns the Created field value
func (o *ApiTelemetryReportGet200ResponseReport) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetCreated(v string) {
	o.Created = v
}

// GetCrush returns the Crush field value
func (o *ApiTelemetryReportGet200ResponseReport) GetCrush() ApiTelemetryReportGet200ResponseReportCrush {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportCrush
		return ret
	}

	return o.Crush
}

// GetCrushOk returns a tuple with the Crush field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetCrushOk() (*ApiTelemetryReportGet200ResponseReportCrush, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Crush, true
}

// SetCrush sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetCrush(v ApiTelemetryReportGet200ResponseReportCrush) {
	o.Crush = v
}

// GetFs returns the Fs field value
func (o *ApiTelemetryReportGet200ResponseReport) GetFs() ApiTelemetryReportGet200ResponseReportFs {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportFs
		return ret
	}

	return o.Fs
}

// GetFsOk returns a tuple with the Fs field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetFsOk() (*ApiTelemetryReportGet200ResponseReportFs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fs, true
}

// SetFs sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetFs(v ApiTelemetryReportGet200ResponseReportFs) {
	o.Fs = v
}

// GetHosts returns the Hosts field value
func (o *ApiTelemetryReportGet200ResponseReport) GetHosts() ApiTelemetryReportGet200ResponseReportHosts {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportHosts
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetHostsOk() (*ApiTelemetryReportGet200ResponseReportHosts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// SetHosts sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetHosts(v ApiTelemetryReportGet200ResponseReportHosts) {
	o.Hosts = v
}

// GetLeaderboard returns the Leaderboard field value
func (o *ApiTelemetryReportGet200ResponseReport) GetLeaderboard() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Leaderboard
}

// GetLeaderboardOk returns a tuple with the Leaderboard field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetLeaderboardOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leaderboard, true
}

// SetLeaderboard sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetLeaderboard(v bool) {
	o.Leaderboard = v
}

// GetLicense returns the License field value
func (o *ApiTelemetryReportGet200ResponseReport) GetLicense() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.License
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetLicenseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.License, true
}

// SetLicense sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetLicense(v string) {
	o.License = v
}

// GetMetadata returns the Metadata field value
func (o *ApiTelemetryReportGet200ResponseReport) GetMetadata() ApiTelemetryReportGet200ResponseReportMetadata {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetMetadataOk() (*ApiTelemetryReportGet200ResponseReportMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetMetadata(v ApiTelemetryReportGet200ResponseReportMetadata) {
	o.Metadata = v
}

// GetMon returns the Mon field value
func (o *ApiTelemetryReportGet200ResponseReport) GetMon() ApiTelemetryReportGet200ResponseReportMon {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportMon
		return ret
	}

	return o.Mon
}

// GetMonOk returns a tuple with the Mon field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetMonOk() (*ApiTelemetryReportGet200ResponseReportMon, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mon, true
}

// SetMon sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetMon(v ApiTelemetryReportGet200ResponseReportMon) {
	o.Mon = v
}

// GetOsd returns the Osd field value
func (o *ApiTelemetryReportGet200ResponseReport) GetOsd() ApiTelemetryReportGet200ResponseReportOsd {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportOsd
		return ret
	}

	return o.Osd
}

// GetOsdOk returns a tuple with the Osd field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetOsdOk() (*ApiTelemetryReportGet200ResponseReportOsd, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Osd, true
}

// SetOsd sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetOsd(v ApiTelemetryReportGet200ResponseReportOsd) {
	o.Osd = v
}

// GetPools returns the Pools field value
func (o *ApiTelemetryReportGet200ResponseReport) GetPools() []ApiTelemetryReportGet200ResponseReportPoolsInner {
	if o == nil {
		var ret []ApiTelemetryReportGet200ResponseReportPoolsInner
		return ret
	}

	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetPoolsOk() ([]ApiTelemetryReportGet200ResponseReportPoolsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pools, true
}

// SetPools sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetPools(v []ApiTelemetryReportGet200ResponseReportPoolsInner) {
	o.Pools = v
}

// GetRbd returns the Rbd field value
func (o *ApiTelemetryReportGet200ResponseReport) GetRbd() ApiTelemetryReportGet200ResponseReportRbd {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportRbd
		return ret
	}

	return o.Rbd
}

// GetRbdOk returns a tuple with the Rbd field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetRbdOk() (*ApiTelemetryReportGet200ResponseReportRbd, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rbd, true
}

// SetRbd sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetRbd(v ApiTelemetryReportGet200ResponseReportRbd) {
	o.Rbd = v
}

// GetReportId returns the ReportId field value
func (o *ApiTelemetryReportGet200ResponseReport) GetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportId, true
}

// SetReportId sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetReportId(v string) {
	o.ReportId = v
}

// GetReportTimestamp returns the ReportTimestamp field value
func (o *ApiTelemetryReportGet200ResponseReport) GetReportTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportTimestamp
}

// GetReportTimestampOk returns a tuple with the ReportTimestamp field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetReportTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportTimestamp, true
}

// SetReportTimestamp sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetReportTimestamp(v string) {
	o.ReportTimestamp = v
}

// GetReportVersion returns the ReportVersion field value
func (o *ApiTelemetryReportGet200ResponseReport) GetReportVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReportVersion
}

// GetReportVersionOk returns a tuple with the ReportVersion field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetReportVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportVersion, true
}

// SetReportVersion sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetReportVersion(v int32) {
	o.ReportVersion = v
}

// GetRgw returns the Rgw field value
func (o *ApiTelemetryReportGet200ResponseReport) GetRgw() ApiTelemetryReportGet200ResponseReportRgw {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportRgw
		return ret
	}

	return o.Rgw
}

// GetRgwOk returns a tuple with the Rgw field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetRgwOk() (*ApiTelemetryReportGet200ResponseReportRgw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rgw, true
}

// SetRgw sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetRgw(v ApiTelemetryReportGet200ResponseReportRgw) {
	o.Rgw = v
}

// GetServices returns the Services field value
func (o *ApiTelemetryReportGet200ResponseReport) GetServices() ApiTelemetryReportGet200ResponseReportServices {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportServices
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetServicesOk() (*ApiTelemetryReportGet200ResponseReportServices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetServices(v ApiTelemetryReportGet200ResponseReportServices) {
	o.Services = v
}

// GetUsage returns the Usage field value
func (o *ApiTelemetryReportGet200ResponseReport) GetUsage() ApiTelemetryReportGet200ResponseReportUsage {
	if o == nil {
		var ret ApiTelemetryReportGet200ResponseReportUsage
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReport) GetUsageOk() (*ApiTelemetryReportGet200ResponseReportUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *ApiTelemetryReportGet200ResponseReport) SetUsage(v ApiTelemetryReportGet200ResponseReportUsage) {
	o.Usage = v
}

func (o ApiTelemetryReportGet200ResponseReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balancer"] = o.Balancer
	toSerialize["channels"] = o.Channels
	toSerialize["channels_available"] = o.ChannelsAvailable
	toSerialize["config"] = o.Config
	toSerialize["crashes"] = o.Crashes
	toSerialize["created"] = o.Created
	toSerialize["crush"] = o.Crush
	toSerialize["fs"] = o.Fs
	toSerialize["hosts"] = o.Hosts
	toSerialize["leaderboard"] = o.Leaderboard
	toSerialize["license"] = o.License
	toSerialize["metadata"] = o.Metadata
	toSerialize["mon"] = o.Mon
	toSerialize["osd"] = o.Osd
	toSerialize["pools"] = o.Pools
	toSerialize["rbd"] = o.Rbd
	toSerialize["report_id"] = o.ReportId
	toSerialize["report_timestamp"] = o.ReportTimestamp
	toSerialize["report_version"] = o.ReportVersion
	toSerialize["rgw"] = o.Rgw
	toSerialize["services"] = o.Services
	toSerialize["usage"] = o.Usage
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"balancer",
		"channels",
		"channels_available",
		"config",
		"crashes",
		"created",
		"crush",
		"fs",
		"hosts",
		"leaderboard",
		"license",
		"metadata",
		"mon",
		"osd",
		"pools",
		"rbd",
		"report_id",
		"report_timestamp",
		"report_version",
		"rgw",
		"services",
		"usage",
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

	varApiTelemetryReportGet200ResponseReport := _ApiTelemetryReportGet200ResponseReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReport)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReport(varApiTelemetryReportGet200ResponseReport)

	return err
}

type NullableApiTelemetryReportGet200ResponseReport struct {
	value *ApiTelemetryReportGet200ResponseReport
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReport) Get() *ApiTelemetryReportGet200ResponseReport {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReport) Set(val *ApiTelemetryReportGet200ResponseReport) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReport) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReport(val *ApiTelemetryReportGet200ResponseReport) *NullableApiTelemetryReportGet200ResponseReport {
	return &NullableApiTelemetryReportGet200ResponseReport{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


