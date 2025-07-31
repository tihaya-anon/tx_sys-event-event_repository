# TopicMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the topic | [optional] 
**Configs** | Pointer to **map[string]string** | Per-topic configuration overrides | [optional] 
**Partitions** | Pointer to [**[]PartitionMetadata**](PartitionMetadata.md) |  | [optional] 

## Methods

### NewTopicMetadata

`func NewTopicMetadata() *TopicMetadata`

NewTopicMetadata instantiates a new TopicMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicMetadataWithDefaults

`func NewTopicMetadataWithDefaults() *TopicMetadata`

NewTopicMetadataWithDefaults instantiates a new TopicMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TopicMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopicMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopicMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopicMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigs

`func (o *TopicMetadata) GetConfigs() map[string]string`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *TopicMetadata) GetConfigsOk() (*map[string]string, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *TopicMetadata) SetConfigs(v map[string]string)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *TopicMetadata) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetPartitions

`func (o *TopicMetadata) GetPartitions() []PartitionMetadata`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *TopicMetadata) GetPartitionsOk() (*[]PartitionMetadata, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *TopicMetadata) SetPartitions(v []PartitionMetadata)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *TopicMetadata) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


