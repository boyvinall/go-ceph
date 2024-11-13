# \ClusterAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiClusterGet**](ClusterAPI.md#ApiClusterGet) | **Get** /api/cluster | Get the cluster status
[**ApiClusterPut**](ClusterAPI.md#ApiClusterPut) | **Put** /api/cluster | Update the cluster status
[**ApiClusterUserExportPost**](ClusterAPI.md#ApiClusterUserExportPost) | **Post** /api/cluster/user/export | Export Ceph Users
[**ApiClusterUserGet**](ClusterAPI.md#ApiClusterUserGet) | **Get** /api/cluster/user | Get list of ceph users
[**ApiClusterUserPost**](ClusterAPI.md#ApiClusterUserPost) | **Post** /api/cluster/user | Create Ceph User
[**ApiClusterUserPut**](ClusterAPI.md#ApiClusterUserPut) | **Put** /api/cluster/user | Edit Ceph User Capabilities
[**ApiClusterUserUserEntityDelete**](ClusterAPI.md#ApiClusterUserUserEntityDelete) | **Delete** /api/cluster/user/{user_entity} | Delete Ceph User



## ApiClusterGet

> ApiClusterGet(ctx).Execute()

Get the cluster status

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
	r, err := apiClient.ClusterAPI.ApiClusterGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiClusterGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v0.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiClusterPut

> ApiClusterPut(ctx).ApiClusterPutRequest(apiClusterPutRequest).Execute()

Update the cluster status

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
	apiClusterPutRequest := *openapiclient.NewApiClusterPutRequest("Status_example") // ApiClusterPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClusterAPI.ApiClusterPut(context.Background()).ApiClusterPutRequest(apiClusterPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiClusterPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiClusterPutRequest** | [**ApiClusterPutRequest**](ApiClusterPutRequest.md) |  | 

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


## ApiClusterUserExportPost

> ApiClusterUserExportPost(ctx).ApiClusterUserExportPostRequest(apiClusterUserExportPostRequest).Execute()

Export Ceph Users

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
	apiClusterUserExportPostRequest := *openapiclient.NewApiClusterUserExportPostRequest("Entities_example") // ApiClusterUserExportPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClusterAPI.ApiClusterUserExportPost(context.Background()).ApiClusterUserExportPostRequest(apiClusterUserExportPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiClusterUserExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUserExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiClusterUserExportPostRequest** | [**ApiClusterUserExportPostRequest**](ApiClusterUserExportPostRequest.md) |  | 

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


## ApiClusterUserGet

> ApiClusterUserGet(ctx).Execute()

Get list of ceph users



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
	r, err := apiClient.ClusterAPI.ApiClusterUserGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiClusterUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUserGetRequest struct via the builder pattern


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


## ApiClusterUserPost

> ApiClusterUserPost(ctx).ApiClusterUserPostRequest(apiClusterUserPostRequest).Execute()

Create Ceph User



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
	apiClusterUserPostRequest := *openapiclient.NewApiClusterUserPostRequest() // ApiClusterUserPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClusterAPI.ApiClusterUserPost(context.Background()).ApiClusterUserPostRequest(apiClusterUserPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiClusterUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiClusterUserPostRequest** | [**ApiClusterUserPostRequest**](ApiClusterUserPostRequest.md) |  | 

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


## ApiClusterUserPut

> ApiClusterUserPut(ctx).ApiClusterUserPutRequest(apiClusterUserPutRequest).Execute()

Edit Ceph User Capabilities



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
	apiClusterUserPutRequest := *openapiclient.NewApiClusterUserPutRequest() // ApiClusterUserPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClusterAPI.ApiClusterUserPut(context.Background()).ApiClusterUserPutRequest(apiClusterUserPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiClusterUserPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUserPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiClusterUserPutRequest** | [**ApiClusterUserPutRequest**](ApiClusterUserPutRequest.md) |  | 

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


## ApiClusterUserUserEntityDelete

> ApiClusterUserUserEntityDelete(ctx, userEntity).Execute()

Delete Ceph User



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
	userEntity := "userEntity_example" // string | Entity to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClusterAPI.ApiClusterUserUserEntityDelete(context.Background(), userEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiClusterUserUserEntityDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEntity** | **string** | Entity to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUserUserEntityDeleteRequest struct via the builder pattern


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

