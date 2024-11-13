# ApiOsdFlagsIndividualPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | [**ApiOsdFlagsIndividualPutRequestFlags**](ApiOsdFlagsIndividualPutRequestFlags.md) |  | 
**Ids** | **[]int32** | List of OSD ids the flags should be applied to. | 

## Methods

### NewApiOsdFlagsIndividualPutRequest

`func NewApiOsdFlagsIndividualPutRequest(flags ApiOsdFlagsIndividualPutRequestFlags, ids []int32, ) *ApiOsdFlagsIndividualPutRequest`

NewApiOsdFlagsIndividualPutRequest instantiates a new ApiOsdFlagsIndividualPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOsdFlagsIndividualPutRequestWithDefaults

`func NewApiOsdFlagsIndividualPutRequestWithDefaults() *ApiOsdFlagsIndividualPutRequest`

NewApiOsdFlagsIndividualPutRequestWithDefaults instantiates a new ApiOsdFlagsIndividualPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *ApiOsdFlagsIndividualPutRequest) GetFlags() ApiOsdFlagsIndividualPutRequestFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ApiOsdFlagsIndividualPutRequest) GetFlagsOk() (*ApiOsdFlagsIndividualPutRequestFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ApiOsdFlagsIndividualPutRequest) SetFlags(v ApiOsdFlagsIndividualPutRequestFlags)`

SetFlags sets Flags field to given value.


### GetIds

`func (o *ApiOsdFlagsIndividualPutRequest) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ApiOsdFlagsIndividualPutRequest) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ApiOsdFlagsIndividualPutRequest) SetIds(v []int32)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


