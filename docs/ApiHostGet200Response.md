# ApiHostGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | **string** | Host address | 
**CephVersion** | **string** | Ceph version | 
**Hostname** | **string** | Hostname | 
**Labels** | **[]string** | Labels related to the host | 
**ServiceInstances** | [**[]ApiHostGet200ResponseServiceInstancesInner**](ApiHostGet200ResponseServiceInstancesInner.md) | Service instances related to the host | 
**ServiceType** | **string** |  | 
**Services** | [**[]ApiHostGet200ResponseServicesInner**](ApiHostGet200ResponseServicesInner.md) | Services related to the host | 
**Sources** | [**ApiHostGet200ResponseSources**](ApiHostGet200ResponseSources.md) |  | 
**Status** | **string** |  | 

## Methods

### NewApiHostGet200Response

`func NewApiHostGet200Response(addr string, cephVersion string, hostname string, labels []string, serviceInstances []ApiHostGet200ResponseServiceInstancesInner, serviceType string, services []ApiHostGet200ResponseServicesInner, sources ApiHostGet200ResponseSources, status string, ) *ApiHostGet200Response`

NewApiHostGet200Response instantiates a new ApiHostGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHostGet200ResponseWithDefaults

`func NewApiHostGet200ResponseWithDefaults() *ApiHostGet200Response`

NewApiHostGet200ResponseWithDefaults instantiates a new ApiHostGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *ApiHostGet200Response) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *ApiHostGet200Response) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *ApiHostGet200Response) SetAddr(v string)`

SetAddr sets Addr field to given value.


### GetCephVersion

`func (o *ApiHostGet200Response) GetCephVersion() string`

GetCephVersion returns the CephVersion field if non-nil, zero value otherwise.

### GetCephVersionOk

`func (o *ApiHostGet200Response) GetCephVersionOk() (*string, bool)`

GetCephVersionOk returns a tuple with the CephVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephVersion

`func (o *ApiHostGet200Response) SetCephVersion(v string)`

SetCephVersion sets CephVersion field to given value.


### GetHostname

`func (o *ApiHostGet200Response) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApiHostGet200Response) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApiHostGet200Response) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetLabels

`func (o *ApiHostGet200Response) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiHostGet200Response) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiHostGet200Response) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetServiceInstances

`func (o *ApiHostGet200Response) GetServiceInstances() []ApiHostGet200ResponseServiceInstancesInner`

GetServiceInstances returns the ServiceInstances field if non-nil, zero value otherwise.

### GetServiceInstancesOk

`func (o *ApiHostGet200Response) GetServiceInstancesOk() (*[]ApiHostGet200ResponseServiceInstancesInner, bool)`

GetServiceInstancesOk returns a tuple with the ServiceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstances

`func (o *ApiHostGet200Response) SetServiceInstances(v []ApiHostGet200ResponseServiceInstancesInner)`

SetServiceInstances sets ServiceInstances field to given value.


### GetServiceType

`func (o *ApiHostGet200Response) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ApiHostGet200Response) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ApiHostGet200Response) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetServices

`func (o *ApiHostGet200Response) GetServices() []ApiHostGet200ResponseServicesInner`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ApiHostGet200Response) GetServicesOk() (*[]ApiHostGet200ResponseServicesInner, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ApiHostGet200Response) SetServices(v []ApiHostGet200ResponseServicesInner)`

SetServices sets Services field to given value.


### GetSources

`func (o *ApiHostGet200Response) GetSources() ApiHostGet200ResponseSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ApiHostGet200Response) GetSourcesOk() (*ApiHostGet200ResponseSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ApiHostGet200Response) SetSources(v ApiHostGet200ResponseSources)`

SetSources sets Sources field to given value.


### GetStatus

`func (o *ApiHostGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiHostGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiHostGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


