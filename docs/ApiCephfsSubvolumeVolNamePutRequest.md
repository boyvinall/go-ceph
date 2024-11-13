# ApiCephfsSubvolumeVolNamePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** |  | [optional] [default to ""]
**Size** | **int32** |  | 
**SubvolName** | **string** |  | 

## Methods

### NewApiCephfsSubvolumeVolNamePutRequest

`func NewApiCephfsSubvolumeVolNamePutRequest(size int32, subvolName string, ) *ApiCephfsSubvolumeVolNamePutRequest`

NewApiCephfsSubvolumeVolNamePutRequest instantiates a new ApiCephfsSubvolumeVolNamePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCephfsSubvolumeVolNamePutRequestWithDefaults

`func NewApiCephfsSubvolumeVolNamePutRequestWithDefaults() *ApiCephfsSubvolumeVolNamePutRequest`

NewApiCephfsSubvolumeVolNamePutRequestWithDefaults instantiates a new ApiCephfsSubvolumeVolNamePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *ApiCephfsSubvolumeVolNamePutRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ApiCephfsSubvolumeVolNamePutRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ApiCephfsSubvolumeVolNamePutRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ApiCephfsSubvolumeVolNamePutRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSize

`func (o *ApiCephfsSubvolumeVolNamePutRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApiCephfsSubvolumeVolNamePutRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApiCephfsSubvolumeVolNamePutRequest) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSubvolName

`func (o *ApiCephfsSubvolumeVolNamePutRequest) GetSubvolName() string`

GetSubvolName returns the SubvolName field if non-nil, zero value otherwise.

### GetSubvolNameOk

`func (o *ApiCephfsSubvolumeVolNamePutRequest) GetSubvolNameOk() (*string, bool)`

GetSubvolNameOk returns a tuple with the SubvolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubvolName

`func (o *ApiCephfsSubvolumeVolNamePutRequest) SetSubvolName(v string)`

SetSubvolName sets SubvolName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


