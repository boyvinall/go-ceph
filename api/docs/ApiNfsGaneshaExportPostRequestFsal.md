# ApiNfsGaneshaExportPostRequestFsal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsName** | Pointer to **string** | CephFS filesystem name | [optional] 
**Name** | **string** | name of FSAL | 
**SecLabelXattr** | Pointer to **string** | Name of xattr for security label | [optional] 

## Methods

### NewApiNfsGaneshaExportPostRequestFsal

`func NewApiNfsGaneshaExportPostRequestFsal(name string, ) *ApiNfsGaneshaExportPostRequestFsal`

NewApiNfsGaneshaExportPostRequestFsal instantiates a new ApiNfsGaneshaExportPostRequestFsal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNfsGaneshaExportPostRequestFsalWithDefaults

`func NewApiNfsGaneshaExportPostRequestFsalWithDefaults() *ApiNfsGaneshaExportPostRequestFsal`

NewApiNfsGaneshaExportPostRequestFsalWithDefaults instantiates a new ApiNfsGaneshaExportPostRequestFsal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsName

`func (o *ApiNfsGaneshaExportPostRequestFsal) GetFsName() string`

GetFsName returns the FsName field if non-nil, zero value otherwise.

### GetFsNameOk

`func (o *ApiNfsGaneshaExportPostRequestFsal) GetFsNameOk() (*string, bool)`

GetFsNameOk returns a tuple with the FsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsName

`func (o *ApiNfsGaneshaExportPostRequestFsal) SetFsName(v string)`

SetFsName sets FsName field to given value.

### HasFsName

`func (o *ApiNfsGaneshaExportPostRequestFsal) HasFsName() bool`

HasFsName returns a boolean if a field has been set.

### GetName

`func (o *ApiNfsGaneshaExportPostRequestFsal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiNfsGaneshaExportPostRequestFsal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiNfsGaneshaExportPostRequestFsal) SetName(v string)`

SetName sets Name field to given value.


### GetSecLabelXattr

`func (o *ApiNfsGaneshaExportPostRequestFsal) GetSecLabelXattr() string`

GetSecLabelXattr returns the SecLabelXattr field if non-nil, zero value otherwise.

### GetSecLabelXattrOk

`func (o *ApiNfsGaneshaExportPostRequestFsal) GetSecLabelXattrOk() (*string, bool)`

GetSecLabelXattrOk returns a tuple with the SecLabelXattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecLabelXattr

`func (o *ApiNfsGaneshaExportPostRequestFsal) SetSecLabelXattr(v string)`

SetSecLabelXattr sets SecLabelXattr field to given value.

### HasSecLabelXattr

`func (o *ApiNfsGaneshaExportPostRequestFsal) HasSecLabelXattr() bool`

HasSecLabelXattr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


