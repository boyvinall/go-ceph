# ApiCephfsRenamePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Existing FS Name | 
**NewName** | **string** | New FS Name | 

## Methods

### NewApiCephfsRenamePutRequest

`func NewApiCephfsRenamePutRequest(name string, newName string, ) *ApiCephfsRenamePutRequest`

NewApiCephfsRenamePutRequest instantiates a new ApiCephfsRenamePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCephfsRenamePutRequestWithDefaults

`func NewApiCephfsRenamePutRequestWithDefaults() *ApiCephfsRenamePutRequest`

NewApiCephfsRenamePutRequestWithDefaults instantiates a new ApiCephfsRenamePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiCephfsRenamePutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiCephfsRenamePutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiCephfsRenamePutRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *ApiCephfsRenamePutRequest) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *ApiCephfsRenamePutRequest) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *ApiCephfsRenamePutRequest) SetNewName(v string)`

SetNewName sets NewName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


