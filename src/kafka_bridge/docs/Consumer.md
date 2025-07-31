# Consumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name for the consumer instance. The name is unique within the scope of the consumer group. The name is used in URLs. If a name is not specified, a randomly generated name is assigned. | [optional] 
**Format** | Pointer to **string** | The allowable message format for the consumer, which can be &#x60;binary&#x60; (default) or &#x60;json&#x60;. The messages are converted into a JSON format.  | [optional] 
**AutoOffsetReset** | Pointer to **string** | Resets the offset position for the consumer. If set to &#x60;latest&#x60; (default), messages are read from the latest offset. If set to &#x60;earliest&#x60;, messages are read from the first offset. | [optional] 
**FetchMinBytes** | Pointer to **int32** | Sets the minimum amount of data, in bytes, for the consumer to receive. The broker waits until the data to send exceeds this amount. Default is &#x60;1&#x60; byte. | [optional] 
**ConsumerRequestTimeoutMs** | Pointer to **int32** | Sets the maximum amount of time, in milliseconds, for the consumer to wait for messages for a request. If the timeout period is reached without a response, an error is returned. Default is &#x60;30000&#x60; (30 seconds). | [optional] 
**EnableAutoCommit** | Pointer to **bool** | If set to &#x60;true&#x60; (default), message offsets are committed automatically for the consumer. If set to &#x60;false&#x60;, message offsets must be committed manually. | [optional] 
**IsolationLevel** | Pointer to **string** | If set to &#x60;read_uncommitted&#x60; (default), all transaction records are retrieved, indpendent of any transaction outcome. If set to &#x60;read_committed&#x60;, the records from committed transactions are retrieved. | [optional] 

## Methods

### NewConsumer

`func NewConsumer() *Consumer`

NewConsumer instantiates a new Consumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerWithDefaults

`func NewConsumerWithDefaults() *Consumer`

NewConsumerWithDefaults instantiates a new Consumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Consumer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Consumer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Consumer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Consumer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormat

`func (o *Consumer) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Consumer) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Consumer) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *Consumer) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetAutoOffsetReset

`func (o *Consumer) GetAutoOffsetReset() string`

GetAutoOffsetReset returns the AutoOffsetReset field if non-nil, zero value otherwise.

### GetAutoOffsetResetOk

`func (o *Consumer) GetAutoOffsetResetOk() (*string, bool)`

GetAutoOffsetResetOk returns a tuple with the AutoOffsetReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOffsetReset

`func (o *Consumer) SetAutoOffsetReset(v string)`

SetAutoOffsetReset sets AutoOffsetReset field to given value.

### HasAutoOffsetReset

`func (o *Consumer) HasAutoOffsetReset() bool`

HasAutoOffsetReset returns a boolean if a field has been set.

### GetFetchMinBytes

`func (o *Consumer) GetFetchMinBytes() int32`

GetFetchMinBytes returns the FetchMinBytes field if non-nil, zero value otherwise.

### GetFetchMinBytesOk

`func (o *Consumer) GetFetchMinBytesOk() (*int32, bool)`

GetFetchMinBytesOk returns a tuple with the FetchMinBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchMinBytes

`func (o *Consumer) SetFetchMinBytes(v int32)`

SetFetchMinBytes sets FetchMinBytes field to given value.

### HasFetchMinBytes

`func (o *Consumer) HasFetchMinBytes() bool`

HasFetchMinBytes returns a boolean if a field has been set.

### GetConsumerRequestTimeoutMs

`func (o *Consumer) GetConsumerRequestTimeoutMs() int32`

GetConsumerRequestTimeoutMs returns the ConsumerRequestTimeoutMs field if non-nil, zero value otherwise.

### GetConsumerRequestTimeoutMsOk

`func (o *Consumer) GetConsumerRequestTimeoutMsOk() (*int32, bool)`

GetConsumerRequestTimeoutMsOk returns a tuple with the ConsumerRequestTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerRequestTimeoutMs

`func (o *Consumer) SetConsumerRequestTimeoutMs(v int32)`

SetConsumerRequestTimeoutMs sets ConsumerRequestTimeoutMs field to given value.

### HasConsumerRequestTimeoutMs

`func (o *Consumer) HasConsumerRequestTimeoutMs() bool`

HasConsumerRequestTimeoutMs returns a boolean if a field has been set.

### GetEnableAutoCommit

`func (o *Consumer) GetEnableAutoCommit() bool`

GetEnableAutoCommit returns the EnableAutoCommit field if non-nil, zero value otherwise.

### GetEnableAutoCommitOk

`func (o *Consumer) GetEnableAutoCommitOk() (*bool, bool)`

GetEnableAutoCommitOk returns a tuple with the EnableAutoCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoCommit

`func (o *Consumer) SetEnableAutoCommit(v bool)`

SetEnableAutoCommit sets EnableAutoCommit field to given value.

### HasEnableAutoCommit

`func (o *Consumer) HasEnableAutoCommit() bool`

HasEnableAutoCommit returns a boolean if a field has been set.

### GetIsolationLevel

`func (o *Consumer) GetIsolationLevel() string`

GetIsolationLevel returns the IsolationLevel field if non-nil, zero value otherwise.

### GetIsolationLevelOk

`func (o *Consumer) GetIsolationLevelOk() (*string, bool)`

GetIsolationLevelOk returns a tuple with the IsolationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationLevel

`func (o *Consumer) SetIsolationLevel(v string)`

SetIsolationLevel sets IsolationLevel field to given value.

### HasIsolationLevel

`func (o *Consumer) HasIsolationLevel() bool`

HasIsolationLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


