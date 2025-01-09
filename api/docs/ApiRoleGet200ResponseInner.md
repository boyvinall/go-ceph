# ApiRoleGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Role Descriptions | [optional] 
**Name** | Pointer to **string** | Role Name | [optional] 
**ScopesPermissions** | Pointer to [**ApiRoleGet200ResponseInnerScopesPermissions**](ApiRoleGet200ResponseInnerScopesPermissions.md) |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiRoleGet200ResponseInner

`func NewApiRoleGet200ResponseInner() *ApiRoleGet200ResponseInner`

NewApiRoleGet200ResponseInner instantiates a new ApiRoleGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRoleGet200ResponseInnerWithDefaults

`func NewApiRoleGet200ResponseInnerWithDefaults() *ApiRoleGet200ResponseInner`

NewApiRoleGet200ResponseInnerWithDefaults instantiates a new ApiRoleGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApiRoleGet200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiRoleGet200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiRoleGet200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiRoleGet200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ApiRoleGet200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiRoleGet200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiRoleGet200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiRoleGet200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopesPermissions

`func (o *ApiRoleGet200ResponseInner) GetScopesPermissions() ApiRoleGet200ResponseInnerScopesPermissions`

GetScopesPermissions returns the ScopesPermissions field if non-nil, zero value otherwise.

### GetScopesPermissionsOk

`func (o *ApiRoleGet200ResponseInner) GetScopesPermissionsOk() (*ApiRoleGet200ResponseInnerScopesPermissions, bool)`

GetScopesPermissionsOk returns a tuple with the ScopesPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesPermissions

`func (o *ApiRoleGet200ResponseInner) SetScopesPermissions(v ApiRoleGet200ResponseInnerScopesPermissions)`

SetScopesPermissions sets ScopesPermissions field to given value.

### HasScopesPermissions

`func (o *ApiRoleGet200ResponseInner) HasScopesPermissions() bool`

HasScopesPermissions returns a boolean if a field has been set.

### GetSystem

`func (o *ApiRoleGet200ResponseInner) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ApiRoleGet200ResponseInner) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ApiRoleGet200ResponseInner) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ApiRoleGet200ResponseInner) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


