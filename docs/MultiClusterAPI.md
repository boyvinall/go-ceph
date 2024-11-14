# \MultiClusterAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiMultiClusterAuthPost**](MultiClusterAPI.md#ApiMultiClusterAuthPost) | **Post** /api/multi-cluster/auth | Authenticate to a remote cluster
[**ApiMultiClusterCheckTokenStatusGet**](MultiClusterAPI.md#ApiMultiClusterCheckTokenStatusGet) | **Get** /api/multi-cluster/check_token_status | 
[**ApiMultiClusterDeleteClusterClusterNameClusterUserDelete**](MultiClusterAPI.md#ApiMultiClusterDeleteClusterClusterNameClusterUserDelete) | **Delete** /api/multi-cluster/delete_cluster/{cluster_name}/{cluster_user} | 
[**ApiMultiClusterEditClusterPut**](MultiClusterAPI.md#ApiMultiClusterEditClusterPut) | **Put** /api/multi-cluster/edit_cluster | 
[**ApiMultiClusterGetConfigGet**](MultiClusterAPI.md#ApiMultiClusterGetConfigGet) | **Get** /api/multi-cluster/get_config | 
[**ApiMultiClusterGetPrometheusApiUrlGet**](MultiClusterAPI.md#ApiMultiClusterGetPrometheusApiUrlGet) | **Get** /api/multi-cluster/get_prometheus_api_url | 
[**ApiMultiClusterReconnectClusterPut**](MultiClusterAPI.md#ApiMultiClusterReconnectClusterPut) | **Put** /api/multi-cluster/reconnect_cluster | 
[**ApiMultiClusterSecurityConfigGet**](MultiClusterAPI.md#ApiMultiClusterSecurityConfigGet) | **Get** /api/multi-cluster/security_config | 
[**ApiMultiClusterSetConfigPut**](MultiClusterAPI.md#ApiMultiClusterSetConfigPut) | **Put** /api/multi-cluster/set_config | 



## ApiMultiClusterAuthPost

> ApiMultiClusterAuthPost(ctx).ApiMultiClusterAuthPostRequest(apiMultiClusterAuthPostRequest).Execute()

Authenticate to a remote cluster

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
	apiMultiClusterAuthPostRequest := *openapiclient.NewApiMultiClusterAuthPostRequest("ClusterAlias_example", "Url_example", "Username_example") // ApiMultiClusterAuthPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiClusterAPI.ApiMultiClusterAuthPost(context.Background()).ApiMultiClusterAuthPostRequest(apiMultiClusterAuthPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ApiMultiClusterAuthPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiClusterAuthPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiMultiClusterAuthPostRequest** | [**ApiMultiClusterAuthPostRequest**](ApiMultiClusterAuthPostRequest.md) |  | 

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


## ApiMultiClusterCheckTokenStatusGet

> ApiMultiClusterCheckTokenStatusGet(ctx).ClustersTokenMap(clustersTokenMap).Execute()



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
	clustersTokenMap := "clustersTokenMap_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiClusterAPI.ApiMultiClusterCheckTokenStatusGet(context.Background()).ClustersTokenMap(clustersTokenMap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ApiMultiClusterCheckTokenStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiClusterCheckTokenStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clustersTokenMap** | **string** |  | 

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


## ApiMultiClusterDeleteClusterClusterNameClusterUserDelete

> ApiMultiClusterDeleteClusterClusterNameClusterUserDelete(ctx, clusterName, clusterUser).Execute()



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
	clusterName := "clusterName_example" // string | 
	clusterUser := "clusterUser_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiClusterAPI.ApiMultiClusterDeleteClusterClusterNameClusterUserDelete(context.Background(), clusterName, clusterUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ApiMultiClusterDeleteClusterClusterNameClusterUserDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** |  | 
**clusterUser** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiClusterDeleteClusterClusterNameClusterUserDeleteRequest struct via the builder pattern


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


## ApiMultiClusterEditClusterPut

> ApiMultiClusterEditClusterPut(ctx).ApiMultiClusterEditClusterPutRequest(apiMultiClusterEditClusterPutRequest).Execute()



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
	apiMultiClusterEditClusterPutRequest := *openapiclient.NewApiMultiClusterEditClusterPutRequest("ClusterAlias_example", "Name_example", "Url_example", "Username_example") // ApiMultiClusterEditClusterPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiClusterAPI.ApiMultiClusterEditClusterPut(context.Background()).ApiMultiClusterEditClusterPutRequest(apiMultiClusterEditClusterPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ApiMultiClusterEditClusterPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiClusterEditClusterPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiMultiClusterEditClusterPutRequest** | [**ApiMultiClusterEditClusterPutRequest**](ApiMultiClusterEditClusterPutRequest.md) |  | 

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


## ApiMultiClusterGetConfigGet

> ApiMultiClusterGetConfigGet(ctx).Execute()



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
	r, err := apiClient.MultiClusterAPI.ApiMultiClusterGetConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ApiMultiClusterGetConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiClusterGetConfigGetRequest struct via the builder pattern


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


## ApiMultiClusterGetPrometheusApiUrlGet

> ApiMultiClusterGetPrometheusApiUrlGet(ctx).Execute()



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
	r, err := apiClient.MultiClusterAPI.ApiMultiClusterGetPrometheusApiUrlGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ApiMultiClusterGetPrometheusApiUrlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiClusterGetPrometheusApiUrlGetRequest struct via the builder pattern


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


## ApiMultiClusterReconnectClusterPut

> ApiMultiClusterReconnectClusterPut(ctx).ApiMultiClusterReconnectClusterPutRequest(apiMultiClusterReconnectClusterPutRequest).Execute()



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
	apiMultiClusterReconnectClusterPutRequest := *openapiclient.NewApiMultiClusterReconnectClusterPutRequest("Url_example") // ApiMultiClusterReconnectClusterPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiClusterAPI.ApiMultiClusterReconnectClusterPut(context.Background()).ApiMultiClusterReconnectClusterPutRequest(apiMultiClusterReconnectClusterPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ApiMultiClusterReconnectClusterPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiClusterReconnectClusterPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiMultiClusterReconnectClusterPutRequest** | [**ApiMultiClusterReconnectClusterPutRequest**](ApiMultiClusterReconnectClusterPutRequest.md) |  | 

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


## ApiMultiClusterSecurityConfigGet

> ApiMultiClusterSecurityConfigGet(ctx).Execute()



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
	r, err := apiClient.MultiClusterAPI.ApiMultiClusterSecurityConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ApiMultiClusterSecurityConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiClusterSecurityConfigGetRequest struct via the builder pattern


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


## ApiMultiClusterSetConfigPut

> ApiMultiClusterSetConfigPut(ctx).ApiMgrModuleModuleNamePutRequest(apiMgrModuleModuleNamePutRequest).Execute()



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
	apiMgrModuleModuleNamePutRequest := *openapiclient.NewApiMgrModuleModuleNamePutRequest("Config_example") // ApiMgrModuleModuleNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultiClusterAPI.ApiMultiClusterSetConfigPut(context.Background()).ApiMgrModuleModuleNamePutRequest(apiMgrModuleModuleNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiClusterAPI.ApiMultiClusterSetConfigPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiMultiClusterSetConfigPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiMgrModuleModuleNamePutRequest** | [**ApiMgrModuleModuleNamePutRequest**](ApiMgrModuleModuleNamePutRequest.md) |  | 

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

