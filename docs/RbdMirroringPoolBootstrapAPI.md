# \RbdMirroringPoolBootstrapAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBlockMirroringPoolPoolNameBootstrapPeerPost**](RbdMirroringPoolBootstrapAPI.md#ApiBlockMirroringPoolPoolNameBootstrapPeerPost) | **Post** /api/block/mirroring/pool/{pool_name}/bootstrap/peer | 
[**ApiBlockMirroringPoolPoolNameBootstrapTokenPost**](RbdMirroringPoolBootstrapAPI.md#ApiBlockMirroringPoolPoolNameBootstrapTokenPost) | **Post** /api/block/mirroring/pool/{pool_name}/bootstrap/token | 



## ApiBlockMirroringPoolPoolNameBootstrapPeerPost

> ApiBlockMirroringPoolPoolNameBootstrapPeerPost(ctx, poolName).ApiBlockMirroringPoolPoolNameBootstrapPeerPostRequest(apiBlockMirroringPoolPoolNameBootstrapPeerPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {
	poolName := "poolName_example" // string | 
	apiBlockMirroringPoolPoolNameBootstrapPeerPostRequest := *openapiclient.NewApiBlockMirroringPoolPoolNameBootstrapPeerPostRequest("Direction_example", "Token_example") // ApiBlockMirroringPoolPoolNameBootstrapPeerPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdMirroringPoolBootstrapAPI.ApiBlockMirroringPoolPoolNameBootstrapPeerPost(context.Background(), poolName).ApiBlockMirroringPoolPoolNameBootstrapPeerPostRequest(apiBlockMirroringPoolPoolNameBootstrapPeerPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringPoolBootstrapAPI.ApiBlockMirroringPoolPoolNameBootstrapPeerPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiBlockMirroringPoolPoolNameBootstrapPeerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiBlockMirroringPoolPoolNameBootstrapPeerPostRequest** | [**ApiBlockMirroringPoolPoolNameBootstrapPeerPostRequest**](ApiBlockMirroringPoolPoolNameBootstrapPeerPostRequest.md) |  | 

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


## ApiBlockMirroringPoolPoolNameBootstrapTokenPost

> ApiBlockMirroringPoolPoolNameBootstrapTokenPost(ctx, poolName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {
	poolName := "poolName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdMirroringPoolBootstrapAPI.ApiBlockMirroringPoolPoolNameBootstrapTokenPost(context.Background(), poolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringPoolBootstrapAPI.ApiBlockMirroringPoolPoolNameBootstrapTokenPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiBlockMirroringPoolPoolNameBootstrapTokenPostRequest struct via the builder pattern


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

