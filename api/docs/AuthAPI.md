# \AuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAuthCheckPost**](AuthAPI.md#ApiAuthCheckPost) | **Post** /api/auth/check | Check token Authentication
[**ApiAuthLogoutPost**](AuthAPI.md#ApiAuthLogoutPost) | **Post** /api/auth/logout | 
[**ApiAuthPost**](AuthAPI.md#ApiAuthPost) | **Post** /api/auth | Dashboard Authentication



## ApiAuthCheckPost

> ApiAuthCheckPost201Response ApiAuthCheckPost(ctx).Token(token).ApiAuthCheckPostRequest(apiAuthCheckPostRequest).Execute()

Check token Authentication

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
	token := "token_example" // string | Authentication Token
	apiAuthCheckPostRequest := *openapiclient.NewApiAuthCheckPostRequest("Token_example") // ApiAuthCheckPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.ApiAuthCheckPost(context.Background()).Token(token).ApiAuthCheckPostRequest(apiAuthCheckPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.ApiAuthCheckPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiAuthCheckPost`: ApiAuthCheckPost201Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.ApiAuthCheckPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthCheckPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Authentication Token | 
 **apiAuthCheckPostRequest** | [**ApiAuthCheckPostRequest**](ApiAuthCheckPostRequest.md) |  | 

### Return type

[**ApiAuthCheckPost201Response**](ApiAuthCheckPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthLogoutPost

> ApiAuthLogoutPost(ctx).Execute()



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
	r, err := apiClient.AuthAPI.ApiAuthLogoutPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.ApiAuthLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthLogoutPostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthPost

> ApiAuthPost201Response ApiAuthPost(ctx).ApiAuthPostRequest(apiAuthPostRequest).Execute()

Dashboard Authentication

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
	apiAuthPostRequest := *openapiclient.NewApiAuthPostRequest("Password_example", "Username_example") // ApiAuthPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.ApiAuthPost(context.Background()).ApiAuthPostRequest(apiAuthPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.ApiAuthPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiAuthPost`: ApiAuthPost201Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.ApiAuthPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAuthPostRequest** | [**ApiAuthPostRequest**](ApiAuthPostRequest.md) |  | 

### Return type

[**ApiAuthPost201Response**](ApiAuthPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

