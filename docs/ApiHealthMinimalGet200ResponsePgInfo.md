# ApiHealthMinimalGet200ResponsePgInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectStats** | [**ApiHealthMinimalGet200ResponsePgInfoObjectStats**](ApiHealthMinimalGet200ResponsePgInfoObjectStats.md) |  | 
**PgsPerOsd** | **int32** |  | 
**Statuses** | **string** |  | 

## Methods

### NewApiHealthMinimalGet200ResponsePgInfo

`func NewApiHealthMinimalGet200ResponsePgInfo(objectStats ApiHealthMinimalGet200ResponsePgInfoObjectStats, pgsPerOsd int32, statuses string, ) *ApiHealthMinimalGet200ResponsePgInfo`

NewApiHealthMinimalGet200ResponsePgInfo instantiates a new ApiHealthMinimalGet200ResponsePgInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHealthMinimalGet200ResponsePgInfoWithDefaults

`func NewApiHealthMinimalGet200ResponsePgInfoWithDefaults() *ApiHealthMinimalGet200ResponsePgInfo`

NewApiHealthMinimalGet200ResponsePgInfoWithDefaults instantiates a new ApiHealthMinimalGet200ResponsePgInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectStats

`func (o *ApiHealthMinimalGet200ResponsePgInfo) GetObjectStats() ApiHealthMinimalGet200ResponsePgInfoObjectStats`

GetObjectStats returns the ObjectStats field if non-nil, zero value otherwise.

### GetObjectStatsOk

`func (o *ApiHealthMinimalGet200ResponsePgInfo) GetObjectStatsOk() (*ApiHealthMinimalGet200ResponsePgInfoObjectStats, bool)`

GetObjectStatsOk returns a tuple with the ObjectStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStats

`func (o *ApiHealthMinimalGet200ResponsePgInfo) SetObjectStats(v ApiHealthMinimalGet200ResponsePgInfoObjectStats)`

SetObjectStats sets ObjectStats field to given value.


### GetPgsPerOsd

`func (o *ApiHealthMinimalGet200ResponsePgInfo) GetPgsPerOsd() int32`

GetPgsPerOsd returns the PgsPerOsd field if non-nil, zero value otherwise.

### GetPgsPerOsdOk

`func (o *ApiHealthMinimalGet200ResponsePgInfo) GetPgsPerOsdOk() (*int32, bool)`

GetPgsPerOsdOk returns a tuple with the PgsPerOsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgsPerOsd

`func (o *ApiHealthMinimalGet200ResponsePgInfo) SetPgsPerOsd(v int32)`

SetPgsPerOsd sets PgsPerOsd field to given value.


### GetStatuses

`func (o *ApiHealthMinimalGet200ResponsePgInfo) GetStatuses() string`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ApiHealthMinimalGet200ResponsePgInfo) GetStatusesOk() (*string, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ApiHealthMinimalGet200ResponsePgInfo) SetStatuses(v string)`

SetStatuses sets Statuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


