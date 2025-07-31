# ProducerRecordToPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**NullableRecordValue**](RecordValue.md) |  | 
**Key** | Pointer to [**RecordKey**](RecordKey.md) |  | [optional] 
**Headers** | Pointer to [**[]KafkaHeader**](KafkaHeader.md) |  | [optional] 

## Methods

### NewProducerRecordToPartition

`func NewProducerRecordToPartition(value NullableRecordValue, ) *ProducerRecordToPartition`

NewProducerRecordToPartition instantiates a new ProducerRecordToPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProducerRecordToPartitionWithDefaults

`func NewProducerRecordToPartitionWithDefaults() *ProducerRecordToPartition`

NewProducerRecordToPartitionWithDefaults instantiates a new ProducerRecordToPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ProducerRecordToPartition) GetValue() RecordValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProducerRecordToPartition) GetValueOk() (*RecordValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProducerRecordToPartition) SetValue(v RecordValue)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ProducerRecordToPartition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProducerRecordToPartition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetKey

`func (o *ProducerRecordToPartition) GetKey() RecordKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProducerRecordToPartition) GetKeyOk() (*RecordKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProducerRecordToPartition) SetKey(v RecordKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProducerRecordToPartition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetHeaders

`func (o *ProducerRecordToPartition) GetHeaders() []KafkaHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ProducerRecordToPartition) GetHeadersOk() (*[]KafkaHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ProducerRecordToPartition) SetHeaders(v []KafkaHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ProducerRecordToPartition) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


