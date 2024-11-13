# \HardwareAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiHardwareSummaryGet**](HardwareAPI.md#ApiHardwareSummaryGet) | **Get** /api/hardware/summary | Retrieve a summary of the hardware health status



## ApiHardwareSummaryGet

> ApiHardwareSummaryGet(ctx).Categories(categories).Hostname(hostname).Execute()

Retrieve a summary of the hardware health status



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
	categories := "categories_example" // string |  (optional)
	hostname := "hostname_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HardwareAPI.ApiHardwareSummaryGet(context.Background()).Categories(categories).Hostname(hostname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HardwareAPI.ApiHardwareSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHardwareSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categories** | **string** |  | 
 **hostname** | **string** |  | 

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

