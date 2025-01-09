# ApiRgwZonegroupZonegroupNamePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddZones** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **string** |  | [optional] [default to ""]
**Master** | Pointer to **string** |  | [optional] [default to ""]
**NewZonegroupName** | **string** |  | 
**PlacementTargets** | Pointer to **string** |  | [optional] 
**RealmName** | **string** |  | 
**RemoveZones** | Pointer to **string** |  | [optional] 
**ZonegroupEndpoints** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewApiRgwZonegroupZonegroupNamePutRequest

`func NewApiRgwZonegroupZonegroupNamePutRequest(newZonegroupName string, realmName string, ) *ApiRgwZonegroupZonegroupNamePutRequest`

NewApiRgwZonegroupZonegroupNamePutRequest instantiates a new ApiRgwZonegroupZonegroupNamePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwZonegroupZonegroupNamePutRequestWithDefaults

`func NewApiRgwZonegroupZonegroupNamePutRequestWithDefaults() *ApiRgwZonegroupZonegroupNamePutRequest`

NewApiRgwZonegroupZonegroupNamePutRequestWithDefaults instantiates a new ApiRgwZonegroupZonegroupNamePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddZones

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetAddZones() string`

GetAddZones returns the AddZones field if non-nil, zero value otherwise.

### GetAddZonesOk

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetAddZonesOk() (*string, bool)`

GetAddZonesOk returns a tuple with the AddZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddZones

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetAddZones(v string)`

SetAddZones sets AddZones field to given value.

### HasAddZones

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasAddZones() bool`

HasAddZones returns a boolean if a field has been set.

### GetDefault

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetMaster

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetMaster() string`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetMasterOk() (*string, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetMaster(v string)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetNewZonegroupName

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetNewZonegroupName() string`

GetNewZonegroupName returns the NewZonegroupName field if non-nil, zero value otherwise.

### GetNewZonegroupNameOk

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetNewZonegroupNameOk() (*string, bool)`

GetNewZonegroupNameOk returns a tuple with the NewZonegroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewZonegroupName

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetNewZonegroupName(v string)`

SetNewZonegroupName sets NewZonegroupName field to given value.


### GetPlacementTargets

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetPlacementTargets() string`

GetPlacementTargets returns the PlacementTargets field if non-nil, zero value otherwise.

### GetPlacementTargetsOk

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetPlacementTargetsOk() (*string, bool)`

GetPlacementTargetsOk returns a tuple with the PlacementTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementTargets

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetPlacementTargets(v string)`

SetPlacementTargets sets PlacementTargets field to given value.

### HasPlacementTargets

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasPlacementTargets() bool`

HasPlacementTargets returns a boolean if a field has been set.

### GetRealmName

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.


### GetRemoveZones

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetRemoveZones() string`

GetRemoveZones returns the RemoveZones field if non-nil, zero value otherwise.

### GetRemoveZonesOk

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetRemoveZonesOk() (*string, bool)`

GetRemoveZonesOk returns a tuple with the RemoveZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveZones

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetRemoveZones(v string)`

SetRemoveZones sets RemoveZones field to given value.

### HasRemoveZones

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasRemoveZones() bool`

HasRemoveZones returns a boolean if a field has been set.

### GetZonegroupEndpoints

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetZonegroupEndpoints() string`

GetZonegroupEndpoints returns the ZonegroupEndpoints field if non-nil, zero value otherwise.

### GetZonegroupEndpointsOk

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) GetZonegroupEndpointsOk() (*string, bool)`

GetZonegroupEndpointsOk returns a tuple with the ZonegroupEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonegroupEndpoints

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) SetZonegroupEndpoints(v string)`

SetZonegroupEndpoints sets ZonegroupEndpoints field to given value.

### HasZonegroupEndpoints

`func (o *ApiRgwZonegroupZonegroupNamePutRequest) HasZonegroupEndpoints() bool`

HasZonegroupEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


