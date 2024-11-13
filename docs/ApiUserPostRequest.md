# ApiUserPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PwdExpirationDate** | Pointer to **string** |  | [optional] 
**PwdUpdateRequired** | Pointer to **bool** |  | [optional] [default to true]
**Roles** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewApiUserPostRequest

`func NewApiUserPostRequest() *ApiUserPostRequest`

NewApiUserPostRequest instantiates a new ApiUserPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserPostRequestWithDefaults

`func NewApiUserPostRequestWithDefaults() *ApiUserPostRequest`

NewApiUserPostRequestWithDefaults instantiates a new ApiUserPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ApiUserPostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApiUserPostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApiUserPostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApiUserPostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiUserPostRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiUserPostRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiUserPostRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiUserPostRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *ApiUserPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiUserPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiUserPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiUserPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *ApiUserPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiUserPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiUserPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiUserPostRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPwdExpirationDate

`func (o *ApiUserPostRequest) GetPwdExpirationDate() string`

GetPwdExpirationDate returns the PwdExpirationDate field if non-nil, zero value otherwise.

### GetPwdExpirationDateOk

`func (o *ApiUserPostRequest) GetPwdExpirationDateOk() (*string, bool)`

GetPwdExpirationDateOk returns a tuple with the PwdExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdExpirationDate

`func (o *ApiUserPostRequest) SetPwdExpirationDate(v string)`

SetPwdExpirationDate sets PwdExpirationDate field to given value.

### HasPwdExpirationDate

`func (o *ApiUserPostRequest) HasPwdExpirationDate() bool`

HasPwdExpirationDate returns a boolean if a field has been set.

### GetPwdUpdateRequired

`func (o *ApiUserPostRequest) GetPwdUpdateRequired() bool`

GetPwdUpdateRequired returns the PwdUpdateRequired field if non-nil, zero value otherwise.

### GetPwdUpdateRequiredOk

`func (o *ApiUserPostRequest) GetPwdUpdateRequiredOk() (*bool, bool)`

GetPwdUpdateRequiredOk returns a tuple with the PwdUpdateRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdUpdateRequired

`func (o *ApiUserPostRequest) SetPwdUpdateRequired(v bool)`

SetPwdUpdateRequired sets PwdUpdateRequired field to given value.

### HasPwdUpdateRequired

`func (o *ApiUserPostRequest) HasPwdUpdateRequired() bool`

HasPwdUpdateRequired returns a boolean if a field has been set.

### GetRoles

`func (o *ApiUserPostRequest) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiUserPostRequest) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiUserPostRequest) SetRoles(v string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiUserPostRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUsername

`func (o *ApiUserPostRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiUserPostRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiUserPostRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiUserPostRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


