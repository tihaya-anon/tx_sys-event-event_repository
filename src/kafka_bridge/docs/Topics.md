# Topics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topics** | Pointer to **[]string** |  | [optional] 
**TopicPattern** | Pointer to **string** | A regex topic pattern for matching multiple topics | [optional] 

## Methods

### NewTopics

`func NewTopics() *Topics`

NewTopics instantiates a new Topics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicsWithDefaults

`func NewTopicsWithDefaults() *Topics`

NewTopicsWithDefaults instantiates a new Topics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopics

`func (o *Topics) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *Topics) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *Topics) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *Topics) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetTopicPattern

`func (o *Topics) GetTopicPattern() string`

GetTopicPattern returns the TopicPattern field if non-nil, zero value otherwise.

### GetTopicPatternOk

`func (o *Topics) GetTopicPatternOk() (*string, bool)`

GetTopicPatternOk returns a tuple with the TopicPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicPattern

`func (o *Topics) SetTopicPattern(v string)`

SetTopicPattern sets TopicPattern field to given value.

### HasTopicPattern

`func (o *Topics) HasTopicPattern() bool`

HasTopicPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


