# ApiFeedbackPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**Project** | **string** |  | 
**Subject** | **string** |  | 
**Tracker** | **string** |  | 

## Methods

### NewApiFeedbackPostRequest

`func NewApiFeedbackPostRequest(description string, project string, subject string, tracker string, ) *ApiFeedbackPostRequest`

NewApiFeedbackPostRequest instantiates a new ApiFeedbackPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFeedbackPostRequestWithDefaults

`func NewApiFeedbackPostRequestWithDefaults() *ApiFeedbackPostRequest`

NewApiFeedbackPostRequestWithDefaults instantiates a new ApiFeedbackPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *ApiFeedbackPostRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ApiFeedbackPostRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ApiFeedbackPostRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ApiFeedbackPostRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetDescription

`func (o *ApiFeedbackPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiFeedbackPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiFeedbackPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProject

`func (o *ApiFeedbackPostRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ApiFeedbackPostRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ApiFeedbackPostRequest) SetProject(v string)`

SetProject sets Project field to given value.


### GetSubject

`func (o *ApiFeedbackPostRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ApiFeedbackPostRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ApiFeedbackPostRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTracker

`func (o *ApiFeedbackPostRequest) GetTracker() string`

GetTracker returns the Tracker field if non-nil, zero value otherwise.

### GetTrackerOk

`func (o *ApiFeedbackPostRequest) GetTrackerOk() (*string, bool)`

GetTrackerOk returns a tuple with the Tracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracker

`func (o *ApiFeedbackPostRequest) SetTracker(v string)`

SetTracker sets Tracker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


