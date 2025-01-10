# \NVMeOFSubsystemListenerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiNvmeofSubsystemNqnListenerGet**](NVMeOFSubsystemListenerAPI.md#ApiNvmeofSubsystemNqnListenerGet) | **Get** /api/nvmeof/subsystem/{nqn}/listener | List all NVMeoF listeners
[**ApiNvmeofSubsystemNqnListenerHostNameTraddrDelete**](NVMeOFSubsystemListenerAPI.md#ApiNvmeofSubsystemNqnListenerHostNameTraddrDelete) | **Delete** /api/nvmeof/subsystem/{nqn}/listener/{host_name}/{traddr} | Delete an existing NVMeoF listener
[**ApiNvmeofSubsystemNqnListenerPost**](NVMeOFSubsystemListenerAPI.md#ApiNvmeofSubsystemNqnListenerPost) | **Post** /api/nvmeof/subsystem/{nqn}/listener | Create a new NVMeoF listener



## ApiNvmeofSubsystemNqnListenerGet

> ApiNvmeofSubsystemNqnListenerGet(ctx, nqn).GwGroup(gwGroup).Execute()

List all NVMeoF listeners

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
	nqn := "nqn_example" // string | NVMeoF subsystem NQN
	gwGroup := "gwGroup_example" // string | NVMeoF gateway group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemListenerAPI.ApiNvmeofSubsystemNqnListenerGet(context.Background(), nqn).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemListenerAPI.ApiNvmeofSubsystemNqnListenerGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnListenerGetRequest struct via the builder pattern


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


## ApiNvmeofSubsystemNqnListenerHostNameTraddrDelete

> ApiNvmeofSubsystemNqnListenerHostNameTraddrDelete(ctx, nqn, hostName, traddr).Trsvcid(trsvcid).Adrfam(adrfam).Force(force).GwGroup(gwGroup).Execute()

Delete an existing NVMeoF listener

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
	nqn := "nqn_example" // string | NVMeoF subsystem NQN
	hostName := "hostName_example" // string | NVMeoF hostname
	traddr := "traddr_example" // string | NVMeoF transport address
	trsvcid := int32(56) // int32 | NVMeoF transport service port (optional)
	adrfam := int32(56) // int32 | NVMeoF address family (0 - IPv4, 1 - IPv6) (optional)
	force := true // bool |  (optional)
	gwGroup := "gwGroup_example" // string | NVMeoF gateway group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemListenerAPI.ApiNvmeofSubsystemNqnListenerHostNameTraddrDelete(context.Background(), nqn, hostName, traddr).Trsvcid(trsvcid).Adrfam(adrfam).Force(force).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemListenerAPI.ApiNvmeofSubsystemNqnListenerHostNameTraddrDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nqn** | **string** | NVMeoF subsystem NQN | 
**hostName** | **string** | NVMeoF hostname | 
**traddr** | **string** | NVMeoF transport address | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnListenerHostNameTraddrDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **trsvcid** | **int32** | NVMeoF transport service port | 
 **adrfam** | **int32** | NVMeoF address family (0 - IPv4, 1 - IPv6) | 
 **force** | **bool** |  | 
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


## ApiNvmeofSubsystemNqnListenerPost

> ApiNvmeofSubsystemNqnListenerPost(ctx, nqn).ApiNvmeofSubsystemNqnListenerPostRequest(apiNvmeofSubsystemNqnListenerPostRequest).Execute()

Create a new NVMeoF listener

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
	nqn := "nqn_example" // string | NVMeoF subsystem NQN
	apiNvmeofSubsystemNqnListenerPostRequest := *openapiclient.NewApiNvmeofSubsystemNqnListenerPostRequest("HostName_example", "Traddr_example") // ApiNvmeofSubsystemNqnListenerPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemListenerAPI.ApiNvmeofSubsystemNqnListenerPost(context.Background(), nqn).ApiNvmeofSubsystemNqnListenerPostRequest(apiNvmeofSubsystemNqnListenerPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemListenerAPI.ApiNvmeofSubsystemNqnListenerPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnListenerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiNvmeofSubsystemNqnListenerPostRequest** | [**ApiNvmeofSubsystemNqnListenerPostRequest**](ApiNvmeofSubsystemNqnListenerPostRequest.md) |  | 

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

