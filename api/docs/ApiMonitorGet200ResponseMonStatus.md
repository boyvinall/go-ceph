# ApiMonitorGet200ResponseMonStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElectionEpoch** | **int32** |  | 
**ExtraProbePeers** | **[]string** |  | 
**FeatureMap** | [**ApiMonitorGet200ResponseMonStatusFeatureMap**](ApiMonitorGet200ResponseMonStatusFeatureMap.md) |  | 
**Features** | [**ApiMonitorGet200ResponseMonStatusFeatures**](ApiMonitorGet200ResponseMonStatusFeatures.md) |  | 
**Monmap** | [**ApiMonitorGet200ResponseMonStatusMonmap**](ApiMonitorGet200ResponseMonStatusMonmap.md) |  | 
**Name** | **string** |  | 
**OutsideQuorum** | **[]string** |  | 
**Quorum** | **[]int32** |  | 
**QuorumAge** | **int32** |  | 
**Rank** | **int32** |  | 
**State** | **string** |  | 
**SyncProvider** | **[]string** |  | 

## Methods

### NewApiMonitorGet200ResponseMonStatus

`func NewApiMonitorGet200ResponseMonStatus(electionEpoch int32, extraProbePeers []string, featureMap ApiMonitorGet200ResponseMonStatusFeatureMap, features ApiMonitorGet200ResponseMonStatusFeatures, monmap ApiMonitorGet200ResponseMonStatusMonmap, name string, outsideQuorum []string, quorum []int32, quorumAge int32, rank int32, state string, syncProvider []string, ) *ApiMonitorGet200ResponseMonStatus`

NewApiMonitorGet200ResponseMonStatus instantiates a new ApiMonitorGet200ResponseMonStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMonitorGet200ResponseMonStatusWithDefaults

`func NewApiMonitorGet200ResponseMonStatusWithDefaults() *ApiMonitorGet200ResponseMonStatus`

NewApiMonitorGet200ResponseMonStatusWithDefaults instantiates a new ApiMonitorGet200ResponseMonStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElectionEpoch

`func (o *ApiMonitorGet200ResponseMonStatus) GetElectionEpoch() int32`

GetElectionEpoch returns the ElectionEpoch field if non-nil, zero value otherwise.

### GetElectionEpochOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetElectionEpochOk() (*int32, bool)`

GetElectionEpochOk returns a tuple with the ElectionEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectionEpoch

`func (o *ApiMonitorGet200ResponseMonStatus) SetElectionEpoch(v int32)`

SetElectionEpoch sets ElectionEpoch field to given value.


### GetExtraProbePeers

`func (o *ApiMonitorGet200ResponseMonStatus) GetExtraProbePeers() []string`

GetExtraProbePeers returns the ExtraProbePeers field if non-nil, zero value otherwise.

### GetExtraProbePeersOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetExtraProbePeersOk() (*[]string, bool)`

GetExtraProbePeersOk returns a tuple with the ExtraProbePeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProbePeers

`func (o *ApiMonitorGet200ResponseMonStatus) SetExtraProbePeers(v []string)`

SetExtraProbePeers sets ExtraProbePeers field to given value.


### GetFeatureMap

`func (o *ApiMonitorGet200ResponseMonStatus) GetFeatureMap() ApiMonitorGet200ResponseMonStatusFeatureMap`

GetFeatureMap returns the FeatureMap field if non-nil, zero value otherwise.

### GetFeatureMapOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetFeatureMapOk() (*ApiMonitorGet200ResponseMonStatusFeatureMap, bool)`

GetFeatureMapOk returns a tuple with the FeatureMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureMap

`func (o *ApiMonitorGet200ResponseMonStatus) SetFeatureMap(v ApiMonitorGet200ResponseMonStatusFeatureMap)`

SetFeatureMap sets FeatureMap field to given value.


### GetFeatures

`func (o *ApiMonitorGet200ResponseMonStatus) GetFeatures() ApiMonitorGet200ResponseMonStatusFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetFeaturesOk() (*ApiMonitorGet200ResponseMonStatusFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ApiMonitorGet200ResponseMonStatus) SetFeatures(v ApiMonitorGet200ResponseMonStatusFeatures)`

SetFeatures sets Features field to given value.


### GetMonmap

`func (o *ApiMonitorGet200ResponseMonStatus) GetMonmap() ApiMonitorGet200ResponseMonStatusMonmap`

GetMonmap returns the Monmap field if non-nil, zero value otherwise.

### GetMonmapOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetMonmapOk() (*ApiMonitorGet200ResponseMonStatusMonmap, bool)`

GetMonmapOk returns a tuple with the Monmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonmap

`func (o *ApiMonitorGet200ResponseMonStatus) SetMonmap(v ApiMonitorGet200ResponseMonStatusMonmap)`

SetMonmap sets Monmap field to given value.


### GetName

`func (o *ApiMonitorGet200ResponseMonStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMonitorGet200ResponseMonStatus) SetName(v string)`

SetName sets Name field to given value.


### GetOutsideQuorum

`func (o *ApiMonitorGet200ResponseMonStatus) GetOutsideQuorum() []string`

GetOutsideQuorum returns the OutsideQuorum field if non-nil, zero value otherwise.

### GetOutsideQuorumOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetOutsideQuorumOk() (*[]string, bool)`

GetOutsideQuorumOk returns a tuple with the OutsideQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideQuorum

`func (o *ApiMonitorGet200ResponseMonStatus) SetOutsideQuorum(v []string)`

SetOutsideQuorum sets OutsideQuorum field to given value.


### GetQuorum

`func (o *ApiMonitorGet200ResponseMonStatus) GetQuorum() []int32`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetQuorumOk() (*[]int32, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *ApiMonitorGet200ResponseMonStatus) SetQuorum(v []int32)`

SetQuorum sets Quorum field to given value.


### GetQuorumAge

`func (o *ApiMonitorGet200ResponseMonStatus) GetQuorumAge() int32`

GetQuorumAge returns the QuorumAge field if non-nil, zero value otherwise.

### GetQuorumAgeOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetQuorumAgeOk() (*int32, bool)`

GetQuorumAgeOk returns a tuple with the QuorumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorumAge

`func (o *ApiMonitorGet200ResponseMonStatus) SetQuorumAge(v int32)`

SetQuorumAge sets QuorumAge field to given value.


### GetRank

`func (o *ApiMonitorGet200ResponseMonStatus) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ApiMonitorGet200ResponseMonStatus) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetState

`func (o *ApiMonitorGet200ResponseMonStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiMonitorGet200ResponseMonStatus) SetState(v string)`

SetState sets State field to given value.


### GetSyncProvider

`func (o *ApiMonitorGet200ResponseMonStatus) GetSyncProvider() []string`

GetSyncProvider returns the SyncProvider field if non-nil, zero value otherwise.

### GetSyncProviderOk

`func (o *ApiMonitorGet200ResponseMonStatus) GetSyncProviderOk() (*[]string, bool)`

GetSyncProviderOk returns a tuple with the SyncProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncProvider

`func (o *ApiMonitorGet200ResponseMonStatus) SetSyncProvider(v []string)`

SetSyncProvider sets SyncProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


