# \RGWAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRgwRolesGet**](RGWAPI.md#ApiRgwRolesGet) | **Get** /api/rgw/roles | List RGW roles
[**ApiRgwRolesPost**](RGWAPI.md#ApiRgwRolesPost) | **Post** /api/rgw/roles | Create RGW role
[**ApiRgwRolesPut**](RGWAPI.md#ApiRgwRolesPut) | **Put** /api/rgw/roles | Edit RGW role
[**ApiRgwRolesRoleNameDelete**](RGWAPI.md#ApiRgwRolesRoleNameDelete) | **Delete** /api/rgw/roles/{role_name} | Delete RGW role



## ApiRgwRolesGet

> ApiRgwRolesGet(ctx).Execute()

List RGW roles

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
	r, err := apiClient.RGWAPI.ApiRgwRolesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RGWAPI.ApiRgwRolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRolesGetRequest struct via the builder pattern


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


## ApiRgwRolesPost

> ApiRgwRolesPost(ctx).ApiRgwRolesPostRequest(apiRgwRolesPostRequest).Execute()

Create RGW role

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
	apiRgwRolesPostRequest := *openapiclient.NewApiRgwRolesPostRequest() // ApiRgwRolesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RGWAPI.ApiRgwRolesPost(context.Background()).ApiRgwRolesPostRequest(apiRgwRolesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RGWAPI.ApiRgwRolesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwRolesPostRequest** | [**ApiRgwRolesPostRequest**](ApiRgwRolesPostRequest.md) |  | 

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


## ApiRgwRolesPut

> ApiRgwRolesPut(ctx).ApiRgwRolesPutRequest(apiRgwRolesPutRequest).Execute()

Edit RGW role

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
	apiRgwRolesPutRequest := *openapiclient.NewApiRgwRolesPutRequest("MaxSessionDuration_example", "RoleName_example") // ApiRgwRolesPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RGWAPI.ApiRgwRolesPut(context.Background()).ApiRgwRolesPutRequest(apiRgwRolesPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RGWAPI.ApiRgwRolesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRolesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwRolesPutRequest** | [**ApiRgwRolesPutRequest**](ApiRgwRolesPutRequest.md) |  | 

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


## ApiRgwRolesRoleNameDelete

> ApiRgwRolesRoleNameDelete(ctx, roleName).Execute()

Delete RGW role

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
	roleName := "roleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RGWAPI.ApiRgwRolesRoleNameDelete(context.Background(), roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RGWAPI.ApiRgwRolesRoleNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRolesRoleNameDeleteRequest struct via the builder pattern


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

