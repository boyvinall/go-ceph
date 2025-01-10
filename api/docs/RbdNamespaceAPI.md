# \RbdNamespaceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBlockPoolPoolNameNamespaceGet**](RbdNamespaceAPI.md#ApiBlockPoolPoolNameNamespaceGet) | **Get** /api/block/pool/{pool_name}/namespace | 
[**ApiBlockPoolPoolNameNamespaceNamespaceDelete**](RbdNamespaceAPI.md#ApiBlockPoolPoolNameNamespaceNamespaceDelete) | **Delete** /api/block/pool/{pool_name}/namespace/{namespace} | 
[**ApiBlockPoolPoolNameNamespacePost**](RbdNamespaceAPI.md#ApiBlockPoolPoolNameNamespacePost) | **Post** /api/block/pool/{pool_name}/namespace | 



## ApiBlockPoolPoolNameNamespaceGet

> ApiBlockPoolPoolNameNamespaceGet(ctx, poolName).Execute()



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
	r, err := apiClient.RbdNamespaceAPI.ApiBlockPoolPoolNameNamespaceGet(context.Background(), poolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdNamespaceAPI.ApiBlockPoolPoolNameNamespaceGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiBlockPoolPoolNameNamespaceGetRequest struct via the builder pattern


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


## ApiBlockPoolPoolNameNamespaceNamespaceDelete

> ApiBlockPoolPoolNameNamespaceNamespaceDelete(ctx, poolName, namespace).Execute()



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
	namespace := "namespace_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdNamespaceAPI.ApiBlockPoolPoolNameNamespaceNamespaceDelete(context.Background(), poolName, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdNamespaceAPI.ApiBlockPoolPoolNameNamespaceNamespaceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolName** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockPoolPoolNameNamespaceNamespaceDeleteRequest struct via the builder pattern


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


## ApiBlockPoolPoolNameNamespacePost

> ApiBlockPoolPoolNameNamespacePost(ctx, poolName).ApiBlockPoolPoolNameNamespacePostRequest(apiBlockPoolPoolNameNamespacePostRequest).Execute()



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
	apiBlockPoolPoolNameNamespacePostRequest := *openapiclient.NewApiBlockPoolPoolNameNamespacePostRequest("Namespace_example") // ApiBlockPoolPoolNameNamespacePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdNamespaceAPI.ApiBlockPoolPoolNameNamespacePost(context.Background(), poolName).ApiBlockPoolPoolNameNamespacePostRequest(apiBlockPoolPoolNameNamespacePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdNamespaceAPI.ApiBlockPoolPoolNameNamespacePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiBlockPoolPoolNameNamespacePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiBlockPoolPoolNameNamespacePostRequest** | [**ApiBlockPoolPoolNameNamespacePostRequest**](ApiBlockPoolPoolNameNamespacePostRequest.md) |  | 

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

