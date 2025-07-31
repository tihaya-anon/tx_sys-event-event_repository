# ConsumerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**RecordKey**](RecordKey.md) |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Partition** | Pointer to **int32** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**NullableRecordValue**](RecordValue.md) |  | [optional] 
**Headers** | Pointer to [**[]KafkaHeader**](KafkaHeader.md) |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewConsumerRecord

`func NewConsumerRecord() *ConsumerRecord`

NewConsumerRecord instantiates a new ConsumerRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerRecordWithDefaults

`func NewConsumerRecordWithDefaults() *ConsumerRecord`

NewConsumerRecordWithDefaults instantiates a new ConsumerRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ConsumerRecord) GetKey() RecordKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConsumerRecord) GetKeyOk() (*RecordKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConsumerRecord) SetKey(v RecordKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *ConsumerRecord) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOffset

`func (o *ConsumerRecord) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ConsumerRecord) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ConsumerRecord) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ConsumerRecord) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPartition

`func (o *ConsumerRecord) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ConsumerRecord) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ConsumerRecord) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *ConsumerRecord) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetTopic

`func (o *ConsumerRecord) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ConsumerRecord) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ConsumerRecord) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ConsumerRecord) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetValue

`func (o *ConsumerRecord) GetValue() RecordValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConsumerRecord) GetValueOk() (*RecordValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConsumerRecord) SetValue(v RecordValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConsumerRecord) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ConsumerRecord) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ConsumerRecord) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetHeaders

`func (o *ConsumerRecord) GetHeaders() []KafkaHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ConsumerRecord) GetHeadersOk() (*[]KafkaHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ConsumerRecord) SetHeaders(v []KafkaHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ConsumerRecord) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetTimestamp

`func (o *ConsumerRecord) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ConsumerRecord) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ConsumerRecord) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ConsumerRecord) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


