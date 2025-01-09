# \UserPasswordPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUserValidatePasswordPost**](UserPasswordPolicyAPI.md#ApiUserValidatePasswordPost) | **Post** /api/user/validate_password | 



## ApiUserValidatePasswordPost

> ApiUserValidatePasswordPost(ctx).ApiUserValidatePasswordPostRequest(apiUserValidatePasswordPostRequest).Execute()





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
	apiUserValidatePasswordPostRequest := *openapiclient.NewApiUserValidatePasswordPostRequest("Password_example") // ApiUserValidatePasswordPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserPasswordPolicyAPI.ApiUserValidatePasswordPost(context.Background()).ApiUserValidatePasswordPostRequest(apiUserValidatePasswordPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPasswordPolicyAPI.ApiUserValidatePasswordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUserValidatePasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiUserValidatePasswordPostRequest** | [**ApiUserValidatePasswordPostRequest**](ApiUserValidatePasswordPostRequest.md) |  | 

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

