# ApiHealthMinimalGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientPerf** | [**ApiHealthMinimalGet200ResponseClientPerf**](ApiHealthMinimalGet200ResponseClientPerf.md) |  | 
**Df** | [**ApiHealthMinimalGet200ResponseDf**](ApiHealthMinimalGet200ResponseDf.md) |  | 
**FsMap** | [**ApiHealthMinimalGet200ResponseFsMap**](ApiHealthMinimalGet200ResponseFsMap.md) |  | 
**Health** | [**ApiHealthMinimalGet200ResponseHealth**](ApiHealthMinimalGet200ResponseHealth.md) |  | 
**Hosts** | **int32** |  | 
**IscsiDaemons** | [**ApiHealthMinimalGet200ResponseIscsiDaemons**](ApiHealthMinimalGet200ResponseIscsiDaemons.md) |  | 
**MgrMap** | [**ApiHealthMinimalGet200ResponseMgrMap**](ApiHealthMinimalGet200ResponseMgrMap.md) |  | 
**MonStatus** | [**ApiHealthMinimalGet200ResponseMonStatus**](ApiHealthMinimalGet200ResponseMonStatus.md) |  | 
**OsdMap** | [**ApiHealthMinimalGet200ResponseOsdMap**](ApiHealthMinimalGet200ResponseOsdMap.md) |  | 
**PgInfo** | [**ApiHealthMinimalGet200ResponsePgInfo**](ApiHealthMinimalGet200ResponsePgInfo.md) |  | 
**Pools** | **string** |  | 
**Rgw** | **int32** |  | 
**ScrubStatus** | **string** |  | 

## Methods

### NewApiHealthMinimalGet200Response

`func NewApiHealthMinimalGet200Response(clientPerf ApiHealthMinimalGet200ResponseClientPerf, df ApiHealthMinimalGet200ResponseDf, fsMap ApiHealthMinimalGet200ResponseFsMap, health ApiHealthMinimalGet200ResponseHealth, hosts int32, iscsiDaemons ApiHealthMinimalGet200ResponseIscsiDaemons, mgrMap ApiHealthMinimalGet200ResponseMgrMap, monStatus ApiHealthMinimalGet200ResponseMonStatus, osdMap ApiHealthMinimalGet200ResponseOsdMap, pgInfo ApiHealthMinimalGet200ResponsePgInfo, pools string, rgw int32, scrubStatus string, ) *ApiHealthMinimalGet200Response`

NewApiHealthMinimalGet200Response instantiates a new ApiHealthMinimalGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHealthMinimalGet200ResponseWithDefaults

`func NewApiHealthMinimalGet200ResponseWithDefaults() *ApiHealthMinimalGet200Response`

NewApiHealthMinimalGet200ResponseWithDefaults instantiates a new ApiHealthMinimalGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientPerf

`func (o *ApiHealthMinimalGet200Response) GetClientPerf() ApiHealthMinimalGet200ResponseClientPerf`

GetClientPerf returns the ClientPerf field if non-nil, zero value otherwise.

### GetClientPerfOk

`func (o *ApiHealthMinimalGet200Response) GetClientPerfOk() (*ApiHealthMinimalGet200ResponseClientPerf, bool)`

GetClientPerfOk returns a tuple with the ClientPerf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPerf

`func (o *ApiHealthMinimalGet200Response) SetClientPerf(v ApiHealthMinimalGet200ResponseClientPerf)`

SetClientPerf sets ClientPerf field to given value.


### GetDf

`func (o *ApiHealthMinimalGet200Response) GetDf() ApiHealthMinimalGet200ResponseDf`

GetDf returns the Df field if non-nil, zero value otherwise.

### GetDfOk

`func (o *ApiHealthMinimalGet200Response) GetDfOk() (*ApiHealthMinimalGet200ResponseDf, bool)`

GetDfOk returns a tuple with the Df field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDf

`func (o *ApiHealthMinimalGet200Response) SetDf(v ApiHealthMinimalGet200ResponseDf)`

SetDf sets Df field to given value.


### GetFsMap

`func (o *ApiHealthMinimalGet200Response) GetFsMap() ApiHealthMinimalGet200ResponseFsMap`

GetFsMap returns the FsMap field if non-nil, zero value otherwise.

### GetFsMapOk

`func (o *ApiHealthMinimalGet200Response) GetFsMapOk() (*ApiHealthMinimalGet200ResponseFsMap, bool)`

GetFsMapOk returns a tuple with the FsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsMap

`func (o *ApiHealthMinimalGet200Response) SetFsMap(v ApiHealthMinimalGet200ResponseFsMap)`

SetFsMap sets FsMap field to given value.


### GetHealth

`func (o *ApiHealthMinimalGet200Response) GetHealth() ApiHealthMinimalGet200ResponseHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ApiHealthMinimalGet200Response) GetHealthOk() (*ApiHealthMinimalGet200ResponseHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ApiHealthMinimalGet200Response) SetHealth(v ApiHealthMinimalGet200ResponseHealth)`

SetHealth sets Health field to given value.


