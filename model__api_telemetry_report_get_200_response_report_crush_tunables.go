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

// checks if the ApiTelemetryReportGet200ResponseReportCrushTunables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTelemetryReportGet200ResponseReportCrushTunables{}

// ApiTelemetryReportGet200ResponseReportCrushTunables 
type ApiTelemetryReportGet200ResponseReportCrushTunables struct {
	// 
	AllowedBucketAlgs int32 `json:"allowed_bucket_algs"`
	// 
	ChooseLocalFallbackTries int32 `json:"choose_local_fallback_tries"`
	// 
	ChooseLocalTries int32 `json:"choose_local_tries"`
	// 
	ChooseTotalTries int32 `json:"choose_total_tries"`
	// 
	ChooseleafDescendOnce int32 `json:"chooseleaf_descend_once"`
	// 
	ChooseleafStable int32 `json:"chooseleaf_stable"`
	// 
	ChooseleafVaryR int32 `json:"chooseleaf_vary_r"`
	// 
	HasV2Rules int32 `json:"has_v2_rules"`
	// 
	HasV3Rules int32 `json:"has_v3_rules"`
	// 
	HasV4Buckets int32 `json:"has_v4_buckets"`
	// 
	HasV5Rules int32 `json:"has_v5_rules"`
	// 
	LegacyTunables int32 `json:"legacy_tunables"`
	// 
	MinimumRequiredVersion string `json:"minimum_required_version"`
	// 
	OptimalTunables int32 `json:"optimal_tunables"`
	// 
	Profile string `json:"profile"`
	// 
	RequireFeatureTunables int32 `json:"require_feature_tunables"`
	// 
	RequireFeatureTunables2 int32 `json:"require_feature_tunables2"`
	// 
	RequireFeatureTunables3 int32 `json:"require_feature_tunables3"`
	// 
	RequireFeatureTunables5 int32 `json:"require_feature_tunables5"`
	// 
	StrawCalcVersion int32 `json:"straw_calc_version"`
}

type _ApiTelemetryReportGet200ResponseReportCrushTunables ApiTelemetryReportGet200ResponseReportCrushTunables

// NewApiTelemetryReportGet200ResponseReportCrushTunables instantiates a new ApiTelemetryReportGet200ResponseReportCrushTunables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTelemetryReportGet200ResponseReportCrushTunables(allowedBucketAlgs int32, chooseLocalFallbackTries int32, chooseLocalTries int32, chooseTotalTries int32, chooseleafDescendOnce int32, chooseleafStable int32, chooseleafVaryR int32, hasV2Rules int32, hasV3Rules int32, hasV4Buckets int32, hasV5Rules int32, legacyTunables int32, minimumRequiredVersion string, optimalTunables int32, profile string, requireFeatureTunables int32, requireFeatureTunables2 int32, requireFeatureTunables3 int32, requireFeatureTunables5 int32, strawCalcVersion int32) *ApiTelemetryReportGet200ResponseReportCrushTunables {
	this := ApiTelemetryReportGet200ResponseReportCrushTunables{}
	this.AllowedBucketAlgs = allowedBucketAlgs
	this.ChooseLocalFallbackTries = chooseLocalFallbackTries
	this.ChooseLocalTries = chooseLocalTries
	this.ChooseTotalTries = chooseTotalTries
	this.ChooseleafDescendOnce = chooseleafDescendOnce
	this.ChooseleafStable = chooseleafStable
	this.ChooseleafVaryR = chooseleafVaryR
	this.HasV2Rules = hasV2Rules
	this.HasV3Rules = hasV3Rules
	this.HasV4Buckets = hasV4Buckets
	this.HasV5Rules = hasV5Rules
	this.LegacyTunables = legacyTunables
	this.MinimumRequiredVersion = minimumRequiredVersion
	this.OptimalTunables = optimalTunables
	this.Profile = profile
	this.RequireFeatureTunables = requireFeatureTunables
	this.RequireFeatureTunables2 = requireFeatureTunables2
	this.RequireFeatureTunables3 = requireFeatureTunables3
	this.RequireFeatureTunables5 = requireFeatureTunables5
	this.StrawCalcVersion = strawCalcVersion
	return &this
}

