# \NVMeOFSubsystemAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiNvmeofSubsystemGet**](NVMeOFSubsystemAPI.md#ApiNvmeofSubsystemGet) | **Get** /api/nvmeof/subsystem | List all NVMeoF subsystems
[**ApiNvmeofSubsystemNqnDelete**](NVMeOFSubsystemAPI.md#ApiNvmeofSubsystemNqnDelete) | **Delete** /api/nvmeof/subsystem/{nqn} | Delete an existing NVMeoF subsystem
[**ApiNvmeofSubsystemNqnGet**](NVMeOFSubsystemAPI.md#ApiNvmeofSubsystemNqnGet) | **Get** /api/nvmeof/subsystem/{nqn} | Get information from a specific NVMeoF subsystem
[**ApiNvmeofSubsystemPost**](NVMeOFSubsystemAPI.md#ApiNvmeofSubsystemPost) | **Post** /api/nvmeof/subsystem | Create a new NVMeoF subsystem



## ApiNvmeofSubsystemGet

> ApiNvmeofSubsystemGet(ctx).GwGroup(gwGroup).Execute()

List all NVMeoF subsystems

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
	gwGroup := "gwGroup_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemAPI.ApiNvmeofSubsystemGet(context.Background()).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemAPI.ApiNvmeofSubsystemGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gwGroup** | **string** |  | 

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


## ApiNvmeofSubsystemNqnDelete

> ApiNvmeofSubsystemNqnDelete(ctx, nqn).Force(force).GwGroup(gwGroup).Execute()

Delete an existing NVMeoF subsystem

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
	nqn := "nqn_example" // string | NVMeoF subsystem NQN
	force := true // bool | Force delete (optional)
	gwGroup := "gwGroup_example" // string | NVMeoF gateway group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemAPI.ApiNvmeofSubsystemNqnDelete(context.Background(), nqn).Force(force).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemAPI.ApiNvmeofSubsystemNqnDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nqn** | **string** | NVMeoF subsystem NQN | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force delete | 
 **gwGroup** | **string** | NVMeoF gateway group | 

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


## ApiNvmeofSubsystemNqnGet

> ApiNvmeofSubsystemNqnGet(ctx, nqn).GwGroup(gwGroup).Execute()

Get information from a specific NVMeoF subsystem

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
	nqn := "nqn_example" // string | NVMeoF subsystem NQN
	gwGroup := "gwGroup_example" // string | NVMeoF gateway group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemAPI.ApiNvmeofSubsystemNqnGet(context.Background(), nqn).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemAPI.ApiNvmeofSubsystemNqnGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nqn** | **string** | NVMeoF subsystem NQN | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gwGroup** | **string** | NVMeoF gateway group | 

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


## ApiNvmeofSubsystemPost

> ApiNvmeofSubsystemPost(ctx).ApiNvmeofSubsystemPostRequest(apiNvmeofSubsystemPostRequest).Execute()

Create a new NVMeoF subsystem

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
	apiNvmeofSubsystemPostRequest := *openapiclient.NewApiNvmeofSubsystemPostRequest(false, "Nqn_example") // ApiNvmeofSubsystemPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemAPI.ApiNvmeofSubsystemPost(context.Background()).ApiNvmeofSubsystemPostRequest(apiNvmeofSubsystemPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemAPI.ApiNvmeofSubsystemPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiNvmeofSubsystemPostRequest** | [**ApiNvmeofSubsystemPostRequest**](ApiNvmeofSubsystemPostRequest.md) |  | 

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

