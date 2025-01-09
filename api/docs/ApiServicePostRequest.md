# ApiServicePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** |  | 
**ServiceSpec** | **string** |  | 

## Methods

### NewApiServicePostRequest

`func NewApiServicePostRequest(serviceName string, serviceSpec string, ) *ApiServicePostRequest`

NewApiServicePostRequest instantiates a new ApiServicePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServicePostRequestWithDefaults

`func NewApiServicePostRequestWithDefaults() *ApiServicePostRequest`

NewApiServicePostRequestWithDefaults instantiates a new ApiServicePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *ApiServicePostRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ApiServicePostRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ApiServicePostRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServiceSpec

`func (o *ApiServicePostRequest) GetServiceSpec() string`

GetServiceSpec returns the ServiceSpec field if non-nil, zero value otherwise.

### GetServiceSpecOk

`func (o *ApiServicePostRequest) GetServiceSpecOk() (*string, bool)`

GetServiceSpecOk returns a tuple with the ServiceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSpec

`func (o *ApiServicePostRequest) SetServiceSpec(v string)`

SetServiceSpec sets ServiceSpec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


