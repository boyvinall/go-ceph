# ApiUserGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | User email address | 
**Enabled** | **bool** | Is the user enabled? | 
**LastUpdate** | **int32** | Details last updated | 
**Name** | **string** | User Name | 
**PwdExpirationDate** | **string** | Password Expiration date | 
**PwdUpdateRequired** | **bool** | Is Password Update Required? | 
**Roles** | **[]string** | User Roles | 
**Username** | **string** | Username of the user | 

## Methods

### NewApiUserGet200Response

`func NewApiUserGet200Response(email string, enabled bool, lastUpdate int32, name string, pwdExpirationDate string, pwdUpdateRequired bool, roles []string, username string, ) *ApiUserGet200Response`

NewApiUserGet200Response instantiates a new ApiUserGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserGet200ResponseWithDefaults

`func NewApiUserGet200ResponseWithDefaults() *ApiUserGet200Response`

NewApiUserGet200ResponseWithDefaults instantiates a new ApiUserGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ApiUserGet200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApiUserGet200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApiUserGet200Response) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnabled

`func (o *ApiUserGet200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiUserGet200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiUserGet200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLastUpdate

`func (o *ApiUserGet200Response) GetLastUpdate() int32`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *ApiUserGet200Response) GetLastUpdateOk() (*int32, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *ApiUserGet200Response) SetLastUpdate(v int32)`

SetLastUpdate sets LastUpdate field to given value.


### GetName

`func (o *ApiUserGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiUserGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiUserGet200Response) SetName(v string)`

SetName sets Name field to given value.


### GetPwdExpirationDate

`func (o *ApiUserGet200Response) GetPwdExpirationDate() string`

GetPwdExpirationDate returns the PwdExpirationDate field if non-nil, zero value otherwise.

### GetPwdExpirationDateOk

`func (o *ApiUserGet200Response) GetPwdExpirationDateOk() (*string, bool)`

GetPwdExpirationDateOk returns a tuple with the PwdExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdExpirationDate

`func (o *ApiUserGet200Response) SetPwdExpirationDate(v string)`

SetPwdExpirationDate sets PwdExpirationDate field to given value.


### GetPwdUpdateRequired

`func (o *ApiUserGet200Response) GetPwdUpdateRequired() bool`

GetPwdUpdateRequired returns the PwdUpdateRequired field if non-nil, zero value otherwise.

### GetPwdUpdateRequiredOk

`func (o *ApiUserGet200Response) GetPwdUpdateRequiredOk() (*bool, bool)`

GetPwdUpdateRequiredOk returns a tuple with the PwdUpdateRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdUpdateRequired

`func (o *ApiUserGet200Response) SetPwdUpdateRequired(v bool)`

SetPwdUpdateRequired sets PwdUpdateRequired field to given value.


### GetRoles

`func (o *ApiUserGet200Response) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiUserGet200Response) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiUserGet200Response) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetUsername

`func (o *ApiUserGet200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiUserGet200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiUserGet200Response) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


