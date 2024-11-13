# ApiHealthMinimalGet200ResponseMonStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monmap** | [**ApiHealthMinimalGet200ResponseMonStatusMonmap**](ApiHealthMinimalGet200ResponseMonStatusMonmap.md) |  | 
**Quorum** | **[]int32** |  | 

## Methods

### NewApiHealthMinimalGet200ResponseMonStatus

`func NewApiHealthMinimalGet200ResponseMonStatus(monmap ApiHealthMinimalGet200ResponseMonStatusMonmap, quorum []int32, ) *ApiHealthMinimalGet200ResponseMonStatus`

NewApiHealthMinimalGet200ResponseMonStatus instantiates a new ApiHealthMinimalGet200ResponseMonStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHealthMinimalGet200ResponseMonStatusWithDefaults

`func NewApiHealthMinimalGet200ResponseMonStatusWithDefaults() *ApiHealthMinimalGet200ResponseMonStatus`

NewApiHealthMinimalGet200ResponseMonStatusWithDefaults instantiates a new ApiHealthMinimalGet200ResponseMonStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonmap

`func (o *ApiHealthMinimalGet200ResponseMonStatus) GetMonmap() ApiHealthMinimalGet200ResponseMonStatusMonmap`

GetMonmap returns the Monmap field if non-nil, zero value otherwise.

### GetMonmapOk

`func (o *ApiHealthMinimalGet200ResponseMonStatus) GetMonmapOk() (*ApiHealthMinimalGet200ResponseMonStatusMonmap, bool)`

GetMonmapOk returns a tuple with the Monmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonmap

`func (o *ApiHealthMinimalGet200ResponseMonStatus) SetMonmap(v ApiHealthMinimalGet200ResponseMonStatusMonmap)`

SetMonmap sets Monmap field to given value.


### GetQuorum

`func (o *ApiHealthMinimalGet200ResponseMonStatus) GetQuorum() []int32`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *ApiHealthMinimalGet200ResponseMonStatus) GetQuorumOk() (*[]int32, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *ApiHealthMinimalGet200ResponseMonStatus) SetQuorum(v []int32)`

SetQuorum sets Quorum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


