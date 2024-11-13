# ApiMultiClusterAuthPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterAlias** | **string** |  | 
**HubUrl** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SslCertificate** | Pointer to **string** |  | [optional] 
**SslVerify** | Pointer to **bool** |  | [optional] [default to false]
**Ttl** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Username** | **string** |  | 

## Methods

### NewApiMultiClusterAuthPostRequest

`func NewApiMultiClusterAuthPostRequest(clusterAlias string, url string, username string, ) *ApiMultiClusterAuthPostRequest`

NewApiMultiClusterAuthPostRequest instantiates a new ApiMultiClusterAuthPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMultiClusterAuthPostRequestWithDefaults

`func NewApiMultiClusterAuthPostRequestWithDefaults() *ApiMultiClusterAuthPostRequest`

NewApiMultiClusterAuthPostRequestWithDefaults instantiates a new ApiMultiClusterAuthPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterAlias

`func (o *ApiMultiClusterAuthPostRequest) GetClusterAlias() string`

GetClusterAlias returns the ClusterAlias field if non-nil, zero value otherwise.

### GetClusterAliasOk

`func (o *ApiMultiClusterAuthPostRequest) GetClusterAliasOk() (*string, bool)`

GetClusterAliasOk returns a tuple with the ClusterAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAlias

`func (o *ApiMultiClusterAuthPostRequest) SetClusterAlias(v string)`

SetClusterAlias sets ClusterAlias field to given value.


### GetHubUrl

`func (o *ApiMultiClusterAuthPostRequest) GetHubUrl() string`

GetHubUrl returns the HubUrl field if non-nil, zero value otherwise.

### GetHubUrlOk

`func (o *ApiMultiClusterAuthPostRequest) GetHubUrlOk() (*string, bool)`

GetHubUrlOk returns a tuple with the HubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubUrl

`func (o *ApiMultiClusterAuthPostRequest) SetHubUrl(v string)`

SetHubUrl sets HubUrl field to given value.

### HasHubUrl

`func (o *ApiMultiClusterAuthPostRequest) HasHubUrl() bool`

HasHubUrl returns a boolean if a field has been set.

### GetPassword

`func (o *ApiMultiClusterAuthPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiMultiClusterAuthPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiMultiClusterAuthPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiMultiClusterAuthPostRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSslCertificate

`func (o *ApiMultiClusterAuthPostRequest) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *ApiMultiClusterAuthPostRequest) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *ApiMultiClusterAuthPostRequest) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *ApiMultiClusterAuthPostRequest) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetSslVerify

`func (o *ApiMultiClusterAuthPostRequest) GetSslVerify() bool`

GetSslVerify returns the SslVerify field if non-nil, zero value otherwise.

### GetSslVerifyOk

`func (o *ApiMultiClusterAuthPostRequest) GetSslVerifyOk() (*bool, bool)`

GetSslVerifyOk returns a tuple with the SslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerify

`func (o *ApiMultiClusterAuthPostRequest) SetSslVerify(v bool)`

SetSslVerify sets SslVerify field to given value.

### HasSslVerify

`func (o *ApiMultiClusterAuthPostRequest) HasSslVerify() bool`

HasSslVerify returns a boolean if a field has been set.

### GetTtl

`func (o *ApiMultiClusterAuthPostRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ApiMultiClusterAuthPostRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ApiMultiClusterAuthPostRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ApiMultiClusterAuthPostRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUrl

`func (o *ApiMultiClusterAuthPostRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiMultiClusterAuthPostRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiMultiClusterAuthPostRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *ApiMultiClusterAuthPostRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiMultiClusterAuthPostRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiMultiClusterAuthPostRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


