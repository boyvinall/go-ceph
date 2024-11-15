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

// checks if the ApiIscsiTargetTargetIqnPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiIscsiTargetTargetIqnPutRequest{}

// ApiIscsiTargetTargetIqnPutRequest struct for ApiIscsiTargetTargetIqnPutRequest
type ApiIscsiTargetTargetIqnPutRequest struct {
	AclEnabled *string `json:"acl_enabled,omitempty"`
	Auth *string `json:"auth,omitempty"`
	Clients *string `json:"clients,omitempty"`
	Disks *string `json:"disks,omitempty"`
	Groups *string `json:"groups,omitempty"`
	NewTargetIqn *string `json:"new_target_iqn,omitempty"`
	Portals *string `json:"portals,omitempty"`
	TargetControls *string `json:"target_controls,omitempty"`
}

// NewApiIscsiTargetTargetIqnPutRequest instantiates a new ApiIscsiTargetTargetIqnPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiIscsiTargetTargetIqnPutRequest() *ApiIscsiTargetTargetIqnPutRequest {
	this := ApiIscsiTargetTargetIqnPutRequest{}
	return &this
}

// NewApiIscsiTargetTargetIqnPutRequestWithDefaults instantiates a new ApiIscsiTargetTargetIqnPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiIscsiTargetTargetIqnPutRequestWithDefaults() *ApiIscsiTargetTargetIqnPutRequest {
	this := ApiIscsiTargetTargetIqnPutRequest{}
	return &this
}

// GetAclEnabled returns the AclEnabled field value if set, zero value otherwise.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetAclEnabled() string {
	if o == nil || IsNil(o.AclEnabled) {
		var ret string
		return ret
	}
	return *o.AclEnabled
}

// GetAclEnabledOk returns a tuple with the AclEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetAclEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.AclEnabled) {
		return nil, false
	}
	return o.AclEnabled, true
}

// HasAclEnabled returns a boolean if a field has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) HasAclEnabled() bool {
	if o != nil && !IsNil(o.AclEnabled) {
		return true
	}

	return false
}

