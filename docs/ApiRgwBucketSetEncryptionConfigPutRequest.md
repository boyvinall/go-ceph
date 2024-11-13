# ApiRgwBucketSetEncryptionConfigPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AuthMethod** | Pointer to **string** |  | [optional] 
**ClientCert** | Pointer to **string** |  | [optional] 
**ClientKey** | Pointer to **string** |  | [optional] 
**DaemonName** | Pointer to **string** |  | [optional] 
**EncryptionType** | Pointer to **string** |  | [optional] 
**KmsProvider** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] [default to ""]
**Owner** | Pointer to **string** |  | [optional] 
**SecretEngine** | Pointer to **string** |  | [optional] 
**SecretPath** | Pointer to **string** |  | [optional] [default to ""]
**SslCert** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewApiRgwBucketSetEncryptionConfigPutRequest

`func NewApiRgwBucketSetEncryptionConfigPutRequest() *ApiRgwBucketSetEncryptionConfigPutRequest`

NewApiRgwBucketSetEncryptionConfigPutRequest instantiates a new ApiRgwBucketSetEncryptionConfigPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwBucketSetEncryptionConfigPutRequestWithDefaults

`func NewApiRgwBucketSetEncryptionConfigPutRequestWithDefaults() *ApiRgwBucketSetEncryptionConfigPutRequest`

NewApiRgwBucketSetEncryptionConfigPutRequestWithDefaults instantiates a new ApiRgwBucketSetEncryptionConfigPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAuthMethod

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetClientCert

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetClientKey

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetDaemonName

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetDaemonName() string`

GetDaemonName returns the DaemonName field if non-nil, zero value otherwise.

### GetDaemonNameOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetDaemonNameOk() (*string, bool)`

GetDaemonNameOk returns a tuple with the DaemonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonName

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetDaemonName(v string)`

SetDaemonName sets DaemonName field to given value.

### HasDaemonName

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasDaemonName() bool`

HasDaemonName returns a boolean if a field has been set.

### GetEncryptionType

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.

### GetKmsProvider

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetKmsProvider() string`

GetKmsProvider returns the KmsProvider field if non-nil, zero value otherwise.

### GetKmsProviderOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetKmsProviderOk() (*string, bool)`

GetKmsProviderOk returns a tuple with the KmsProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsProvider

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetKmsProvider(v string)`

SetKmsProvider sets KmsProvider field to given value.

### HasKmsProvider

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasKmsProvider() bool`

HasKmsProvider returns a boolean if a field has been set.

### GetNamespace

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOwner

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSecretEngine

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetSecretEngine() string`

GetSecretEngine returns the SecretEngine field if non-nil, zero value otherwise.

### GetSecretEngineOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetSecretEngineOk() (*string, bool)`

GetSecretEngineOk returns a tuple with the SecretEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretEngine

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetSecretEngine(v string)`

SetSecretEngine sets SecretEngine field to given value.

### HasSecretEngine

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasSecretEngine() bool`

HasSecretEngine returns a boolean if a field has been set.

### GetSecretPath

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.

### HasSecretPath

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasSecretPath() bool`

HasSecretPath returns a boolean if a field has been set.

### GetSslCert

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### GetToken

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApiRgwBucketSetEncryptionConfigPutRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


