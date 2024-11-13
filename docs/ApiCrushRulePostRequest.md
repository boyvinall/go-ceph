# ApiCrushRulePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceClass** | Pointer to **string** |  | [optional] 
**FailureDomain** | **string** |  | 
**Name** | **string** |  | 
**PoolType** | Pointer to **string** |  | [optional] [default to "replication"]
**Profile** | Pointer to **string** |  | [optional] 
**Root** | Pointer to **string** |  | [optional] 

## Methods

### NewApiCrushRulePostRequest

`func NewApiCrushRulePostRequest(failureDomain string, name string, ) *ApiCrushRulePostRequest`

NewApiCrushRulePostRequest instantiates a new ApiCrushRulePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCrushRulePostRequestWithDefaults

`func NewApiCrushRulePostRequestWithDefaults() *ApiCrushRulePostRequest`

NewApiCrushRulePostRequestWithDefaults instantiates a new ApiCrushRulePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceClass

`func (o *ApiCrushRulePostRequest) GetDeviceClass() string`

GetDeviceClass returns the DeviceClass field if non-nil, zero value otherwise.

### GetDeviceClassOk

`func (o *ApiCrushRulePostRequest) GetDeviceClassOk() (*string, bool)`

GetDeviceClassOk returns a tuple with the DeviceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClass

`func (o *ApiCrushRulePostRequest) SetDeviceClass(v string)`

SetDeviceClass sets DeviceClass field to given value.

### HasDeviceClass

`func (o *ApiCrushRulePostRequest) HasDeviceClass() bool`

HasDeviceClass returns a boolean if a field has been set.

### GetFailureDomain

`func (o *ApiCrushRulePostRequest) GetFailureDomain() string`

GetFailureDomain returns the FailureDomain field if non-nil, zero value otherwise.

### GetFailureDomainOk

`func (o *ApiCrushRulePostRequest) GetFailureDomainOk() (*string, bool)`

GetFailureDomainOk returns a tuple with the FailureDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDomain

`func (o *ApiCrushRulePostRequest) SetFailureDomain(v string)`

SetFailureDomain sets FailureDomain field to given value.


### GetName

`func (o *ApiCrushRulePostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiCrushRulePostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiCrushRulePostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPoolType

`func (o *ApiCrushRulePostRequest) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *ApiCrushRulePostRequest) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *ApiCrushRulePostRequest) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *ApiCrushRulePostRequest) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetProfile

`func (o *ApiCrushRulePostRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ApiCrushRulePostRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ApiCrushRulePostRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ApiCrushRulePostRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRoot

`func (o *ApiCrushRulePostRequest) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *ApiCrushRulePostRequest) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *ApiCrushRulePostRequest) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *ApiCrushRulePostRequest) HasRoot() bool`

HasRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