// SetAclEnabled gets a reference to the given string and assigns it to the AclEnabled field.
func (o *ApiIscsiTargetTargetIqnPutRequest) SetAclEnabled(v string) {
	o.AclEnabled = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetAuth() string {
	if o == nil || IsNil(o.Auth) {
		var ret string
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetAuthOk() (*string, bool) {
	if o == nil || IsNil(o.Auth) {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) HasAuth() bool {
	if o != nil && !IsNil(o.Auth) {
		return true
	}

	return false
}

// SetAuth gets a reference to the given string and assigns it to the Auth field.
func (o *ApiIscsiTargetTargetIqnPutRequest) SetAuth(v string) {
	o.Auth = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetClients() string {
	if o == nil || IsNil(o.Clients) {
		var ret string
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetClientsOk() (*string, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given string and assigns it to the Clients field.
func (o *ApiIscsiTargetTargetIqnPutRequest) SetClients(v string) {
	o.Clients = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetDisks() string {
	if o == nil || IsNil(o.Disks) {
		var ret string
		return ret
	}
	return *o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetDisksOk() (*string, bool) {
	if o == nil || IsNil(o.Disks) {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) HasDisks() bool {
	if o != nil && !IsNil(o.Disks) {
		return true
	}

	return false
}

// SetDisks gets a reference to the given string and assigns it to the Disks field.
func (o *ApiIscsiTargetTargetIqnPutRequest) SetDisks(v string) {
	o.Disks = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetGroups() string {
	if o == nil || IsNil(o.Groups) {
		var ret string
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetGroupsOk() (*string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given string and assigns it to the Groups field.
func (o *ApiIscsiTargetTargetIqnPutRequest) SetGroups(v string) {
	o.Groups = &v
}

// GetNewTargetIqn returns the NewTargetIqn field value if set, zero value otherwise.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetNewTargetIqn() string {
	if o == nil || IsNil(o.NewTargetIqn) {
		var ret string
		return ret
	}
	return *o.NewTargetIqn
}

// GetNewTargetIqnOk returns a tuple with the NewTargetIqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetNewTargetIqnOk() (*string, bool) {
	if o == nil || IsNil(o.NewTargetIqn) {
		return nil, false
	}
	return o.NewTargetIqn, true
}

// HasNewTargetIqn returns a boolean if a field has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) HasNewTargetIqn() bool {
	if o != nil && !IsNil(o.NewTargetIqn) {
		return true
	}

	return false
}

// SetNewTargetIqn gets a reference to the given string and assigns it to the NewTargetIqn field.
func (o *ApiIscsiTargetTargetIqnPutRequest) SetNewTargetIqn(v string) {
	o.NewTargetIqn = &v
}

// GetPortals returns the Portals field value if set, zero value otherwise.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetPortals() string {
	if o == nil || IsNil(o.Portals) {
		var ret string
		return ret
	}
	return *o.Portals
}

// GetPortalsOk returns a tuple with the Portals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetPortalsOk() (*string, bool) {
	if o == nil || IsNil(o.Portals) {
		return nil, false
	}
	return o.Portals, true
}

// HasPortals returns a boolean if a field has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) HasPortals() bool {
	if o != nil && !IsNil(o.Portals) {
		return true
	}

	return false
}

// SetPortals gets a reference to the given string and assigns it to the Portals field.
func (o *ApiIscsiTargetTargetIqnPutRequest) SetPortals(v string) {
	o.Portals = &v
}

// GetTargetControls returns the TargetControls field value if set, zero value otherwise.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetTargetControls() string {
	if o == nil || IsNil(o.TargetControls) {
		var ret string
		return ret
	}
	return *o.TargetControls
}

// GetTargetControlsOk returns a tuple with the TargetControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) GetTargetControlsOk() (*string, bool) {
	if o == nil || IsNil(o.TargetControls) {
		return nil, false
	}
	return o.TargetControls, true
}

// HasTargetControls returns a boolean if a field has been set.
func (o *ApiIscsiTargetTargetIqnPutRequest) HasTargetControls() bool {
	if o != nil && !IsNil(o.TargetControls) {
		return true
	}

	return false
}

// SetTargetControls gets a reference to the given string and assigns it to the TargetControls field.
func (o *ApiIscsiTargetTargetIqnPutRequest) SetTargetControls(v string) {
	o.TargetControls = &v
}

func (o ApiIscsiTargetTargetIqnPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiIscsiTargetTargetIqnPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AclEnabled) {
		toSerialize["acl_enabled"] = o.AclEnabled
	}
	if !IsNil(o.Auth) {
		toSerialize["auth"] = o.Auth
	}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !IsNil(o.Disks) {
		toSerialize["disks"] = o.Disks
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.NewTargetIqn) {
		toSerialize["new_target_iqn"] = o.NewTargetIqn
	}
	if !IsNil(o.Portals) {
		toSerialize["portals"] = o.Portals
	}
	if !IsNil(o.TargetControls) {
		toSerialize["target_controls"] = o.TargetControls
	}
	return toSerialize, nil
}

type NullableApiIscsiTargetTargetIqnPutRequest struct {
	value *ApiIscsiTargetTargetIqnPutRequest
	isSet bool
}

func (v NullableApiIscsiTargetTargetIqnPutRequest) Get() *ApiIscsiTargetTargetIqnPutRequest {
	return v.value
}

func (v *NullableApiIscsiTargetTargetIqnPutRequest) Set(val *ApiIscsiTargetTargetIqnPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiIscsiTargetTargetIqnPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiIscsiTargetTargetIqnPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiIscsiTargetTargetIqnPutRequest(val *ApiIscsiTargetTargetIqnPutRequest) *NullableApiIscsiTargetTargetIqnPutRequest {
	return &NullableApiIscsiTargetTargetIqnPutRequest{value: val, isSet: true}
}

func (v NullableApiIscsiTargetTargetIqnPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiIscsiTargetTargetIqnPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


