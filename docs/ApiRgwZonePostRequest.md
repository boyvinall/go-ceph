# ApiRgwZonePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] [default to false]
**Master** | Pointer to **bool** |  | [optional] [default to false]
**SecretKey** | Pointer to **string** |  | [optional] 
**ZoneEndpoints** | Pointer to **string** |  | [optional] 
**ZoneName** | **string** |  | 
**ZonegroupName** | Pointer to **string** |  | [optional] 

## Methods

### NewApiRgwZonePostRequest

`func NewApiRgwZonePostRequest(zoneName string, ) *ApiRgwZonePostRequest`

NewApiRgwZonePostRequest instantiates a new ApiRgwZonePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwZonePostRequestWithDefaults

`func NewApiRgwZonePostRequestWithDefaults() *ApiRgwZonePostRequest`

NewApiRgwZonePostRequestWithDefaults instantiates a new ApiRgwZonePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *ApiRgwZonePostRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ApiRgwZonePostRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ApiRgwZonePostRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ApiRgwZonePostRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetDefault

`func (o *ApiRgwZonePostRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ApiRgwZonePostRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ApiRgwZonePostRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ApiRgwZonePostRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetMaster

`func (o *ApiRgwZonePostRequest) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ApiRgwZonePostRequest) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ApiRgwZonePostRequest) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ApiRgwZonePostRequest) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetSecretKey

`func (o *ApiRgwZonePostRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ApiRgwZonePostRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ApiRgwZonePostRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ApiRgwZonePostRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetZoneEndpoints

`func (o *ApiRgwZonePostRequest) GetZoneEndpoints() string`

GetZoneEndpoints returns the ZoneEndpoints field if non-nil, zero value otherwise.

### GetZoneEndpointsOk

`func (o *ApiRgwZonePostRequest) GetZoneEndpointsOk() (*string, bool)`

GetZoneEndpointsOk returns a tuple with the ZoneEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneEndpoints

`func (o *ApiRgwZonePostRequest) SetZoneEndpoints(v string)`

SetZoneEndpoints sets ZoneEndpoints field to given value.

### HasZoneEndpoints

`func (o *ApiRgwZonePostRequest) HasZoneEndpoints() bool`

HasZoneEndpoints returns a boolean if a field has been set.

### GetZoneName

`func (o *ApiRgwZonePostRequest) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ApiRgwZonePostRequest) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ApiRgwZonePostRequest) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetZonegroupName

`func (o *ApiRgwZonePostRequest) GetZonegroupName() string`

GetZonegroupName returns the ZonegroupName field if non-nil, zero value otherwise.

### GetZonegroupNameOk

`func (o *ApiRgwZonePostRequest) GetZonegroupNameOk() (*string, bool)`

GetZonegroupNameOk returns a tuple with the ZonegroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonegroupName

`func (o *ApiRgwZonePostRequest) SetZonegroupName(v string)`

SetZonegroupName sets ZonegroupName field to given value.

### HasZonegroupName

`func (o *ApiRgwZonePostRequest) HasZonegroupName() bool`

HasZonegroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


