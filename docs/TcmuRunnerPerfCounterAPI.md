# \TcmuRunnerPerfCounterAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPerfCountersTcmuRunnerServiceIdGet**](TcmuRunnerPerfCounterAPI.md#ApiPerfCountersTcmuRunnerServiceIdGet) | **Get** /api/perf_counters/tcmu-runner/{service_id} | 



## ApiPerfCountersTcmuRunnerServiceIdGet

> ApiPerfCountersTcmuRunnerServiceIdGet(ctx, serviceId).Execute()



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
	serviceId := "serviceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TcmuRunnerPerfCounterAPI.ApiPerfCountersTcmuRunnerServiceIdGet(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TcmuRunnerPerfCounterAPI.ApiPerfCountersTcmuRunnerServiceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPerfCountersTcmuRunnerServiceIdGetRequest struct via the builder pattern


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

