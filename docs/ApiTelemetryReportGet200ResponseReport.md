# ApiTelemetryReportGet200ResponseReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balancer** | [**ApiTelemetryReportGet200ResponseReportBalancer**](ApiTelemetryReportGet200ResponseReportBalancer.md) |  | 
**Channels** | **[]string** |  | 
**ChannelsAvailable** | **[]string** |  | 
**Config** | [**ApiTelemetryReportGet200ResponseReportConfig**](ApiTelemetryReportGet200ResponseReportConfig.md) |  | 
**Crashes** | **[]int32** |  | 
**Created** | **string** |  | 
**Crush** | [**ApiTelemetryReportGet200ResponseReportCrush**](ApiTelemetryReportGet200ResponseReportCrush.md) |  | 
**Fs** | [**ApiTelemetryReportGet200ResponseReportFs**](ApiTelemetryReportGet200ResponseReportFs.md) |  | 
**Hosts** | [**ApiTelemetryReportGet200ResponseReportHosts**](ApiTelemetryReportGet200ResponseReportHosts.md) |  | 
**Leaderboard** | **bool** |  | 
**License** | **string** |  | 
**Metadata** | [**ApiTelemetryReportGet200ResponseReportMetadata**](ApiTelemetryReportGet200ResponseReportMetadata.md) |  | 
**Mon** | [**ApiTelemetryReportGet200ResponseReportMon**](ApiTelemetryReportGet200ResponseReportMon.md) |  | 
**Osd** | [**ApiTelemetryReportGet200ResponseReportOsd**](ApiTelemetryReportGet200ResponseReportOsd.md) |  | 
**Pools** | [**[]ApiTelemetryReportGet200ResponseReportPoolsInner**](ApiTelemetryReportGet200ResponseReportPoolsInner.md) |  | 
**Rbd** | [**ApiTelemetryReportGet200ResponseReportRbd**](ApiTelemetryReportGet200ResponseReportRbd.md) |  | 
**ReportId** | **string** |  | 
**ReportTimestamp** | **string** |  | 
**ReportVersion** | **int32** |  | 
**Rgw** | [**ApiTelemetryReportGet200ResponseReportRgw**](ApiTelemetryReportGet200ResponseReportRgw.md) |  | 
**Services** | [**ApiTelemetryReportGet200ResponseReportServices**](ApiTelemetryReportGet200ResponseReportServices.md) |  | 
**Usage** | [**ApiTelemetryReportGet200ResponseReportUsage**](ApiTelemetryReportGet200ResponseReportUsage.md) |  | 

## Methods

### NewApiTelemetryReportGet200ResponseReport

`func NewApiTelemetryReportGet200ResponseReport(balancer ApiTelemetryReportGet200ResponseReportBalancer, channels []string, channelsAvailable []string, config ApiTelemetryReportGet200ResponseReportConfig, crashes []int32, created string, crush ApiTelemetryReportGet200ResponseReportCrush, fs ApiTelemetryReportGet200ResponseReportFs, hosts ApiTelemetryReportGet200ResponseReportHosts, leaderboard bool, license string, metadata ApiTelemetryReportGet200ResponseReportMetadata, mon ApiTelemetryReportGet200ResponseReportMon, osd ApiTelemetryReportGet200ResponseReportOsd, pools []ApiTelemetryReportGet200ResponseReportPoolsInner, rbd ApiTelemetryReportGet200ResponseReportRbd, reportId string, reportTimestamp string, reportVersion int32, rgw ApiTelemetryReportGet200ResponseReportRgw, services ApiTelemetryReportGet200ResponseReportServices, usage ApiTelemetryReportGet200ResponseReportUsage, ) *ApiTelemetryReportGet200ResponseReport`

NewApiTelemetryReportGet200ResponseReport instantiates a new ApiTelemetryReportGet200ResponseReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTelemetryReportGet200ResponseReportWithDefaults

`func NewApiTelemetryReportGet200ResponseReportWithDefaults() *ApiTelemetryReportGet200ResponseReport`

NewApiTelemetryReportGet200ResponseReportWithDefaults instantiates a new ApiTelemetryReportGet200ResponseReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalancer

`func (o *ApiTelemetryReportGet200ResponseReport) GetBalancer() ApiTelemetryReportGet200ResponseReportBalancer`

GetBalancer returns the Balancer field if non-nil, zero value otherwise.

### GetBalancerOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetBalancerOk() (*ApiTelemetryReportGet200ResponseReportBalancer, bool)`

GetBalancerOk returns a tuple with the Balancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancer

`func (o *ApiTelemetryReportGet200ResponseReport) SetBalancer(v ApiTelemetryReportGet200ResponseReportBalancer)`

