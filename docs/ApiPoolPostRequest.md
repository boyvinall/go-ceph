# ApiPoolPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pool** | Pointer to **string** |  | [optional] [default to "rbd-mirror"]

## Methods

### NewApiPoolPostRequest

`func NewApiPoolPostRequest() *ApiPoolPostRequest`

NewApiPoolPostRequest instantiates a new ApiPoolPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPoolPostRequestWithDefaults

`func NewApiPoolPostRequestWithDefaults() *ApiPoolPostRequest`

NewApiPoolPostRequestWithDefaults instantiates a new ApiPoolPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPool

`func (o *ApiPoolPostRequest) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ApiPoolPostRequest) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ApiPoolPostRequest) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ApiPoolPostRequest) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


