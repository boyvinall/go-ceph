# ApiOsdFlagsPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | **[]string** | List of flags to set. The flags &#x60;recovery_deletes&#x60;, &#x60;sortbitwise&#x60; and &#x60;pglog_hardlimit&#x60; cannot be unset. Additionally &#x60;purged_snapshots&#x60; cannot even be set. | 

## Methods

### NewApiOsdFlagsPutRequest

`func NewApiOsdFlagsPutRequest(flags []string, ) *ApiOsdFlagsPutRequest`

NewApiOsdFlagsPutRequest instantiates a new ApiOsdFlagsPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOsdFlagsPutRequestWithDefaults

`func NewApiOsdFlagsPutRequestWithDefaults() *ApiOsdFlagsPutRequest`

NewApiOsdFlagsPutRequestWithDefaults instantiates a new ApiOsdFlagsPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *ApiOsdFlagsPutRequest) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ApiOsdFlagsPutRequest) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ApiOsdFlagsPutRequest) SetFlags(v []string)`

SetFlags sets Flags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


