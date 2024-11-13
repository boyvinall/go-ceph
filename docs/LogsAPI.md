# \LogsAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiLogsAllGet**](LogsAPI.md#ApiLogsAllGet) | **Get** /api/logs/all | Display Logs Configuration



## ApiLogsAllGet

> ApiLogsAllGet200Response ApiLogsAllGet(ctx).Execute()

Display Logs Configuration

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
	resp, r, err := apiClient.LogsAPI.ApiLogsAllGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ApiLogsAllGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiLogsAllGet`: ApiLogsAllGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.ApiLogsAllGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiLogsAllGetRequest struct via the builder pattern


### Return type

[**ApiLogsAllGet200Response**](ApiLogsAllGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

