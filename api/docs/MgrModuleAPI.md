# \MgrModuleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiMgrModuleGet**](MgrModuleAPI.md#ApiMgrModuleGet) | **Get** /api/mgr/module | List Mgr modules
[**ApiMgrModuleModuleNameDisablePost**](MgrModuleAPI.md#ApiMgrModuleModuleNameDisablePost) | **Post** /api/mgr/module/{module_name}/disable | 
[**ApiMgrModuleModuleNameEnablePost**](MgrModuleAPI.md#ApiMgrModuleModuleNameEnablePost) | **Post** /api/mgr/module/{module_name}/enable | 
[**ApiMgrModuleModuleNameGet**](MgrModuleAPI.md#ApiMgrModuleModuleNameGet) | **Get** /api/mgr/module/{module_name} | 
[**ApiMgrModuleModuleNameOptionsGet**](MgrModuleAPI.md#ApiMgrModuleModuleNameOptionsGet) | **Get** /api/mgr/module/{module_name}/options | 
[**ApiMgrModuleModuleNamePut**](MgrModuleAPI.md#ApiMgrModuleModuleNamePut) | **Put** /api/mgr/module/{module_name} | 



## ApiMgrModuleGet

> []ApiMgrModuleGet200ResponseInner ApiMgrModuleGet(ctx).Execute()

List Mgr modules



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
	resp, r, err := apiClient.MgrModuleAPI.ApiMgrModuleGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MgrModuleAPI.ApiMgrModuleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiMgrModuleGet`: []ApiMgrModuleGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MgrModuleAPI.ApiMgrModuleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiMgrModuleGetRequest struct via the builder pattern


### Return type

[**[]ApiMgrModuleGet200ResponseInner**](ApiMgrModuleGet200ResponseInner.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiMgrModuleModuleNameDisablePost

> ApiMgrModuleModuleNameDisablePost(ctx, moduleName).Execute()





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
	moduleName := "moduleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MgrModuleAPI.ApiMgrModuleModuleNameDisablePost(context.Background(), moduleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MgrModuleAPI.ApiMgrModuleModuleNameDisablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMgrModuleModuleNameDisablePostRequest struct via the builder pattern


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


## ApiMgrModuleModuleNameEnablePost

> ApiMgrModuleModuleNameEnablePost(ctx, moduleName).Execute()





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
	moduleName := "moduleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MgrModuleAPI.ApiMgrModuleModuleNameEnablePost(context.Background(), moduleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MgrModuleAPI.ApiMgrModuleModuleNameEnablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMgrModuleModuleNameEnablePostRequest struct via the builder pattern


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


## ApiMgrModuleModuleNameGet

> ApiMgrModuleModuleNameGet(ctx, moduleName).Execute()





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
	moduleName := "moduleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MgrModuleAPI.ApiMgrModuleModuleNameGet(context.Background(), moduleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MgrModuleAPI.ApiMgrModuleModuleNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMgrModuleModuleNameGetRequest struct via the builder pattern


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


## ApiMgrModuleModuleNameOptionsGet

> ApiMgrModuleModuleNameOptionsGet(ctx, moduleName).Execute()





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
	moduleName := "moduleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MgrModuleAPI.ApiMgrModuleModuleNameOptionsGet(context.Background(), moduleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MgrModuleAPI.ApiMgrModuleModuleNameOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMgrModuleModuleNameOptionsGetRequest struct via the builder pattern


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


## ApiMgrModuleModuleNamePut

> ApiMgrModuleModuleNamePut(ctx, moduleName).ApiMgrModuleModuleNamePutRequest(apiMgrModuleModuleNamePutRequest).Execute()





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
	moduleName := "moduleName_example" // string | 
	apiMgrModuleModuleNamePutRequest := *openapiclient.NewApiMgrModuleModuleNamePutRequest("Config_example") // ApiMgrModuleModuleNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MgrModuleAPI.ApiMgrModuleModuleNamePut(context.Background(), moduleName).ApiMgrModuleModuleNamePutRequest(apiMgrModuleModuleNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MgrModuleAPI.ApiMgrModuleModuleNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moduleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiMgrModuleModuleNamePutRequest struct via the builder pattern


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

