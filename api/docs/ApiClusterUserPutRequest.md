# ApiClusterUserPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**[]ApiClusterUserPutRequestCapabilitiesInner**](ApiClusterUserPutRequestCapabilitiesInner.md) | List of updated capabilities to user_entity | [optional] 
**UserEntity** | Pointer to **string** | Entity to edit | [optional] [default to ""]

## Methods

### NewApiClusterUserPutRequest

`func NewApiClusterUserPutRequest() *ApiClusterUserPutRequest`

NewApiClusterUserPutRequest instantiates a new ApiClusterUserPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClusterUserPutRequestWithDefaults

`func NewApiClusterUserPutRequestWithDefaults() *ApiClusterUserPutRequest`

NewApiClusterUserPutRequestWithDefaults instantiates a new ApiClusterUserPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *ApiClusterUserPutRequest) GetCapabilities() []ApiClusterUserPutRequestCapabilitiesInner`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ApiClusterUserPutRequest) GetCapabilitiesOk() (*[]ApiClusterUserPutRequestCapabilitiesInner, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ApiClusterUserPutRequest) SetCapabilities(v []ApiClusterUserPutRequestCapabilitiesInner)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ApiClusterUserPutRequest) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetUserEntity

`func (o *ApiClusterUserPutRequest) GetUserEntity() string`

GetUserEntity returns the UserEntity field if non-nil, zero value otherwise.

### GetUserEntityOk

`func (o *ApiClusterUserPutRequest) GetUserEntityOk() (*string, bool)`

GetUserEntityOk returns a tuple with the UserEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEntity

`func (o *ApiClusterUserPutRequest) SetUserEntity(v string)`

SetUserEntity sets UserEntity field to given value.

### HasUserEntity

`func (o *ApiClusterUserPutRequest) HasUserEntity() bool`

HasUserEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


