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

// checks if the ApiHostGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHostGet200Response{}

// ApiHostGet200Response struct for ApiHostGet200Response
type ApiHostGet200Response struct {
	// Host address
	Addr string `json:"addr"`
	// Ceph version
	CephVersion string `json:"ceph_version"`
	// Hostname
	Hostname string `json:"hostname"`
	// Labels related to the host
	Labels []string `json:"labels"`
	// Service instances related to the host
	ServiceInstances []ApiHostGet200ResponseServiceInstancesInner `json:"service_instances"`
	// 
	ServiceType string `json:"service_type"`
	// Services related to the host
	Services []ApiHostGet200ResponseServicesInner `json:"services"`
	Sources ApiHostGet200ResponseSources `json:"sources"`
	// 
	Status string `json:"status"`
}

type _ApiHostGet200Response ApiHostGet200Response

// NewApiHostGet200Response instantiates a new ApiHostGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHostGet200Response(addr string, cephVersion string, hostname string, labels []string, serviceInstances []ApiHostGet200ResponseServiceInstancesInner, serviceType string, services []ApiHostGet200ResponseServicesInner, sources ApiHostGet200ResponseSources, status string) *ApiHostGet200Response {
	this := ApiHostGet200Response{}
	this.Addr = addr
	this.CephVersion = cephVersion
	this.Hostname = hostname
	this.Labels = labels
	this.ServiceInstances = serviceInstances
	this.ServiceType = serviceType
	this.Services = services
	this.Sources = sources
	this.Status = status
	return &this
}

// NewApiHostGet200ResponseWithDefaults instantiates a new ApiHostGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHostGet200ResponseWithDefaults() *ApiHostGet200Response {
	this := ApiHostGet200Response{}
	return &this
}

// GetAddr returns the Addr field value
func (o *ApiHostGet200Response) GetAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Addr
}

// GetAddrOk returns a tuple with the Addr field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200Response) GetAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addr, true
}

// SetAddr sets field value
func (o *ApiHostGet200Response) SetAddr(v string) {
	o.Addr = v
}

// GetCephVersion returns the CephVersion field value
func (o *ApiHostGet200Response) GetCephVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CephVersion
}

// GetCephVersionOk returns a tuple with the CephVersion field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200Response) GetCephVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CephVersion, true
}

// SetCephVersion sets field value
func (o *ApiHostGet200Response) SetCephVersion(v string) {
	o.CephVersion = v
}

// GetHostname returns the Hostname field value
func (o *ApiHostGet200Response) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200Response) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *ApiHostGet200Response) SetHostname(v string) {
	o.Hostname = v
}

// GetLabels returns the Labels field value
func (o *ApiHostGet200Response) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200Response) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *ApiHostGet200Response) SetLabels(v []string) {
	o.Labels = v
}

// GetServiceInstances returns the ServiceInstances field value
func (o *ApiHostGet200Response) GetServiceInstances() []ApiHostGet200ResponseServiceInstancesInner {
	if o == nil {
		var ret []ApiHostGet200ResponseServiceInstancesInner
		return ret
	}

	return o.ServiceInstances
}

// GetServiceInstancesOk returns a tuple with the ServiceInstances field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200Response) GetServiceInstancesOk() ([]ApiHostGet200ResponseServiceInstancesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceInstances, true
}

// SetServiceInstances sets field value
func (o *ApiHostGet200Response) SetServiceInstances(v []ApiHostGet200ResponseServiceInstancesInner) {
	o.ServiceInstances = v
}

// GetServiceType returns the ServiceType field value
func (o *ApiHostGet200Response) GetServiceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200Response) GetServiceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *ApiHostGet200Response) SetServiceType(v string) {
	o.ServiceType = v
}

// GetServices returns the Services field value
func (o *ApiHostGet200Response) GetServices() []ApiHostGet200ResponseServicesInner {
	if o == nil {
		var ret []ApiHostGet200ResponseServicesInner
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200Response) GetServicesOk() ([]ApiHostGet200ResponseServicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Services, true
}

// SetServices sets field value
func (o *ApiHostGet200Response) SetServices(v []ApiHostGet200ResponseServicesInner) {
	o.Services = v
}

// GetSources returns the Sources field value
func (o *ApiHostGet200Response) GetSources() ApiHostGet200ResponseSources {
	if o == nil {
		var ret ApiHostGet200ResponseSources
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200Response) GetSourcesOk() (*ApiHostGet200ResponseSources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value
func (o *ApiHostGet200Response) SetSources(v ApiHostGet200ResponseSources) {
	o.Sources = v
}

// GetStatus returns the Status field value
func (o *ApiHostGet200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiHostGet200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiHostGet200Response) SetStatus(v string) {
	o.Status = v
}

func (o ApiHostGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHostGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addr"] = o.Addr
	toSerialize["ceph_version"] = o.CephVersion
	toSerialize["hostname"] = o.Hostname
	toSerialize["labels"] = o.Labels
	toSerialize["service_instances"] = o.ServiceInstances
	toSerialize["service_type"] = o.ServiceType
	toSerialize["services"] = o.Services
	toSerialize["sources"] = o.Sources
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ApiHostGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addr",
		"ceph_version",
		"hostname",
		"labels",
		"service_instances",
		"service_type",
		"services",
		"sources",
		"status",
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

	varApiHostGet200Response := _ApiHostGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiHostGet200Response)

	if err != nil {
		return err
	}

	*o = ApiHostGet200Response(varApiHostGet200Response)

	return err
}

type NullableApiHostGet200Response struct {
	value *ApiHostGet200Response
	isSet bool
}

func (v NullableApiHostGet200Response) Get() *ApiHostGet200Response {
	return v.value
}

func (v *NullableApiHostGet200Response) Set(val *ApiHostGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHostGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHostGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHostGet200Response(val *ApiHostGet200Response) *NullableApiHostGet200Response {
	return &NullableApiHostGet200Response{value: val, isSet: true}
}

func (v NullableApiHostGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHostGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

