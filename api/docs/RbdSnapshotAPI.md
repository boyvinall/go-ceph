# \RbdSnapshotAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBlockImageImageSpecSnapPost**](RbdSnapshotAPI.md#ApiBlockImageImageSpecSnapPost) | **Post** /api/block/image/{image_spec}/snap | 
[**ApiBlockImageImageSpecSnapSnapshotNameClonePost**](RbdSnapshotAPI.md#ApiBlockImageImageSpecSnapSnapshotNameClonePost) | **Post** /api/block/image/{image_spec}/snap/{snapshot_name}/clone | 
[**ApiBlockImageImageSpecSnapSnapshotNameDelete**](RbdSnapshotAPI.md#ApiBlockImageImageSpecSnapSnapshotNameDelete) | **Delete** /api/block/image/{image_spec}/snap/{snapshot_name} | 
[**ApiBlockImageImageSpecSnapSnapshotNamePut**](RbdSnapshotAPI.md#ApiBlockImageImageSpecSnapSnapshotNamePut) | **Put** /api/block/image/{image_spec}/snap/{snapshot_name} | 
[**ApiBlockImageImageSpecSnapSnapshotNameRollbackPost**](RbdSnapshotAPI.md#ApiBlockImageImageSpecSnapSnapshotNameRollbackPost) | **Post** /api/block/image/{image_spec}/snap/{snapshot_name}/rollback | 



## ApiBlockImageImageSpecSnapPost

> ApiBlockImageImageSpecSnapPost(ctx, imageSpec).ApiBlockImageImageSpecSnapPostRequest(apiBlockImageImageSpecSnapPostRequest).Execute()



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
	imageSpec := "imageSpec_example" // string | 
	apiBlockImageImageSpecSnapPostRequest := *openapiclient.NewApiBlockImageImageSpecSnapPostRequest("MirrorImageSnapshot_example", "SnapshotName_example") // ApiBlockImageImageSpecSnapPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdSnapshotAPI.ApiBlockImageImageSpecSnapPost(context.Background(), imageSpec).ApiBlockImageImageSpecSnapPostRequest(apiBlockImageImageSpecSnapPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdSnapshotAPI.ApiBlockImageImageSpecSnapPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecSnapPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiBlockImageImageSpecSnapPostRequest** | [**ApiBlockImageImageSpecSnapPostRequest**](ApiBlockImageImageSpecSnapPostRequest.md) |  | 

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


## ApiBlockImageImageSpecSnapSnapshotNameClonePost

> ApiBlockImageImageSpecSnapSnapshotNameClonePost(ctx, imageSpec, snapshotName).ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest(apiBlockImageImageSpecSnapSnapshotNameClonePostRequest).Execute()





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
	imageSpec := "imageSpec_example" // string | 
	snapshotName := "snapshotName_example" // string | 
	apiBlockImageImageSpecSnapSnapshotNameClonePostRequest := *openapiclient.NewApiBlockImageImageSpecSnapSnapshotNameClonePostRequest("ChildImageName_example", "ChildPoolName_example") // ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdSnapshotAPI.ApiBlockImageImageSpecSnapSnapshotNameClonePost(context.Background(), imageSpec, snapshotName).ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest(apiBlockImageImageSpecSnapSnapshotNameClonePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdSnapshotAPI.ApiBlockImageImageSpecSnapSnapshotNameClonePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** |  | 
**snapshotName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecSnapSnapshotNameClonePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiBlockImageImageSpecSnapSnapshotNameClonePostRequest** | [**ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest**](ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest.md) |  | 

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


## ApiBlockImageImageSpecSnapSnapshotNameDelete

> ApiBlockImageImageSpecSnapSnapshotNameDelete(ctx, imageSpec, snapshotName).Execute()



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
	imageSpec := "imageSpec_example" // string | 
	snapshotName := "snapshotName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdSnapshotAPI.ApiBlockImageImageSpecSnapSnapshotNameDelete(context.Background(), imageSpec, snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdSnapshotAPI.ApiBlockImageImageSpecSnapSnapshotNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** |  | 
**snapshotName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecSnapSnapshotNameDeleteRequest struct via the builder pattern


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


## ApiBlockImageImageSpecSnapSnapshotNamePut

> ApiBlockImageImageSpecSnapSnapshotNamePut(ctx, imageSpec, snapshotName).ApiBlockImageImageSpecSnapSnapshotNamePutRequest(apiBlockImageImageSpecSnapSnapshotNamePutRequest).Execute()



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
	imageSpec := "imageSpec_example" // string | 
	snapshotName := "snapshotName_example" // string | 
	apiBlockImageImageSpecSnapSnapshotNamePutRequest := *openapiclient.NewApiBlockImageImageSpecSnapSnapshotNamePutRequest() // ApiBlockImageImageSpecSnapSnapshotNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdSnapshotAPI.ApiBlockImageImageSpecSnapSnapshotNamePut(context.Background(), imageSpec, snapshotName).ApiBlockImageImageSpecSnapSnapshotNamePutRequest(apiBlockImageImageSpecSnapSnapshotNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdSnapshotAPI.ApiBlockImageImageSpecSnapSnapshotNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** |  | 
**snapshotName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecSnapSnapshotNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiBlockImageImageSpecSnapSnapshotNamePutRequest** | [**ApiBlockImageImageSpecSnapSnapshotNamePutRequest**](ApiBlockImageImageSpecSnapSnapshotNamePutRequest.md) |  | 

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


## ApiBlockImageImageSpecSnapSnapshotNameRollbackPost

> ApiBlockImageImageSpecSnapSnapshotNameRollbackPost(ctx, imageSpec, snapshotName).Execute()



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
	imageSpec := "imageSpec_example" // string | 
	snapshotName := "snapshotName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdSnapshotAPI.ApiBlockImageImageSpecSnapSnapshotNameRollbackPost(context.Background(), imageSpec, snapshotName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdSnapshotAPI.ApiBlockImageImageSpecSnapSnapshotNameRollbackPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** |  | 
**snapshotName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecSnapSnapshotNameRollbackPostRequest struct via the builder pattern


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

