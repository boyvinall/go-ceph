# ApiCephfsSnapshotSchedulePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fs** | **string** |  | 
**Group** | Pointer to **string** |  | [optional] 
**Path** | **string** |  | 
**RetentionPolicy** | Pointer to **string** |  | [optional] 
**SnapSchedule** | **string** |  | 
**Start** | **string** |  | 
**Subvol** | Pointer to **string** |  | [optional] 

## Methods

### NewApiCephfsSnapshotSchedulePostRequest

`func NewApiCephfsSnapshotSchedulePostRequest(fs string, path string, snapSchedule string, start string, ) *ApiCephfsSnapshotSchedulePostRequest`

NewApiCephfsSnapshotSchedulePostRequest instantiates a new ApiCephfsSnapshotSchedulePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCephfsSnapshotSchedulePostRequestWithDefaults

`func NewApiCephfsSnapshotSchedulePostRequestWithDefaults() *ApiCephfsSnapshotSchedulePostRequest`

NewApiCephfsSnapshotSchedulePostRequestWithDefaults instantiates a new ApiCephfsSnapshotSchedulePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFs

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetFs() string`

GetFs returns the Fs field if non-nil, zero value otherwise.

### GetFsOk

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetFsOk() (*string, bool)`

GetFsOk returns a tuple with the Fs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFs

`func (o *ApiCephfsSnapshotSchedulePostRequest) SetFs(v string)`

SetFs sets Fs field to given value.


### GetGroup

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ApiCephfsSnapshotSchedulePostRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ApiCephfsSnapshotSchedulePostRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPath

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiCephfsSnapshotSchedulePostRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetRetentionPolicy

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetRetentionPolicy() string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetRetentionPolicyOk() (*string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *ApiCephfsSnapshotSchedulePostRequest) SetRetentionPolicy(v string)`

SetRetentionPolicy sets RetentionPolicy field to given value.

### HasRetentionPolicy

`func (o *ApiCephfsSnapshotSchedulePostRequest) HasRetentionPolicy() bool`

HasRetentionPolicy returns a boolean if a field has been set.

### GetSnapSchedule

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetSnapSchedule() string`

GetSnapSchedule returns the SnapSchedule field if non-nil, zero value otherwise.

### GetSnapScheduleOk

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetSnapScheduleOk() (*string, bool)`

GetSnapScheduleOk returns a tuple with the SnapSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapSchedule

`func (o *ApiCephfsSnapshotSchedulePostRequest) SetSnapSchedule(v string)`

SetSnapSchedule sets SnapSchedule field to given value.


### GetStart

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ApiCephfsSnapshotSchedulePostRequest) SetStart(v string)`

SetStart sets Start field to given value.


### GetSubvol

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetSubvol() string`

GetSubvol returns the Subvol field if non-nil, zero value otherwise.

### GetSubvolOk

`func (o *ApiCephfsSnapshotSchedulePostRequest) GetSubvolOk() (*string, bool)`

GetSubvolOk returns a tuple with the Subvol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubvol

`func (o *ApiCephfsSnapshotSchedulePostRequest) SetSubvol(v string)`

SetSubvol sets Subvol field to given value.

### HasSubvol

`func (o *ApiCephfsSnapshotSchedulePostRequest) HasSubvol() bool`

HasSubvol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


