# \CephfsSubvolumeSnapshotAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCephfsSubvolumeSnapshotPost**](CephfsSubvolumeSnapshotAPI.md#ApiCephfsSubvolumeSnapshotPost) | **Post** /api/cephfs/subvolume/snapshot | 
[**ApiCephfsSubvolumeSnapshotVolNameSubvolNameDelete**](CephfsSubvolumeSnapshotAPI.md#ApiCephfsSubvolumeSnapshotVolNameSubvolNameDelete) | **Delete** /api/cephfs/subvolume/snapshot/{vol_name}/{subvol_name} | 
[**ApiCephfsSubvolumeSnapshotVolNameSubvolNameGet**](CephfsSubvolumeSnapshotAPI.md#ApiCephfsSubvolumeSnapshotVolNameSubvolNameGet) | **Get** /api/cephfs/subvolume/snapshot/{vol_name}/{subvol_name} | 
[**ApiCephfsSubvolumeSnapshotVolNameSubvolNameInfoGet**](CephfsSubvolumeSnapshotAPI.md#ApiCephfsSubvolumeSnapshotVolNameSubvolNameInfoGet) | **Get** /api/cephfs/subvolume/snapshot/{vol_name}/{subvol_name}/info | 



## ApiCephfsSubvolumeSnapshotPost

> ApiCephfsSubvolumeSnapshotPost(ctx).ApiCephfsSubvolumeSnapshotPostRequest(apiCephfsSubvolumeSnapshotPostRequest).Execute()



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
	apiCephfsSubvolumeSnapshotPostRequest := *openapiclient.NewApiCephfsSubvolumeSnapshotPostRequest("SnapName_example", "SubvolName_example", "VolName_example") // ApiCephfsSubvolumeSnapshotPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotPost(context.Background()).ApiCephfsSubvolumeSnapshotPostRequest(apiCephfsSubvolumeSnapshotPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeSnapshotPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCephfsSubvolumeSnapshotPostRequest** | [**ApiCephfsSubvolumeSnapshotPostRequest**](ApiCephfsSubvolumeSnapshotPostRequest.md) |  | 

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


## ApiCephfsSubvolumeSnapshotVolNameSubvolNameDelete

> ApiCephfsSubvolumeSnapshotVolNameSubvolNameDelete(ctx, volName, subvolName).SnapName(snapName).GroupName(groupName).Force(force).Execute()



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
	volName := "volName_example" // string | 
	subvolName := "subvolName_example" // string | 
	snapName := "snapName_example" // string | 
	groupName := "groupName_example" // string |  (optional)
	force := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotVolNameSubvolNameDelete(context.Background(), volName, subvolName).SnapName(snapName).GroupName(groupName).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotVolNameSubvolNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volName** | **string** |  | 
**subvolName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeSnapshotVolNameSubvolNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **snapName** | **string** |  | 
 **groupName** | **string** |  | 
 **force** | **bool** |  | 

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


## ApiCephfsSubvolumeSnapshotVolNameSubvolNameGet

> ApiCephfsSubvolumeSnapshotVolNameSubvolNameGet(ctx, volName, subvolName).GroupName(groupName).Info(info).Execute()



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
	volName := "volName_example" // string | 
	subvolName := "subvolName_example" // string | 
	groupName := "groupName_example" // string |  (optional)
	info := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotVolNameSubvolNameGet(context.Background(), volName, subvolName).GroupName(groupName).Info(info).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotVolNameSubvolNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volName** | **string** |  | 
**subvolName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeSnapshotVolNameSubvolNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupName** | **string** |  | 
 **info** | **bool** |  | 

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


## ApiCephfsSubvolumeSnapshotVolNameSubvolNameInfoGet

> ApiCephfsSubvolumeSnapshotVolNameSubvolNameInfoGet(ctx, volName, subvolName).SnapName(snapName).GroupName(groupName).Execute()



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
	volName := "volName_example" // string | 
	subvolName := "subvolName_example" // string | 
	snapName := "snapName_example" // string | 
	groupName := "groupName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotVolNameSubvolNameInfoGet(context.Background(), volName, subvolName).SnapName(snapName).GroupName(groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotVolNameSubvolNameInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volName** | **string** |  | 
**subvolName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeSnapshotVolNameSubvolNameInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **snapName** | **string** |  | 
 **groupName** | **string** |  | 

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

