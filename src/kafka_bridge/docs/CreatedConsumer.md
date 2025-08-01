# CreatedConsumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **string** | Unique ID for the consumer instance in the group. | [optional] 
**BaseUri** | Pointer to **string** | Base URI used to construct URIs for subsequent requests against this consumer instance. | [optional] 

## Methods

### NewCreatedConsumer

`func NewCreatedConsumer() *CreatedConsumer`

NewCreatedConsumer instantiates a new CreatedConsumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedConsumerWithDefaults

`func NewCreatedConsumerWithDefaults() *CreatedConsumer`

NewCreatedConsumerWithDefaults instantiates a new CreatedConsumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *CreatedConsumer) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CreatedConsumer) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CreatedConsumer) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *CreatedConsumer) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetBaseUri

`func (o *CreatedConsumer) GetBaseUri() string`

GetBaseUri returns the BaseUri field if non-nil, zero value otherwise.

### GetBaseUriOk

`func (o *CreatedConsumer) GetBaseUriOk() (*string, bool)`

GetBaseUriOk returns a tuple with the BaseUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUri

`func (o *CreatedConsumer) SetBaseUri(v string)`

SetBaseUri sets BaseUri field to given value.

### HasBaseUri

`func (o *CreatedConsumer) HasBaseUri() bool`

HasBaseUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


