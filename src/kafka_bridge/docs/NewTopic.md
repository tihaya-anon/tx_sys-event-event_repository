# NewTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicName** | **string** | Name of the topic to create. | 
**PartitionsCount** | Pointer to **NullableInt32** | Number of partitions for the topic. | [optional] 
**ReplicationFactor** | Pointer to **NullableInt32** | Number of replicas for each partition. | [optional] 

## Methods

### NewNewTopic

`func NewNewTopic(topicName string, ) *NewTopic`

NewNewTopic instantiates a new NewTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewTopicWithDefaults

`func NewNewTopicWithDefaults() *NewTopic`

NewNewTopicWithDefaults instantiates a new NewTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicName

`func (o *NewTopic) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *NewTopic) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *NewTopic) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPartitionsCount

`func (o *NewTopic) GetPartitionsCount() int32`

GetPartitionsCount returns the PartitionsCount field if non-nil, zero value otherwise.

### GetPartitionsCountOk

`func (o *NewTopic) GetPartitionsCountOk() (*int32, bool)`

GetPartitionsCountOk returns a tuple with the PartitionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionsCount

`func (o *NewTopic) SetPartitionsCount(v int32)`

SetPartitionsCount sets PartitionsCount field to given value.

### HasPartitionsCount

`func (o *NewTopic) HasPartitionsCount() bool`

HasPartitionsCount returns a boolean if a field has been set.

### SetPartitionsCountNil

`func (o *NewTopic) SetPartitionsCountNil(b bool)`

 SetPartitionsCountNil sets the value for PartitionsCount to be an explicit nil

### UnsetPartitionsCount
`func (o *NewTopic) UnsetPartitionsCount()`

UnsetPartitionsCount ensures that no value is present for PartitionsCount, not even an explicit nil
### GetReplicationFactor

`func (o *NewTopic) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *NewTopic) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *NewTopic) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *NewTopic) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### SetReplicationFactorNil

`func (o *NewTopic) SetReplicationFactorNil(b bool)`

 SetReplicationFactorNil sets the value for ReplicationFactor to be an explicit nil

### UnsetReplicationFactor
`func (o *NewTopic) UnsetReplicationFactor()`

UnsetReplicationFactor ensures that no value is present for ReplicationFactor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


