# ApiCephfsAuthPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caps** | **string** | Path and given capabilities | 
**ClientId** | **string** | Cephx user ID | 
**FsName** | **string** | File system name | 
**RootSquash** | **string** | File System Identifier | 

## Methods

### NewApiCephfsAuthPutRequest

`func NewApiCephfsAuthPutRequest(caps string, clientId string, fsName string, rootSquash string, ) *ApiCephfsAuthPutRequest`

NewApiCephfsAuthPutRequest instantiates a new ApiCephfsAuthPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCephfsAuthPutRequestWithDefaults

`func NewApiCephfsAuthPutRequestWithDefaults() *ApiCephfsAuthPutRequest`

NewApiCephfsAuthPutRequestWithDefaults instantiates a new ApiCephfsAuthPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaps

`func (o *ApiCephfsAuthPutRequest) GetCaps() string`

GetCaps returns the Caps field if non-nil, zero value otherwise.

### GetCapsOk

`func (o *ApiCephfsAuthPutRequest) GetCapsOk() (*string, bool)`

GetCapsOk returns a tuple with the Caps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaps

`func (o *ApiCephfsAuthPutRequest) SetCaps(v string)`

SetCaps sets Caps field to given value.


### GetClientId

`func (o *ApiCephfsAuthPutRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApiCephfsAuthPutRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApiCephfsAuthPutRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetFsName

`func (o *ApiCephfsAuthPutRequest) GetFsName() string`

GetFsName returns the FsName field if non-nil, zero value otherwise.

### GetFsNameOk

`func (o *ApiCephfsAuthPutRequest) GetFsNameOk() (*string, bool)`

GetFsNameOk returns a tuple with the FsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsName

`func (o *ApiCephfsAuthPutRequest) SetFsName(v string)`

SetFsName sets FsName field to given value.


### GetRootSquash

`func (o *ApiCephfsAuthPutRequest) GetRootSquash() string`

GetRootSquash returns the RootSquash field if non-nil, zero value otherwise.

### GetRootSquashOk

`func (o *ApiCephfsAuthPutRequest) GetRootSquashOk() (*string, bool)`

GetRootSquashOk returns a tuple with the RootSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSquash

`func (o *ApiCephfsAuthPutRequest) SetRootSquash(v string)`

SetRootSquash sets RootSquash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


