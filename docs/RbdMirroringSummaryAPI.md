# \RbdMirroringSummaryAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBlockMirroringSummaryGet**](RbdMirroringSummaryAPI.md#ApiBlockMirroringSummaryGet) | **Get** /api/block/mirroring/summary | Display Rbd Mirroring Summary



## ApiBlockMirroringSummaryGet

> ApiBlockMirroringSummaryGet200Response ApiBlockMirroringSummaryGet(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdMirroringSummaryAPI.ApiBlockMirroringSummaryGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringSummaryAPI.ApiBlockMirroringSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiBlockMirroringSummaryGet`: ApiBlockMirroringSummaryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RbdMirroringSummaryAPI.ApiBlockMirroringSummaryGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockMirroringSummaryGetRequest struct via the builder pattern


### Return type

[**ApiBlockMirroringSummaryGet200Response**](ApiBlockMirroringSummaryGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

