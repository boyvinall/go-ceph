# ApiHostHostnamePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drain** | Pointer to **bool** | Drain Host | [optional] [default to false]
**Force** | Pointer to **bool** | Force Enter Maintenance | [optional] [default to false]
**Labels** | Pointer to **[]string** | Host Labels | [optional] 
**Maintenance** | Pointer to **bool** | Enter/Exit Maintenance | [optional] [default to false]
**UpdateLabels** | Pointer to **bool** | Update Labels | [optional] [default to false]

## Methods

### NewApiHostHostnamePutRequest

`func NewApiHostHostnamePutRequest() *ApiHostHostnamePutRequest`

NewApiHostHostnamePutRequest instantiates a new ApiHostHostnamePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHostHostnamePutRequestWithDefaults

`func NewApiHostHostnamePutRequestWithDefaults() *ApiHostHostnamePutRequest`

NewApiHostHostnamePutRequestWithDefaults instantiates a new ApiHostHostnamePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrain

`func (o *ApiHostHostnamePutRequest) GetDrain() bool`

GetDrain returns the Drain field if non-nil, zero value otherwise.

### GetDrainOk

`func (o *ApiHostHostnamePutRequest) GetDrainOk() (*bool, bool)`

GetDrainOk returns a tuple with the Drain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrain

`func (o *ApiHostHostnamePutRequest) SetDrain(v bool)`

SetDrain sets Drain field to given value.

### HasDrain

`func (o *ApiHostHostnamePutRequest) HasDrain() bool`

HasDrain returns a boolean if a field has been set.

### GetForce

`func (o *ApiHostHostnamePutRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *ApiHostHostnamePutRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *ApiHostHostnamePutRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *ApiHostHostnamePutRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetLabels

`func (o *ApiHostHostnamePutRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiHostHostnamePutRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiHostHostnamePutRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiHostHostnamePutRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMaintenance

`func (o *ApiHostHostnamePutRequest) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *ApiHostHostnamePutRequest) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *ApiHostHostnamePutRequest) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *ApiHostHostnamePutRequest) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### GetUpdateLabels

`func (o *ApiHostHostnamePutRequest) GetUpdateLabels() bool`

GetUpdateLabels returns the UpdateLabels field if non-nil, zero value otherwise.

### GetUpdateLabelsOk

`func (o *ApiHostHostnamePutRequest) GetUpdateLabelsOk() (*bool, bool)`

GetUpdateLabelsOk returns a tuple with the UpdateLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLabels

`func (o *ApiHostHostnamePutRequest) SetUpdateLabels(v bool)`

SetUpdateLabels sets UpdateLabels field to given value.

### HasUpdateLabels

`func (o *ApiHostHostnamePutRequest) HasUpdateLabels() bool`

HasUpdateLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


