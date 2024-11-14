# \CephfsSubvolumeGroupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCephfsSubvolumeGroupPost**](CephfsSubvolumeGroupAPI.md#ApiCephfsSubvolumeGroupPost) | **Post** /api/cephfs/subvolume/group | 
[**ApiCephfsSubvolumeGroupVolNameDelete**](CephfsSubvolumeGroupAPI.md#ApiCephfsSubvolumeGroupVolNameDelete) | **Delete** /api/cephfs/subvolume/group/{vol_name} | 
[**ApiCephfsSubvolumeGroupVolNameGet**](CephfsSubvolumeGroupAPI.md#ApiCephfsSubvolumeGroupVolNameGet) | **Get** /api/cephfs/subvolume/group/{vol_name} | 
[**ApiCephfsSubvolumeGroupVolNameInfoGet**](CephfsSubvolumeGroupAPI.md#ApiCephfsSubvolumeGroupVolNameInfoGet) | **Get** /api/cephfs/subvolume/group/{vol_name}/info | 
[**ApiCephfsSubvolumeGroupVolNamePut**](CephfsSubvolumeGroupAPI.md#ApiCephfsSubvolumeGroupVolNamePut) | **Put** /api/cephfs/subvolume/group/{vol_name} | 



## ApiCephfsSubvolumeGroupPost

> ApiCephfsSubvolumeGroupPost(ctx).ApiCephfsSubvolumeGroupPostRequest(apiCephfsSubvolumeGroupPostRequest).Execute()



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
	apiCephfsSubvolumeGroupPostRequest := *openapiclient.NewApiCephfsSubvolumeGroupPostRequest("GroupName_example", "VolName_example") // ApiCephfsSubvolumeGroupPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsSubvolumeGroupAPI.ApiCephfsSubvolumeGroupPost(context.Background()).ApiCephfsSubvolumeGroupPostRequest(apiCephfsSubvolumeGroupPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsSubvolumeGroupAPI.ApiCephfsSubvolumeGroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCephfsSubvolumeGroupPostRequest** | [**ApiCephfsSubvolumeGroupPostRequest**](ApiCephfsSubvolumeGroupPostRequest.md) |  | 

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


## ApiCephfsSubvolumeGroupVolNameDelete

> ApiCephfsSubvolumeGroupVolNameDelete(ctx, volName).GroupName(groupName).Execute()



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
	groupName := "groupName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsSubvolumeGroupAPI.ApiCephfsSubvolumeGroupVolNameDelete(context.Background(), volName).GroupName(groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsSubvolumeGroupAPI.ApiCephfsSubvolumeGroupVolNameDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeGroupVolNameDeleteRequest struct via the builder pattern


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


## ApiCephfsSubvolumeGroupVolNameGet

> ApiCephfsSubvolumeGroupVolNameGet(ctx, volName).Info(info).Execute()



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
	info := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsSubvolumeGroupAPI.ApiCephfsSubvolumeGroupVolNameGet(context.Background(), volName).Info(info).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsSubvolumeGroupAPI.ApiCephfsSubvolumeGroupVolNameGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeGroupVolNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ApiCephfsSubvolumeGroupVolNameInfoGet

> ApiCephfsSubvolumeGroupVolNameInfoGet(ctx, volName).GroupName(groupName).Execute()



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
	groupName := "groupName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsSubvolumeGroupAPI.ApiCephfsSubvolumeGroupVolNameInfoGet(context.Background(), volName).GroupName(groupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsSubvolumeGroupAPI.ApiCephfsSubvolumeGroupVolNameInfoGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeGroupVolNameInfoGetRequest struct via the builder pattern


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


## ApiCephfsSubvolumeGroupVolNamePut

> ApiCephfsSubvolumeGroupVolNamePut(ctx, volName).ApiCephfsSubvolumeGroupVolNamePutRequest(apiCephfsSubvolumeGroupVolNamePutRequest).Execute()



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
	apiCephfsSubvolumeGroupVolNamePutRequest := *openapiclient.NewApiCephfsSubvolumeGroupVolNamePutRequest("GroupName_example", int32(123)) // ApiCephfsSubvolumeGroupVolNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsSubvolumeGroupAPI.ApiCephfsSubvolumeGroupVolNamePut(context.Background(), volName).ApiCephfsSubvolumeGroupVolNamePutRequest(apiCephfsSubvolumeGroupVolNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsSubvolumeGroupAPI.ApiCephfsSubvolumeGroupVolNamePut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeGroupVolNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiCephfsSubvolumeGroupVolNamePutRequest** | [**ApiCephfsSubvolumeGroupVolNamePutRequest**](ApiCephfsSubvolumeGroupVolNamePutRequest.md) |  | 

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

