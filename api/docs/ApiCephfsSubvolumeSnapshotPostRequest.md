# ApiCephfsSubvolumeSnapshotPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** |  | [optional] [default to ""]
**SnapName** | **string** |  | 
**SubvolName** | **string** |  | 
**VolName** | **string** |  | 

## Methods

### NewApiCephfsSubvolumeSnapshotPostRequest

`func NewApiCephfsSubvolumeSnapshotPostRequest(snapName string, subvolName string, volName string, ) *ApiCephfsSubvolumeSnapshotPostRequest`

NewApiCephfsSubvolumeSnapshotPostRequest instantiates a new ApiCephfsSubvolumeSnapshotPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCephfsSubvolumeSnapshotPostRequestWithDefaults

`func NewApiCephfsSubvolumeSnapshotPostRequestWithDefaults() *ApiCephfsSubvolumeSnapshotPostRequest`

NewApiCephfsSubvolumeSnapshotPostRequestWithDefaults instantiates a new ApiCephfsSubvolumeSnapshotPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSnapName

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) GetSnapName() string`

GetSnapName returns the SnapName field if non-nil, zero value otherwise.

### GetSnapNameOk

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) GetSnapNameOk() (*string, bool)`

GetSnapNameOk returns a tuple with the SnapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapName

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) SetSnapName(v string)`

SetSnapName sets SnapName field to given value.


### GetSubvolName

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) GetSubvolName() string`

GetSubvolName returns the SubvolName field if non-nil, zero value otherwise.

### GetSubvolNameOk

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) GetSubvolNameOk() (*string, bool)`

GetSubvolNameOk returns a tuple with the SubvolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubvolName

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) SetSubvolName(v string)`

SetSubvolName sets SubvolName field to given value.


### GetVolName

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) GetVolName() string`

GetVolName returns the VolName field if non-nil, zero value otherwise.

### GetVolNameOk

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) GetVolNameOk() (*string, bool)`

GetVolNameOk returns a tuple with the VolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolName

`func (o *ApiCephfsSubvolumeSnapshotPostRequest) SetVolName(v string)`

SetVolName sets VolName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


