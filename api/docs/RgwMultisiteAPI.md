# \RgwMultisiteAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRgwMultisiteSyncFlowFlowIdFlowTypeGroupIdDelete**](RgwMultisiteAPI.md#ApiRgwMultisiteSyncFlowFlowIdFlowTypeGroupIdDelete) | **Delete** /api/rgw/multisite/sync-flow/{flow_id}/{flow_type}/{group_id} | Remove the sync flow
[**ApiRgwMultisiteSyncFlowPut**](RgwMultisiteAPI.md#ApiRgwMultisiteSyncFlowPut) | **Put** /api/rgw/multisite/sync-flow | Create or update the sync flow
[**ApiRgwMultisiteSyncPipeGroupIdPipeIdDelete**](RgwMultisiteAPI.md#ApiRgwMultisiteSyncPipeGroupIdPipeIdDelete) | **Delete** /api/rgw/multisite/sync-pipe/{group_id}/{pipe_id} | Remove the sync pipe
[**ApiRgwMultisiteSyncPipePut**](RgwMultisiteAPI.md#ApiRgwMultisiteSyncPipePut) | **Put** /api/rgw/multisite/sync-pipe | Create or update the sync pipe
[**ApiRgwMultisiteSyncPolicyGet**](RgwMultisiteAPI.md#ApiRgwMultisiteSyncPolicyGet) | **Get** /api/rgw/multisite/sync-policy | Get the sync policy
[**ApiRgwMultisiteSyncPolicyGroupGroupIdDelete**](RgwMultisiteAPI.md#ApiRgwMultisiteSyncPolicyGroupGroupIdDelete) | **Delete** /api/rgw/multisite/sync-policy-group/{group_id} | Remove the sync policy group
[**ApiRgwMultisiteSyncPolicyGroupGroupIdGet**](RgwMultisiteAPI.md#ApiRgwMultisiteSyncPolicyGroupGroupIdGet) | **Get** /api/rgw/multisite/sync-policy-group/{group_id} | Get the sync policy group
[**ApiRgwMultisiteSyncPolicyGroupPost**](RgwMultisiteAPI.md#ApiRgwMultisiteSyncPolicyGroupPost) | **Post** /api/rgw/multisite/sync-policy-group | Create the sync policy group
[**ApiRgwMultisiteSyncPolicyGroupPut**](RgwMultisiteAPI.md#ApiRgwMultisiteSyncPolicyGroupPut) | **Put** /api/rgw/multisite/sync-policy-group | Update the sync policy group
[**ApiRgwMultisiteSyncStatusGet**](RgwMultisiteAPI.md#ApiRgwMultisiteSyncStatusGet) | **Get** /api/rgw/multisite/sync_status | Get the sync status



## ApiRgwMultisiteSyncFlowFlowIdFlowTypeGroupIdDelete

> ApiRgwMultisiteSyncFlowFlowIdFlowTypeGroupIdDelete(ctx, flowId, flowType, groupId).SourceZone(sourceZone).DestinationZone(destinationZone).Zones(zones).BucketName(bucketName).Execute()

Remove the sync flow

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	flowId := "flowId_example" // string | 
	flowType := "flowType_example" // string | 
	groupId := "groupId_example" // string | 
	sourceZone := "sourceZone_example" // string |  (optional)
	destinationZone := "destinationZone_example" // string |  (optional)
	zones := "zones_example" // string |  (optional)
	bucketName := "bucketName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncFlowFlowIdFlowTypeGroupIdDelete(context.Background(), flowId, flowType, groupId).SourceZone(sourceZone).DestinationZone(destinationZone).Zones(zones).BucketName(bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwMultisiteAPI.ApiRgwMultisiteSyncFlowFlowIdFlowTypeGroupIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 
**flowType** | **string** |  | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwMultisiteSyncFlowFlowIdFlowTypeGroupIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sourceZone** | **string** |  | 
 **destinationZone** | **string** |  | 
 **zones** | **string** |  | 
 **bucketName** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwMultisiteSyncFlowPut

> ApiRgwMultisiteSyncFlowPut(ctx).ApiRgwMultisiteSyncFlowPutRequest(apiRgwMultisiteSyncFlowPutRequest).Execute()

Create or update the sync flow

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	apiRgwMultisiteSyncFlowPutRequest := *openapiclient.NewApiRgwMultisiteSyncFlowPutRequest("FlowId_example", "FlowType_example", "GroupId_example") // ApiRgwMultisiteSyncFlowPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncFlowPut(context.Background()).ApiRgwMultisiteSyncFlowPutRequest(apiRgwMultisiteSyncFlowPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwMultisiteAPI.ApiRgwMultisiteSyncFlowPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwMultisiteSyncFlowPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwMultisiteSyncFlowPutRequest** | [**ApiRgwMultisiteSyncFlowPutRequest**](ApiRgwMultisiteSyncFlowPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwMultisiteSyncPipeGroupIdPipeIdDelete

> ApiRgwMultisiteSyncPipeGroupIdPipeIdDelete(ctx, groupId, pipeId).SourceZones(sourceZones).DestinationZones(destinationZones).BucketName(bucketName).Execute()

Remove the sync pipe

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	groupId := "groupId_example" // string | 
	pipeId := "pipeId_example" // string | 
	sourceZones := "sourceZones_example" // string |  (optional)
	destinationZones := "destinationZones_example" // string |  (optional)
	bucketName := "bucketName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPipeGroupIdPipeIdDelete(context.Background(), groupId, pipeId).SourceZones(sourceZones).DestinationZones(destinationZones).BucketName(bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwMultisiteAPI.ApiRgwMultisiteSyncPipeGroupIdPipeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**pipeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwMultisiteSyncPipeGroupIdPipeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sourceZones** | **string** |  | 
 **destinationZones** | **string** |  | 
 **bucketName** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwMultisiteSyncPipePut

> ApiRgwMultisiteSyncPipePut(ctx).ApiRgwMultisiteSyncPipePutRequest(apiRgwMultisiteSyncPipePutRequest).Execute()

Create or update the sync pipe

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	apiRgwMultisiteSyncPipePutRequest := *openapiclient.NewApiRgwMultisiteSyncPipePutRequest("DestinationZones_example", "GroupId_example", "PipeId_example", "SourceZones_example") // ApiRgwMultisiteSyncPipePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPipePut(context.Background()).ApiRgwMultisiteSyncPipePutRequest(apiRgwMultisiteSyncPipePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwMultisiteAPI.ApiRgwMultisiteSyncPipePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwMultisiteSyncPipePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwMultisiteSyncPipePutRequest** | [**ApiRgwMultisiteSyncPipePutRequest**](ApiRgwMultisiteSyncPipePutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwMultisiteSyncPolicyGet

> ApiRgwMultisiteSyncPolicyGet(ctx).BucketName(bucketName).ZonegroupName(zonegroupName).AllPolicy(allPolicy).Execute()

Get the sync policy

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	bucketName := "bucketName_example" // string |  (optional)
	zonegroupName := "zonegroupName_example" // string |  (optional)
	allPolicy := "allPolicy_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGet(context.Background()).BucketName(bucketName).ZonegroupName(zonegroupName).AllPolicy(allPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwMultisiteSyncPolicyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketName** | **string** |  | 
 **zonegroupName** | **string** |  | 
 **allPolicy** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwMultisiteSyncPolicyGroupGroupIdDelete

> ApiRgwMultisiteSyncPolicyGroupGroupIdDelete(ctx, groupId).BucketName(bucketName).Execute()

Remove the sync policy group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	groupId := "groupId_example" // string | 
	bucketName := "bucketName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupGroupIdDelete(context.Background(), groupId).BucketName(bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupGroupIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwMultisiteSyncPolicyGroupGroupIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketName** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwMultisiteSyncPolicyGroupGroupIdGet

> ApiRgwMultisiteSyncPolicyGroupGroupIdGet(ctx, groupId).BucketName(bucketName).Execute()

Get the sync policy group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	groupId := "groupId_example" // string | 
	bucketName := "bucketName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupGroupIdGet(context.Background(), groupId).BucketName(bucketName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupGroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwMultisiteSyncPolicyGroupGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucketName** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwMultisiteSyncPolicyGroupPost

> ApiRgwMultisiteSyncPolicyGroupPost(ctx).ApiRgwMultisiteSyncPolicyGroupPutRequest(apiRgwMultisiteSyncPolicyGroupPutRequest).Execute()

Create the sync policy group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	apiRgwMultisiteSyncPolicyGroupPutRequest := *openapiclient.NewApiRgwMultisiteSyncPolicyGroupPutRequest("GroupId_example", "Status_example") // ApiRgwMultisiteSyncPolicyGroupPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupPost(context.Background()).ApiRgwMultisiteSyncPolicyGroupPutRequest(apiRgwMultisiteSyncPolicyGroupPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwMultisiteSyncPolicyGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwMultisiteSyncPolicyGroupPutRequest** | [**ApiRgwMultisiteSyncPolicyGroupPutRequest**](ApiRgwMultisiteSyncPolicyGroupPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwMultisiteSyncPolicyGroupPut

> ApiRgwMultisiteSyncPolicyGroupPut(ctx).ApiRgwMultisiteSyncPolicyGroupPutRequest(apiRgwMultisiteSyncPolicyGroupPutRequest).Execute()

Update the sync policy group

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	apiRgwMultisiteSyncPolicyGroupPutRequest := *openapiclient.NewApiRgwMultisiteSyncPolicyGroupPutRequest("GroupId_example", "Status_example") // ApiRgwMultisiteSyncPolicyGroupPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupPut(context.Background()).ApiRgwMultisiteSyncPolicyGroupPutRequest(apiRgwMultisiteSyncPolicyGroupPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwMultisiteSyncPolicyGroupPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwMultisiteSyncPolicyGroupPutRequest** | [**ApiRgwMultisiteSyncPolicyGroupPutRequest**](ApiRgwMultisiteSyncPolicyGroupPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwMultisiteSyncStatusGet

> ApiRgwMultisiteSyncStatusGet(ctx).DaemonName(daemonName).Execute()

Get the sync status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncStatusGet(context.Background()).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwMultisiteAPI.ApiRgwMultisiteSyncStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwMultisiteSyncStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **daemonName** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

