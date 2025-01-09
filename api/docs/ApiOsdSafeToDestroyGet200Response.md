# ApiOsdSafeToDestroyGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **[]int32** |  | 
**IsSafeToDestroy** | **bool** | Is OSD safe to destroy? | 
**MissingStats** | **[]string** |  | 
**SafeToDestroy** | **[]string** | Is OSD safe to destroy? | 
**StoredPgs** | **[]string** | Stored Pool groups in Osd | 

## Methods

### NewApiOsdSafeToDestroyGet200Response

`func NewApiOsdSafeToDestroyGet200Response(active []int32, isSafeToDestroy bool, missingStats []string, safeToDestroy []string, storedPgs []string, ) *ApiOsdSafeToDestroyGet200Response`

NewApiOsdSafeToDestroyGet200Response instantiates a new ApiOsdSafeToDestroyGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOsdSafeToDestroyGet200ResponseWithDefaults

`func NewApiOsdSafeToDestroyGet200ResponseWithDefaults() *ApiOsdSafeToDestroyGet200Response`

NewApiOsdSafeToDestroyGet200ResponseWithDefaults instantiates a new ApiOsdSafeToDestroyGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApiOsdSafeToDestroyGet200Response) GetActive() []int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiOsdSafeToDestroyGet200Response) GetActiveOk() (*[]int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiOsdSafeToDestroyGet200Response) SetActive(v []int32)`

SetActive sets Active field to given value.


### GetIsSafeToDestroy

`func (o *ApiOsdSafeToDestroyGet200Response) GetIsSafeToDestroy() bool`

GetIsSafeToDestroy returns the IsSafeToDestroy field if non-nil, zero value otherwise.

### GetIsSafeToDestroyOk

`func (o *ApiOsdSafeToDestroyGet200Response) GetIsSafeToDestroyOk() (*bool, bool)`

GetIsSafeToDestroyOk returns a tuple with the IsSafeToDestroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafeToDestroy

`func (o *ApiOsdSafeToDestroyGet200Response) SetIsSafeToDestroy(v bool)`

SetIsSafeToDestroy sets IsSafeToDestroy field to given value.


### GetMissingStats

`func (o *ApiOsdSafeToDestroyGet200Response) GetMissingStats() []string`

GetMissingStats returns the MissingStats field if non-nil, zero value otherwise.

### GetMissingStatsOk

`func (o *ApiOsdSafeToDestroyGet200Response) GetMissingStatsOk() (*[]string, bool)`

GetMissingStatsOk returns a tuple with the MissingStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingStats

`func (o *ApiOsdSafeToDestroyGet200Response) SetMissingStats(v []string)`

SetMissingStats sets MissingStats field to given value.


### GetSafeToDestroy

`func (o *ApiOsdSafeToDestroyGet200Response) GetSafeToDestroy() []string`

GetSafeToDestroy returns the SafeToDestroy field if non-nil, zero value otherwise.

### GetSafeToDestroyOk

`func (o *ApiOsdSafeToDestroyGet200Response) GetSafeToDestroyOk() (*[]string, bool)`

GetSafeToDestroyOk returns a tuple with the SafeToDestroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeToDestroy

`func (o *ApiOsdSafeToDestroyGet200Response) SetSafeToDestroy(v []string)`

SetSafeToDestroy sets SafeToDestroy field to given value.


### GetStoredPgs

`func (o *ApiOsdSafeToDestroyGet200Response) GetStoredPgs() []string`

GetStoredPgs returns the StoredPgs field if non-nil, zero value otherwise.

### GetStoredPgsOk

`func (o *ApiOsdSafeToDestroyGet200Response) GetStoredPgsOk() (*[]string, bool)`

GetStoredPgsOk returns a tuple with the StoredPgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredPgs

`func (o *ApiOsdSafeToDestroyGet200Response) SetStoredPgs(v []string)`

SetStoredPgs sets StoredPgs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


