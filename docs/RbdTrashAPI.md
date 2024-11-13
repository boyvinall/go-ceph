# \RbdTrashAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBlockImageTrashGet**](RbdTrashAPI.md#ApiBlockImageTrashGet) | **Get** /api/block/image/trash | Get RBD Trash Details by pool name
[**ApiBlockImageTrashImageIdSpecDelete**](RbdTrashAPI.md#ApiBlockImageTrashImageIdSpecDelete) | **Delete** /api/block/image/trash/{image_id_spec} | 
[**ApiBlockImageTrashImageIdSpecRestorePost**](RbdTrashAPI.md#ApiBlockImageTrashImageIdSpecRestorePost) | **Post** /api/block/image/trash/{image_id_spec}/restore | 
[**ApiBlockImageTrashPurgePost**](RbdTrashAPI.md#ApiBlockImageTrashPurgePost) | **Post** /api/block/image/trash/purge | 



## ApiBlockImageTrashGet

> []ApiBlockImageTrashGet200ResponseInner ApiBlockImageTrashGet(ctx).PoolName(poolName).Execute()

Get RBD Trash Details by pool name



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
	poolName := "poolName_example" // string | Name of the pool (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdTrashAPI.ApiBlockImageTrashGet(context.Background()).PoolName(poolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdTrashAPI.ApiBlockImageTrashGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiBlockImageTrashGet`: []ApiBlockImageTrashGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RbdTrashAPI.ApiBlockImageTrashGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageTrashGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolName** | **string** | Name of the pool | 

### Return type

[**[]ApiBlockImageTrashGet200ResponseInner**](ApiBlockImageTrashGet200ResponseInner.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBlockImageTrashImageIdSpecDelete

> ApiBlockImageTrashImageIdSpecDelete(ctx, imageIdSpec).Force(force).Execute()





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
	imageIdSpec := "imageIdSpec_example" // string | 
	force := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdTrashAPI.ApiBlockImageTrashImageIdSpecDelete(context.Background(), imageIdSpec).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdTrashAPI.ApiBlockImageTrashImageIdSpecDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageIdSpec** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageTrashImageIdSpecDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | 

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


## ApiBlockImageTrashImageIdSpecRestorePost

> ApiBlockImageTrashImageIdSpecRestorePost(ctx, imageIdSpec).ApiBlockImageTrashImageIdSpecRestorePostRequest(apiBlockImageTrashImageIdSpecRestorePostRequest).Execute()





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
	imageIdSpec := "imageIdSpec_example" // string | 
	apiBlockImageTrashImageIdSpecRestorePostRequest := *openapiclient.NewApiBlockImageTrashImageIdSpecRestorePostRequest("NewImageName_example") // ApiBlockImageTrashImageIdSpecRestorePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdTrashAPI.ApiBlockImageTrashImageIdSpecRestorePost(context.Background(), imageIdSpec).ApiBlockImageTrashImageIdSpecRestorePostRequest(apiBlockImageTrashImageIdSpecRestorePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdTrashAPI.ApiBlockImageTrashImageIdSpecRestorePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageIdSpec** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageTrashImageIdSpecRestorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiBlockImageTrashImageIdSpecRestorePostRequest** | [**ApiBlockImageTrashImageIdSpecRestorePostRequest**](ApiBlockImageTrashImageIdSpecRestorePostRequest.md) |  | 

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


## ApiBlockImageTrashPurgePost

> ApiBlockImageTrashPurgePost(ctx).PoolName(poolName).ApiBlockImageTrashPurgePostRequest(apiBlockImageTrashPurgePostRequest).Execute()





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
	poolName := "poolName_example" // string |  (optional)
	apiBlockImageTrashPurgePostRequest := *openapiclient.NewApiBlockImageTrashPurgePostRequest() // ApiBlockImageTrashPurgePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdTrashAPI.ApiBlockImageTrashPurgePost(context.Background()).PoolName(poolName).ApiBlockImageTrashPurgePostRequest(apiBlockImageTrashPurgePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdTrashAPI.ApiBlockImageTrashPurgePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageTrashPurgePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolName** | **string** |  | 
 **apiBlockImageTrashPurgePostRequest** | [**ApiBlockImageTrashPurgePostRequest**](ApiBlockImageTrashPurgePostRequest.md) |  | 

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

