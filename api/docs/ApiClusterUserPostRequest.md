# ApiClusterUserPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**[]ApiClusterUserPostRequestCapabilitiesInner**](ApiClusterUserPostRequestCapabilitiesInner.md) | List of capabilities to add to user_entity | [optional] 
**ImportData** | Pointer to **string** |  | [optional] [default to ""]
**UserEntity** | Pointer to **string** | Entity to add | [optional] [default to ""]

## Methods

### NewApiClusterUserPostRequest

`func NewApiClusterUserPostRequest() *ApiClusterUserPostRequest`

NewApiClusterUserPostRequest instantiates a new ApiClusterUserPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClusterUserPostRequestWithDefaults

`func NewApiClusterUserPostRequestWithDefaults() *ApiClusterUserPostRequest`

NewApiClusterUserPostRequestWithDefaults instantiates a new ApiClusterUserPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *ApiClusterUserPostRequest) GetCapabilities() []ApiClusterUserPostRequestCapabilitiesInner`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ApiClusterUserPostRequest) GetCapabilitiesOk() (*[]ApiClusterUserPostRequestCapabilitiesInner, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ApiClusterUserPostRequest) SetCapabilities(v []ApiClusterUserPostRequestCapabilitiesInner)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ApiClusterUserPostRequest) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetImportData

`func (o *ApiClusterUserPostRequest) GetImportData() string`

GetImportData returns the ImportData field if non-nil, zero value otherwise.

### GetImportDataOk

`func (o *ApiClusterUserPostRequest) GetImportDataOk() (*string, bool)`

GetImportDataOk returns a tuple with the ImportData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportData

`func (o *ApiClusterUserPostRequest) SetImportData(v string)`

SetImportData sets ImportData field to given value.

### HasImportData

`func (o *ApiClusterUserPostRequest) HasImportData() bool`

HasImportData returns a boolean if a field has been set.

### GetUserEntity

`func (o *ApiClusterUserPostRequest) GetUserEntity() string`

GetUserEntity returns the UserEntity field if non-nil, zero value otherwise.

### GetUserEntityOk

`func (o *ApiClusterUserPostRequest) GetUserEntityOk() (*string, bool)`

GetUserEntityOk returns a tuple with the UserEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEntity

`func (o *ApiClusterUserPostRequest) SetUserEntity(v string)`

SetUserEntity sets UserEntity field to given value.

### HasUserEntity

`func (o *ApiClusterUserPostRequest) HasUserEntity() bool`

HasUserEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


