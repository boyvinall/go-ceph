# \FeatureTogglesEndpointAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiFeatureTogglesGet**](FeatureTogglesEndpointAPI.md#ApiFeatureTogglesGet) | **Get** /api/feature_toggles | Get List Of Features



## ApiFeatureTogglesGet

> ApiFeatureTogglesGet200Response ApiFeatureTogglesGet(ctx).Execute()

Get List Of Features

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
	resp, r, err := apiClient.FeatureTogglesEndpointAPI.ApiFeatureTogglesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureTogglesEndpointAPI.ApiFeatureTogglesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiFeatureTogglesGet`: ApiFeatureTogglesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FeatureTogglesEndpointAPI.ApiFeatureTogglesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiFeatureTogglesGetRequest struct via the builder pattern


### Return type

[**ApiFeatureTogglesGet200Response**](ApiFeatureTogglesGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

