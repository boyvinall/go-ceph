# \SummaryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSummaryGet**](SummaryAPI.md#ApiSummaryGet) | **Get** /api/summary | Display Summary



## ApiSummaryGet

> ApiSummaryGet200Response ApiSummaryGet(ctx).Execute()

Display Summary

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
	resp, r, err := apiClient.SummaryAPI.ApiSummaryGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummaryAPI.ApiSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSummaryGet`: ApiSummaryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SummaryAPI.ApiSummaryGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiSummaryGetRequest struct via the builder pattern


### Return type

[**ApiSummaryGet200Response**](ApiSummaryGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

