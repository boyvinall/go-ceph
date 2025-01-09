# \IscsiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiIscsiDiscoveryauthGet**](IscsiAPI.md#ApiIscsiDiscoveryauthGet) | **Get** /api/iscsi/discoveryauth | Get Iscsi discoveryauth Details
[**ApiIscsiDiscoveryauthPut**](IscsiAPI.md#ApiIscsiDiscoveryauthPut) | **Put** /api/iscsi/discoveryauth | Set Iscsi discoveryauth



## ApiIscsiDiscoveryauthGet

> []ApiIscsiDiscoveryauthGet200ResponseInner ApiIscsiDiscoveryauthGet(ctx).Execute()

Get Iscsi discoveryauth Details

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
	resp, r, err := apiClient.IscsiAPI.ApiIscsiDiscoveryauthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IscsiAPI.ApiIscsiDiscoveryauthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiIscsiDiscoveryauthGet`: []ApiIscsiDiscoveryauthGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `IscsiAPI.ApiIscsiDiscoveryauthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiIscsiDiscoveryauthGetRequest struct via the builder pattern


### Return type

[**[]ApiIscsiDiscoveryauthGet200ResponseInner**](ApiIscsiDiscoveryauthGet200ResponseInner.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiIscsiDiscoveryauthPut

> ApiIscsiDiscoveryauthPut(ctx).User(user).Password(password).MutualUser(mutualUser).MutualPassword(mutualPassword).ApiIscsiDiscoveryauthPutRequest(apiIscsiDiscoveryauthPutRequest).Execute()

Set Iscsi discoveryauth

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
	user := "user_example" // string | Username
	password := "password_example" // string | Password
	mutualUser := "mutualUser_example" // string | Mutual UserName
	mutualPassword := "mutualPassword_example" // string | Mutual Password
	apiIscsiDiscoveryauthPutRequest := *openapiclient.NewApiIscsiDiscoveryauthPutRequest("MutualPassword_example", "MutualUser_example", "Password_example", "User_example") // ApiIscsiDiscoveryauthPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IscsiAPI.ApiIscsiDiscoveryauthPut(context.Background()).User(user).Password(password).MutualUser(mutualUser).MutualPassword(mutualPassword).ApiIscsiDiscoveryauthPutRequest(apiIscsiDiscoveryauthPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IscsiAPI.ApiIscsiDiscoveryauthPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiIscsiDiscoveryauthPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | Username | 
 **password** | **string** | Password | 
 **mutualUser** | **string** | Mutual UserName | 
 **mutualPassword** | **string** | Mutual Password | 
 **apiIscsiDiscoveryauthPutRequest** | [**ApiIscsiDiscoveryauthPutRequest**](ApiIscsiDiscoveryauthPutRequest.md) |  | 

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

