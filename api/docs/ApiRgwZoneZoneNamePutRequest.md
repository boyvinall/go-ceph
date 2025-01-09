# ApiRgwZoneZoneNamePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] [default to ""]
**Compression** | Pointer to **string** |  | [optional] [default to ""]
**DataExtraPool** | Pointer to **string** |  | [optional] [default to ""]
**DataPool** | Pointer to **string** |  | [optional] [default to ""]
**DataPoolClass** | Pointer to **string** |  | [optional] [default to ""]
**Default** | Pointer to **string** |  | [optional] [default to ""]
**IndexPool** | Pointer to **string** |  | [optional] [default to ""]
**Master** | Pointer to **string** |  | [optional] [default to ""]
**NewZoneName** | **string** |  | 
**PlacementTarget** | Pointer to **string** |  | [optional] [default to ""]
**SecretKey** | Pointer to **string** |  | [optional] [default to ""]
**StorageClass** | Pointer to **string** |  | [optional] [default to ""]
**ZoneEndpoints** | Pointer to **string** |  | [optional] [default to ""]
**ZonegroupName** | **string** |  | 

## Methods

### NewApiRgwZoneZoneNamePutRequest

`func NewApiRgwZoneZoneNamePutRequest(newZoneName string, zonegroupName string, ) *ApiRgwZoneZoneNamePutRequest`

NewApiRgwZoneZoneNamePutRequest instantiates a new ApiRgwZoneZoneNamePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwZoneZoneNamePutRequestWithDefaults

`func NewApiRgwZoneZoneNamePutRequestWithDefaults() *ApiRgwZoneZoneNamePutRequest`

NewApiRgwZoneZoneNamePutRequestWithDefaults instantiates a new ApiRgwZoneZoneNamePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *ApiRgwZoneZoneNamePutRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ApiRgwZoneZoneNamePutRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ApiRgwZoneZoneNamePutRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetCompression

`func (o *ApiRgwZoneZoneNamePutRequest) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *ApiRgwZoneZoneNamePutRequest) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *ApiRgwZoneZoneNamePutRequest) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetDataExtraPool

`func (o *ApiRgwZoneZoneNamePutRequest) GetDataExtraPool() string`

GetDataExtraPool returns the DataExtraPool field if non-nil, zero value otherwise.

### GetDataExtraPoolOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetDataExtraPoolOk() (*string, bool)`

GetDataExtraPoolOk returns a tuple with the DataExtraPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataExtraPool

`func (o *ApiRgwZoneZoneNamePutRequest) SetDataExtraPool(v string)`

SetDataExtraPool sets DataExtraPool field to given value.

### HasDataExtraPool

`func (o *ApiRgwZoneZoneNamePutRequest) HasDataExtraPool() bool`

HasDataExtraPool returns a boolean if a field has been set.

### GetDataPool

`func (o *ApiRgwZoneZoneNamePutRequest) GetDataPool() string`

GetDataPool returns the DataPool field if non-nil, zero value otherwise.

### GetDataPoolOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetDataPoolOk() (*string, bool)`

GetDataPoolOk returns a tuple with the DataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPool

`func (o *ApiRgwZoneZoneNamePutRequest) SetDataPool(v string)`

SetDataPool sets DataPool field to given value.

### HasDataPool

`func (o *ApiRgwZoneZoneNamePutRequest) HasDataPool() bool`

HasDataPool returns a boolean if a field has been set.

### GetDataPoolClass

`func (o *ApiRgwZoneZoneNamePutRequest) GetDataPoolClass() string`

GetDataPoolClass returns the DataPoolClass field if non-nil, zero value otherwise.

### GetDataPoolClassOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetDataPoolClassOk() (*string, bool)`

GetDataPoolClassOk returns a tuple with the DataPoolClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoolClass

`func (o *ApiRgwZoneZoneNamePutRequest) SetDataPoolClass(v string)`

SetDataPoolClass sets DataPoolClass field to given value.

### HasDataPoolClass

`func (o *ApiRgwZoneZoneNamePutRequest) HasDataPoolClass() bool`

