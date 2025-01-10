# \GrafanaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiGrafanaDashboardsPost**](GrafanaAPI.md#ApiGrafanaDashboardsPost) | **Post** /api/grafana/dashboards | 
[**ApiGrafanaUrlGet**](GrafanaAPI.md#ApiGrafanaUrlGet) | **Get** /api/grafana/url | List Grafana URL Instance
[**ApiGrafanaValidationParamsGet**](GrafanaAPI.md#ApiGrafanaValidationParamsGet) | **Get** /api/grafana/validation/{params} | 



## ApiGrafanaDashboardsPost

> ApiGrafanaDashboardsPost(ctx).Execute()



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
	r, err := apiClient.GrafanaAPI.ApiGrafanaDashboardsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrafanaAPI.ApiGrafanaDashboardsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiGrafanaDashboardsPostRequest struct via the builder pattern


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


## ApiGrafanaUrlGet

> ApiGrafanaUrlGet200Response ApiGrafanaUrlGet(ctx).Execute()

List Grafana URL Instance

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
	resp, r, err := apiClient.GrafanaAPI.ApiGrafanaUrlGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrafanaAPI.ApiGrafanaUrlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiGrafanaUrlGet`: ApiGrafanaUrlGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GrafanaAPI.ApiGrafanaUrlGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiGrafanaUrlGetRequest struct via the builder pattern


### Return type

[**ApiGrafanaUrlGet200Response**](ApiGrafanaUrlGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiGrafanaValidationParamsGet

> ApiGrafanaValidationParamsGet(ctx, params).Execute()



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
	params := "params_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GrafanaAPI.ApiGrafanaValidationParamsGet(context.Background(), params).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrafanaAPI.ApiGrafanaValidationParamsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**params** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiGrafanaValidationParamsGetRequest struct via the builder pattern


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

