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

// checks if the ApiMultiClusterAuthPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMultiClusterAuthPostRequest{}

// ApiMultiClusterAuthPostRequest struct for ApiMultiClusterAuthPostRequest
type ApiMultiClusterAuthPostRequest struct {
	ClusterAlias string `json:"cluster_alias"`
	HubUrl *string `json:"hub_url,omitempty"`
	Password *string `json:"password,omitempty"`
	SslCertificate *string `json:"ssl_certificate,omitempty"`
	SslVerify *bool `json:"ssl_verify,omitempty"`
	Ttl *string `json:"ttl,omitempty"`
	Url string `json:"url"`
	Username string `json:"username"`
}

type _ApiMultiClusterAuthPostRequest ApiMultiClusterAuthPostRequest

// NewApiMultiClusterAuthPostRequest instantiates a new ApiMultiClusterAuthPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMultiClusterAuthPostRequest(clusterAlias string, url string, username string) *ApiMultiClusterAuthPostRequest {
	this := ApiMultiClusterAuthPostRequest{}
	this.ClusterAlias = clusterAlias
	var sslVerify bool = false
	this.SslVerify = &sslVerify
	this.Url = url
	this.Username = username
	return &this
}

// NewApiMultiClusterAuthPostRequestWithDefaults instantiates a new ApiMultiClusterAuthPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMultiClusterAuthPostRequestWithDefaults() *ApiMultiClusterAuthPostRequest {
	this := ApiMultiClusterAuthPostRequest{}
	var sslVerify bool = false
	this.SslVerify = &sslVerify
	return &this
}

// GetClusterAlias returns the ClusterAlias field value
func (o *ApiMultiClusterAuthPostRequest) GetClusterAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterAlias
}

// GetClusterAliasOk returns a tuple with the ClusterAlias field value
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterAuthPostRequest) GetClusterAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterAlias, true
}

// SetClusterAlias sets field value
func (o *ApiMultiClusterAuthPostRequest) SetClusterAlias(v string) {
	o.ClusterAlias = v
}

// GetHubUrl returns the HubUrl field value if set, zero value otherwise.
func (o *ApiMultiClusterAuthPostRequest) GetHubUrl() string {
	if o == nil || IsNil(o.HubUrl) {
		var ret string
		return ret
	}
	return *o.HubUrl
}

// GetHubUrlOk returns a tuple with the HubUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterAuthPostRequest) GetHubUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HubUrl) {
		return nil, false
	}
	return o.HubUrl, true
}

// HasHubUrl returns a boolean if a field has been set.
func (o *ApiMultiClusterAuthPostRequest) HasHubUrl() bool {
	if o != nil && !IsNil(o.HubUrl) {
		return true
	}

	return false
}

// SetHubUrl gets a reference to the given string and assigns it to the HubUrl field.
func (o *ApiMultiClusterAuthPostRequest) SetHubUrl(v string) {
	o.HubUrl = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiMultiClusterAuthPostRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterAuthPostRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiMultiClusterAuthPostRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiMultiClusterAuthPostRequest) SetPassword(v string) {
	o.Password = &v
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise.
func (o *ApiMultiClusterAuthPostRequest) GetSslCertificate() string {
	if o == nil || IsNil(o.SslCertificate) {
		var ret string
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterAuthPostRequest) GetSslCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.SslCertificate) {
		return nil, false
	}
	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *ApiMultiClusterAuthPostRequest) HasSslCertificate() bool {
	if o != nil && !IsNil(o.SslCertificate) {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given string and assigns it to the SslCertificate field.
func (o *ApiMultiClusterAuthPostRequest) SetSslCertificate(v string) {
	o.SslCertificate = &v
}

// GetSslVerify returns the SslVerify field value if set, zero value otherwise.
func (o *ApiMultiClusterAuthPostRequest) GetSslVerify() bool {
	if o == nil || IsNil(o.SslVerify) {
		var ret bool
		return ret
	}
	return *o.SslVerify
}

// GetSslVerifyOk returns a tuple with the SslVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterAuthPostRequest) GetSslVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.SslVerify) {
		return nil, false
	}
	return o.SslVerify, true
}

// HasSslVerify returns a boolean if a field has been set.
func (o *ApiMultiClusterAuthPostRequest) HasSslVerify() bool {
	if o != nil && !IsNil(o.SslVerify) {
		return true
	}

	return false
}

// SetSslVerify gets a reference to the given bool and assigns it to the SslVerify field.
func (o *ApiMultiClusterAuthPostRequest) SetSslVerify(v bool) {
	o.SslVerify = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ApiMultiClusterAuthPostRequest) GetTtl() string {
	if o == nil || IsNil(o.Ttl) {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterAuthPostRequest) GetTtlOk() (*string, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ApiMultiClusterAuthPostRequest) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *ApiMultiClusterAuthPostRequest) SetTtl(v string) {
	o.Ttl = &v
}

// GetUrl returns the Url field value
func (o *ApiMultiClusterAuthPostRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterAuthPostRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ApiMultiClusterAuthPostRequest) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value
func (o *ApiMultiClusterAuthPostRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterAuthPostRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ApiMultiClusterAuthPostRequest) SetUsername(v string) {
	o.Username = v
}

func (o ApiMultiClusterAuthPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMultiClusterAuthPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster_alias"] = o.ClusterAlias
	if !IsNil(o.HubUrl) {
		toSerialize["hub_url"] = o.HubUrl
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.SslCertificate) {
		toSerialize["ssl_certificate"] = o.SslCertificate
	}
	if !IsNil(o.SslVerify) {
		toSerialize["ssl_verify"] = o.SslVerify
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	toSerialize["url"] = o.Url
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *ApiMultiClusterAuthPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cluster_alias",
		"url",
		"username",
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

	varApiMultiClusterAuthPostRequest := _ApiMultiClusterAuthPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiMultiClusterAuthPostRequest)

	if err != nil {
		return err
	}

	*o = ApiMultiClusterAuthPostRequest(varApiMultiClusterAuthPostRequest)

	return err
}

type NullableApiMultiClusterAuthPostRequest struct {
	value *ApiMultiClusterAuthPostRequest
	isSet bool
}

func (v NullableApiMultiClusterAuthPostRequest) Get() *ApiMultiClusterAuthPostRequest {
	return v.value
}

func (v *NullableApiMultiClusterAuthPostRequest) Set(val *ApiMultiClusterAuthPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMultiClusterAuthPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMultiClusterAuthPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMultiClusterAuthPostRequest(val *ApiMultiClusterAuthPostRequest) *NullableApiMultiClusterAuthPostRequest {
	return &NullableApiMultiClusterAuthPostRequest{value: val, isSet: true}
}

func (v NullableApiMultiClusterAuthPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMultiClusterAuthPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