HasDataPoolClass returns a boolean if a field has been set.

### GetDefault

`func (o *ApiRgwZoneZoneNamePutRequest) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ApiRgwZoneZoneNamePutRequest) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ApiRgwZoneZoneNamePutRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetIndexPool

`func (o *ApiRgwZoneZoneNamePutRequest) GetIndexPool() string`

GetIndexPool returns the IndexPool field if non-nil, zero value otherwise.

### GetIndexPoolOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetIndexPoolOk() (*string, bool)`

GetIndexPoolOk returns a tuple with the IndexPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPool

`func (o *ApiRgwZoneZoneNamePutRequest) SetIndexPool(v string)`

SetIndexPool sets IndexPool field to given value.

### HasIndexPool

`func (o *ApiRgwZoneZoneNamePutRequest) HasIndexPool() bool`

HasIndexPool returns a boolean if a field has been set.

### GetMaster

`func (o *ApiRgwZoneZoneNamePutRequest) GetMaster() string`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetMasterOk() (*string, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ApiRgwZoneZoneNamePutRequest) SetMaster(v string)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ApiRgwZoneZoneNamePutRequest) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetNewZoneName

`func (o *ApiRgwZoneZoneNamePutRequest) GetNewZoneName() string`

GetNewZoneName returns the NewZoneName field if non-nil, zero value otherwise.

### GetNewZoneNameOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetNewZoneNameOk() (*string, bool)`

GetNewZoneNameOk returns a tuple with the NewZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewZoneName

`func (o *ApiRgwZoneZoneNamePutRequest) SetNewZoneName(v string)`

SetNewZoneName sets NewZoneName field to given value.


### GetPlacementTarget

`func (o *ApiRgwZoneZoneNamePutRequest) GetPlacementTarget() string`

GetPlacementTarget returns the PlacementTarget field if non-nil, zero value otherwise.

### GetPlacementTargetOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetPlacementTargetOk() (*string, bool)`

GetPlacementTargetOk returns a tuple with the PlacementTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementTarget

`func (o *ApiRgwZoneZoneNamePutRequest) SetPlacementTarget(v string)`

SetPlacementTarget sets PlacementTarget field to given value.

### HasPlacementTarget

`func (o *ApiRgwZoneZoneNamePutRequest) HasPlacementTarget() bool`

HasPlacementTarget returns a boolean if a field has been set.

### GetSecretKey

`func (o *ApiRgwZoneZoneNamePutRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ApiRgwZoneZoneNamePutRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ApiRgwZoneZoneNamePutRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetStorageClass

`func (o *ApiRgwZoneZoneNamePutRequest) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *ApiRgwZoneZoneNamePutRequest) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *ApiRgwZoneZoneNamePutRequest) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetZoneEndpoints

`func (o *ApiRgwZoneZoneNamePutRequest) GetZoneEndpoints() string`

GetZoneEndpoints returns the ZoneEndpoints field if non-nil, zero value otherwise.

### GetZoneEndpointsOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetZoneEndpointsOk() (*string, bool)`

GetZoneEndpointsOk returns a tuple with the ZoneEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneEndpoints

`func (o *ApiRgwZoneZoneNamePutRequest) SetZoneEndpoints(v string)`

SetZoneEndpoints sets ZoneEndpoints field to given value.

### HasZoneEndpoints

`func (o *ApiRgwZoneZoneNamePutRequest) HasZoneEndpoints() bool`

HasZoneEndpoints returns a boolean if a field has been set.

### GetZonegroupName

`func (o *ApiRgwZoneZoneNamePutRequest) GetZonegroupName() string`

GetZonegroupName returns the ZonegroupName field if non-nil, zero value otherwise.

### GetZonegroupNameOk

`func (o *ApiRgwZoneZoneNamePutRequest) GetZonegroupNameOk() (*string, bool)`

GetZonegroupNameOk returns a tuple with the ZonegroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonegroupName

`func (o *ApiRgwZoneZoneNamePutRequest) SetZonegroupName(v string)`

SetZonegroupName sets ZonegroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


