# ApiAuthPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Password | 
**Ttl** | Pointer to **int32** | Token Time to Live (in hours) | [optional] 
**Username** | **string** | Username | 

## Methods

### NewApiAuthPostRequest

`func NewApiAuthPostRequest(password string, username string, ) *ApiAuthPostRequest`

NewApiAuthPostRequest instantiates a new ApiAuthPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAuthPostRequestWithDefaults

`func NewApiAuthPostRequestWithDefaults() *ApiAuthPostRequest`

NewApiAuthPostRequestWithDefaults instantiates a new ApiAuthPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ApiAuthPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiAuthPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiAuthPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTtl

`func (o *ApiAuthPostRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ApiAuthPostRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ApiAuthPostRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ApiAuthPostRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUsername

`func (o *ApiAuthPostRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiAuthPostRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiAuthPostRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


