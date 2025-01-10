# \MonitorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiMonitorGet**](MonitorAPI.md#ApiMonitorGet) | **Get** /api/monitor | Get Monitor Details



## ApiMonitorGet

> ApiMonitorGet200Response ApiMonitorGet(ctx).Execute()

Get Monitor Details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorAPI.ApiMonitorGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorAPI.ApiMonitorGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiMonitorGet`: ApiMonitorGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MonitorAPI.ApiMonitorGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiMonitorGetRequest struct via the builder pattern


### Return type

[**ApiMonitorGet200Response**](ApiMonitorGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

