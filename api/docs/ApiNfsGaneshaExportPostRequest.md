# ApiNfsGaneshaExportPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | **string** | Export access type | 
**Clients** | [**[]ApiNfsGaneshaExportGet200ResponseInnerClientsInner**](ApiNfsGaneshaExportGet200ResponseInnerClientsInner.md) | List of client configurations | 
**ClusterId** | **string** | Cluster identifier | 
**Fsal** | [**ApiNfsGaneshaExportPostRequestFsal**](ApiNfsGaneshaExportPostRequestFsal.md) |  | 
**Path** | **string** | Export path | 
**Protocols** | **[]int32** | List of protocol types | 
**Pseudo** | **string** | Pseudo FS path | 
**SecurityLabel** | **string** | Security label | 
**Squash** | **string** | Export squash policy | 
**Transports** | **[]string** | List of transport types | 

## Methods

### NewApiNfsGaneshaExportPostRequest

`func NewApiNfsGaneshaExportPostRequest(accessType string, clients []ApiNfsGaneshaExportGet200ResponseInnerClientsInner, clusterId string, fsal ApiNfsGaneshaExportPostRequestFsal, path string, protocols []int32, pseudo string, securityLabel string, squash string, transports []string, ) *ApiNfsGaneshaExportPostRequest`

NewApiNfsGaneshaExportPostRequest instantiates a new ApiNfsGaneshaExportPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNfsGaneshaExportPostRequestWithDefaults

`func NewApiNfsGaneshaExportPostRequestWithDefaults() *ApiNfsGaneshaExportPostRequest`

NewApiNfsGaneshaExportPostRequestWithDefaults instantiates a new ApiNfsGaneshaExportPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *ApiNfsGaneshaExportPostRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *ApiNfsGaneshaExportPostRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *ApiNfsGaneshaExportPostRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetClients

`func (o *ApiNfsGaneshaExportPostRequest) GetClients() []ApiNfsGaneshaExportGet200ResponseInnerClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ApiNfsGaneshaExportPostRequest) GetClientsOk() (*[]ApiNfsGaneshaExportGet200ResponseInnerClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ApiNfsGaneshaExportPostRequest) SetClients(v []ApiNfsGaneshaExportGet200ResponseInnerClientsInner)`

SetClients sets Clients field to given value.


### GetClusterId

`func (o *ApiNfsGaneshaExportPostRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ApiNfsGaneshaExportPostRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ApiNfsGaneshaExportPostRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetFsal

`func (o *ApiNfsGaneshaExportPostRequest) GetFsal() ApiNfsGaneshaExportPostRequestFsal`

GetFsal returns the Fsal field if non-nil, zero value otherwise.

### GetFsalOk

`func (o *ApiNfsGaneshaExportPostRequest) GetFsalOk() (*ApiNfsGaneshaExportPostRequestFsal, bool)`

GetFsalOk returns a tuple with the Fsal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsal

`func (o *ApiNfsGaneshaExportPostRequest) SetFsal(v ApiNfsGaneshaExportPostRequestFsal)`

SetFsal sets Fsal field to given value.


### GetPath

`func (o *ApiNfsGaneshaExportPostRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiNfsGaneshaExportPostRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiNfsGaneshaExportPostRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetProtocols

`func (o *ApiNfsGaneshaExportPostRequest) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ApiNfsGaneshaExportPostRequest) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ApiNfsGaneshaExportPostRequest) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetPseudo

`func (o *ApiNfsGaneshaExportPostRequest) GetPseudo() string`

GetPseudo returns the Pseudo field if non-nil, zero value otherwise.

### GetPseudoOk

`func (o *ApiNfsGaneshaExportPostRequest) GetPseudoOk() (*string, bool)`

GetPseudoOk returns a tuple with the Pseudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudo

`func (o *ApiNfsGaneshaExportPostRequest) SetPseudo(v string)`

SetPseudo sets Pseudo field to given value.


### GetSecurityLabel

`func (o *ApiNfsGaneshaExportPostRequest) GetSecurityLabel() string`

GetSecurityLabel returns the SecurityLabel field if non-nil, zero value otherwise.

### GetSecurityLabelOk

`func (o *ApiNfsGaneshaExportPostRequest) GetSecurityLabelOk() (*string, bool)`

GetSecurityLabelOk returns a tuple with the SecurityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLabel

`func (o *ApiNfsGaneshaExportPostRequest) SetSecurityLabel(v string)`

SetSecurityLabel sets SecurityLabel field to given value.


### GetSquash

`func (o *ApiNfsGaneshaExportPostRequest) GetSquash() string`

GetSquash returns the Squash field if non-nil, zero value otherwise.

### GetSquashOk

`func (o *ApiNfsGaneshaExportPostRequest) GetSquashOk() (*string, bool)`

GetSquashOk returns a tuple with the Squash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquash

`func (o *ApiNfsGaneshaExportPostRequest) SetSquash(v string)`

SetSquash sets Squash field to given value.


### GetTransports

`func (o *ApiNfsGaneshaExportPostRequest) GetTransports() []string`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *ApiNfsGaneshaExportPostRequest) GetTransportsOk() (*[]string, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *ApiNfsGaneshaExportPostRequest) SetTransports(v []string)`

SetTransports sets Transports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


