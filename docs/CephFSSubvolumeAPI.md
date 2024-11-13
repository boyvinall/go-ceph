# \CephFSSubvolumeAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCephfsSubvolumePost**](CephFSSubvolumeAPI.md#ApiCephfsSubvolumePost) | **Post** /api/cephfs/subvolume | 
[**ApiCephfsSubvolumeVolNameDelete**](CephFSSubvolumeAPI.md#ApiCephfsSubvolumeVolNameDelete) | **Delete** /api/cephfs/subvolume/{vol_name} | 
[**ApiCephfsSubvolumeVolNameExistsGet**](CephFSSubvolumeAPI.md#ApiCephfsSubvolumeVolNameExistsGet) | **Get** /api/cephfs/subvolume/{vol_name}/exists | 
[**ApiCephfsSubvolumeVolNameGet**](CephFSSubvolumeAPI.md#ApiCephfsSubvolumeVolNameGet) | **Get** /api/cephfs/subvolume/{vol_name} | 
[**ApiCephfsSubvolumeVolNameInfoGet**](CephFSSubvolumeAPI.md#ApiCephfsSubvolumeVolNameInfoGet) | **Get** /api/cephfs/subvolume/{vol_name}/info | 
[**ApiCephfsSubvolumeVolNamePut**](CephFSSubvolumeAPI.md#ApiCephfsSubvolumeVolNamePut) | **Put** /api/cephfs/subvolume/{vol_name} | 



## ApiCephfsSubvolumePost

> ApiCephfsSubvolumePost(ctx).ApiCephfsSubvolumePostRequest(apiCephfsSubvolumePostRequest).Execute()



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
	apiCephfsSubvolumePostRequest := *openapiclient.NewApiCephfsSubvolumePostRequest("SubvolName_example", "VolName_example") // ApiCephfsSubvolumePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSubvolumeAPI.ApiCephfsSubvolumePost(context.Background()).ApiCephfsSubvolumePostRequest(apiCephfsSubvolumePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSubvolumeAPI.ApiCephfsSubvolumePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCephfsSubvolumePostRequest** | [**ApiCephfsSubvolumePostRequest**](ApiCephfsSubvolumePostRequest.md) |  | 

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


## ApiCephfsSubvolumeVolNameDelete

> ApiCephfsSubvolumeVolNameDelete(ctx, volName).SubvolName(subvolName).GroupName(groupName).RetainSnapshots(retainSnapshots).Execute()



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
	volName := "volName_example" // string | 
	subvolName := "subvolName_example" // string | 
	groupName := "groupName_example" // string |  (optional)
	retainSnapshots := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSubvolumeAPI.ApiCephfsSubvolumeVolNameDelete(context.Background(), volName).SubvolName(subvolName).GroupName(groupName).RetainSnapshots(retainSnapshots).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSubvolumeAPI.ApiCephfsSubvolumeVolNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeVolNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subvolName** | **string** |  | 
 **groupName** | **string** |  | 
 **retainSnapshots** | **bool** |  | 

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


## ApiCephfsSubvolumeVolNameExistsGet

> ApiCephfsSubvolumeVolNameExistsGet(ctx, volName).GroupName(groupName).Execute()



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
	volName := "volName_example" // string | 
	groupName := "groupName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSubvolumeAPI.ApiCephfsSubvolumeVolNameExistsGet(context.Background(), volName).GroupName(groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSubvolumeAPI.ApiCephfsSubvolumeVolNameExistsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeVolNameExistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ApiCephfsSubvolumeVolNameGet

> ApiCephfsSubvolumeVolNameGet(ctx, volName).GroupName(groupName).Info(info).Execute()



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
	volName := "volName_example" // string | 
	groupName := "groupName_example" // string |  (optional)
	info := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSubvolumeAPI.ApiCephfsSubvolumeVolNameGet(context.Background(), volName).GroupName(groupName).Info(info).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSubvolumeAPI.ApiCephfsSubvolumeVolNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeVolNameGetRequest struct via the builder pattern


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


## ApiCephfsSubvolumeVolNameInfoGet

> ApiCephfsSubvolumeVolNameInfoGet(ctx, volName).SubvolName(subvolName).GroupName(groupName).Execute()



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
	volName := "volName_example" // string | 
	subvolName := "subvolName_example" // string | 
	groupName := "groupName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSubvolumeAPI.ApiCephfsSubvolumeVolNameInfoGet(context.Background(), volName).SubvolName(subvolName).GroupName(groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSubvolumeAPI.ApiCephfsSubvolumeVolNameInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeVolNameInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subvolName** | **string** |  | 
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


## ApiCephfsSubvolumeVolNamePut

> ApiCephfsSubvolumeVolNamePut(ctx, volName).ApiCephfsSubvolumeVolNamePutRequest(apiCephfsSubvolumeVolNamePutRequest).Execute()



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
	volName := "volName_example" // string | 
	apiCephfsSubvolumeVolNamePutRequest := *openapiclient.NewApiCephfsSubvolumeVolNamePutRequest(int32(123), "SubvolName_example") // ApiCephfsSubvolumeVolNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephFSSubvolumeAPI.ApiCephfsSubvolumeVolNamePut(context.Background(), volName).ApiCephfsSubvolumeVolNamePutRequest(apiCephfsSubvolumeVolNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephFSSubvolumeAPI.ApiCephfsSubvolumeVolNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeVolNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiCephfsSubvolumeVolNamePutRequest** | [**ApiCephfsSubvolumeVolNamePutRequest**](ApiCephfsSubvolumeVolNamePutRequest.md) |  | 

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

