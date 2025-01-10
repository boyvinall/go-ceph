# \TelemetryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTelemetryPut**](TelemetryAPI.md#ApiTelemetryPut) | **Put** /api/telemetry | 
[**ApiTelemetryReportGet**](TelemetryAPI.md#ApiTelemetryReportGet) | **Get** /api/telemetry/report | Get Detailed Telemetry report



## ApiTelemetryPut

> ApiTelemetryPut(ctx).ApiTelemetryPutRequest(apiTelemetryPutRequest).Execute()





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
	apiTelemetryPutRequest := *openapiclient.NewApiTelemetryPutRequest() // ApiTelemetryPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TelemetryAPI.ApiTelemetryPut(context.Background()).ApiTelemetryPutRequest(apiTelemetryPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelemetryAPI.ApiTelemetryPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTelemetryPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiTelemetryPutRequest** | [**ApiTelemetryPutRequest**](ApiTelemetryPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTelemetryReportGet

> ApiTelemetryReportGet200Response ApiTelemetryReportGet(ctx).Execute()

Get Detailed Telemetry report



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
	resp, r, err := apiClient.TelemetryAPI.ApiTelemetryReportGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelemetryAPI.ApiTelemetryReportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiTelemetryReportGet`: ApiTelemetryReportGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TelemetryAPI.ApiTelemetryReportGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiTelemetryReportGetRequest struct via the builder pattern


### Return type

[**ApiTelemetryReportGet200Response**](ApiTelemetryReportGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

