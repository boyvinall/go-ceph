# \IscsiTargetAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiIscsiTargetGet**](IscsiTargetAPI.md#ApiIscsiTargetGet) | **Get** /api/iscsi/target | 
[**ApiIscsiTargetPost**](IscsiTargetAPI.md#ApiIscsiTargetPost) | **Post** /api/iscsi/target | 
[**ApiIscsiTargetTargetIqnDelete**](IscsiTargetAPI.md#ApiIscsiTargetTargetIqnDelete) | **Delete** /api/iscsi/target/{target_iqn} | 
[**ApiIscsiTargetTargetIqnGet**](IscsiTargetAPI.md#ApiIscsiTargetTargetIqnGet) | **Get** /api/iscsi/target/{target_iqn} | 
[**ApiIscsiTargetTargetIqnPut**](IscsiTargetAPI.md#ApiIscsiTargetTargetIqnPut) | **Put** /api/iscsi/target/{target_iqn} | 



## ApiIscsiTargetGet

> ApiIscsiTargetGet(ctx).Execute()



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
	r, err := apiClient.IscsiTargetAPI.ApiIscsiTargetGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IscsiTargetAPI.ApiIscsiTargetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiIscsiTargetGetRequest struct via the builder pattern


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


## ApiIscsiTargetPost

> ApiIscsiTargetPost(ctx).ApiIscsiTargetPostRequest(apiIscsiTargetPostRequest).Execute()



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
	apiIscsiTargetPostRequest := *openapiclient.NewApiIscsiTargetPostRequest() // ApiIscsiTargetPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IscsiTargetAPI.ApiIscsiTargetPost(context.Background()).ApiIscsiTargetPostRequest(apiIscsiTargetPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IscsiTargetAPI.ApiIscsiTargetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiIscsiTargetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiIscsiTargetPostRequest** | [**ApiIscsiTargetPostRequest**](ApiIscsiTargetPostRequest.md) |  | 

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


## ApiIscsiTargetTargetIqnDelete

> ApiIscsiTargetTargetIqnDelete(ctx, targetIqn).Execute()



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
	targetIqn := "targetIqn_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IscsiTargetAPI.ApiIscsiTargetTargetIqnDelete(context.Background(), targetIqn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IscsiTargetAPI.ApiIscsiTargetTargetIqnDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetIqn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiIscsiTargetTargetIqnDeleteRequest struct via the builder pattern


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


## ApiIscsiTargetTargetIqnGet

> ApiIscsiTargetTargetIqnGet(ctx, targetIqn).Execute()



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
	targetIqn := "targetIqn_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IscsiTargetAPI.ApiIscsiTargetTargetIqnGet(context.Background(), targetIqn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IscsiTargetAPI.ApiIscsiTargetTargetIqnGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetIqn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiIscsiTargetTargetIqnGetRequest struct via the builder pattern


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


## ApiIscsiTargetTargetIqnPut

> ApiIscsiTargetTargetIqnPut(ctx, targetIqn).ApiIscsiTargetTargetIqnPutRequest(apiIscsiTargetTargetIqnPutRequest).Execute()



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
	targetIqn := "targetIqn_example" // string | 
	apiIscsiTargetTargetIqnPutRequest := *openapiclient.NewApiIscsiTargetTargetIqnPutRequest() // ApiIscsiTargetTargetIqnPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IscsiTargetAPI.ApiIscsiTargetTargetIqnPut(context.Background(), targetIqn).ApiIscsiTargetTargetIqnPutRequest(apiIscsiTargetTargetIqnPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IscsiTargetAPI.ApiIscsiTargetTargetIqnPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetIqn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiIscsiTargetTargetIqnPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiIscsiTargetTargetIqnPutRequest** | [**ApiIscsiTargetTargetIqnPutRequest**](ApiIscsiTargetTargetIqnPutRequest.md) |  | 

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

