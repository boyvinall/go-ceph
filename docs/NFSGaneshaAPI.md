# \NFSGaneshaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiNfsGaneshaClusterGet**](NFSGaneshaAPI.md#ApiNfsGaneshaClusterGet) | **Get** /api/nfs-ganesha/cluster | 
[**ApiNfsGaneshaExportClusterIdExportIdDelete**](NFSGaneshaAPI.md#ApiNfsGaneshaExportClusterIdExportIdDelete) | **Delete** /api/nfs-ganesha/export/{cluster_id}/{export_id} | Deletes an NFS-Ganesha export
[**ApiNfsGaneshaExportClusterIdExportIdGet**](NFSGaneshaAPI.md#ApiNfsGaneshaExportClusterIdExportIdGet) | **Get** /api/nfs-ganesha/export/{cluster_id}/{export_id} | Get an NFS-Ganesha export
[**ApiNfsGaneshaExportClusterIdExportIdPut**](NFSGaneshaAPI.md#ApiNfsGaneshaExportClusterIdExportIdPut) | **Put** /api/nfs-ganesha/export/{cluster_id}/{export_id} | Updates an NFS-Ganesha export
[**ApiNfsGaneshaExportGet**](NFSGaneshaAPI.md#ApiNfsGaneshaExportGet) | **Get** /api/nfs-ganesha/export | List all NFS-Ganesha exports
[**ApiNfsGaneshaExportPost**](NFSGaneshaAPI.md#ApiNfsGaneshaExportPost) | **Post** /api/nfs-ganesha/export | Creates a new NFS-Ganesha export



## ApiNfsGaneshaClusterGet

> ApiNfsGaneshaClusterGet(ctx).Execute()



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
	r, err := apiClient.NFSGaneshaAPI.ApiNfsGaneshaClusterGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFSGaneshaAPI.ApiNfsGaneshaClusterGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiNfsGaneshaClusterGetRequest struct via the builder pattern


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


## ApiNfsGaneshaExportClusterIdExportIdDelete

> ApiNfsGaneshaExportClusterIdExportIdDelete(ctx, clusterId, exportId).Execute()

Deletes an NFS-Ganesha export

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
	clusterId := "clusterId_example" // string | Cluster identifier
	exportId := int32(56) // int32 | Export ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NFSGaneshaAPI.ApiNfsGaneshaExportClusterIdExportIdDelete(context.Background(), clusterId, exportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFSGaneshaAPI.ApiNfsGaneshaExportClusterIdExportIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster identifier | 
**exportId** | **int32** | Export ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNfsGaneshaExportClusterIdExportIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiNfsGaneshaExportClusterIdExportIdGet

> ApiNfsGaneshaExportPost201Response ApiNfsGaneshaExportClusterIdExportIdGet(ctx, clusterId, exportId).Execute()

Get an NFS-Ganesha export

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
	clusterId := "clusterId_example" // string | Cluster identifier
	exportId := "exportId_example" // string | Export ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NFSGaneshaAPI.ApiNfsGaneshaExportClusterIdExportIdGet(context.Background(), clusterId, exportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFSGaneshaAPI.ApiNfsGaneshaExportClusterIdExportIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiNfsGaneshaExportClusterIdExportIdGet`: ApiNfsGaneshaExportPost201Response
	fmt.Fprintf(os.Stdout, "Response from `NFSGaneshaAPI.ApiNfsGaneshaExportClusterIdExportIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster identifier | 
**exportId** | **string** | Export ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNfsGaneshaExportClusterIdExportIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiNfsGaneshaExportPost201Response**](ApiNfsGaneshaExportPost201Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiNfsGaneshaExportClusterIdExportIdPut

> ApiNfsGaneshaExportPost201Response ApiNfsGaneshaExportClusterIdExportIdPut(ctx, clusterId, exportId).ApiNfsGaneshaExportClusterIdExportIdPutRequest(apiNfsGaneshaExportClusterIdExportIdPutRequest).Execute()

Updates an NFS-Ganesha export

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
	clusterId := "clusterId_example" // string | Cluster identifier
	exportId := int32(56) // int32 | Export ID
	apiNfsGaneshaExportClusterIdExportIdPutRequest := *openapiclient.NewApiNfsGaneshaExportClusterIdExportIdPutRequest("AccessType_example", []openapiclient.ApiNfsGaneshaExportGet200ResponseInnerClientsInner{*openapiclient.NewApiNfsGaneshaExportGet200ResponseInnerClientsInner("AccessType_example", []string{"Addresses_example"}, "Squash_example")}, *openapiclient.NewApiNfsGaneshaExportPostRequestFsal("Name_example"), "Path_example", []int32{int32(123)}, "Pseudo_example", "SecurityLabel_example", "Squash_example", []string{"Transports_example"}) // ApiNfsGaneshaExportClusterIdExportIdPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NFSGaneshaAPI.ApiNfsGaneshaExportClusterIdExportIdPut(context.Background(), clusterId, exportId).ApiNfsGaneshaExportClusterIdExportIdPutRequest(apiNfsGaneshaExportClusterIdExportIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFSGaneshaAPI.ApiNfsGaneshaExportClusterIdExportIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiNfsGaneshaExportClusterIdExportIdPut`: ApiNfsGaneshaExportPost201Response
	fmt.Fprintf(os.Stdout, "Response from `NFSGaneshaAPI.ApiNfsGaneshaExportClusterIdExportIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster identifier | 
**exportId** | **int32** | Export ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNfsGaneshaExportClusterIdExportIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiNfsGaneshaExportClusterIdExportIdPutRequest** | [**ApiNfsGaneshaExportClusterIdExportIdPutRequest**](ApiNfsGaneshaExportClusterIdExportIdPutRequest.md) |  | 

### Return type

[**ApiNfsGaneshaExportPost201Response**](ApiNfsGaneshaExportPost201Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiNfsGaneshaExportGet

> []ApiNfsGaneshaExportGet200ResponseInner ApiNfsGaneshaExportGet(ctx).Execute()

List all NFS-Ganesha exports

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
	resp, r, err := apiClient.NFSGaneshaAPI.ApiNfsGaneshaExportGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFSGaneshaAPI.ApiNfsGaneshaExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiNfsGaneshaExportGet`: []ApiNfsGaneshaExportGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NFSGaneshaAPI.ApiNfsGaneshaExportGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiNfsGaneshaExportGetRequest struct via the builder pattern


### Return type

[**[]ApiNfsGaneshaExportGet200ResponseInner**](ApiNfsGaneshaExportGet200ResponseInner.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiNfsGaneshaExportPost

> ApiNfsGaneshaExportPost201Response ApiNfsGaneshaExportPost(ctx).ApiNfsGaneshaExportPostRequest(apiNfsGaneshaExportPostRequest).Execute()

Creates a new NFS-Ganesha export

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
	apiNfsGaneshaExportPostRequest := *openapiclient.NewApiNfsGaneshaExportPostRequest("AccessType_example", []openapiclient.ApiNfsGaneshaExportGet200ResponseInnerClientsInner{*openapiclient.NewApiNfsGaneshaExportGet200ResponseInnerClientsInner("AccessType_example", []string{"Addresses_example"}, "Squash_example")}, "ClusterId_example", *openapiclient.NewApiNfsGaneshaExportPostRequestFsal("Name_example"), "Path_example", []int32{int32(123)}, "Pseudo_example", "SecurityLabel_example", "Squash_example", []string{"Transports_example"}) // ApiNfsGaneshaExportPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NFSGaneshaAPI.ApiNfsGaneshaExportPost(context.Background()).ApiNfsGaneshaExportPostRequest(apiNfsGaneshaExportPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFSGaneshaAPI.ApiNfsGaneshaExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiNfsGaneshaExportPost`: ApiNfsGaneshaExportPost201Response
	fmt.Fprintf(os.Stdout, "Response from `NFSGaneshaAPI.ApiNfsGaneshaExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiNfsGaneshaExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiNfsGaneshaExportPostRequest** | [**ApiNfsGaneshaExportPostRequest**](ApiNfsGaneshaExportPostRequest.md) |  | 

### Return type

[**ApiNfsGaneshaExportPost201Response**](ApiNfsGaneshaExportPost201Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

