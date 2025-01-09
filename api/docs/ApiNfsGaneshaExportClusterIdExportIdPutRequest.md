# ApiNfsGaneshaExportClusterIdExportIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | **string** | Export access type | 
**Clients** | [**[]ApiNfsGaneshaExportGet200ResponseInnerClientsInner**](ApiNfsGaneshaExportGet200ResponseInnerClientsInner.md) | List of client configurations | 
**Fsal** | [**ApiNfsGaneshaExportPostRequestFsal**](ApiNfsGaneshaExportPostRequestFsal.md) |  | 
**Path** | **string** | Export path | 
**Protocols** | **[]int32** | List of protocol types | 
**Pseudo** | **string** | Pseudo FS path | 
**SecurityLabel** | **string** | Security label | 
**Squash** | **string** | Export squash policy | 
**Transports** | **[]string** | List of transport types | 

## Methods

### NewApiNfsGaneshaExportClusterIdExportIdPutRequest

`func NewApiNfsGaneshaExportClusterIdExportIdPutRequest(accessType string, clients []ApiNfsGaneshaExportGet200ResponseInnerClientsInner, fsal ApiNfsGaneshaExportPostRequestFsal, path string, protocols []int32, pseudo string, securityLabel string, squash string, transports []string, ) *ApiNfsGaneshaExportClusterIdExportIdPutRequest`

NewApiNfsGaneshaExportClusterIdExportIdPutRequest instantiates a new ApiNfsGaneshaExportClusterIdExportIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNfsGaneshaExportClusterIdExportIdPutRequestWithDefaults

`func NewApiNfsGaneshaExportClusterIdExportIdPutRequestWithDefaults() *ApiNfsGaneshaExportClusterIdExportIdPutRequest`

NewApiNfsGaneshaExportClusterIdExportIdPutRequestWithDefaults instantiates a new ApiNfsGaneshaExportClusterIdExportIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetClients

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetClients() []ApiNfsGaneshaExportGet200ResponseInnerClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetClientsOk() (*[]ApiNfsGaneshaExportGet200ResponseInnerClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) SetClients(v []ApiNfsGaneshaExportGet200ResponseInnerClientsInner)`

SetClients sets Clients field to given value.


### GetFsal

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetFsal() ApiNfsGaneshaExportPostRequestFsal`

GetFsal returns the Fsal field if non-nil, zero value otherwise.

### GetFsalOk

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetFsalOk() (*ApiNfsGaneshaExportPostRequestFsal, bool)`

GetFsalOk returns a tuple with the Fsal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsal

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) SetFsal(v ApiNfsGaneshaExportPostRequestFsal)`

SetFsal sets Fsal field to given value.


### GetPath

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetProtocols

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.


### GetPseudo

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetPseudo() string`

GetPseudo returns the Pseudo field if non-nil, zero value otherwise.

### GetPseudoOk

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetPseudoOk() (*string, bool)`

GetPseudoOk returns a tuple with the Pseudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudo

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) SetPseudo(v string)`

SetPseudo sets Pseudo field to given value.


### GetSecurityLabel

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetSecurityLabel() string`

GetSecurityLabel returns the SecurityLabel field if non-nil, zero value otherwise.

### GetSecurityLabelOk

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetSecurityLabelOk() (*string, bool)`

GetSecurityLabelOk returns a tuple with the SecurityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLabel

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) SetSecurityLabel(v string)`

SetSecurityLabel sets SecurityLabel field to given value.


### GetSquash

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetSquash() string`

GetSquash returns the Squash field if non-nil, zero value otherwise.

### GetSquashOk

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetSquashOk() (*string, bool)`

GetSquashOk returns a tuple with the Squash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquash

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) SetSquash(v string)`

SetSquash sets Squash field to given value.


### GetTransports

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetTransports() []string`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) GetTransportsOk() (*[]string, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *ApiNfsGaneshaExportClusterIdExportIdPutRequest) SetTransports(v []string)`

SetTransports sets Transports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


