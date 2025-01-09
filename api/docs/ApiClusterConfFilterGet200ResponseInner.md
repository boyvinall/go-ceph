# ApiClusterConfFilterGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanUpdateAtRuntime** | Pointer to **bool** | Check if can update at runtime | [optional] 
**DaemonDefault** | Pointer to **string** | Daemon specific default value | [optional] 
**Default** | Pointer to **string** | Default value for the config option | [optional] 
**Desc** | Pointer to **string** | Description of the configuration | [optional] 
**EnumValues** | Pointer to **[]string** | List of enums allowed | [optional] 
**Flags** | Pointer to **[]string** | List of flags associated | [optional] 
**Level** | Pointer to **string** | Config option level | [optional] 
**LongDesc** | Pointer to **string** | Elaborated description | [optional] 
**Max** | Pointer to **string** | Maximum value | [optional] 
**Min** | Pointer to **string** | Minimum value | [optional] 
**Name** | Pointer to **string** | Name of the config option | [optional] 
**SeeAlso** | Pointer to **[]string** | Related config options | [optional] 
**Services** | Pointer to **[]string** | Services associated with the config option | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the cluster | [optional] 
**Type** | Pointer to **string** | Config option type | [optional] 

## Methods

### NewApiClusterConfFilterGet200ResponseInner

`func NewApiClusterConfFilterGet200ResponseInner() *ApiClusterConfFilterGet200ResponseInner`

NewApiClusterConfFilterGet200ResponseInner instantiates a new ApiClusterConfFilterGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClusterConfFilterGet200ResponseInnerWithDefaults

`func NewApiClusterConfFilterGet200ResponseInnerWithDefaults() *ApiClusterConfFilterGet200ResponseInner`

NewApiClusterConfFilterGet200ResponseInnerWithDefaults instantiates a new ApiClusterConfFilterGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanUpdateAtRuntime

`func (o *ApiClusterConfFilterGet200ResponseInner) GetCanUpdateAtRuntime() bool`

GetCanUpdateAtRuntime returns the CanUpdateAtRuntime field if non-nil, zero value otherwise.

### GetCanUpdateAtRuntimeOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetCanUpdateAtRuntimeOk() (*bool, bool)`

GetCanUpdateAtRuntimeOk returns a tuple with the CanUpdateAtRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpdateAtRuntime

`func (o *ApiClusterConfFilterGet200ResponseInner) SetCanUpdateAtRuntime(v bool)`

SetCanUpdateAtRuntime sets CanUpdateAtRuntime field to given value.

### HasCanUpdateAtRuntime

`func (o *ApiClusterConfFilterGet200ResponseInner) HasCanUpdateAtRuntime() bool`

HasCanUpdateAtRuntime returns a boolean if a field has been set.

### GetDaemonDefault

`func (o *ApiClusterConfFilterGet200ResponseInner) GetDaemonDefault() string`

GetDaemonDefault returns the DaemonDefault field if non-nil, zero value otherwise.

### GetDaemonDefaultOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetDaemonDefaultOk() (*string, bool)`

GetDaemonDefaultOk returns a tuple with the DaemonDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonDefault

`func (o *ApiClusterConfFilterGet200ResponseInner) SetDaemonDefault(v string)`

SetDaemonDefault sets DaemonDefault field to given value.

### HasDaemonDefault

`func (o *ApiClusterConfFilterGet200ResponseInner) HasDaemonDefault() bool`

HasDaemonDefault returns a boolean if a field has been set.

### GetDefault

`func (o *ApiClusterConfFilterGet200ResponseInner) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ApiClusterConfFilterGet200ResponseInner) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ApiClusterConfFilterGet200ResponseInner) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDesc

`func (o *ApiClusterConfFilterGet200ResponseInner) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ApiClusterConfFilterGet200ResponseInner) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *ApiClusterConfFilterGet200ResponseInner) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetEnumValues

`func (o *ApiClusterConfFilterGet200ResponseInner) GetEnumValues() []string`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetEnumValuesOk() (*[]string, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *ApiClusterConfFilterGet200ResponseInner) SetEnumValues(v []string)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *ApiClusterConfFilterGet200ResponseInner) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### GetFlags

`func (o *ApiClusterConfFilterGet200ResponseInner) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ApiClusterConfFilterGet200ResponseInner) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ApiClusterConfFilterGet200ResponseInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetLevel

`func (o *ApiClusterConfFilterGet200ResponseInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ApiClusterConfFilterGet200ResponseInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ApiClusterConfFilterGet200ResponseInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLongDesc

`func (o *ApiClusterConfFilterGet200ResponseInner) GetLongDesc() string`

GetLongDesc returns the LongDesc field if non-nil, zero value otherwise.

### GetLongDescOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetLongDescOk() (*string, bool)`

GetLongDescOk returns a tuple with the LongDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongDesc

`func (o *ApiClusterConfFilterGet200ResponseInner) SetLongDesc(v string)`

SetLongDesc sets LongDesc field to given value.

### HasLongDesc

`func (o *ApiClusterConfFilterGet200ResponseInner) HasLongDesc() bool`

HasLongDesc returns a boolean if a field has been set.

### GetMax

`func (o *ApiClusterConfFilterGet200ResponseInner) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ApiClusterConfFilterGet200ResponseInner) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *ApiClusterConfFilterGet200ResponseInner) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *ApiClusterConfFilterGet200ResponseInner) GetMin() string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetMinOk() (*string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ApiClusterConfFilterGet200ResponseInner) SetMin(v string)`

SetMin sets Min field to given value.

### HasMin

`func (o *ApiClusterConfFilterGet200ResponseInner) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetName

`func (o *ApiClusterConfFilterGet200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiClusterConfFilterGet200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiClusterConfFilterGet200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeeAlso

`func (o *ApiClusterConfFilterGet200ResponseInner) GetSeeAlso() []string`

GetSeeAlso returns the SeeAlso field if non-nil, zero value otherwise.

### GetSeeAlsoOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetSeeAlsoOk() (*[]string, bool)`

GetSeeAlsoOk returns a tuple with the SeeAlso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeAlso

`func (o *ApiClusterConfFilterGet200ResponseInner) SetSeeAlso(v []string)`

SetSeeAlso sets SeeAlso field to given value.

### HasSeeAlso

`func (o *ApiClusterConfFilterGet200ResponseInner) HasSeeAlso() bool`

HasSeeAlso returns a boolean if a field has been set.

### GetServices

`func (o *ApiClusterConfFilterGet200ResponseInner) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ApiClusterConfFilterGet200ResponseInner) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *ApiClusterConfFilterGet200ResponseInner) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTags

`func (o *ApiClusterConfFilterGet200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApiClusterConfFilterGet200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApiClusterConfFilterGet200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *ApiClusterConfFilterGet200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiClusterConfFilterGet200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiClusterConfFilterGet200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiClusterConfFilterGet200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


