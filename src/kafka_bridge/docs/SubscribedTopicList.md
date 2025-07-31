# SubscribedTopicList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topics** | Pointer to [**Topics**](Topics.md) |  | [optional] 
**Partitions** | Pointer to [**[]map[string][]int32**](map[string][]int32.md) |  | [optional] 

## Methods

### NewSubscribedTopicList

`func NewSubscribedTopicList() *SubscribedTopicList`

NewSubscribedTopicList instantiates a new SubscribedTopicList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribedTopicListWithDefaults

`func NewSubscribedTopicListWithDefaults() *SubscribedTopicList`

NewSubscribedTopicListWithDefaults instantiates a new SubscribedTopicList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopics

`func (o *SubscribedTopicList) GetTopics() Topics`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *SubscribedTopicList) GetTopicsOk() (*Topics, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *SubscribedTopicList) SetTopics(v Topics)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *SubscribedTopicList) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetPartitions

`func (o *SubscribedTopicList) GetPartitions() []map[string][]int32`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *SubscribedTopicList) GetPartitionsOk() (*[]map[string][]int32, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *SubscribedTopicList) SetPartitions(v []map[string][]int32)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *SubscribedTopicList) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


