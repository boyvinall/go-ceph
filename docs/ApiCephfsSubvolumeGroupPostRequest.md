# ApiCephfsSubvolumeGroupPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** |  | 
**VolName** | **string** |  | 

## Methods

### NewApiCephfsSubvolumeGroupPostRequest

`func NewApiCephfsSubvolumeGroupPostRequest(groupName string, volName string, ) *ApiCephfsSubvolumeGroupPostRequest`

NewApiCephfsSubvolumeGroupPostRequest instantiates a new ApiCephfsSubvolumeGroupPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCephfsSubvolumeGroupPostRequestWithDefaults

`func NewApiCephfsSubvolumeGroupPostRequestWithDefaults() *ApiCephfsSubvolumeGroupPostRequest`

NewApiCephfsSubvolumeGroupPostRequestWithDefaults instantiates a new ApiCephfsSubvolumeGroupPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *ApiCephfsSubvolumeGroupPostRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ApiCephfsSubvolumeGroupPostRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ApiCephfsSubvolumeGroupPostRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetVolName

`func (o *ApiCephfsSubvolumeGroupPostRequest) GetVolName() string`

GetVolName returns the VolName field if non-nil, zero value otherwise.

### GetVolNameOk

`func (o *ApiCephfsSubvolumeGroupPostRequest) GetVolNameOk() (*string, bool)`

GetVolNameOk returns a tuple with the VolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolName

`func (o *ApiCephfsSubvolumeGroupPostRequest) SetVolName(v string)`

SetVolName sets VolName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


