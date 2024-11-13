# ApiNfsGaneshaExportGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** | Export access type | [optional] 
**Clients** | Pointer to [**[]ApiNfsGaneshaExportGet200ResponseInnerClientsInner**](ApiNfsGaneshaExportGet200ResponseInnerClientsInner.md) | List of client configurations | [optional] 
**ClusterId** | Pointer to **string** | Cluster identifier | [optional] 
**ExportId** | Pointer to **int32** | Export ID | [optional] 
**Fsal** | Pointer to [**ApiNfsGaneshaExportGet200ResponseInnerFsal**](ApiNfsGaneshaExportGet200ResponseInnerFsal.md) |  | [optional] 
**Path** | Pointer to **string** | Export path | [optional] 
**Protocols** | Pointer to **[]int32** | List of protocol types | [optional] 
**Pseudo** | Pointer to **string** | Pseudo FS path | [optional] 
**SecurityLabel** | Pointer to **string** | Security label | [optional] 
**Squash** | Pointer to **string** | Export squash policy | [optional] 
**Transports** | Pointer to **[]string** | List of transport types | [optional] 

## Methods

### NewApiNfsGaneshaExportGet200ResponseInner

`func NewApiNfsGaneshaExportGet200ResponseInner() *ApiNfsGaneshaExportGet200ResponseInner`

NewApiNfsGaneshaExportGet200ResponseInner instantiates a new ApiNfsGaneshaExportGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNfsGaneshaExportGet200ResponseInnerWithDefaults

`func NewApiNfsGaneshaExportGet200ResponseInnerWithDefaults() *ApiNfsGaneshaExportGet200ResponseInner`

NewApiNfsGaneshaExportGet200ResponseInnerWithDefaults instantiates a new ApiNfsGaneshaExportGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetClients

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetClients() []ApiNfsGaneshaExportGet200ResponseInnerClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetClientsOk() (*[]ApiNfsGaneshaExportGet200ResponseInnerClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetClients(v []ApiNfsGaneshaExportGet200ResponseInnerClientsInner)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetClusterId

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetExportId

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetExportId() int32`

GetExportId returns the ExportId field if non-nil, zero value otherwise.

### GetExportIdOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetExportIdOk() (*int32, bool)`

GetExportIdOk returns a tuple with the ExportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportId

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetExportId(v int32)`

SetExportId sets ExportId field to given value.

### HasExportId

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasExportId() bool`

HasExportId returns a boolean if a field has been set.

### GetFsal

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetFsal() ApiNfsGaneshaExportGet200ResponseInnerFsal`

GetFsal returns the Fsal field if non-nil, zero value otherwise.

### GetFsalOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetFsalOk() (*ApiNfsGaneshaExportGet200ResponseInnerFsal, bool)`

GetFsalOk returns a tuple with the Fsal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsal

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetFsal(v ApiNfsGaneshaExportGet200ResponseInnerFsal)`

SetFsal sets Fsal field to given value.

### HasFsal

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasFsal() bool`

HasFsal returns a boolean if a field has been set.

### GetPath

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProtocols

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetProtocols() []int32`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetProtocolsOk() (*[]int32, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetProtocols(v []int32)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetPseudo

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetPseudo() string`

GetPseudo returns the Pseudo field if non-nil, zero value otherwise.

### GetPseudoOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetPseudoOk() (*string, bool)`

GetPseudoOk returns a tuple with the Pseudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudo

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetPseudo(v string)`

SetPseudo sets Pseudo field to given value.

### HasPseudo

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasPseudo() bool`

HasPseudo returns a boolean if a field has been set.

### GetSecurityLabel

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetSecurityLabel() string`

GetSecurityLabel returns the SecurityLabel field if non-nil, zero value otherwise.

### GetSecurityLabelOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetSecurityLabelOk() (*string, bool)`

GetSecurityLabelOk returns a tuple with the SecurityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLabel

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetSecurityLabel(v string)`

SetSecurityLabel sets SecurityLabel field to given value.

### HasSecurityLabel

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasSecurityLabel() bool`

HasSecurityLabel returns a boolean if a field has been set.

### GetSquash

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetSquash() string`

GetSquash returns the Squash field if non-nil, zero value otherwise.

### GetSquashOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetSquashOk() (*string, bool)`

GetSquashOk returns a tuple with the Squash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquash

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetSquash(v string)`

SetSquash sets Squash field to given value.

### HasSquash

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasSquash() bool`

HasSquash returns a boolean if a field has been set.

### GetTransports

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetTransports() []string`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *ApiNfsGaneshaExportGet200ResponseInner) GetTransportsOk() (*[]string, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *ApiNfsGaneshaExportGet200ResponseInner) SetTransports(v []string)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *ApiNfsGaneshaExportGet200ResponseInner) HasTransports() bool`

HasTransports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


