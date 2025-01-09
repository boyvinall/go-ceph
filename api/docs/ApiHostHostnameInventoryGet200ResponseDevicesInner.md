# ApiHostHostnameInventoryGet200ResponseDevicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | If the device can be provisioned to an OSD | 
**DeviceId** | **string** | Device&#39;s udev ID | 
**HumanReadableType** | **string** | Device type. ssd or hdd | 
**LsmData** | [**ApiHostHostnameInventoryGet200ResponseDevicesInnerLsmData**](ApiHostHostnameInventoryGet200ResponseDevicesInnerLsmData.md) |  | 
**Lvs** | [**[]ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner**](ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner.md) |  | 
**OsdIds** | **[]int32** | Device OSD IDs | 
**Path** | **string** | Device path | 
**RejectedReasons** | **[]string** |  | 
**SysApi** | [**ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi**](ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi.md) |  | 

## Methods

### NewApiHostHostnameInventoryGet200ResponseDevicesInner

`func NewApiHostHostnameInventoryGet200ResponseDevicesInner(available bool, deviceId string, humanReadableType string, lsmData ApiHostHostnameInventoryGet200ResponseDevicesInnerLsmData, lvs []ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner, osdIds []int32, path string, rejectedReasons []string, sysApi ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi, ) *ApiHostHostnameInventoryGet200ResponseDevicesInner`

NewApiHostHostnameInventoryGet200ResponseDevicesInner instantiates a new ApiHostHostnameInventoryGet200ResponseDevicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHostHostnameInventoryGet200ResponseDevicesInnerWithDefaults

`func NewApiHostHostnameInventoryGet200ResponseDevicesInnerWithDefaults() *ApiHostHostnameInventoryGet200ResponseDevicesInner`

NewApiHostHostnameInventoryGet200ResponseDevicesInnerWithDefaults instantiates a new ApiHostHostnameInventoryGet200ResponseDevicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetDeviceId

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetHumanReadableType

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetHumanReadableType() string`

GetHumanReadableType returns the HumanReadableType field if non-nil, zero value otherwise.

### GetHumanReadableTypeOk

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetHumanReadableTypeOk() (*string, bool)`

GetHumanReadableTypeOk returns a tuple with the HumanReadableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanReadableType

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) SetHumanReadableType(v string)`

SetHumanReadableType sets HumanReadableType field to given value.


### GetLsmData

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetLsmData() ApiHostHostnameInventoryGet200ResponseDevicesInnerLsmData`

GetLsmData returns the LsmData field if non-nil, zero value otherwise.

### GetLsmDataOk

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetLsmDataOk() (*ApiHostHostnameInventoryGet200ResponseDevicesInnerLsmData, bool)`

GetLsmDataOk returns a tuple with the LsmData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLsmData

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) SetLsmData(v ApiHostHostnameInventoryGet200ResponseDevicesInnerLsmData)`

SetLsmData sets LsmData field to given value.


### GetLvs

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetLvs() []ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner`

GetLvs returns the Lvs field if non-nil, zero value otherwise.

### GetLvsOk

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetLvsOk() (*[]ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner, bool)`

GetLvsOk returns a tuple with the Lvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvs

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) SetLvs(v []ApiHostHostnameInventoryGet200ResponseDevicesInnerLvsInner)`

SetLvs sets Lvs field to given value.


### GetOsdIds

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetOsdIds() []int32`

GetOsdIds returns the OsdIds field if non-nil, zero value otherwise.

### GetOsdIdsOk

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetOsdIdsOk() (*[]int32, bool)`

GetOsdIdsOk returns a tuple with the OsdIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdIds

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) SetOsdIds(v []int32)`

SetOsdIds sets OsdIds field to given value.


### GetPath

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetRejectedReasons

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetRejectedReasons() []string`

GetRejectedReasons returns the RejectedReasons field if non-nil, zero value otherwise.

### GetRejectedReasonsOk

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetRejectedReasonsOk() (*[]string, bool)`

GetRejectedReasonsOk returns a tuple with the RejectedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedReasons

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) SetRejectedReasons(v []string)`

SetRejectedReasons sets RejectedReasons field to given value.


### GetSysApi

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetSysApi() ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi`

GetSysApi returns the SysApi field if non-nil, zero value otherwise.

### GetSysApiOk

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) GetSysApiOk() (*ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi, bool)`

GetSysApiOk returns a tuple with the SysApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysApi

`func (o *ApiHostHostnameInventoryGet200ResponseDevicesInner) SetSysApi(v ApiHostHostnameInventoryGet200ResponseDevicesInnerSysApi)`

SetSysApi sets SysApi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


