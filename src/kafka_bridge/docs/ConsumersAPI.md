# \ConsumersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Assign**](ConsumersAPI.md#Assign) | **Post** /consumers/{groupid}/instances/{name}/assignments | 
[**Commit**](ConsumersAPI.md#Commit) | **Post** /consumers/{groupid}/instances/{name}/offsets | 
[**CreateConsumer**](ConsumersAPI.md#CreateConsumer) | **Post** /consumers/{groupid} | 
[**DeleteConsumer**](ConsumersAPI.md#DeleteConsumer) | **Delete** /consumers/{groupid}/instances/{name} | 
[**ListSubscriptions**](ConsumersAPI.md#ListSubscriptions) | **Get** /consumers/{groupid}/instances/{name}/subscription | 
[**Poll**](ConsumersAPI.md#Poll) | **Get** /consumers/{groupid}/instances/{name}/records | 
[**Seek**](ConsumersAPI.md#Seek) | **Post** /consumers/{groupid}/instances/{name}/positions | 
[**SeekToBeginning**](ConsumersAPI.md#SeekToBeginning) | **Post** /consumers/{groupid}/instances/{name}/positions/beginning | 
[**SeekToEnd**](ConsumersAPI.md#SeekToEnd) | **Post** /consumers/{groupid}/instances/{name}/positions/end | 
[**Subscribe**](ConsumersAPI.md#Subscribe) | **Post** /consumers/{groupid}/instances/{name}/subscription | 
[**Unsubscribe**](ConsumersAPI.md#Unsubscribe) | **Delete** /consumers/{groupid}/instances/{name}/subscription | 



## Assign

> Assign(ctx, groupid, name).Partitions(partitions).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group to which the consumer belongs.
	name := "name_example" // string | Name of the consumer to assign topic partitions to.
	partitions := *openapiclient.NewPartitions() // Partitions | List of topic partitions to assign to the consumer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumersAPI.Assign(context.Background(), groupid, name).Partitions(partitions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.Assign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the consumer belongs. | 
**name** | **string** | Name of the consumer to assign topic partitions to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **partitions** | [**Partitions**](Partitions.md) | List of topic partitions to assign to the consumer. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Commit

> Commit(ctx, groupid, name).OffsetCommitSeekList(offsetCommitSeekList).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group to which the consumer belongs.
	name := "name_example" // string | Name of the consumer.
	offsetCommitSeekList := *openapiclient.NewOffsetCommitSeekList() // OffsetCommitSeekList | List of consumer offsets to commit to the consumer offsets commit log. You can specify one or more topic partitions to commit offsets for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumersAPI.Commit(context.Background(), groupid, name).OffsetCommitSeekList(offsetCommitSeekList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.Commit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the consumer belongs. | 
**name** | **string** | Name of the consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offsetCommitSeekList** | [**OffsetCommitSeekList**](OffsetCommitSeekList.md) | List of consumer offsets to commit to the consumer offsets commit log. You can specify one or more topic partitions to commit offsets for. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConsumer

> CreatedConsumer CreateConsumer(ctx, groupid).Consumer(consumer).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group in which to create the consumer.
	consumer := *openapiclient.NewConsumer() // Consumer | Name and configuration of the consumer. The name is unique within the scope of the consumer group. If a name is not specified, a randomly generated name is assigned. All parameters are optional. The supported configuration options are shown in the following example. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumersAPI.CreateConsumer(context.Background(), groupid).Consumer(consumer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.CreateConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConsumer`: CreatedConsumer
	fmt.Fprintf(os.Stdout, "Response from `ConsumersAPI.CreateConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group in which to create the consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumer** | [**Consumer**](Consumer.md) | Name and configuration of the consumer. The name is unique within the scope of the consumer group. If a name is not specified, a randomly generated name is assigned. All parameters are optional. The supported configuration options are shown in the following example. | 

### Return type

[**CreatedConsumer**](CreatedConsumer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConsumer

> DeleteConsumer(ctx, groupid, name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group to which the consumer belongs.
	name := "name_example" // string | Name of the consumer to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumersAPI.DeleteConsumer(context.Background(), groupid, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.DeleteConsumer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the consumer belongs. | 
**name** | **string** | Name of the consumer to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptions

> SubscribedTopicList ListSubscriptions(ctx, groupid, name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group to which the subscribed consumer belongs.
	name := "name_example" // string | Name of the subscribed consumer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumersAPI.ListSubscriptions(context.Background(), groupid, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.ListSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubscriptions`: SubscribedTopicList
	fmt.Fprintf(os.Stdout, "Response from `ConsumersAPI.ListSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the subscribed consumer belongs. | 
**name** | **string** | Name of the subscribed consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubscribedTopicList**](SubscribedTopicList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Poll

> []ConsumerRecord Poll(ctx, groupid, name).Timeout(timeout).MaxBytes(maxBytes).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group to which the subscribed consumer belongs.
	name := "name_example" // string | Name of the subscribed consumer to retrieve records from.
	timeout := int32(56) // int32 | The maximum amount of time, in milliseconds, that the HTTP Bridge spends retrieving records before timing out the request. (optional)
	maxBytes := int32(56) // int32 | The maximum size, in bytes, of unencoded keys and values that can be included in the response. Otherwise, an error response with code 422 is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsumersAPI.Poll(context.Background(), groupid, name).Timeout(timeout).MaxBytes(maxBytes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.Poll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Poll`: []ConsumerRecord
	fmt.Fprintf(os.Stdout, "Response from `ConsumersAPI.Poll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the subscribed consumer belongs. | 
**name** | **string** | Name of the subscribed consumer to retrieve records from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeout** | **int32** | The maximum amount of time, in milliseconds, that the HTTP Bridge spends retrieving records before timing out the request. | 
 **maxBytes** | **int32** | The maximum size, in bytes, of unencoded keys and values that can be included in the response. Otherwise, an error response with code 422 is returned. | 

### Return type

[**[]ConsumerRecord**](ConsumerRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.json.v2+json, application/vnd.kafka.binary.v2+json, application/vnd.kafka.text.v2+json, application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Seek

> Seek(ctx, groupid, name).OffsetCommitSeekList(offsetCommitSeekList).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group to which the consumer belongs.
	name := "name_example" // string | Name of the subscribed consumer.
	offsetCommitSeekList := *openapiclient.NewOffsetCommitSeekList() // OffsetCommitSeekList | List of partition offsets from which the subscribed consumer will next fetch records.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumersAPI.Seek(context.Background(), groupid, name).OffsetCommitSeekList(offsetCommitSeekList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.Seek``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the consumer belongs. | 
**name** | **string** | Name of the subscribed consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSeekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offsetCommitSeekList** | [**OffsetCommitSeekList**](OffsetCommitSeekList.md) | List of partition offsets from which the subscribed consumer will next fetch records. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeekToBeginning

> SeekToBeginning(ctx, groupid, name).Partitions(partitions).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group to which the subscribed consumer belongs.
	name := "name_example" // string | Name of the subscribed consumer.
	partitions := *openapiclient.NewPartitions() // Partitions | List of topic partitions to which the consumer is subscribed. The consumer will seek the first offset in the specified partitions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumersAPI.SeekToBeginning(context.Background(), groupid, name).Partitions(partitions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.SeekToBeginning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the subscribed consumer belongs. | 
**name** | **string** | Name of the subscribed consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSeekToBeginningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **partitions** | [**Partitions**](Partitions.md) | List of topic partitions to which the consumer is subscribed. The consumer will seek the first offset in the specified partitions. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeekToEnd

> SeekToEnd(ctx, groupid, name).Partitions(partitions).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group to which the subscribed consumer belongs.
	name := "name_example" // string | Name of the subscribed consumer.
	partitions := *openapiclient.NewPartitions() // Partitions | List of topic partitions to which the consumer is subscribed. The consumer will seek the last offset in the specified partitions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumersAPI.SeekToEnd(context.Background(), groupid, name).Partitions(partitions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.SeekToEnd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the subscribed consumer belongs. | 
**name** | **string** | Name of the subscribed consumer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSeekToEndRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **partitions** | [**Partitions**](Partitions.md) | List of topic partitions to which the consumer is subscribed. The consumer will seek the last offset in the specified partitions. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Subscribe

> Subscribe(ctx, groupid, name).Topics(topics).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group to which the subscribed consumer belongs.
	name := "name_example" // string | Name of the consumer to subscribe to topics.
	topics := *openapiclient.NewTopics() // Topics | List of topics to which the consumer will subscribe.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumersAPI.Subscribe(context.Background(), groupid, name).Topics(topics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.Subscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the subscribed consumer belongs. | 
**name** | **string** | Name of the consumer to subscribe to topics. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **topics** | [**Topics**](Topics.md) | List of topics to which the consumer will subscribe. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unsubscribe

> Unsubscribe(ctx, groupid, name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
)

func main() {
	groupid := "groupid_example" // string | ID of the consumer group to which the subscribed consumer belongs.
	name := "name_example" // string | Name of the consumer to unsubscribe from topics.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsumersAPI.Unsubscribe(context.Background(), groupid, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsumersAPI.Unsubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupid** | **string** | ID of the consumer group to which the subscribed consumer belongs. | 
**name** | **string** | Name of the consumer to unsubscribe from topics. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

