# \UserChangePasswordAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUserUsernameChangePasswordPost**](UserChangePasswordAPI.md#ApiUserUsernameChangePasswordPost) | **Post** /api/user/{username}/change_password | 



## ApiUserUsernameChangePasswordPost

> ApiUserUsernameChangePasswordPost(ctx, username).ApiUserUsernameChangePasswordPostRequest(apiUserUsernameChangePasswordPostRequest).Execute()



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
	username := "username_example" // string | 
	apiUserUsernameChangePasswordPostRequest := *openapiclient.NewApiUserUsernameChangePasswordPostRequest("NewPassword_example", "OldPassword_example") // ApiUserUsernameChangePasswordPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserChangePasswordAPI.ApiUserUsernameChangePasswordPost(context.Background(), username).ApiUserUsernameChangePasswordPostRequest(apiUserUsernameChangePasswordPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserChangePasswordAPI.ApiUserUsernameChangePasswordPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiUserUsernameChangePasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiUserUsernameChangePasswordPostRequest** | [**ApiUserUsernameChangePasswordPostRequest**](ApiUserUsernameChangePasswordPostRequest.md) |  | 

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

