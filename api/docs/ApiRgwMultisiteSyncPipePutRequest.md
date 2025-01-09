# ApiRgwMultisiteSyncPipePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** |  | [optional] [default to ""]
**DestinationBucket** | Pointer to **string** |  | [optional] [default to ""]
**DestinationZones** | **string** |  | 
**GroupId** | **string** |  | 
**Mode** | Pointer to **string** |  | [optional] [default to ""]
**PipeId** | **string** |  | 
**SourceBucket** | Pointer to **string** |  | [optional] [default to ""]
**SourceZones** | **string** |  | 
**User** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewApiRgwMultisiteSyncPipePutRequest

`func NewApiRgwMultisiteSyncPipePutRequest(destinationZones string, groupId string, pipeId string, sourceZones string, ) *ApiRgwMultisiteSyncPipePutRequest`

NewApiRgwMultisiteSyncPipePutRequest instantiates a new ApiRgwMultisiteSyncPipePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwMultisiteSyncPipePutRequestWithDefaults

`func NewApiRgwMultisiteSyncPipePutRequestWithDefaults() *ApiRgwMultisiteSyncPipePutRequest`

NewApiRgwMultisiteSyncPipePutRequestWithDefaults instantiates a new ApiRgwMultisiteSyncPipePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ApiRgwMultisiteSyncPipePutRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *ApiRgwMultisiteSyncPipePutRequest) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetDestinationBucket

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetDestinationBucket() string`

GetDestinationBucket returns the DestinationBucket field if non-nil, zero value otherwise.

### GetDestinationBucketOk

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetDestinationBucketOk() (*string, bool)`

GetDestinationBucketOk returns a tuple with the DestinationBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationBucket

`func (o *ApiRgwMultisiteSyncPipePutRequest) SetDestinationBucket(v string)`

SetDestinationBucket sets DestinationBucket field to given value.

### HasDestinationBucket

`func (o *ApiRgwMultisiteSyncPipePutRequest) HasDestinationBucket() bool`

HasDestinationBucket returns a boolean if a field has been set.

### GetDestinationZones

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetDestinationZones() string`

GetDestinationZones returns the DestinationZones field if non-nil, zero value otherwise.

### GetDestinationZonesOk

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetDestinationZonesOk() (*string, bool)`

GetDestinationZonesOk returns a tuple with the DestinationZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationZones

`func (o *ApiRgwMultisiteSyncPipePutRequest) SetDestinationZones(v string)`

SetDestinationZones sets DestinationZones field to given value.


### GetGroupId

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiRgwMultisiteSyncPipePutRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetMode

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApiRgwMultisiteSyncPipePutRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ApiRgwMultisiteSyncPipePutRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPipeId

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetPipeId() string`

GetPipeId returns the PipeId field if non-nil, zero value otherwise.

### GetPipeIdOk

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetPipeIdOk() (*string, bool)`

GetPipeIdOk returns a tuple with the PipeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeId

`func (o *ApiRgwMultisiteSyncPipePutRequest) SetPipeId(v string)`

SetPipeId sets PipeId field to given value.


### GetSourceBucket

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetSourceBucket() string`

GetSourceBucket returns the SourceBucket field if non-nil, zero value otherwise.

### GetSourceBucketOk

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetSourceBucketOk() (*string, bool)`

GetSourceBucketOk returns a tuple with the SourceBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBucket

`func (o *ApiRgwMultisiteSyncPipePutRequest) SetSourceBucket(v string)`

SetSourceBucket sets SourceBucket field to given value.

### HasSourceBucket

`func (o *ApiRgwMultisiteSyncPipePutRequest) HasSourceBucket() bool`

HasSourceBucket returns a boolean if a field has been set.

### GetSourceZones

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetSourceZones() string`

GetSourceZones returns the SourceZones field if non-nil, zero value otherwise.

### GetSourceZonesOk

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetSourceZonesOk() (*string, bool)`

GetSourceZonesOk returns a tuple with the SourceZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceZones

`func (o *ApiRgwMultisiteSyncPipePutRequest) SetSourceZones(v string)`

SetSourceZones sets SourceZones field to given value.


### GetUser

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApiRgwMultisiteSyncPipePutRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApiRgwMultisiteSyncPipePutRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ApiRgwMultisiteSyncPipePutRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


