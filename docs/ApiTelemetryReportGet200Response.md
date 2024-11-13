# ApiTelemetryReportGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceReport** | **string** |  | 
**Report** | [**ApiTelemetryReportGet200ResponseReport**](ApiTelemetryReportGet200ResponseReport.md) |  | 

## Methods

### NewApiTelemetryReportGet200Response

`func NewApiTelemetryReportGet200Response(deviceReport string, report ApiTelemetryReportGet200ResponseReport, ) *ApiTelemetryReportGet200Response`

NewApiTelemetryReportGet200Response instantiates a new ApiTelemetryReportGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTelemetryReportGet200ResponseWithDefaults

`func NewApiTelemetryReportGet200ResponseWithDefaults() *ApiTelemetryReportGet200Response`

NewApiTelemetryReportGet200ResponseWithDefaults instantiates a new ApiTelemetryReportGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceReport

`func (o *ApiTelemetryReportGet200Response) GetDeviceReport() string`

GetDeviceReport returns the DeviceReport field if non-nil, zero value otherwise.

### GetDeviceReportOk

`func (o *ApiTelemetryReportGet200Response) GetDeviceReportOk() (*string, bool)`

GetDeviceReportOk returns a tuple with the DeviceReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceReport

`func (o *ApiTelemetryReportGet200Response) SetDeviceReport(v string)`

SetDeviceReport sets DeviceReport field to given value.


### GetReport

`func (o *ApiTelemetryReportGet200Response) GetReport() ApiTelemetryReportGet200ResponseReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *ApiTelemetryReportGet200Response) GetReportOk() (*ApiTelemetryReportGet200ResponseReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *ApiTelemetryReportGet200Response) SetReport(v ApiTelemetryReportGet200ResponseReport)`

SetReport sets Report field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


