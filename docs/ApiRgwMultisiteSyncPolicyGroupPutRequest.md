# ApiRgwMultisiteSyncPolicyGroupPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** |  | [optional] [default to ""]
**GroupId** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewApiRgwMultisiteSyncPolicyGroupPutRequest

`func NewApiRgwMultisiteSyncPolicyGroupPutRequest(groupId string, status string, ) *ApiRgwMultisiteSyncPolicyGroupPutRequest`

NewApiRgwMultisiteSyncPolicyGroupPutRequest instantiates a new ApiRgwMultisiteSyncPolicyGroupPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwMultisiteSyncPolicyGroupPutRequestWithDefaults

`func NewApiRgwMultisiteSyncPolicyGroupPutRequestWithDefaults() *ApiRgwMultisiteSyncPolicyGroupPutRequest`

NewApiRgwMultisiteSyncPolicyGroupPutRequestWithDefaults instantiates a new ApiRgwMultisiteSyncPolicyGroupPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetGroupId

`func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetStatus

`func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiRgwMultisiteSyncPolicyGroupPutRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


