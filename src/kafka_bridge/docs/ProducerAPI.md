# \ProducerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Send**](ProducerAPI.md#Send) | **Post** /topics/{topicname} | 
[**SendToPartition**](ProducerAPI.md#SendToPartition) | **Post** /topics/{topicname}/partitions/{partitionid} | 



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
	resp, r, err := apiClient.ProducerAPI.Send(context.Background(), topicname).ProducerRecordList(producerRecordList).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProducerAPI.Send``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Send`: OffsetRecordSentList
	fmt.Fprintf(os.Stdout, "Response from `ProducerAPI.Send`: %v\n", resp)
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
	resp, r, err := apiClient.ProducerAPI.SendToPartition(context.Background(), topicname, partitionid).ProducerRecordToPartitionList(producerRecordToPartitionList).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProducerAPI.SendToPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendToPartition`: OffsetRecordSentList
	fmt.Fprintf(os.Stdout, "Response from `ProducerAPI.SendToPartition`: %v\n", resp)
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

