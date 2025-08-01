# OffsetCommitSeek

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | **int32** |  | 
**Offset** | **int64** |  | 
**Topic** | **string** |  | 

## Methods

### NewOffsetCommitSeek

`func NewOffsetCommitSeek(partition int32, offset int64, topic string, ) *OffsetCommitSeek`

NewOffsetCommitSeek instantiates a new OffsetCommitSeek object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffsetCommitSeekWithDefaults

`func NewOffsetCommitSeekWithDefaults() *OffsetCommitSeek`

NewOffsetCommitSeekWithDefaults instantiates a new OffsetCommitSeek object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *OffsetCommitSeek) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *OffsetCommitSeek) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *OffsetCommitSeek) SetPartition(v int32)`

SetPartition sets Partition field to given value.


### GetOffset

`func (o *OffsetCommitSeek) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *OffsetCommitSeek) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *OffsetCommitSeek) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetTopic

`func (o *OffsetCommitSeek) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *OffsetCommitSeek) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *OffsetCommitSeek) SetTopic(v string)`

SetTopic sets Topic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


