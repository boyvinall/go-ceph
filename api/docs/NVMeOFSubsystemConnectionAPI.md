# \NVMeOFSubsystemConnectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiNvmeofSubsystemNqnConnectionGet**](NVMeOFSubsystemConnectionAPI.md#ApiNvmeofSubsystemNqnConnectionGet) | **Get** /api/nvmeof/subsystem/{nqn}/connection | List all NVMeoF Subsystem Connections



## ApiNvmeofSubsystemNqnConnectionGet

> ApiNvmeofSubsystemNqnConnectionGet(ctx, nqn).GwGroup(gwGroup).Execute()

List all NVMeoF Subsystem Connections

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
	r, err := apiClient.NVMeOFSubsystemConnectionAPI.ApiNvmeofSubsystemNqnConnectionGet(context.Background(), nqn).GwGroup(gwGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NVMeOFSubsystemConnectionAPI.ApiNvmeofSubsystemNqnConnectionGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiNvmeofSubsystemNqnConnectionGetRequest struct via the builder pattern


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

