# ApiRgwUserUidSubuserPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** |  | 
**AccessKey** | Pointer to **string** |  | [optional] 
**DaemonName** | Pointer to **string** |  | [optional] 
**GenerateSecret** | Pointer to **string** |  | [optional] [default to "true"]
**KeyType** | Pointer to **string** |  | [optional] [default to "s3"]
**SecretKey** | Pointer to **string** |  | [optional] 
**Subuser** | **string** |  | 

## Methods

### NewApiRgwUserUidSubuserPostRequest

`func NewApiRgwUserUidSubuserPostRequest(access string, subuser string, ) *ApiRgwUserUidSubuserPostRequest`

NewApiRgwUserUidSubuserPostRequest instantiates a new ApiRgwUserUidSubuserPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwUserUidSubuserPostRequestWithDefaults

`func NewApiRgwUserUidSubuserPostRequestWithDefaults() *ApiRgwUserUidSubuserPostRequest`

NewApiRgwUserUidSubuserPostRequestWithDefaults instantiates a new ApiRgwUserUidSubuserPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ApiRgwUserUidSubuserPostRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ApiRgwUserUidSubuserPostRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ApiRgwUserUidSubuserPostRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAccessKey

`func (o *ApiRgwUserUidSubuserPostRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ApiRgwUserUidSubuserPostRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ApiRgwUserUidSubuserPostRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ApiRgwUserUidSubuserPostRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetDaemonName

`func (o *ApiRgwUserUidSubuserPostRequest) GetDaemonName() string`

GetDaemonName returns the DaemonName field if non-nil, zero value otherwise.

### GetDaemonNameOk

`func (o *ApiRgwUserUidSubuserPostRequest) GetDaemonNameOk() (*string, bool)`

GetDaemonNameOk returns a tuple with the DaemonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonName

`func (o *ApiRgwUserUidSubuserPostRequest) SetDaemonName(v string)`

SetDaemonName sets DaemonName field to given value.

### HasDaemonName

`func (o *ApiRgwUserUidSubuserPostRequest) HasDaemonName() bool`

HasDaemonName returns a boolean if a field has been set.

### GetGenerateSecret

`func (o *ApiRgwUserUidSubuserPostRequest) GetGenerateSecret() string`

GetGenerateSecret returns the GenerateSecret field if non-nil, zero value otherwise.

### GetGenerateSecretOk

`func (o *ApiRgwUserUidSubuserPostRequest) GetGenerateSecretOk() (*string, bool)`

GetGenerateSecretOk returns a tuple with the GenerateSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSecret

`func (o *ApiRgwUserUidSubuserPostRequest) SetGenerateSecret(v string)`

SetGenerateSecret sets GenerateSecret field to given value.

### HasGenerateSecret

`func (o *ApiRgwUserUidSubuserPostRequest) HasGenerateSecret() bool`

HasGenerateSecret returns a boolean if a field has been set.

### GetKeyType

`func (o *ApiRgwUserUidSubuserPostRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ApiRgwUserUidSubuserPostRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ApiRgwUserUidSubuserPostRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ApiRgwUserUidSubuserPostRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetSecretKey

`func (o *ApiRgwUserUidSubuserPostRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ApiRgwUserUidSubuserPostRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ApiRgwUserUidSubuserPostRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ApiRgwUserUidSubuserPostRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSubuser

`func (o *ApiRgwUserUidSubuserPostRequest) GetSubuser() string`

GetSubuser returns the Subuser field if non-nil, zero value otherwise.

### GetSubuserOk

`func (o *ApiRgwUserUidSubuserPostRequest) GetSubuserOk() (*string, bool)`

GetSubuserOk returns a tuple with the Subuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubuser

`func (o *ApiRgwUserUidSubuserPostRequest) SetSubuser(v string)`

SetSubuser sets Subuser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


