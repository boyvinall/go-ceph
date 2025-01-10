# \PoolAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPoolGet**](PoolAPI.md#ApiPoolGet) | **Get** /api/pool | Display Pool List
[**ApiPoolPoolNameConfigurationGet**](PoolAPI.md#ApiPoolPoolNameConfigurationGet) | **Get** /api/pool/{pool_name}/configuration | 
[**ApiPoolPoolNameDelete**](PoolAPI.md#ApiPoolPoolNameDelete) | **Delete** /api/pool/{pool_name} | 
[**ApiPoolPoolNameGet**](PoolAPI.md#ApiPoolPoolNameGet) | **Get** /api/pool/{pool_name} | 
[**ApiPoolPoolNamePut**](PoolAPI.md#ApiPoolPoolNamePut) | **Put** /api/pool/{pool_name} | 
[**ApiPoolPost**](PoolAPI.md#ApiPoolPost) | **Post** /api/pool | 



## ApiPoolGet

> []ApiPoolGet200ResponseInner ApiPoolGet(ctx).Attrs(attrs).Stats(stats).Execute()

Display Pool List

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
	attrs := "attrs_example" // string | Pool Attributes (optional)
	stats := true // bool | Pool Stats (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolAPI.ApiPoolGet(context.Background()).Attrs(attrs).Stats(stats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ApiPoolGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPoolGet`: []ApiPoolGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PoolAPI.ApiPoolGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPoolGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attrs** | **string** | Pool Attributes | 
 **stats** | **bool** | Pool Stats | 

### Return type

[**[]ApiPoolGet200ResponseInner**](ApiPoolGet200ResponseInner.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPoolPoolNameConfigurationGet

> ApiPoolPoolNameConfigurationGet(ctx, poolName).Execute()



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
	r, err := apiClient.PoolAPI.ApiPoolPoolNameConfigurationGet(context.Background(), poolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ApiPoolPoolNameConfigurationGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiPoolPoolNameConfigurationGetRequest struct via the builder pattern


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


## ApiPoolPoolNameDelete

> ApiPoolPoolNameDelete(ctx, poolName).Execute()



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
	r, err := apiClient.PoolAPI.ApiPoolPoolNameDelete(context.Background(), poolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ApiPoolPoolNameDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiPoolPoolNameDeleteRequest struct via the builder pattern


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


## ApiPoolPoolNameGet

> ApiPoolPoolNameGet(ctx, poolName).Attrs(attrs).Stats(stats).Execute()



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
	attrs := "attrs_example" // string |  (optional)
	stats := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoolAPI.ApiPoolPoolNameGet(context.Background(), poolName).Attrs(attrs).Stats(stats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ApiPoolPoolNameGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiPoolPoolNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attrs** | **string** |  | 
 **stats** | **bool** |  | 

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


## ApiPoolPoolNamePut

> ApiPoolPoolNamePut(ctx, poolName).ApiPoolPoolNamePutRequest(apiPoolPoolNamePutRequest).Execute()



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
	apiPoolPoolNamePutRequest := *openapiclient.NewApiPoolPoolNamePutRequest() // ApiPoolPoolNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoolAPI.ApiPoolPoolNamePut(context.Background(), poolName).ApiPoolPoolNamePutRequest(apiPoolPoolNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ApiPoolPoolNamePut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiPoolPoolNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiPoolPoolNamePutRequest** | [**ApiPoolPoolNamePutRequest**](ApiPoolPoolNamePutRequest.md) |  | 

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


## ApiPoolPost

> ApiPoolPost(ctx).ApiPoolPostRequest(apiPoolPostRequest).Execute()



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
	apiPoolPostRequest := *openapiclient.NewApiPoolPostRequest() // ApiPoolPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoolAPI.ApiPoolPost(context.Background()).ApiPoolPostRequest(apiPoolPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolAPI.ApiPoolPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPoolPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiPoolPostRequest** | [**ApiPoolPostRequest**](ApiPoolPostRequest.md) |  | 

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