### GetHosts

`func (o *ApiHealthMinimalGet200Response) GetHosts() int32`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ApiHealthMinimalGet200Response) GetHostsOk() (*int32, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ApiHealthMinimalGet200Response) SetHosts(v int32)`

SetHosts sets Hosts field to given value.


### GetIscsiDaemons

`func (o *ApiHealthMinimalGet200Response) GetIscsiDaemons() ApiHealthMinimalGet200ResponseIscsiDaemons`

GetIscsiDaemons returns the IscsiDaemons field if non-nil, zero value otherwise.

### GetIscsiDaemonsOk

`func (o *ApiHealthMinimalGet200Response) GetIscsiDaemonsOk() (*ApiHealthMinimalGet200ResponseIscsiDaemons, bool)`

GetIscsiDaemonsOk returns a tuple with the IscsiDaemons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiDaemons

`func (o *ApiHealthMinimalGet200Response) SetIscsiDaemons(v ApiHealthMinimalGet200ResponseIscsiDaemons)`

SetIscsiDaemons sets IscsiDaemons field to given value.


### GetMgrMap

`func (o *ApiHealthMinimalGet200Response) GetMgrMap() ApiHealthMinimalGet200ResponseMgrMap`

GetMgrMap returns the MgrMap field if non-nil, zero value otherwise.

### GetMgrMapOk

`func (o *ApiHealthMinimalGet200Response) GetMgrMapOk() (*ApiHealthMinimalGet200ResponseMgrMap, bool)`

GetMgrMapOk returns a tuple with the MgrMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgrMap

`func (o *ApiHealthMinimalGet200Response) SetMgrMap(v ApiHealthMinimalGet200ResponseMgrMap)`

SetMgrMap sets MgrMap field to given value.


### GetMonStatus

`func (o *ApiHealthMinimalGet200Response) GetMonStatus() ApiHealthMinimalGet200ResponseMonStatus`

GetMonStatus returns the MonStatus field if non-nil, zero value otherwise.

### GetMonStatusOk

`func (o *ApiHealthMinimalGet200Response) GetMonStatusOk() (*ApiHealthMinimalGet200ResponseMonStatus, bool)`

GetMonStatusOk returns a tuple with the MonStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonStatus

`func (o *ApiHealthMinimalGet200Response) SetMonStatus(v ApiHealthMinimalGet200ResponseMonStatus)`

SetMonStatus sets MonStatus field to given value.


### GetOsdMap

`func (o *ApiHealthMinimalGet200Response) GetOsdMap() ApiHealthMinimalGet200ResponseOsdMap`

GetOsdMap returns the OsdMap field if non-nil, zero value otherwise.

### GetOsdMapOk

`func (o *ApiHealthMinimalGet200Response) GetOsdMapOk() (*ApiHealthMinimalGet200ResponseOsdMap, bool)`

GetOsdMapOk returns a tuple with the OsdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdMap

`func (o *ApiHealthMinimalGet200Response) SetOsdMap(v ApiHealthMinimalGet200ResponseOsdMap)`

SetOsdMap sets OsdMap field to given value.


### GetPgInfo

`func (o *ApiHealthMinimalGet200Response) GetPgInfo() ApiHealthMinimalGet200ResponsePgInfo`

GetPgInfo returns the PgInfo field if non-nil, zero value otherwise.

### GetPgInfoOk

`func (o *ApiHealthMinimalGet200Response) GetPgInfoOk() (*ApiHealthMinimalGet200ResponsePgInfo, bool)`

GetPgInfoOk returns a tuple with the PgInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgInfo

`func (o *ApiHealthMinimalGet200Response) SetPgInfo(v ApiHealthMinimalGet200ResponsePgInfo)`

SetPgInfo sets PgInfo field to given value.


### GetPools

`func (o *ApiHealthMinimalGet200Response) GetPools() string`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *ApiHealthMinimalGet200Response) GetPoolsOk() (*string, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *ApiHealthMinimalGet200Response) SetPools(v string)`

SetPools sets Pools field to given value.


### GetRgw

`func (o *ApiHealthMinimalGet200Response) GetRgw() int32`

GetRgw returns the Rgw field if non-nil, zero value otherwise.

### GetRgwOk

`func (o *ApiHealthMinimalGet200Response) GetRgwOk() (*int32, bool)`

GetRgwOk returns a tuple with the Rgw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRgw

`func (o *ApiHealthMinimalGet200Response) SetRgw(v int32)`

SetRgw sets Rgw field to given value.


### GetScrubStatus

`func (o *ApiHealthMinimalGet200Response) GetScrubStatus() string`

GetScrubStatus returns the ScrubStatus field if non-nil, zero value otherwise.

### GetScrubStatusOk

`func (o *ApiHealthMinimalGet200Response) GetScrubStatusOk() (*string, bool)`

GetScrubStatusOk returns a tuple with the ScrubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubStatus

`func (o *ApiHealthMinimalGet200Response) SetScrubStatus(v string)`

SetScrubStatus sets ScrubStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


