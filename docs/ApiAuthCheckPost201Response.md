# ApiAuthCheckPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**ApiAuthPost201ResponsePermissions**](ApiAuthPost201ResponsePermissions.md) |  | 
**PwdUpdateRequired** | **bool** | Is password update required? | 
**Sso** | **bool** | Uses single sign on? | 
**Username** | **string** | Username | 

## Methods

### NewApiAuthCheckPost201Response

`func NewApiAuthCheckPost201Response(permissions ApiAuthPost201ResponsePermissions, pwdUpdateRequired bool, sso bool, username string, ) *ApiAuthCheckPost201Response`

NewApiAuthCheckPost201Response instantiates a new ApiAuthCheckPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAuthCheckPost201ResponseWithDefaults

`func NewApiAuthCheckPost201ResponseWithDefaults() *ApiAuthCheckPost201Response`

NewApiAuthCheckPost201ResponseWithDefaults instantiates a new ApiAuthCheckPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *ApiAuthCheckPost201Response) GetPermissions() ApiAuthPost201ResponsePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiAuthCheckPost201Response) GetPermissionsOk() (*ApiAuthPost201ResponsePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiAuthCheckPost201Response) SetPermissions(v ApiAuthPost201ResponsePermissions)`

SetPermissions sets Permissions field to given value.


### GetPwdUpdateRequired

`func (o *ApiAuthCheckPost201Response) GetPwdUpdateRequired() bool`

GetPwdUpdateRequired returns the PwdUpdateRequired field if non-nil, zero value otherwise.

### GetPwdUpdateRequiredOk

`func (o *ApiAuthCheckPost201Response) GetPwdUpdateRequiredOk() (*bool, bool)`

GetPwdUpdateRequiredOk returns a tuple with the PwdUpdateRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwdUpdateRequired

`func (o *ApiAuthCheckPost201Response) SetPwdUpdateRequired(v bool)`

SetPwdUpdateRequired sets PwdUpdateRequired field to given value.


### GetSso

`func (o *ApiAuthCheckPost201Response) GetSso() bool`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *ApiAuthCheckPost201Response) GetSsoOk() (*bool, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *ApiAuthCheckPost201Response) SetSso(v bool)`

SetSso sets Sso field to given value.


### GetUsername

`func (o *ApiAuthCheckPost201Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiAuthCheckPost201Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiAuthCheckPost201Response) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


