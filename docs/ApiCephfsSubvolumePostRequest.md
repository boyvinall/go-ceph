# ApiCephfsSubvolumePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubvolName** | **string** |  | 
**VolName** | **string** |  | 

## Methods

### NewApiCephfsSubvolumePostRequest

`func NewApiCephfsSubvolumePostRequest(subvolName string, volName string, ) *ApiCephfsSubvolumePostRequest`

NewApiCephfsSubvolumePostRequest instantiates a new ApiCephfsSubvolumePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCephfsSubvolumePostRequestWithDefaults

`func NewApiCephfsSubvolumePostRequestWithDefaults() *ApiCephfsSubvolumePostRequest`

NewApiCephfsSubvolumePostRequestWithDefaults instantiates a new ApiCephfsSubvolumePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubvolName

`func (o *ApiCephfsSubvolumePostRequest) GetSubvolName() string`

GetSubvolName returns the SubvolName field if non-nil, zero value otherwise.

### GetSubvolNameOk

`func (o *ApiCephfsSubvolumePostRequest) GetSubvolNameOk() (*string, bool)`

GetSubvolNameOk returns a tuple with the SubvolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubvolName

`func (o *ApiCephfsSubvolumePostRequest) SetSubvolName(v string)`

SetSubvolName sets SubvolName field to given value.


### GetVolName

`func (o *ApiCephfsSubvolumePostRequest) GetVolName() string`

GetVolName returns the VolName field if non-nil, zero value otherwise.

### GetVolNameOk

`func (o *ApiCephfsSubvolumePostRequest) GetVolNameOk() (*string, bool)`

GetVolNameOk returns a tuple with the VolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolName

`func (o *ApiCephfsSubvolumePostRequest) SetVolName(v string)`

SetVolName sets VolName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