SetBalancer sets Balancer field to given value.


### GetChannels

`func (o *ApiTelemetryReportGet200ResponseReport) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ApiTelemetryReportGet200ResponseReport) SetChannels(v []string)`

SetChannels sets Channels field to given value.


### GetChannelsAvailable

`func (o *ApiTelemetryReportGet200ResponseReport) GetChannelsAvailable() []string`

GetChannelsAvailable returns the ChannelsAvailable field if non-nil, zero value otherwise.

### GetChannelsAvailableOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetChannelsAvailableOk() (*[]string, bool)`

GetChannelsAvailableOk returns a tuple with the ChannelsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelsAvailable

`func (o *ApiTelemetryReportGet200ResponseReport) SetChannelsAvailable(v []string)`

SetChannelsAvailable sets ChannelsAvailable field to given value.


### GetConfig

`func (o *ApiTelemetryReportGet200ResponseReport) GetConfig() ApiTelemetryReportGet200ResponseReportConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetConfigOk() (*ApiTelemetryReportGet200ResponseReportConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiTelemetryReportGet200ResponseReport) SetConfig(v ApiTelemetryReportGet200ResponseReportConfig)`

SetConfig sets Config field to given value.


### GetCrashes

`func (o *ApiTelemetryReportGet200ResponseReport) GetCrashes() []int32`

GetCrashes returns the Crashes field if non-nil, zero value otherwise.

### GetCrashesOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetCrashesOk() (*[]int32, bool)`

GetCrashesOk returns a tuple with the Crashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrashes

`func (o *ApiTelemetryReportGet200ResponseReport) SetCrashes(v []int32)`

SetCrashes sets Crashes field to given value.


### GetCreated

`func (o *ApiTelemetryReportGet200ResponseReport) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiTelemetryReportGet200ResponseReport) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetCrush

`func (o *ApiTelemetryReportGet200ResponseReport) GetCrush() ApiTelemetryReportGet200ResponseReportCrush`

GetCrush returns the Crush field if non-nil, zero value otherwise.

### GetCrushOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetCrushOk() (*ApiTelemetryReportGet200ResponseReportCrush, bool)`

GetCrushOk returns a tuple with the Crush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrush

`func (o *ApiTelemetryReportGet200ResponseReport) SetCrush(v ApiTelemetryReportGet200ResponseReportCrush)`

SetCrush sets Crush field to given value.


### GetFs

`func (o *ApiTelemetryReportGet200ResponseReport) GetFs() ApiTelemetryReportGet200ResponseReportFs`

GetFs returns the Fs field if non-nil, zero value otherwise.

### GetFsOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetFsOk() (*ApiTelemetryReportGet200ResponseReportFs, bool)`

GetFsOk returns a tuple with the Fs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFs

`func (o *ApiTelemetryReportGet200ResponseReport) SetFs(v ApiTelemetryReportGet200ResponseReportFs)`

SetFs sets Fs field to given value.


### GetHosts

`func (o *ApiTelemetryReportGet200ResponseReport) GetHosts() ApiTelemetryReportGet200ResponseReportHosts`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetHostsOk() (*ApiTelemetryReportGet200ResponseReportHosts, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ApiTelemetryReportGet200ResponseReport) SetHosts(v ApiTelemetryReportGet200ResponseReportHosts)`

SetHosts sets Hosts field to given value.


### GetLeaderboard

`func (o *ApiTelemetryReportGet200ResponseReport) GetLeaderboard() bool`

GetLeaderboard returns the Leaderboard field if non-nil, zero value otherwise.

### GetLeaderboardOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetLeaderboardOk() (*bool, bool)`

GetLeaderboardOk returns a tuple with the Leaderboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboard

`func (o *ApiTelemetryReportGet200ResponseReport) SetLeaderboard(v bool)`

SetLeaderboard sets Leaderboard field to given value.


### GetLicense

`func (o *ApiTelemetryReportGet200ResponseReport) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ApiTelemetryReportGet200ResponseReport) SetLicense(v string)`

SetLicense sets License field to given value.


### GetMetadata

`func (o *ApiTelemetryReportGet200ResponseReport) GetMetadata() ApiTelemetryReportGet200ResponseReportMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetMetadataOk() (*ApiTelemetryReportGet200ResponseReportMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApiTelemetryReportGet200ResponseReport) SetMetadata(v ApiTelemetryReportGet200ResponseReportMetadata)`

SetMetadata sets Metadata field to given value.


### GetMon

`func (o *ApiTelemetryReportGet200ResponseReport) GetMon() ApiTelemetryReportGet200ResponseReportMon`

