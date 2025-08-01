# PartitionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | Pointer to **int32** |  | [optional] 
**Leader** | Pointer to **int32** |  | [optional] 
**Replicas** | Pointer to [**[]Replica**](Replica.md) |  | [optional] 

## Methods

### NewPartitionMetadata

`func NewPartitionMetadata() *PartitionMetadata`

NewPartitionMetadata instantiates a new PartitionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionMetadataWithDefaults

`func NewPartitionMetadataWithDefaults() *PartitionMetadata`

NewPartitionMetadataWithDefaults instantiates a new PartitionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *PartitionMetadata) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *PartitionMetadata) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *PartitionMetadata) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *PartitionMetadata) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetLeader

`func (o *PartitionMetadata) GetLeader() int32`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *PartitionMetadata) GetLeaderOk() (*int32, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *PartitionMetadata) SetLeader(v int32)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *PartitionMetadata) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetReplicas

`func (o *PartitionMetadata) GetReplicas() []Replica`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PartitionMetadata) GetReplicasOk() (*[]Replica, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PartitionMetadata) SetReplicas(v []Replica)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PartitionMetadata) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


