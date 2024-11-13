# ApiLogsAllGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLog** | [**[]ApiLogsAllGet200ResponseAuditLogInner**](ApiLogsAllGet200ResponseAuditLogInner.md) | Audit log | 
**Clog** | **[]string** |  | 

## Methods

### NewApiLogsAllGet200Response

`func NewApiLogsAllGet200Response(auditLog []ApiLogsAllGet200ResponseAuditLogInner, clog []string, ) *ApiLogsAllGet200Response`

NewApiLogsAllGet200Response instantiates a new ApiLogsAllGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLogsAllGet200ResponseWithDefaults

`func NewApiLogsAllGet200ResponseWithDefaults() *ApiLogsAllGet200Response`

NewApiLogsAllGet200ResponseWithDefaults instantiates a new ApiLogsAllGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLog

`func (o *ApiLogsAllGet200Response) GetAuditLog() []ApiLogsAllGet200ResponseAuditLogInner`

GetAuditLog returns the AuditLog field if non-nil, zero value otherwise.

### GetAuditLogOk

`func (o *ApiLogsAllGet200Response) GetAuditLogOk() (*[]ApiLogsAllGet200ResponseAuditLogInner, bool)`

GetAuditLogOk returns a tuple with the AuditLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLog

`func (o *ApiLogsAllGet200Response) SetAuditLog(v []ApiLogsAllGet200ResponseAuditLogInner)`

SetAuditLog sets AuditLog field to given value.


### GetClog

`func (o *ApiLogsAllGet200Response) GetClog() []string`

GetClog returns the Clog field if non-nil, zero value otherwise.

### GetClogOk

`func (o *ApiLogsAllGet200Response) GetClogOk() (*[]string, bool)`

GetClogOk returns a tuple with the Clog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClog

`func (o *ApiLogsAllGet200Response) SetClog(v []string)`

SetClog sets Clog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


