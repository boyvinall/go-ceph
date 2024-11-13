# \DaemonAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDaemonDaemonNamePut**](DaemonAPI.md#ApiDaemonDaemonNamePut) | **Put** /api/daemon/{daemon_name} | 
[**ApiDaemonGet**](DaemonAPI.md#ApiDaemonGet) | **Get** /api/daemon | 



## ApiDaemonDaemonNamePut

> ApiDaemonDaemonNamePut(ctx, daemonName).ApiDaemonDaemonNamePutRequest(apiDaemonDaemonNamePutRequest).Execute()



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
	daemonName := "daemonName_example" // string | 
	apiDaemonDaemonNamePutRequest := *openapiclient.NewApiDaemonDaemonNamePutRequest() // ApiDaemonDaemonNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaemonAPI.ApiDaemonDaemonNamePut(context.Background(), daemonName).ApiDaemonDaemonNamePutRequest(apiDaemonDaemonNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaemonAPI.ApiDaemonDaemonNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**daemonName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDaemonDaemonNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiDaemonDaemonNamePutRequest** | [**ApiDaemonDaemonNamePutRequest**](ApiDaemonDaemonNamePutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v0.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDaemonGet

> ApiDaemonGet(ctx).DaemonTypes(daemonTypes).Execute()





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
	daemonTypes := "daemonTypes_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaemonAPI.ApiDaemonGet(context.Background()).DaemonTypes(daemonTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaemonAPI.ApiDaemonGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDaemonGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **daemonTypes** | **string** |  | 

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

