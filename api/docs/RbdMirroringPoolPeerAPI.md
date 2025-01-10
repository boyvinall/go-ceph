# \RbdMirroringPoolPeerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBlockMirroringPoolPoolNamePeerGet**](RbdMirroringPoolPeerAPI.md#ApiBlockMirroringPoolPoolNamePeerGet) | **Get** /api/block/mirroring/pool/{pool_name}/peer | 
[**ApiBlockMirroringPoolPoolNamePeerPeerUuidDelete**](RbdMirroringPoolPeerAPI.md#ApiBlockMirroringPoolPoolNamePeerPeerUuidDelete) | **Delete** /api/block/mirroring/pool/{pool_name}/peer/{peer_uuid} | 
[**ApiBlockMirroringPoolPoolNamePeerPeerUuidGet**](RbdMirroringPoolPeerAPI.md#ApiBlockMirroringPoolPoolNamePeerPeerUuidGet) | **Get** /api/block/mirroring/pool/{pool_name}/peer/{peer_uuid} | 
[**ApiBlockMirroringPoolPoolNamePeerPeerUuidPut**](RbdMirroringPoolPeerAPI.md#ApiBlockMirroringPoolPoolNamePeerPeerUuidPut) | **Put** /api/block/mirroring/pool/{pool_name}/peer/{peer_uuid} | 
[**ApiBlockMirroringPoolPoolNamePeerPost**](RbdMirroringPoolPeerAPI.md#ApiBlockMirroringPoolPoolNamePeerPost) | **Post** /api/block/mirroring/pool/{pool_name}/peer | 



## ApiBlockMirroringPoolPoolNamePeerGet

> ApiBlockMirroringPoolPoolNamePeerGet(ctx, poolName).Execute()



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
	poolName := "poolName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerGet(context.Background(), poolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockMirroringPoolPoolNamePeerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiBlockMirroringPoolPoolNamePeerPeerUuidDelete

> ApiBlockMirroringPoolPoolNamePeerPeerUuidDelete(ctx, poolName, peerUuid).Execute()



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
	poolName := "poolName_example" // string | 
	peerUuid := "peerUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPeerUuidDelete(context.Background(), poolName, peerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPeerUuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolName** | **string** |  | 
**peerUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockMirroringPoolPoolNamePeerPeerUuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ApiBlockMirroringPoolPoolNamePeerPeerUuidGet

> ApiBlockMirroringPoolPoolNamePeerPeerUuidGet(ctx, poolName, peerUuid).Execute()



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
	poolName := "poolName_example" // string | 
	peerUuid := "peerUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPeerUuidGet(context.Background(), poolName, peerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPeerUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolName** | **string** |  | 
**peerUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockMirroringPoolPoolNamePeerPeerUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ApiBlockMirroringPoolPoolNamePeerPeerUuidPut

> ApiBlockMirroringPoolPoolNamePeerPeerUuidPut(ctx, poolName, peerUuid).ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest(apiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest).Execute()



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
	poolName := "poolName_example" // string | 
	peerUuid := "peerUuid_example" // string | 
	apiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest := *openapiclient.NewApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest() // ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPeerUuidPut(context.Background(), poolName, peerUuid).ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest(apiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPeerUuidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolName** | **string** |  | 
**peerUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest** | [**ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest**](ApiBlockMirroringPoolPoolNamePeerPeerUuidPutRequest.md) |  | 

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


## ApiBlockMirroringPoolPoolNamePeerPost

> ApiBlockMirroringPoolPoolNamePeerPost(ctx, poolName).ApiBlockMirroringPoolPoolNamePeerPostRequest(apiBlockMirroringPoolPoolNamePeerPostRequest).Execute()



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
	poolName := "poolName_example" // string | 
	apiBlockMirroringPoolPoolNamePeerPostRequest := *openapiclient.NewApiBlockMirroringPoolPoolNamePeerPostRequest("ClientId_example", "ClusterName_example") // ApiBlockMirroringPoolPoolNamePeerPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPost(context.Background(), poolName).ApiBlockMirroringPoolPoolNamePeerPostRequest(apiBlockMirroringPoolPoolNamePeerPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockMirroringPoolPoolNamePeerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiBlockMirroringPoolPoolNamePeerPostRequest** | [**ApiBlockMirroringPoolPoolNamePeerPostRequest**](ApiBlockMirroringPoolPoolNamePeerPostRequest.md) |  | 

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

