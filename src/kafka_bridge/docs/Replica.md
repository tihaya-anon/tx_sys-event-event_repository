# Replica

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Broker** | Pointer to **int32** |  | [optional] 
**Leader** | Pointer to **bool** |  | [optional] 
**InSync** | Pointer to **bool** |  | [optional] 

## Methods

### NewReplica

`func NewReplica() *Replica`

NewReplica instantiates a new Replica object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaWithDefaults

`func NewReplicaWithDefaults() *Replica`

NewReplicaWithDefaults instantiates a new Replica object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroker

`func (o *Replica) GetBroker() int32`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *Replica) GetBrokerOk() (*int32, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *Replica) SetBroker(v int32)`

SetBroker sets Broker field to given value.

### HasBroker

`func (o *Replica) HasBroker() bool`

HasBroker returns a boolean if a field has been set.

### GetLeader

`func (o *Replica) GetLeader() bool`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *Replica) GetLeaderOk() (*bool, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *Replica) SetLeader(v bool)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *Replica) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetInSync

`func (o *Replica) GetInSync() bool`

GetInSync returns the InSync field if non-nil, zero value otherwise.

### GetInSyncOk

`func (o *Replica) GetInSyncOk() (*bool, bool)`

GetInSyncOk returns a tuple with the InSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSync

`func (o *Replica) SetInSync(v bool)`

SetInSync sets InSync field to given value.

### HasInSync

`func (o *Replica) HasInSync() bool`

HasInSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


