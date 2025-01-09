# ApiRgwUserUidCapabilityPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaemonName** | Pointer to **string** |  | [optional] 
**Perm** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewApiRgwUserUidCapabilityPostRequest

`func NewApiRgwUserUidCapabilityPostRequest(perm string, type_ string, ) *ApiRgwUserUidCapabilityPostRequest`

NewApiRgwUserUidCapabilityPostRequest instantiates a new ApiRgwUserUidCapabilityPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwUserUidCapabilityPostRequestWithDefaults

`func NewApiRgwUserUidCapabilityPostRequestWithDefaults() *ApiRgwUserUidCapabilityPostRequest`

NewApiRgwUserUidCapabilityPostRequestWithDefaults instantiates a new ApiRgwUserUidCapabilityPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaemonName

`func (o *ApiRgwUserUidCapabilityPostRequest) GetDaemonName() string`

GetDaemonName returns the DaemonName field if non-nil, zero value otherwise.

### GetDaemonNameOk

`func (o *ApiRgwUserUidCapabilityPostRequest) GetDaemonNameOk() (*string, bool)`

GetDaemonNameOk returns a tuple with the DaemonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonName

`func (o *ApiRgwUserUidCapabilityPostRequest) SetDaemonName(v string)`

SetDaemonName sets DaemonName field to given value.

### HasDaemonName

`func (o *ApiRgwUserUidCapabilityPostRequest) HasDaemonName() bool`

HasDaemonName returns a boolean if a field has been set.

### GetPerm

`func (o *ApiRgwUserUidCapabilityPostRequest) GetPerm() string`

GetPerm returns the Perm field if non-nil, zero value otherwise.

### GetPermOk

`func (o *ApiRgwUserUidCapabilityPostRequest) GetPermOk() (*string, bool)`

GetPermOk returns a tuple with the Perm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerm

`func (o *ApiRgwUserUidCapabilityPostRequest) SetPerm(v string)`

SetPerm sets Perm field to given value.


### GetType

`func (o *ApiRgwUserUidCapabilityPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiRgwUserUidCapabilityPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiRgwUserUidCapabilityPostRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


