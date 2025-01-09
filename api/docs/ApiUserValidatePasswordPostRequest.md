# ApiUserValidatePasswordPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldPassword** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewApiUserValidatePasswordPostRequest

`func NewApiUserValidatePasswordPostRequest(password string, ) *ApiUserValidatePasswordPostRequest`

NewApiUserValidatePasswordPostRequest instantiates a new ApiUserValidatePasswordPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserValidatePasswordPostRequestWithDefaults

`func NewApiUserValidatePasswordPostRequestWithDefaults() *ApiUserValidatePasswordPostRequest`

NewApiUserValidatePasswordPostRequestWithDefaults instantiates a new ApiUserValidatePasswordPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldPassword

`func (o *ApiUserValidatePasswordPostRequest) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *ApiUserValidatePasswordPostRequest) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *ApiUserValidatePasswordPostRequest) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *ApiUserValidatePasswordPostRequest) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.

### GetPassword

`func (o *ApiUserValidatePasswordPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiUserValidatePasswordPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiUserValidatePasswordPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *ApiUserValidatePasswordPostRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiUserValidatePasswordPostRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiUserValidatePasswordPostRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiUserValidatePasswordPostRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


