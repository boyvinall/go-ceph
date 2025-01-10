# \CephFSSnapshotScheduleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCephfsSnapshotScheduleFsGet**](CephFSSnapshotScheduleAPI.md#ApiCephfsSnapshotScheduleFsGet) | **Get** /api/cephfs/snapshot/schedule/{fs} | 
[**ApiCephfsSnapshotScheduleFsPathActivatePost**](CephFSSnapshotScheduleAPI.md#ApiCephfsSnapshotScheduleFsPathActivatePost) | **Post** /api/cephfs/snapshot/schedule/{fs}/{path}/activate | 
[**ApiCephfsSnapshotScheduleFsPathDeactivatePost**](CephFSSnapshotScheduleAPI.md#ApiCephfsSnapshotScheduleFsPathDeactivatePost) | **Post** /api/cephfs/snapshot/schedule/{fs}/{path}/deactivate | 
[**ApiCephfsSnapshotScheduleFsPathDeleteSnapshotDelete**](CephFSSnapshotScheduleAPI.md#ApiCephfsSnapshotScheduleFsPathDeleteSnapshotDelete) | **Delete** /api/cephfs/snapshot/schedule/{fs}/{path}/delete_snapshot | 
[**ApiCephfsSnapshotScheduleFsPathPut**](CephFSSnapshotScheduleAPI.md#ApiCephfsSnapshotScheduleFsPathPut) | **Put** /api/cephfs/snapshot/schedule/{fs}/{path} | 
[**ApiCephfsSnapshotSchedulePost**](CephFSSnapshotScheduleAPI.md#ApiCephfsSnapshotSchedulePost) | **Post** /api/cephfs/snapshot/schedule | 



## ApiCephfsSnapshotScheduleFsGet

> ApiCephfsSnapshotScheduleFsGet(ctx, fs).Path(path).Recursive(recursive).Execute()



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
	fs := "fs_example" // string | 
	path := "path_example" // string |  (optional)
	recursive := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSnapshotScheduleAPI.ApiCephfsSnapshotScheduleFsGet(context.Background(), fs).Path(path).Recursive(recursive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSnapshotScheduleAPI.ApiCephfsSnapshotScheduleFsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fs** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSnapshotScheduleFsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** |  | 
 **recursive** | **bool** |  | 

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


## ApiCephfsSnapshotScheduleFsPathActivatePost

> ApiCephfsSnapshotScheduleFsPathActivatePost(ctx, fs, path).ApiCephfsSnapshotScheduleFsPathActivatePostRequest(apiCephfsSnapshotScheduleFsPathActivatePostRequest).Execute()



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
	fs := "fs_example" // string | 
	path := "path_example" // string | 
	apiCephfsSnapshotScheduleFsPathActivatePostRequest := *openapiclient.NewApiCephfsSnapshotScheduleFsPathActivatePostRequest("Schedule_example", "Start_example") // ApiCephfsSnapshotScheduleFsPathActivatePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSnapshotScheduleAPI.ApiCephfsSnapshotScheduleFsPathActivatePost(context.Background(), fs, path).ApiCephfsSnapshotScheduleFsPathActivatePostRequest(apiCephfsSnapshotScheduleFsPathActivatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSnapshotScheduleAPI.ApiCephfsSnapshotScheduleFsPathActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fs** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSnapshotScheduleFsPathActivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiCephfsSnapshotScheduleFsPathActivatePostRequest** | [**ApiCephfsSnapshotScheduleFsPathActivatePostRequest**](ApiCephfsSnapshotScheduleFsPathActivatePostRequest.md) |  | 

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


## ApiCephfsSnapshotScheduleFsPathDeactivatePost

> ApiCephfsSnapshotScheduleFsPathDeactivatePost(ctx, fs, path).ApiCephfsSnapshotScheduleFsPathActivatePostRequest(apiCephfsSnapshotScheduleFsPathActivatePostRequest).Execute()



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
	fs := "fs_example" // string | 
	path := "path_example" // string | 
	apiCephfsSnapshotScheduleFsPathActivatePostRequest := *openapiclient.NewApiCephfsSnapshotScheduleFsPathActivatePostRequest("Schedule_example", "Start_example") // ApiCephfsSnapshotScheduleFsPathActivatePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSnapshotScheduleAPI.ApiCephfsSnapshotScheduleFsPathDeactivatePost(context.Background(), fs, path).ApiCephfsSnapshotScheduleFsPathActivatePostRequest(apiCephfsSnapshotScheduleFsPathActivatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSnapshotScheduleAPI.ApiCephfsSnapshotScheduleFsPathDeactivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fs** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSnapshotScheduleFsPathDeactivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiCephfsSnapshotScheduleFsPathActivatePostRequest** | [**ApiCephfsSnapshotScheduleFsPathActivatePostRequest**](ApiCephfsSnapshotScheduleFsPathActivatePostRequest.md) |  | 

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


## ApiCephfsSnapshotScheduleFsPathDeleteSnapshotDelete

> ApiCephfsSnapshotScheduleFsPathDeleteSnapshotDelete(ctx, fs, path).Schedule(schedule).Start(start).RetentionPolicy(retentionPolicy).Subvol(subvol).Group(group).Execute()



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
	fs := "fs_example" // string | 
	path := "path_example" // string | 
	schedule := "schedule_example" // string | 
	start := "start_example" // string | 
	retentionPolicy := "retentionPolicy_example" // string |  (optional)
	subvol := "subvol_example" // string |  (optional)
	group := "group_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSnapshotScheduleAPI.ApiCephfsSnapshotScheduleFsPathDeleteSnapshotDelete(context.Background(), fs, path).Schedule(schedule).Start(start).RetentionPolicy(retentionPolicy).Subvol(subvol).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSnapshotScheduleAPI.ApiCephfsSnapshotScheduleFsPathDeleteSnapshotDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fs** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSnapshotScheduleFsPathDeleteSnapshotDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **schedule** | **string** |  | 
 **start** | **string** |  | 
 **retentionPolicy** | **string** |  | 
 **subvol** | **string** |  | 
 **group** | **string** |  | 

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


## ApiCephfsSnapshotScheduleFsPathPut

> ApiCephfsSnapshotScheduleFsPathPut(ctx, fs, path).ApiCephfsSnapshotScheduleFsPathPutRequest(apiCephfsSnapshotScheduleFsPathPutRequest).Execute()



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
	fs := "fs_example" // string | 
	path := "path_example" // string | 
	apiCephfsSnapshotScheduleFsPathPutRequest := *openapiclient.NewApiCephfsSnapshotScheduleFsPathPutRequest() // ApiCephfsSnapshotScheduleFsPathPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSnapshotScheduleAPI.ApiCephfsSnapshotScheduleFsPathPut(context.Background(), fs, path).ApiCephfsSnapshotScheduleFsPathPutRequest(apiCephfsSnapshotScheduleFsPathPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSnapshotScheduleAPI.ApiCephfsSnapshotScheduleFsPathPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fs** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSnapshotScheduleFsPathPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiCephfsSnapshotScheduleFsPathPutRequest** | [**ApiCephfsSnapshotScheduleFsPathPutRequest**](ApiCephfsSnapshotScheduleFsPathPutRequest.md) |  | 

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


## ApiCephfsSnapshotSchedulePost

> ApiCephfsSnapshotSchedulePost(ctx).ApiCephfsSnapshotSchedulePostRequest(apiCephfsSnapshotSchedulePostRequest).Execute()



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
	apiCephfsSnapshotSchedulePostRequest := *openapiclient.NewApiCephfsSnapshotSchedulePostRequest("Fs_example", "Path_example", "SnapSchedule_example", "Start_example") // ApiCephfsSnapshotSchedulePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSnapshotScheduleAPI.ApiCephfsSnapshotSchedulePost(context.Background()).ApiCephfsSnapshotSchedulePostRequest(apiCephfsSnapshotSchedulePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSnapshotScheduleAPI.ApiCephfsSnapshotSchedulePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSnapshotSchedulePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCephfsSnapshotSchedulePostRequest** | [**ApiCephfsSnapshotSchedulePostRequest**](ApiCephfsSnapshotSchedulePostRequest.md) |  | 

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

