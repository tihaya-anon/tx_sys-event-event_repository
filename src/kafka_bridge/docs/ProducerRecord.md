# ProducerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Value** | [**NullableRecordValue**](RecordValue.md) |  | 
**Key** | Pointer to [**RecordKey**](RecordKey.md) |  | [optional] 
**Headers** | Pointer to [**[]KafkaHeader**](KafkaHeader.md) |  | [optional] 

## Methods

### NewProducerRecord

`func NewProducerRecord(value NullableRecordValue, ) *ProducerRecord`

NewProducerRecord instantiates a new ProducerRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProducerRecordWithDefaults

`func NewProducerRecordWithDefaults() *ProducerRecord`

NewProducerRecordWithDefaults instantiates a new ProducerRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *ProducerRecord) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *ProducerRecord) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *ProducerRecord) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *ProducerRecord) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetTimestamp

`func (o *ProducerRecord) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ProducerRecord) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ProducerRecord) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ProducerRecord) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *ProducerRecord) GetValue() RecordValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProducerRecord) GetValueOk() (*RecordValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProducerRecord) SetValue(v RecordValue)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ProducerRecord) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProducerRecord) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetKey

`func (o *ProducerRecord) GetKey() RecordKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProducerRecord) GetKeyOk() (*RecordKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProducerRecord) SetKey(v RecordKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProducerRecord) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetHeaders

`func (o *ProducerRecord) GetHeaders() []KafkaHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ProducerRecord) GetHeadersOk() (*[]KafkaHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ProducerRecord) SetHeaders(v []KafkaHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ProducerRecord) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


