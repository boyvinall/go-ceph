# \NVMeOFSubsystemNamespaceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiNvmeofSubsystemNqnNamespaceGet**](NVMeOFSubsystemNamespaceAPI.md#ApiNvmeofSubsystemNqnNamespaceGet) | **Get** /api/nvmeof/subsystem/{nqn}/namespace | List all NVMeoF namespaces in a subsystem
[**ApiNvmeofSubsystemNqnNamespaceNsidDelete**](NVMeOFSubsystemNamespaceAPI.md#ApiNvmeofSubsystemNqnNamespaceNsidDelete) | **Delete** /api/nvmeof/subsystem/{nqn}/namespace/{nsid} | Delete an existing NVMeoF namespace
[**ApiNvmeofSubsystemNqnNamespaceNsidGet**](NVMeOFSubsystemNamespaceAPI.md#ApiNvmeofSubsystemNqnNamespaceNsidGet) | **Get** /api/nvmeof/subsystem/{nqn}/namespace/{nsid} | Get info from specified NVMeoF namespace
[**ApiNvmeofSubsystemNqnNamespaceNsidIoStatsGet**](NVMeOFSubsystemNamespaceAPI.md#ApiNvmeofSubsystemNqnNamespaceNsidIoStatsGet) | **Get** /api/nvmeof/subsystem/{nqn}/namespace/{nsid}/io_stats | Get IO stats from specified NVMeoF namespace
[**ApiNvmeofSubsystemNqnNamespaceNsidPatch**](NVMeOFSubsystemNamespaceAPI.md#ApiNvmeofSubsystemNqnNamespaceNsidPatch) | **Patch** /api/nvmeof/subsystem/{nqn}/namespace/{nsid} | Update an existing NVMeoF namespace
[**ApiNvmeofSubsystemNqnNamespacePost**](NVMeOFSubsystemNamespaceAPI.md#ApiNvmeofSubsystemNqnNamespacePost) | **Post** /api/nvmeof/subsystem/{nqn}/namespace | Create a new NVMeoF namespace



## ApiNvmeofSubsystemNqnNamespaceGet

> ApiNvmeofSubsystemNqnNamespaceGet(ctx, nqn).GwGroup(gwGroup).Execute()

List all NVMeoF namespaces in a subsystem

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
	r, err := apiClient.NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespaceGet(context.Background(), nqn).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespaceGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnNamespaceGetRequest struct via the builder pattern


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


## ApiNvmeofSubsystemNqnNamespaceNsidDelete

> ApiNvmeofSubsystemNqnNamespaceNsidDelete(ctx, nqn, nsid).GwGroup(gwGroup).Execute()

Delete an existing NVMeoF namespace

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
	nsid := "nsid_example" // string | NVMeoF Namespace ID
	gwGroup := "gwGroup_example" // string | NVMeoF gateway group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespaceNsidDelete(context.Background(), nqn, nsid).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespaceNsidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nqn** | **string** | NVMeoF subsystem NQN | 
**nsid** | **string** | NVMeoF Namespace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnNamespaceNsidDeleteRequest struct via the builder pattern


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


## ApiNvmeofSubsystemNqnNamespaceNsidGet

> ApiNvmeofSubsystemNqnNamespaceNsidGet(ctx, nqn, nsid).GwGroup(gwGroup).Execute()

Get info from specified NVMeoF namespace

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
	nsid := "nsid_example" // string | NVMeoF Namespace ID
	gwGroup := "gwGroup_example" // string | NVMeoF gateway group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespaceNsidGet(context.Background(), nqn, nsid).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespaceNsidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nqn** | **string** | NVMeoF subsystem NQN | 
**nsid** | **string** | NVMeoF Namespace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnNamespaceNsidGetRequest struct via the builder pattern


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


## ApiNvmeofSubsystemNqnNamespaceNsidIoStatsGet

> ApiNvmeofSubsystemNqnNamespaceNsidIoStatsGet(ctx, nqn, nsid).GwGroup(gwGroup).Execute()

Get IO stats from specified NVMeoF namespace

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
	nsid := "nsid_example" // string | NVMeoF Namespace ID
	gwGroup := "gwGroup_example" // string | NVMeoF gateway group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespaceNsidIoStatsGet(context.Background(), nqn, nsid).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespaceNsidIoStatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nqn** | **string** | NVMeoF subsystem NQN | 
**nsid** | **string** | NVMeoF Namespace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnNamespaceNsidIoStatsGetRequest struct via the builder pattern


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


## ApiNvmeofSubsystemNqnNamespaceNsidPatch

> ApiNvmeofSubsystemNqnNamespaceNsidPatch(ctx, nqn, nsid).ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest(apiNvmeofSubsystemNqnNamespaceNsidPatchRequest).Execute()

Update an existing NVMeoF namespace

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
	nsid := "nsid_example" // string | NVMeoF Namespace ID
	apiNvmeofSubsystemNqnNamespaceNsidPatchRequest := *openapiclient.NewApiNvmeofSubsystemNqnNamespaceNsidPatchRequest() // ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespaceNsidPatch(context.Background(), nqn, nsid).ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest(apiNvmeofSubsystemNqnNamespaceNsidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespaceNsidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nqn** | **string** | NVMeoF subsystem NQN | 
**nsid** | **string** | NVMeoF Namespace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnNamespaceNsidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiNvmeofSubsystemNqnNamespaceNsidPatchRequest** | [**ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest**](ApiNvmeofSubsystemNqnNamespaceNsidPatchRequest.md) |  | 

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


## ApiNvmeofSubsystemNqnNamespacePost

> ApiNvmeofSubsystemNqnNamespacePost(ctx, nqn).ApiNvmeofSubsystemNqnNamespacePostRequest(apiNvmeofSubsystemNqnNamespacePostRequest).Execute()

Create a new NVMeoF namespace

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
	apiNvmeofSubsystemNqnNamespacePostRequest := *openapiclient.NewApiNvmeofSubsystemNqnNamespacePostRequest("RbdImageName_example") // ApiNvmeofSubsystemNqnNamespacePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespacePost(context.Background(), nqn).ApiNvmeofSubsystemNqnNamespacePostRequest(apiNvmeofSubsystemNqnNamespacePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemNamespaceAPI.ApiNvmeofSubsystemNqnNamespacePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnNamespacePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiNvmeofSubsystemNqnNamespacePostRequest** | [**ApiNvmeofSubsystemNqnNamespacePostRequest**](ApiNvmeofSubsystemNqnNamespacePostRequest.md) |  | 

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

