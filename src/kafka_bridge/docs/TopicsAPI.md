# \TopicsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTopic**](TopicsAPI.md#CreateTopic) | **Post** /admin/topics | 
[**GetOffsets**](TopicsAPI.md#GetOffsets) | **Get** /topics/{topicname}/partitions/{partitionid}/offsets | 
[**GetPartition**](TopicsAPI.md#GetPartition) | **Get** /topics/{topicname}/partitions/{partitionid} | 
[**GetTopic**](TopicsAPI.md#GetTopic) | **Get** /topics/{topicname} | 
[**ListPartitions**](TopicsAPI.md#ListPartitions) | **Get** /topics/{topicname}/partitions | 
[**ListTopics**](TopicsAPI.md#ListTopics) | **Get** /topics | 
[**Send**](TopicsAPI.md#Send) | **Post** /topics/{topicname} | 
[**SendToPartition**](TopicsAPI.md#SendToPartition) | **Post** /topics/{topicname}/partitions/{partitionid} | 



## CreateTopic

> CreateTopic(ctx).NewTopic(newTopic).Execute()





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
	newTopic := *openapiclient.NewNewTopic("TopicName_example") // NewTopic | Creates a topic with given name, partitions count, and replication factor.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TopicsAPI.CreateTopic(context.Background()).NewTopic(newTopic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.CreateTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newTopic** | [**NewTopic**](NewTopic.md) | Creates a topic with given name, partitions count, and replication factor. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.json.v2+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffsets

> OffsetsSummary GetOffsets(ctx, topicname, partitionid).Execute()





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
	topicname := "topicname_example" // string | Name of the topic containing the partition.
	partitionid := int32(56) // int32 | ID of the partition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopicsAPI.GetOffsets(context.Background(), topicname, partitionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.GetOffsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOffsets`: OffsetsSummary
	fmt.Fprintf(os.Stdout, "Response from `TopicsAPI.GetOffsets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic containing the partition. | 
**partitionid** | **int32** | ID of the partition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOffsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OffsetsSummary**](OffsetsSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartition

> PartitionMetadata GetPartition(ctx, topicname, partitionid).Execute()





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
	topicname := "topicname_example" // string | Name of the topic to send records to or retrieve metadata from.
	partitionid := int32(56) // int32 | ID of the partition to send records to or retrieve metadata from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopicsAPI.GetPartition(context.Background(), topicname, partitionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPartition`: PartitionMetadata
	fmt.Fprintf(os.Stdout, "Response from `TopicsAPI.GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic to send records to or retrieve metadata from. | 
**partitionid** | **int32** | ID of the partition to send records to or retrieve metadata from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PartitionMetadata**](PartitionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopic

> TopicMetadata GetTopic(ctx, topicname).Execute()





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
	topicname := "topicname_example" // string | Name of the topic to send records to or retrieve metadata from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopicsAPI.GetTopic(context.Background(), topicname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.GetTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopic`: TopicMetadata
	fmt.Fprintf(os.Stdout, "Response from `TopicsAPI.GetTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic to send records to or retrieve metadata from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopicMetadata**](TopicMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPartitions

> []PartitionMetadata ListPartitions(ctx, topicname).Execute()





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
	topicname := "topicname_example" // string | Name of the topic to send records to or retrieve metadata from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopicsAPI.ListPartitions(context.Background(), topicname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.ListPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPartitions`: []PartitionMetadata
	fmt.Fprintf(os.Stdout, "Response from `TopicsAPI.ListPartitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic to send records to or retrieve metadata from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PartitionMetadata**](PartitionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTopics

> []string ListTopics(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopicsAPI.ListTopics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.ListTopics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTopics`: []string
	fmt.Fprintf(os.Stdout, "Response from `TopicsAPI.ListTopics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTopicsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Send

> OffsetRecordSentList Send(ctx, topicname).ProducerRecordList(producerRecordList).Async(async).Execute()





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
	topicname := "topicname_example" // string | Name of the topic to send records to or retrieve metadata from.
	producerRecordList := *openapiclient.NewProducerRecordList() // ProducerRecordList | 
	async := true // bool | Ignore metadata as result of the sending operation, not returning them to the client. If not specified it is false, metadata returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopicsAPI.Send(context.Background(), topicname).ProducerRecordList(producerRecordList).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.Send``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Send`: OffsetRecordSentList
	fmt.Fprintf(os.Stdout, "Response from `TopicsAPI.Send`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic to send records to or retrieve metadata from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **producerRecordList** | [**ProducerRecordList**](ProducerRecordList.md) |  | 
 **async** | **bool** | Ignore metadata as result of the sending operation, not returning them to the client. If not specified it is false, metadata returned. | 

### Return type

[**OffsetRecordSentList**](OffsetRecordSentList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.json.v2+json, application/vnd.kafka.binary.v2+json, application/vnd.kafka.text.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendToPartition

> OffsetRecordSentList SendToPartition(ctx, topicname, partitionid).ProducerRecordToPartitionList(producerRecordToPartitionList).Async(async).Execute()





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
	topicname := "topicname_example" // string | Name of the topic to send records to or retrieve metadata from.
	partitionid := int32(56) // int32 | ID of the partition to send records to or retrieve metadata from.
	producerRecordToPartitionList := *openapiclient.NewProducerRecordToPartitionList() // ProducerRecordToPartitionList | List of records to send to a given topic partition, including a value (required) and a key (optional).
	async := true // bool | Whether to return immediately upon sending records, instead of waiting for metadata. No offsets will be returned if specified. Defaults to false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TopicsAPI.SendToPartition(context.Background(), topicname, partitionid).ProducerRecordToPartitionList(producerRecordToPartitionList).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.SendToPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendToPartition`: OffsetRecordSentList
	fmt.Fprintf(os.Stdout, "Response from `TopicsAPI.SendToPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicname** | **string** | Name of the topic to send records to or retrieve metadata from. | 
**partitionid** | **int32** | ID of the partition to send records to or retrieve metadata from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendToPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **producerRecordToPartitionList** | [**ProducerRecordToPartitionList**](ProducerRecordToPartitionList.md) | List of records to send to a given topic partition, including a value (required) and a key (optional). | 
 **async** | **bool** | Whether to return immediately upon sending records, instead of waiting for metadata. No offsets will be returned if specified. Defaults to false. | 

### Return type

[**OffsetRecordSentList**](OffsetRecordSentList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.kafka.json.v2+json, application/vnd.kafka.binary.v2+json, application/vnd.kafka.text.v2+json
- **Accept**: application/vnd.kafka.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

