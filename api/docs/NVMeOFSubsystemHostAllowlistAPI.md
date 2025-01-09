# \NVMeOFSubsystemHostAllowlistAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiNvmeofSubsystemNqnHostGet**](NVMeOFSubsystemHostAllowlistAPI.md#ApiNvmeofSubsystemNqnHostGet) | **Get** /api/nvmeof/subsystem/{nqn}/host | List all allowed hosts for an NVMeoF subsystem
[**ApiNvmeofSubsystemNqnHostHostNqnDelete**](NVMeOFSubsystemHostAllowlistAPI.md#ApiNvmeofSubsystemNqnHostHostNqnDelete) | **Delete** /api/nvmeof/subsystem/{nqn}/host/{host_nqn} | Disallow hosts from accessing an NVMeoF subsystem
[**ApiNvmeofSubsystemNqnHostPost**](NVMeOFSubsystemHostAllowlistAPI.md#ApiNvmeofSubsystemNqnHostPost) | **Post** /api/nvmeof/subsystem/{nqn}/host | Allow hosts to access an NVMeoF subsystem



## ApiNvmeofSubsystemNqnHostGet

> ApiNvmeofSubsystemNqnHostGet(ctx, nqn).GwGroup(gwGroup).Execute()

List all allowed hosts for an NVMeoF subsystem

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
	r, err := apiClient.NVMeOFSubsystemHostAllowlistAPI.ApiNvmeofSubsystemNqnHostGet(context.Background(), nqn).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemHostAllowlistAPI.ApiNvmeofSubsystemNqnHostGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnHostGetRequest struct via the builder pattern


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


## ApiNvmeofSubsystemNqnHostHostNqnDelete

> ApiNvmeofSubsystemNqnHostHostNqnDelete(ctx, nqn, hostNqn).GwGroup(gwGroup).Execute()

Disallow hosts from accessing an NVMeoF subsystem

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
	hostNqn := "hostNqn_example" // string | NVMeoF host NQN. Use \"*\" to disallow any host.
	gwGroup := "gwGroup_example" // string | NVMeoF gateway group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemHostAllowlistAPI.ApiNvmeofSubsystemNqnHostHostNqnDelete(context.Background(), nqn, hostNqn).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemHostAllowlistAPI.ApiNvmeofSubsystemNqnHostHostNqnDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nqn** | **string** | NVMeoF subsystem NQN | 
**hostNqn** | **string** | NVMeoF host NQN. Use \&quot;*\&quot; to disallow any host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnHostHostNqnDeleteRequest struct via the builder pattern


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


## ApiNvmeofSubsystemNqnHostPost

> ApiNvmeofSubsystemNqnHostPost(ctx, nqn).ApiNvmeofSubsystemNqnHostPostRequest(apiNvmeofSubsystemNqnHostPostRequest).Execute()

Allow hosts to access an NVMeoF subsystem

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
	apiNvmeofSubsystemNqnHostPostRequest := *openapiclient.NewApiNvmeofSubsystemNqnHostPostRequest("HostNqn_example") // ApiNvmeofSubsystemNqnHostPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemHostAllowlistAPI.ApiNvmeofSubsystemNqnHostPost(context.Background(), nqn).ApiNvmeofSubsystemNqnHostPostRequest(apiNvmeofSubsystemNqnHostPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemHostAllowlistAPI.ApiNvmeofSubsystemNqnHostPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnHostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiNvmeofSubsystemNqnHostPostRequest** | [**ApiNvmeofSubsystemNqnHostPostRequest**](ApiNvmeofSubsystemNqnHostPostRequest.md) |  | 

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

