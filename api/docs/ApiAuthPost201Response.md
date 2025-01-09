# ApiAuthPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**ApiAuthPost201ResponsePermissions**](ApiAuthPost201ResponsePermissions.md) |  | 
**PwdExpirationDate** | **string** | Password expiration date | 
**PwdUpdateRequired** | **bool** | Is password update required? | 
**Sso** | **bool** | Uses single sign on? | 
**Token** | **string** | Authentication Token | 
**Username** | **string** | Username | 

## Methods

### NewApiAuthPost201Response

`func NewApiAuthPost201Response(permissions ApiAuthPost201ResponsePermissions, pwdExpirationDate string, pwdUpdateRequired bool, sso bool, token string, username string, ) *ApiAuthPost201Response`

NewApiAuthPost201Response instantiates a new ApiAuthPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAuthPost201ResponseWithDefaults

`func NewApiAuthPost201ResponseWithDefaults() *ApiAuthPost201Response`

NewApiAuthPost201ResponseWithDefaults instantiates a new ApiAuthPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *ApiAuthPost201Response) GetPermissions() ApiAuthPost201ResponsePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiAuthPost201Response) GetPermissionsOk() (*ApiAuthPost201ResponsePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiAuthPost201Response) SetPermissions(v ApiAuthPost201ResponsePermissions)`

SetPermissions sets Permissions field to given value.


### GetPwdExpirationDate

`func (o *ApiAuthPost201Response) GetPwdExpirationDate() string`

GetPwdExpirationDate returns the PwdExpirationDate field if non-nil, zero value otherwise.

### GetPwdExpirationDateOk

`func (o *ApiAuthPost201Response) GetPwdExpirationDateOk() (*string, bool)`

GetPwdExpirationDateOk returns a tuple with the PwdExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdExpirationDate

`func (o *ApiAuthPost201Response) SetPwdExpirationDate(v string)`

SetPwdExpirationDate sets PwdExpirationDate field to given value.


### GetPwdUpdateRequired

`func (o *ApiAuthPost201Response) GetPwdUpdateRequired() bool`

GetPwdUpdateRequired returns the PwdUpdateRequired field if non-nil, zero value otherwise.

### GetPwdUpdateRequiredOk

`func (o *ApiAuthPost201Response) GetPwdUpdateRequiredOk() (*bool, bool)`

GetPwdUpdateRequiredOk returns a tuple with the PwdUpdateRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdUpdateRequired

`func (o *ApiAuthPost201Response) SetPwdUpdateRequired(v bool)`

SetPwdUpdateRequired sets PwdUpdateRequired field to given value.


### GetSso

`func (o *ApiAuthPost201Response) GetSso() bool`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *ApiAuthPost201Response) GetSsoOk() (*bool, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *ApiAuthPost201Response) SetSso(v bool)`

SetSso sets Sso field to given value.


### GetToken

`func (o *ApiAuthPost201Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiAuthPost201Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiAuthPost201Response) SetToken(v string)`

SetToken sets Token field to given value.


### GetUsername

`func (o *ApiAuthPost201Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiAuthPost201Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiAuthPost201Response) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


