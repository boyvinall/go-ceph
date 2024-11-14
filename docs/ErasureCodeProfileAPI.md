# \ErasureCodeProfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiErasureCodeProfileGet**](ErasureCodeProfileAPI.md#ApiErasureCodeProfileGet) | **Get** /api/erasure_code_profile | List Erasure Code Profile Information
[**ApiErasureCodeProfileNameDelete**](ErasureCodeProfileAPI.md#ApiErasureCodeProfileNameDelete) | **Delete** /api/erasure_code_profile/{name} | 
[**ApiErasureCodeProfileNameGet**](ErasureCodeProfileAPI.md#ApiErasureCodeProfileNameGet) | **Get** /api/erasure_code_profile/{name} | 
[**ApiErasureCodeProfilePost**](ErasureCodeProfileAPI.md#ApiErasureCodeProfilePost) | **Post** /api/erasure_code_profile | 



## ApiErasureCodeProfileGet

> []ApiErasureCodeProfileGet200ResponseInner ApiErasureCodeProfileGet(ctx).Execute()

List Erasure Code Profile Information

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
	resp, r, err := apiClient.ErasureCodeProfileAPI.ApiErasureCodeProfileGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErasureCodeProfileAPI.ApiErasureCodeProfileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiErasureCodeProfileGet`: []ApiErasureCodeProfileGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ErasureCodeProfileAPI.ApiErasureCodeProfileGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiErasureCodeProfileGetRequest struct via the builder pattern


### Return type

[**[]ApiErasureCodeProfileGet200ResponseInner**](ApiErasureCodeProfileGet200ResponseInner.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiErasureCodeProfileNameDelete

> ApiErasureCodeProfileNameDelete(ctx, name).Execute()



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ErasureCodeProfileAPI.ApiErasureCodeProfileNameDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErasureCodeProfileAPI.ApiErasureCodeProfileNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiErasureCodeProfileNameDeleteRequest struct via the builder pattern


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


## ApiErasureCodeProfileNameGet

> ApiErasureCodeProfileNameGet(ctx, name).Execute()



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ErasureCodeProfileAPI.ApiErasureCodeProfileNameGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErasureCodeProfileAPI.ApiErasureCodeProfileNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiErasureCodeProfileNameGetRequest struct via the builder pattern


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


## ApiErasureCodeProfilePost

> ApiErasureCodeProfilePost(ctx).ApiErasureCodeProfilePostRequest(apiErasureCodeProfilePostRequest).Execute()



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
	apiErasureCodeProfilePostRequest := *openapiclient.NewApiErasureCodeProfilePostRequest("Name_example") // ApiErasureCodeProfilePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ErasureCodeProfileAPI.ApiErasureCodeProfilePost(context.Background()).ApiErasureCodeProfilePostRequest(apiErasureCodeProfilePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ErasureCodeProfileAPI.ApiErasureCodeProfilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiErasureCodeProfilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiErasureCodeProfilePostRequest** | [**ApiErasureCodeProfilePostRequest**](ApiErasureCodeProfilePostRequest.md) |  | 

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

