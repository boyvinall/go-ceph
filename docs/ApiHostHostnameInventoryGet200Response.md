# ApiHostHostnameInventoryGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | **string** | Host address | 
**Devices** | [**[]ApiHostHostnameInventoryGet200ResponseDevicesInner**](ApiHostHostnameInventoryGet200ResponseDevicesInner.md) | Host devices | 
**Labels** | **[]string** | Host labels | 
**Name** | **string** | Hostname | 

## Methods

### NewApiHostHostnameInventoryGet200Response

`func NewApiHostHostnameInventoryGet200Response(addr string, devices []ApiHostHostnameInventoryGet200ResponseDevicesInner, labels []string, name string, ) *ApiHostHostnameInventoryGet200Response`

NewApiHostHostnameInventoryGet200Response instantiates a new ApiHostHostnameInventoryGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHostHostnameInventoryGet200ResponseWithDefaults

`func NewApiHostHostnameInventoryGet200ResponseWithDefaults() *ApiHostHostnameInventoryGet200Response`

NewApiHostHostnameInventoryGet200ResponseWithDefaults instantiates a new ApiHostHostnameInventoryGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *ApiHostHostnameInventoryGet200Response) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *ApiHostHostnameInventoryGet200Response) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *ApiHostHostnameInventoryGet200Response) SetAddr(v string)`

SetAddr sets Addr field to given value.


### GetDevices

`func (o *ApiHostHostnameInventoryGet200Response) GetDevices() []ApiHostHostnameInventoryGet200ResponseDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ApiHostHostnameInventoryGet200Response) GetDevicesOk() (*[]ApiHostHostnameInventoryGet200ResponseDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ApiHostHostnameInventoryGet200Response) SetDevices(v []ApiHostHostnameInventoryGet200ResponseDevicesInner)`

SetDevices sets Devices field to given value.


### GetLabels

`func (o *ApiHostHostnameInventoryGet200Response) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiHostHostnameInventoryGet200Response) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiHostHostnameInventoryGet200Response) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *ApiHostHostnameInventoryGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiHostHostnameInventoryGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiHostHostnameInventoryGet200Response) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


