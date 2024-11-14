# \CephfsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCephfsAuthPut**](CephfsAPI.md#ApiCephfsAuthPut) | **Put** /api/cephfs/auth | Set Ceph authentication capabilities for the specified user ID in the given path
[**ApiCephfsFsIdClientClientIdDelete**](CephfsAPI.md#ApiCephfsFsIdClientClientIdDelete) | **Delete** /api/cephfs/{fs_id}/client/{client_id} | 
[**ApiCephfsFsIdClientsGet**](CephfsAPI.md#ApiCephfsFsIdClientsGet) | **Get** /api/cephfs/{fs_id}/clients | 
[**ApiCephfsFsIdGet**](CephfsAPI.md#ApiCephfsFsIdGet) | **Get** /api/cephfs/{fs_id} | 
[**ApiCephfsFsIdGetRootDirectoryGet**](CephfsAPI.md#ApiCephfsFsIdGetRootDirectoryGet) | **Get** /api/cephfs/{fs_id}/get_root_directory | 
[**ApiCephfsFsIdLsDirGet**](CephfsAPI.md#ApiCephfsFsIdLsDirGet) | **Get** /api/cephfs/{fs_id}/ls_dir | 
[**ApiCephfsFsIdMdsCountersGet**](CephfsAPI.md#ApiCephfsFsIdMdsCountersGet) | **Get** /api/cephfs/{fs_id}/mds_counters | 
[**ApiCephfsFsIdQuotaGet**](CephfsAPI.md#ApiCephfsFsIdQuotaGet) | **Get** /api/cephfs/{fs_id}/quota | Get Cephfs Quotas of the specified path
[**ApiCephfsFsIdQuotaPut**](CephfsAPI.md#ApiCephfsFsIdQuotaPut) | **Put** /api/cephfs/{fs_id}/quota | 
[**ApiCephfsFsIdRenamePathPut**](CephfsAPI.md#ApiCephfsFsIdRenamePathPut) | **Put** /api/cephfs/{fs_id}/rename-path | 
[**ApiCephfsFsIdSnapshotDelete**](CephfsAPI.md#ApiCephfsFsIdSnapshotDelete) | **Delete** /api/cephfs/{fs_id}/snapshot | 
[**ApiCephfsFsIdSnapshotPost**](CephfsAPI.md#ApiCephfsFsIdSnapshotPost) | **Post** /api/cephfs/{fs_id}/snapshot | 
[**ApiCephfsFsIdStatfsGet**](CephfsAPI.md#ApiCephfsFsIdStatfsGet) | **Get** /api/cephfs/{fs_id}/statfs | Get Cephfs statfs of the specified path
[**ApiCephfsFsIdTreeDelete**](CephfsAPI.md#ApiCephfsFsIdTreeDelete) | **Delete** /api/cephfs/{fs_id}/tree | 
[**ApiCephfsFsIdTreePost**](CephfsAPI.md#ApiCephfsFsIdTreePost) | **Post** /api/cephfs/{fs_id}/tree | 
[**ApiCephfsFsIdUnlinkDelete**](CephfsAPI.md#ApiCephfsFsIdUnlinkDelete) | **Delete** /api/cephfs/{fs_id}/unlink | 
[**ApiCephfsFsIdWriteToFilePost**](CephfsAPI.md#ApiCephfsFsIdWriteToFilePost) | **Post** /api/cephfs/{fs_id}/write_to_file | 
[**ApiCephfsGet**](CephfsAPI.md#ApiCephfsGet) | **Get** /api/cephfs | 
[**ApiCephfsPost**](CephfsAPI.md#ApiCephfsPost) | **Post** /api/cephfs | 
[**ApiCephfsRemoveNameDelete**](CephfsAPI.md#ApiCephfsRemoveNameDelete) | **Delete** /api/cephfs/remove/{name} | Remove CephFS Volume
[**ApiCephfsRenamePut**](CephfsAPI.md#ApiCephfsRenamePut) | **Put** /api/cephfs/rename | Rename CephFS Volume



## ApiCephfsAuthPut

> ApiCephfsAuthPut(ctx).ApiCephfsAuthPutRequest(apiCephfsAuthPutRequest).Execute()

Set Ceph authentication capabilities for the specified user ID in the given path

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
	apiCephfsAuthPutRequest := *openapiclient.NewApiCephfsAuthPutRequest("Caps_example", "ClientId_example", "FsName_example", "RootSquash_example") // ApiCephfsAuthPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsAuthPut(context.Background()).ApiCephfsAuthPutRequest(apiCephfsAuthPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsAuthPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsAuthPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCephfsAuthPutRequest** | [**ApiCephfsAuthPutRequest**](ApiCephfsAuthPutRequest.md) |  | 

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


## ApiCephfsFsIdClientClientIdDelete

> ApiCephfsFsIdClientClientIdDelete(ctx, fsId, clientId).Execute()



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
	fsId := "fsId_example" // string | 
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdClientClientIdDelete(context.Background(), fsId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdClientClientIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdClientClientIdDeleteRequest struct via the builder pattern


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


## ApiCephfsFsIdClientsGet

> ApiCephfsFsIdClientsGet(ctx, fsId).Execute()



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
	fsId := "fsId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdClientsGet(context.Background(), fsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdClientsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdClientsGetRequest struct via the builder pattern


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


## ApiCephfsFsIdGet

> ApiCephfsFsIdGet(ctx, fsId).Execute()



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
	fsId := "fsId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdGet(context.Background(), fsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdGetRequest struct via the builder pattern


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


## ApiCephfsFsIdGetRootDirectoryGet

> ApiCephfsFsIdGetRootDirectoryGet(ctx, fsId).Execute()





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
	fsId := "fsId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdGetRootDirectoryGet(context.Background(), fsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdGetRootDirectoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdGetRootDirectoryGetRequest struct via the builder pattern


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


## ApiCephfsFsIdLsDirGet

> ApiCephfsFsIdLsDirGet(ctx, fsId).Path(path).Depth(depth).Execute()





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
	fsId := "fsId_example" // string | 
	path := "path_example" // string |  (optional)
	depth := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdLsDirGet(context.Background(), fsId).Path(path).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdLsDirGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdLsDirGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** |  | 
 **depth** | **int32** |  | 

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


## ApiCephfsFsIdMdsCountersGet

> ApiCephfsFsIdMdsCountersGet(ctx, fsId).Counters(counters).Execute()



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
	fsId := "fsId_example" // string | 
	counters := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdMdsCountersGet(context.Background(), fsId).Counters(counters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdMdsCountersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdMdsCountersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **counters** | **int32** |  | 

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


## ApiCephfsFsIdQuotaGet

> ApiCephfsFsIdQuotaGet200Response ApiCephfsFsIdQuotaGet(ctx, fsId).Path(path).Execute()

Get Cephfs Quotas of the specified path



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
	fsId := "fsId_example" // string | File System Identifier
	path := "path_example" // string | File System Path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CephfsAPI.ApiCephfsFsIdQuotaGet(context.Background(), fsId).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdQuotaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCephfsFsIdQuotaGet`: ApiCephfsFsIdQuotaGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CephfsAPI.ApiCephfsFsIdQuotaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** | File System Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdQuotaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | File System Path | 

### Return type

[**ApiCephfsFsIdQuotaGet200Response**](ApiCephfsFsIdQuotaGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCephfsFsIdQuotaPut

> ApiCephfsFsIdQuotaPut(ctx, fsId).ApiCephfsFsIdQuotaPutRequest(apiCephfsFsIdQuotaPutRequest).Execute()





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
	fsId := "fsId_example" // string | 
	apiCephfsFsIdQuotaPutRequest := *openapiclient.NewApiCephfsFsIdQuotaPutRequest("Path_example") // ApiCephfsFsIdQuotaPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdQuotaPut(context.Background(), fsId).ApiCephfsFsIdQuotaPutRequest(apiCephfsFsIdQuotaPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdQuotaPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdQuotaPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiCephfsFsIdQuotaPutRequest** | [**ApiCephfsFsIdQuotaPutRequest**](ApiCephfsFsIdQuotaPutRequest.md) |  | 

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


## ApiCephfsFsIdRenamePathPut

> ApiCephfsFsIdRenamePathPut(ctx, fsId).ApiCephfsFsIdRenamePathPutRequest(apiCephfsFsIdRenamePathPutRequest).Execute()





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
	fsId := "fsId_example" // string | 
	apiCephfsFsIdRenamePathPutRequest := *openapiclient.NewApiCephfsFsIdRenamePathPutRequest("DstPath_example", "SrcPath_example") // ApiCephfsFsIdRenamePathPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdRenamePathPut(context.Background(), fsId).ApiCephfsFsIdRenamePathPutRequest(apiCephfsFsIdRenamePathPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdRenamePathPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdRenamePathPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiCephfsFsIdRenamePathPutRequest** | [**ApiCephfsFsIdRenamePathPutRequest**](ApiCephfsFsIdRenamePathPutRequest.md) |  | 

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


## ApiCephfsFsIdSnapshotDelete

> ApiCephfsFsIdSnapshotDelete(ctx, fsId).Path(path).Name(name).Execute()





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
	fsId := "fsId_example" // string | 
	path := "path_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdSnapshotDelete(context.Background(), fsId).Path(path).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdSnapshotDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdSnapshotDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** |  | 
 **name** | **string** |  | 

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


## ApiCephfsFsIdSnapshotPost

> ApiCephfsFsIdSnapshotPost(ctx, fsId).ApiCephfsFsIdSnapshotPostRequest(apiCephfsFsIdSnapshotPostRequest).Execute()





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
	fsId := "fsId_example" // string | 
	apiCephfsFsIdSnapshotPostRequest := *openapiclient.NewApiCephfsFsIdSnapshotPostRequest("Path_example") // ApiCephfsFsIdSnapshotPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdSnapshotPost(context.Background(), fsId).ApiCephfsFsIdSnapshotPostRequest(apiCephfsFsIdSnapshotPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdSnapshotPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdSnapshotPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiCephfsFsIdSnapshotPostRequest** | [**ApiCephfsFsIdSnapshotPostRequest**](ApiCephfsFsIdSnapshotPostRequest.md) |  | 

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


## ApiCephfsFsIdStatfsGet

> ApiCephfsFsIdStatfsGet200Response ApiCephfsFsIdStatfsGet(ctx, fsId).Path(path).Execute()

Get Cephfs statfs of the specified path



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
	fsId := "fsId_example" // string | File System Identifier
	path := "path_example" // string | File System Path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CephfsAPI.ApiCephfsFsIdStatfsGet(context.Background(), fsId).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdStatfsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCephfsFsIdStatfsGet`: ApiCephfsFsIdStatfsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CephfsAPI.ApiCephfsFsIdStatfsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** | File System Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdStatfsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | File System Path | 

### Return type

[**ApiCephfsFsIdStatfsGet200Response**](ApiCephfsFsIdStatfsGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCephfsFsIdTreeDelete

> ApiCephfsFsIdTreeDelete(ctx, fsId).Path(path).Execute()





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
	fsId := "fsId_example" // string | 
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdTreeDelete(context.Background(), fsId).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdTreeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdTreeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** |  | 

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


## ApiCephfsFsIdTreePost

> ApiCephfsFsIdTreePost(ctx, fsId).ApiCephfsFsIdTreePostRequest(apiCephfsFsIdTreePostRequest).Execute()





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
	fsId := "fsId_example" // string | 
	apiCephfsFsIdTreePostRequest := *openapiclient.NewApiCephfsFsIdTreePostRequest("Path_example") // ApiCephfsFsIdTreePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdTreePost(context.Background(), fsId).ApiCephfsFsIdTreePostRequest(apiCephfsFsIdTreePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdTreePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdTreePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiCephfsFsIdTreePostRequest** | [**ApiCephfsFsIdTreePostRequest**](ApiCephfsFsIdTreePostRequest.md) |  | 

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


## ApiCephfsFsIdUnlinkDelete

> ApiCephfsFsIdUnlinkDelete(ctx, fsId).Path(path).Execute()





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
	fsId := "fsId_example" // string | 
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdUnlinkDelete(context.Background(), fsId).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdUnlinkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdUnlinkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** |  | 

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


## ApiCephfsFsIdWriteToFilePost

> ApiCephfsFsIdWriteToFilePost(ctx, fsId).ApiCephfsFsIdWriteToFilePostRequest(apiCephfsFsIdWriteToFilePostRequest).Execute()





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
	fsId := "fsId_example" // string | 
	apiCephfsFsIdWriteToFilePostRequest := *openapiclient.NewApiCephfsFsIdWriteToFilePostRequest("Buf_example", "Path_example") // ApiCephfsFsIdWriteToFilePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsFsIdWriteToFilePost(context.Background(), fsId).ApiCephfsFsIdWriteToFilePostRequest(apiCephfsFsIdWriteToFilePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsFsIdWriteToFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsFsIdWriteToFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiCephfsFsIdWriteToFilePostRequest** | [**ApiCephfsFsIdWriteToFilePostRequest**](ApiCephfsFsIdWriteToFilePostRequest.md) |  | 

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


## ApiCephfsGet

> ApiCephfsGet(ctx).Execute()



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
	r, err := apiClient.CephfsAPI.ApiCephfsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsGetRequest struct via the builder pattern


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


## ApiCephfsPost

> ApiCephfsPost(ctx).ApiCephfsPostRequest(apiCephfsPostRequest).Execute()



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
	apiCephfsPostRequest := *openapiclient.NewApiCephfsPostRequest("Name_example", "ServiceSpec_example") // ApiCephfsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsPost(context.Background()).ApiCephfsPostRequest(apiCephfsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCephfsPostRequest** | [**ApiCephfsPostRequest**](ApiCephfsPostRequest.md) |  | 

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


## ApiCephfsRemoveNameDelete

> ApiCephfsRemoveNameDelete(ctx, name).Execute()

Remove CephFS Volume

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
	name := "name_example" // string | File System Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsRemoveNameDelete(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsRemoveNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | File System Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsRemoveNameDeleteRequest struct via the builder pattern


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


## ApiCephfsRenamePut

> ApiCephfsRenamePut(ctx).ApiCephfsRenamePutRequest(apiCephfsRenamePutRequest).Execute()

Rename CephFS Volume

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
	apiCephfsRenamePutRequest := *openapiclient.NewApiCephfsRenamePutRequest("Name_example", "NewName_example") // ApiCephfsRenamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsAPI.ApiCephfsRenamePut(context.Background()).ApiCephfsRenamePutRequest(apiCephfsRenamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsAPI.ApiCephfsRenamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsRenamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCephfsRenamePutRequest** | [**ApiCephfsRenamePutRequest**](ApiCephfsRenamePutRequest.md) |  | 

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

