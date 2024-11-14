# \PerfCountersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPerfCountersGet**](PerfCountersAPI.md#ApiPerfCountersGet) | **Get** /api/perf_counters | Display Perf Counters



## ApiPerfCountersGet

> ApiPerfCountersGet200Response ApiPerfCountersGet(ctx).Execute()

Display Perf Counters

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
	resp, r, err := apiClient.PerfCountersAPI.ApiPerfCountersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerfCountersAPI.ApiPerfCountersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPerfCountersGet`: ApiPerfCountersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PerfCountersAPI.ApiPerfCountersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiPerfCountersGetRequest struct via the builder pattern


### Return type

[**ApiPerfCountersGet200Response**](ApiPerfCountersGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

