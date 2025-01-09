# ApiOsdFlagsIndividualPut200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | **[]string** | List of added flags | 
**Ids** | **[]int32** | List of updated OSDs | 
**Removed** | **[]string** | List of removed flags | 

## Methods

### NewApiOsdFlagsIndividualPut200Response

`func NewApiOsdFlagsIndividualPut200Response(added []string, ids []int32, removed []string, ) *ApiOsdFlagsIndividualPut200Response`

NewApiOsdFlagsIndividualPut200Response instantiates a new ApiOsdFlagsIndividualPut200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOsdFlagsIndividualPut200ResponseWithDefaults

`func NewApiOsdFlagsIndividualPut200ResponseWithDefaults() *ApiOsdFlagsIndividualPut200Response`

NewApiOsdFlagsIndividualPut200ResponseWithDefaults instantiates a new ApiOsdFlagsIndividualPut200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *ApiOsdFlagsIndividualPut200Response) GetAdded() []string`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *ApiOsdFlagsIndividualPut200Response) GetAddedOk() (*[]string, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *ApiOsdFlagsIndividualPut200Response) SetAdded(v []string)`

SetAdded sets Added field to given value.


### GetIds

`func (o *ApiOsdFlagsIndividualPut200Response) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ApiOsdFlagsIndividualPut200Response) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ApiOsdFlagsIndividualPut200Response) SetIds(v []int32)`

SetIds sets Ids field to given value.


### GetRemoved

`func (o *ApiOsdFlagsIndividualPut200Response) GetRemoved() []string`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *ApiOsdFlagsIndividualPut200Response) GetRemovedOk() (*[]string, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *ApiOsdFlagsIndividualPut200Response) SetRemoved(v []string)`

SetRemoved sets Removed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