// NewApiTelemetryReportGet200ResponseReportCrushTunablesWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReportCrushTunables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTelemetryReportGet200ResponseReportCrushTunablesWithDefaults() *ApiTelemetryReportGet200ResponseReportCrushTunables {
	this := ApiTelemetryReportGet200ResponseReportCrushTunables{}
	return &this
}

// GetAllowedBucketAlgs returns the AllowedBucketAlgs field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetAllowedBucketAlgs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AllowedBucketAlgs
}

// GetAllowedBucketAlgsOk returns a tuple with the AllowedBucketAlgs field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetAllowedBucketAlgsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowedBucketAlgs, true
}

// SetAllowedBucketAlgs sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetAllowedBucketAlgs(v int32) {
	o.AllowedBucketAlgs = v
}

// GetChooseLocalFallbackTries returns the ChooseLocalFallbackTries field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseLocalFallbackTries() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChooseLocalFallbackTries
}

// GetChooseLocalFallbackTriesOk returns a tuple with the ChooseLocalFallbackTries field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseLocalFallbackTriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChooseLocalFallbackTries, true
}

// SetChooseLocalFallbackTries sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetChooseLocalFallbackTries(v int32) {
	o.ChooseLocalFallbackTries = v
}

// GetChooseLocalTries returns the ChooseLocalTries field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseLocalTries() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChooseLocalTries
}

// GetChooseLocalTriesOk returns a tuple with the ChooseLocalTries field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseLocalTriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChooseLocalTries, true
}

// SetChooseLocalTries sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetChooseLocalTries(v int32) {
	o.ChooseLocalTries = v
}

// GetChooseTotalTries returns the ChooseTotalTries field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseTotalTries() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChooseTotalTries
}

// GetChooseTotalTriesOk returns a tuple with the ChooseTotalTries field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseTotalTriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChooseTotalTries, true
}

// SetChooseTotalTries sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetChooseTotalTries(v int32) {
	o.ChooseTotalTries = v
}

// GetChooseleafDescendOnce returns the ChooseleafDescendOnce field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseleafDescendOnce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChooseleafDescendOnce
}

// GetChooseleafDescendOnceOk returns a tuple with the ChooseleafDescendOnce field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseleafDescendOnceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChooseleafDescendOnce, true
}

// SetChooseleafDescendOnce sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetChooseleafDescendOnce(v int32) {
	o.ChooseleafDescendOnce = v
}

// GetChooseleafStable returns the ChooseleafStable field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseleafStable() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChooseleafStable
}

// GetChooseleafStableOk returns a tuple with the ChooseleafStable field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseleafStableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChooseleafStable, true
}

// SetChooseleafStable sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetChooseleafStable(v int32) {
	o.ChooseleafStable = v
}

// GetChooseleafVaryR returns the ChooseleafVaryR field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseleafVaryR() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChooseleafVaryR
}

// GetChooseleafVaryROk returns a tuple with the ChooseleafVaryR field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetChooseleafVaryROk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChooseleafVaryR, true
}

// SetChooseleafVaryR sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetChooseleafVaryR(v int32) {
	o.ChooseleafVaryR = v
}

// GetHasV2Rules returns the HasV2Rules field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetHasV2Rules() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HasV2Rules
}

// GetHasV2RulesOk returns a tuple with the HasV2Rules field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetHasV2RulesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasV2Rules, true
}

// SetHasV2Rules sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetHasV2Rules(v int32) {
	o.HasV2Rules = v
}

// GetHasV3Rules returns the HasV3Rules field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetHasV3Rules() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HasV3Rules
}

// GetHasV3RulesOk returns a tuple with the HasV3Rules field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetHasV3RulesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasV3Rules, true
}

