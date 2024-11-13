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

// checks if the ApiMultiClusterReconnectClusterPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMultiClusterReconnectClusterPutRequest{}

// ApiMultiClusterReconnectClusterPutRequest struct for ApiMultiClusterReconnectClusterPutRequest
type ApiMultiClusterReconnectClusterPutRequest struct {
	ClusterToken *string `json:"cluster_token,omitempty"`
	Password *string `json:"password,omitempty"`
	SslCertificate *string `json:"ssl_certificate,omitempty"`
	SslVerify *bool `json:"ssl_verify,omitempty"`
	Ttl *string `json:"ttl,omitempty"`
	Url string `json:"url"`
	Username *string `json:"username,omitempty"`
}

type _ApiMultiClusterReconnectClusterPutRequest ApiMultiClusterReconnectClusterPutRequest

// NewApiMultiClusterReconnectClusterPutRequest instantiates a new ApiMultiClusterReconnectClusterPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMultiClusterReconnectClusterPutRequest(url string) *ApiMultiClusterReconnectClusterPutRequest {
	this := ApiMultiClusterReconnectClusterPutRequest{}
	var sslVerify bool = false
	this.SslVerify = &sslVerify
	this.Url = url
	return &this
}

// NewApiMultiClusterReconnectClusterPutRequestWithDefaults instantiates a new ApiMultiClusterReconnectClusterPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMultiClusterReconnectClusterPutRequestWithDefaults() *ApiMultiClusterReconnectClusterPutRequest {
	this := ApiMultiClusterReconnectClusterPutRequest{}
	var sslVerify bool = false
	this.SslVerify = &sslVerify
	return &this
}

// GetClusterToken returns the ClusterToken field value if set, zero value otherwise.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetClusterToken() string {
	if o == nil || IsNil(o.ClusterToken) {
		var ret string
		return ret
	}
	return *o.ClusterToken
}

// GetClusterTokenOk returns a tuple with the ClusterToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetClusterTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterToken) {
		return nil, false
	}
	return o.ClusterToken, true
}

// HasClusterToken returns a boolean if a field has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) HasClusterToken() bool {
	if o != nil && !IsNil(o.ClusterToken) {
		return true
	}

	return false
}

// SetClusterToken gets a reference to the given string and assigns it to the ClusterToken field.
func (o *ApiMultiClusterReconnectClusterPutRequest) SetClusterToken(v string) {
	o.ClusterToken = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiMultiClusterReconnectClusterPutRequest) SetPassword(v string) {
	o.Password = &v
}

// GetSslCertificate returns the SslCertificate field value if set, zero value otherwise.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetSslCertificate() string {
	if o == nil || IsNil(o.SslCertificate) {
		var ret string
		return ret
	}
	return *o.SslCertificate
}

// GetSslCertificateOk returns a tuple with the SslCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetSslCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.SslCertificate) {
		return nil, false
	}
	return o.SslCertificate, true
}

// HasSslCertificate returns a boolean if a field has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) HasSslCertificate() bool {
	if o != nil && !IsNil(o.SslCertificate) {
		return true
	}

	return false
}

// SetSslCertificate gets a reference to the given string and assigns it to the SslCertificate field.
func (o *ApiMultiClusterReconnectClusterPutRequest) SetSslCertificate(v string) {
	o.SslCertificate = &v
}

// GetSslVerify returns the SslVerify field value if set, zero value otherwise.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetSslVerify() bool {
	if o == nil || IsNil(o.SslVerify) {
		var ret bool
		return ret
	}
	return *o.SslVerify
}

// GetSslVerifyOk returns a tuple with the SslVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetSslVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.SslVerify) {
		return nil, false
	}
	return o.SslVerify, true
}

// HasSslVerify returns a boolean if a field has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) HasSslVerify() bool {
	if o != nil && !IsNil(o.SslVerify) {
		return true
	}

	return false
}

// SetSslVerify gets a reference to the given bool and assigns it to the SslVerify field.
func (o *ApiMultiClusterReconnectClusterPutRequest) SetSslVerify(v bool) {
	o.SslVerify = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetTtl() string {
	if o == nil || IsNil(o.Ttl) {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetTtlOk() (*string, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *ApiMultiClusterReconnectClusterPutRequest) SetTtl(v string) {
	o.Ttl = &v
}

// GetUrl returns the Url field value
func (o *ApiMultiClusterReconnectClusterPutRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ApiMultiClusterReconnectClusterPutRequest) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiMultiClusterReconnectClusterPutRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiMultiClusterReconnectClusterPutRequest) SetUsername(v string) {
	o.Username = &v
}

func (o ApiMultiClusterReconnectClusterPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMultiClusterReconnectClusterPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterToken) {
		toSerialize["cluster_token"] = o.ClusterToken
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
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

func (o *ApiMultiClusterReconnectClusterPutRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varApiMultiClusterReconnectClusterPutRequest := _ApiMultiClusterReconnectClusterPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiMultiClusterReconnectClusterPutRequest)

	if err != nil {
		return err
	}

	*o = ApiMultiClusterReconnectClusterPutRequest(varApiMultiClusterReconnectClusterPutRequest)

	return err
}

type NullableApiMultiClusterReconnectClusterPutRequest struct {
	value *ApiMultiClusterReconnectClusterPutRequest
	isSet bool
}

func (v NullableApiMultiClusterReconnectClusterPutRequest) Get() *ApiMultiClusterReconnectClusterPutRequest {
	return v.value
}

func (v *NullableApiMultiClusterReconnectClusterPutRequest) Set(val *ApiMultiClusterReconnectClusterPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMultiClusterReconnectClusterPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMultiClusterReconnectClusterPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMultiClusterReconnectClusterPutRequest(val *ApiMultiClusterReconnectClusterPutRequest) *NullableApiMultiClusterReconnectClusterPutRequest {
	return &NullableApiMultiClusterReconnectClusterPutRequest{value: val, isSet: true}
}

func (v NullableApiMultiClusterReconnectClusterPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMultiClusterReconnectClusterPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


