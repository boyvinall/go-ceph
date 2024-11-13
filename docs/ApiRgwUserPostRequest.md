# ApiRgwUserPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**DaemonName** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**GenerateKey** | Pointer to **string** |  | [optional] 
**MaxBuckets** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**Suspended** | Pointer to **string** |  | [optional] 
**System** | Pointer to **string** |  | [optional] 
**Uid** | **string** |  | 

## Methods

### NewApiRgwUserPostRequest

`func NewApiRgwUserPostRequest(displayName string, uid string, ) *ApiRgwUserPostRequest`

NewApiRgwUserPostRequest instantiates a new ApiRgwUserPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwUserPostRequestWithDefaults

`func NewApiRgwUserPostRequestWithDefaults() *ApiRgwUserPostRequest`

NewApiRgwUserPostRequestWithDefaults instantiates a new ApiRgwUserPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *ApiRgwUserPostRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ApiRgwUserPostRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ApiRgwUserPostRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ApiRgwUserPostRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetDaemonName

`func (o *ApiRgwUserPostRequest) GetDaemonName() string`

GetDaemonName returns the DaemonName field if non-nil, zero value otherwise.

### GetDaemonNameOk

`func (o *ApiRgwUserPostRequest) GetDaemonNameOk() (*string, bool)`

GetDaemonNameOk returns a tuple with the DaemonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonName

`func (o *ApiRgwUserPostRequest) SetDaemonName(v string)`

SetDaemonName sets DaemonName field to given value.

### HasDaemonName

`func (o *ApiRgwUserPostRequest) HasDaemonName() bool`

HasDaemonName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiRgwUserPostRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiRgwUserPostRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiRgwUserPostRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmail

`func (o *ApiRgwUserPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApiRgwUserPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApiRgwUserPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApiRgwUserPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGenerateKey

`func (o *ApiRgwUserPostRequest) GetGenerateKey() string`

GetGenerateKey returns the GenerateKey field if non-nil, zero value otherwise.

### GetGenerateKeyOk

`func (o *ApiRgwUserPostRequest) GetGenerateKeyOk() (*string, bool)`

GetGenerateKeyOk returns a tuple with the GenerateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateKey

`func (o *ApiRgwUserPostRequest) SetGenerateKey(v string)`

SetGenerateKey sets GenerateKey field to given value.

### HasGenerateKey

`func (o *ApiRgwUserPostRequest) HasGenerateKey() bool`

HasGenerateKey returns a boolean if a field has been set.

### GetMaxBuckets

`func (o *ApiRgwUserPostRequest) GetMaxBuckets() string`

GetMaxBuckets returns the MaxBuckets field if non-nil, zero value otherwise.

### GetMaxBucketsOk

`func (o *ApiRgwUserPostRequest) GetMaxBucketsOk() (*string, bool)`

GetMaxBucketsOk returns a tuple with the MaxBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBuckets

`func (o *ApiRgwUserPostRequest) SetMaxBuckets(v string)`

SetMaxBuckets sets MaxBuckets field to given value.

### HasMaxBuckets

`func (o *ApiRgwUserPostRequest) HasMaxBuckets() bool`

HasMaxBuckets returns a boolean if a field has been set.

### GetSecretKey

`func (o *ApiRgwUserPostRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ApiRgwUserPostRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ApiRgwUserPostRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ApiRgwUserPostRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSuspended

`func (o *ApiRgwUserPostRequest) GetSuspended() string`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *ApiRgwUserPostRequest) GetSuspendedOk() (*string, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *ApiRgwUserPostRequest) SetSuspended(v string)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *ApiRgwUserPostRequest) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystem

`func (o *ApiRgwUserPostRequest) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ApiRgwUserPostRequest) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ApiRgwUserPostRequest) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ApiRgwUserPostRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetUid

`func (o *ApiRgwUserPostRequest) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ApiRgwUserPostRequest) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ApiRgwUserPostRequest) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


