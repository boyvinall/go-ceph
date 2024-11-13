# ApiUserUsernamePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PwdExpirationDate** | Pointer to **string** |  | [optional] 
**PwdUpdateRequired** | Pointer to **bool** |  | [optional] [default to false]
**Roles** | Pointer to **string** |  | [optional] 

## Methods

### NewApiUserUsernamePutRequest

`func NewApiUserUsernamePutRequest() *ApiUserUsernamePutRequest`

NewApiUserUsernamePutRequest instantiates a new ApiUserUsernamePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserUsernamePutRequestWithDefaults

`func NewApiUserUsernamePutRequestWithDefaults() *ApiUserUsernamePutRequest`

NewApiUserUsernamePutRequestWithDefaults instantiates a new ApiUserUsernamePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ApiUserUsernamePutRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApiUserUsernamePutRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApiUserUsernamePutRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApiUserUsernamePutRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiUserUsernamePutRequest) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiUserUsernamePutRequest) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiUserUsernamePutRequest) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiUserUsernamePutRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *ApiUserUsernamePutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiUserUsernamePutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiUserUsernamePutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiUserUsernamePutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *ApiUserUsernamePutRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiUserUsernamePutRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiUserUsernamePutRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiUserUsernamePutRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPwdExpirationDate

`func (o *ApiUserUsernamePutRequest) GetPwdExpirationDate() string`

GetPwdExpirationDate returns the PwdExpirationDate field if non-nil, zero value otherwise.

### GetPwdExpirationDateOk

`func (o *ApiUserUsernamePutRequest) GetPwdExpirationDateOk() (*string, bool)`

GetPwdExpirationDateOk returns a tuple with the PwdExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdExpirationDate

`func (o *ApiUserUsernamePutRequest) SetPwdExpirationDate(v string)`

SetPwdExpirationDate sets PwdExpirationDate field to given value.

### HasPwdExpirationDate

`func (o *ApiUserUsernamePutRequest) HasPwdExpirationDate() bool`

HasPwdExpirationDate returns a boolean if a field has been set.

### GetPwdUpdateRequired

`func (o *ApiUserUsernamePutRequest) GetPwdUpdateRequired() bool`

GetPwdUpdateRequired returns the PwdUpdateRequired field if non-nil, zero value otherwise.

### GetPwdUpdateRequiredOk

`func (o *ApiUserUsernamePutRequest) GetPwdUpdateRequiredOk() (*bool, bool)`

GetPwdUpdateRequiredOk returns a tuple with the PwdUpdateRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdUpdateRequired

`func (o *ApiUserUsernamePutRequest) SetPwdUpdateRequired(v bool)`

SetPwdUpdateRequired sets PwdUpdateRequired field to given value.

### HasPwdUpdateRequired

`func (o *ApiUserUsernamePutRequest) HasPwdUpdateRequired() bool`

HasPwdUpdateRequired returns a boolean if a field has been set.

### GetRoles

`func (o *ApiUserUsernamePutRequest) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiUserUsernamePutRequest) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiUserUsernamePutRequest) SetRoles(v string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiUserUsernamePutRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


