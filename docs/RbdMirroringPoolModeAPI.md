# \RbdMirroringPoolModeAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBlockMirroringPoolPoolNameGet**](RbdMirroringPoolModeAPI.md#ApiBlockMirroringPoolPoolNameGet) | **Get** /api/block/mirroring/pool/{pool_name} | Display Rbd Mirroring Summary
[**ApiBlockMirroringPoolPoolNamePut**](RbdMirroringPoolModeAPI.md#ApiBlockMirroringPoolPoolNamePut) | **Put** /api/block/mirroring/pool/{pool_name} | 



## ApiBlockMirroringPoolPoolNameGet

> ApiBlockMirroringPoolPoolNameGet200Response ApiBlockMirroringPoolPoolNameGet(ctx, poolName).Execute()

Display Rbd Mirroring Summary

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
	poolName := "poolName_example" // string | Pool Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdMirroringPoolModeAPI.ApiBlockMirroringPoolPoolNameGet(context.Background(), poolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringPoolModeAPI.ApiBlockMirroringPoolPoolNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiBlockMirroringPoolPoolNameGet`: ApiBlockMirroringPoolPoolNameGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RbdMirroringPoolModeAPI.ApiBlockMirroringPoolPoolNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolName** | **string** | Pool Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockMirroringPoolPoolNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiBlockMirroringPoolPoolNameGet200Response**](ApiBlockMirroringPoolPoolNameGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBlockMirroringPoolPoolNamePut

> ApiBlockMirroringPoolPoolNamePut(ctx, poolName).ApiBlockMirroringPoolPoolNamePutRequest(apiBlockMirroringPoolPoolNamePutRequest).Execute()



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
	poolName := "poolName_example" // string | 
	apiBlockMirroringPoolPoolNamePutRequest := *openapiclient.NewApiBlockMirroringPoolPoolNamePutRequest() // ApiBlockMirroringPoolPoolNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdMirroringPoolModeAPI.ApiBlockMirroringPoolPoolNamePut(context.Background(), poolName).ApiBlockMirroringPoolPoolNamePutRequest(apiBlockMirroringPoolPoolNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringPoolModeAPI.ApiBlockMirroringPoolPoolNamePut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiBlockMirroringPoolPoolNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiBlockMirroringPoolPoolNamePutRequest** | [**ApiBlockMirroringPoolPoolNamePutRequest**](ApiBlockMirroringPoolPoolNamePutRequest.md) |  | 

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