GetMon returns the Mon field if non-nil, zero value otherwise.

### GetMonOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetMonOk() (*ApiTelemetryReportGet200ResponseReportMon, bool)`

GetMonOk returns a tuple with the Mon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMon

`func (o *ApiTelemetryReportGet200ResponseReport) SetMon(v ApiTelemetryReportGet200ResponseReportMon)`

SetMon sets Mon field to given value.


### GetOsd

`func (o *ApiTelemetryReportGet200ResponseReport) GetOsd() ApiTelemetryReportGet200ResponseReportOsd`

GetOsd returns the Osd field if non-nil, zero value otherwise.

### GetOsdOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetOsdOk() (*ApiTelemetryReportGet200ResponseReportOsd, bool)`

GetOsdOk returns a tuple with the Osd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsd

`func (o *ApiTelemetryReportGet200ResponseReport) SetOsd(v ApiTelemetryReportGet200ResponseReportOsd)`

SetOsd sets Osd field to given value.


### GetPools

`func (o *ApiTelemetryReportGet200ResponseReport) GetPools() []ApiTelemetryReportGet200ResponseReportPoolsInner`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetPoolsOk() (*[]ApiTelemetryReportGet200ResponseReportPoolsInner, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *ApiTelemetryReportGet200ResponseReport) SetPools(v []ApiTelemetryReportGet200ResponseReportPoolsInner)`

SetPools sets Pools field to given value.


### GetRbd

`func (o *ApiTelemetryReportGet200ResponseReport) GetRbd() ApiTelemetryReportGet200ResponseReportRbd`

GetRbd returns the Rbd field if non-nil, zero value otherwise.

### GetRbdOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetRbdOk() (*ApiTelemetryReportGet200ResponseReportRbd, bool)`

GetRbdOk returns a tuple with the Rbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbd

`func (o *ApiTelemetryReportGet200ResponseReport) SetRbd(v ApiTelemetryReportGet200ResponseReportRbd)`

SetRbd sets Rbd field to given value.


### GetReportId

`func (o *ApiTelemetryReportGet200ResponseReport) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *ApiTelemetryReportGet200ResponseReport) SetReportId(v string)`

SetReportId sets ReportId field to given value.


### GetReportTimestamp

`func (o *ApiTelemetryReportGet200ResponseReport) GetReportTimestamp() string`

GetReportTimestamp returns the ReportTimestamp field if non-nil, zero value otherwise.

### GetReportTimestampOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetReportTimestampOk() (*string, bool)`

GetReportTimestampOk returns a tuple with the ReportTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTimestamp

`func (o *ApiTelemetryReportGet200ResponseReport) SetReportTimestamp(v string)`

SetReportTimestamp sets ReportTimestamp field to given value.


### GetReportVersion

`func (o *ApiTelemetryReportGet200ResponseReport) GetReportVersion() int32`

GetReportVersion returns the ReportVersion field if non-nil, zero value otherwise.

### GetReportVersionOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetReportVersionOk() (*int32, bool)`

GetReportVersionOk returns a tuple with the ReportVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportVersion

`func (o *ApiTelemetryReportGet200ResponseReport) SetReportVersion(v int32)`

SetReportVersion sets ReportVersion field to given value.


### GetRgw

`func (o *ApiTelemetryReportGet200ResponseReport) GetRgw() ApiTelemetryReportGet200ResponseReportRgw`

GetRgw returns the Rgw field if non-nil, zero value otherwise.

### GetRgwOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetRgwOk() (*ApiTelemetryReportGet200ResponseReportRgw, bool)`

GetRgwOk returns a tuple with the Rgw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRgw

`func (o *ApiTelemetryReportGet200ResponseReport) SetRgw(v ApiTelemetryReportGet200ResponseReportRgw)`

SetRgw sets Rgw field to given value.


### GetServices

`func (o *ApiTelemetryReportGet200ResponseReport) GetServices() ApiTelemetryReportGet200ResponseReportServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetServicesOk() (*ApiTelemetryReportGet200ResponseReportServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ApiTelemetryReportGet200ResponseReport) SetServices(v ApiTelemetryReportGet200ResponseReportServices)`

SetServices sets Services field to given value.


### GetUsage

`func (o *ApiTelemetryReportGet200ResponseReport) GetUsage() ApiTelemetryReportGet200ResponseReportUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ApiTelemetryReportGet200ResponseReport) GetUsageOk() (*ApiTelemetryReportGet200ResponseReportUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ApiTelemetryReportGet200ResponseReport) SetUsage(v ApiTelemetryReportGet200ResponseReportUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


