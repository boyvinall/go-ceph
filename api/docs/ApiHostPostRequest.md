# ApiHostPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | Pointer to **string** | Network Address | [optional] 
**Hostname** | **string** | Hostname | 
**Labels** | Pointer to **[]string** | Host Labels | [optional] 
**Status** | Pointer to **string** | Host Status | [optional] 

## Methods

### NewApiHostPostRequest

`func NewApiHostPostRequest(hostname string, ) *ApiHostPostRequest`

NewApiHostPostRequest instantiates a new ApiHostPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHostPostRequestWithDefaults

`func NewApiHostPostRequestWithDefaults() *ApiHostPostRequest`

NewApiHostPostRequestWithDefaults instantiates a new ApiHostPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *ApiHostPostRequest) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *ApiHostPostRequest) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *ApiHostPostRequest) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *ApiHostPostRequest) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetHostname

`func (o *ApiHostPostRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApiHostPostRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApiHostPostRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetLabels

`func (o *ApiHostPostRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiHostPostRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiHostPostRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiHostPostRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetStatus

`func (o *ApiHostPostRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiHostPostRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiHostPostRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiHostPostRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


