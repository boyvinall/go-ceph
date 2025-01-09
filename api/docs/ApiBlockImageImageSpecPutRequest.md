# ApiBlockImageImageSpecPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **string** |  | [optional] 
**EnableMirror** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **string** |  | [optional] 
**Force** | Pointer to **bool** |  | [optional] [default to false]
**ImageMirrorMode** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**MirrorMode** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Primary** | Pointer to **string** |  | [optional] 
**RemoveScheduling** | Pointer to **bool** |  | [optional] [default to false]
**Resync** | Pointer to **bool** |  | [optional] [default to false]
**ScheduleInterval** | Pointer to **string** |  | [optional] [default to ""]
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiBlockImageImageSpecPutRequest

`func NewApiBlockImageImageSpecPutRequest() *ApiBlockImageImageSpecPutRequest`

NewApiBlockImageImageSpecPutRequest instantiates a new ApiBlockImageImageSpecPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBlockImageImageSpecPutRequestWithDefaults

`func NewApiBlockImageImageSpecPutRequestWithDefaults() *ApiBlockImageImageSpecPutRequest`

NewApiBlockImageImageSpecPutRequestWithDefaults instantiates a new ApiBlockImageImageSpecPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *ApiBlockImageImageSpecPutRequest) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ApiBlockImageImageSpecPutRequest) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ApiBlockImageImageSpecPutRequest) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ApiBlockImageImageSpecPutRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetEnableMirror

`func (o *ApiBlockImageImageSpecPutRequest) GetEnableMirror() string`

GetEnableMirror returns the EnableMirror field if non-nil, zero value otherwise.

### GetEnableMirrorOk

`func (o *ApiBlockImageImageSpecPutRequest) GetEnableMirrorOk() (*string, bool)`

GetEnableMirrorOk returns a tuple with the EnableMirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMirror

`func (o *ApiBlockImageImageSpecPutRequest) SetEnableMirror(v string)`

SetEnableMirror sets EnableMirror field to given value.

### HasEnableMirror

`func (o *ApiBlockImageImageSpecPutRequest) HasEnableMirror() bool`

HasEnableMirror returns a boolean if a field has been set.

### GetFeatures

`func (o *ApiBlockImageImageSpecPutRequest) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ApiBlockImageImageSpecPutRequest) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ApiBlockImageImageSpecPutRequest) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ApiBlockImageImageSpecPutRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetForce

`func (o *ApiBlockImageImageSpecPutRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *ApiBlockImageImageSpecPutRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *ApiBlockImageImageSpecPutRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *ApiBlockImageImageSpecPutRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetImageMirrorMode

`func (o *ApiBlockImageImageSpecPutRequest) GetImageMirrorMode() string`

GetImageMirrorMode returns the ImageMirrorMode field if non-nil, zero value otherwise.

### GetImageMirrorModeOk

`func (o *ApiBlockImageImageSpecPutRequest) GetImageMirrorModeOk() (*string, bool)`

GetImageMirrorModeOk returns a tuple with the ImageMirrorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMirrorMode

`func (o *ApiBlockImageImageSpecPutRequest) SetImageMirrorMode(v string)`

SetImageMirrorMode sets ImageMirrorMode field to given value.

### HasImageMirrorMode

`func (o *ApiBlockImageImageSpecPutRequest) HasImageMirrorMode() bool`

HasImageMirrorMode returns a boolean if a field has been set.

### GetMetadata

`func (o *ApiBlockImageImageSpecPutRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApiBlockImageImageSpecPutRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApiBlockImageImageSpecPutRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApiBlockImageImageSpecPutRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMirrorMode

`func (o *ApiBlockImageImageSpecPutRequest) GetMirrorMode() string`

GetMirrorMode returns the MirrorMode field if non-nil, zero value otherwise.

### GetMirrorModeOk

`func (o *ApiBlockImageImageSpecPutRequest) GetMirrorModeOk() (*string, bool)`

GetMirrorModeOk returns a tuple with the MirrorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorMode

`func (o *ApiBlockImageImageSpecPutRequest) SetMirrorMode(v string)`

SetMirrorMode sets MirrorMode field to given value.

### HasMirrorMode

`func (o *ApiBlockImageImageSpecPutRequest) HasMirrorMode() bool`

HasMirrorMode returns a boolean if a field has been set.

### GetName

`func (o *ApiBlockImageImageSpecPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiBlockImageImageSpecPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiBlockImageImageSpecPutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiBlockImageImageSpecPutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimary

`func (o *ApiBlockImageImageSpecPutRequest) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *ApiBlockImageImageSpecPutRequest) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *ApiBlockImageImageSpecPutRequest) SetPrimary(v string)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *ApiBlockImageImageSpecPutRequest) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetRemoveScheduling

`func (o *ApiBlockImageImageSpecPutRequest) GetRemoveScheduling() bool`

GetRemoveScheduling returns the RemoveScheduling field if non-nil, zero value otherwise.

### GetRemoveSchedulingOk

`func (o *ApiBlockImageImageSpecPutRequest) GetRemoveSchedulingOk() (*bool, bool)`

GetRemoveSchedulingOk returns a tuple with the RemoveScheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveScheduling

`func (o *ApiBlockImageImageSpecPutRequest) SetRemoveScheduling(v bool)`

SetRemoveScheduling sets RemoveScheduling field to given value.

### HasRemoveScheduling

`func (o *ApiBlockImageImageSpecPutRequest) HasRemoveScheduling() bool`

HasRemoveScheduling returns a boolean if a field has been set.

### GetResync

`func (o *ApiBlockImageImageSpecPutRequest) GetResync() bool`

GetResync returns the Resync field if non-nil, zero value otherwise.

### GetResyncOk

`func (o *ApiBlockImageImageSpecPutRequest) GetResyncOk() (*bool, bool)`

GetResyncOk returns a tuple with the Resync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResync

`func (o *ApiBlockImageImageSpecPutRequest) SetResync(v bool)`

SetResync sets Resync field to given value.

### HasResync

`func (o *ApiBlockImageImageSpecPutRequest) HasResync() bool`

HasResync returns a boolean if a field has been set.

### GetScheduleInterval

`func (o *ApiBlockImageImageSpecPutRequest) GetScheduleInterval() string`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *ApiBlockImageImageSpecPutRequest) GetScheduleIntervalOk() (*string, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *ApiBlockImageImageSpecPutRequest) SetScheduleInterval(v string)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *ApiBlockImageImageSpecPutRequest) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetSize

`func (o *ApiBlockImageImageSpecPutRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApiBlockImageImageSpecPutRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApiBlockImageImageSpecPutRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ApiBlockImageImageSpecPutRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


