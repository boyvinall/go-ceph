# ApiMultiClusterEditClusterPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterAlias** | **string** |  | 
**Name** | **string** |  | 
**SslCertificate** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Username** | **string** |  | 
**Verify** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewApiMultiClusterEditClusterPutRequest

`func NewApiMultiClusterEditClusterPutRequest(clusterAlias string, name string, url string, username string, ) *ApiMultiClusterEditClusterPutRequest`

NewApiMultiClusterEditClusterPutRequest instantiates a new ApiMultiClusterEditClusterPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMultiClusterEditClusterPutRequestWithDefaults

`func NewApiMultiClusterEditClusterPutRequestWithDefaults() *ApiMultiClusterEditClusterPutRequest`

NewApiMultiClusterEditClusterPutRequestWithDefaults instantiates a new ApiMultiClusterEditClusterPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterAlias

`func (o *ApiMultiClusterEditClusterPutRequest) GetClusterAlias() string`

GetClusterAlias returns the ClusterAlias field if non-nil, zero value otherwise.

### GetClusterAliasOk

`func (o *ApiMultiClusterEditClusterPutRequest) GetClusterAliasOk() (*string, bool)`

GetClusterAliasOk returns a tuple with the ClusterAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAlias

`func (o *ApiMultiClusterEditClusterPutRequest) SetClusterAlias(v string)`

SetClusterAlias sets ClusterAlias field to given value.


### GetName

`func (o *ApiMultiClusterEditClusterPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMultiClusterEditClusterPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMultiClusterEditClusterPutRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSslCertificate

`func (o *ApiMultiClusterEditClusterPutRequest) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *ApiMultiClusterEditClusterPutRequest) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *ApiMultiClusterEditClusterPutRequest) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *ApiMultiClusterEditClusterPutRequest) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetUrl

`func (o *ApiMultiClusterEditClusterPutRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiMultiClusterEditClusterPutRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiMultiClusterEditClusterPutRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *ApiMultiClusterEditClusterPutRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiMultiClusterEditClusterPutRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiMultiClusterEditClusterPutRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetVerify

`func (o *ApiMultiClusterEditClusterPutRequest) GetVerify() bool`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *ApiMultiClusterEditClusterPutRequest) GetVerifyOk() (*bool, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *ApiMultiClusterEditClusterPutRequest) SetVerify(v bool)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *ApiMultiClusterEditClusterPutRequest) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


