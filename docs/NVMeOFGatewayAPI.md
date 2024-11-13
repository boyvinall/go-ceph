# \NVMeOFGatewayAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiNvmeofGatewayGet**](NVMeOFGatewayAPI.md#ApiNvmeofGatewayGet) | **Get** /api/nvmeof/gateway | Get information about the NVMeoF gateway
[**ApiNvmeofGatewayGroupGet**](NVMeOFGatewayAPI.md#ApiNvmeofGatewayGroupGet) | **Get** /api/nvmeof/gateway/group | 



## ApiNvmeofGatewayGet

> ApiNvmeofGatewayGet(ctx).GwGroup(gwGroup).Execute()

Get information about the NVMeoF gateway

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
	gwGroup := "gwGroup_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFGatewayAPI.ApiNvmeofGatewayGet(context.Background()).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFGatewayAPI.ApiNvmeofGatewayGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofGatewayGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gwGroup** | **string** |  | 

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


## ApiNvmeofGatewayGroupGet

> ApiNvmeofGatewayGroupGet(ctx).Execute()



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
	r, err := apiClient.NVMeOFGatewayAPI.ApiNvmeofGatewayGroupGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFGatewayAPI.ApiNvmeofGatewayGroupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofGatewayGroupGetRequest struct via the builder pattern


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

