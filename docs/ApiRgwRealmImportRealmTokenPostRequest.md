# ApiRgwRealmImportRealmTokenPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlacementSpec** | Pointer to **string** |  | [optional] 
**Port** | **string** |  | 
**RealmToken** | **string** |  | 
**ZoneName** | **string** |  | 

## Methods

### NewApiRgwRealmImportRealmTokenPostRequest

`func NewApiRgwRealmImportRealmTokenPostRequest(port string, realmToken string, zoneName string, ) *ApiRgwRealmImportRealmTokenPostRequest`

NewApiRgwRealmImportRealmTokenPostRequest instantiates a new ApiRgwRealmImportRealmTokenPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRgwRealmImportRealmTokenPostRequestWithDefaults

`func NewApiRgwRealmImportRealmTokenPostRequestWithDefaults() *ApiRgwRealmImportRealmTokenPostRequest`

NewApiRgwRealmImportRealmTokenPostRequestWithDefaults instantiates a new ApiRgwRealmImportRealmTokenPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacementSpec

`func (o *ApiRgwRealmImportRealmTokenPostRequest) GetPlacementSpec() string`

GetPlacementSpec returns the PlacementSpec field if non-nil, zero value otherwise.

### GetPlacementSpecOk

`func (o *ApiRgwRealmImportRealmTokenPostRequest) GetPlacementSpecOk() (*string, bool)`

GetPlacementSpecOk returns a tuple with the PlacementSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementSpec

`func (o *ApiRgwRealmImportRealmTokenPostRequest) SetPlacementSpec(v string)`

SetPlacementSpec sets PlacementSpec field to given value.

### HasPlacementSpec

`func (o *ApiRgwRealmImportRealmTokenPostRequest) HasPlacementSpec() bool`

HasPlacementSpec returns a boolean if a field has been set.

### GetPort

`func (o *ApiRgwRealmImportRealmTokenPostRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiRgwRealmImportRealmTokenPostRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiRgwRealmImportRealmTokenPostRequest) SetPort(v string)`

SetPort sets Port field to given value.


### GetRealmToken

`func (o *ApiRgwRealmImportRealmTokenPostRequest) GetRealmToken() string`

GetRealmToken returns the RealmToken field if non-nil, zero value otherwise.

### GetRealmTokenOk

`func (o *ApiRgwRealmImportRealmTokenPostRequest) GetRealmTokenOk() (*string, bool)`

GetRealmTokenOk returns a tuple with the RealmToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmToken

`func (o *ApiRgwRealmImportRealmTokenPostRequest) SetRealmToken(v string)`

SetRealmToken sets RealmToken field to given value.


### GetZoneName

`func (o *ApiRgwRealmImportRealmTokenPostRequest) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ApiRgwRealmImportRealmTokenPostRequest) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ApiRgwRealmImportRealmTokenPostRequest) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


