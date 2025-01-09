# ApiNvmeofSubsystemNqnListenerPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adrfam** | Pointer to **int32** | NVMeoF address family (0 - IPv4, 1 - IPv6) | [optional] [default to 0]
**GwGroup** | Pointer to **string** | NVMeoF gateway group | [optional] 
**HostName** | **string** | NVMeoF hostname | 
**Traddr** | **string** | NVMeoF transport address | 
**Trsvcid** | Pointer to **int32** | NVMeoF transport service port | [optional] [default to 4420]

## Methods

### NewApiNvmeofSubsystemNqnListenerPostRequest

`func NewApiNvmeofSubsystemNqnListenerPostRequest(hostName string, traddr string, ) *ApiNvmeofSubsystemNqnListenerPostRequest`

NewApiNvmeofSubsystemNqnListenerPostRequest instantiates a new ApiNvmeofSubsystemNqnListenerPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNvmeofSubsystemNqnListenerPostRequestWithDefaults

`func NewApiNvmeofSubsystemNqnListenerPostRequestWithDefaults() *ApiNvmeofSubsystemNqnListenerPostRequest`

NewApiNvmeofSubsystemNqnListenerPostRequestWithDefaults instantiates a new ApiNvmeofSubsystemNqnListenerPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdrfam

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) GetAdrfam() int32`

GetAdrfam returns the Adrfam field if non-nil, zero value otherwise.

### GetAdrfamOk

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) GetAdrfamOk() (*int32, bool)`

GetAdrfamOk returns a tuple with the Adrfam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrfam

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) SetAdrfam(v int32)`

SetAdrfam sets Adrfam field to given value.

### HasAdrfam

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) HasAdrfam() bool`

HasAdrfam returns a boolean if a field has been set.

### GetGwGroup

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) GetGwGroup() string`

GetGwGroup returns the GwGroup field if non-nil, zero value otherwise.

### GetGwGroupOk

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) GetGwGroupOk() (*string, bool)`

GetGwGroupOk returns a tuple with the GwGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwGroup

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) SetGwGroup(v string)`

SetGwGroup sets GwGroup field to given value.

### HasGwGroup

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) HasGwGroup() bool`

HasGwGroup returns a boolean if a field has been set.

### GetHostName

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetTraddr

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) GetTraddr() string`

GetTraddr returns the Traddr field if non-nil, zero value otherwise.

### GetTraddrOk

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) GetTraddrOk() (*string, bool)`

GetTraddrOk returns a tuple with the Traddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraddr

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) SetTraddr(v string)`

SetTraddr sets Traddr field to given value.


### GetTrsvcid

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) GetTrsvcid() int32`

GetTrsvcid returns the Trsvcid field if non-nil, zero value otherwise.

### GetTrsvcidOk

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) GetTrsvcidOk() (*int32, bool)`

GetTrsvcidOk returns a tuple with the Trsvcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrsvcid

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) SetTrsvcid(v int32)`

SetTrsvcid sets Trsvcid field to given value.

### HasTrsvcid

`func (o *ApiNvmeofSubsystemNqnListenerPostRequest) HasTrsvcid() bool`

HasTrsvcid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