// SetHasV3Rules sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetHasV3Rules(v int32) {
	o.HasV3Rules = v
}

// GetHasV4Buckets returns the HasV4Buckets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetHasV4Buckets() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HasV4Buckets
}

// GetHasV4BucketsOk returns a tuple with the HasV4Buckets field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetHasV4BucketsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasV4Buckets, true
}

// SetHasV4Buckets sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetHasV4Buckets(v int32) {
	o.HasV4Buckets = v
}

// GetHasV5Rules returns the HasV5Rules field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetHasV5Rules() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HasV5Rules
}

// GetHasV5RulesOk returns a tuple with the HasV5Rules field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetHasV5RulesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasV5Rules, true
}

// SetHasV5Rules sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetHasV5Rules(v int32) {
	o.HasV5Rules = v
}

// GetLegacyTunables returns the LegacyTunables field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetLegacyTunables() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LegacyTunables
}

// GetLegacyTunablesOk returns a tuple with the LegacyTunables field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetLegacyTunablesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegacyTunables, true
}

// SetLegacyTunables sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetLegacyTunables(v int32) {
	o.LegacyTunables = v
}

// GetMinimumRequiredVersion returns the MinimumRequiredVersion field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetMinimumRequiredVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumRequiredVersion
}

// GetMinimumRequiredVersionOk returns a tuple with the MinimumRequiredVersion field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetMinimumRequiredVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumRequiredVersion, true
}

// SetMinimumRequiredVersion sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetMinimumRequiredVersion(v string) {
	o.MinimumRequiredVersion = v
}

// GetOptimalTunables returns the OptimalTunables field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetOptimalTunables() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OptimalTunables
}

// GetOptimalTunablesOk returns a tuple with the OptimalTunables field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetOptimalTunablesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimalTunables, true
}

// SetOptimalTunables sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetOptimalTunables(v int32) {
	o.OptimalTunables = v
}

// GetProfile returns the Profile field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetProfile(v string) {
	o.Profile = v
}

// GetRequireFeatureTunables returns the RequireFeatureTunables field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetRequireFeatureTunables() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequireFeatureTunables
}

// GetRequireFeatureTunablesOk returns a tuple with the RequireFeatureTunables field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetRequireFeatureTunablesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireFeatureTunables, true
}

// SetRequireFeatureTunables sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetRequireFeatureTunables(v int32) {
	o.RequireFeatureTunables = v
}

// GetRequireFeatureTunables2 returns the RequireFeatureTunables2 field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetRequireFeatureTunables2() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequireFeatureTunables2
}

// GetRequireFeatureTunables2Ok returns a tuple with the RequireFeatureTunables2 field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetRequireFeatureTunables2Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireFeatureTunables2, true
}

// SetRequireFeatureTunables2 sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetRequireFeatureTunables2(v int32) {
	o.RequireFeatureTunables2 = v
}

// GetRequireFeatureTunables3 returns the RequireFeatureTunables3 field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetRequireFeatureTunables3() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequireFeatureTunables3
}

// GetRequireFeatureTunables3Ok returns a tuple with the RequireFeatureTunables3 field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetRequireFeatureTunables3Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireFeatureTunables3, true
}

// SetRequireFeatureTunables3 sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetRequireFeatureTunables3(v int32) {
	o.RequireFeatureTunables3 = v
}

// GetRequireFeatureTunables5 returns the RequireFeatureTunables5 field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetRequireFeatureTunables5() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequireFeatureTunables5
}

// GetRequireFeatureTunables5Ok returns a tuple with the RequireFeatureTunables5 field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetRequireFeatureTunables5Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireFeatureTunables5, true
}

// SetRequireFeatureTunables5 sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetRequireFeatureTunables5(v int32) {
	o.RequireFeatureTunables5 = v
}

// GetStrawCalcVersion returns the StrawCalcVersion field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetStrawCalcVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StrawCalcVersion
}

