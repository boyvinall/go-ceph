# ApiTelemetryPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] [default to true]
**LicenseName** | Pointer to **string** |  | [optional] 

## Methods

### NewApiTelemetryPutRequest

`func NewApiTelemetryPutRequest() *ApiTelemetryPutRequest`

NewApiTelemetryPutRequest instantiates a new ApiTelemetryPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTelemetryPutRequestWithDefaults

`func NewApiTelemetryPutRequestWithDefaults() *ApiTelemetryPutRequest`

NewApiTelemetryPutRequestWithDefaults instantiates a new ApiTelemetryPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *ApiTelemetryPutRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ApiTelemetryPutRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ApiTelemetryPutRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ApiTelemetryPutRequest) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetLicenseName

`func (o *ApiTelemetryPutRequest) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *ApiTelemetryPutRequest) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *ApiTelemetryPutRequest) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *ApiTelemetryPutRequest) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


