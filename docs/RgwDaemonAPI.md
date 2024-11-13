# \RgwDaemonAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRgwDaemonGet**](RgwDaemonAPI.md#ApiRgwDaemonGet) | **Get** /api/rgw/daemon | Display RGW Daemons
[**ApiRgwDaemonSetMultisiteConfigPut**](RgwDaemonAPI.md#ApiRgwDaemonSetMultisiteConfigPut) | **Put** /api/rgw/daemon/set_multisite_config | 
[**ApiRgwDaemonSvcIdGet**](RgwDaemonAPI.md#ApiRgwDaemonSvcIdGet) | **Get** /api/rgw/daemon/{svc_id} | 



## ApiRgwDaemonGet

> []ApiRgwDaemonGet200ResponseInner ApiRgwDaemonGet(ctx).Execute()

Display RGW Daemons

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RgwDaemonAPI.ApiRgwDaemonGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwDaemonAPI.ApiRgwDaemonGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiRgwDaemonGet`: []ApiRgwDaemonGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RgwDaemonAPI.ApiRgwDaemonGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwDaemonGetRequest struct via the builder pattern


### Return type

[**[]ApiRgwDaemonGet200ResponseInner**](ApiRgwDaemonGet200ResponseInner.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwDaemonSetMultisiteConfigPut

> ApiRgwDaemonSetMultisiteConfigPut(ctx).ApiRgwDaemonSetMultisiteConfigPutRequest(apiRgwDaemonSetMultisiteConfigPutRequest).Execute()



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
	apiRgwDaemonSetMultisiteConfigPutRequest := *openapiclient.NewApiRgwDaemonSetMultisiteConfigPutRequest() // ApiRgwDaemonSetMultisiteConfigPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwDaemonAPI.ApiRgwDaemonSetMultisiteConfigPut(context.Background()).ApiRgwDaemonSetMultisiteConfigPutRequest(apiRgwDaemonSetMultisiteConfigPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwDaemonAPI.ApiRgwDaemonSetMultisiteConfigPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwDaemonSetMultisiteConfigPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwDaemonSetMultisiteConfigPutRequest** | [**ApiRgwDaemonSetMultisiteConfigPutRequest**](ApiRgwDaemonSetMultisiteConfigPutRequest.md) |  | 

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


## ApiRgwDaemonSvcIdGet

> ApiRgwDaemonSvcIdGet(ctx, svcId).Execute()



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
	svcId := "svcId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwDaemonAPI.ApiRgwDaemonSvcIdGet(context.Background(), svcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwDaemonAPI.ApiRgwDaemonSvcIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwDaemonSvcIdGetRequest struct via the builder pattern


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

