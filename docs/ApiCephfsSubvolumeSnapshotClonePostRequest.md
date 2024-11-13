# ApiCephfsSubvolumeSnapshotClonePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneName** | **string** |  | 
**GroupName** | Pointer to **string** |  | [optional] [default to ""]
**SnapName** | **string** |  | 
**SubvolName** | **string** |  | 
**TargetGroupName** | Pointer to **string** |  | [optional] [default to ""]
**VolName** | **string** |  | 

## Methods

### NewApiCephfsSubvolumeSnapshotClonePostRequest

`func NewApiCephfsSubvolumeSnapshotClonePostRequest(cloneName string, snapName string, subvolName string, volName string, ) *ApiCephfsSubvolumeSnapshotClonePostRequest`

NewApiCephfsSubvolumeSnapshotClonePostRequest instantiates a new ApiCephfsSubvolumeSnapshotClonePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCephfsSubvolumeSnapshotClonePostRequestWithDefaults

`func NewApiCephfsSubvolumeSnapshotClonePostRequestWithDefaults() *ApiCephfsSubvolumeSnapshotClonePostRequest`

NewApiCephfsSubvolumeSnapshotClonePostRequestWithDefaults instantiates a new ApiCephfsSubvolumeSnapshotClonePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetCloneName() string`

GetCloneName returns the CloneName field if non-nil, zero value otherwise.

### GetCloneNameOk

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetCloneNameOk() (*string, bool)`

GetCloneNameOk returns a tuple with the CloneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetCloneName(v string)`

SetCloneName sets CloneName field to given value.


### GetGroupName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSnapName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetSnapName() string`

GetSnapName returns the SnapName field if non-nil, zero value otherwise.

### GetSnapNameOk

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetSnapNameOk() (*string, bool)`

GetSnapNameOk returns a tuple with the SnapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetSnapName(v string)`

SetSnapName sets SnapName field to given value.


### GetSubvolName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetSubvolName() string`

GetSubvolName returns the SubvolName field if non-nil, zero value otherwise.

### GetSubvolNameOk

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetSubvolNameOk() (*string, bool)`

GetSubvolNameOk returns a tuple with the SubvolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubvolName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetSubvolName(v string)`

SetSubvolName sets SubvolName field to given value.


### GetTargetGroupName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetTargetGroupName() string`

GetTargetGroupName returns the TargetGroupName field if non-nil, zero value otherwise.

### GetTargetGroupNameOk

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetTargetGroupNameOk() (*string, bool)`

GetTargetGroupNameOk returns a tuple with the TargetGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetTargetGroupName(v string)`

SetTargetGroupName sets TargetGroupName field to given value.

### HasTargetGroupName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) HasTargetGroupName() bool`

HasTargetGroupName returns a boolean if a field has been set.

### GetVolName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetVolName() string`

GetVolName returns the VolName field if non-nil, zero value otherwise.

### GetVolNameOk

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) GetVolNameOk() (*string, bool)`

GetVolNameOk returns a tuple with the VolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolName

`func (o *ApiCephfsSubvolumeSnapshotClonePostRequest) SetVolName(v string)`

SetVolName sets VolName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


