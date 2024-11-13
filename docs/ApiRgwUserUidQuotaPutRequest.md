# ApiRgwUserUidQuotaPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaemonName** | Pointer to **string** |  | [optional] 
**Enabled** | **string** |  | 
**MaxObjects** | **string** |  | 
**MaxSizeKb** | **int32** |  | 
**QuotaType** | **string** |  | 

## Methods

### NewApiRgwUserUidQuotaPutRequest

`func NewApiRgwUserUidQuotaPutRequest(enabled string, maxObjects string, maxSizeKb int32, quotaType string, ) *ApiRgwUserUidQuotaPutRequest`

NewApiRgwUserUidQuotaPutRequest instantiates a new ApiRgwUserUidQuotaPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwUserUidQuotaPutRequestWithDefaults

`func NewApiRgwUserUidQuotaPutRequestWithDefaults() *ApiRgwUserUidQuotaPutRequest`

NewApiRgwUserUidQuotaPutRequestWithDefaults instantiates a new ApiRgwUserUidQuotaPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaemonName

`func (o *ApiRgwUserUidQuotaPutRequest) GetDaemonName() string`

GetDaemonName returns the DaemonName field if non-nil, zero value otherwise.

### GetDaemonNameOk

`func (o *ApiRgwUserUidQuotaPutRequest) GetDaemonNameOk() (*string, bool)`

GetDaemonNameOk returns a tuple with the DaemonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonName

`func (o *ApiRgwUserUidQuotaPutRequest) SetDaemonName(v string)`

SetDaemonName sets DaemonName field to given value.

### HasDaemonName

`func (o *ApiRgwUserUidQuotaPutRequest) HasDaemonName() bool`

HasDaemonName returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiRgwUserUidQuotaPutRequest) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiRgwUserUidQuotaPutRequest) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiRgwUserUidQuotaPutRequest) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.


### GetMaxObjects

`func (o *ApiRgwUserUidQuotaPutRequest) GetMaxObjects() string`

GetMaxObjects returns the MaxObjects field if non-nil, zero value otherwise.

### GetMaxObjectsOk

`func (o *ApiRgwUserUidQuotaPutRequest) GetMaxObjectsOk() (*string, bool)`

GetMaxObjectsOk returns a tuple with the MaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxObjects

`func (o *ApiRgwUserUidQuotaPutRequest) SetMaxObjects(v string)`

SetMaxObjects sets MaxObjects field to given value.


### GetMaxSizeKb

`func (o *ApiRgwUserUidQuotaPutRequest) GetMaxSizeKb() int32`

GetMaxSizeKb returns the MaxSizeKb field if non-nil, zero value otherwise.

### GetMaxSizeKbOk

`func (o *ApiRgwUserUidQuotaPutRequest) GetMaxSizeKbOk() (*int32, bool)`

GetMaxSizeKbOk returns a tuple with the MaxSizeKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSizeKb

`func (o *ApiRgwUserUidQuotaPutRequest) SetMaxSizeKb(v int32)`

SetMaxSizeKb sets MaxSizeKb field to given value.


### GetQuotaType

`func (o *ApiRgwUserUidQuotaPutRequest) GetQuotaType() string`

GetQuotaType returns the QuotaType field if non-nil, zero value otherwise.

### GetQuotaTypeOk

`func (o *ApiRgwUserUidQuotaPutRequest) GetQuotaTypeOk() (*string, bool)`

GetQuotaTypeOk returns a tuple with the QuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaType

`func (o *ApiRgwUserUidQuotaPutRequest) SetQuotaType(v string)`

SetQuotaType sets QuotaType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