// GetStrawCalcVersionOk returns a tuple with the StrawCalcVersion field value
// and a boolean to check if the value has been set.
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) GetStrawCalcVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StrawCalcVersion, true
}

// SetStrawCalcVersion sets field value
func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) SetStrawCalcVersion(v int32) {
	o.StrawCalcVersion = v
}

func (o ApiTelemetryReportGet200ResponseReportCrushTunables) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTelemetryReportGet200ResponseReportCrushTunables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowed_bucket_algs"] = o.AllowedBucketAlgs
	toSerialize["choose_local_fallback_tries"] = o.ChooseLocalFallbackTries
	toSerialize["choose_local_tries"] = o.ChooseLocalTries
	toSerialize["choose_total_tries"] = o.ChooseTotalTries
	toSerialize["chooseleaf_descend_once"] = o.ChooseleafDescendOnce
	toSerialize["chooseleaf_stable"] = o.ChooseleafStable
	toSerialize["chooseleaf_vary_r"] = o.ChooseleafVaryR
	toSerialize["has_v2_rules"] = o.HasV2Rules
	toSerialize["has_v3_rules"] = o.HasV3Rules
	toSerialize["has_v4_buckets"] = o.HasV4Buckets
	toSerialize["has_v5_rules"] = o.HasV5Rules
	toSerialize["legacy_tunables"] = o.LegacyTunables
	toSerialize["minimum_required_version"] = o.MinimumRequiredVersion
	toSerialize["optimal_tunables"] = o.OptimalTunables
	toSerialize["profile"] = o.Profile
	toSerialize["require_feature_tunables"] = o.RequireFeatureTunables
	toSerialize["require_feature_tunables2"] = o.RequireFeatureTunables2
	toSerialize["require_feature_tunables3"] = o.RequireFeatureTunables3
	toSerialize["require_feature_tunables5"] = o.RequireFeatureTunables5
	toSerialize["straw_calc_version"] = o.StrawCalcVersion
	return toSerialize, nil
}

func (o *ApiTelemetryReportGet200ResponseReportCrushTunables) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allowed_bucket_algs",
		"choose_local_fallback_tries",
		"choose_local_tries",
		"choose_total_tries",
		"chooseleaf_descend_once",
		"chooseleaf_stable",
		"chooseleaf_vary_r",
		"has_v2_rules",
		"has_v3_rules",
		"has_v4_buckets",
		"has_v5_rules",
		"legacy_tunables",
		"minimum_required_version",
		"optimal_tunables",
		"profile",
		"require_feature_tunables",
		"require_feature_tunables2",
		"require_feature_tunables3",
		"require_feature_tunables5",
		"straw_calc_version",
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

	varApiTelemetryReportGet200ResponseReportCrushTunables := _ApiTelemetryReportGet200ResponseReportCrushTunables{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTelemetryReportGet200ResponseReportCrushTunables)

	if err != nil {
		return err
	}

	*o = ApiTelemetryReportGet200ResponseReportCrushTunables(varApiTelemetryReportGet200ResponseReportCrushTunables)

	return err
}

type NullableApiTelemetryReportGet200ResponseReportCrushTunables struct {
	value *ApiTelemetryReportGet200ResponseReportCrushTunables
	isSet bool
}

func (v NullableApiTelemetryReportGet200ResponseReportCrushTunables) Get() *ApiTelemetryReportGet200ResponseReportCrushTunables {
	return v.value
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrushTunables) Set(val *ApiTelemetryReportGet200ResponseReportCrushTunables) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTelemetryReportGet200ResponseReportCrushTunables) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrushTunables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTelemetryReportGet200ResponseReportCrushTunables(val *ApiTelemetryReportGet200ResponseReportCrushTunables) *NullableApiTelemetryReportGet200ResponseReportCrushTunables {
	return &NullableApiTelemetryReportGet200ResponseReportCrushTunables{value: val, isSet: true}
}

func (v NullableApiTelemetryReportGet200ResponseReportCrushTunables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTelemetryReportGet200ResponseReportCrushTunables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


