# ApiMonitorGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InQuorum** | [**[]ApiMonitorGet200ResponseInQuorumInner**](ApiMonitorGet200ResponseInQuorumInner.md) |  | 
**MonStatus** | [**ApiMonitorGet200ResponseMonStatus**](ApiMonitorGet200ResponseMonStatus.md) |  | 
**OutQuorum** | **[]int32** |  | 

## Methods

### NewApiMonitorGet200Response

`func NewApiMonitorGet200Response(inQuorum []ApiMonitorGet200ResponseInQuorumInner, monStatus ApiMonitorGet200ResponseMonStatus, outQuorum []int32, ) *ApiMonitorGet200Response`

NewApiMonitorGet200Response instantiates a new ApiMonitorGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMonitorGet200ResponseWithDefaults

`func NewApiMonitorGet200ResponseWithDefaults() *ApiMonitorGet200Response`

NewApiMonitorGet200ResponseWithDefaults instantiates a new ApiMonitorGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInQuorum

`func (o *ApiMonitorGet200Response) GetInQuorum() []ApiMonitorGet200ResponseInQuorumInner`

GetInQuorum returns the InQuorum field if non-nil, zero value otherwise.

### GetInQuorumOk

`func (o *ApiMonitorGet200Response) GetInQuorumOk() (*[]ApiMonitorGet200ResponseInQuorumInner, bool)`

GetInQuorumOk returns a tuple with the InQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInQuorum

`func (o *ApiMonitorGet200Response) SetInQuorum(v []ApiMonitorGet200ResponseInQuorumInner)`

SetInQuorum sets InQuorum field to given value.


### GetMonStatus

`func (o *ApiMonitorGet200Response) GetMonStatus() ApiMonitorGet200ResponseMonStatus`

GetMonStatus returns the MonStatus field if non-nil, zero value otherwise.

### GetMonStatusOk

`func (o *ApiMonitorGet200Response) GetMonStatusOk() (*ApiMonitorGet200ResponseMonStatus, bool)`

GetMonStatusOk returns a tuple with the MonStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonStatus

`func (o *ApiMonitorGet200Response) SetMonStatus(v ApiMonitorGet200ResponseMonStatus)`

SetMonStatus sets MonStatus field to given value.


### GetOutQuorum

`func (o *ApiMonitorGet200Response) GetOutQuorum() []int32`

GetOutQuorum returns the OutQuorum field if non-nil, zero value otherwise.

### GetOutQuorumOk

`func (o *ApiMonitorGet200Response) GetOutQuorumOk() (*[]int32, bool)`

GetOutQuorumOk returns a tuple with the OutQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutQuorum

`func (o *ApiMonitorGet200Response) SetOutQuorum(v []int32)`

SetOutQuorum sets OutQuorum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


