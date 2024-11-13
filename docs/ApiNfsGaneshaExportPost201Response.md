# ApiNfsGaneshaExportPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | **string** | Export access type | 
**Clients** | [**[]ApiNfsGaneshaExportGet200ResponseInnerClientsInner**](ApiNfsGaneshaExportGet200ResponseInnerClientsInner.md) | List of client configurations | 
**ClusterId** | **string** | Cluster identifier | 
**ExportId** | **int32** | Export ID | 
**Fsal** | [**ApiNfsGaneshaExportGet200ResponseInnerFsal**](ApiNfsGaneshaExportGet200ResponseInnerFsal.md) |  | 
**Path** | **string** | Export path | 
**Protocols** | **[]int32** | List of protocol types | 
**Pseudo** | **string** | Pseudo FS path | 
**SecurityLabel** | **string** | Security label | 
**Squash** | **string** | Export squash policy | 
**Transports** | **[]string** | List of transport types | 

## Methods

### NewApiNfsGaneshaExportPost201Response

`func NewApiNfsGaneshaExportPost201Response(accessType string, clients []ApiNfsGaneshaExportGet200ResponseInnerClientsInner, clusterId string, exportId int32, fsal ApiNfsGaneshaExportGet200ResponseInnerFsal, path string, protocols []int32, pseudo string, securityLabel string, squash string, transports []string, ) *ApiNfsGaneshaExportPost201Response`

NewApiNfsGaneshaExportPost201Response instantiates a new ApiNfsGaneshaExportPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNfsGaneshaExportPost201ResponseWithDefaults

`func NewApiNfsGaneshaExportPost201ResponseWithDefaults() *ApiNfsGaneshaExportPost201Response`

NewApiNfsGaneshaExportPost201ResponseWithDefaults instantiates a new ApiNfsGaneshaExportPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *ApiNfsGaneshaExportPost201Response) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *ApiNfsGaneshaExportPost201Response) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *ApiNfsGaneshaExportPost201Response) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetClients

`func (o *ApiNfsGaneshaExportPost201Response) GetClients() []ApiNfsGaneshaExportGet200ResponseInnerClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ApiNfsGaneshaExportPost201Response) GetClientsOk() (*[]ApiNfsGaneshaExportGet200ResponseInnerClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ApiNfsGaneshaExportPost201Response) SetClients(v []ApiNfsGaneshaExportGet200ResponseInnerClientsInner)`

SetClients sets Clients field to given value.


### GetClusterId

`func (o *ApiNfsGaneshaExportPost201Response) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ApiNfsGaneshaExportPost201Response) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ApiNfsGaneshaExportPost201Response) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetExportId

`func (o *ApiNfsGaneshaExportPost201Response) GetExportId() int32`

GetExportId returns the ExportId field if non-nil, zero value otherwise.

### GetExportIdOk

`func (o *ApiNfsGaneshaExportPost201Response) GetExportIdOk() (*int32, bool)`

GetExportIdOk returns a tuple with the ExportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportId

`func (o *ApiNfsGaneshaExportPost201Response) SetExportId(v int32)`

SetExportId sets ExportId field to given value.


### GetFsal

`func (o *ApiNfsGaneshaExportPost201Response) GetFsal() ApiNfsGaneshaExportGet200ResponseInnerFsal`

GetFsal returns the Fsal field if non-nil, zero value otherwise.

### GetFsalOk

`func (o *ApiNfsGaneshaExportPost201Response) GetFsalOk() (*ApiNfsGaneshaExportGet200ResponseInnerFsal, bool)`

GetFsalOk returns a tuple with the Fsal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsal

`func (o *ApiNfsGaneshaExportPost201Response) SetFsal(v ApiNfsGaneshaExportGet200ResponseInnerFsal)`

SetFsal sets Fsal field to given value.


### GetPath

`func (o *ApiNfsGaneshaExportPost201Response) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiNfsGaneshaExportPost201Response) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiNfsGaneshaExportPost201Response) SetPath(v string)`

SetPath sets Path field to given value.


### GetProtocols

`func (o *ApiNfsGaneshaExportPost201Response) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ApiNfsGaneshaExportPost201Response) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ApiNfsGaneshaExportPost201Response) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetPseudo

`func (o *ApiNfsGaneshaExportPost201Response) GetPseudo() string`

GetPseudo returns the Pseudo field if non-nil, zero value otherwise.

### GetPseudoOk

`func (o *ApiNfsGaneshaExportPost201Response) GetPseudoOk() (*string, bool)`

GetPseudoOk returns a tuple with the Pseudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudo

`func (o *ApiNfsGaneshaExportPost201Response) SetPseudo(v string)`

SetPseudo sets Pseudo field to given value.


### GetSecurityLabel

`func (o *ApiNfsGaneshaExportPost201Response) GetSecurityLabel() string`

GetSecurityLabel returns the SecurityLabel field if non-nil, zero value otherwise.

### GetSecurityLabelOk

`func (o *ApiNfsGaneshaExportPost201Response) GetSecurityLabelOk() (*string, bool)`

GetSecurityLabelOk returns a tuple with the SecurityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLabel

`func (o *ApiNfsGaneshaExportPost201Response) SetSecurityLabel(v string)`

SetSecurityLabel sets SecurityLabel field to given value.


### GetSquash

`func (o *ApiNfsGaneshaExportPost201Response) GetSquash() string`

GetSquash returns the Squash field if non-nil, zero value otherwise.

### GetSquashOk

`func (o *ApiNfsGaneshaExportPost201Response) GetSquashOk() (*string, bool)`

GetSquashOk returns a tuple with the Squash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquash

`func (o *ApiNfsGaneshaExportPost201Response) SetSquash(v string)`

SetSquash sets Squash field to given value.


### GetTransports

`func (o *ApiNfsGaneshaExportPost201Response) GetTransports() []string`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *ApiNfsGaneshaExportPost201Response) GetTransportsOk() (*[]string, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *ApiNfsGaneshaExportPost201Response) SetTransports(v []string)`

SetTransports sets Transports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


