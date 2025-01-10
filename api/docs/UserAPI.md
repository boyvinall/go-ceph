# \UserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUserGet**](UserAPI.md#ApiUserGet) | **Get** /api/user | Get List Of Users
[**ApiUserPost**](UserAPI.md#ApiUserPost) | **Post** /api/user | 
[**ApiUserUsernameDelete**](UserAPI.md#ApiUserUsernameDelete) | **Delete** /api/user/{username} | 
[**ApiUserUsernameGet**](UserAPI.md#ApiUserUsernameGet) | **Get** /api/user/{username} | 
[**ApiUserUsernamePut**](UserAPI.md#ApiUserUsernamePut) | **Put** /api/user/{username} | 



## ApiUserGet

> ApiUserGet200Response ApiUserGet(ctx).Execute()

Get List Of Users

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
	resp, r, err := apiClient.UserAPI.ApiUserGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ApiUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiUserGet`: ApiUserGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ApiUserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserGetRequest struct via the builder pattern


### Return type

[**ApiUserGet200Response**](ApiUserGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserPost

> ApiUserPost(ctx).ApiUserPostRequest(apiUserPostRequest).Execute()



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
	apiUserPostRequest := *openapiclient.NewApiUserPostRequest() // ApiUserPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.ApiUserPost(context.Background()).ApiUserPostRequest(apiUserPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ApiUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiUserPostRequest** | [**ApiUserPostRequest**](ApiUserPostRequest.md) |  | 

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


## ApiUserUsernameDelete

> ApiUserUsernameDelete(ctx, username).Execute()



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
	username := "username_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.ApiUserUsernameDelete(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ApiUserUsernameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserUsernameDeleteRequest struct via the builder pattern


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


## ApiUserUsernameGet

> ApiUserUsernameGet(ctx, username).Execute()



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
	username := "username_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.ApiUserUsernameGet(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ApiUserUsernameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserUsernameGetRequest struct via the builder pattern


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


## ApiUserUsernamePut

> ApiUserUsernamePut(ctx, username).ApiUserUsernamePutRequest(apiUserUsernamePutRequest).Execute()



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
	username := "username_example" // string | 
	apiUserUsernamePutRequest := *openapiclient.NewApiUserUsernamePutRequest() // ApiUserUsernamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.ApiUserUsernamePut(context.Background(), username).ApiUserUsernamePutRequest(apiUserUsernamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ApiUserUsernamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserUsernamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiUserUsernamePutRequest** | [**ApiUserUsernamePutRequest**](ApiUserUsernamePutRequest.md) |  | 

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

